# flux:field

Form field container with labels, descriptions, and validation messaging.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `variant` | string | block | `block`, `inline` |

## Child Components

### flux:label

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `badge` | string | - | Badge text (e.g., "Required", "Optional") |

**Slots:** `default` (label text), `trailing` (end text)

### flux:description

Helper text for form inputs.

**Slots:** `default` (description text)

### flux:error

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `name` | string | - | Field name for validation errors |
| `message` | string | - | Custom error message |
| `bag` | string | default | Error bag reference |
| `icon` | string | exclamation-triangle | Icon displayed (false to hide) |

**Slots:** `default` (custom error content)

### flux:fieldset

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `legend` | string | - | Fieldset heading |
| `description` | string | - | Fieldset description |

### flux:legend

Heading for fieldset groups.

**Slots:** `default` (heading text)

## Basic Usage

```blade
<flux:field>
    <flux:label>Email</flux:label>
    <flux:input type="email" wire:model="email" />
    <flux:error name="email" />
</flux:field>
```

## With Description

```blade
<flux:field>
    <flux:label>Password</flux:label>
    <flux:description>Must be at least 8 characters.</flux:description>
    <flux:input type="password" wire:model="password" />
    <flux:error name="password" />
</flux:field>
```

## Trailing Description

```blade
<flux:field>
    <flux:label>Username</flux:label>
    <flux:input wire:model="username" />
    <flux:description>This will be your public display name.</flux:description>
</flux:field>
```

## With Badge

```blade
<flux:field>
    <flux:label badge="Required">Email</flux:label>
    <flux:input type="email" wire:model="email" />
</flux:field>
```

## Shorthand (Recommended)

Most input components accept `label` and `description` directly:

```blade
<flux:input
    label="Email"
    description="We'll never share your email."
    wire:model="email"
/>
```

## Fieldset Grouping

```blade
<flux:fieldset legend="Personal Information" description="Tell us about yourself.">
    <flux:input label="First name" wire:model="firstName" />
    <flux:input label="Last name" wire:model="lastName" />
</flux:fieldset>
```

## Split Layout

```blade
<div class="grid grid-cols-2 gap-4">
    <flux:input label="First name" wire:model="firstName" />
    <flux:input label="Last name" wire:model="lastName" />
</div>
```

## Custom Error Message

```blade
<flux:field>
    <flux:label>Email</flux:label>
    <flux:input type="email" wire:model="email" />
    <flux:error name="email" message="Please enter a valid email address." />
</flux:field>
```
