# flux:button

Powerful, composable button component with variants, icons, loading states, and tooltips.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `as` | string | button | Render as: `button`, `a`, `div` |
| `href` | string | - | URL (renders as link) |
| `type` | string | button | `button`, `submit` |
| `variant` | string | outline | `outline`, `primary`, `filled`, `danger`, `ghost`, `subtle` |
| `size` | string | base | `base`, `sm`, `xs` |
| `color` | string | - | Tailwind colour name |
| `icon` | string | - | Leading icon name |
| `icon:trailing` | string | - | Trailing icon name |
| `icon:variant` | string | micro | `outline`, `solid`, `mini`, `micro` |
| `square` | boolean | false | Equal width/height (auto for icon-only) |
| `align` | string | center | `start`, `center`, `end` |
| `inset` | string | - | Negative margin: `top`, `bottom`, `left`, `right` (combinable) |
| `loading` | boolean | true | Show loading indicator on `wire:click` |
| `disabled` | boolean | false | Disable interaction |
| `tooltip` | string | - | Hover tooltip text |
| `tooltip:position` | string | top | `top`, `bottom`, `left`, `right` |
| `tooltip:kbd` | string | - | Keyboard shortcut in tooltip |
| `kbd` | string | - | Keyboard shortcut hint |

### Colour Options

`zinc`, `red`, `orange`, `amber`, `yellow`, `lime`, `green`, `emerald`, `teal`, `cyan`, `sky`, `blue`, `indigo`, `violet`, `purple`, `fuchsia`, `pink`, `rose`

---

## Basic Usage

```blade
<flux:button>Button</flux:button>
```

## Variants

```blade
<flux:button variant="outline">Outline</flux:button>
<flux:button variant="primary">Primary</flux:button>
<flux:button variant="filled">Filled</flux:button>
<flux:button variant="danger">Danger</flux:button>
<flux:button variant="ghost">Ghost</flux:button>
<flux:button variant="subtle">Subtle</flux:button>
```

## Sizes

```blade
<flux:button size="xs">Extra small</flux:button>
<flux:button size="sm">Small</flux:button>
<flux:button>Base (default)</flux:button>
```

## Colours

```blade
<flux:button color="indigo">Indigo</flux:button>
<flux:button variant="primary" color="green">Green Primary</flux:button>
<flux:button variant="filled" color="red">Red Filled</flux:button>
```

## With Icons

```blade
<flux:button icon="check">Save</flux:button>
<flux:button icon:trailing="arrow-right">Next</flux:button>
<flux:button icon="pencil" icon:trailing="chevron-down">Edit</flux:button>
```

## Icon Only

```blade
<flux:button icon="plus" />
<flux:button icon="trash" variant="danger" />
<flux:button icon="cog-6-tooth" variant="ghost" />
```

## As Link

```blade
<flux:button href="/dashboard">Go to Dashboard</flux:button>
<flux:button href="/docs" icon="book-open">Documentation</flux:button>
```

## Full Width

```blade
<flux:button class="w-full">Full Width Button</flux:button>
```

## Loading States

Automatic loading indicator with `wire:click`:

```blade
<flux:button wire:click="save">Save changes</flux:button>
```

Disable auto-loading:

```blade
<flux:button wire:click="save" loading="false">Save</flux:button>
```

## With Tooltip

```blade
<flux:button icon="cog-6-tooth" tooltip="Settings" />
<flux:button icon="trash" tooltip="Delete item" tooltip:position="bottom" />
```

## Keyboard Shortcuts

In button:

```blade
<flux:button kbd="⌘S">Save</flux:button>
```

In tooltip:

```blade
<flux:button icon="magnifying-glass" tooltip="Search" tooltip:kbd="⌘K" />
```

## Disabled

```blade
<flux:button disabled>Disabled</flux:button>
<flux:button variant="primary" disabled>Disabled Primary</flux:button>
```

## Inset (Negative Margin)

For buttons at edges of containers:

```blade
<flux:button variant="ghost" inset="left">Back</flux:button>
<flux:button variant="ghost" inset="right">Next</flux:button>
```

## Form Submit

```blade
<form wire:submit="save">
    <flux:input wire:model="name" label="Name" />
    <flux:button type="submit" variant="primary">Save</flux:button>
</form>
```

---

## Button Group

Group multiple buttons with shared borders:

```blade
<flux:button.group>
    <flux:button>Left</flux:button>
    <flux:button>Centre</flux:button>
    <flux:button>Right</flux:button>
</flux:button.group>
```

With icons:

```blade
<flux:button.group>
    <flux:button icon="bold" />
    <flux:button icon="italic" />
    <flux:button icon="underline" />
</flux:button.group>
```

With variants:

```blade
<flux:button.group>
    <flux:button variant="primary">Save</flux:button>
    <flux:button icon="chevron-down" />
</flux:button.group>
```

---

## Common Patterns

### Save/Cancel

```blade
<div class="flex gap-2">
    <flux:button variant="ghost">Cancel</flux:button>
    <flux:button variant="primary">Save</flux:button>
</div>
```

### Danger Action

```blade
<flux:button variant="danger" icon="trash">Delete</flux:button>
```

### Icon with Dropdown

```blade
<flux:dropdown>
    <flux:button icon="ellipsis-horizontal" variant="ghost" />
    <flux:menu>
        <flux:menu.item icon="pencil">Edit</flux:menu.item>
        <flux:menu.item icon="trash" variant="danger">Delete</flux:menu.item>
    </flux:menu>
</flux:dropdown>
```
