---
title: Agent
description: Agent orchestration platform — Claude Code plugins, Go agent framework, and Laravel agent package
---

# Agent

`forge.lthn.ai/core/agent`

A monorepo containing the agent orchestration platform for Host UK. Combines Go and PHP into a single codebase that powers AI agent tooling across the stack.

## What's Inside

### Claude Code Plugins

A collection of Claude Code plugins for the Host UK federated monorepo:

| Plugin | Description |
|--------|-------------|
| **code** | Core development — hooks, scripts, data collection |
| **review** | Code review automation |

### Go Agent Framework

131 Go files providing the agent runtime, session management, and orchestration layer.

### Laravel Agent Package

231 PHP files providing the Laravel integration — agent sessions, plans, tool handlers, and the agentic portal backend.

## Repository

- **Source**: [forge.lthn.ai/core/agent](https://forge.lthn.ai/core/agent)
- **Go module**: `forge.lthn.ai/core/agent`
- **Composer**: `lthn/agentic`

## Absorbed Archives

This repo consolidates code from the now-archived `go-agent`, `go-agentic`, and `php-agentic` packages. The `php-devops` functionality is also planned to be implemented here.
