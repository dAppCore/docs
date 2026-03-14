---
title: GUI
description: IPC-based desktop GUI framework abstracting Wails v3 through typed services.
---

# GUI

**Module**: `forge.lthn.ai/core/gui`
**Language**: Go 1.25
**Licence**: EUPL-1.2

CoreGUI is an abstraction layer over Wails v3 — a "display server" that provides a stable API contract for desktop applications. Apps never import Wails directly; CoreGUI defines Platform interfaces that insulate all Wails types behind adapter boundaries. If Wails breaks, it is fixed in one place.

## Architecture

CoreGUI follows a three-layer stack:

```
IPC Bus (core/go ACTION / QUERY / PERFORM)
    |
Service (core.ServiceRuntime + business logic)
    |
Platform Interface (Wails v3 adapter, injected at startup)
```

Each feature area is a `core.Service` registered with the DI container. Services communicate via typed IPC messages — queries return data, tasks mutate state, and actions are fire-and-forget broadcasts. No service imports another directly; the display orchestrator bridges IPC events to a WebSocket pub/sub channel for TypeScript frontends.

```
pkg/display (orchestrator)
+-- imports pkg/window, pkg/systray, pkg/menu         (message types only)
+-- imports pkg/clipboard, pkg/dialog, pkg/notification
+-- imports pkg/screen, pkg/environment, pkg/dock, pkg/lifecycle
+-- imports pkg/keybinding, pkg/contextmenu, pkg/browser
+-- imports pkg/webview, pkg/mcp
+-- imports core/go (DI, IPC) + go-config
```

No circular dependencies. Sub-packages do not import each other or the orchestrator.

## Packages

| Package | Description |
|---------|-------------|
| `pkg/display` | Orchestrator — owns Wails app, config, WSEventManager bridge |
| `pkg/window` | Window lifecycle, tiling, snapping, layouts, state persistence |
| `pkg/screen` | Screen enumeration, primary detection, point-to-screen queries |
| `pkg/clipboard` | Clipboard read/write (text) |
| `pkg/dialog` | File open/save, directory select, message dialogs |
| `pkg/notification` | Native system notifications with dialog fallback |
| `pkg/systray` | System tray icon, tooltip, dynamic menus, panel window |
| `pkg/menu` | Application menu builder (structure only, handlers injected) |
| `pkg/keybinding` | Global keyboard shortcuts with accelerator syntax |
| `pkg/contextmenu` | Right-click context menu registration and management |
| `pkg/dock` | macOS dock icon visibility, badge (taskbar badge on Windows) |
| `pkg/lifecycle` | Application lifecycle events (startup, shutdown, theme change) |
| `pkg/browser` | Open URLs and files in the system default browser |
| `pkg/environment` | OS/platform info, dark mode detection, accent colour, theme events |
| `pkg/webview` | CDP-based WebView interaction — JS eval, DOM queries, screenshots |
| `pkg/mcp` | MCP display subsystem exposing ~74 tools across all packages |

## Service Registration

Each package exposes a `Register(platform)` factory. The orchestrator creates Wails adapters and passes them to each service:

```go
wailsApp := application.New(application.Options{...})

core.New(
    core.WithService(display.Register(wailsApp)),
    core.WithService(window.Register(window.NewWailsPlatform(wailsApp))),
    core.WithService(systray.Register(systray.NewWailsPlatform(wailsApp))),
    core.WithService(menu.Register(menu.NewWailsPlatform(wailsApp))),
    core.WithService(clipboard.Register(clipPlatform)),
    core.WithService(screen.Register(screenPlatform)),
    // ... remaining services
    core.WithServiceLock(),
)
```

Display registers first (owns config via `go-config`). Sub-services query their config section during `OnStartup`. Shutdown runs in reverse order.

## Platform Insulation

Each sub-package defines a `Platform` interface — the adapter contract. Wails types never leak past this boundary:

```go
// pkg/window/platform.go
type Platform interface {
    CreateWindow(opts PlatformWindowOptions) PlatformWindow
    GetWindows() []PlatformWindow
}
```

Wails adapter implementations live alongside each package (e.g. `pkg/window/wails.go`). Mock implementations enable testing without a Wails runtime.

## IPC Message Pattern

Services define typed message structs in a `messages.go` file:

- **Query** — read-only, returns data (e.g. `QueryWindowList`, `QueryTheme`, `QueryAll`)
- **Task** — side-effects, returns result (e.g. `TaskOpenWindow`, `TaskSetText`, `TaskSend`)
- **Action** — fire-and-forget broadcast (e.g. `ActionWindowOpened`, `ActionThemeChanged`)

The display orchestrator's `HandleIPCEvents` converts IPC actions to WebSocket events for TypeScript apps. Inbound WebSocket messages are translated to IPC tasks/queries, with request ID correlation for responses.

## Config

Configuration lives at `~/.core/gui/config.yaml`, loaded via `go-config`. Top-level keys map to service names:

```yaml
window:
  state_file: window_state.json
  default_width: 1024
  default_height: 768
systray:
  icon: apptray.png
  tooltip: "Core GUI"
menu:
  show_dev_tools: true
```

The display orchestrator is the single writer to disk. Sub-services read via `QueryConfig` and save via `TaskSaveConfig`.

## MCP Integration

`pkg/mcp` is an MCP subsystem that translates tool calls into IPC messages across all packages. It structurally satisfies `core/mcp`'s `Subsystem` interface (no import required):

```go
guiSub := guimcp.New(coreInstance)
mcpSvc, _ := coremcp.New(coremcp.WithSubsystem(guiSub))
```

Tool categories include window management, layout control, screen queries, clipboard, dialogs, notifications, tray, environment, keybinding, context menus, dock, lifecycle, browser, and full WebView interaction (eval, click, type, navigate, screenshot, DOM queries).

## Repository

- **Source**: [forge.lthn.ai/core/gui](https://forge.lthn.ai/core/gui)
