---
title: OpenAI Codex Plugins
description: OpenAI Codex plugin structure and command reference for the Core platform.
---

# OpenAI Codex Plugins

Core Agent includes a `codex/` directory that mirrors the Claude plugin structure for OpenAI Codex compatibility, with additional plugins for ethics, guardrails, and performance.


## Structure

Codex plugins use `AGENTS.md` files for agent instructions (equivalent to Claude's `CLAUDE.md`). Each plugin directory contains commands and scripts.

```
codex/
+-- AGENTS.md
+-- code/          # Core development
+-- review/        # Code review
+-- verify/        # Work verification
+-- qa/            # QA pipeline
+-- ci/            # CI/CD management
+-- issue/         # Issue tracking
+-- coolify/       # Coolify deployment
+-- api/           # API generation
+-- ethics/        # Ethics guardrails
+-- guardrails/    # Safety checks
+-- perf/          # Performance analysis
+-- awareness/     # Context awareness
+-- collect/       # Data collection
+-- core/          # Core framework
```


## Plugins

### Development

| Plugin | Purpose | Key Commands |
|--------|---------|--------------|
| **code** | Development workflow | Research, data collection |
| **review** | Code review | Review, security, PR review |
| **verify** | Work verification | Verify, ready check, tests |

### Quality

| Plugin | Purpose | Key Commands |
|--------|---------|--------------|
| **qa** | QA pipeline | `qa`, `fix`, `check`, `lint` |
| **ci** | CI/CD management | `ci`, `status`, `run`, `fix`, `workflow` |

### Operations

| Plugin | Purpose | Key Commands |
|--------|---------|--------------|
| **issue** | Issue tracking | `list`, `view`, `start`, `close` |
| **coolify** | Coolify deployment | `deploy`, `status` |
| **api** | API generation | `generate` |

### Safety

| Plugin | Purpose |
|--------|---------|
| **ethics** | Ethical guardrails for agent actions |
| **guardrails** | Safety checks and boundaries |
| **perf** | Performance analysis and monitoring |
| **awareness** | Context awareness and self-reflection |


## Configuration

Codex reads `AGENTS.md` at the repository root and within each plugin directory. No marketplace or install step is needed -- Codex discovers `AGENTS.md` files automatically.


## Licence

EUPL-1.2
