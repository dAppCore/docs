---
title: go-log
description: Structured logger with Core DI container integration
---

# go-log

`forge.lthn.ai/core/go-log`

Structured logging library that integrates with the Core dependency injection container. Supports log rotation, configurable output formats, and error wrapping with contextual fields.

## Key Types

- `Logger` — structured logger with levelled output and field support
- `Err` — contextual error wrapper that attaches log fields
- `RotationOptions` — configuration for log file rotation (size, age, count)
- `Options` — logger configuration (level, format, output targets)
