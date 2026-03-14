---
title: go-process
description: Process management with Core IPC integration for Go applications.
---

# go-process

`forge.lthn.ai/core/go-process` is a process management library that provides
spawning, monitoring, and controlling external processes with real-time output
streaming via the Core ACTION (IPC) system. It integrates directly with the
[Core DI framework](https://forge.lthn.ai/core/go) as a first-class service.

## Features

- Spawn and manage external processes with full lifecycle tracking
- Real-time stdout/stderr streaming via Core IPC actions
- Ring buffer output capture (default 1 MB, configurable)
- Process pipeline runner with dependency graphs, sequential, and parallel modes
- Daemon mode with PID file locking, health check HTTP server, and graceful shutdown
- Daemon registry for tracking running instances across the system
- Lightweight `exec` sub-package for one-shot command execution with logging
- Thread-safe throughout; designed for concurrent use

## Quick Start

### Register with Core

```go
import (
    "context"
    framework "forge.lthn.ai/core/go/pkg/core"
    "forge.lthn.ai/core/go-process"
)

// Create a Core instance with the process service
c, err := framework.New(
    framework.WithName("process", process.NewService(process.Options{})),
)
if err != nil {
    log.Fatal(err)
}

// Retrieve the typed service
svc, err := framework.ServiceFor[*process.Service](c, "process")
if err != nil {
    log.Fatal(err)
}
```

### Run a Command

```go
// Fire-and-forget (async)
proc, err := svc.Start(ctx, "go", "test", "./...")
if err != nil {
    return err
}
<-proc.Done()
fmt.Println(proc.Output())

// Synchronous convenience
output, err := svc.Run(ctx, "echo", "hello world")
```

### Listen for Events

Process lifecycle events are broadcast through Core's ACTION system:

```go
c.RegisterAction(func(c *framework.Core, msg framework.Message) error {
    switch m := msg.(type) {
    case process.ActionProcessStarted:
        fmt.Printf("Started: %s (PID %d)\n", m.Command, m.PID)
    case process.ActionProcessOutput:
        fmt.Print(m.Line)
    case process.ActionProcessExited:
        fmt.Printf("Exit code: %d (%s)\n", m.ExitCode, m.Duration)
    case process.ActionProcessKilled:
        fmt.Printf("Killed with %s\n", m.Signal)
    }
    return nil
})
```

### Global Convenience API

For applications that only need a single process service, a global singleton
is available:

```go
// Initialise once at startup
process.Init(coreInstance)

// Then use package-level functions anywhere
proc, _ := process.Start(ctx, "ls", "-la")
output, _ := process.Run(ctx, "date")
procs := process.List()
running := process.Running()
```

## Daemon mode

go-process also manages *this process* as a long-running service. Where
`Process` manages child processes, `Daemon` manages the current process's own
lifecycle -- PID file locking, health endpoints, signal handling, and graceful
shutdown.

These types were extracted from `core/cli` to give any Go service daemon
capabilities without depending on the full CLI framework.

### PID file

`PIDFile` enforces single-instance execution. It writes the current PID on
`Acquire()`, detects stale lock files, and cleans up on `Release()`.

```go
pf := process.NewPIDFile("/var/run/myapp.pid")
if err := pf.Acquire(); err != nil {
    log.Fatal("another instance is running")
}
defer pf.Release()
```

### Health server

`HealthServer` provides HTTP `/health` and `/ready` endpoints. Custom health
checks can be added and the ready state toggled independently.

```go
hs := process.NewHealthServer("127.0.0.1:9000")
hs.AddCheck(func() error { return db.Ping() })
hs.Start()
defer hs.Stop(ctx)
hs.SetReady(true)
```

### Daemon orchestration

`Daemon` combines PID file, health server, and signal handling into a single
struct. It listens for `SIGTERM`/`SIGINT` and calls registered shutdown hooks.

```go
d := process.NewDaemon(process.DaemonOptions{
    PIDFile:         "/var/run/myapp.pid",
    HealthAddr:      "127.0.0.1:9000",
    ShutdownTimeout: 30 * time.Second,
})
d.Start()
d.SetReady(true)
d.Run(ctx) // blocks until signal
```

### Daemon registry

The `Registry` tracks all running daemons across the system via JSON files
in `~/.core/daemons/`. When a `Daemon` is configured with a `Registry`, it
auto-registers on start and auto-unregisters on stop.

```go
reg := process.DefaultRegistry()

// Manual registration
reg.Register(process.DaemonEntry{
    Code: "my-app", Daemon: "serve", PID: os.Getpid(),
    Health: "127.0.0.1:9000", Project: "/path/to/project",
})

// List all live daemons (stale entries are pruned automatically)
entries, _ := reg.List()

// Auto-registration via Daemon
d := process.NewDaemon(process.DaemonOptions{
    Registry: reg,
    RegistryEntry: process.DaemonEntry{
        Code: "my-app", Daemon: "serve",
    },
})
```

The registry is consumed by `core start/stop/list` CLI commands for
project-level daemon management.

## Package layout

| Path | Description |
|------|-------------|
| `*.go` (root) | Process service, types, actions, runner, daemon, health, PID file, registry |
| `exec/` | Lightweight command wrapper with fluent API and structured logging |

Key files:

| File | Purpose |
|------|---------|
| `daemon.go` | `Daemon`, `DaemonOptions`, `Mode`, `DetectMode()` |
| `pidfile.go` | `PIDFile` (acquire, release, stale detection) |
| `health.go` | `HealthServer` with `/health` and `/ready` endpoints |
| `registry.go` | `Registry`, `DaemonEntry`, `DefaultRegistry()` |

## Module information

| Field | Value |
|-------|-------|
| Module path | `forge.lthn.ai/core/go-process` |
| Go version | 1.26.0 |
| Licence | EUPL-1.2 |

## Dependencies

| Module | Purpose |
|--------|---------|
| `forge.lthn.ai/core/go` | Core DI framework (`ServiceRuntime`, `Core.ACTION`, lifecycle interfaces) |
| `github.com/stretchr/testify` | Test assertions (test-only) |

The package has no other runtime dependencies beyond the Go standard library
and the Core framework.
