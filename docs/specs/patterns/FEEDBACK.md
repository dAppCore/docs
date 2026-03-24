# Feedback Patterns (Tailwind+)

Consolidated list of all feedback patterns from Tailwind+ to convert to Flux UI.

**Total: 12 variants across 2 categories**

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/feedback/

---

## Alerts (6 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/feedback/alerts

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | With description | Icon + title + description text |
| 2 | With list | Alert with bullet point list |
| 3 | With actions | Primary/secondary action buttons |
| 4 | With link on right | Inline link action |
| 5 | With accent border | Left border accent color |
| 6 | With dismiss button | X button to close |

---

## Empty States (6 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/feedback/empty-states

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Simple | Icon + message + CTA button |
| 2 | With dashed border | Dashed border container |
| 3 | With starting points | Multiple getting-started options |
| 4 | With recommendations | Suggested items/actions |
| 5 | With templates | Template selection grid |
| 6 | With recommendations grid | Grid of recommendation cards |

---

## Pattern Categories (for Flux UI organisation)

### By Severity (Alerts)
- **Info**: Blue - informational messages
- **Success**: Green - positive confirmations
- **Warning**: Yellow - caution messages
- **Error**: Red - error states

### By Complexity (Empty States)
- **Simple**: Just message + action
- **Rich**: With suggestions, templates, or guides

### By Features
- **Dismissible**: Can be closed
- **Actionable**: Has buttons/links
- **Informational**: Display only

---

## Conversion Priority

1. **Alerts - With description** - most common alert pattern
2. **Alerts - With dismiss button** - toast-style notifications
3. **Empty States - Simple** - basic placeholder
4. **Alerts - With actions** - confirmation dialogs
5. **Empty States - With recommendations** - onboarding
