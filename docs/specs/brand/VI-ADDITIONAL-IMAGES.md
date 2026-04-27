# Vi's Additional Image Requirements

**Companion Document to VI-IMAGE-BRIEF.md**
*Written by Vi*

---

Right then. You've got the core brief sorted—error pages, welcomes, empty states. But I've been thinking about the micro-moments where a bit of personality makes all the difference. These are the states between the states, the gentle nudges, the "oh, you noticed that" moments.

This isn't about decoration. Every image here serves a specific interaction, reduces friction, or makes something clearer. British restraint applies throughout—no shouting, no jazz hands.

---

## 1. Interaction States

These live in the spaces between clicks. Hover states, confirmations, the gentle feedback that tells you the system's listening.

### Hover State — Curious Peek

| Attribute | Specification |
|-----------|---------------|
| **Context** | Button hover states, clickable cards, service tiles |
| **Pose** | Leaning forward slightly from edge of frame, looking directly at viewer with curious head tilt. One wing raised as if about to tap the thing you're hovering over. |
| **Expression** | Attentive, "go on then" energy |
| **Size** | 64×64px |
| **Props** | None needed—it's all in the lean and tilt |
| **Mood** | Gentle encouragement without pressure |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style, Duolingo quality.

Pose: Leaning forward from edge of frame, head tilted curiously, one wing raised as if about to tap something. Looking directly at viewer with attentive expression.

Mood: Curious, encouraging, "go on then" energy. British restraint—no over-excitement.

Format: 64x64px, transparent background.
```

![HoverState—CuriousPeek.png](violet/images/interactions/HoverState%E2%80%94CuriousPeek.png)
![HoverState—CuriousPeek2.png](violet/images/interactions/HoverState%E2%80%94CuriousPeek2.png)

---

### Click Confirmation — Gentle Nod

| Attribute | Specification |
|-----------|---------------|
| **Context** | Split-second feedback when user clicks a button |
| **Pose** | Small satisfied nod, eyes briefly closed in approval, one wing giving subtle "thumbs up" (primary feathers extended upward) |
| **Expression** | Quiet satisfaction, confirming action |
| **Size** | 48×48px |
| **Animation** | 2 frames: neutral → nod → back to neutral over 0.3 seconds |
| **Props** | None |
| **Mood** | "Yes, got it" — confirmation without fanfare |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Pose: Small satisfied nod with eyes briefly closed in approval. One wing giving subtle thumbs-up gesture (primary feathers extended upward).

Expression: Quiet satisfaction, confirming user action.

Mood: "Yes, got it" — British confirmation without fanfare.

Format: 48x48px, transparent background. Design for 2-frame animation.
```

---

### Processing — Thinking Gesture

| Attribute | Specification |
|-----------|---------------|
| **Context** | Short processing tasks (3-10 seconds), form submissions |
| **Pose** | Wing to chin in "thinking" pose, looking slightly upward as if considering something. Other wing holding tiny notepad or clipboard. |
| **Expression** | Thoughtful, working it out |
| **Size** | 120×120px |
| **Animation** | Gentle rock or sway, or blinking every 2 seconds |
| **Props** | Tiny notepad/clipboard (optional) |
| **Mood** | Active processing, not passive waiting |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Pose: One wing to chin in thinking gesture, looking slightly upward thoughtfully. Other wing holding tiny notepad or clipboard.

Expression: Thoughtful, actively working it out.

Props: Small notepad or clipboard in wing.

Mood: Active processing, intelligent consideration. British professionalism.

Format: 120x120px, transparent background.
```

---

## 2. Notification & Alert Variants

Toasts and alerts need personality that matches their severity. I'm your friendly messenger here, never alarming.

### Info Notification — Helpful Tap

| Attribute | Specification |
|-----------|---------------|
| **Context** | Informational toasts, "FYI" messages |
| **Pose** | Perched with one wing raised as if pointing to the message text. Friendly, sharing information. Small ℹ️ icon floating nearby. |
| **Expression** | Helpful, "thought you'd like to know" |
| **Size** | 64×64px |
| **Props** | Small info icon (circle with "i") |
| **Mood** | Sharing knowledge, not demanding attention |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Pose: Perched with one wing raised, pointing helpfully toward message area. Small circular info icon floating nearby.

Expression: Friendly, helpful, "thought you'd like to know" energy.

Mood: Sharing information without urgency. Approachable.

Format: 64x64px, transparent background.
```

---

### Warning Notification — Gentle Caution

| Attribute | Specification |
|-----------|---------------|
| **Context** | Warning toasts, non-critical alerts, "are you sure?" confirmations |
| **Pose** | Standing with both wings raised slightly in gentle "hold on" gesture. Not alarmed, just checking. Small amber triangle nearby. |
| **Expression** | Concerned but calm, "just double-checking with you" |
| **Size** | 64×64px |
| **Props** | Small amber warning triangle (not red—we're not panicking) |
| **Colour accent** | Amber/orange for triangle, keeps Vi purple |
| **Mood** | Cautious friend, not alarm system |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Pose: Standing with both wings raised slightly in gentle "hold on" gesture. Small amber/orange warning triangle floating nearby.

Expression: Concerned but calm, "just double-checking" energy.

Mood: Cautious friend, not alarm. British restraint—no panic.

Colours: Purple Vi with amber/orange accent for warning symbol.

Format: 64x64px, transparent background.
```

---

### Success Notification — Quiet Pride

| Attribute | Specification |
|-----------|---------------|
| **Context** | Success toasts, confirmations, completed actions |
| **Pose** | Small satisfied hop or standing tall with one wing giving subtle thumbs-up. Green checkmark nearby. British celebration (no confetti cannons). |
| **Expression** | Pleased, proud of the user |
| **Size** | 64×64px |
| **Props** | Small green checkmark or tick |
| **Mood** | "Well done" — contained joy |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Pose: Small satisfied hop or standing tall, one wing giving subtle thumbs-up. Small green checkmark floating nearby.

Expression: Pleased, proud of the user's success.

Mood: "Well done" — British celebration, contained joy. No over-the-top excitement.

Colours: Purple Vi with green checkmark accent.

Format: 64x64px, transparent background.
```

---

### Error Notification — Sympathetic Shrug

| Attribute | Specification |
|-----------|---------------|
| **Context** | Gentle error messages, failed validations, "that didn't work" moments |
| **Pose** | Small apologetic shrug with both wings, sympathetic head tilt. Small red X or cross nearby (muted red, not screaming). |
| **Expression** | "Sorry, that didn't work" — empathetic, not judgemental |
| **Size** | 64×64px |
| **Props** | Muted red X or cross symbol |
| **Mood** | Gentle apology, "let's try again together" |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Pose: Small apologetic shrug with both wings raised, sympathetic head tilt. Small muted red X or cross floating nearby.

Expression: "Sorry, that didn't work" — empathetic, understanding, not judgemental.

Mood: Gentle apology, never scolding. "Let's try again together."

Colours: Purple Vi with muted red (not screaming red) accent for error symbol.

Format: 64x64px, transparent background.
```

---

## 3. Feature Discovery & Tooltips

These pop up when you're learning the system. I'm your patient guide here.

### "Did You Know?" Tooltip

| Attribute | Specification |
|-----------|---------------|
| **Context** | Feature discovery tooltips, helpful hints, onboarding tips |
| **Pose** | Leaning in from edge of tooltip bubble, one wing pointing to the feature being explained. Small lightbulb icon above head. |
| **Expression** | Enthusiastic teacher, "let me show you this neat thing" |
| **Size** | 80×80px |
| **Props** | Small glowing lightbulb (soft yellow, not harsh) |
| **Mood** | Helpful discovery, not annoying tutorial |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Pose: Leaning in from edge of frame, one wing pointing helpfully to demonstrate a feature. Small glowing lightbulb above head.

Expression: Enthusiastic but restrained, "let me show you this" energy. Helpful teacher.

Props: Soft yellow glowing lightbulb (not harsh bright).

Mood: Feature discovery, patient guide. Never annoying.

Format: 80x80px, transparent background suitable for tooltip overlays.
```

---

### New Feature Callout

| Attribute | Specification |
|-----------|---------------|
| **Context** | Announcing new features, changelog highlights, "what's new" sections |
| **Pose** | Presenting with both wings gesturing toward the new feature. Subtle sparkle or star near one wing. Standing proud but not shouting. |
| **Expression** | Proud to show you, "built this for you" |
| **Size** | 200×150px |
| **Props** | Single subtle sparkle or "new" badge (understated) |
| **Mood** | Quiet pride in craftsmanship |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Pose: Standing proud, both wings gesturing toward feature area as if presenting. Single subtle sparkle or star near one wing.

Expression: Proud but understated, "built this for you" energy.

Props: Single subtle sparkle (not multiple—British restraint). Optional small "new" badge.

Mood: Quiet pride in craftsmanship. Not shouting about features.

Format: 200x150px, transparent background.
```

---

## 4. Emotional Range

Sometimes you just need to know I understand what you're feeling.

### Confused — Puzzled Tilt

| Attribute | Specification |
|-----------|---------------|
| **Context** | Help modals, unclear states, "what's this?" moments |
| **Pose** | Head tilted nearly 90 degrees (classic confused bird), one wing scratching head gently. Small question mark floating. |
| **Expression** | Genuinely puzzled but not distressed |
| **Size** | 120×120px |
| **Props** | Small question mark |
| **Mood** | "I'm a bit lost too, let's figure this out" |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Pose: Head tilted nearly 90 degrees (classic confused bird pose), one wing gently scratching head. Small question mark floating nearby.

Expression: Genuinely puzzled but not distressed. Relatable confusion.

Mood: "I'm a bit lost too, let's figure this out together."

Format: 120x120px, transparent background.
```

---

### Eureka Moment — Discovery

| Attribute | Specification |
|-----------|---------------|
| **Context** | Successful troubleshooting, solutions found, lightbulb moments |
| **Pose** | Wings spread in moment of realisation, looking up with brightened expression. Single sparkle above head. |
| **Expression** | "Ah, there it is" — satisfying discovery |
| **Size** | 150×150px |
| **Props** | Single sparkle or glowing lightbulb |
| **Mood** | Problem-solving satisfaction |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Pose: Wings spread in moment of realisation, looking upward with brightened expression. Single sparkle or glowing lightbulb above head.

Expression: "Ah, there it is" — satisfying discovery, eureka moment.

Mood: Problem-solving satisfaction. British restraint—one sparkle only.

Format: 150x150px, transparent background.
```

---

### Sympathetic Listener — Understanding

| Attribute | Specification |
|-----------|---------------|
| **Context** | Error states, support forms, "tell us what happened" |
| **Pose** | Perched attentively, both wings folded, leaning forward slightly. Direct eye contact, full attention. |
| **Expression** | "I'm listening, I understand" — empathetic |
| **Size** | 120×120px |
| **Props** | None needed—it's all in the posture and eye contact |
| **Mood** | Patient support, genuine care |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Pose: Perched attentively with both wings folded, leaning forward slightly. Direct eye contact with viewer, giving full attention.

Expression: "I'm listening, I understand" — empathetic, patient.

Mood: Genuine care and support. Therapist energy (but still a raven).

Format: 120x120px, transparent background.
```

---

### Celebrating Together — Shared Success

| Attribute | Specification |
|-----------|---------------|
| **Context** | Milestone achievements, first published post, account milestones |
| **Pose** | Small hop with wings slightly spread, looking at viewer with genuine shared joy. Tiny party hat (single colour, not garish). |
| **Expression** | "We did it" — inclusive celebration |
| **Size** | 200×200px |
| **Props** | Single-colour tiny party hat (purple or gold) |
| **Mood** | British celebration—genuinely pleased but not American-level enthusiasm |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Pose: Small celebratory hop with wings slightly spread, looking at viewer with genuine shared joy. Wearing tiny simple party hat (single colour—purple or gold, not striped garish).

Expression: "We did it together" — inclusive celebration.

Mood: British celebration level—genuinely pleased but restrained. No confetti cannons.

Props: Simple tiny party hat.

Format: 200x200px, transparent background.
```

---

## 5. Seasonal & Event Variants

The brief mentioned these. Here are the specific prompts.

### Winter Vi

| Attribute | Specification |
|-----------|---------------|
| **Context** | December through February UI variant |
| **Pose** | Standard friendly perch, wearing small knitted scarf (purple/gold striped). Holding tiny steaming mug in wing. |
| **Expression** | Cosy, content |
| **Size** | 400×400px (profile variant) |
| **Props** | Small knitted scarf, steaming mug of tea |
| **Mood** | Warm despite the cold, hygge energy |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Pose: Standard friendly perch, wearing small knitted scarf in purple and gold stripes. Holding tiny steaming mug of tea in one wing.

Expression: Cosy, content, warm despite winter.

Props: Knitted scarf (purple/gold), small mug with steam rising.

Mood: Hygge energy, British winter comfort. Understated cosiness.

Season: Winter variant for December-February.

Format: 400x400px, transparent or subtle winter background.
```

---

### Spring Vi

| Attribute | Specification |
|-----------|---------------|
| **Context** | March through May UI variant |
| **Pose** | Standard perch with tiny flower crown (violet flowers, naturally). Small butterfly nearby. |
| **Expression** | Fresh, renewed, hopeful |
| **Size** | 400×400px |
| **Props** | Delicate flower crown (violet/purple blooms), single butterfly |
| **Mood** | New beginnings, gentle optimism |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Pose: Standard friendly perch wearing tiny delicate flower crown made of violet/purple flowers. Small butterfly nearby.

Expression: Fresh, renewed, hopeful.

Props: Flower crown (violet blooms), single butterfly.

Mood: Spring renewal, gentle optimism. British springtime.

Season: Spring variant for March-May.

Format: 400x400px, transparent or subtle spring background.
```

---

### Summer Vi

| Attribute | Specification |
|-----------|---------------|
| **Context** | June through August UI variant |
| **Pose** | Perched with small round sunglasses (gold frames), holding tiny ice cream cone (purple/lavender flavour). |
| **Expression** | Relaxed, enjoying summer |
| **Size** | 400×400px |
| **Props** | Small round sunglasses (gold frames), tiny ice cream cone |
| **Mood** | British summer—pleasant but not scorching |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Pose: Perched wearing small round sunglasses with gold frames. Holding tiny ice cream cone in one wing (purple/lavender flavour).

Expression: Relaxed, enjoying British summer.

Props: Round sunglasses (gold frames), small ice cream cone (purple colour).

Mood: Pleasant summer day. British summer—nice but not too hot.

Season: Summer variant for June-August.

Format: 400x400px, transparent or subtle summer background.
```

---

### Autumn Vi

| Attribute | Specification |
|-----------|---------------|
| **Context** | September through November UI variant |
| **Pose** | Perched surrounded by gently falling leaves (amber, gold, brown). One wing catching a falling leaf. |
| **Expression** | Peaceful, contemplative |
| **Size** | 400×400px |
| **Props** | Falling autumn leaves in warm colours |
| **Mood** | British autumn—crisp, beautiful, slightly melancholic |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Pose: Perched with one wing raised, gently catching a falling autumn leaf. Surrounded by falling leaves in amber, gold, and brown tones.

Expression: Peaceful, contemplative.

Props: Falling autumn leaves (warm colours—amber, gold, brown).

Mood: British autumn—crisp air, beautiful but slightly melancholic.

Season: Autumn variant for September-November.

Format: 400x400px, transparent or subtle autumn background.
```

---

### Christmas Vi (Restrained)

| Attribute | Specification |
|-----------|---------------|
| **Context** | Mid-December through Christmas Day |
| **Pose** | Standard perch wearing single red Santa hat (not garish, properly fitted). Small wrapped gift beside. |
| **Expression** | Festive but dignified |
| **Size** | 400×400px |
| **Props** | Single red Santa hat (fits properly, not comically oversized), small wrapped gift (purple paper, gold bow) |
| **Mood** | British Christmas—festive but not over-the-top |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Pose: Standard friendly perch wearing single red Santa hat (properly fitted, not comically oversized). Small wrapped gift sitting beside (purple paper, gold bow).

Expression: Festive but dignified. British Christmas spirit.

Props: Red Santa hat (fitted properly), small wrapped gift.

Mood: British Christmas—festive, warm, but restrained. Not American-level Christmas enthusiasm.

Season: Christmas variant for mid-December through 25th.

Format: 400x400px, transparent or subtle winter background.
```

---

## 6. Service-Specific Variants (Extended)

The brief covered SocialHost and AnalyticsHost. Here are the others.

### BioHost Vi

| Attribute | Specification |
|-----------|---------------|
| **Context** | BioHost (bio link pages) service |
| **Pose** | Presenting or gesturing toward a stylised bio page mockup. Chain link icons connecting elements. Standing proud of the creation. |
| **Expression** | Artistic, proud designer |
| **Size** | 600×400px |
| **Props** | Abstract bio page mockup, chain link symbols, artistic flourishes |
| **Colour accents** | Purple with link chain silver/grey |
| **Mood** | Creative professional showing their portfolio |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Context: BioHost service (bio link pages).

Pose: Standing beside or presenting stylised bio page mockup. Chain link icons connecting page elements. One wing gesturing proudly toward creation.

Expression: Artistic, proud designer showing their work.

Props: Abstract bio page mockup (simple webpage shape), chain link symbols (silver/grey), artistic flourishes.

Mood: Creative professional, portfolio presentation energy.

Colours: Purple Vi with silver/grey chain link accents.

Format: 600x400px, suitable for service branding.
```

---

### TrustHost Vi

| Attribute | Specification |
|-----------|---------------|
| **Context** | TrustHost (social proof widgets) service |
| **Pose** | Wing-thumbs-up surrounded by floating five-star reviews, quote bubbles, and trust badges. Confident, endorsing. |
| **Expression** | Trustworthy, "you can trust this" |
| **Size** | 600×400px |
| **Props** | Five gold stars, quote bubbles with testimonials, trust badge icons |
| **Colour accents** | Purple with gold stars |
| **Mood** | Building credibility, authentic endorsement |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Context: TrustHost service (social proof and review widgets).

Pose: Wing giving thumbs-up gesture, surrounded by floating five-star reviews, testimonial quote bubbles, and trust badge icons.

Expression: Trustworthy, confident endorsement. "You can trust this."

Props: Five gold stars, quote bubbles (with abstract text lines), trust badge icons (shields, checkmarks).

Mood: Building credibility through authentic endorsement. Reliable.

Colours: Purple Vi with gold star accents.

Format: 600x400px, suitable for service branding.
```

---

### NotifyHost Vi

| Attribute | Specification |
|-----------|---------------|
| **Context** | NotifyHost (push notification service) |
| **Pose** | Gently ringing small notification bell with one wing. Other wing welcoming gesture. Return arrow symbol nearby (bringing users back). |
| **Expression** | Friendly reminder, gentle nudge |
| **Size** | 600×400px |
| **Props** | Small notification bell (gold), return/back arrow icon |
| **Colour accents** | Purple with gold bell |
| **Mood** | Helpful reminder without being annoying |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Context: NotifyHost service (push notifications, re-engagement).

Pose: Gently ringing small notification bell with one wing. Other wing in welcoming gesture. Return/back arrow symbol floating nearby.

Expression: Friendly reminder, gentle nudge. Never annoying.

Props: Small gold notification bell, return arrow icon.

Mood: Helpful reminder that brings users back without being pushy.

Colours: Purple Vi with gold bell accent.

Format: 600x400px, suitable for service branding.
```

---

### MailHost Vi

| Attribute | Specification |
|-----------|---------------|
| **Context** | MailHost (email service) |
| **Pose** | Postal worker vibe—sorting tiny envelopes efficiently. Small postal bag or sorting tray nearby. @ symbol floating. |
| **Expression** | Organised, reliable mail carrier |
| **Size** | 600×400px |
| **Props** | Tiny envelopes, postal bag or sorting tray, @ symbol |
| **Colour accents** | Purple with classic postal blue/red trim (subtle) |
| **Mood** | British postal reliability—"it'll get there" |

**Google Whisk Prompt:**
```
A friendly cartoon raven mascot named Vi. Royal purple feathers (#663399), large expressive eyes with golden highlights, orange-gold beak. Clean modern vector style.

Context: MailHost service (email platform).

Pose: Postal worker vibe—standing with small postal bag or sorting through tiny envelopes. @ symbol floating nearby.

Expression: Organised, reliable mail carrier. "It'll get there."

Props: Tiny envelopes (various colours), small postal bag or sorting tray, @ symbol.

Mood: British postal reliability. Royal Mail energy but modern.

Colours: Purple Vi with subtle postal blue/red trim accents.

Format: 600x400px, suitable for service branding.
```

---

## Priority & Implementation

### Phase 1 (High Value, Low Effort)

1. **Notification variants** (info, warning, success, error) — 64×64px each
2. **Hover state** (curious peek) — 64×64px
3. **Click confirmation** (gentle nod) — 48×48px, 2 frames

These add personality to every interaction without major design lift.

### Phase 2 (Feature Discovery)

4. **Tooltip "Did You Know?"** — 80×80px
5. **New feature callout** — 200×150px
6. **Confused state** — 120×120px

Help users learn the system naturally.

### Phase 3 (Emotional Range)

7. **Eureka moment** — 150×150px
8. **Sympathetic listener** — 120×120px
9. **Celebrating together** — 200×200px

Build emotional connection during key moments.

### Phase 4 (Seasonal Rotation)

10. **Seasonal variants** — 400×400px each (5 total: winter, spring, summer, autumn, Christmas)

Refresh the dashboard throughout the year.

### Phase 5 (Service Branding)

11. **Service-specific variants** — 600×400px each (BioHost, TrustHost, NotifyHost, MailHost)

Complete the service family branding.

---

## Technical Notes

### Character Consistency with Google Whisk

When using Google Whisk's character creator:

1. **Create master character first**: Upload the existing Vi reference image and create the base character
2. **Use character reference**: Apply the character to all subsequent prompts
3. **Specify pose variations**: The character stays consistent whilst poses change
4. **Maintain colour accuracy**: Emphasise #663399 purple in every prompt
5. **Export high resolution**: Generate at 4x final size, then scale down

### Animation Frames

For states requiring animation:
- **Click confirmation**: 2 frames (neutral → nod)
- **Loading tea**: 3 frames (page turning)
- **Thinking**: Subtle sway or blink every 2 seconds

Export each frame separately, then implement in CSS or JavaScript.

### Accessibility Considerations

Every image needs:
- **Alt text** written in my voice
- **Sufficient contrast** against both light and dark backgrounds
- **Dark mode variant** where background isn't transparent
- **Meaningful content** (never "decorative")

---

## Why These Matter

You might think micro-interactions don't need personality. But that's precisely where personality matters most. When someone hovers over a button and I lean in curiously, that split-second feedback says "the system is alive and paying attention."

When a post fails to publish and I give a sympathetic shrug instead of a cold error icon, that's the difference between frustration and "alright, let's try again."

Every moment is a conversation. These images are my half of it.

---

*Generate these in phases as capacity allows. Consistency matters more than speed—I'd rather wait for proper Vi than get rushed clip art with a beak.*

*And yes, I'll have that cuppa whilst we work through these.*

— Vi
