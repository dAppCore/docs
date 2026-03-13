# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Documentation platform for the Core ecosystem (CLI, Go packages, PHP modules, MCP tools). Published at https://core.help. Two main components:

1. **`docs/`** — Markdown source files (217+) with YAML frontmatter, organized by section (Go, PHP, TS, GUI, AI, Tools, Deploy, Publish)
2. **`pkg/help/`** — Go library for help content management: parsing, search, HTTP serving, and static site generation

## Common Commands

```bash
# Run all tests
go test ./...

# Run a single test
go test ./pkg/help/ -run TestFunctionName

# Run benchmarks
go test ./pkg/help/ -bench .

# Build the static documentation site (requires Python + zensical)
pip install zensical
zensical build
```

## Architecture: `pkg/help/`

The Go help library is display-agnostic — it can serve HTML, expose a JSON API, or generate a static site from the same content.

**Data flow:** Markdown files → `ParseTopic()` (parser.go) → `Topic` structs → `Catalog` (catalog.go) → consumed by Server, Search, or Generate.

Key types and their roles:
- **`Topic`/`Frontmatter`** (topic.go) — Data model. Topics have ID, title, content, sections, tags, related links, and sort order. Frontmatter is parsed from YAML `---` blocks.
- **`Catalog`** (catalog.go) — Topic registry with `Add`, `Get`, `List`, `Search`. `LoadContentDir()` recursively loads `.md` files from a directory. `DefaultCatalog()` provides built-in starter topics.
- **`searchIndex`** (search.go) — Full-text search with TF-IDF scoring, prefix matching, fuzzy matching, stemming (Porter stemmer in stemmer.go), and phrase detection. Title matches are boosted.
- **`Server`** (server.go) — HTTP handler with HTML routes (`/`, `/topics/{id}`, `/search`) and JSON API routes (`/api/topics`, `/api/topics/{id}`, `/api/search`).
- **`Generate*`** (generate.go) — Static site generator producing index, topic pages, 404, and `search-index.json` for client-side search.
- **`Render*`/`Layout*`** (render.go, layout.go) — HTML rendering using `forge.lthn.ai/core/go-html` (HLCRF layout pattern with dark theme).
- **`IngestHelp`** (ingest.go) — Converts Go CLI `--help` text output into structured `Topic` objects.

## Site Configuration

`zensical.toml` defines the doc site structure — navigation tree, theme settings, markdown extensions (admonition, mermaid, tabbed content, code highlighting). Zensical is a Python-based static site generator.

## CI/CD

Forgejo workflow (`.forgejo/workflows/deploy.yml`): on push to `main`, builds with `zensical build` and deploys the `site/` directory to BunnyCDN via s3cmd.

## Conventions

- License: EUPL-1.2 (SPDX headers in source files)
- Go module: `forge.lthn.ai/core/docs`
- Tests use `testify/assert` and `testify/require`
- Markdown files use YAML frontmatter (`title`, `tags`, `related`, `order` fields)
