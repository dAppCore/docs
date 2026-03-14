---
title: core/lint
description: Pattern catalog, regex-based code checker, and quality assurance toolkit for Go and PHP projects
---

# core/lint

`forge.lthn.ai/core/lint` is a standalone pattern catalog and code quality toolkit. It ships a YAML-based rule catalog for detecting security issues, correctness bugs, and modernisation opportunities in Go source code. It also provides a full PHP quality assurance pipeline and a suite of developer tooling wrappers.

The library is designed to be embedded into other tools. The YAML rule files are compiled into the binary at build time via `go:embed`, so there are no runtime file dependencies.

## Module Path

```
forge.lthn.ai/core/lint
```

Requires Go 1.26+.

## Quick Start

### As a Library

```go
import (
    lint "forge.lthn.ai/core/lint"
    lintpkg "forge.lthn.ai/core/lint/pkg/lint"
)

// Load the embedded rule catalog.
cat, err := lint.LoadEmbeddedCatalog()
if err != nil {
    log.Fatal(err)
}

// Filter rules for Go, severity medium and above.
rules := cat.ForLanguage("go")
filtered := (&lintpkg.Catalog{Rules: rules}).AtSeverity("medium")

// Create a scanner and scan a directory.
scanner, err := lintpkg.NewScanner(filtered)
if err != nil {
    log.Fatal(err)
}

findings, err := scanner.ScanDir("./src")
if err != nil {
    log.Fatal(err)
}

// Output results.
lintpkg.WriteText(os.Stdout, findings)
```

### As a CLI

```bash
# Build the binary
core build          # produces ./bin/core-lint

# Scan the current directory with all rules
core-lint lint check

# Scan with filters
core-lint lint check --lang go --severity high ./pkg/...

# Output as JSON
core-lint lint check --format json .

# Browse the catalog
core-lint lint catalog list
core-lint lint catalog list --lang go
core-lint lint catalog show go-sec-001
```

### QA Commands

The `qa` command group provides workflow-level quality assurance:

```bash
# Go-focused
core qa watch              # Monitor GitHub Actions after a push
core qa review             # PR review status with actionable next steps
core qa health             # Aggregate CI health across all repos
core qa issues             # Intelligent issue triage
core qa docblock           # Check Go docblock coverage

# PHP-focused
core qa fmt                # Format PHP code with Laravel Pint
core qa stan               # Run PHPStan/Larastan static analysis
core qa psalm              # Run Psalm static analysis
core qa audit              # Audit composer and npm dependencies
core qa security           # Security checks (.env, filesystem, deps)
core qa rector             # Automated code refactoring
core qa infection          # Mutation testing
core qa test               # Run Pest or PHPUnit tests
```

## Package Layout

| Package | Path | Description |
|---------|------|-------------|
| `lint` (root) | `lint.go` | Embeds YAML catalogs and exposes `LoadEmbeddedCatalog()` |
| `pkg/lint` | `pkg/lint/` | Core library: Rule, Catalog, Matcher, Scanner, Report, Complexity, Coverage, VulnCheck, Toolkit |
| `pkg/detect` | `pkg/detect/` | Project type detection (Go, PHP) by filesystem markers |
| `pkg/php` | `pkg/php/` | PHP quality tools: format, analyse, audit, security, refactor, mutation, test, pipeline, runner |
| `cmd/core-lint` | `cmd/core-lint/` | CLI binary (`core-lint lint check`, `core-lint lint catalog`) |
| `cmd/qa` | `cmd/qa/` | QA workflow commands (watch, review, health, issues, docblock, PHP tools) |
| `catalog/` | `catalog/` | YAML rule definitions (embedded at compile time) |

## Rule Catalogs

Three built-in YAML catalogs ship with the module:

| File | Rules | Focus |
|------|-------|-------|
| `go-security.yaml` | 6 | SQL injection, path traversal, XSS, timing attacks, log injection, secret leaks |
| `go-correctness.yaml` | 7 | Unsynchronised goroutines, silent error swallowing, panics in library code, file deletion |
| `go-modernise.yaml` | 5 | Replace legacy patterns with modern stdlib (`slices.Clone`, `slices.Sort`, `maps.Keys`, `errgroup`) |

Total: **18 rules** across 4 severity levels. All rules currently target Go. The catalog is extensible -- add more YAML files to `catalog/` and they are embedded automatically at build time.

### Rule Schema

Each rule is defined in YAML with the following fields:

```yaml
- id: go-sec-001
  title: "SQL wildcard injection in LIKE clauses"
  severity: high          # info, low, medium, high, critical
  languages: [go]
  tags: [security, injection]
  pattern: 'LIKE\s+\?.*["%].*\+'
  exclude_pattern: 'EscapeLike'              # suppress if this also matches
  fix: "Use parameterised LIKE with EscapeLike() helper"
  found_in: [go-store]                       # repos where first discovered
  example_bad: |
    db.Query("SELECT * FROM users WHERE name LIKE ?", "%"+input+"%")
  example_good: |
    db.Query("SELECT * FROM users WHERE name LIKE ?", "%"+store.EscapeLike(input)+"%")
  first_seen: "2026-03-09"
  detection: regex        # regex (only type currently supported)
  auto_fixable: false
```

| Field | Required | Description |
|-------|----------|-------------|
| `id` | yes | Unique identifier (e.g. `go-sec-001`) |
| `title` | yes | Human-readable description |
| `severity` | yes | One of: `info`, `low`, `medium`, `high`, `critical` |
| `languages` | yes | Target languages (e.g. `[go]`, `[go, php]`) |
| `tags` | no | Categorisation tags (e.g. `security`, `concurrency`) |
| `pattern` | yes | Regex pattern to match against each line |
| `exclude_pattern` | no | Regex that suppresses the match if it also matches the line |
| `fix` | no | Recommended fix |
| `found_in` | no | Repos where the pattern was first discovered |
| `example_bad` | no | Code example that triggers the rule |
| `example_good` | no | Code example showing the correct approach |
| `first_seen` | no | Date the rule was added |
| `detection` | yes | Detection method (currently only `regex`) |
| `auto_fixable` | no | Whether automated fixing is supported |

### Rule Reference

#### Security Rules (`go-security.yaml`)

| ID | Title | Severity |
|----|-------|----------|
| `go-sec-001` | SQL wildcard injection in LIKE clauses | high |
| `go-sec-002` | Path traversal via `filepath.Join` | high |
| `go-sec-003` | XSS via unescaped HTML in `fmt.Sprintf` | high |
| `go-sec-004` | Non-constant-time authentication comparison | critical |
| `go-sec-005` | Log injection via string concatenation | medium |
| `go-sec-006` | Secrets leaked in log output | critical |

#### Correctness Rules (`go-correctness.yaml`)

| ID | Title | Severity |
|----|-------|----------|
| `go-cor-001` | Goroutine launched without synchronisation | medium |
| `go-cor-002` | `WaitGroup.Wait` without timeout or context | low |
| `go-cor-003` | Silent error swallowing with blank identifier | medium |
| `go-cor-004` | Panic in library code | high |
| `go-cor-005` | File deletion without path validation | high |
| `go-cor-006` | HTTP response error discarded | high |
| `go-cor-007` | Wrong signal type (`syscall.Signal` instead of `os.Signal`) | low |

#### Modernisation Rules (`go-modernise.yaml`)

| ID | Title | Severity |
|----|-------|----------|
| `go-mod-001` | Manual slice clone (use `slices.Clone`) | info |
| `go-mod-002` | Legacy sort functions (use `slices.Sort`) | info |
| `go-mod-003` | Manual reverse loop (use `slices.Reverse`) | info |
| `go-mod-004` | Manual WaitGroup Add+Done (use `errgroup`) | info |
| `go-mod-005` | Manual map key collection (use `maps.Keys`) | info |

### Adding New Rules

Create a new YAML file in `catalog/` or add entries to an existing file. Rules are validated at load time -- the regex patterns must compile, and all required fields must be present. New catalog files are automatically discovered and embedded on the next build.

## Scanner

The scanner walks directory trees, detects file languages by extension, and matches rules against each line. Directories named `vendor`, `node_modules`, `.git`, `testdata`, and `.core` are excluded by default.

Supported file extensions:

| Extension | Language |
|-----------|----------|
| `.go` | go |
| `.php` | php |
| `.ts`, `.tsx` | ts |
| `.js`, `.jsx` | js |
| `.cpp`, `.cc`, `.c`, `.h` | cpp |
| `.py` | py |

### Matching Behaviour

For each file, the scanner:

1. Detects the language from the file extension
2. Filters rules to those targeting that language
3. Checks each line against each rule's `pattern` regex
4. Suppresses the match if the line also matches the rule's `exclude_pattern`
5. Records a `Finding` with rule ID, severity, file, line number, matched text, and fix

### Output Formats

Findings can be written in three formats:

- **text** (default): `file:line [severity] title (rule-id)`
- **json**: Pretty-printed JSON array
- **jsonl**: Newline-delimited JSON, one finding per line (compatible with `~/.core/ai/metrics/`)

## QA Commands

The `core qa` command group provides workflow-level quality assurance for both Go and PHP projects. It is registered as a CLI plugin via `cmd/qa/` and appears as `core qa` when the lint module is linked into the CLI binary.

### Go Commands

| Command | Description |
|---------|-------------|
| `core qa watch` | Monitor GitHub Actions after a push |
| `core qa review` | PR review status with actionable next steps |
| `core qa health` | Aggregate CI health across all repos |
| `core qa issues` | Intelligent issue triage |
| `core qa docblock` | Check Go docblock coverage |

### PHP Commands

All PHP commands auto-detect the project by checking for `composer.json`. They shell out to the relevant PHP tool (which must be installed via Composer).

| Command | Tool | Description |
|---------|------|-------------|
| `core qa fmt` | Laravel Pint | Format PHP code (`--fix` to apply, `--diff` to preview) |
| `core qa stan` | PHPStan/Larastan | Static analysis (`--level 0-9`, `--json`, `--sarif`) |
| `core qa psalm` | Psalm | Deep type-level analysis (`--fix`, `--baseline`) |
| `core qa audit` | Composer/npm | Audit dependencies for vulnerabilities |
| `core qa security` | Built-in | Check `.env` exposure, debug mode, filesystem permissions, HTTP headers |
| `core qa rector` | Rector | Automated refactoring (`--fix` to apply) |
| `core qa infection` | Infection | Mutation testing (`--min-msi`, `--threads`) |
| `core qa test` | Pest/PHPUnit | Run tests (`--parallel`, `--coverage`, `--filter`) |

### Project Detection

The `pkg/detect` package identifies project types by filesystem markers:

| Marker | Type |
|--------|------|
| `go.mod` | Go |
| `composer.json` | PHP |

`detect.DetectAll()` returns all detected types for a directory, enabling mixed Go+PHP projects to use the full QA suite.

## Dependencies

Direct dependencies:

| Module | Purpose |
|--------|---------|
| `forge.lthn.ai/core/cli` | CLI framework (`cli.Main()`, command registration, TUI styles) |
| `forge.lthn.ai/core/go-i18n` | Internationalisation for CLI strings |
| `forge.lthn.ai/core/go-io` | Filesystem abstraction for registry loading |
| `forge.lthn.ai/core/go-log` | Structured logging and error wrapping |
| `forge.lthn.ai/core/go-scm` | Repository registry (`repos.yaml`) for multi-repo commands |
| `github.com/stretchr/testify` | Test assertions |
| `gopkg.in/yaml.v3` | YAML parsing for rule catalogs |

The `pkg/lint` sub-package has minimal dependencies (only `gopkg.in/yaml.v3` and standard library). The heavier CLI and SCM dependencies live in `cmd/`.

## Licence

EUPL-1.2
