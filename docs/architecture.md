---
title: Architecture
tags: [meta, architecture]
order: 2
---

# Architecture

Two-layer split: Markdown content under `docs/`, Go library under `go/pkg/help/`.

## Components

- **Markdown corpus** (`docs/`): 217+ topic files with YAML frontmatter (`title`, `tags`, `related`, `order`). Organised by section under subdirectories matching the navigation tree in `zensical.toml`.
- **Help library** (`go/pkg/help/`): display-agnostic Go package. Same `Catalog` powers HTTP server, JSON API, and static-site generator.

## Data flow

```
*.md files
  → ParseTopic() (parser.go)
  → Topic structs
  → Catalog (catalog.go)
  → Server | Search | Generate
```

## Key types

| Type | File | Role |
|---|---|---|
| `Topic`, `Frontmatter` | `topic.go` | Data model — id, title, content, sections, tags, related, sort order |
| `Catalog` | `catalog.go` | Topic registry with `Add`, `Get`, `List`, `Search`. `LoadContentDir()` recursively loads `.md` |
| `searchIndex` | `search.go` + `stemmer.go` | TF-IDF scoring, prefix + fuzzy matching, Porter stemming, title boost |
| `Server` | `server.go` | HTTP handler — HTML routes (`/`, `/topics/{id}`, `/search`), JSON API (`/api/...`) |
| `Generate*` | `generate.go` | Static site generator — index, topic pages, 404, `search-index.json` |
| `Render*`, `Layout*` | `render.go`, `layout.go` | HTML rendering via `forge.lthn.ai/core/go-html` (HLCRF + dark theme) |
| `IngestHelp` | `ingest.go` | Converts Go CLI `--help` output into structured `Topic` objects |

## Site build

`zensical.toml` defines navigation tree, theme, and Markdown extensions (admonition, mermaid, tabbed, code highlighting). Zensical is a Python static-site generator; it consumes the Markdown corpus and emits the published `site/`.
