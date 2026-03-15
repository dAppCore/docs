# Configuration

Core uses `.core/` directory for project configuration. Config files are auto-discovered — commands need zero arguments.

## Quick Reference

| File | Scope | Package | Purpose |
|------|-------|---------|---------|
| `~/.core/config.yaml` | User | go-config | Global settings (Viper) |
| `.core/build.yaml` | Project | go-build | Build targets, flags, signing |
| `.core/release.yaml` | Project | go-build | Publishers, changelog, SDK gen |
| `.core/php.yaml` | Project | go-build | PHP/Laravel dev, test, deploy |
| `.core/test.yaml` | Project | go-container | Named test commands |
| `.core/manifest.yaml` | App | go-scm | Providers, daemons, permissions |
| `.core/workspace.yaml` | Workspace | agent | Active package, paths |
| `.core/repos.yaml` | Workspace | go-scm | Repo registry + dependencies |
| `.core/work.yaml` | Workspace | go-scm | Sync policy, agent heartbeat |
| `.core/git.yaml` | Machine | go-scm | Local git state (gitignored) |
| `.core/kb.yaml` | Workspace | go-scm | Wiki mirror, Qdrant search |
| `.core/linuxkit/*.yml` | Project | go-container | VM templates |

**Scopes:**

- **User** (`~/.core/`) — global settings, persists across all projects
- **Project** (`{repo}/.core/`) — per-repository config, checked into git
- **Workspace** (`{workspace}/.core/`) — multi-repo workspace config, checked into git
- **Machine** (`{workspace}/.core/`) — per-machine state, gitignored

**Discovery patterns:**

- **Fixed path** — `build.yaml`, `release.yaml`, `test.yaml`, `manifest.yaml`
- **Walk-up** — `workspace.yaml`, `repos.yaml` (search current dir → parents → home)
- **Direct load** — `work.yaml`, `git.yaml`, `kb.yaml` (from workspace root)

## Directory Structure

```
~/.core/                      # User-level (global)
├── config.yaml               # Global settings
├── plugins/                  # Plugin discovery
├── known_hosts               # SSH known hosts
└── linuxkit/                 # User LinuxKit templates

{workspace}/.core/            # Workspace-level (shared)
├── workspace.yaml            # Active package, paths
├── repos.yaml                # Repository registry
├── work.yaml                 # Sync policy, agent heartbeat
├── git.yaml                  # Machine-local git state (gitignored)
└── kb.yaml                   # Knowledge base config

{project}/.core/              # Project-level (per-repo)
├── build.yaml                # Build configuration
├── release.yaml              # Release configuration
├── php.yaml                  # PHP/Laravel configuration
├── test.yaml                 # Test commands
├── manifest.yaml             # Application manifest
└── linuxkit/                 # LinuxKit templates
    ├── server.yml
    └── dev.yml
```

## release.yaml

Full release configuration reference:

```yaml
version: 1

project:
  name: myapp
  repository: myorg/myapp

build:
  targets:
    - os: linux
      arch: amd64
    - os: linux
      arch: arm64
    - os: darwin
      arch: amd64
    - os: darwin
      arch: arm64
    - os: windows
      arch: amd64

publishers:
  # GitHub Releases (required - others reference these artifacts)
  - type: github
    prerelease: false
    draft: false

  # npm binary wrapper
  - type: npm
    package: "@myorg/myapp"
    access: public  # or "restricted"

  # Homebrew formula
  - type: homebrew
    tap: myorg/homebrew-tap
    formula: myapp
    official:
      enabled: false
      output: dist/homebrew

  # Scoop manifest (Windows)
  - type: scoop
    bucket: myorg/scoop-bucket
    official:
      enabled: false
      output: dist/scoop

  # AUR (Arch Linux)
  - type: aur
    maintainer: "Name <email>"

  # Chocolatey (Windows)
  - type: chocolatey
    push: false  # true to publish

  # Docker multi-arch
  - type: docker
    registry: ghcr.io
    image: myorg/myapp
    dockerfile: Dockerfile
    platforms:
      - linux/amd64
      - linux/arm64
    tags:
      - latest
      - "{{.Version}}"
    build_args:
      VERSION: "{{.Version}}"

  # LinuxKit images
  - type: linuxkit
    config: .core/linuxkit/server.yml
    formats:
      - iso
      - qcow2
      - docker
    platforms:
      - linux/amd64
      - linux/arm64

changelog:
  include:
    - feat
    - fix
    - perf
    - refactor
  exclude:
    - chore
    - docs
    - style
    - test
    - ci
```

## build.yaml

Optional build configuration:

```yaml
version: 1

project:
  name: myapp
  binary: myapp

build:
  main: ./cmd/myapp
  env:
    CGO_ENABLED: "0"
  flags:
    - -trimpath
  ldflags:
    - -s -w
    - -X main.version={{.Version}}
    - -X main.commit={{.Commit}}

targets:
  - os: linux
    arch: amd64
  - os: darwin
    arch: arm64
```

## php.yaml

PHP/Laravel configuration:

```yaml
version: 1

dev:
  domain: myapp.test
  ssl: true
  port: 8000
  services:
    - frankenphp
    - vite
    - horizon
    - reverb
    - redis

test:
  parallel: true
  coverage: false

deploy:
  coolify:
    server: https://coolify.example.com
    project: my-project
    environment: production
```

## LinuxKit Templates

LinuxKit YAML configuration:

```yaml
kernel:
  image: linuxkit/kernel:6.6
  cmdline: "console=tty0 console=ttyS0"

init:
  - linuxkit/init:latest
  - linuxkit/runc:latest
  - linuxkit/containerd:latest
  - linuxkit/ca-certificates:latest

onboot:
  - name: sysctl
    image: linuxkit/sysctl:latest

services:
  - name: dhcpcd
    image: linuxkit/dhcpcd:latest
  - name: sshd
    image: linuxkit/sshd:latest
  - name: myapp
    image: myorg/myapp:latest
    capabilities:
      - CAP_NET_BIND_SERVICE

files:
  - path: /etc/myapp/config.yaml
    contents: |
      server:
        port: 8080
```

## repos.yaml

Package registry for multi-repo workspaces:

```yaml
# Organisation name (used for GitHub URLs)
org: host-uk

# Base path for cloning (default: current directory)
base_path: .

# Default settings for all repos
defaults:
  ci: github
  license: EUPL-1.2
  branch: main

# Repository definitions
repos:
  # Foundation packages (no dependencies)
  core-php:
    type: foundation
    description: Foundation framework

  core-devops:
    type: foundation
    description: Development environment
    clone: false  # Skip during setup (already exists)

  # Module packages (depend on foundation)
  core-tenant:
    type: module
    depends_on: [core-php]
    description: Multi-tenancy module

  core-admin:
    type: module
    depends_on: [core-php, core-tenant]
    description: Admin panel

  core-api:
    type: module
    depends_on: [core-php]
    description: REST API framework

  # Product packages (user-facing applications)
  core-bio:
    type: product
    depends_on: [core-php, core-tenant]
    description: Link-in-bio product
    domain: bio.host.uk.com

  core-social:
    type: product
    depends_on: [core-php, core-tenant]
    description: Social scheduling
    domain: social.host.uk.com

  # Templates
  core-template:
    type: template
    description: Starter template for new projects
```

### repos.yaml Fields

| Field | Required | Description |
|-------|----------|-------------|
| `org` | Yes | GitHub organisation name |
| `base_path` | No | Directory for cloning (default: `.`) |
| `defaults` | No | Default settings applied to all repos |
| `repos` | Yes | Map of repository definitions |

### Repository Fields

| Field | Required | Description |
|-------|----------|-------------|
| `type` | Yes | `foundation`, `module`, `product`, or `template` |
| `description` | No | Human-readable description |
| `depends_on` | No | List of package dependencies |
| `clone` | No | Set `false` to skip during setup |
| `domain` | No | Production domain (for products) |
| `branch` | No | Override default branch |

### Package Types

| Type | Description | Dependencies |
|------|-------------|--------------|
| `foundation` | Core framework packages | None |
| `module` | Reusable modules | Foundation packages |
| `product` | User-facing applications | Foundation + modules |
| `template` | Starter templates | Any |

## workspace.yaml

Workspace-level configuration. Discovered by walking up from CWD.

```yaml
version: 1

# Active package for unified commands
active: core-php

# Default package types for setup
default_only:
  - foundation
  - module

# Paths
packages_dir: ./packages

# Workspace settings
settings:
  suggest_core_commands: true
  show_active_in_prompt: true
```

**Package:** `forge.lthn.ai/core/agent` · **Discovery:** walk-up from CWD

## work.yaml

Team sync policy. Checked into git (shared across team).

```yaml
version: 1

sync:
  interval: 5m
  auto_pull: true
  auto_push: false
  clone_missing: true

agent:
  heartbeat_interval: 30s
  stale_after: 10m
  overlap_warning: true

triggers:
  on_activate: sync
  on_commit: push
  scheduled: "*/5 * * * *"
```

**Package:** `forge.lthn.ai/core/go-scm/repos` · **Discovery:** `{workspaceRoot}/.core/work.yaml`

## git.yaml

Machine-local git state. **Gitignored** — not shared across machines.

```yaml
version: 1

repos:
  core-php:
    branch: main
    remote: origin
    last_pull: "2026-03-15T10:00:00Z"
    last_push: "2026-03-15T09:45:00Z"
    ahead: 0
    behind: 0
  core-tenant:
    branch: main
    remote: origin
    last_pull: "2026-03-15T10:00:00Z"

agent:
  name: cladius
  last_heartbeat: "2026-03-15T10:05:00Z"
```

**Package:** `forge.lthn.ai/core/go-scm/repos` · **Discovery:** `{workspaceRoot}/.core/git.yaml`

## kb.yaml

Knowledge base configuration. Controls wiki mirroring and vector search.

```yaml
version: 1

wiki:
  enabled: true
  directory: kb          # Relative to .core/
  remote: "ssh://git@forge.lthn.ai:2223/core/wiki.git"

search:
  qdrant:
    host: qdrant.lthn.sh
    port: 6334
    collection: openbrain
  ollama:
    url: http://ollama.lthn.sh
    model: embeddinggemma
  top_k: 10
```

**Package:** `forge.lthn.ai/core/go-scm/repos` · **Discovery:** `{workspaceRoot}/.core/kb.yaml`

## test.yaml

Named test commands per project. Auto-detected if not present.

```yaml
version: 1

commands:
  unit:
    run: composer test -- --filter=Unit
    env:
      APP_ENV: testing
  integration:
    run: composer test -- --filter=Integration
    env:
      APP_ENV: testing
      DB_DATABASE: test_db
  all:
    run: composer test
```

**Auto-detection chain** (if no `test.yaml`): `composer.json` → `package.json` → `go.mod` → `pytest` → `Taskfile`

**Package:** `forge.lthn.ai/core/go-container/devenv` · **Discovery:** `{projectDir}/.core/test.yaml`

## manifest.yaml

Application manifest for providers, daemons, and permissions. Supports ed25519 signature verification.

```yaml
version: 1

app:
  name: my-provider
  namespace: my-provider
  description: Custom service provider

providers:
  - namespace: my-provider
    port: 9900
    binary: ./bin/my-provider
    args: ["serve"]
    elements:
      - tag: my-provider-panel
        source: /assets/my-provider.js

daemons:
  - name: worker
    command: ./bin/worker
    restart: always

permissions:
  - net.listen
  - fs.read
```

**Package:** `forge.lthn.ai/core/go-scm/manifest` · **Discovery:** `{appRoot}/.core/manifest.yaml`

---

## Environment Variables

Complete reference of environment variables used by Core CLI.

### Authentication

| Variable | Used By | Description |
|----------|---------|-------------|
| `GITHUB_TOKEN` | `core ci`, `core dev` | GitHub API authentication |
| `ANTHROPIC_API_KEY` | `core ai`, `core dev claude` | Claude API key |
| `AGENTIC_TOKEN` | `core ai task*` | Agentic API authentication |
| `AGENTIC_BASE_URL` | `core ai task*` | Agentic API endpoint |

### Publishing

| Variable | Used By | Description |
|----------|---------|-------------|
| `NPM_TOKEN` | `core ci` (npm publisher) | npm registry auth token |
| `CHOCOLATEY_API_KEY` | `core ci` (chocolatey publisher) | Chocolatey API key |
| `DOCKER_USERNAME` | `core ci` (docker publisher) | Docker registry username |
| `DOCKER_PASSWORD` | `core ci` (docker publisher) | Docker registry password |

### Deployment

| Variable | Used By | Description |
|----------|---------|-------------|
| `COOLIFY_URL` | `core php deploy` | Coolify server URL |
| `COOLIFY_TOKEN` | `core php deploy` | Coolify API token |
| `COOLIFY_APP_ID` | `core php deploy` | Production application ID |
| `COOLIFY_STAGING_APP_ID` | `core php deploy --staging` | Staging application ID |

### Build

| Variable | Used By | Description |
|----------|---------|-------------|
| `CGO_ENABLED` | `core build`, `core go *` | Enable/disable CGO (default: 0) |
| `GOOS` | `core build` | Target operating system |
| `GOARCH` | `core build` | Target architecture |

### Configuration Paths

| Variable | Description |
|----------|-------------|
| `CORE_CONFIG` | Override config directory (default: `~/.core/`) |
| `CORE_REGISTRY` | Override repos.yaml path |

---

## Defaults

If no configuration exists, sensible defaults are used:

- **Targets**: linux/amd64, linux/arm64, darwin/amd64, darwin/arm64, windows/amd64
- **Publishers**: GitHub only
- **Changelog**: feat, fix, perf, refactor included
