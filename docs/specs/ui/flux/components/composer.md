# flux:composer

Configurable message input for chat interfaces and AI prompts.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Binds to Livewire property |
| `name` | string | - | Name for validation |
| `placeholder` | string | - | Placeholder text |
| `label` | string | - | Label text |
| `label:sr-only` | boolean | false | Screen reader only label |
| `description` | string | - | Help text |
| `rows` | number | 2 | Visible text lines |
| `max-rows` | number | - | Maximum expandable rows |
| `inline` | boolean | false | Single-row action layout |
| `submit` | string | cmd+enter | `cmd+enter`, `enter` |
| `disabled` | boolean | false | Prevents interaction |
| `invalid` | boolean | false | Error styling |
| `variant` | string | - | `input` |

## Slots

| Slot | Description |
|------|-------------|
| `input` | Custom input (e.g., rich text editor) |
| `header` | Content above input |
| `footer` | Content below input |
| `actionsLeading` | Start-side action buttons |
| `actionsTrailing` | End-side action buttons |

## Basic Usage

```blade
<flux:composer wire:model="prompt" placeholder="How can I help you today?">
    <x-slot name="actionsTrailing">
        <flux:button type="submit" icon="paper-airplane" />
    </x-slot>
</flux:composer>
```

## With Leading Actions

```blade
<flux:composer wire:model="message">
    <x-slot name="actionsLeading">
        <flux:button icon="photo" variant="ghost" />
        <flux:button icon="paper-clip" variant="ghost" />
    </x-slot>
    <x-slot name="actionsTrailing">
        <flux:button type="submit" variant="primary">Send</flux:button>
    </x-slot>
</flux:composer>
```

## Inline Layout

```blade
<flux:composer wire:model="search" inline placeholder="Ask a question...">
    <x-slot name="actionsTrailing">
        <flux:button type="submit" icon="arrow-right" />
    </x-slot>
</flux:composer>
```

## With Rich Text Editor

```blade
<flux:composer wire:model="content">
    <x-slot name="input">
        <flux:editor wire:model="content" />
    </x-slot>
</flux:composer>
```
