[![Go Reference](https://pkg.go.dev/badge/dappco.re/go/core/docs.svg)](https://pkg.go.dev/dappco.re/go/core/docs)
[![License: EUPL-1.2](https://img.shields.io/badge/License-EUPL--1.2-blue.svg)](LICENCE)
[![Go Version](https://img.shields.io/badge/Go-1.26-00ADD8?style=flat&logo=go)](go/go.mod)

# core/docs

Documentation platform for the Core ecosystem (CLI, Go packages, PHP modules, MCP tools). Published at https://core.help.

**Module**: `dappco.re/go/core/docs`
**Licence**: EUPL-1.2
**Language**: Go 1.26 (library) + Python 3 (site build via zensical)

## Components

1. **`docs/`** — Markdown source files (217+) with YAML frontmatter, organised by section (Go, PHP, TS, GUI, AI, Tools, Deploy, Publish).
2. **`go/pkg/help/`** — Go library for help content management: parsing, search, HTTP serving, and static site generation.

## Quickstart

### Build the static site

```bash
pip install zensical
zensical build
# Output: site/
```

### Use the Go help library

```go
import help "dappco.re/go/core/docs/pkg/help"

cat, err := help.LoadContentDir("docs/")
if err != nil { panic(err) }

results := cat.Search("install")
for _, r := range results {
    fmt.Println(r.Topic.ID, r.Score)
}
```

### Run as an HTTP server

```go
srv := help.NewServer(cat)
http.ListenAndServe(":8080", srv)
// GET /, /topics/{id}, /search?q=X
// GET /api/topics, /api/topics/{id}, /api/search?q=X
```

## Layout

```
docs/                       Markdown content (217+ files)
go/pkg/help/                Help library (parser, catalog, search, server, generate, render)
external/                   Submodule deps
zensical.toml               Site build config (Python)
.forgejo/workflows/         Forgejo CI (deploy.yml builds + pushes site/ to BunnyCDN)
```

See [`docs/architecture.md`](docs/architecture.md) for the help-library data flow and [`docs/development.md`](docs/development.md) for contributor setup.

## Licence

EUPL-1.2 — see [LICENCE](LICENCE).
