# flux:modal

Overlay content layer for confirmations, alerts, and forms.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `name` | string | required | Unique modal identifier |
| `flyout` | boolean | false | Opens as flyout panel |
| `variant` | string | default | `default`, `floating`, `bare` |
| `position` | string | right | Flyout direction: `right`, `left`, `bottom` |
| `dismissible` | boolean | true | Allow closing by clicking outside |
| `closable` | boolean | true | Show close button |
| `wire:model` | string | - | Bind to Livewire property |

**Slots:** `default` (content), `footer` (for floating variant)

**Events:**
- `@close` / `wire:close` - Triggered when modal closes
- `@cancel` / `wire:cancel` - Triggered when dismissed

## Child Components

### flux:modal.trigger

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `name` | string | - | Must match target modal name |
| `shortcut` | string | - | Keyboard shortcut (e.g., cmd.k) |

### flux:modal.close

Closes parent modal when activated.

## Control Methods

### From Livewire (PHP)

```php
Flux::modal('confirm')->show();
Flux::modal('confirm')->close();
Flux::modals()->close(); // Close all

// Component-scoped
$this->modal('confirm')->show();
```

### From Alpine.js

```javascript
$flux.modal('confirm').show()
$flux.modal('confirm').close()
$flux.modals().close()
```

### From JavaScript

```javascript
Flux.modal('confirm').show()
Flux.modal('confirm').close()
```

## Basic Usage

```blade
<flux:modal.trigger name="confirm">
    <flux:button>Open Modal</flux:button>
</flux:modal.trigger>

<flux:modal name="confirm" class="md:w-96">
    <flux:heading size="lg">Confirm Action</flux:heading>
    <flux:text class="mt-2">Are you sure you want to proceed?</flux:text>

    <div class="mt-6 flex gap-2">
        <flux:modal.close>
            <flux:button variant="ghost">Cancel</flux:button>
        </flux:modal.close>
        <flux:button variant="primary">Confirm</flux:button>
    </div>
</flux:modal>
```

## With wire:model

```blade
<flux:modal wire:model.self="showModal" name="edit">
    {{-- Content --}}
</flux:modal>
```

```php
public bool $showModal = false;

public function openEdit()
{
    $this->showModal = true;
}
```

## Non-dismissible

```blade
<flux:modal name="required" :dismissible="false" :closable="false">
    {{-- Must complete action to close --}}
</flux:modal>
```

## Flyout Panel

```blade
<flux:modal name="sidebar" flyout>
    {{-- Slides in from right --}}
</flux:modal>

<flux:modal name="left-panel" flyout position="left">
    {{-- Slides in from left --}}
</flux:modal>

<flux:modal name="drawer" flyout position="bottom">
    {{-- Slides up from bottom --}}
</flux:modal>
```

## Floating Flyout (Pro)

```blade
<flux:modal name="float" flyout variant="floating">
    <flux:heading>Floating Panel</flux:heading>
    <x-slot name="footer">
        <flux:button variant="primary">Save</flux:button>
    </x-slot>
</flux:modal>
```

## Bare Variant (for Command Palette)

```blade
<flux:modal name="command" variant="bare">
    <flux:command class="w-full max-w-lg">
        {{-- Command palette content --}}
    </flux:command>
</flux:modal>
```

## With Close Listener

```blade
<flux:modal name="form" @close="$wire.resetForm()">
    {{-- Form content --}}
</flux:modal>
```

## Keyboard Shortcut Trigger

```blade
<flux:modal.trigger name="search" shortcut="cmd.k">
    <flux:button icon="magnifying-glass">Search</flux:button>
</flux:modal.trigger>
```

## Dynamic Modal Names (in loops)

```blade
@foreach ($items as $item)
    <flux:modal.trigger :name="'edit-' . $item->id">
        <flux:button size="sm">Edit</flux:button>
    </flux:modal.trigger>

    <flux:modal :name="'edit-' . $item->id">
        {{-- Edit form for $item --}}
    </flux:modal>
@endforeach
```
