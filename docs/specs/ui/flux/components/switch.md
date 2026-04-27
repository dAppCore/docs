# flux:switch

Toggle control for binary options with auto-save capability.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Binds state to Livewire property |
| `label` | string | - | Label text (wraps in field) |
| `description` | string | - | Help text |
| `align` | string | right\|start | `right\|start`, `left\|end` |
| `disabled` | boolean | false | Prevents interaction |

**Attributes:**
- `data-flux-switch` - Root element
- `data-checked` - Applied when on

## Basic Usage

```blade
<flux:switch wire:model="notifications" />
```

## With Label

```blade
<flux:field variant="inline">
    <flux:label>Enable notifications</flux:label>
    <flux:switch wire:model.live="notifications" />
</flux:field>
```

## With Description

```blade
<flux:field variant="inline">
    <flux:label>Dark mode</flux:label>
    <flux:description>Use dark theme throughout the app.</flux:description>
    <flux:switch wire:model.live="darkMode" />
</flux:field>
```

## Left Aligned

```blade
<flux:field variant="inline">
    <flux:switch wire:model.live="enabled" align="left" />
    <flux:label>Enable feature</flux:label>
</flux:field>
```

## Settings Group

```blade
<flux:fieldset legend="Notifications" description="Manage your notification preferences.">
    <flux:field variant="inline">
        <flux:label>Email notifications</flux:label>
        <flux:switch wire:model.live="emailNotifications" />
    </flux:field>

    <flux:separator />

    <flux:field variant="inline">
        <flux:label>Push notifications</flux:label>
        <flux:switch wire:model.live="pushNotifications" />
    </flux:field>

    <flux:separator />

    <flux:field variant="inline">
        <flux:label>SMS notifications</flux:label>
        <flux:switch wire:model.live="smsNotifications" />
    </flux:field>
</flux:fieldset>
```

## Disabled State

```blade
<flux:field variant="inline">
    <flux:label>Premium feature</flux:label>
    <flux:description>Upgrade to access this feature.</flux:description>
    <flux:switch disabled />
</flux:field>
```

## Auto-save Pattern

Switches are ideal for settings that save immediately:

```blade
<flux:switch wire:model.live="autoSave" />
```

```php
public bool $autoSave = true;

public function updatedAutoSave($value)
{
    auth()->user()->update(['auto_save' => $value]);
    $this->dispatch('notify', 'Settings saved');
}
```
