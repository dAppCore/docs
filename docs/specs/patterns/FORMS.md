# Forms Patterns (Tailwind+)

Consolidated list of all form patterns from Tailwind+ to convert to Flux UI.

**Total: 74 variants across 10 categories**

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/forms/

---

## Form Layouts (4 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/forms/form-layouts

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Stacked | Single column, labels above inputs |
| 2 | Two-column | Split layout for larger forms |
| 3 | Two-column with cards | Sections in card containers |
| 4 | Labels on left | Horizontal label/input layout |

---

## Input Groups (21 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/forms/input-groups

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Input with label | Basic labeled input |
| 2 | Input with label and help text | Label + helper text |
| 3 | Input with validation error | Error state styling |
| 4 | Input with disabled state | Disabled input |
| 5 | Input with hidden label | Screen reader only label |
| 6 | Input with corner hint | Optional/required hint |
| 7 | Input with leading icon | Icon before input |
| 8 | Input with trailing icon | Icon after input |
| 9 | Input with add-on | Prefix/suffix text |
| 10 | Input with inline add-on | Integrated add-on |
| 11 | Input with inline leading and trailing add-ons | Both add-ons |
| 12 | Input with inline leading dropdown | Dropdown prefix |
| 13 | Input with inline leading add-on and trailing dropdown | Mixed add-ons |
| 14 | Input with leading icon and trailing button | Icon + button combo |
| 15 | Inputs with shared borders | Grouped inputs |
| 16 | Input with inset label | Label inside input |
| 17 | Inputs with inset labels and shared borders | Grouped inset inputs |
| 18 | Input with overlapping label | Floating label effect |
| 19 | Input with pill shape | Rounded pill input |
| 20 | Input with gray background and bottom border | Underline style |
| 21 | Input with keyboard shortcut | Shows keyboard hint |

---

## Select Menus (7 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/forms/select-menus

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Simple native | Browser default select |
| 2 | Simple custom | Custom styled dropdown |
| 3 | Custom with check on left | Check mark alignment |
| 4 | Custom with status indicator | Status dots |
| 5 | Custom with avatar | User/item avatars |
| 6 | With secondary text | Additional info text |
| 7 | Branded with supported text | Brand styling |

---

## Sign-in and Registration (4 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/forms/sign-in-forms

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Simple | Basic centered form |
| 2 | Simple no labels | Placeholder-only inputs |
| 3 | Split screen | Image + form layout |
| 4 | Simple card | Form in card container |

---

## Textareas (5 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/forms/textareas

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Simple | Basic textarea |
| 2 | With avatar and actions | Comment-style with user avatar |
| 3 | With underline and actions | Minimal underline style |
| 4 | With title and pill actions | Rich text editor style |
| 5 | With preview button | Markdown preview toggle |

---

## Radio Groups (12 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/forms/radio-groups

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Simple list | Basic vertical list |
| 2 | Simple inline list | Horizontal layout |
| 3 | List with description | Description per option |
| 4 | List with inline description | Compact descriptions |
| 5 | List with radio on right | Right-aligned radios |
| 6 | Simple list with radio on right | Right-aligned simple |
| 7 | Simple table | Table layout |
| 8 | List with descriptions in panel | Card-style options |
| 9 | Color picker | Color swatch selection |
| 10 | Cards | Option cards |
| 11 | Small cards | Compact option cards |
| 12 | Stacked cards | Full-width cards |

---

## Checkboxes (4 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/forms/checkboxes

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | List with description | Description per checkbox |
| 2 | List with inline description | Compact layout |
| 3 | List with checkbox on right | Right-aligned checkboxes |
| 4 | Simple list with heading | Grouped with heading |

---

## Toggles (5 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/forms/toggles

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Simple toggle | Basic switch |
| 2 | Short toggle | Compact switch |
| 3 | Toggle with icon | Icon inside toggle |
| 4 | With left label and description | Label + description |
| 5 | With right label | Right-aligned label |

---

## Action Panels (8 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/forms/action-panels

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Simple | Basic action panel |
| 2 | With link | Link action |
| 3 | With button on right | Right-aligned button |
| 4 | With button at top right | Top-right button |
| 5 | With toggle | Toggle action |
| 6 | With input | Input field action |
| 7 | Simple well | Well background |
| 8 | With well | Complex well layout |

---

## Comboboxes (4 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/forms/comboboxes

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Simple | Basic autocomplete |
| 2 | With status indicator | Status dots in options |
| 3 | With image | Avatars in options |
| 4 | With secondary text | Additional info text |

---

## Pattern Categories (for Flux UI organisation)

### By Input Type
- **Text inputs**: Input groups, textareas, comboboxes
- **Selection**: Select menus, radio groups, checkboxes
- **Boolean**: Toggles
- **Composite**: Action panels, sign-in forms

### By Layout
- **Stacked**: Vertical arrangement
- **Inline/Horizontal**: Side-by-side layout
- **Grid**: Multi-column forms
- **Cards**: Card-wrapped inputs

### By State
- **Default**: Normal state
- **Error**: Validation errors
- **Disabled**: Non-interactive

---

## Conversion Priority

1. **Input with label** - essential for all forms
2. **Simple toggle** - boolean inputs
3. **Simple custom select** - dropdowns
4. **Simple radio list** - single selection
5. **Form layout - Stacked** - basic form structure
6. **Sign-in - Simple** - authentication pages
