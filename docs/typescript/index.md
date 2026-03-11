---
title: CoreTS
description: TypeScript runtime management — Go service wrapping Deno with lifecycle integration
---

# CoreTS

`forge.lthn.ai/core/ts`

Go service that manages a Deno TypeScript runtime as a sidecar process. Provides lifecycle integration with the Core framework, permission management, and a client interface for communicating with TypeScript modules from Go.

## Key Types

- `Sidecar` — manages the Deno process lifecycle
- `DenoClient` — communicates with running Deno instances
- `Options` — configuration for the TypeScript runtime
- `Permissions` — Deno permission mapping (network, read, write, env)
- `ModulePermissions` — per-module permission scoping

## Repository

- **Source**: [forge.lthn.ai/core/ts](https://forge.lthn.ai/core/ts)
- **Go module**: `forge.lthn.ai/core/ts`
