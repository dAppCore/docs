---
title: LEM Agent Integration
description: Lethean Evaluation Model agent integration for the Core platform.
---

# LEM Agent Integration

LEM (Lethean Evaluation Model) is a locally-hosted AI model trained on the Core ecosystem. It integrates with the agent orchestration layer for on-device inference without external API dependencies.


## Overview

LEM runs as a native macOS application (LEM Lab) or as a backend service, providing:

- **Local inference** -- MLX-based, runs on Apple Silicon with no cloud dependency
- **Ecosystem knowledge** -- trained on Core framework documentation, patterns, and conventions
- **Cascade scoring** -- EaaS (Evaluation as a Service) for grading agent outputs
- **Poindexter spatial indexing** -- KDTree/cosine similarity for knowledge gap detection


## Integration Points

### Agent Dispatch

The agent orchestration layer (`pkg/lifecycle/`) can dispatch work to LEM instances alongside Claude, Gemini, and Codex. LEM is registered as an agent in the fleet registry.

### MCP Tools

LEM exposes tools via the Model Context Protocol, making them available to other agents in the fleet:

| Tool | Description |
|------|-------------|
| `lem_evaluate` | Score a response using the EaaS cascade |
| `lem_embed` | Generate embeddings for knowledge indexing |
| `lem_chat` | Local chat completion |

### Scorer Binary

The LEM scorer binary (`core-lem`) can be invoked as a subprocess from the PHP platform:

```bash
# Score a response against criteria
core-lem score --input response.txt --criteria accuracy,completeness
```

The PHP module (`app/Mod/Lem/`) wraps this via `proc_open` for integration with the Laravel application.


## Community Compute

LEM is designed for distributed evaluation across community-donated compute:

1. **Local-first** -- each contributor runs LEM on their own hardware
2. **No API keys required** -- MLX inference runs entirely on-device
3. **Federated scoring** -- results are aggregated without sharing raw data
4. **Google Developer credits** -- contributors can donate their API credits for cloud-based model access


## Requirements

- macOS with Apple Silicon (M1+) for MLX inference
- 16 GB RAM minimum (32 GB recommended for larger models)
- Core CLI (`core` binary) for build and orchestration


## Licence

EUPL-1.2
