# RFC: HLCRF Compositor

**Status:** Implemented
**Created:** 2026-01-15
**Authors:** Host UK Engineering

---

## Abstract

The HLCRF Compositor is a hierarchical layout system where each composite contains up to five regions—Header, Left, Content, Right, and Footer. Composites nest infinitely: any region can contain another composite, which can contain another, and so on.

The core innovation is **inline sub-structure declaration**: a single string like `H[LC]C[HCF]F` declares the entire nested hierarchy. No configuration files, no database schema, no separate definitions—parse the string and you have the complete structure.

Just as Markdown made document formatting a human-readable string, HLCRF makes layout structure a portable, self-describing data type that can be stored, transmitted, validated, and rendered anywhere.

Path-based element IDs (`L-H-0`, `C-F-C-2`) encode the full hierarchy, eliminating database lookups to resolve structure. The system supports responsive breakpoints, block-based content, and shortcode integration.

---

## Motivation

Traditional layout systems require separate templates for each layout variation. A page with a left sidebar needs one template; without it, another. Add responsive behaviour, and the combinations multiply quickly.

The HLCRF Compositor addresses this through:

1. **Data-driven layouts** — A single compositor handles all layout permutations via variant strings
2. **Nested composition** — Layouts can contain other layouts, with automatic path tracking for unique identification
3. **Responsive design** — Breakpoint-aware rendering collapses regions appropriately for different devices
4. **Block-based content** — Content populates regions as discrete blocks, enabling conditional display and reordering

The approach treats layout as data rather than markup, allowing the same content to adapt to different structural requirements without template duplication.

---

## Terminology

### HLCRF

**H**ierarchical **L**ayer **C**ompositing **R**ender **F**rame.

The acronym also serves as a mnemonic for the five possible regions:

| Letter | Region | Semantic element | Purpose |
|--------|--------|------------------|---------|
| **H** | Header | `<header>` | Top navigation, branding |
| **L** | Left | `<aside>` | Left sidebar, secondary navigation |
| **C** | Content | `<main>` | Primary content area |
| **R** | Right | `<aside>` | Right sidebar, supplementary content |
| **F** | Footer | `<footer>` | Site footer, links, legal |

### Variant string

A string of 1–5 characters from the set `{H, L, C, R, F}` that defines which regions are active. The string `HCF` produces a layout with Header, Content, and Footer. The string `HLCRF` enables all five regions.

A flat variant like `HLCRF` renders as:

```
┌─────────────────────────────────────┐
│                  H                  │  ← Header
├─────────┬───────────────┬───────────┤
│    L    │       C       │     R     │  ← Body row
├─────────┴───────────────┴───────────┤
│                  F                  │  ← Footer
└─────────────────────────────────────┘
```

A nested variant like `H[LCR]CF` renders differently—the body row is **inside** the Header:

```
┌─────────────────────────────────────┐
│ H ┌─────────┬─────────┬───────────┐ │
│   │   H-L   │   H-C   │    H-R    │ │  ← Body row nested IN Header
│   └─────────┴─────────┴───────────┘ │
├─────────────────────────────────────┤
│                  C                  │  ← Root Content
├─────────────────────────────────────┤
│                  F                  │  ← Root Footer
└─────────────────────────────────────┘
```

With blocks placed, element IDs become addresses. A typical website header `H[LCR]CF`:

```
┌───────────────────────────────────────────────────────────────┐
│ H ┌───────────┬───────────────────────────────┬─────────────┐ │
│   │   H-L-0   │  H-C-0   H-C-1   H-C-2   H-C-3│    H-R-0    │ │
│   │   [Logo]  │  [Home]  [About] [Blog] [Shop]│   [Login]   │ │
│   └───────────┴───────────────────────────────┴─────────────┘ │
├───────────────────────────────────────────────────────────────┤
│                            C-0                                │
│                       [Page Content]                          │
├───────────────────────────────────────────────────────────────┤
│              F-0                         F-1                  │
│           [© 2026]                    [Legal]                 │
└───────────────────────────────────────────────────────────────┘
```

Every element has a unique, deterministic address. Computers count from zero—deal with it.

**Key principle:** What's missing defines the layout type. Brackets define nesting.

### Path

A hierarchical identifier tracking a layout's position within nested structures. The root layout has an empty path. A layout nested within the Left region of the root receives path `L-`. Further nesting appends to this path.

### Slot

A named region within a layout that accepts content. Each slot corresponds to one HLCRF letter.

### Block

A discrete unit of content assigned to a region. Blocks have their own ordering and can conditionally display based on breakpoint or other conditions.

### Breakpoint

A device category determining layout behaviour:

| Breakpoint | Target | Typical behaviour |
|------------|--------|-------------------|
| `phone` | < 768px | Single column, stacked |
| `tablet` | 768px–1023px | Content only, sidebars hidden |
| `desktop` | ≥ 1024px | Full layout with all regions |

---

## Specification

### Layout variant strings

#### Valid variants

Any combination of the letters H, L, C, R, F, in that order. Common variants:

| Variant | Description | Use case |
|---------|-------------|----------|
| `C` | Content only | Embedded widgets, minimal layouts |
| `HCF` | Header, Content, Footer | Standard page layout |
| `HCR` | Header, Content, Right | Dashboard with right sidebar |
| `HLC` | Header, Left, Content | Admin panel with navigation |
| `HLCF` | Header, Left, Content, Footer | Admin with footer |
| `HLCR` | Header, Left, Content, Right | Three-column dashboard |
| `HLCRF` | All regions | Full-featured layouts |

The variant string is case-insensitive. The compositor normalises to uppercase.

#### Inline sub-structure declaration

Variant strings support **inline nesting** using bracket notation. Each region letter can be followed by brackets containing its nested layout:

```
H[LC]L[HC]C[HCF]F[LCF]
```

This declares the entire hierarchy in a single string:

| Segment | Meaning |
|---------|---------|
| `H[LC]` | Header region contains a Left-Content layout |
| `L[HC]` | Left region contains a Header-Content layout |
| `C[HCF]` | Content region contains a Header-Content-Footer layout |
| `F[LCF]` | Footer region contains a Left-Content-Footer layout |

Brackets nest recursively. A complex declaration like `H[L[C]C]CF` means:
- Header contains a nested layout
- That nested layout's Left region contains yet another layout (Content-only)
- Root also has Content and Footer at the top level

This syntax is particularly useful for:
- **Shortcodes** declaring their expected structure
- **Templates** defining reusable page scaffolds
- **Configuration** specifying layout contracts

The string `H[LC]L[HC]C[HCF]F[LCF]` is a complete website declaration—no additional nesting configuration needed.

#### Region requirements

- **Content (C)** is implicitly included when any body region (L, C, R) is present
- Regions render only when the variant includes them AND content has been added
- An empty region does not render, even if specified in the variant

### Region hierarchy

The compositor enforces a fixed spatial hierarchy:

```
Row 1: Header (full width)
Row 2: Left | Content | Right (body row)
Row 3: Footer (full width)
```

This structure maps to CSS Grid areas:

```css
grid-template-areas:
    "header"
    "body"
    "footer";
```

The body row uses a nested grid or flexbox for the three-column layout.

### Nesting and path context

Layouts can be nested within any region. The compositor automatically manages path context to ensure unique slot identifiers.

#### Path generation

When a layout renders, it assigns each slot an ID based on its path:

- Root layout, Header slot: `H`
- Root layout, Left slot: `L`
- Nested layout within Left, Header slot: `L-H`
- Nested layout within Left, Content slot: `L-C`
- Further nested within that Content slot: `L-C-C`

#### Block identifiers

Within each slot, blocks receive indexed identifiers:

- First block in Header: `H-0`
- Second block in Header: `H-1`
- First block in nested Content: `L-C-0`

This scheme enables precise targeting for styling, JavaScript, and debugging.

### Responsive breakpoints

The compositor supports breakpoint-specific layout variants. A page might use `HLCRF` on desktop but collapse to `HCF` on tablet and `C` on phone.

#### Configuration schema

```json
{
    "layout_config": {
        "layout_type": {
            "desktop": "HLCRF",
            "tablet": "HCF",
            "phone": "CF"
        },
        "regions": {
            "desktop": {
                "left": { "width": 280 },
                "content": { "max_width": 680 },
                "right": { "width": 280 }
            }
        }
    }
}
```

#### CSS breakpoint handling

The default CSS collapses sidebars at tablet breakpoint and stacks content at phone breakpoint:

```css
/* Tablet: Hide sidebars */
@media (max-width: 1023px) {
    .hlcrf-body {
        grid-template-columns: minmax(0, var(--content-max-width));
        grid-template-areas: "content";
    }
    .hlcrf-left, .hlcrf-right { display: none; }
}

/* Phone: Full width, stacked */
@media (max-width: 767px) {
    .hlcrf-body {
        grid-template-columns: 1fr;
        padding: 0 1rem;
    }
}
```

### Block visibility

Blocks can define per-breakpoint visibility:

```json
{
    "breakpoint_visibility": {
        "desktop": true,
        "tablet": true,
        "phone": false
    }
}
```

A block with `phone: false` does not render on mobile devices, regardless of which region it belongs to.

### Deep nesting

The HLCRF system is **infinitely nestable**. Any region can contain another complete HLCRF layout, which can itself contain further nested layouts. The path-based ID scheme ensures every element remains uniquely addressable regardless of nesting depth.

#### Path reading convention

Paths read left-to-right, describing the journey from root to element:

```
L-H-0
│ │ └─ Block index (first block)
│ └─── Region in nested layout (Header)
└───── Region in root layout (Left)
```

This means: "The first block in the Header region of a layout nested within the Left region of the root."

#### Multi-level path construction

Paths concatenate as layouts nest. Consider this structure:

- Root layout: `HLCRF`
- Nested in Content: another `HCF` layout
- Nested in that layout's Footer: a `C`-only layout with a button block

The button receives the path: `C-F-C-0`

Reading left to right:
1. `C` — Content region of root
2. `F` — Footer region of nested layout
3. `C` — Content region of deepest layout
4. `0` — First block in that region

#### Three-level nesting example

```
┌─────────────────────────────────────────────────────────┐
│ H (root header)                                         │
├────────┬────────────────────────────────────┬───────────┤
│   L    │                C                   │     R     │
│        │  ┌───────────────────────────────┐ │           │
│        │  │ C-H (nested header)           │ │           │
│        │  ├─────┬─────────────┬───────────┤ │           │
│        │  │ C-L │    C-C      │    C-R    │ │           │
│        │  │     │ ┌─────────┐ │           │ │           │
│        │  │     │ │ C-C-C   │ │           │ │           │
│        │  │     │ │(deepest)│ │           │ │           │
│        │  │     │ └─────────┘ │           │ │           │
│        │  ├─────┴─────────────┴───────────┤ │           │
│        │  │ C-F (nested footer)           │ │           │
│        │  └───────────────────────────────┘ │           │
├────────┴────────────────────────────────────┴───────────┤
│ F (root footer)                                         │
└─────────────────────────────────────────────────────────┘
```

In this diagram:
- Root regions: `H`, `L`, `C`, `R`, `F`
- Second level (nested in C): `C-H`, `C-L`, `C-C`, `C-R`, `C-F`
- Third level (nested in C-C): `C-C-C`

A block placed in the deepest Content region would receive ID `C-C-C-0`.

#### Path examples at each nesting level

| Nesting depth | Example path | Meaning |
|---------------|--------------|---------|
| 1 (root) | `H-0` | First block in root Header |
| 1 (root) | `L-2` | Third block in root Left sidebar |
| 2 (nested) | `L-H-0` | First block in Header of layout nested in Left |
| 2 (nested) | `C-C-1` | Second block in Content of layout nested in Content |
| 3 (deep) | `L-C-H-0` | First block in Header of layout nested in Content, nested in Left |
| 4+ | `C-L-C-R-0` | Paths continue indefinitely |

The path length equals the nesting depth plus one (for the block index).

#### Practical example: sidebar with nested layout

```php
$sidebar = Layout::make('HCF')
    ->h('<h3>Widget Panel</h3>')
    ->c(view('widgets.list'))
    ->f('<a href="#">Manage widgets</a>');

$page = Layout::make('HLCRF')
    ->h(view('header'))
    ->l($sidebar)  // Nested layout in Left
    ->c(view('main-content'))
    ->f(view('footer'));
```

The sidebar's regions receive paths:
- Header: `L-H`
- Content: `L-C`
- Footer: `L-F`

Blocks within the sidebar's Content would be `L-C-0`, `L-C-1`, etc.

#### Why infinite nesting matters

Deep nesting enables:

1. **Component encapsulation** — A reusable component can define its own internal layout without knowing where it will be placed
2. **Recursive structures** — Tree views, nested comments, or hierarchical navigation can use consistent layout patterns at each level
3. **Micro-layouts** — Small UI sections (cards, panels, modals) can use HLCRF internally whilst remaining composable

---

## API reference

### `Layout` class

**Namespace:** `Core\Front\Components`

#### Factory method

```php
Layout::make(string $variant = 'HCF', string $path = ''): static
```

Creates a new layout instance.

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `$variant` | string | `'HCF'` | Layout variant string |
| `$path` | string | `''` | Hierarchical path (typically managed automatically) |

#### Slot methods

Each region has a variadic method accepting any renderable content:

```php
public function h(mixed ...$items): static  // Header
public function l(mixed ...$items): static  // Left
public function c(mixed ...$items): static  // Content
public function r(mixed ...$items): static  // Right
public function f(mixed ...$items): static  // Footer
```

Alias methods provide readability for explicit code:

```php
public function addHeader(mixed ...$items): static
public function addLeft(mixed ...$items): static
public function addContent(mixed ...$items): static
public function addRight(mixed ...$items): static
public function addFooter(mixed ...$items): static
```

#### Content types

Slot methods accept:

- **Strings** — Raw HTML or text
- **`Htmlable`** — Objects implementing `toHtml()`
- **`Renderable`** — Objects implementing `render()`
- **`View`** — Laravel view instances
- **`Layout`** — Nested layout instances (path context injected automatically)
- **Callables** — Functions returning any of the above

#### Attribute methods

```php
public function attributes(array $attributes): static
```

Merge HTML attributes onto the layout container.

```php
public function class(string $class): static
```

Append a CSS class to the container.

#### Rendering

```php
public function render(): string
public function toHtml(): string
public function __toString(): string
```

All three methods return the compiled HTML. The class implements `Htmlable` and `Renderable` for framework integration.

---

## Examples

### Basic page layout

```php
use Core\Front\Components\Layout;

$page = Layout::make('HCF')
    ->h(view('components.header'))
    ->c('<article>Page content here</article>')
    ->f(view('components.footer'));

echo $page;
```

### Admin dashboard with sidebar

```php
$dashboard = Layout::make('HLCF')
    ->class('min-h-screen bg-gray-100')
    ->h(view('admin.header'))
    ->l(view('admin.sidebar'))
    ->c($content)
    ->f(view('admin.footer'));
```

### Nested layouts

```php
// Outer layout with left sidebar
$outer = Layout::make('HLC')
    ->h('<nav>Main Navigation</nav>')
    ->l('<aside>Sidebar</aside>')
    ->c(
        // Inner layout nested in content area
        Layout::make('HCF')
            ->h('<h1>Section Title</h1>')
            ->c('<div>Inner content</div>')
            ->f('<p>Section footer</p>')
    );
```

The inner layout receives path context `C-`, so its slots become `C-H`, `C-C`, and `C-F`.

### Multiple blocks per region

```php
$page = Layout::make('HLCF')
    ->h(view('header.logo'), view('header.navigation'), view('header.search'))
    ->l(view('sidebar.menu'), view('sidebar.widgets'))
    ->c(view('content.hero'), view('content.features'), view('content.cta'))
    ->f(view('footer.links'), view('footer.legal'));
```

Each item becomes a separate block with a unique identifier.

### Responsive rendering

```php
// In a service or controller
$breakpoint = $this->detectBreakpoint($request);
$layoutType = $page->getLayoutTypeFor($breakpoint);

$layout = Layout::make($layoutType)
    ->class('bio-page')
    ->h($headerBlocks)
    ->c($contentBlocks)
    ->f($footerBlocks);

// Sidebars only added on desktop
if ($breakpoint === 'desktop') {
    $layout->l($leftBlocks)->r($rightBlocks);
}
```

---

## Implementation notes

### CSS Grid structure

The compositor generates a grid-based structure:

```html
<div class="hlcrf-layout" data-layout="root">
    <header class="hlcrf-header" data-slot="H">...</header>
    <div class="hlcrf-body flex flex-1">
        <aside class="hlcrf-left shrink-0" data-slot="L">...</aside>
        <main class="hlcrf-content flex-1" data-slot="C">...</main>
        <aside class="hlcrf-right shrink-0" data-slot="R">...</aside>
    </div>
    <footer class="hlcrf-footer" data-slot="F">...</footer>
</div>
```

The base CSS uses CSS Grid for the outer structure and Flexbox for the body row.

### Semantic HTML

The compositor uses appropriate semantic elements:

- `<header>` for the Header region
- `<aside>` for Left and Right sidebars
- `<main>` for the Content region
- `<footer>` for the Footer region

This provides accessibility benefits and proper document outline.

### Accessibility considerations

- **Landmark regions** — Semantic elements create implicit ARIA landmarks
- **Skip links** — Consider adding skip-to-content links in the Header
- **Focus management** — Nested layouts maintain sensible tab order
- **Screen reader compatibility** — Block `data-block` attributes aid debugging but do not affect accessibility tree

### Data attributes

The compositor adds data attributes for debugging and JavaScript integration:

| Attribute | Location | Purpose |
|-----------|----------|---------|
| `data-layout` | Container | Layout path identifier (`root`, `L-`, etc.) |
| `data-slot` | Region | Slot identifier (`H`, `L-C`, etc.) |
| `data-block` | Block wrapper | Block identifier (`H-0`, `L-C-2`, etc.) |

### Database schema

For persisted layouts, the schema includes:

**Pages/Biolinks table:**
- `layout_config` (JSON) — Layout type per breakpoint, region dimensions

**Blocks table:**
- `region` (string) — Target region: `header`, `left`, `content`, `right`, `footer`
- `region_order` (integer) — Sort order within region
- `breakpoint_visibility` (JSON) — Per-breakpoint visibility flags

### Theme integration

The renderer can generate CSS custom properties for theming:

```css
:root {
    --biolink-bg: #f9fafb;
    --biolink-text: #111827;
    --biolink-font: 'Inter', system-ui, sans-serif;
}
```

These integrate with the compositor's class-based styling.

---

## Integration patterns

This section describes how the HLCRF system integrates with common web development patterns and technologies.

### CSS box model parallel

HLCRF mirrors the CSS box model conceptually, which aids developer intuition:

```
CSS Box Model          HLCRF Layout
┌──────────────┐       ┌──────────────┐
│   margin     │       │      H       │  ← Block-level, full-width
├──────────────┤       ├──────────────┤
│ │ padding  │ │       │ L │  C  │ R  │  ← Content row with "sidebars"
├──────────────┤       ├──────────────┤
│   margin     │       │      F       │  ← Block-level, full-width
└──────────────┘       └──────────────┘
```

The mapping:
- **H/F** behave like block-level elements spanning the full width, similar to how top/bottom margins frame content
- **L/R** act as the "padding" on either side of the content, creating gutters or sidebars
- **C** is the content itself—the innermost box

This mental model helps developers predict layout behaviour:
- Adding `L` or `R` is like adding horizontal padding
- Adding `H` or `F` is like adding vertical margins
- The `[LCR]` row always forms the content layer, with `C` as the primary content area

When nesting layouts, the analogy extends recursively—a nested layout's `H/F` become block-level elements within their parent region, and its `[LCR]` row subdivides that space further.

### Shortcode structure definitions

HLCRF enables shortcodes to define complete structural layouts through their variant string. The variant becomes a **structural contract** that the shortcode guarantees to fulfil.

#### Example: Hero shortcode

```php
// Shortcode definition
class HeroShortcode extends Shortcode
{
    public string $layout = 'HCF';

    public function render(): Layout
    {
        return Layout::make($this->layout)
            ->h($this->renderTitle())      // H: Title region
            ->c($this->renderContent())    // C: Main hero content
            ->f($this->renderCta());       // F: Call-to-action region
    }
}
```

Usage in content:

```
[hero layout="HCF" title="Welcome" cta="Get Started"]
    Your hero content here.
[/hero]
```

The shortcode author declares which regions exist; content authors populate them. The variant string serves as documentation and constraint simultaneously.

#### Variant as capability declaration

Different shortcode variants expose different capabilities:

| Shortcode | Variant | Regions | Purpose |
|-----------|---------|---------|---------|
| `[hero]` | `HCF` | Title, Content, CTA | Landing page hero |
| `[sidebar-panel]` | `HLC` | Title, Actions, Content | Dashboard widget |
| `[card]` | `HCF` | Header, Body, Footer | Content card |
| `[split]` | `LCR` | Left, Centre, Right | Comparison layout |

#### Nested shortcode structures

Shortcodes can nest within each other, inheriting the path context:

```
[dashboard layout="HLCF"]
    [widget layout="HCF" slot="L"]
        Widget content here
    [/widget]
    Main dashboard content
[/dashboard]
```

The widget's regions receive paths `L-H`, `L-C`, `L-F` because it renders within the dashboard's Left region. This happens automatically—shortcode authors need not manage paths manually.

### HTML5 slots integration

The path-based ID system integrates naturally with HTML5 `<slot>` elements, enabling Web Components to define HLCRF structures.

#### Slot element mapping

```html
<template id="hlcrf-component">
    <div class="hlcrf-layout">
        <header data-slot="H">
            <slot name="H"></slot>
        </header>
        <div class="hlcrf-body">
            <aside data-slot="L">
                <slot name="L"></slot>
            </aside>
            <main data-slot="C">
                <slot name="C"></slot>
            </main>
            <aside data-slot="R">
                <slot name="R"></slot>
            </aside>
        </div>
        <footer data-slot="F">
            <slot name="F"></slot>
        </footer>
    </div>
</template>
```

#### Nested slot paths

For nested layouts, slot names follow the path convention:

```html
<div data-slot="L-C">
    <slot name="L-C"></slot>
    <!-- Content injected into nested layout's Content region -->
</div>
```

The `data-slot` attribute and slot `name` always match, enabling both CSS targeting and content projection:

```html
<!-- Nested layout within the Left region -->
<aside data-slot="L">
    <div class="hlcrf-layout" data-layout="L-">
        <header data-slot="L-H">
            <slot name="L-H"></slot>
        </header>
        <main data-slot="L-C">
            <slot name="L-C"></slot>
        </main>
        <footer data-slot="L-F">
            <slot name="L-F"></slot>
        </footer>
    </div>
</aside>
```

Content authors inject into specific nested regions using the slot attribute:

```html
<my-layout-component>
    <h1 slot="H">Page Title</h1>
    <nav slot="L-H">Sidebar Navigation</nav>
    <div slot="L-C">Sidebar Content</div>
    <article slot="C">Main Content</article>
</my-layout-component>
```

#### Progressive enhancement

Slots enable progressive enhancement patterns:

1. **Server-rendered baseline** — PHP compositor renders complete HTML
2. **Client enhancement** — JavaScript can relocate content between slots
3. **Framework agnostic** — Works with vanilla JS, Alpine, Vue, or React

```html
<!-- Server-rendered -->
<main data-slot="C">
    <article>Content here</article>
</main>

<!-- JavaScript enhancement -->
<script>
    // Move content to different region based on viewport
    const content = document.querySelector('[data-slot="C"] article');
    if (viewport.isMobile) {
        document.querySelector('[data-slot="L"]').appendChild(content);
    }
</script>
```

### Alpine.js integration

The compositor's data attributes work naturally with Alpine.js:

```html
<div class="hlcrf-layout" x-data="{ activeRegion: 'C' }">
    <aside data-slot="L" x-show="activeRegion === 'L' || $screen('lg')">
        <!-- Sidebar content -->
    </aside>
    <main data-slot="C" @click="activeRegion = 'C'">
        <!-- Main content -->
    </main>
</div>
```

### Livewire component boundaries

HLCRF regions can serve as Livewire component boundaries:

```php
$layout = Layout::make('HLCF')
    ->h(livewire('header-nav'))
    ->l(livewire('sidebar-menu'))
    ->c(livewire('main-content'))
    ->f(livewire('footer-links'));
```

Each region becomes an independent Livewire component with its own state and lifecycle.

### Path-based event targeting

The hierarchical path system enables precise event targeting:

```javascript
// Listen for events in a specific nested region
document.querySelector('[data-slot="L-C"]')
    .addEventListener('block:added', (e) => {
        console.log(`Block added to left sidebar content: ${e.detail.blockId}`);
    });

// Broadcast to all blocks in a path
function notifyRegion(path, event) {
    document.querySelectorAll(`[data-slot^="${path}"]`)
        .forEach(el => el.dispatchEvent(new CustomEvent(event)));
}
```

### Server-side rendering integration

The compositor works with SSR frameworks:

```php
// Inertia.js integration
return Inertia::render('Dashboard', [
    'layout' => [
        'variant' => 'HLCF',
        'regions' => [
            'H' => $headerData,
            'L' => $sidebarData,
            'C' => $contentData,
            'F' => $footerData,
        ],
    ],
]);
```

The frontend receives structured data and renders using the same HLCRF conventions.

---

## Related files

- `app/Core/Front/Components/Layout.php` — Core compositor class
- `app/Core/Front/Components/View/Blade/layout.blade.php` — Blade component variant
- `app/Mod/Bio/Services/HlcrfRenderer.php` — Bio page rendering service
- `app/Mod/Bio/Migrations/2026_01_14_100000_add_hlcrf_support.php` — Database schema

---

## Version history

| Version | Date | Changes |
|---------|------|---------|
| 1.0 | 2026-01-15 | Initial RFC |
