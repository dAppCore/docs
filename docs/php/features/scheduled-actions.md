# Scheduled Actions

Declare schedules directly on Action classes using PHP attributes. No manual `routes/console.php` entries needed — the framework discovers, persists, and executes them automatically.

## Overview

Scheduled Actions combine the [Actions pattern](actions.md) with PHP 8.1 attributes to create a database-backed scheduling system. Actions declare their default schedule via `#[Scheduled]`, a sync command persists them to a `scheduled_actions` table, and a service provider wires them into Laravel's scheduler at runtime.

```
artisan schedule:sync         artisan schedule:run
        │                             │
  ScheduledActionScanner        ScheduleServiceProvider
        │                             │
  Discovers #[Scheduled]        Reads scheduled_actions
  attributes via reflection     table, wires into Schedule
        │                             │
  Upserts scheduled_actions     Calls Action::run() at
  table rows                    configured frequency
```

## Basic Usage

Add the `#[Scheduled]` attribute to any Action class:

```php
<?php

declare(strict_types=1);

namespace Mod\Social\Actions;

use Core\Actions\Action;
use Core\Actions\Scheduled;

#[Scheduled(frequency: 'dailyAt:09:00', timezone: 'Europe/London')]
class PublishDiscordDigest
{
    use Action;

    public function handle(): void
    {
        // Gather yesterday's commits, summarise, post to Discord
    }
}
```

No Boot registration needed. No `routes/console.php` entry. The scanner discovers it, `schedule:sync` persists it, and the scheduler runs it.

## The `#[Scheduled]` Attribute

```php
#[Attribute(Attribute::TARGET_CLASS)]
class Scheduled
{
    public function __construct(
        public string $frequency,
        public ?string $timezone = null,
        public bool $withoutOverlapping = true,
        public bool $runInBackground = true,
    ) {}
}
```

### Frequency Strings

The `frequency` string maps directly to Laravel Schedule methods. Arguments are colon-separated, with multiple arguments comma-separated:

| Frequency String | Laravel Equivalent |
|---|---|
| `everyMinute` | `->everyMinute()` |
| `hourly` | `->hourly()` |
| `dailyAt:09:00` | `->dailyAt('09:00')` |
| `weeklyOn:1,09:00` | `->weeklyOn(1, '09:00')` |
| `monthlyOn:1,00:00` | `->monthlyOn(1, '00:00')` |
| `cron:*/5 * * * *` | `->cron('*/5 * * * *')` |

Numeric arguments are automatically cast to integers, so `weeklyOn:1,09:00` correctly passes `(int) 1` and `'09:00'`.

## Syncing Schedules

The `schedule:sync` command scans for `#[Scheduled]` attributes and persists them to the database:

```bash
php artisan schedule:sync
# Schedule sync complete: 3 added, 1 disabled, 12 unchanged.
```

### Behaviour

- **New classes** are inserted with their attribute defaults
- **Existing rows** are preserved (manual edits to frequency are not overwritten)
- **Removed classes** are disabled (`is_enabled = false`), not deleted
- **Idempotent** — safe to run on every deploy

Run this command as part of your deployment pipeline, after migrations.

### Scan Paths

By default, the scanner checks:

- `app/Core`, `app/Mod`, `app/Website` (application code)
- `src/Core`, `src/Mod` (framework code)

Override with the `core.scheduled_action_paths` config key:

```php
// config/core.php
'scheduled_action_paths' => [
    app_path('Core'),
    app_path('Mod'),
],
```

## The `ScheduledAction` Model

Each discovered action is persisted as a `ScheduledAction` row:

| Column | Type | Description |
|---|---|---|
| `action_class` | `string` (unique) | Fully qualified class name |
| `frequency` | `string` | Schedule frequency string |
| `timezone` | `string` (nullable) | Timezone override |
| `without_overlapping` | `boolean` | Prevent concurrent runs |
| `run_in_background` | `boolean` | Run in background process |
| `is_enabled` | `boolean` | Toggle on/off |
| `last_run_at` | `timestamp` (nullable) | Last execution time |
| `next_run_at` | `timestamp` (nullable) | Computed next run |

### Querying

```php
use Core\Actions\ScheduledAction;

// All enabled actions
$active = ScheduledAction::enabled()->get();

// Check last run
$action = ScheduledAction::where('action_class', MyAction::class)->first();
echo $action->last_run_at?->diffForHumans(); // "2 hours ago"

// Parse frequency
$action->frequencyMethod(); // 'dailyAt'
$action->frequencyArgs();   // ['09:00']
```

## Runtime Execution

The `ScheduleServiceProvider` boots in console context and wires all enabled rows into Laravel's scheduler. It validates each action before registering:

- **Namespace allowlist** — only classes in `App\`, `Core\`, or `Mod\` namespaces are accepted
- **Action trait check** — the class must use the `Core\Actions\Action` trait
- **Frequency allowlist** — only recognised Laravel Schedule methods are permitted

After each run, `last_run_at` is updated automatically.

## Admin Control

The `scheduled_actions` table is designed for admin visibility. You can:

- **Disable** an action by setting `is_enabled = false` — it will not be re-enabled by subsequent syncs
- **Change frequency** by editing the `frequency` column — manual edits are preserved across syncs
- **Monitor** via `last_run_at` — see when each action last executed

## Migration Strategy

- Existing `routes/console.php` commands remain untouched
- New scheduled work uses `#[Scheduled]` actions
- Existing commands can be migrated to actions gradually at natural touch points

## Examples

### Every-minute health check

```php
#[Scheduled(frequency: 'everyMinute', withoutOverlapping: true)]
class CheckServiceHealth
{
    use Action;

    public function handle(): void
    {
        // Ping upstream services, alert on failure
    }
}
```

### Weekly report with timezone

```php
#[Scheduled(frequency: 'weeklyOn:1,09:00', timezone: 'Europe/London')]
class SendWeeklyReport
{
    use Action;

    public function handle(): void
    {
        // Compile and email weekly metrics
    }
}
```

### Cron expression

```php
#[Scheduled(frequency: 'cron:0 */6 * * *')]
class SyncExternalData
{
    use Action;

    public function handle(): void
    {
        // Pull data from external API every 6 hours
    }
}
```

## Testing

```php
use Core\Actions\Scheduled;
use Core\Actions\ScheduledAction;
use Core\Actions\ScheduledActionScanner;

it('discovers scheduled actions', function () {
    $scanner = new ScheduledActionScanner();
    $results = $scanner->scan([app_path('Mod')]);

    expect($results)->not->toBeEmpty();
    expect(array_values($results)[0])->toBeInstanceOf(Scheduled::class);
});

it('syncs scheduled actions to database', function () {
    $this->artisan('schedule:sync')->assertSuccessful();

    expect(ScheduledAction::enabled()->count())->toBeGreaterThan(0);
});
```

## Learn More

- [Actions Pattern](actions.md)
- [Module System](/php/framework/modules)
- [Lifecycle Events](/php/framework/events)
