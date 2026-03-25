---
title: Core Go Framework
description: Dependency injection, service lifecycle, and permission framework for Go.
---

# Core Go Framework

Core (`dappco.re/go/core`) is a dependency injection, service lifecycle, and permission framework for Go. It provides a typed service registry, lifecycle hooks, a message-passing bus, named actions with task composition, and an entitlement primitive for permission gating.

This is the foundation layer of the ecosystem. It has no CLI, no GUI, and minimal dependencies (stdlib + go-io + go-log).

## Installation

```bash
go get dappco.re/go/core
```

Requires Go 1.26 or later.

## Quick Example

```go
package main

import "dappco.re/go/core"

func main() {
    c := core.New(
        core.WithOption("name", "my-app"),
        core.WithService(mypackage.Register),
        core.WithService(monitor.Register),
        core.WithServiceLock(),
    )
    c.Run()
}
```

`core.New()` returns `*Core`. Services register via factory functions. `Run()` calls `ServiceStartup`, runs the CLI, then `ServiceShutdown`. For error handling use `RunE()` which returns `error` instead of calling `os.Exit`.

## Service Registration

```go
func Register(c *core.Core) core.Result {
    svc := &MyService{
        ServiceRuntime: core.NewServiceRuntime(c, MyOptions{}),
    }
    return core.Result{Value: svc, OK: true}
}

// In main:
core.New(core.WithService(mypackage.Register))
```

Services implement `Startable` and/or `Stoppable` for lifecycle hooks:

```go
type Startable interface {
    OnStartup(ctx context.Context) Result
}

type Stoppable interface {
    OnShutdown(ctx context.Context) Result
}
```

## Subsystem Accessors

```go
c.Options()      // *Options  — input configuration
c.App()          // *App      — application identity
c.Config()       // *Config   — runtime settings, feature flags
c.Data()         // *Data     — embedded assets (Registry[*Embed])
c.Drive()        // *Drive    — transport handles (Registry[*DriveHandle])
c.Fs()           // *Fs       — filesystem I/O (sandboxable)
c.Cli()          // *Cli      — CLI command framework
c.IPC()          // *Ipc      — message bus internals
c.Log()          // *ErrorLog — structured logging
c.Error()        // *ErrorPanic — panic recovery
c.I18n()         // *I18n     — internationalisation
c.Process()      // *Process  — managed execution (Action sugar)
c.Context()      // context.Context — Core's lifecycle context
c.Env(key)       // string    — environment variable (cached at init)
```

## Primitive Types

```go
// Option — the atom
core.Option{Key: "name", Value: "brain"}

// Options — universal input
opts := core.NewOptions(
    core.Option{Key: "name", Value: "myapp"},
    core.Option{Key: "port", Value: 8080},
)
opts.String("name")  // "myapp"
opts.Int("port")     // 8080

// Result — universal output
core.Result{Value: svc, OK: true}
```

## IPC — Message Passing

```go
// Broadcast (fire-and-forget)
c.ACTION(messages.AgentCompleted{Agent: "codex", Status: "completed"})

// Query (first responder)
r := c.QUERY(MyQuery{Name: "brain"})

// Register handler
c.RegisterAction(func(c *core.Core, msg core.Message) core.Result {
    if ev, ok := msg.(messages.AgentCompleted); ok { /* handle */ }
    return core.Result{OK: true}
})
```

## Named Actions

```go
// Register
c.Action("git.log", func(ctx context.Context, opts core.Options) core.Result {
    dir := opts.String("dir")
    return c.Process().RunIn(ctx, dir, "git", "log", "--oneline")
})

// Invoke
r := c.Action("git.log").Run(ctx, core.NewOptions(
    core.Option{Key: "dir", Value: "/repo"},
))

// Check capability
if c.Action("process.run").Exists() { /* can run commands */ }
```

## Task Composition

```go
c.Task("deploy", core.TaskDef{
    Steps: []core.Step{
        {Action: "go.build"},
        {Action: "go.test"},
        {Action: "docker.push"},
        {Action: "ansible.deploy", Async: true},
    },
})

r := c.Task("deploy").Run(ctx, c, opts)
```

Sequential steps stop on failure. Async steps fire without blocking. `Input: "previous"` pipes the last step's output to the next.

## Process Primitive

```go
// Run a command (delegates to Action "process.run")
r := c.Process().Run(ctx, "git", "log", "--oneline")
r := c.Process().RunIn(ctx, "/repo", "go", "test", "./...")

// Permission by registration:
// No go-process registered → c.Process().Run() returns Result{OK: false}
// go-process registered → executes the command
```

## Registry[T]

Thread-safe named collection — the universal brick for all registries.

```go
r := core.NewRegistry[*MyService]()
r.Set("brain", brainSvc)
r.Get("brain")              // Result{brainSvc, true}
r.Has("brain")              // true
r.Names()                   // insertion order
r.List("process.*")         // glob match
r.Each(func(name string, svc *MyService) { ... })
r.Lock()                    // fully frozen
r.Seal()                    // no new keys, updates OK

// Cross-cutting queries
c.RegistryOf("services").Names()
c.RegistryOf("actions").List("process.*")
```

## Commands

```go
c.Command("deploy/to/homelab", core.Command{
    Action:  handler,
    Managed: "process.daemon",  // go-process provides lifecycle
})
```

Path = hierarchy. `deploy/to/homelab` becomes `myapp deploy to homelab` in CLI.

## Utilities

```go
core.ID()                    // "id-1-a3f2b1" — unique identifier
core.ValidateName("brain")  // Result{OK: true}
core.SanitisePath("../../x") // "x"
core.E("op", "msg", err)    // structured error

fs.WriteAtomic(path, data)   // write-to-temp-then-rename
fs.NewUnrestricted()         // full filesystem access
fs.Root()                    // sandbox root path
```

## Error Handling

All errors use `core.E()`:

```go
return core.E("service.Method", "what failed", underlyingErr)
```

Never use `fmt.Errorf`, `errors.New`, or `log.*`. Core handles all error reporting.

## Documentation

| Page | Covers |
|------|--------|
| [Getting Started](getting-started.md) | Installing the Core CLI, first build |
| [Configuration](configuration.md) | Config files, environment variables |
| [Workflows](workflows.md) | Common task sequences |
| [Packages](packages/index.md) | Ecosystem package reference |

## API Specification

The full API contract lives in [`docs/RFC.md`](https://forge.lthn.ai/core/go/src/branch/dev/docs/RFC.md) — 3,800+ lines covering all 21 sections, 108 findings, and implementation plans.

## Dependencies

Core is deliberately minimal:

- `dappco.re/go/core/io` — abstract storage (local, S3, SFTP, WebDAV)
- `dappco.re/go/core/log` — structured logging
- `github.com/stretchr/testify` — test assertions (test-only)

## Licence

EUPL-1.2
