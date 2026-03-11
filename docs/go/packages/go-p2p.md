---
title: go-p2p
description: P2P mesh networking for Lethean with encrypted transport and intent routing
---

# go-p2p

`forge.lthn.ai/core/go-p2p`

Peer-to-peer mesh networking library for the Lethean network. Provides Ed25519 node identity, encrypted WebSocket transport with HMAC-SHA256 challenge-response authentication, and KD-tree based peer selection optimised for latency, hop count, geography, and reliability.

Implements the UEPS wire protocol (RFC-021) with a TLV builder/reader, intent routing with a threat circuit breaker, and TIM deployment bundle encryption with Zip Slip and decompression-bomb defences.
