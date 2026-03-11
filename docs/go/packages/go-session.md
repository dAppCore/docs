---
title: go-session
description: Claude Code JSONL transcript parser, analytics engine, and HTML timeline renderer
---

# go-session

`forge.lthn.ai/core/go-session`

Parser and analytics engine for Claude Code JSONL session transcripts. Extracts per-tool usage statistics, generates self-contained HTML timeline visualisations, and produces VHS tape scripts for terminal recording. Useful for understanding agent behaviour and tool usage patterns across sessions.

## Key Types

- `SessionAnalytics` — aggregated statistics from one or more sessions (tool counts, durations, token usage)
- `Event` — single parsed event from a JSONL transcript
- `Session` — complete parsed session with ordered events and metadata
