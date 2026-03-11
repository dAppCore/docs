---
title: go-ws
description: WebSocket hub for real-time streaming with channel pub/sub and Redis bridge
---

# go-ws

`forge.lthn.ai/core/go-ws`

WebSocket server implementing the hub pattern for real-time streaming. Provides named channel pub/sub, token-based authentication on the upgrade handshake, automatic reconnection with exponential backoff on the client side, and a Redis pub/sub bridge for multi-instance coordination.

## Key Types

- `AuthResult` — outcome of a WebSocket authentication attempt
- `APIKeyAuthenticator` — authenticates connections via API key header or query parameter
- `BearerTokenAuth` — authenticates connections via Bearer token
- `RedisConfig` — configuration for the Redis pub/sub bridge
