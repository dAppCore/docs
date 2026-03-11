---
title: go-config
description: Configuration service for the Core DI container
---

# go-config

`forge.lthn.ai/core/go-config`

Configuration management service that integrates with the Core dependency injection container. Handles loading, merging, and accessing typed configuration values for services registered in the Core runtime.

## Key Types

- `Config` — configuration data holder with typed accessors
- `Service` — Core-integrated service that manages configuration lifecycle
- `ServiceOptions` — options for configuring the service registration
