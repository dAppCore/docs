# go-html

HLCRF DOM compositor with grammar pipeline integration.

**Module**: `forge.lthn.ai/core/go-html`

Provides a type-safe node tree (El, Text, Raw, If, Each, Switch, Entitled), a five-slot Header/Left/Content/Right/Footer layout compositor with deterministic `data-block` path IDs and ARIA roles, a responsive multi-variant wrapper, a server-side grammar pipeline (StripTags, GrammarImprint via go-i18n reversal, CompareVariants), a build-time Web Component codegen CLI, and a WASM module (2.90 MB raw, 842 KB gzip) exposing `renderToString()`.

## Quick Start

```go
import "forge.lthn.ai/core/go-html"

page := html.NewLayout("HCF").
    H(html.El("nav", html.Text("i18n.label.navigation"))).
    C(html.El("main",
        html.El("h1", html.Text("i18n.label.welcome")),
        html.Each(items, func(item Item) html.Node {
            return html.El("li", html.Text(item.Name))
        }),
    )).
    F(html.El("footer", html.Text("i18n.label.copyright")))

rendered := page.Render(html.NewContext("en-GB"))
```

## Layout Slots

| Slot | Method | ARIA Role |
|------|--------|-----------|
| Header | `.H()` | `banner` |
| Left | `.L()` | `navigation` |
| Content | `.C()` | `main` |
| Right | `.R()` | `complementary` |
| Footer | `.F()` | `contentinfo` |
