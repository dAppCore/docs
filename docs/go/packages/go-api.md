---
title: go-api
description: REST framework with OpenAPI SDK generation, built on Gin
---

# go-api

`forge.lthn.ai/core/go-api`

HTTP REST framework built on Gin with automatic OpenAPI spec generation and client SDK output. Provides a route group interface for modular endpoint registration and includes health check endpoints out of the box.

## Key Types

- `Engine` — HTTP server wrapping Gin with OpenAPI integration and middleware pipeline
- `healthGroup` — built-in health check route group for liveness and readiness probes
- `RouteGroup` — interface for modular route registration
