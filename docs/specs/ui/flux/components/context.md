# flux:context

Right-click context menu functionality.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Binds menu state to Livewire property |
| `position` | string | bottom end | `[vertical] [horizontal]` (top/bottom, start/center/end) |
| `gap` | number | 4 | Distance from click position |
| `offset` | string | - | Additional offset `[x] [y]` |
| `target` | string | - | ID of external menu element |
| `detail` | mixed | - | Custom value for styling/behaviour |
| `disabled` | boolean | false | Prevents context menu |

## Slots

| Slot | Description |
|------|-------------|
| `default` | First child = trigger area, second = `flux:menu` |

## Basic Usage

```blade
<flux:context>
    <flux:card class="border-dashed border-2 px-16">
        <flux:text>Right click here</flux:text>
    </flux:card>

    <flux:menu>
        <flux:menu.item icon="plus">New post</flux:menu.item>
        <flux:menu.separator />
        <flux:menu.submenu heading="Sort by">
            <flux:menu.radio.group>
                <flux:menu.radio checked>Name</flux:menu.radio>
                <flux:menu.radio>Date</flux:menu.radio>
            </flux:menu.radio.group>
        </flux:menu.submenu>
        <flux:menu.separator />
        <flux:menu.item variant="danger" icon="trash">Delete</flux:menu.item>
    </flux:menu>
</flux:context>
```
