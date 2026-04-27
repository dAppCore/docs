# flux:checkbox

Selection of one or multiple options with groups, descriptions, and visual variants.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Binds to Livewire property |
| `label` | string | - | Label text |
| `description` | string | - | Help text |
| `value` | string | - | Value when checked in group |
| `checked` | boolean | false | Default checked state |
| `indeterminate` | boolean | false | Partial selection (dash icon) |
| `disabled` | boolean | false | Prevents interaction |
| `invalid` | boolean | false | Error styling |

## Child Components

### flux:checkbox.group

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Array of selected values |
| `label` | string | - | Group heading |
| `description` | string | - | Help text |
| `variant` | string | default | `default`, `cards`, `pills`, `buttons` |
| `disabled` | boolean | false | Disable all |
| `invalid` | boolean | false | Error all |

### flux:checkbox.all

Master checkbox controlling all group members.

## Basic

```blade
<flux:checkbox wire:model="terms" label="I agree to terms" />
```

## Group

```blade
<flux:checkbox.group wire:model="notifications" label="Notifications">
    <flux:checkbox label="Push" value="push" checked />
    <flux:checkbox label="Email" value="email" />
    <flux:checkbox label="SMS" value="sms" />
</flux:checkbox.group>
```

## Variants (Pro)

```blade
{{-- Cards --}}
<flux:checkbox.group wire:model="plan" variant="cards">
    <flux:checkbox value="basic" label="Basic" description="For individuals" />
    <flux:checkbox value="pro" label="Pro" description="For teams" />
</flux:checkbox.group>

{{-- Pills --}}
<flux:checkbox.group wire:model="tags" variant="pills">
    <flux:checkbox value="php">PHP</flux:checkbox>
    <flux:checkbox value="laravel">Laravel</flux:checkbox>
</flux:checkbox.group>

{{-- Buttons --}}
<flux:checkbox.group wire:model="format" variant="buttons">
    <flux:checkbox value="bold" icon="bold" />
    <flux:checkbox value="italic" icon="italic" />
</flux:checkbox.group>
```

## Check All

```blade
<flux:checkbox.group>
    <flux:checkbox.all label="Select all" />
    <flux:checkbox value="1" label="Option 1" />
    <flux:checkbox value="2" label="Option 2" />
</flux:checkbox.group>
```
