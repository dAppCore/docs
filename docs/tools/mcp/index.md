---
title: MCP
description: Model Context Protocol — Go MCP server and Laravel MCP package with 49 tools
---

# MCP

`forge.lthn.ai/core/mcp`

Model Context Protocol server combining a Go MCP server with a Laravel MCP package. Produces the `core-mcp` binary. 49 MCP tools covering brain, RAG, ML, IDE bridge, and more.

## What's Inside

### Go MCP Server

43 Go files in `pkg/mcp/` providing the MCP server implementation and `cmd/` entry point for the `core-mcp` binary.

### Laravel MCP Package

145 PHP files in `src/php/` providing the Laravel integration — tool handlers, workspace management, database querying, analytics, quotas, and security.

## Repository

- **Source**: [forge.lthn.ai/core/mcp](https://forge.lthn.ai/core/mcp)
- **Go module**: `forge.lthn.ai/core/mcp`
- **Composer**: `lthn/mcp`

## Absorbed Archives

This repo consolidates code from the now-archived `php-mcp` package.
