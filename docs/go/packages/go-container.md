---
title: go-container
description: Container runtime, LinuxKit builder, and dev environment management
---

# go-container

`forge.lthn.ai/core/go-container`

Container runtime abstraction supporting multiple hypervisors for running LinuxKit-based images and managing development environments. Handles image building, VM lifecycle, and port forwarding across QEMU and HyperKit backends.

## Key Types

- `Container` — manages container lifecycle (create, start, stop, remove)
- `RunOptions` — configuration for container execution (ports, volumes, environment)
- `HypervisorOptions` — shared configuration across hypervisor backends
- `QemuHypervisor` — QEMU-based VM backend
- `HyperkitHypervisor` — HyperKit-based VM backend (macOS)
