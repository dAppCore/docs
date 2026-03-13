---
title: Gemini CLI Extension
description: Google Gemini CLI extension for the Core platform.
---

# Gemini CLI Extension

Core Agent includes a Gemini CLI extension (`google/gemini-cli/`) that provides tool integration for Google's Gemini CLI.


## Structure

```
google/gemini-cli/
+-- gemini-extension.json    # Extension manifest
+-- GEMINI.md                # Agent instructions
+-- package.json             # npm dependencies
+-- src/                     # TypeScript source
+-- commands/                # Slash commands
+-- hooks/                   # Event hooks
```


## Installation

The Gemini extension is a TypeScript package. Install dependencies and register it with Gemini CLI:

```bash
cd core/agent/google/gemini-cli
npm install
```

Gemini CLI reads `GEMINI.md` at the repository root for agent instructions, similar to how Claude Code reads `CLAUDE.md`.


## MCP Server

A separate HTTP MCP server (`google/mcp/`) provides Gemini-compatible tools:

| Tool | Description |
|------|-------------|
| `core_go_test` | Run Go tests |
| `core_dev_health` | Check development environment health |
| `core_dev_commit` | Commit changes across repos |

The MCP server listens on `:8080` and uses HTTP transport.

```bash
# Start the MCP server
go run google/mcp/main.go
```


## Licence

EUPL-1.2
