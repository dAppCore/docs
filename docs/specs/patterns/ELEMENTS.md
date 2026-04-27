# Elements Patterns (Tailwind+)

Consolidated list of all element patterns from Tailwind+ to convert to Flux UI.

**Total: 45 variants across 5 categories**

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/elements/

---

## Avatars (11 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/elements/avatars

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Circular avatars | Round avatar images, multiple sizes |
| 2 | Rounded avatars | Rounded square avatars |
| 3 | Circular avatars with top notification | Status dot top-right |
| 4 | Rounded avatars with top notification | Rounded + top status |
| 5 | Circular avatars with bottom notification | Status dot bottom-right |
| 6 | Rounded avatars with bottom notification | Rounded + bottom status |
| 7 | Circular avatars with placeholder icon | User icon fallback |
| 8 | Circular avatars with placeholder initials | Initials fallback |
| 9 | Avatar group stacked bottom to top | Overlapping group, last on top |
| 10 | Avatar group stacked top to bottom | Overlapping group, first on top |
| 11 | With text | Avatar + name/description |

---

## Badges (16 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/elements/badges

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | With border | Outlined badges, multiple colors |
| 2 | With border and dot | Border + status dot |
| 3 | Pill with border | Rounded pill shape + border |
| 4 | Pill with border and dot | Pill + border + dot |
| 5 | With border and remove button | Dismissible badge |
| 6 | Flat | Solid background, no border |
| 7 | Flat pill | Solid pill shape |
| 8 | Flat with dot | Solid + status dot |
| 9 | Flat pill with dot | Solid pill + dot |
| 10 | Flat with remove button | Solid dismissible |
| 11 | Small with border | Compact bordered |
| 12 | Small flat | Compact solid |
| 13 | Small pill with border | Compact pill bordered |
| 14 | Small flat pill | Compact solid pill |
| 15 | Small flat with dot | Compact solid + dot |
| 16 | Small flat pill with dot | Compact solid pill + dot |

---

## Dropdowns (5 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/elements/dropdowns

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Simple | Basic dropdown menu |
| 2 | With dividers | Grouped menu items |
| 3 | With icons | Icons beside menu items |
| 4 | With minimal menu icon | Three-dot trigger |
| 5 | With simple header | Header text above items |

---

## Buttons (8 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/elements/buttons

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Primary buttons | Solid primary color, multiple sizes |
| 2 | Secondary buttons | Outlined/subtle style |
| 3 | Soft buttons | Light background tint |
| 4 | Buttons with leading icon | Icon before text |
| 5 | Buttons with trailing icon | Icon after text |
| 6 | Rounded primary buttons | Pill-shaped primary |
| 7 | Rounded secondary buttons | Pill-shaped secondary |
| 8 | Circular buttons | Icon-only round buttons |

---

## Button Groups (5 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/elements/button-groups

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Basic | Connected button group |
| 2 | Icon only | Icon buttons grouped |
| 3 | With stat | Button + count badge |
| 4 | With dropdown | Button + dropdown trigger |
| 5 | With checkbox and dropdown | Select + dropdown combo |

---

## Pattern Categories (for Flux UI organisation)

### By Type
- **Display**: Avatars, badges (non-interactive display)
- **Interactive**: Buttons, dropdowns, button groups

### By Size
- **Standard**: Regular size elements
- **Small/Compact**: Reduced size variants
- **Multiple sizes**: Size scale (xs, sm, md, lg, xl)

### By Style
- **Solid/Flat**: Filled background
- **Outlined/Border**: Border with transparent fill
- **Soft**: Light tinted background
- **Pill/Rounded**: Full border-radius

### By Features
- **With icons**: Leading/trailing icons
- **With status**: Notification dots, status indicators
- **With actions**: Remove buttons, dropdowns
- **Grouped**: Stacked avatars, button groups

---

## Conversion Priority

1. **Primary buttons** - essential for all UIs
2. **Badges - Flat** - status indicators
3. **Avatars - Circular** - user display
4. **Dropdowns - Simple** - action menus
5. **Button groups - Basic** - toolbar patterns
