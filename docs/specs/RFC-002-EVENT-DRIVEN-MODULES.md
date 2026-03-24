# RFC: Event-Driven Module Loading

**Status:** Implemented
**Created:** 2026-01-15
**Authors:** Host UK Engineering

---

## Abstract

The Event-Driven Module Loading system enables lazy instantiation of modules based on lifecycle events. Instead of eagerly booting all modules at application startup, modules declare interest in specific events via static `$listens` arrays. The module is only instantiated when its events fire.

This provides:
- Faster boot times (only load what's needed)
- Context-aware loading (CLI gets CLI modules, web gets web modules)
- Clean separation between infrastructure and modules
- Testable event-based architecture

---

## Core Components

### Architecture Overview

```
┌─────────────────────────────────────────────────────────────────┐
│                     Application Bootstrap                        │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│   LifecycleEventProvider                                        │
│   └── ModuleRegistry                                            │
│       └── ModuleScanner (reads $listens via reflection)         │
│           └── LazyModuleListener (defers instantiation)         │
│                                                                  │
├─────────────────────────────────────────────────────────────────┤
│                     Frontages (fire events)                      │
├─────────────────────────────────────────────────────────────────┤
│   Front/Web/Boot ──────────▶ WebRoutesRegistering               │
│   Front/Admin/Boot ────────▶ AdminPanelBooting                  │
│   Front/Api/Boot ──────────▶ ApiRoutesRegistering               │
│   Front/Cli/Boot ──────────▶ ConsoleBooting                     │
│   Mcp/Server ──────────────▶ McpToolsRegistering                │
│   Queue Worker ────────────▶ QueueWorkerBooting                 │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

### ModuleScanner

Reads Boot.php files and extracts `$listens` arrays via reflection without instantiating the modules.

```php
namespace Core;

class ModuleScanner
{
    public function scan(array $paths): array
    {
        // Returns: [EventClass => [ModuleClass => 'methodName']]
    }

    public function extractListens(string $class): array
    {
        // Uses ReflectionClass to read static $listens property
        // Returns empty array if missing/invalid
    }
}
```

### ModuleRegistry

Wires up lazy listeners for all scanned modules.

```php
namespace Core;

class ModuleRegistry
{
    public function register(array $paths): void
    {
        $mappings = $this->scanner->scan($paths);

        foreach ($mappings as $event => $listeners) {
            foreach ($listeners as $moduleClass => $method) {
                Event::listen($event, new LazyModuleListener($moduleClass, $method));
            }
        }
    }
}
```

### LazyModuleListener

Defers module instantiation until the event fires.

```php
namespace Core;

class LazyModuleListener
{
    public function __invoke(object $event): void
    {
        $module = $this->resolveModule();
        $module->{$this->method}($event);
    }

    private function resolveModule(): object
    {
        // Handles ServiceProvider subclasses correctly
        if (is_subclass_of($this->moduleClass, ServiceProvider::class)) {
            return app()->resolveProvider($this->moduleClass);
        }
        return app()->make($this->moduleClass);
    }
}
```

### LifecycleEvent Base Class

Events collect requests from modules without immediately applying them.

```php
namespace Core\Events;

abstract class LifecycleEvent
{
    public function routes(callable $callback): void;
    public function views(string $namespace, string $path): void;
    public function livewire(string $alias, string $class): void;
    public function command(string $class): void;
    public function middleware(string $alias, string $class): void;
    public function navigation(array $item): void;
    public function translations(string $namespace, string $path): void;
    public function policy(string $model, string $policy): void;

    // Getters for processing
    public function routeRequests(): array;
    public function viewRequests(): array;
    // etc.
}
```

---

## Available Events

| Event | Context | Fired By |
|-------|---------|----------|
| `AdminPanelBooting` | Admin panel requests | `Front\Admin\Boot` |
| `WebRoutesRegistering` | Web requests | `Front\Web\Boot` |
| `ApiRoutesRegistering` | API requests | `Front\Api\Boot` |
| `ConsoleBooting` | CLI commands | `Front\Cli\Boot` |
| `McpToolsRegistering` | MCP server | Mcp module |
| `QueueWorkerBooting` | Queue workers | `LifecycleEventProvider` |
| `FrameworkBooted` | All contexts (post-boot) | `LifecycleEventProvider` |
| `MediaRequested` | Media serving | Core media handler |
| `SearchRequested` | Search operations | Core search handler |
| `MailSending` | Mail dispatch | Core mail handler |

---

## Module Implementation

### Declaring Listeners

Modules declare interest in events via the static `$listens` property:

```php
namespace Mod\Commerce;

use Core\Events\AdminPanelBooting;
use Core\Events\ConsoleBooting;
use Core\Events\WebRoutesRegistering;

class Boot extends ServiceProvider
{
    public static array $listens = [
        AdminPanelBooting::class => 'onAdminPanel',
        WebRoutesRegistering::class => 'onWebRoutes',
        ConsoleBooting::class => 'onConsole',
    ];

    public function onAdminPanel(AdminPanelBooting $event): void
    {
        $event->views('commerce', __DIR__.'/View/Blade');
        $event->livewire('commerce.checkout', Components\Checkout::class);
        $event->routes(fn () => require __DIR__.'/Routes/admin.php');
    }

    public function onWebRoutes(WebRoutesRegistering $event): void
    {
        $event->views('commerce', __DIR__.'/View/Blade');
        $event->routes(fn () => require __DIR__.'/Routes/web.php');
    }

    public function onConsole(ConsoleBooting $event): void
    {
        $event->command(Commands\ProcessPayments::class);
        $event->command(Commands\SyncSubscriptions::class);
    }
}
```

### What Stays in boot()

Some registrations must remain in the traditional `boot()` method:

| Registration | Reason |
|--------------|--------|
| `loadMigrationsFrom()` | Needed early for `artisan migrate` |
| `AdminMenuRegistry->register()` | Uses interface pattern (AdminMenuProvider) |
| Laravel event listeners | Standard Laravel events, not lifecycle events |

```php
public function boot(): void
{
    $this->loadMigrationsFrom(__DIR__.'/Migrations');

    // Interface-based registration
    app(AdminMenuRegistry::class)->register($this);

    // Standard Laravel events (not lifecycle events)
    Event::listen(OrderPlaced::class, SendOrderConfirmation::class);
}
```

---

## Request Processing

Frontages fire events and process collected requests:

```php
// In Front/Web/Boot
public static function fireWebRoutes(): void
{
    $event = new WebRoutesRegistering;
    event($event);

    // Process view namespaces
    foreach ($event->viewRequests() as [$namespace, $path]) {
        view()->addNamespace($namespace, $path);
    }

    // Process Livewire components
    foreach ($event->livewireRequests() as [$alias, $class]) {
        Livewire::component($alias, $class);
    }

    // Process routes with web middleware
    foreach ($event->routeRequests() as $callback) {
        Route::middleware('web')->group($callback);
    }
}
```

This "collect then process" pattern ensures:
1. Modules cannot directly mutate infrastructure
2. Core validates and controls registration order
3. Easy to add cross-cutting concerns (logging, validation)

---

## Testing

### Unit Tests

Test ModuleScanner reflection without Laravel app:

```php
it('extracts $listens from a class with public static property', function () {
    $scanner = new ModuleScanner;
    $listens = $scanner->extractListens(ModuleWithListens::class);

    expect($listens)->toBe([
        'SomeEvent' => 'handleSomeEvent',
    ]);
});

it('returns empty array when $listens is not public', function () {
    $scanner = new ModuleScanner;
    $listens = $scanner->extractListens(ModuleWithPrivateListens::class);

    expect($listens)->toBe([]);
});
```

### Integration Tests

Test real module scanning with Laravel app:

```php
it('scans the Mod directory and finds modules', function () {
    $scanner = new ModuleScanner;
    $result = $scanner->scan([app_path('Mod')]);

    expect($result)->toHaveKey(AdminPanelBooting::class);
    expect($result)->toHaveKey(WebRoutesRegistering::class);
});
```

---

## Performance

### Lazy Loading Benefits

| Context | Modules Loaded | Without Lazy Loading |
|---------|----------------|----------------------|
| Web request | 6-8 modules | All 16+ modules |
| Admin request | 10-12 modules | All 16+ modules |
| CLI command | 4-6 modules | All 16+ modules |
| API request | 3-5 modules | All 16+ modules |

### Memory Impact

Modules not needed for the current context are never instantiated:
- No class autoloading
- No service binding
- No config merging
- No route registration

---

## Files

### Core Infrastructure

| File | Purpose |
|------|---------|
| `Core/ModuleScanner.php` | Scans Boot.php files for $listens |
| `Core/ModuleRegistry.php` | Wires up lazy listeners |
| `Core/LazyModuleListener.php` | Defers module instantiation |
| `Core/LifecycleEventProvider.php` | Orchestrates scanning and events |
| `Core/Events/LifecycleEvent.php` | Base class for all lifecycle events |

### Events

| File | Purpose |
|------|---------|
| `Core/Events/AdminPanelBooting.php` | Admin panel context |
| `Core/Events/WebRoutesRegistering.php` | Web context |
| `Core/Events/ApiRoutesRegistering.php` | API context |
| `Core/Events/ConsoleBooting.php` | CLI context |
| `Core/Events/McpToolsRegistering.php` | MCP server context |
| `Core/Events/QueueWorkerBooting.php` | Queue worker context |
| `Core/Events/FrameworkBooted.php` | Post-boot event |

### Frontages

| File | Purpose |
|------|---------|
| `Core/Front/Web/Boot.php` | Fires WebRoutesRegistering |
| `Core/Front/Admin/Boot.php` | Fires AdminPanelBooting |
| `Core/Front/Api/Boot.php` | Fires ApiRoutesRegistering |
| `Core/Front/Cli/Boot.php` | Fires ConsoleBooting |

---

## Migration Guide

### Before (Legacy)

```php
class Boot extends ServiceProvider
{
    public function boot(): void
    {
        $this->registerRoutes();
        $this->registerViews();
        $this->registerLivewireComponents();
        $this->registerCommands();
    }

    private function registerRoutes(): void
    {
        Route::middleware('web')->group(__DIR__.'/Routes/web.php');
    }

    private function registerViews(): void
    {
        $this->loadViewsFrom(__DIR__.'/View/Blade', 'mymodule');
    }
    // etc.
}
```

### After (Event-Driven)

```php
class Boot extends ServiceProvider
{
    public static array $listens = [
        WebRoutesRegistering::class => 'onWebRoutes',
        ConsoleBooting::class => 'onConsole',
    ];

    public function boot(): void
    {
        $this->loadMigrationsFrom(__DIR__.'/Migrations');
    }

    public function onWebRoutes(WebRoutesRegistering $event): void
    {
        $event->views('mymodule', __DIR__.'/View/Blade');
        $event->routes(fn () => require __DIR__.'/Routes/web.php');
    }

    public function onConsole(ConsoleBooting $event): void
    {
        $event->command(Commands\MyCommand::class);
    }
}
```

---

## Future Considerations

1. **Event Caching**: Cache scanned mappings in production for faster boot
2. **Module Dependencies**: Declare dependencies between modules for ordered loading
3. **Hot Module Reloading**: In development, detect changes and re-scan
4. **Event Priorities**: Allow modules to specify listener priority
