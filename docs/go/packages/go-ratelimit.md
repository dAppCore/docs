---
title: go-ratelimit
description: Provider-agnostic sliding window rate limiter for LLM API calls
---

# go-ratelimit

`forge.lthn.ai/core/go-ratelimit`

Provider-agnostic rate limiter for LLM API calls using sliding window counters. Enforces RPM (requests per minute), TPM (tokens per minute), and RPD (requests per day) quotas on a per-model basis. Ships with default profiles for Gemini, OpenAI, Anthropic, and local inference. Supports YAML or SQLite persistence for quota state.

## Key Types

- `ModelQuota` — per-model rate limits (RPM, TPM, RPD)
- `ProviderProfile` — named provider with its model quotas
- `Config` — rate limiter configuration (providers, persistence backend)
- `UsageStats` — current usage counters and remaining quota
