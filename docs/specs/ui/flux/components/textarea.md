# flux:textarea

Multi-line text input for comments, descriptions, and feedback.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Binds to Livewire property |
| `placeholder` | string | - | Hint text when empty |
| `label` | string | - | Label above textarea |
| `description` | string | - | Help text (above textarea) |
| `description:trailing` | string | - | Help text below textarea |
| `badge` | string | - | Badge in label |
| `rows` | number/string | 4 | Visible lines, or `auto` |
| `resize` | string | vertical | `vertical`, `horizontal`, `both`, `none` |
| `invalid` | boolean | false | Error styling |

## Basic Usage

```blade
<flux:textarea wire:model="message" />
```

## With Label

```blade
<flux:textarea wire:model="notes" label="Notes" />
```

## With Placeholder

```blade
<flux:textarea
    wire:model="notes"
    label="Order notes"
    placeholder="No lettuce, tomato, or onion..."
/>
```

## Fixed Row Height

```blade
<flux:textarea wire:model="summary" rows="2" label="Short summary" />
<flux:textarea wire:model="description" rows="6" label="Full description" />
```

## Auto-sizing

Automatically grows with content:

```blade
<flux:textarea wire:model="content" rows="auto" />
```

## Resize Options

```blade
{{-- Default: vertical only --}}
<flux:textarea wire:model="notes" resize="vertical" />

{{-- Horizontal only --}}
<flux:textarea wire:model="code" resize="horizontal" />

{{-- Both directions --}}
<flux:textarea wire:model="content" resize="both" />

{{-- No resizing --}}
<flux:textarea wire:model="fixed" resize="none" />
```

## With Description

```blade
<flux:textarea
    wire:model="bio"
    label="Bio"
    description="Write a short bio for your profile."
/>
```

## Trailing Description

```blade
<flux:textarea
    wire:model="message"
    label="Message"
    description:trailing="Maximum 500 characters."
/>
```

## With Badge

```blade
<flux:textarea
    wire:model="feedback"
    label="Feedback"
    badge="Optional"
/>
```

## Invalid State

```blade
<flux:textarea
    wire:model="content"
    label="Content"
    :invalid="$errors->has('content')"
/>
<flux:error name="content" />
```

## In Form

```blade
<form wire:submit="submit">
    <flux:input wire:model="subject" label="Subject" />
    <flux:textarea wire:model="message" label="Message" rows="6" class="mt-4" />
    <flux:button type="submit" class="mt-4">Send</flux:button>
</form>
```
