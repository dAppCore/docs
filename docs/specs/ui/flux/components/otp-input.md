# flux:otp

One-time password input with individual fields, masking, and auto-submit.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Binds value to Livewire property |
| `value` | string | - | Current value as string |
| `length` | number | - | Number of input fields |
| `mode` | string | numeric | `numeric`, `alphanumeric`, `alpha` |
| `private` | boolean | false | Masks input values |
| `submit` | string | - | `auto` submits when filled |
| `autocomplete` | string | one-time-code | Autocomplete attribute |
| `label` | string | - | Field label |
| `label:sr-only` | boolean | false | Screen-reader only label |
| `error:icon` | boolean | - | Error icon display |
| `error:class` | string | - | Custom error classes |

## Child Components

### flux:otp.input

Individual input field within the group.

### flux:otp.separator

Visual separator between field groups.

### flux:otp.group

Container grouping multiple input fields.

## Basic Usage

```blade
<flux:otp wire:model="code" length="6" />
```

## Auto-submit

Automatically submits form when all fields are filled:

```blade
<form wire:submit="verify">
    <flux:otp wire:model="code" length="6" submit="auto" />
</form>
```

## Alphanumeric Mode

For license keys or mixed codes:

```blade
<flux:otp wire:model="licenseKey" length="10" mode="alphanumeric" />
```

## Private Mode (PIN codes)

Masks input for sensitive values:

```blade
<flux:otp wire:model="pin" length="4" private />
```

## With Label

```blade
<flux:otp wire:model="code" length="6" label="Verification code" />
```

## Custom Separators

Group inputs with visual dividers:

```blade
<flux:otp wire:model="code">
    <flux:otp.input />
    <flux:otp.input />
    <flux:otp.input />
    <flux:otp.separator />
    <flux:otp.input />
    <flux:otp.input />
    <flux:otp.input />
</flux:otp>
```

## Grouped Fields

Organise inputs into logical groups:

```blade
<flux:otp wire:model="code">
    <flux:otp.group>
        <flux:otp.input />
        <flux:otp.input />
        <flux:otp.input />
    </flux:otp.group>
    <flux:otp.separator />
    <flux:otp.group>
        <flux:otp.input />
        <flux:otp.input />
        <flux:otp.input />
    </flux:otp.group>
</flux:otp>
```

## Complete Example

```blade
<div class="space-y-4">
    <flux:heading size="lg">Enter verification code</flux:heading>
    <flux:text>We sent a 6-digit code to your email.</flux:text>

    <form wire:submit="verify">
        <flux:otp wire:model="code" length="6" submit="auto" />
        <flux:error name="code" />

        <flux:button type="submit" class="mt-4 w-full">
            Verify
        </flux:button>
    </form>

    <flux:text size="sm" class="text-center">
        Didn't receive a code?
        <flux:link wire:click="resend">Resend</flux:link>
    </flux:text>
</div>
```
