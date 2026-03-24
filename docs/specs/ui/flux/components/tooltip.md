# flux:tooltip

Contextual information on hover or focus.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `content` | string | - | Tooltip text |
| `position` | string | top | `top`, `right`, `bottom`, `left` |
| `align` | string | center | `center`, `start`, `end` |
| `disabled` | boolean | false | Prevents tooltip |
| `gap` | string | 5px | Space from trigger |
| `offset` | string | 0px | Offset from trigger |
| `toggleable` | boolean | false | Click instead of hover |
| `interactive` | boolean | false | ARIA for interactive content |
| `kbd` | string | - | Keyboard shortcut hint |

## Child Components

### flux:tooltip.content

Custom tooltip content with optional keyboard hint.

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `kbd` | string | - | Keyboard shortcut |

## Basic Usage

```blade
<flux:tooltip content="Settings">
    <flux:button icon="cog-6-tooth" />
</flux:tooltip>
```

## Button Shorthand

```blade
<flux:button tooltip="Settings" icon="cog-6-tooth" />
```

## Positioning

```blade
<flux:tooltip content="Above" position="top">
    <flux:button>Top</flux:button>
</flux:tooltip>

<flux:tooltip content="Right side" position="right">
    <flux:button>Right</flux:button>
</flux:tooltip>

<flux:tooltip content="Below" position="bottom">
    <flux:button>Bottom</flux:button>
</flux:tooltip>

<flux:tooltip content="Left side" position="left">
    <flux:button>Left</flux:button>
</flux:tooltip>
```

## Alignment

```blade
<flux:tooltip content="Start aligned" position="bottom" align="start">
    <flux:button>Start</flux:button>
</flux:tooltip>

<flux:tooltip content="End aligned" position="bottom" align="end">
    <flux:button>End</flux:button>
</flux:tooltip>
```

## With Keyboard Shortcut

```blade
<flux:tooltip>
    <flux:button icon="magnifying-glass" />
    <flux:tooltip.content kbd="⌘K">Search</flux:tooltip.content>
</flux:tooltip>
```

## Toggleable (Touch Devices)

```blade
<flux:tooltip toggleable>
    <flux:icon.information-circle class="text-zinc-400" />
    <flux:tooltip.content>
        Essential information that needs to be accessible on touch devices.
    </flux:tooltip.content>
</flux:tooltip>
```

## Interactive Content

```blade
<flux:tooltip interactive>
    <flux:button>Help</flux:button>
    <flux:tooltip.content>
        <flux:heading size="sm">Need help?</flux:heading>
        <flux:text size="sm">Contact support for assistance.</flux:text>
    </flux:tooltip.content>
</flux:tooltip>
```

## Custom Gap/Offset

```blade
<flux:tooltip content="Spaced tooltip" gap="10px" offset="5px">
    <flux:button>Hover</flux:button>
</flux:tooltip>
```

## Disabled Tooltip

```blade
<flux:tooltip content="Won't show" :disabled="true">
    <flux:button>No tooltip</flux:button>
</flux:tooltip>
```

## On Disabled Button

Wrap disabled buttons in a div (pointer events limitation):

```blade
<flux:tooltip content="Feature not available">
    <div>
        <flux:button disabled>Unavailable</flux:button>
    </div>
</flux:tooltip>
```

## Toolbar Example

```blade
<div class="flex gap-1">
    <flux:tooltip content="Bold (⌘B)">
        <flux:button icon="bold" variant="ghost" size="sm" />
    </flux:tooltip>
    <flux:tooltip content="Italic (⌘I)">
        <flux:button icon="italic" variant="ghost" size="sm" />
    </flux:tooltip>
    <flux:tooltip content="Underline (⌘U)">
        <flux:button icon="underline" variant="ghost" size="sm" />
    </flux:tooltip>
</div>
```
