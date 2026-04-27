# RFC-0001: Lethean Network Overview

```
RFC:            0001
Title:          Lethean Network Overview
Status:         Standard
Category:       Informational
Authors:        Darbs, Snider
License:        EUPL-1.2
Created:        2026-02-01
Replaces:       N/A
```

---

## Abstract

This document describes the Lethean Network, a decentralized Virtual Private Network (VPN) and proxy service built on blockchain technology. The network enables peer-to-peer connectivity services where providers (Exit Nodes) offer bandwidth and users pay directly using LTHN cryptocurrency, without intermediaries.

---

## 1. Introduction

### 1.1 Purpose

Lethean provides privacy-preserving internet access through a decentralized network of Exit Nodes. Unlike traditional VPN services that rely on centralized providers, Lethean distributes trust across independent node operators who are compensated directly by users via blockchain transactions.

### 1.2 Design Goals

1. **Decentralization** - No single point of control or failure
2. **Privacy** - Untraceable payments and connections
3. **Censorship Resistance** - No central registration authority
4. **Economic Sustainability** - Direct provider compensation
5. **Open Source** - Fully auditable codebase (EUPL-1.2)

### 1.3 Terminology

| Term | Definition |
|------|------------|
| **Exit Node** | Server providing VPN/proxy services to clients |
| **SDP** | Service Descriptor Protocol - discovery mechanism |
| **LTHN** | Lethean token - native cryptocurrency |
| **Dispatcher** | Exit Node component handling authentication and routing |
| **VDP** | Virtual Data Provider - node metadata record |

---

## 2. Network Architecture

### 2.1 Layer Model

```
┌─────────────────────────────────────────────────────────────┐
│                    APPLICATION LAYER                        │
│         (Lethean Wallet GUI, CLI clients, lvpnc)           │
├─────────────────────────────────────────────────────────────┤
│                     SERVICE LAYER                           │
│              (SDP Discovery, VDP Management)                │
├─────────────────────────────────────────────────────────────┤
│                    TRANSPORT LAYER                          │
│        (Exit Nodes: HAProxy, TinyProxy, OpenVPN)           │
├─────────────────────────────────────────────────────────────┤
│                    PAYMENT LAYER                            │
│           (LTHN Blockchain, Wallet RPC, Dispatcher)        │
├─────────────────────────────────────────────────────────────┤
│                   CONSENSUS LAYER                           │
│              (Hybrid PoW/PoS, letheand daemon)             │
└─────────────────────────────────────────────────────────────┘
```

### 2.2 Network Participants

#### 2.2.1 Users (Clients)
- Discover available Exit Nodes via SDP
- Pay for services using LTHN
- Connect through selected Exit Nodes

#### 2.2.2 Exit Node Operators (Providers)
- Run infrastructure providing VPN/proxy services
- Publish service offerings to SDP
- Receive direct LTHN payments from users

#### 2.2.3 Daemon Operators
- Run blockchain nodes (letheand)
- Maintain network consensus
- Eligible for daemon rewards (see RFC-0004)

#### 2.2.4 Governors
- Participate in network governance
- Vote on protocol changes
- Manage community direction

### 2.3 Core Components

| Component | Purpose | Reference |
|-----------|---------|-----------|
| letheand | Blockchain daemon | Consensus layer |
| lethean-wallet-cli | Wallet operations | Payment layer |
| Dispatcher (lthnvpnd) | Exit Node orchestration | RFC-0003 |
| SDP Server | Service discovery | RFC-0002 |
| lvpnc | Client software | RFC-0005 |

---

## 3. Blockchain Foundation

### 3.1 Consensus Mechanism

Lethean uses a hybrid Proof-of-Work / Proof-of-Stake consensus:

- **Algorithm**: Argon2id (Chukwav2 variant)
- **Block Time**: ~120 seconds
- **Privacy**: CryptoNote-based (untraceable transactions)

### 3.2 Token Economics

| Parameter | Value |
|-----------|-------|
| Ticker | LTHN |
| Total Supply | 1,000,000,000 (1B) |
| Decimals | 8 |
| Emission | Decreasing curve |

### 3.3 Blockchain Role

The blockchain provides:

1. **Payment Settlement** - Direct user-to-provider transactions
2. **Identity** - Wallet addresses as pseudonymous identities
3. **Agreements** - Smart contract capability for service terms
4. **Audit Trail** - Immutable record of provider SLA performance

---

## 4. Data Flow

### 4.1 Service Discovery

```
User                    SDP Server                Exit Node
  │                          │                        │
  │  GET /v1/services/search │                        │
  │─────────────────────────>│                        │
  │                          │                        │
  │  Provider List (JSON)    │                        │
  │<─────────────────────────│                        │
  │                          │                        │
  │  Select Provider         │                        │
  │                          │                        │
```

### 4.2 Payment & Connection

```
User                    Blockchain               Exit Node
  │                          │                        │
  │  Send LTHN Payment       │                        │
  │─────────────────────────>│                        │
  │                          │  Transaction Confirmed │
  │                          │───────────────────────>│
  │                          │                        │
  │  Connect (TCP 8880)      │                        │
  │─────────────────────────────────────────────────>│
  │                          │                        │
  │  Dispatcher Validates    │                        │
  │<─────────────────────────────────────────────────│
  │                          │                        │
  │  Tunnel Established      │                        │
  │<════════════════════════════════════════════════>│
```

### 4.3 Provider Registration

```
Exit Node               VDP Manager              SDP Server
  │                          │                        │
  │  Generate VDP            │                        │
  │  (lvmgmt --generate-sdp) │                        │
  │                          │                        │
  │  Push VDP                │                        │
  │─────────────────────────>│                        │
  │                          │                        │
  │                          │  Sync (hourly)         │
  │                          │───────────────────────>│
  │                          │                        │
  │  Refresh (every 3500s)   │                        │
  │─────────────────────────>│                        │
```

---

## 5. Security Model

### 5.1 Threat Model

Lethean protects against:

1. **Surveillance** - ISPs cannot see destination traffic
2. **Censorship** - No central authority to block
3. **Payment Tracking** - Untraceable LTHN transactions
4. **Provider Collusion** - Users choose their own providers

### 5.2 Trust Assumptions

- Users trust their selected Exit Node operator
- Exit Node operators trust the blockchain for payments
- No global trust authority required

### 5.3 Privacy Properties

| Property | Mechanism |
|----------|-----------|
| Payment Privacy | CryptoNote ring signatures |
| Connection Privacy | TLS + VPN/Proxy tunneling |
| Identity Privacy | Pseudonymous wallet addresses |

---

## 6. Governance

### 6.1 Current Structure

The Lethean project is maintained by an Open Source Software (OSS) team:

- **Darbs** - Co-owner, technical architecture
- **Snider** - Co-owner, project lead

### 6.2 License

All Lethean software is released under the European Union Public License (EUPL-1.2), ensuring:

- Freedom to use, modify, and distribute
- Copyleft protection for derivatives
- Compatibility with other OSS licenses

### 6.3 Decision Making

Protocol changes follow this process:

1. RFC proposal submitted
2. Community discussion period
3. Implementation by maintainers
4. Network upgrade coordination

---

## 7. Related RFCs

| RFC | Title | Status |
|-----|-------|--------|
| RFC-0002 | Service Descriptor Protocol (SDP) | Standard |
| RFC-0003 | Exit Node Architecture | Standard |
| RFC-0004 | Payment & Dispatcher Protocol | Standard |
| RFC-0005 | Client Protocol | Standard |

---

## 8. References

### 8.1 Implementations

- **letheand**: https://github.com/letheanVPN/lethean
- **Exit Node**: https://github.com/letheanVPN/lvpn
- **dAppServer**: https://github.com/dAppServer

### 8.2 External Dependencies

- CryptoNote protocol (privacy layer)
- OpenVPN (VPN implementation)
- WireGuard (tunnel transport)
- HAProxy (load balancing)

---

## 9. Changelog

| Version | Date | Changes |
|---------|------|---------|
| 1.0 | 2026-02-01 | Initial RFC specification |

---

## Appendix A: Network Diagram

```
                              ┌─────────────────────┐
                              │    Internet         │
                              │   "Clearnet"        │
                              └──────────┬──────────┘
                                         │
              ┌──────────────────────────┼──────────────────────────┐
              │                          │                          │
              ▼                          ▼                          ▼
     ┌─────────────┐           ┌─────────────┐            ┌─────────────┐
     │  Exit Node  │           │  Exit Node  │            │  Exit Node  │
     │  (Europe)   │           │   (Asia)    │            │  (Americas) │
     └──────┬──────┘           └──────┬──────┘            └──────┬──────┘
            │                         │                          │
            └─────────────────────────┼──────────────────────────┘
                                      │
                              ┌───────┴───────┐
                              │   Lethernet   │
                              │   Network     │
                              └───────┬───────┘
                                      │
                    ┌─────────────────┼─────────────────┐
                    │                 │                 │
                    ▼                 ▼                 ▼
           ┌─────────────┐   ┌─────────────┐   ┌─────────────┐
           │    SDP      │   │  Blockchain │   │   Kevacoin  │
           │  Discovery  │   │   (LTHN)    │   │   Storage   │
           └─────────────┘   └─────────────┘   └─────────────┘
                    │                 │                 │
                    └─────────────────┼─────────────────┘
                                      │
                              ┌───────┴───────┐
                              │    Users      │
                              │  (Clients)    │
                              └───────────────┘
```
