# Data Display Patterns (Tailwind+)

Consolidated list of all data display patterns from Tailwind+ to convert to Flux UI.

**Total: 19 variants across 3 categories**

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/data-display/

---

## Description Lists (6 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/data-display/description-lists

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Left-aligned | Labels left, values right, simple rows |
| 2 | Left-aligned in card | Same layout wrapped in card container |
| 3 | Left-aligned striped | Alternating row backgrounds |
| 4 | Two-column | Data displayed in 2-column grid |
| 5 | Left-aligned with inline actions | Edit/Update buttons per row |
| 6 | Narrow with hidden labels | Compact invoice-style with icons |

---

## Stats (5 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/data-display/stats

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | With trending | Stats with up/down trend indicators |
| 2 | Simple | Basic stat cards, no trends |
| 3 | Simple in cards | Stats in card containers |
| 4 | With brand icon | Colored icons per stat |
| 5 | With shared borders | Stats in single card with dividers |

---

## Calendars (8 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/data-display/calendars

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Small with meetings | Compact calendar + meeting list |
| 2 | Month view | Full month grid with events |
| 3 | Week view | 7-day view with time slots |
| 4 | Day view | Single day with time slots + mini calendar |
| 5 | Year view | 12-month overview grid |
| 6 | Double | Two months side-by-side + upcoming events |
| 7 | Borderless stacked | Calendar above, schedule below |
| 8 | Borderless side-by-side | Calendar left, schedule right |

---

## Pattern Categories (for Flux UI organisation)

### By Complexity
- **Simple**: Basic data display (description lists, simple stats)
- **Interactive**: With actions/buttons (inline actions)
- **Complex**: Multi-component (calendars with events)

### By Layout
- **Single column**: Description lists, stats
- **Multi-column**: Two-column description, calendar views
- **Grid**: Year view, month view

### By Features
- **With actions**: Inline edit/update buttons
- **With trends**: Up/down indicators
- **With icons**: Brand icons, status icons
- **With events**: Calendar events, meetings

---

## Conversion Priority

1. **Description list - Left-aligned** - most common data display
2. **Stats - Simple in cards** - dashboard essential
3. **Calendar - Month view** - scheduling features
4. **Stats - With trending** - analytics dashboards
5. **Description list - With inline actions** - editable data
