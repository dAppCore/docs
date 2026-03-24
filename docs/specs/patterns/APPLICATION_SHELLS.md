# Application Shells (Tailwind+)

Consolidated list of all application shell patterns from Tailwind+ to convert to Flux UI.

**Total: 23 variants across 3 categories**

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/application-shells/

---

## Stacked Layouts (9 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/application-shells/stacked

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | With lighter page header | Light header on light background |
| 2 | With bottom border | Header with bottom border separator |
| 3 | On subtle background | Gray/muted background variant |
| 4 | Branded nav with compact lighter page header | Brand color nav + compact page header |
| 5 | With overlap | Content overlaps header area |
| 6 | Brand nav with overlap | Branded nav + overlap effect |
| 7 | Branded nav with lighter page header | Full branded nav + light page header |
| 8 | With compact lighter page header | Minimal height page header |
| 9 | Two-row navigation with overlap | Dual nav rows + overlap |

---

## Sidebar Layouts (8 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/application-shells/sidebar

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Simple sidebar | Basic sidebar navigation |
| 2 | Simple dark sidebar | Dark themed sidebar |
| 3 | Sidebar with header | Sidebar + top header bar |
| 4 | Dark sidebar with header | Dark sidebar + header |
| 5 | With constrained content area | Max-width content container |
| 6 | With off-white background | Subtle background color |
| 7 | Simple brand sidebar | Brand colored sidebar |
| 8 | Brand sidebar with header | Brand sidebar + header |

---

## Multi-Column Layouts (6 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/application-shells/multi-column

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Full-width three-column | Three columns, edge-to-edge |
| 2 | Full-width secondary column on right | Main + secondary right column |
| 3 | Constrained three column | Max-width three-column layout |
| 4 | Constrained with sticky columns | Sticky side columns on scroll |
| 5 | Full-width with narrow sidebar | Narrow left sidebar + main |
| 6 | Full-width with narrow sidebar and header | Narrow sidebar + top header |

---

## Pattern Categories (for Flux UI organisation)

### By Structure
- **Stacked**: Vertical layout, no sidebar (variants 1-9 stacked)
- **Sidebar**: Persistent side navigation (variants 1-8 sidebar)
- **Multi-column**: 2-3 column layouts (variants 1-6 multi-column)

### By Theme
- **Light**: Standard light theme (most variants)
- **Dark**: Dark themed components (simple dark sidebar, dark sidebar with header)
- **Branded**: Custom brand colors (branded nav variants, brand sidebar variants)

### By Features
- **With header**: Includes top navigation bar
- **With overlap**: Content overlaps into header area
- **Constrained**: Max-width container
- **Full-width**: Edge-to-edge layout
- **Sticky**: Fixed positioning on scroll

---

## Conversion Priority

1. **Simple sidebar** - most common dashboard layout
2. **Sidebar with header** - dashboard with top nav
3. **Stacked with lighter page header** - simple app layout
4. **Full-width with narrow sidebar** - compact navigation
5. **Constrained three column** - complex dashboards
