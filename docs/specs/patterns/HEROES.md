# Hero Patterns (Tailwind+)

Consolidated list of all hero section patterns from Tailwind+ to convert to Flux UI.

**Total: 12 variants**

Source: https://tailwindcss.com/plus/ui-blocks/marketing/sections/heroes

---

## Variants

| # | Variant | Key Features |
|---|---------|--------------|
| 1 | Simple centered | Announcement badge, headline, description, dual CTAs, centered layout |
| 2 | Split with screenshot | Text left, app screenshot right, two-column layout |
| 3 | Split with bordered screenshot | Text left, bordered app screenshot right |
| 4 | Split with code example | Text left, code snippet/terminal right |
| 5 | Simple centered with background image | Full background photo with overlay, centered text |
| 6 | With bordered app screenshot | Centered hero text, bordered screenshot below |
| 7 | With app screenshot | Centered hero text, screenshot below (no border) |
| 8 | With phone mockup | Text left, mobile device mockup right |
| 9 | Split with image | Text left, photo/image right |
| 10 | With angled image on right | Diagonal/angled photo positioning |
| 11 | With image tiles | Text left, multiple image grid on right |
| 12 | With offset image | Asymmetric/offset photo positioning |

---

## Pattern Categories (for Flux UI organisation)

### By Layout
- **Centered**: Single column, centered content (variants 1, 5, 6, 7)
- **Split/Two-column**: Content left, media right (variants 2, 3, 4, 8, 9, 10, 11, 12)

### By Media Type
- **App screenshots**: Browser/app interface mockups (variants 2, 3, 6, 7)
- **Phone mockups**: Mobile device frames (variant 8)
- **Photos/Images**: Real photography (variants 5, 9, 10, 11, 12)
- **Code examples**: Terminal/code snippets (variant 4)
- **No media**: Text-only with CTAs (variant 1)

### By Background
- **Plain/gradient**: Standard background (most variants)
- **Full image background**: Photo with overlay (variant 5)

### By Features
- **Announcement badge**: Top badge/pill (variant 1)
- **Dual CTAs**: Primary + secondary buttons (most variants)
- **Image grids**: Multiple images composed (variant 11)
- **Angled/offset media**: Creative positioning (variants 10, 12)

---

## Common Elements

All heroes typically include:
- Headline (large, bold)
- Description/subheadline
- Primary CTA button
- Secondary CTA (link or ghost button)

Most heroes include:
- Visual element (screenshot, mockup, image)

Some heroes include:
- Announcement badge/pill
- Background image/gradient
- Multiple images/grid layout

---

## Conversion Priority

1. **Simple centered** - baseline for all marketing sites
2. **Split with screenshot** - SaaS/app marketing essential
3. **With app screenshot** - product showcase
4. **Simple centered with background image** - visual impact
5. **Split with image** - general marketing
