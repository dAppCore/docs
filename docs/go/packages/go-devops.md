# go-devops

Infrastructure and build automation library for the Lethean ecosystem.

**Module**: `forge.lthn.ai/core/go-devops`

Provides a native Go Ansible playbook executor (~30 modules over SSH without shelling out), a multi-target build pipeline with project type auto-detection (Go, Wails, Docker, C++, LinuxKit, Taskfile), code signing (macOS codesign, GPG, Windows signtool), release orchestration with changelog generation and eight publisher backends (GitHub Releases, Docker, Homebrew, npm, AUR, Scoop, Chocolatey, LinuxKit), Hetzner Cloud and Robot API clients, CloudNS DNS management, container/VM management via QEMU and Hyperkit, an OpenAPI SDK generator (TypeScript, Python, Go, PHP), and a developer toolkit with cyclomatic complexity analysis, vulnerability scanning, and coverage trending.

## Quick Start

```go
import (
    "forge.lthn.ai/core/go-devops/ansible"
    "forge.lthn.ai/core/go-devops/build"
    "forge.lthn.ai/core/go-devops/release"
)

// Run an Ansible playbook over SSH
pb, _ := ansible.ParsePlaybook("playbooks/deploy.yml")
inv, _ := ansible.ParseInventory("inventory.yml")
pb.Run(ctx, inv)

// Build and release
artifacts, _ := build.Build(ctx, ".")
release.Publish(ctx, releaseCfg, false)
```

## Key Packages

| Package | Description |
|---------|-------------|
| `ansible` | Native Go Ansible executor (~30 modules) |
| `build` | Multi-target build pipeline with auto-detection |
| `release` | Changelog generation + 8 publisher backends |
| `hetzner` | Hetzner Cloud and Robot API clients |
| `cloudns` | CloudNS DNS management |
| `sdk` | OpenAPI SDK generator |
| `devkit` | Complexity analysis, vuln scanning, coverage |
