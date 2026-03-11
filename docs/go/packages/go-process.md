---
title: go-process
description: Process and daemon management with PID files, health checks, and orchestration
---

# go-process

`forge.lthn.ai/core/go-process`

Process and daemon management library. Handles PID file creation, health checks, process lifecycle events, and daemon orchestration through a registry. Communicates process state via Core's action/IPC system.

## Key Types

- `ActionProcessStarted` — IPC event fired when a managed process starts
- `ActionProcessOutput` — IPC event carrying stdout/stderr output from a process
- `ActionProcessExited` — IPC event fired when a managed process exits
- `RingBuffer` — bounded circular buffer for capturing recent process output
