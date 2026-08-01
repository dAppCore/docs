---
title: Development
tags: [meta, contributing]
order: 3
---

# Development

## Repository layout

```
docs/                       Markdown content (217+ topic files)
go/pkg/help/                Go help library
external/                   Submodule deps (core/go etc.)
zensical.toml               Site build config (Python)
.forgejo/workflows/         Forgejo CI
LICENCE                     EUPL-1.2
```

## Local setup

### Go library

```bash
cd go
GOWORK=off go build ./...
GOWORK=off go test -count=1 ./...
GOWORK=off go vet ./...
```

### Static site

```bash
pip install zensical
zensical build      # Output: site/
zensical serve      # Local preview server
```

## Adding a topic

1. Create `docs/<section>/<slug>.md` with YAML frontmatter:

   ```markdown
   ---
   title: Your Topic Title
   tags: [section, keyword]
   related: [other-topic-id]
   order: 5
   ---

   # Your Topic Title

   Body content...
   ```

2. The site build picks it up automatically via the navigation tree in `zensical.toml`.
3. The Go `Catalog.LoadContentDir("docs")` also picks it up — no registration needed.

## CI/CD

Forgejo workflow at `.forgejo/workflows/deploy.yml`:
- Trigger: push to `main`
- Build: `zensical build`
- Deploy: `s3cmd sync site/ s3://core-help-bucket/` (BunnyCDN)

## Conventions

- Licence: EUPL-1.2 (SPDX headers in source files)
- Go module: `dappco.re/go/core/docs`
- Markdown frontmatter fields: `title`, `tags`, `related`, `order`
- See `AGENTS.md` for agent-specific workflow notes
