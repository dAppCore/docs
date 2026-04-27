# flux:popover

Floating overlay for displaying additional content on click or hover.

## Props (via flux:dropdown)

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `position` | string | bottom | `top`, `right`, `bottom`, `left` |
| `align` | string | start | `start`, `center`, `end` |
| `hover` | boolean | false | Opens on hover instead of click |
| `offset` | number | - | Shifts along alignment axis |
| `gap` | number | - | Distance between trigger and popover |
| `wire:model` | string | - | Bind open/closed state |

## Structure

Popover must be wrapped in `flux:dropdown`:

```blade
<flux:dropdown>
    <flux:button>Trigger</flux:button>
    <flux:popover>
        Content here
    </flux:popover>
</flux:dropdown>
```

## Basic Usage

```blade
<flux:dropdown>
    <flux:button icon="information-circle" variant="ghost" />
    <flux:popover class="max-w-xs p-4">
        <flux:heading size="sm">Information</flux:heading>
        <flux:text size="sm" class="mt-2">
            This is additional context about the feature.
        </flux:text>
    </flux:popover>
</flux:dropdown>
```

## Positioning

```blade
{{-- Top --}}
<flux:dropdown position="top">
    <flux:button>Top</flux:button>
    <flux:popover>Content</flux:popover>
</flux:dropdown>

{{-- Right aligned to end --}}
<flux:dropdown position="right" align="end">
    <flux:button>Right End</flux:button>
    <flux:popover>Content</flux:popover>
</flux:dropdown>
```

## Hover Trigger

```blade
<flux:dropdown hover>
    <flux:button>Hover me</flux:button>
    <flux:popover class="p-4">
        <flux:text>This appears on hover.</flux:text>
    </flux:popover>
</flux:dropdown>
```

## With Form Controls

```blade
<flux:dropdown>
    <flux:button>Filter</flux:button>
    <flux:popover class="w-80 p-4 space-y-4">
        <flux:heading size="sm">Filter Options</flux:heading>

        <flux:checkbox.group wire:model="filters">
            <flux:checkbox value="active" label="Active" />
            <flux:checkbox value="pending" label="Pending" />
            <flux:checkbox value="archived" label="Archived" />
        </flux:checkbox.group>

        <div class="flex justify-end gap-2">
            <flux:button variant="ghost" size="sm">Reset</flux:button>
            <flux:button variant="primary" size="sm">Apply</flux:button>
        </div>
    </flux:popover>
</flux:dropdown>
```

## Programmatic Control

```blade
<flux:dropdown wire:model="showPopover">
    <flux:button>Toggle</flux:button>
    <flux:popover>Content</flux:popover>
</flux:dropdown>
```

```php
public bool $showPopover = false;

public function togglePopover()
{
    $this->showPopover = !$this->showPopover;
}
```

## Gap and Offset

```blade
<flux:dropdown gap="8" offset="4">
    <flux:button>Spaced</flux:button>
    <flux:popover>Content with custom spacing</flux:popover>
</flux:dropdown>
```
