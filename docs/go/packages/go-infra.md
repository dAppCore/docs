---
title: go-infra
description: Infrastructure management with DNS provider integrations
---

# go-infra

`forge.lthn.ai/core/go-infra`

Infrastructure management library with provider integrations for DNS and cloud resources. Currently supports CloudNS as a DNS provider with full zone and record management. Includes retry logic and structured API error handling.

## Key Types

- `RetryConfig` — retry policy for transient API failures
- `APIClient` — base HTTP client with authentication and retry support
- `CloudNSClient` — CloudNS API client for DNS operations
- `CloudNSZone` — DNS zone representation
- `CloudNSRecord` — individual DNS record (A, AAAA, CNAME, MX, TXT, etc.)
