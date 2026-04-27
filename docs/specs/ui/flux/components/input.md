# flux:input

Text input with icons, masks, keyboard hints, and interactive features.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Binds to Livewire property |
| `label` | string | - | Wraps in field with label |
| `description` | string | - | Help text between label and input |
| `description:trailing` | string | - | Help text below input |
| `placeholder` | string | - | Text shown when empty |
| `size` | string | - | `sm`, `xs` |
| `variant` | string | outline | `filled`, `outline` |
| `disabled` | boolean | false | Prevents interaction |
| `readonly` | boolean | false | Locks during submission |
| `invalid` | boolean | false | Error styling |
| `multiple` | boolean | false | File input: multiple files |
| `mask` | string | - | Input mask pattern |
| `mask:dynamic` | string | - | Dynamic mask pattern |
| `icon` | string | - | Leading icon name |
| `icon:trailing` | string | - | Trailing icon name |
| `kbd` | string | - | Keyboard shortcut hint |
| `clearable` | boolean | false | Show clear button when filled |
| `copyable` | boolean | false | Add copy button (HTTPS only) |
| `viewable` | boolean | false | Password visibility toggle |
| `as` | string | input | `button`, `input` |
| `class:input` | string | - | Classes on input element |

## Slots

| Slot | Description |
|------|-------------|
| `icon` / `icon:leading` | Custom leading content |
| `icon:trailing` | Custom trailing content |

## Child Components

### flux:input.group

Container for grouped inputs with prefix/suffix.

### flux:input.group.prefix

Text/content before input.

### flux:input.group.suffix

Text/content after input.

## Basic Usage

```blade
<flux:input wire:model="name" placeholder="Enter your name" />
```

## With Label (Shorthand)

```blade
<flux:input label="Email" type="email" wire:model="email" />
```

## With Description

```blade
<flux:input
    label="Username"
    description="Letters and numbers only."
    wire:model="username"
/>
```

## Input Types

```blade
<flux:input type="text" wire:model="name" />
<flux:input type="email" wire:model="email" />
<flux:input type="password" wire:model="password" />
<flux:input type="date" wire:model="date" />
<flux:input type="file" wire:model="file" />
<flux:input type="file" wire:model="files" multiple />
```

## With Icons

```blade
<flux:input icon="envelope" wire:model="email" />
<flux:input icon:trailing="magnifying-glass" wire:model="search" />
```

## Keyboard Shortcut Hint

```blade
<flux:input icon="magnifying-glass" kbd="⌘K" placeholder="Search..." />
```

## Clearable

```blade
<flux:input wire:model="search" clearable />
```

## Copyable

```blade
<flux:input wire:model="apiKey" copyable readonly />
```

## Password with Toggle

```blade
<flux:input type="password" wire:model="password" viewable />
```

## Size Variants

```blade
<flux:input size="sm" wire:model="small" />
<flux:input size="xs" wire:model="extraSmall" />
```

## Input Masking

```blade
{{-- Phone number --}}
<flux:input mask="(999) 999-9999" wire:model="phone" />

{{-- Credit card --}}
<flux:input mask="9999 9999 9999 9999" wire:model="card" />

{{-- Dynamic mask --}}
<flux:input mask:dynamic="['99.999.999/9999-99', '999.999.999-99']" wire:model="document" />
```

## Input Group with Prefix/Suffix

```blade
<flux:input.group>
    <flux:input.group.prefix>https://</flux:input.group.prefix>
    <flux:input wire:model="domain" />
    <flux:input.group.suffix>.com</flux:input.group.suffix>
</flux:input.group>
```

## With Button

```blade
<flux:input.group>
    <flux:input wire:model="email" placeholder="Enter email" />
    <flux:button type="submit">Subscribe</flux:button>
</flux:input.group>
```

## Custom Icon Slot

```blade
<flux:input wire:model="search">
    <x-slot name="iconTrailing">
        <flux:button size="sm" icon="magnifying-glass" />
    </x-slot>
</flux:input>
```
