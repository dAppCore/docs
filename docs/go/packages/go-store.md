---
title: go-store
description: Group-namespaced SQLite key-value store with TTL, quotas, and reactive events
---

# go-store

`forge.lthn.ai/core/go-store`

SQLite-backed key-value store with group namespacing, TTL expiry, and quota enforcement. Operates in WAL mode for concurrent read performance. Supports scoped stores for multi-tenant isolation, reactive Watch/Unwatch subscriptions, and OnChange callbacks. Designed as an integration point for go-ws real-time streaming.

## Key Types

- `Event` — change event emitted on key create, update, or delete
- `Watcher` — subscription handle for reactive change notifications
- `QuotaConfig` — per-namespace limits (max keys, max value size, total size)
- `ScopedStore` — namespace-isolated view of the store for multi-tenant use
