# RFC-021: Core Platform Architecture

```
RFC:            021
Title:          Core Platform Architecture
Status:         Standards Track
Category:       Informational
Authors:        Snider, Cladius Maximus
License:        EUPL-1.2
Created:        2026-03-14
Requires:       RFC-0001 through RFC-0005, RFC-001 through RFC-020
```

---

## Abstract

This document specifies how the 25 preceding RFCs compose into a single coherent platform via the Core package ecosystem. It defines the Service Provider as the universal unit of functionality, the TRIX container as the universal distribution format, HLCRF as the universal layout primitive, UEPS as the universal consent layer, and OpenAPI as the universal polyglot contract. Together these form a Web3 application standard where identity is sovereign, compute is local, and every component is replaceable.

---

## 1. Design Principles

1. **The binary IS the platform** — a single Go binary embeds all layers
2. **Providers are the universal unit** — everything is a service provider
3. **OpenAPI is the polyglot contract** — Go, PHP, TypeScript speak the same API
4. **TRIX is the universal envelope** — all data shares one container format
5. **HLCRF is the universal layout** — one string describes any UI structure
6. **UEPS is the universal consent** — every packet carries intent metadata
7. **Language is irrelevant** — the interface is the boundary, not the runtime

---

## 2. Layer Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│ Layer 7: RENDERING                                               │
│   HLCRF compositor (RFC-001), Angular custom elements,           │
│   go-html WASM, Web Components, CoreDeno sidecar (core/ts)      │
├─────────────────────────────────────────────────────────────────┤
│ Layer 6: ANALYSIS                                                │
│   LEM ethics scoring, go-i18n GrammarImprint,                    │
│   Poindexter spatial indexing, OpenBrain vector store            │
├─────────────────────────────────────────────────────────────────┤
│ Layer 5: STORAGE                                                 │
│   DataNode in-memory FS (RFC-013), go-io Medium (local/S3/SSH), │
│   go-store SQLite KV, Borg blob storage                          │
├─────────────────────────────────────────────────────────────────┤
│ Layer 4: COMPUTE                                                 │
│   TIM containers (RFC-014), go-process daemon registry,          │
│   CoreDeno sandbox, FrankenPHP embedded runtime                  │
├─────────────────────────────────────────────────────────────────┤
│ Layer 3: CRYPTO                                                  │
│   Enchantrix sigils (RFC-009), TRIX containers (RFC-010),        │
│   SMSG media (RFC-012), STIM encrypted TIM (RFC-015),            │
│   STMF secure forms (RFC-019), LTHN hash (RFC-007)               │
├─────────────────────────────────────────────────────────────────┤
│ Layer 2: PROTOCOL                                                │
│   UEPS consent-gated TLV, SDP service discovery (RFC-0002),     │
│   Payment dispatcher (RFC-0004), MCP tool protocol               │
├─────────────────────────────────────────────────────────────────┤
│ Layer 1: IDENTITY                                                │
│   Wallet-as-identity, Ed25519 signing, HNS TLD addressing,      │
│   go-crypt trust policies                                        │
└─────────────────────────────────────────────────────────────────┘
```

Each layer has implementations in Go, PHP, and/or TypeScript. The layer boundaries are defined by interfaces, not languages.

---

## 3. The Service Provider

### 3.1 Definition

A Service Provider is any component that:

1. Declares an OpenAPI spec (the contract)
2. Registers routes on the API engine (the implementation)
3. Optionally declares an HLCRF layout (the UI)
4. Optionally emits WebSocket events (real-time data)
5. Optionally declares MCP tool descriptions (AI integration)

### 3.2 Go Providers

Go providers implement `api.RouteGroup` directly:

```go
type Provider interface {
    Name() string
    BasePath() string
    RegisterRoutes(rg *gin.RouterGroup)
}
```

Extensions: `Streamable` (WS events), `Describable` (OpenAPI), `Renderable` (custom element tag).

### 3.3 PHP Providers

PHP providers are Laravel modules implementing the provider contract. The PHP provider runs inside FrankenPHP (embedded in the Go binary) or as a standalone Laravel app. The Go API layer discovers its OpenAPI spec and creates a reverse proxy route group.

### 3.4 TypeScript Providers

TypeScript providers run inside CoreDeno (core/ts sidecar) and expose routes via the gRPC bridge. Same pattern — OpenAPI spec + reverse proxy.

### 3.5 Provider Distribution

Providers are packaged as TRIX containers:

```
provider.trix
├── .core/view.yml          # Manifest: name, HLCRF variant, permissions
├── openapi.json            # OpenAPI spec (the contract)
├── element.js              # Custom element bundle (Renderable)
├── src/                    # Implementation (Go/PHP/TS source)
└── sign                    # Ed25519 signature
```

For secure distribution: `provider.stim` (encrypted TRIX via RFC-015).

### 3.6 Provider Discovery

```yaml
# .core/config.yaml
providers:
  brain:
    enabled: true
  vpn:
    enabled: true
    endpoint: http://localhost:8774
  studio:
    enabled: true
    package: forge.lthn.ai/core/php-studio
```

The registry loads providers from:
1. Go packages registered via `engine.Register()`
2. `.core/providers/*.yaml` for polyglot providers
3. Marketplace (git-based, signed manifests)

---

## 4. The Application Shell (core/ide)

### 4.1 Architecture

core/ide is a Wails 3 systray application that assembles providers:

```
┌─────────────────────────────────────────────────────────┐
│                    Wails 3 (Systray)                     │
│  ┌──────────┐  ┌──────────────────────────────────────┐ │
│  │ Systray  │  │ Angular Shell                        │ │
│  │ (life-   │  │  ┌─────────────────────────────────┐ │ │
│  │  cycle)  │  │  │ HLCRF Layout                    │ │ │
│  │          │  │  │  H: [nav-bar]                   │ │ │
│  │          │  │  │  L: [provider-list]              │ │ │
│  │          │  │  │  C: [<active-provider-panel>]    │ │ │
│  │          │  │  │  F: [status-bar]                │ │ │
│  │          │  │  └─────────────────────────────────┘ │ │
│  └──────────┘  └──────────────────────────────────────┘ │
│                                                          │
│  ┌──────────────────────────────────────────────────────┐│
│  │ core.Core (DI container + IPC bus)                   ││
│  │  ├─ core/api Engine (provider registry + Gin router) ││
│  │  ├─ core/mcp Service (MCP server + brain + tools)    ││
│  │  ├─ core/gui display.Service (Wails platform bridge) ││
│  │  └─ go-ws Hub (Angular ↔ providers real-time)        ││
│  └──────────────────────────────────────────────────────┘│
└──────────────────────────────────────────────────────────┘
```

### 4.2 Three Access Patterns

The same providers are accessible three ways:

| Access | Transport | Consumer |
|--------|-----------|----------|
| **REST API** | HTTP (Gin) | Web apps, curl, SDKs |
| **MCP** | stdio/TCP/Unix | Claude Code, AI agents |
| **GUI** | WebSocket + custom elements | Angular shell in Wails |

### 4.3 Dual Identity

core/ide serves two roles simultaneously:

1. **Developer tool** — MCP server for AI agents, OpenBrain recall, build/deploy providers, process management dashboard
2. **Network client** — VPN tunnel manager (RFC-0005), payment wallet (RFC-0004), service discovery (RFC-0002), media player (RFC-011)

Both roles use the same provider framework. VPN is just another provider.

---

## 5. RFC Mapping to Core Packages

### 5.1 Network Protocol

| RFC | Package | Implementation |
|-----|---------|----------------|
| 0001: Network | go-p2p | P2P mesh, UEPS wire protocol |
| 0002: SDP | go-scm/forge | Service descriptor queries |
| 0003: Exit Node | go-process | Managed daemon processes |
| 0004: Payment | go-blockchain | Wallet RPC, transaction submission |
| 0005: Client | **core/ide** | The application itself |

### 5.2 Platform

| RFC | Package | Implementation |
|-----|---------|----------------|
| 001: HLCRF | core/php + go-html + core/gui | Polyglot layout |
| 002: Events | core/php + core/go IPC bus | Module lifecycle |
| 003: Config | core/php + go-config | Configuration |
| 004: Entitlements | core/php-tenant | Feature gating |
| 005: Commerce | core/php-commerce | Billing, subscriptions |
| 006: Compound SKU | core/php-commerce | Product structure |

### 5.3 Cryptography

| RFC | Package | Implementation |
|-----|---------|----------------|
| 007: LTHN Hash | go-crypt | Quasi-salted hash |
| 008: Pre-obfuscation | go-crypt | AEAD pre-layer |
| 009: Sigil | Enchantrix | Transformation chain |
| 010: TRIX | Enchantrix pkg/trix | Binary container format |
| 016: TRIX PGP | Enchantrix | PGP encryption variant |
| 017: Key derivation | go-crypt | Key stretching |

### 5.4 Containers

| RFC | Package | Implementation |
|-----|---------|----------------|
| 011: OSS-DRM | Borg pkg/smsg | Password-as-license media |
| 012: SMSG | Borg pkg/smsg | Encrypted media container |
| 013: DataNode | Borg pkg/datanode + go-io | In-memory filesystem |
| 014: TIM | Borg pkg/tim | OCI container bundles |
| 015: STIM | Borg pkg/tim | Encrypted containers |
| 018: Borgfile | Borg cmd/compile | Container compilation |

### 5.5 Interfaces

| RFC | Package | Implementation |
|-----|---------|----------------|
| 019: STMF | Borg pkg/stmf + WASM | Secure form encryption |
| 020: WASM API | Borg pkg/wasm + go-html | Browser crypto API |

---

## 6. UEPS Integration

The Unified Ethical Protocol Stack (Mining pkg/ueps) wraps every inter-provider communication:

```
┌─────────────────────────────────────────┐
│ UEPS Packet                             │
│  Tag 0x01: Version (0x09 = IPv9)        │
│  Tag 0x02: Current Layer (sender)       │
│  Tag 0x03: Target Layer (recipient)     │
│  Tag 0x04: Intent ID (semantic token)   │
│  Tag 0x05: Threat Score (0-65535)       │
│  Tag 0x06: HMAC (SHA-256 signature)     │
│  Tag 0xFF: Payload (the actual data)    │
└─────────────────────────────────────────┘
```

Every API call between providers, every MCP tool invocation, every WS event can carry UEPS metadata. The Intent ID maps to go-i18n's semantic tokens. The Threat Score is updated by LEM's ethics scoring. The HMAC binds the packet to a shared secret (wallet-derived for network, config-derived for local).

Optional for local providers (localhost doesn't need consent gates). Required for network providers (RFC-0003 exit nodes, remote services).

---

## 7. HLCRF as Universal Layout

### 7.1 Three Implementations, One String

The variant string `H[LC]CF` produces identical layouts in:

| Runtime | Implementation | Package |
|---------|---------------|---------|
| **PHP** | `Layout::make('H[LC]CF')` | core/php Front\Components\Layout |
| **Go** | `hlcrf.New("H[LC]CF")` | go-html |
| **TypeScript** | `<core-layout variant="H[LC]CF">` | core/ts Web Component |
| **Angular** | `<app-layout [variant]="'H[LC]CF'">` | core/ide frontend |

### 7.2 Provider Layout Declaration

Each Renderable provider declares its HLCRF variant in `.core/view.yml`:

```yaml
code: brain-panel
name: OpenBrain Panel
element: core-brain-panel
layout: HCF
slots:
  H: search-bar
  C: results-list
  F: status-bar
```

---

## 8. Content Distribution

### 8.1 TRIX Family

All content shares the TRIX container format (RFC-010):

| Magic | Purpose | Encryption | RFC |
|-------|---------|------------|-----|
| TRIX | Generic container | Optional | 010 |
| SMSG | Encrypted media | ChaCha20-Poly1305 | 012 |
| STIM | Encrypted container | Dual ChaCha20-Poly1305 | 015 |
| STMF | Encrypted form data | X25519 + ChaCha20-Poly1305 | 019 |

### 8.2 Provider as SMSG

A provider's custom element + assets can be packaged as SMSG:

- **Manifest**: Provider metadata (name, version, author, links)
- **Payload**: The JS bundle + OpenAPI spec + config
- **License**: Password-as-license for commercial providers
- **Distribution**: CDN, IPFS, Git marketplace — encrypted at rest

---

## 9. The Binary

### 9.1 Composition

```
core-ide binary (~60MB)
├── Go runtime
├── core/go (DI + lifecycle)
├── core/api (Gin + provider registry + OpenAPI)
├── core/mcp (MCP server + brain + subsystems)
├── core/gui (16 IPC packages + display service)
├── core/ts (CoreDeno sidecar manager)
├── go-blockchain (wallet + chain sync)
├── go-process (daemon registry)
├── go-io (sandboxed filesystem)
├── go-crypt (Enchantrix bindings)
├── go-i18n (grammar engine)
├── go-html (HLCRF + WASM codegen)
├── go-p2p (UEPS + network mesh)
├── Wails 3 (WebView2 + systray)
├── FrankenPHP (PHP 8.4 ZTS, optional)
└── Angular frontend (embedded via //go:embed)
```

One binary. No external dependencies. Runs on macOS, Linux, Windows.

### 9.2 Without GUI

go-config `gui.enabled: false` skips Wails. Core still runs all services — MCP server, API engine, brain, providers.

---

## 10. Polyglot Contract

### 10.1 OpenAPI as the Boundary

The only thing that crosses language boundaries is the OpenAPI spec:

```
Go provider:   implements RouteGroup directly
PHP provider:  publishes OpenAPI spec, reverse proxy
TS provider:   publishes OpenAPI spec, CoreDeno gRPC bridge
```

### 10.2 SDK Generation

From the assembled OpenAPI spec, auto-generate client libraries for TypeScript, Python, PHP, and Go. One API, every language.

---

## 11. Security Model

### 11.1 Provider Isolation

| Isolation | Mechanism | RFC |
|-----------|-----------|-----|
| Filesystem | go-io Medium sandbox | — |
| Process | TIM OCI containers | 014 |
| Network | CoreDeno permission gates | — |
| Crypto | Per-provider STIM encryption | 015 |
| Consent | UEPS intent tokens per packet | — |
| Identity | Ed25519 signed manifests | — |

### 11.2 Zero-Trust Distribution

Providers distributed as STIM are encrypted at rest. The marketplace serves encrypted blobs. Only the purchaser's password decrypts. The signature verifies the publisher. The UEPS headers carry consent metadata.

---

## 12. Implementation Status

### 12.1 Complete

All 25 preceding RFCs have implementations. The provider framework (Phase 1) is live. core/ide is modernised. The API polyglot merge is in progress.

### 12.2 In Progress

- core/api polyglot merge (Go + PHP in one repo)
- Provider GUI consumer (Phase 2 — Renderable discovery)
- Polyglot providers (Phase 3 — PHP/TS via OpenAPI)

### 12.3 Future

- Provider marketplace (STIM distribution + git registry)
- Network UEPS enforcement (go-p2p integration)
- VPN client provider (RFC-0005 as a service provider)
- Mobile (Wails mobile or PWA)

---

## 13. Relationship to Existing Standards

| Standard | Relationship |
|----------|-------------|
| OCI Runtime Spec | TIM bundles are OCI-compatible |
| OpenAPI 3.1 | Provider contract format |
| MCP | AI agent integration |
| Web Components v1 | Provider UI elements |
| WireGuard | Network transport option |
| ChaCha20-Poly1305 (RFC 8439) | All encryption |
| X25519 (RFC 7748) | Key exchange (STMF) |
| CSS Grid / Flexbox | HLCRF rendering |
| fs.FS (Go stdlib) | DataNode interface |

---

## 14. References

- RFC-0001 through RFC-0005: Lethean network protocol
- RFC-001 through RFC-020: Implementation specifications
- Borg: forge.lthn.ai/Snider/Borg
- Enchantrix: forge.lthn.ai/Snider/Enchantrix
- Poindexter: forge.lthn.ai/Snider/Poindexter
- Mining: forge.lthn.ai/Snider/Mining
- core.help: Ecosystem documentation

---

## Licence

EUPL-1.2

**Viva La OpenSource**
