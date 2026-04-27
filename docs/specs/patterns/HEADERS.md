# Header Patterns (Tailwind+)

Consolidated list of all header/navbar patterns from Tailwind+ to convert to Flux UI.

**Total: 27 variants across 3 sources**

---

## Marketing Headers (11 variants)

Source: https://tailwindcss.com/plus/ui-blocks/marketing/elements/headers

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | With stacked flyout menu | Mega menu with stacked layout |
| 2 | Constrained | Max-width container |
| 3 | On brand background | Dark/colored background |
| 4 | With full width flyout menu | Full-width mega menu dropdown |
| 5 | Full width | Edge-to-edge header |
| 6 | With call-to-action | CTA button in header |
| 7 | With multiple flyout menus | Multiple mega menu dropdowns |
| 8 | With icons in mobile menu | Icons next to mobile nav items |
| 9 | With left-aligned nav | Logo left, nav items immediately after |
| 10 | With right-aligned nav | Logo left, nav pushed right |
| 11 | With centered logo | Logo center, nav split left/right |

---

## Application UI Navbars (11 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/navigation/navbars

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Simple dark with menu button on left | Dark theme, hamburger menu |
| 2 | Dark with quick action | Dark + primary action button |
| 3 | Simple dark | Minimal dark navbar |
| 4 | Simple with menu button on left | Light theme, hamburger menu |
| 5 | Simple | Minimal light navbar |
| 6 | With quick action | Light + primary action button |
| 7 | Dark with search | Dark + search input |
| 8 | With search | Light + search input |
| 9 | Dark with centered search and secondary links | Two-row: search top, nav below (dark) |
| 10 | With centered search and secondary links | Two-row: search top, nav below (light) |
| 11 | With search in column layout | Compact with wide search bar |

---

## Ecommerce Store Navigation (5 variants)

Source: https://tailwindcss.com/plus/ui-blocks/ecommerce/components/store-navigation

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | With image grid | Mega menu with product images |
| 2 | With simple menu and promo | Promo banner + mega menu |
| 3 | With featured categories | Colored promo banner + featured images in dropdown |
| 4 | With centered logo and featured categories | Centered logo, nav split, featured images |
| 5 | With double column and persistent mobile nav | Double column mega menu |

---

## Pattern Categories (for Flux UI organisation)

When converting, organise by feature rather than Tailwind's categories:

### By Layout
- **Left-aligned**: Logo left, nav follows
- **Right-aligned**: Logo left, nav pushed right
- **Centered logo**: Logo center, nav split
- **Full width**: Edge-to-edge
- **Constrained**: Max-width container

### By Features
- **Simple**: Just nav links
- **With search**: Includes search input
- **With CTA**: Primary action button
- **With mega menu**: Dropdown with columns/images
- **With promo**: Banner/announcement bar
- **Two-row**: Stacked layout (search + nav, or promo + nav)

### By Theme
- **Light**: White/light background
- **Dark**: Dark background
- **On brand**: Custom brand color background

### By Context
- **Marketing**: Public-facing, conversion-focused
- **App/Dashboard**: Authenticated user context
- **Ecommerce**: Shopping-focused with cart

---

## Conversion Priority

1. **Simple** - baseline for all projects
2. **With search** - dashboard essential
3. **With CTA** - marketing essential
4. **With mega menu** - complex sites
5. **Ecommerce variants** - shop features
