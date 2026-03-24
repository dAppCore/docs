# flux:heading

Consistent heading styling with configurable sizes and semantic HTML levels.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `size` | string | base | `base`, `lg`, `xl` |
| `level` | number | - | HTML heading level: `1`, `2`, `3`, `4` (default: div) |
| `accent` | boolean | false | Applies accent colour styling |

## Size Reference

| Size | Pixels | Usage |
|------|--------|-------|
| `base` | 14px | Input labels, toast labels, general use |
| `lg` | 16px | Modal headings, card headings |
| `xl` | 24px | Hero text, page titles |

## Basic Usage

```blade
<flux:heading>User profile</flux:heading>
```

## With Size

```blade
<flux:heading size="base">Small heading</flux:heading>
<flux:heading size="lg">Medium heading</flux:heading>
<flux:heading size="xl">Large heading</flux:heading>
```

## Semantic HTML Level

```blade
<flux:heading level="1">Page Title</flux:heading>
<flux:heading level="2">Section Heading</flux:heading>
<flux:heading level="3">Subsection</flux:heading>
```

## With Accent Colour

```blade
<flux:heading accent>Highlighted heading</flux:heading>
```

## Subheading Pattern

```blade
<div>
    <flux:text>Year to date</flux:text>
    <flux:heading size="xl" class="mb-1">$7,532.16</flux:heading>
</div>
```

## Card Heading

```blade
<flux:card>
    <flux:heading size="lg" class="mb-4">Card Title</flux:heading>
    <flux:text>Card content goes here.</flux:text>
</flux:card>
```

## Related: flux:text

For body text and paragraphs.

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `size` | string | base | `sm`, `base`, `lg`, `xl` |

```blade
<flux:text>Regular paragraph text.</flux:text>
<flux:text size="sm">Smaller text.</flux:text>
<flux:text size="lg">Larger text.</flux:text>
```
