# core docs

Documentation management and help engine for the Core ecosystem.

The `core docs` command collects documentation from across repositories. The `core/docs` Go module also contains `pkg/help`, a full-text search engine and static site generator that powers [core.help](https://core.help).

## Commands

```bash
core docs <command> [flags]
```

| Command | Description |
|---------|-------------|
| `list` | List documentation coverage across repos |
| `sync` | Sync documentation to an output directory |

## docs list

Show documentation coverage across all repos in the workspace.

```bash
core docs list [flags]
```

### Flags

| Flag | Description |
|------|-------------|
| `--registry` | Path to `repos.yaml` |

### Output

```
Repo                  README    CLAUDE    CHANGELOG   docs/
----------------------------------------------------------------------
core                  yes       yes       --          12 files
core-php              yes       yes       yes         8 files
core-images           yes       --        --          --

Coverage: 3 with docs, 0 without
```

## docs sync

Sync documentation from all repos to an output directory.

```bash
core docs sync [flags]
```

### Flags

| Flag | Description |
|------|-------------|
| `--registry` | Path to `repos.yaml` |
| `--output` | Output directory (default: `./docs-build`) |
| `--dry-run` | Show what would be synced |

### What Gets Synced

For each repo, the following files are collected:

| Source | Destination |
|--------|-------------|
| `README.md` | `index.md` |
| `CLAUDE.md` | `claude.md` |
| `CHANGELOG.md` | `changelog.md` |
| `docs/*.md` | `*.md` |

### Example

```bash
# Preview what will be synced
core docs sync --dry-run

# Sync to default output
core docs sync

# Sync to custom directory
core docs sync --output ./site/content
```

---

## Help Engine (`pkg/help`)

The `core/docs` repository is a Go module (`forge.lthn.ai/core/docs`) containing `pkg/help`, a documentation engine that provides:

- **Full-text search** with stemming, fuzzy matching, and phrase queries
- **Static site generation** for deployment to core.help
- **HTTP server** for in-app documentation (embedded in binaries)
- **HLCRF layout** via `go-html` for semantic HTML rendering

### Module Path

```
forge.lthn.ai/core/docs
```

### Architecture

```
core/docs/
  pkg/help/
    catalog.go          Topic store with search indexing
    search.go           Inverted index, stemming, fuzzy matching
    parser.go           YAML frontmatter + section extraction
    render.go           Goldmark Markdown to HTML
    stemmer.go          Porter-style word stemmer
    server.go           HTTP server (HTML + JSON API)
    generate.go         Static site generator
    layout.go           go-html HLCRF page compositor
    ingest.go           CLI help text to Topic conversion
    templates.go        Topic grouping helpers
  docs/                 MkDocs content (this documentation)
  zensical.toml         MkDocs Material configuration
  go.mod
```

### Topics

A `Topic` is the fundamental unit of documentation:

```go
type Topic struct {
    ID       string    // URL-safe slug (e.g. "rate-limiting")
    Title    string    // Display title
    Path     string    // Source file path
    Content  string    // Markdown body
    Sections []Section // Extracted headings with IDs
    Tags     []string  // Categorisation tags
    Related  []string  // Links to related topic IDs
    Order    int       // Sort order
}
```

Topics are parsed from Markdown files with optional YAML frontmatter:

```markdown
---
title: Rate Limiting
tags: [api, security]
related: [authentication]
order: 5
---

## Overview

Rate limiting controls request throughput...
```

Headings are automatically extracted as sections with URL-safe IDs (e.g. `"Rate Limiting"` becomes `rate-limiting`).

### Search

The search engine builds an inverted index from topic content and supports:

- **Keyword search**: Single or multi-word queries
- **Phrase search**: Quoted strings (e.g. `"rate limit"`)
- **Fuzzy matching**: Levenshtein distance (max 2) for typo tolerance
- **Prefix matching**: Partial word completion
- **Stemming**: Porter-style stemming for morphological variants

Relevance scoring weights:

| Signal | Weight |
|--------|--------|
| Title match | 10x |
| Phrase match | 8x |
| Section title match | 5x |
| Tag match | 3x |
| All words present | 2x |
| Exact word | 1x |
| Stemmed word | 0.7x |
| Prefix match | 0.5x |
| Fuzzy match | 0.3x |

### Catalog

The `Catalog` manages a collection of topics with automatic search indexing:

```go
// Load from a directory of Markdown files
catalog, err := help.LoadContentDir("./content")

// Or build programmatically
catalog := help.DefaultCatalog()
catalog.Add(&help.Topic{ID: "my-topic", Title: "My Topic", Content: "..."})

// Search
results := catalog.Search("rate limit")
for _, r := range results {
    fmt.Printf("%s (score: %.1f)\n", r.Topic.Title, r.Score)
}

// Get by ID
topic, err := catalog.Get("rate-limiting")
```

### HTTP Server

`pkg/help` provides an HTTP server with both HTML and JSON API endpoints:

```go
server := help.NewServer(catalog, ":8080")
server.ListenAndServe()
```

#### HTML Routes

| Route | Description |
|-------|-------------|
| `GET /` | Topic listing grouped by tags |
| `GET /topics/{id}` | Single topic page with table of contents |
| `GET /search?q=...` | Search results page |

#### JSON API Routes

| Route | Description |
|-------|-------------|
| `GET /api/topics` | All topics as JSON |
| `GET /api/topics/{id}` | Single topic with sections |
| `GET /api/search?q=...` | Search results with scores and snippets |

### Static Site Generation

The `Generate` function writes a complete static site:

```go
err := help.Generate(catalog, "./dist")
```

Output structure:

```
dist/
  index.html            Topic listing
  search.html           Client-side search (inline JS)
  search-index.json     JSON index for client-side search
  404.html              Not found page
  topics/
    rate-limiting.html  One page per topic
    authentication.html
    ...
```

The static site includes client-side JavaScript that fetches `search-index.json` and performs search without a server. All CSS is inlined -- no external stylesheets are needed.

### HLCRF Layout

Page rendering uses `go-html`'s HLCRF (Header, Left, Content, Right, Footer) compositor for semantic HTML:

| Slot | Element | Content |
|------|---------|---------|
| H | `<header role="banner">` | Nav bar with branding and search |
| L | `<aside role="complementary">` | Topic tree grouped by tag (topic pages only) |
| C | `<main role="main">` | Rendered Markdown with section anchors |
| F | `<footer role="contentinfo">` | Licence and source link |

The layout uses a dark theme with all CSS inlined. Pages are responsive -- the sidebar collapses on narrow viewports.

### Dual-Mode Serving

The help engine supports two deployment modes:

| Mode | Use Case | Details |
|------|----------|---------|
| **In-app** | CoreGUI webview, `core help serve` | `NewServer()` serves from `//go:embed` content, no network |
| **Public** | core.help | Static HTML from `Generate()`, deployed to CDN |

Both modes produce identical URLs and fragment anchors, so deep links work in either context.

## Dependencies

| Module | Purpose |
|--------|---------|
| `forge.lthn.ai/core/go-html` | HLCRF layout compositor |
| `github.com/yuin/goldmark` | Markdown to HTML rendering |
| `gopkg.in/yaml.v3` | YAML frontmatter parsing |

## See Also

- [core/lint](../../lint/) -- Pattern catalog and QA toolkit
