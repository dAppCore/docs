# flux:badge

Highlight information like status, category, or count.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `color` | string | zinc | Background/text colour |
| `size` | string | default | `sm`, default, `lg` |
| `variant` | string | default | `pill`, `solid` |
| `icon` | string | - | Leading icon name |
| `icon:trailing` | string | - | Trailing icon name |
| `icon:variant` | string | mini | `outline`, `solid`, `mini`, `micro` |
| `as` | string | div | `button`, `div` |
| `inset` | string | - | Negative margin: `top`, `bottom`, `left`, `right` (combinable) |

### Colour Options

`zinc`, `red`, `orange`, `amber`, `yellow`, `lime`, `green`, `emerald`, `teal`, `cyan`, `sky`, `blue`, `indigo`, `violet`, `purple`, `fuchsia`, `pink`, `rose`

## Child Component: flux:badge.close

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `icon` | string | x-mark | Close icon name |
| `icon:variant` | string | mini | `outline`, `solid`, `mini`, `micro` |

---

## Basic Usage

```blade
<flux:badge>Default</flux:badge>
<flux:badge color="lime">New</flux:badge>
<flux:badge color="red">Error</flux:badge>
<flux:badge color="green">Success</flux:badge>
```

## Sizes

```blade
<flux:badge size="sm">Small</flux:badge>
<flux:badge>Default</flux:badge>
<flux:badge size="lg">Large</flux:badge>
```

## With Icons

```blade
<flux:badge icon="user-circle">Users</flux:badge>
<flux:badge icon="document-text">Files</flux:badge>
<flux:badge icon:trailing="video-camera">Videos</flux:badge>
```

## Pill Variant

```blade
<flux:badge variant="pill">Default pill</flux:badge>
<flux:badge variant="pill" color="blue">Blue pill</flux:badge>
<flux:badge variant="pill" icon="user">With icon</flux:badge>
```

## Solid Variant

```blade
<flux:badge variant="solid" color="green">Solid green</flux:badge>
<flux:badge variant="solid" color="red">Solid red</flux:badge>
```

## As Button

```blade
<flux:badge as="button" variant="pill" icon="plus" size="lg">
    Add item
</flux:badge>
```

## With Close Button

```blade
<flux:badge>
    Admin
    <flux:badge.close wire:click="removeRole('admin')" />
</flux:badge>

<flux:badge color="blue">
    Tag name
    <flux:badge.close />
</flux:badge>
```

## Inset Spacing

Use `inset` to add negative margin when badge is inline with text:

```blade
<flux:heading>
    Page builder
    <flux:badge color="lime" inset="top bottom">New</flux:badge>
</flux:heading>
```

## Multiple Badges

```blade
<div class="flex gap-2">
    <flux:badge color="green">Active</flux:badge>
    <flux:badge color="blue">Featured</flux:badge>
    <flux:badge color="purple">Premium</flux:badge>
</div>
```

## Status Indicators

```blade
<flux:badge variant="pill" color="green" icon="check-circle">Completed</flux:badge>
<flux:badge variant="pill" color="yellow" icon="clock">Pending</flux:badge>
<flux:badge variant="pill" color="red" icon="x-circle">Failed</flux:badge>
```

## Count Badges

```blade
<flux:badge variant="solid" color="red">99+</flux:badge>
<flux:badge variant="pill" color="blue">12</flux:badge>
```
