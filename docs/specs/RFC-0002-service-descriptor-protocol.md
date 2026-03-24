# RFC-0002: Service Descriptor Protocol (SDP)

```
RFC:            0002
Title:          Service Descriptor Protocol (SDP)
Status:         Standard
Category:       Standards Track
Authors:        Darbs, Snider
License:        EUPL-1.2
Created:        2026-02-01
Requires:       RFC-0001
```

---

## Abstract

This document specifies the Service Descriptor Protocol (SDP), the discovery mechanism that enables Lethean clients to find and connect to Exit Node providers. SDP defines how providers publish their services and how clients query available offerings.

---

## 1. Introduction

### 1.1 Purpose

SDP solves the discovery problem in a decentralized VPN network: how do users find providers without a central authority? SDP provides a standardized format for service advertisements and a query interface for clients.

### 1.2 Design Principles

1. **Provider Autonomy** - Providers control their own listings
2. **Client Choice** - Users select providers based on published criteria
3. **Decentralization Ready** - Designed for distributed hosting
4. **Extensibility** - Support for future service types

---

## 2. Protocol Overview

### 2.1 Components

```
┌─────────────┐          ┌─────────────┐          ┌─────────────┐
│    Client   │          │ SDP Server  │          │  Exit Node  │
│             │  Query   │             │   Sync   │             │
│             │─────────>│             │<─────────│             │
│             │          │             │          │             │
│             │  Results │             │   VDP    │             │
│             │<─────────│             │─────────>│             │
└─────────────┘          └─────────────┘          └─────────────┘
```

### 2.2 Data Flow

1. **Publication**: Exit Nodes generate VDP (Virtual Data Provider) records
2. **Registration**: VDP pushed to VDP Manager
3. **Aggregation**: SDP Server collects VDPs from all managers
4. **Query**: Clients request service listings from SDP endpoint
5. **Selection**: Client chooses provider based on criteria

---

## 3. VDP Record Format

### 3.1 Provider Metadata

```json
{
  "provider": {
    "id": "<64-char hex string>",
    "name": "<display name, max 16 chars>",
    "wallet": "<LTHN wallet address>",
    "terms": "<URL to terms of service>",
    "type": "commercial|community",
    "ca": "<base64 encoded CA certificate>"
  }
}
```

### 3.2 Field Definitions

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| id | string(64) | Yes | Unique provider identifier (hex) |
| name | string(16) | Yes | Human-readable provider name |
| wallet | string | Yes | LTHN payment address |
| terms | string | No | URL or `@filename` for terms |
| type | enum | No | `commercial` or `community` |
| ca | string | Yes | PEM certificate, base64 encoded |

### 3.3 Service Definitions

```json
{
  "services": [
    {
      "id": "1A",
      "type": "proxy",
      "name": "AU Melbourne Proxy 1",
      "endpoint": "proxy.example.com",
      "port": 8080,
      "cost": "0.1",
      "speed": "100mbps",
      "location": {
        "country": "AU",
        "city": "Melbourne"
      }
    },
    {
      "id": "2A",
      "type": "vpn",
      "name": "AU Melbourne VPN 1",
      "endpoint": "vpn.example.com",
      "port": 18080,
      "cost": "0.15",
      "protocol": "openvpn"
    }
  ]
}
```

### 3.4 Service ID Convention

| Range | Type | Description |
|-------|------|-------------|
| 1A-1Z | proxy | HTTP/HTTPS proxy services |
| 2A-2Z | vpn | VPN tunnel services |
| C1A+ | client | Client-side service configs |

---

## 4. DNS Discovery

### 4.1 TXT Record Format

Providers MAY publish a DNS TXT record for decentralized discovery:

```
_lethean.example.com TXT "lv=v3;sdp=https://example.com/sdp.json;id=<provider-id>"
```

### 4.2 TXT Record Fields

| Field | Description |
|-------|-------------|
| lv | Protocol version (currently `v3`) |
| sdp | URL to provider's sdp.json file |
| id | Provider ID (64-char hex) |

### 4.3 Example

```
_lethean.exitnode.example TXT "lv=v3;sdp=https://monitor.lethean.io/sdp.json;id=7b08c778af3b28932185d7cc804b0cf399c05c9149613dc149dff5f30c8cd989"
```

---

## 5. SDP API

### 5.1 Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | /v1/services/search | Query available services |
| GET | /v1/providers/{id} | Get specific provider details |
| POST | /v1/providers | Register/update provider (authenticated) |

### 5.2 Search Query Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| type | string | Filter by service type (`proxy`, `vpn`) |
| country | string | Filter by country code (ISO 3166-1) |
| max_cost | number | Maximum cost in LTHN |
| min_speed | string | Minimum speed (e.g., `10mbps`) |

### 5.3 Search Response

```json
{
  "providers": [
    {
      "id": "efaa812b358956f93a0e324385c8b44469a99e5a82f2de327297b25d8c2ee288",
      "name": "Lethean_Re-Born",
      "wallet": "iz5HSgJUW0max8Hs2TEAacKhKA9LXLLDvc4u7yCV7Lm4iwkgFXTMFBAdtj2mqMpqy7T4BNveDQdW8LBPVxWqy94B2A6sKJXQ7",
      "services": [
        {
          "id": "1A",
          "type": "proxy",
          "name": "AU Melbourne Proxy 1",
          "cost": "0.1",
          "port": 8080
        }
      ],
      "ca": "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----"
    }
  ],
  "total": 1,
  "timestamp": "2026-02-01T12:00:00Z"
}
```

---

## 6. VDP Synchronization

### 6.1 Sync Protocol

Exit Nodes synchronize with the VDP Manager:

```
Exit Node                           VDP Manager
    │                                    │
    │  POST /vdp (VDP record)           │
    │───────────────────────────────────>│
    │                                    │
    │  200 OK + TTL                      │
    │<───────────────────────────────────│
    │                                    │
    │  ... wait (TTL - buffer) ...       │
    │                                    │
    │  POST /vdp (refresh)              │
    │───────────────────────────────────>│
```

### 6.2 Timing Parameters

| Parameter | Value | Description |
|-----------|-------|-------------|
| TTL | 3600 seconds | VDP time-to-live (1 hour) |
| Refresh Buffer | 100 seconds | Refresh before TTL expires |
| Refresh Interval | 3500 seconds | Typical refresh cycle |

### 6.3 VDP Manager Endpoints

| Endpoint | Description |
|----------|-------------|
| mgr.lethean.space | Primary VDP Manager |
| vpn2.lethean.space:8774 | WireGuard tunnel endpoint |

---

## 7. sdp.json File Format

### 7.1 Complete Example

```json
{
  "version": "3",
  "generated": "2026-02-01T12:00:00Z",
  "provider": {
    "id": "efaa812b358956f93a0e324385c8b44469a99e5a82f2de327297b25d8c2ee288",
    "name": "Lethean ReBorn",
    "wallet": "iz5HSgJUW0max8Hs2TEAacKhKA9LXLLDvc...",
    "terms": "https://provider.example/terms",
    "type": "commercial"
  },
  "ca": "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWgAw...\n-----END CERTIFICATE-----",
  "services": [
    {
      "id": "1A",
      "type": "proxy",
      "name": "AU Melbourne Proxy 1",
      "endpoint": "au-mel.provider.example",
      "port": 8080,
      "cost": "0.1",
      "firstVerificationsRequired": 1,
      "subsequentVerificationsRequired": 1,
      "speed": "100mbps",
      "location": {
        "country": "AU",
        "region": "Victoria",
        "city": "Melbourne"
      },
      "restrictions": []
    }
  ]
}
```

### 7.2 Verification Fields

| Field | Type | Description |
|-------|------|-------------|
| firstVerificationsRequired | integer (0-2) | Confirmations for first payment |
| subsequentVerificationsRequired | integer (0-1) | Confirmations for subsequent payments |

Lower values = faster connection, higher risk
Higher values = slower connection, lower risk

---

## 8. Security Considerations

### 8.1 Provider Authenticity

- Provider ID derived from cryptographic key
- CA certificate validates service endpoints
- Wallet address ties provider to blockchain identity

### 8.2 Man-in-the-Middle Protection

- SDP endpoint MUST use HTTPS
- Client MUST verify provider CA before connecting
- DNS TXT records provide secondary verification

### 8.3 Sybil Resistance

- Providers must stake resources (infrastructure)
- SLA monitoring identifies poor performers
- User reviews (future: on-chain reputation)

---

## 9. Implementation Notes

### 9.1 Generating SDP Configuration

```bash
lvmgmt --generate-sdp --wallet-address <LTHN_ADDRESS>
```

Interactive prompts:
1. Provider name (max 16 chars)
2. Service name (max 32 chars)
3. Service type (vpn/proxy)
4. Port number (1-65535)
5. Cost in LTHN
6. Confirmation requirements

Output: `/opt/lthn/etc/sdp.json`

### 9.2 Client Integration

```python
# Pseudocode for SDP query
response = http.get("https://sdp.lethean.io/v1/services/search",
                    params={"type": "proxy", "country": "AU"})
providers = response.json()["providers"]

for provider in providers:
    print(f"{provider['name']}: {provider['services'][0]['cost']} LTHN")
```

---

## 10. Future Considerations

### 10.1 Decentralized SDP

Current implementation uses centralized SDP servers. Future versions may:

- Store VDP records on Kevacoin (key-value blockchain)
- Use DHT (Distributed Hash Table) for discovery
- Enable direct peer-to-peer provider queries

### 10.2 Extended Service Types

SDP is designed to support services beyond VPN/proxy:

- Hosting services
- Storage services
- Computation services
- Custom BYOA (Bring Your Own Application)

---

## 11. References

- RFC-0001: Lethean Network Overview
- RFC-0003: Exit Node Architecture
- Lethean SDP Design (HLA) v0.1

---

## Appendix A: JSON Schema

```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "type": "object",
  "required": ["version", "provider", "services"],
  "properties": {
    "version": {"type": "string", "pattern": "^[0-9]+$"},
    "provider": {
      "type": "object",
      "required": ["id", "name", "wallet"],
      "properties": {
        "id": {"type": "string", "pattern": "^[a-f0-9]{64}$"},
        "name": {"type": "string", "maxLength": 16},
        "wallet": {"type": "string"},
        "terms": {"type": "string"},
        "type": {"enum": ["commercial", "community"]}
      }
    },
    "services": {
      "type": "array",
      "items": {
        "type": "object",
        "required": ["id", "type", "port", "cost"],
        "properties": {
          "id": {"type": "string", "pattern": "^[12C][A-Z0-9]+$"},
          "type": {"enum": ["proxy", "vpn"]},
          "port": {"type": "integer", "minimum": 1, "maximum": 65535},
          "cost": {"type": "string"}
        }
      }
    }
  }
}
```
