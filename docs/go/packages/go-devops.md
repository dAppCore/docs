---
title: go-devops
description: Multi-repo development workflows, deployment, and release snapshot generation for the Lethean ecosystem.
---

# go-devops

`forge.lthn.ai/core/go-devops` provides multi-repo development workflow
commands (`core dev`), deployment orchestration, documentation sync, and
release snapshot generation (`core.json`).

**Module**: `forge.lthn.ai/core/go-devops`
**Go**: 1.26
**Licence**: EUPL-1.2

## Decomposition

go-devops was originally a 31K LOC monolith covering builds, releases,
infrastructure, Ansible, containers, and code quality. It has since been
decomposed into focused, independently-versioned packages:

| Extracted package | Former location | What moved |
|-------------------|-----------------|------------|
| [go-build](go-build.md) | `build/`, `release/`, `sdk/` | Cross-compilation, code signing, release publishing, SDK generation |
| [go-infra](go-infra.md) | `infra/` | Hetzner Cloud/Robot, CloudNS provider APIs, `infra.yaml` config |
| [go-ansible](go-ansible.md) | `ansible/` | Pure Go Ansible playbook engine (41 modules, SSH) |
| [go-container](go-container.md) | `container/`, `devops/` | LinuxKit VM management, dev environments, image sources |

The `devkit/` package (cyclomatic complexity, coverage, vulnerability scanning)
was merged into `core/lint`.

After decomposition, go-devops retains multi-repo orchestration, deployment,
documentation sync, and manifest snapshot generation.

## What it does

| Area | Summary |
|------|---------|
| **Multi-repo workflows** | Status, commit, push, pull across all repos in a `repos.yaml` workspace |
| **GitHub integration** | Issue listing, PR review status, CI workflow checks |
| **Documentation sync** | Collect docs from multi-repo workspaces into a central location |
| **Deployment** | Coolify PaaS integration |
| **Release snapshots** | Generate `core.json` from `.core/manifest.yaml` for marketplace indexing |
| **Setup** | Repository and CI bootstrapping |

## Package layout

```
go-devops/
├── cmd/              CLI command registrations
│   ├── dev/          Multi-repo workflow commands (work, health, commit, push, pull)
│   ├── docs/         Documentation sync and listing
│   ├── deploy/       Coolify deployment commands
│   ├── setup/        Repository and CI bootstrapping
│   └── gitcmd/       Git helpers
├── deploy/           Deployment integrations (Coolify PaaS)
└── snapshot/         Frozen release manifest generation (core.json)
```

## CLI commands

go-devops registers commands into the `core` CLI binary (built from `forge.lthn.ai/core/cli`). Key commands:

```bash
# Multi-repo development
core dev health                # Quick summary across all repos
core dev work                  # Combined status, commit, push workflow
core dev commit                # Claude-assisted commits for dirty repos
core dev push                  # Push repos with unpushed commits
core dev pull                  # Pull repos behind remote

# GitHub integration
core dev issues                # List open issues across repos
core dev reviews               # PRs needing review
core dev ci                    # GitHub Actions status

# Documentation
core docs list                 # Scan repos for docs
core docs sync                 # Copy docs to central location
core docs sync --target gohelp # Sync to go-help format

# Deployment
core deploy servers            # List Coolify servers
core deploy apps               # List Coolify applications

# Setup
core setup repo                # Generate .core/ configuration for a repo
core setup ci                  # Bootstrap CI configuration
```

## Release snapshots

The `snapshot` package generates a frozen `core.json` manifest from
`.core/manifest.yaml`, embedding the git commit SHA, tag, and build
timestamp. This file is consumed by the marketplace for self-describing
package listings.

```json
{
  "schema": 1,
  "code": "photo-browser",
  "name": "Photo Browser",
  "version": "0.1.0",
  "commit": "a1b2c3d4...",
  "tag": "v0.1.0",
  "built": "2026-03-09T15:00:00Z",
  "daemons": { ... },
  "modules": [ ... ]
}
```

## Further reading

- [go-build](go-build.md) -- Build system, release pipeline, SDK generation
- [go-infra](go-infra.md) -- Infrastructure provider APIs
- [go-ansible](go-ansible.md) -- Pure Go Ansible playbook engine
- [go-container](go-container.md) -- LinuxKit VM management
- [Doc Sync](sync.md) -- Documentation sync across multi-repo workspaces
