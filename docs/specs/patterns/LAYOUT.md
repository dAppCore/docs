# Layout Patterns (Tailwind+)

Consolidated list of all layout patterns from Tailwind+ to convert to Flux UI.

**Total: 38 variants across 5 categories**

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/layout/

---

## Containers (5 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/layout/containers

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Full-width on mobile, constrained to breakpoint with padded content above | Responsive container |
| 2 | Constrained to breakpoint with padded content | Fixed-width padded |
| 3 | Full-width on mobile, constrained with padded content above mobile | Mobile-first responsive |
| 4 | Constrained with padded content | Simple constrained |
| 5 | Narrow constrained with padded content | Narrow width variant |

---

## Cards (10 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/layout/cards

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Basic card | Simple card container |
| 2 | Card with header | Header section + body |
| 3 | Card with header and footer | Header + body + footer |
| 4 | Card with gray header | Gray-themed header |
| 5 | Card with gray body | Gray-themed body |
| 6 | Card with gray footer | Gray-themed footer |
| 7 | Card with header and gray footer | Mixed styling |
| 8 | Card with image | Image header card |
| 9 | Card with full-width header and footer | Edge-to-edge sections |
| 10 | Card with divider | Internal divider line |

---

## List Containers (7 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/layout/list-containers

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Simple | Basic list wrapper |
| 2 | Card | List in card container |
| 3 | Card with header | Card list with header |
| 4 | Separate cards | Individual card per item |
| 5 | Separate cards with header above | Header + separate cards |
| 6 | Simple with dividers | Divided list items |
| 7 | Card with dividers | Card list with dividers |

---

## Media Objects (8 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/layout/media-objects

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Basic responsive | Image + content layout |
| 2 | Wide responsive | Wider image variant |
| 3 | Stretched to fit | Full-width stretch |
| 4 | Nested | Nested media objects |
| 5 | Actions dropdown | With action menu |
| 6 | Basic aligned to bottom | Bottom alignment |
| 7 | Basic aligned to center | Center alignment |
| 8 | Wide aligned to center | Wide centered |

---

## Dividers (8 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/layout/dividers

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | With label | Text in center of divider |
| 2 | With title | Title text divider |
| 3 | With title on left | Left-aligned title |
| 4 | With button | Button in divider |
| 5 | With icon | Icon in center |
| 6 | With toolbar | Multiple controls |
| 7 | Simple | Basic horizontal line |
| 8 | With label on left | Left-aligned label |

---

## Pattern Categories (for Flux UI organisation)

### By Type
- **Structural**: Containers (page-level)
- **Content**: Cards, list containers, media objects
- **Decorative**: Dividers

### By Features
- **With headers**: Cards, list containers
- **With footers**: Cards
- **With dividers**: Cards, list containers, divider components
- **Responsive**: Containers, media objects

### By Complexity
- **Simple**: Basic variants
- **Compound**: With multiple sections (header/body/footer)
- **Nested**: Support for nested content

---

## Conversion Priority

1. **Cards - Basic card** - fundamental container
2. **Cards - Card with header and footer** - common dashboard pattern
3. **Containers - Constrained with padded content** - page wrapper
4. **List containers - Card with dividers** - data lists
5. **Dividers - Simple** - section separators
