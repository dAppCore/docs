# Overlays Patterns (Tailwind+)

Consolidated list of all overlay patterns from Tailwind+ to convert to Flux UI.

**Total: 24 variants across 3 categories**

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/overlays/

---

## Modal Dialogs (6 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/overlays/modal-dialogs

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Centered with single action | Single CTA button |
| 2 | Centered with wide buttons | Full-width buttons |
| 3 | Simple alert | Basic alert dialog |
| 4 | Simple with dismiss button | X button to close |
| 5 | Simple with gray footer | Gray footer area |
| 6 | Simple alert with left-aligned buttons | Left-aligned actions |

---

## Drawers (12 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/overlays/drawers

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Empty | Basic empty drawer |
| 2 | Wide empty | Wide empty drawer |
| 3 | With background overlay | Backdrop overlay |
| 4 | With close button on outside | External close button |
| 5 | With branded header | Brand-coloured header |
| 6 | With sticky footer | Fixed footer actions |
| 7 | Create project form example | Form drawer example |
| 8 | Wide create project form example | Wide form drawer |
| 9 | User profile example | Profile display drawer |
| 10 | Wide horizontal user profile example | Wide profile layout |
| 11 | Contact list example | List-based drawer |
| 12 | File details example | Detail view drawer |

---

## Notifications (6 variants)

Source: https://tailwindcss.com/plus/ui-blocks/application-ui/overlays/notifications

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Simple | Basic notification toast |
| 2 | Condensed | Compact notification |
| 3 | With actions below | Action buttons below |
| 4 | With avatar | User avatar included |
| 5 | With split buttons | Split action buttons |
| 6 | With buttons below | Buttons in footer |

---

## Pattern Categories (for Flux UI organisation)

### By Type
- **Modal**: Centered dialogs, alerts
- **Drawer**: Slide-in panels from edge
- **Toast**: Notification popups

### By Complexity
- **Simple**: Basic display, single action
- **Rich**: Multiple actions, forms, complex content

### By Position
- **Center**: Modal dialogs
- **Edge**: Drawers (typically right)
- **Corner**: Notifications (typically top-right)

### By Features
- **Dismissible**: Close button available
- **With overlay**: Background backdrop
- **With actions**: Action buttons included
- **With forms**: Form inputs inside

---

## Conversion Priority

1. **Modal Dialogs - Simple alert** - confirmation dialogs
2. **Notifications - Simple** - toast messages
3. **Drawers - Empty** - slide-out panels
4. **Modal Dialogs - Centered with wide buttons** - action dialogs
5. **Drawers - With sticky footer** - form panels
