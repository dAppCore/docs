---
title: IDE
description: Native desktop IDE built with Wails 3 and Angular 20
---

# IDE

`forge.lthn.ai/core/ide`

Native desktop IDE built with Wails 3 (Go backend) and Angular 20 (frontend). Includes a Claude Code bridge for AI-assisted development and an MCP bridge for tool integration.

## Stack

- **Backend**: Go (Wails 3)
- **Frontend**: Angular 20, Web Awesome, Font Awesome
- **AI**: Claude Code bridge for in-IDE agent interaction
- **Tools**: MCP bridge for protocol tool access

## Key Types

- `ClaudeBridge` — connects the IDE to Claude Code sessions
- `MCPBridge` — exposes MCP tools within the IDE
- `GreetService` — Wails service example/template

## Repository

- **Source**: [forge.lthn.ai/core/ide](https://forge.lthn.ai/core/ide)
- **Go module**: `forge.lthn.ai/core/ide`
