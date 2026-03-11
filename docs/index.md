---
title: core.help
description: Documentation for the Core CLI, Go packages, PHP modules, and MCP tools
---

# Core Platform

Build, deploy, and manage Go and PHP applications with a unified toolkit.

---

| | |
|---|---|
| **[Core GO →](go/index.md)** | **[Core PHP →](php/index.md)** |
| DI framework, service lifecycle, and message-passing bus for Go. | Modular monolith for Laravel with event-driven loading and multi-tenancy. |
| **[CLI →](cli/index.md)** | **[Go Packages →](go/packages/index.md)** |
| Unified `core` command for Go/PHP dev, multi-repo management, builds. | AI, ML, DevOps, crypto, i18n, blockchain, and more. |
| **[Deploy →](deploy/index.md)** | **[Publish →](publish/index.md)** |
| Docker, PHP, and LinuxKit deployment targets with templates. | Release to GitHub, Docker Hub, Homebrew, Scoop, AUR, npm. |

## Quick Start

=== "Go"

    ```bash
    # Install the Core CLI
    go install forge.lthn.ai/core/cli/cmd/core@latest

    # Check your environment
    core doctor

    # Run tests, format, lint
    core go test
    core go fmt
    core go lint

    # Build your project
    core build
    ```

=== "PHP"

    ```bash
    # Install the framework
    composer require lthn/php

    # Create a module
    php artisan make:mod Commerce

    # Start dev environment
    core php dev

    # Run tests
    core php test
    ```

=== "Multi-Repo"

    ```bash
    # Health check across all repos
    core dev health

    # Full workflow: status, commit, push
    core dev work

    # Just show status
    core dev work --status
    ```

## Architecture

The Core platform spans two ecosystems:

**Go** provides the CLI toolchain and infrastructure services — build system, release pipeline, multi-repo management, LinuxKit VMs, and AI/ML integration. The `core` binary is the single entry point.

**PHP** provides the application framework — a Laravel-based modular monolith with event-driven module loading, automatic multi-tenancy, and packages for admin, API, commerce, content, MCP, and developer portals.

Both are connected through the CLI (`core go`, `core php`, `core build`, `core dev`) and share deployment pipelines (`core ci`, `core deploy`).

## Licence

EUPL-1.2 — [European Union Public Licence](https://joinup.ec.europa.eu/collection/eupl/eupl-text-eupl-12)
