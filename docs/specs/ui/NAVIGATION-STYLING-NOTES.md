# Navigation Dropdown Styling - Implementation Notes

> **Status:** Complete (January 2026)
> **Location:** `resources/css/app.css` (lines ~253-370)

## Summary

Custom styling for Flux UI navigation dropdowns with accent-coloured gradients, grid pattern overlays, and consistent hover states.

## Key Implementation Details

### Flux Component Data Attributes

- `[data-flux-navbar-items]` - Note the **plural 's'** - targets navbar item buttons
- `[data-flux-navmenu-item]` - Targets dropdown menu items (singular)
- `[data-flux-navmenu]` - The dropdown menu container itself
- `[data-content]` - Wrapper around text content inside navbar items (has hardcoded `text-sm`)

### Navbar Item Styling

```css
/* Main navbar items - larger text, white colour */
[data-flux-navbar-items] {
    color: white !important;
}

/* Override the hardcoded text-sm on navbar item content */
[data-flux-navbar-items] [data-content] {
    font-size: 1rem !important;
}
```

### Dropdown Menu Card Styling

Each dropdown uses `data-accent` attribute on parent for colour theming:
- `data-accent="purple"` - Services
- `data-accent="orange"` - For
- `data-accent="indigo"` - AI
- `data-accent="cyan"` - Tools
- `data-accent="slate"` - OSS
- `data-accent="amber"` - About
- `data-accent="violet"` - Dashboard

Features:
1. **Asymmetric gradient** - Solid accent colour left, fading to slate-900 right
2. **Accent border** - 3px left border in accent colour
3. **Grid pattern overlay** - Via `::before` pseudo-element with grid.svg
4. **Fading grid** - Uses `mask-image` to fade grid from 50% to right edge
5. **Backdrop blur** - `backdrop-filter: blur(12px)`

### Menu Item Hover Effect

```css
[data-accent] [data-flux-navmenu-item]:hover {
    background: linear-gradient(90deg, rgb(255 255 255 / 0.08), transparent 80%) !important;
    box-shadow: inset 0 0 0 1px rgb(255 255 255 / 0.1) !important;
    color: white !important;
}
```

- Gradient hover that fades to transparent (allows grid to show through)
- Subtle inset border for definition

## File References

- **CSS:** `resources/css/app.css` - Navigation styling section
- **Header:** `resources/views/components/layouts/partials/header.blade.php`
- **Flux stubs:** `vendor/livewire/flux/stubs/resources/views/flux/navbar/item.blade.php`
- **Grid SVG:** `public/vendor/stellar/images/grid.svg`
- **Flux docs:** `doc/ui/flux/navbar.md`

## Browser Notes

- Chrome may cache aggressively - use Shift+Cmd+R for hard refresh when testing CSS changes
- Safari renders backdrop-filter effects more smoothly than Chrome in some cases

## Related Documentation

- `doc/ui/flux/README.md` - Flux UI overview
- `doc/ui/flux/styling.md` - Theming and customisation guide
- `doc/ui/flux/navbar.md` - Navigation component reference
