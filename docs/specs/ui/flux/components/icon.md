# flux:icon

Icons using Heroicons collection with multiple variants and sizes.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `variant` | string | outline | `outline`, `solid`, `mini`, `micro` |
| `name` | string | - | Dynamic icon name (for variable icons) |

## Variants

| Variant | Size | Style |
|---------|------|-------|
| `outline` | 24px | Unfilled (default) |
| `solid` | 24px | Filled |
| `mini` | 20px | Filled, compact |
| `micro` | 16px | Filled, smallest |

## Basic Usage

```blade
<flux:icon.bolt />
```

## With Variants

```blade
<flux:icon.bolt />                    {{-- 24px outline --}}
<flux:icon.bolt variant="solid" />    {{-- 24px filled --}}
<flux:icon.bolt variant="mini" />     {{-- 20px filled --}}
<flux:icon.bolt variant="micro" />    {{-- 16px filled --}}
```

## Sizing

Use Tailwind's `size-*` utilities:

```blade
<flux:icon.bolt class="size-12" />
<flux:icon.bolt class="size-8" />
<flux:icon.bolt class="size-6" />
<flux:icon.bolt class="size-4" />
```

## Styling

Apply colour with Tailwind text utilities:

```blade
<flux:icon.bolt class="text-amber-500" />
<flux:icon.bolt variant="solid" class="text-amber-500 dark:text-amber-300" />
```

## Loading Spinner

```blade
<flux:icon.loading />
```

## Dynamic Icons

When icon name comes from a variable:

```blade
<flux:icon name="bolt" />
<flux:icon :name="$iconName" />
<flux:icon :name="$iconName" variant="solid" />
```

## Import Lucide Icons

Use artisan command to import additional icons:

```bash
# Interactive mode
php artisan flux:icon

# Specific icons
php artisan flux:icon crown grip-vertical github
```

## Custom Icons

Create custom SVG components in `resources/views/flux/icon/`:

```blade
{{-- resources/views/flux/icon/custom-logo.blade.php --}}
@props(['variant' => 'outline'])

<svg {{ $attributes->class([
    'size-6' => $variant === 'outline' || $variant === 'solid',
    'size-5' => $variant === 'mini',
    'size-4' => $variant === 'micro',
]) }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
    <!-- Your SVG paths -->
</svg>
```

Then use:

```blade
<flux:icon.custom-logo />
```

## Common Icons

| Category | Icons |
|----------|-------|
| Actions | `plus`, `minus`, `pencil`, `trash`, `check`, `x-mark` |
| Navigation | `chevron-down`, `chevron-right`, `arrow-left`, `arrow-right` |
| Status | `check-circle`, `x-circle`, `exclamation-triangle`, `information-circle` |
| Objects | `user`, `cog`, `document`, `folder`, `photo`, `link` |
| Communication | `envelope`, `phone`, `chat-bubble-left`, `bell` |

Browse all at [heroicons.com](https://heroicons.com/)
