---
title: go-build
description: Build system, release publishers, and SDK generation
---

# go-build

`forge.lthn.ai/core/go-build`

Build system that auto-detects project types (Go, Wails, Docker, LinuxKit, C++) and produces cross-compiled binaries, archives, and checksums. Includes release publishers for GitHub, Docker registries, npm, Homebrew, Scoop, AUR, and Chocolatey. Also provides SDK generation for API clients.

## Packages

- `pkg/build/` — project-type detection, cross-compilation, archiving
- `pkg/release/` — changelog generation, publisher pipeline
- `pkg/sdk/` — client SDK generation from OpenAPI specs
