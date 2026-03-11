# go-ai

MCP (Model Context Protocol) hub for the Lethean AI stack.

**Module**: `forge.lthn.ai/core/go-ai`

Exposes 49 tools across file operations, directory management, language detection, RAG vector search, ML inference and scoring, process management, WebSocket streaming, browser automation via Chrome DevTools Protocol, JSONL metrics, and an IDE bridge to the Laravel core-agentic backend. The package is a pure library — the Core CLI (`core mcp serve`) imports it and handles transport selection (stdio, TCP, or Unix socket).

## Quick Start

```go
import "forge.lthn.ai/core/go-ai/mcp"

svc, err := mcp.New(
    mcp.WithWorkspaceRoot("/path/to/project"),
    mcp.WithProcessService(ps),
)
// Run as stdio server (default for AI client subprocess integration)
err = svc.Run(ctx)
// Or TCP: MCP_ADDR=127.0.0.1:9100 triggers ServeTCP automatically
```

## Tool Categories

| Category | Tools | Description |
|----------|-------|-------------|
| File | 8 | Read, write, search, glob, patch files |
| Directory | 4 | List, create, move, tree |
| Language | 3 | Detect, grammar, translate |
| RAG | 5 | Ingest, search, embed, index, stats |
| ML | 6 | Generate, score, probe, model management |
| Process | 4 | Start, stop, list, logs |
| WebSocket | 3 | Connect, send, subscribe |
| Browser | 5 | Navigate, click, read, screenshot, evaluate |
| Metrics | 3 | Write, query, dashboard |
| IDE | 8 | Bridge to core-agentic Laravel backend |
