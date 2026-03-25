# RFC Index — Lethean Ecosystem Specifications

> Request For Contribution — design specifications that define how the ecosystem works.
> Each RFC is detailed enough that an agent can implement the described system from the document alone.

## How to Read

Start with the category that matches your task. Each RFC is self-contained — you don't need to read them in order. If you're contributing code, read RFC-025 (Agent Experience) first — it defines the conventions all code must follow.

## Core Framework

| RFC | Title | Status |
|-----|-------|--------|
| [RFC-021](specs/RFC-021-CORE-PLATFORM-ARCHITECTURE.md) | Core Platform Architecture | Draft |
| [RFC-025](specs/RFC-025-AGENT-EXPERIENCE.md) | Agent Experience (AX) Design Principles | Draft |
| [RFC-002](specs/RFC-002-EVENT-DRIVEN-MODULES.md) | Event-Driven Module Loading | Implemented |
| [RFC-003](specs/RFC-003-CONFIG-CHANNELS.md) | Config Channels | Implemented |
| [RFC-004](specs/RFC-004-ENTITLEMENTS.md) | Entitlements and Feature System | Implemented |
| [RFC-024](specs/RFC-024-ISSUE-TRACKER.md) | Issue Tracker and Sprint System | Draft |

## Commerce and Products

| RFC | Title | Status |
|-----|-------|--------|
| [RFC-005](specs/RFC-005-COMMERCE-MATRIX.md) | Commerce Entity Matrix | Implemented |
| [RFC-006](specs/RFC-006-COMPOUND-SKU.md) | Compound SKU Format | Implemented |
| [RFC-001](specs/RFC-001-HLCRF-COMPOSITOR.md) | HLCRF Compositor | Implemented |

## Cryptography and Security

| RFC | Title | Status |
|-----|-------|--------|
| [RFC-011](specs/RFC-011-OSS-DRM.md) | Open Source DRM for Independent Artists | Proposed |
| [RFC-007](specs/RFC-007-LTHN-HASH.md) | LTHN Quasi-Salted Hash Algorithm | Implemented |
| [RFC-008](specs/RFC-008-PRE-OBFUSCATION-LAYER.md) | Pre-Obfuscation Layer Protocol for AEAD Ciphers | Implemented |
| [RFC-009](specs/RFC-009-SIGIL-TRANSFORMATION.md) | Sigil Transformation Framework | Implemented |
| [RFC-010](specs/RFC-010-TRIX-CONTAINER.md) | TRIX Binary Container Format | Implemented |
| [RFC-015](specs/RFC-015-STIM.md) | STIM Encrypted Container Format | Implemented |
| [RFC-016](specs/RFC-016-TRIX-PGP.md) | TRIX PGP Encryption Format | Implemented |
| [RFC-017](specs/RFC-017-LTHN-KEY-DERIVATION.md) | LTHN Key Derivation | Implemented |
| [RFC-019](specs/RFC-019-STMF.md) | STMF Secure To-Me Form | Implemented |
| [RFC-020](specs/RFC-020-WASM-API.md) | WASM Decryption API | Implemented |

## Data and Messaging

| RFC | Title | Status |
|-----|-------|--------|
| [RFC-012](specs/RFC-012-SMSG-FORMAT.md) | SMSG Container Format | Implemented |
| [RFC-013](specs/RFC-013-DATANODE.md) | DataNode In-Memory Filesystem | Implemented |
| [RFC-014](specs/RFC-014-TIM.md) | Terminal Isolation Matrix (TIM) | Implemented |
| [RFC-018](specs/RFC-018-BORGFILE.md) | Borgfile Compilation | Implemented |

## Lethean Network (Legacy)

| RFC | Title | Status |
|-----|-------|--------|
| [RFC-0001](specs/RFC-0001-network-overview.md) | Lethean Network Overview | Implemented |
| [RFC-0002](specs/RFC-0002-service-descriptor-protocol.md) | Service Descriptor Protocol (SDP) | Implemented |
| [RFC-0003](specs/RFC-0003-exit-node-architecture.md) | Exit Node Architecture | Implemented |
| [RFC-0004](specs/RFC-0004-payment-dispatcher-protocol.md) | Payment and Dispatcher Protocol | Implemented |
| [RFC-0005](specs/RFC-0005-client-protocol.md) | Client Protocol | Implemented |

## Contributing

New RFCs follow the numbering scheme `RFC-NNN-TITLE.md` (3-digit, uppercase title). Use RFC-011 (OSS DRM) as the reference for detail level — an agent should be able to implement the system from the document alone.

All contributions must follow [RFC-025: Agent Experience](specs/RFC-025-AGENT-EXPERIENCE.md).
