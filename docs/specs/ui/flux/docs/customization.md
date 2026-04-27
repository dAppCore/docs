# Flux Customisation

Multiple levels of customisation without complete rewrites.

## 1. Tailwind Class Overrides

Pass custom classes directly:

```blade
<flux:select class="max-w-md" />
```

### Handling Conflicts

When classes conflict with Flux internals, use `!` modifier:

```blade
<flux:button class="bg-zinc-800! hover:bg-zinc-700!">
    Custom Background
</flux:button>
```

**Recommendation:** Create new components/variants rather than heavy `!` usage.

## 2. Publishing Components

Full control by publishing to your project:

```bash
# Interactive selection
php artisan flux:publish

# All components
php artisan flux:publish --all
```

**Location:** `resources/views/flux/[component-name].blade.php`

Published components:
- Full control over classes, slots, variants
- No automatic updates from Flux

## 3. Global Style Overrides

Use `data-flux-*` attributes for broad changes:

```css
<style>
    [data-flux-button] {
        @apply bg-zinc-800 dark:bg-zinc-400 hover:bg-zinc-700;
    }

    [data-flux-input] {
        @apply rounded-none;
    }

    [data-flux-card] {
        @apply shadow-lg;
    }
</style>
```

Modifies all matching elements globally.

## Customisation Levels Summary

| Level | Use Case | Scope |
|-------|----------|-------|
| Tailwind classes | Simple styling | Single instance |
| `!` modifier | Override conflicts | Single instance |
| Published components | Major changes | All instances |
| Data attributes | Global styling | All instances |

## Available Data Attributes

Every Flux component has a `data-flux-*` attribute:

- `data-flux-button`
- `data-flux-input`
- `data-flux-card`
- `data-flux-modal`
- `data-flux-dropdown`
- `data-flux-table`
- etc.
