# RFC-0005: Client Protocol

```
RFC:            0005
Title:          Client Protocol
Status:         Standard
Category:       Standards Track
Authors:        Darbs, Snider
License:        EUPL-1.2
Created:        2026-02-01
Requires:       RFC-0001, RFC-0002, RFC-0004
```

---

## Abstract

This document specifies the Lethean VPN Client (lvpnc) protocol, including service discovery, connection establishment, payment submission, and session management. It covers both GUI and CLI interfaces as well as transport layer options.

---

## 1. Introduction

### 1.1 Purpose

The Lethean Client enables users to discover, connect to, and pay for privacy services provided by Exit Nodes. This specification defines the client-side protocols and components required for seamless integration with the Lethean network.

### 1.2 Design Principles

1. **User Sovereignty** - User controls their wallet and connection choices
2. **Privacy First** - No tracking, minimal metadata exposure
3. **Cross-Platform** - Support for Windows, Linux, macOS
4. **Flexibility** - GUI and CLI interfaces for different user preferences

---

## 2. Client Architecture

### 2.1 Component Overview

```
┌────────────────────────────────────────────────────────────────┐
│                     LETHEAN CLIENT (lvpnc)                      │
│                                                                 │
│  ┌───────────────────────────────────────────────────────────┐ │
│  │                         GUI                               │ │
│  │                    (Kivy-based)                           │ │
│  └─────────────────────────┬─────────────────────────────────┘ │
│                            │                                   │
│  ┌─────────────────────────┴─────────────────────────────────┐ │
│  │                    Core Services                          │ │
│  │                                                           │ │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐       │ │
│  │  │   Provider  │  │   Session   │  │   Payment   │       │ │
│  │  │   Manager   │  │   Manager   │  │   Manager   │       │ │
│  │  └──────┬──────┘  └──────┬──────┘  └──────┬──────┘       │ │
│  └─────────┼────────────────┼────────────────┼───────────────┘ │
│            │                │                │                 │
│  ┌─────────┴────────────────┴────────────────┴───────────────┐ │
│  │                  Transport Layer                          │ │
│  │                                                           │ │
│  │  ┌───────────┐  ┌───────────┐  ┌───────────┐            │ │
│  │  │    SSH    │  │ WireGuard │  │   HTTP    │            │ │
│  │  │   Tunnel  │  │   Tunnel  │  │   Proxy   │            │ │
│  │  └───────────┘  └───────────┘  └───────────┘            │ │
│  └───────────────────────────────────────────────────────────┘ │
│                                                                 │
│  ┌───────────────────────────────────────────────────────────┐ │
│  │                  Blockchain Services                      │ │
│  │                                                           │ │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐       │ │
│  │  │   Wallet    │  │   Daemon    │  │   Daemon    │       │ │
│  │  │    RPC      │  │    RPC      │  │   (local)   │       │ │
│  │  └─────────────┘  └─────────────┘  └─────────────┘       │ │
│  └───────────────────────────────────────────────────────────┘ │
└────────────────────────────────────────────────────────────────┘
```

### 2.2 Core Components

| Component | Description |
|-----------|-------------|
| **GUI** | Kivy-based graphical interface |
| **Provider Manager** | Handles SDP queries and VDP caching |
| **Session Manager** | Manages active connections |
| **Payment Manager** | Wallet integration and transaction submission |
| **Transport Layer** | SSH, WireGuard, or HTTP proxy tunnels |
| **Wallet RPC** | Local wallet for payment processing |
| **Daemon RPC** | Connection to blockchain node |

---

## 3. Service Discovery

### 3.1 Discovery Flow

```
┌──────────┐          ┌──────────────┐          ┌──────────┐
│  Client  │  Query   │  SDP Server  │  Cached  │   VDP    │
│          │─────────>│              │─────────>│  Store   │
└──────────┘          └──────────────┘          └──────────┘
     │                       │
     │      Provider List    │
     │<──────────────────────│
     │                       │
     ▼                       │
┌──────────┐                 │
│  User    │                 │
│Selection │                 │
└──────────┘                 │
```

### 3.2 SDP Query

```python
# Query available providers
GET https://sdp.lethean.io/v1/services/search

# Optional filters
?type=proxy        # Service type (proxy, vpn)
?country=AU        # Country code (ISO 3166-1)
?max_cost=0.5      # Maximum cost in LTHN
?min_speed=10mbps  # Minimum speed
```

### 3.3 Provider Selection Criteria

| Criterion | Description |
|-----------|-------------|
| **Location** | Geographic proximity for latency |
| **Cost** | LTHN price per session/hour/GB |
| **Speed** | Advertised bandwidth |
| **Type** | Service type (proxy, VPN) |
| **Reputation** | Future: on-chain reputation |

### 3.4 VDP Caching

```
Directory Structure:
├── providers/        # All known provider VDPs
├── my-providers/     # Providers we've connected to
├── spaces/           # Lethernet space definitions
├── gates/            # Gateway definitions
└── sessions/         # Active session data
```

---

## 4. Connection Protocol

### 4.1 Connection Flow

```
┌────────────┐                                    ┌────────────┐
│   Client   │                                    │ Exit Node  │
└─────┬──────┘                                    └─────┬──────┘
      │                                                 │
      │  1. Query SDP for providers                     │
      │────────────────────────────────────────────────>│
      │                                                 │
      │  2. Select provider, parse VDP                  │
      │<────────────────────────────────────────────────│
      │                                                 │
      │  3. Send payment to provider wallet             │
      │────────────────(Blockchain)────────────────────>│
      │                                                 │
      │  4. Wait for confirmation (0-2 blocks)          │
      │                                                 │
      │  5. Connect via transport (SSH/WG/HTTP)         │
      │────────────────────────────────────────────────>│
      │                                                 │
      │  6. Dispatcher validates payment                │
      │                                                 │
      │  7. Session established                         │
      │<────────────────────────────────────────────────│
      │                                                 │
      │  8. Route traffic through tunnel                │
      │◄───────────────────────────────────────────────►│
```

### 4.2 Transport Options

| Transport | Port | Protocol | Use Case |
|-----------|------|----------|----------|
| **SSH Tunnel** | 8880 | TCP | Default, firewall-friendly |
| **WireGuard** | 8774 | UDP | High performance VPN |
| **HTTP Proxy** | 8080 | TCP | Browser-only proxy |
| **HTTPS/TLS** | 8881 | TCP | Manager-over-TLS |

### 4.3 Exit Node Connection

```
Client connects to Exit Node:

HTTP/8880  ──► haproxy ──► Authenticated proxy access
HTTPS/8881 ──► haproxy ──► Manager-over-TLS channel
UDP/8774   ──► WireGuard ──► VPN tunnel
```

---

## 5. Payment Integration

### 5.1 Payment Flow

```
┌──────────────┐         ┌──────────────┐         ┌──────────────┐
│    Client    │         │  Blockchain  │         │   Provider   │
│    Wallet    │         │   Network    │         │    Wallet    │
└──────┬───────┘         └──────┬───────┘         └──────┬───────┘
       │                        │                        │
       │  1. Create TX          │                        │
       │─────────────────────►  │                        │
       │                        │                        │
       │  2. Broadcast          │                        │
       │                        │─────────────────────►  │
       │                        │                        │
       │  3. Confirm (0-2 blocks)                       │
       │                        │                        │
       │  4. Connect with TX proof                      │
       │───────────────────────────────────────────────►│
       │                        │                        │
       │  5. Dispatcher validates                       │
       │                        │  ◄──────────────────── │
       │                        │                        │
       │  6. Session granted                            │
       │◄───────────────────────────────────────────────│
```

### 5.2 Payment Parameters

| Parameter | Description |
|-----------|-------------|
| `--default-pay-days` | Days to pay for by default |
| `--auto-pay-days` | Auto-pay without GUI confirmation |
| `--free-session-days` | Days to request for free services |
| `--unpaid-expiry` | Seconds before unpaid session expires |
| `--use-tx-pool` | Accept unconfirmed payments |

### 5.3 Wallet Configuration

```ini
# client.ini wallet settings
wallet-rpc-url=http://127.0.0.1:14660/json_rpc
wallet-rpc-port=14660
wallet-rpc-user=client
wallet-rpc-password=<auto-generated>
wallet-name=default
wallet-address=<LTHN_ADDRESS>
```

---

## 6. Session Management

### 6.1 Session States

| State | Description |
|-------|-------------|
| **UNPAID** | Session created, awaiting payment |
| **PENDING** | Payment sent, awaiting confirmation |
| **ACTIVE** | Payment confirmed, connection established |
| **EXPIRED** | Time/data limit reached |
| **DISCONNECTED** | User-initiated disconnect |

### 6.2 Session Lifecycle

```
                    ┌───────────┐
                    │  SELECT   │
                    │  PROVIDER │
                    └─────┬─────┘
                          │
                          ▼
                    ┌───────────┐
                    │  UNPAID   │
                    └─────┬─────┘
                          │ Send payment
                          ▼
                    ┌───────────┐
                    │  PENDING  │◄─────────────────┐
                    └─────┬─────┘                  │
                          │ Confirmed              │ Renew
                          ▼                        │
                    ┌───────────┐                  │
                    │  ACTIVE   │──────────────────┘
                    └─────┬─────┘
                          │ Expire/Disconnect
                          ▼
              ┌───────────┴───────────┐
              ▼                       ▼
        ┌───────────┐           ┌───────────┐
        │  EXPIRED  │           │DISCONNECTED│
        └───────────┘           └───────────┘
```

### 6.3 Auto-Reconnect

```ini
# Reconnection settings
auto-reconnect=30        # Seconds between reconnect attempts
auto-connect=lthn://provider-id/service-id
```

---

## 7. Command Line Interface

### 7.1 Installation

**Windows (PowerShell):**
```powershell
irm https://raw.githubusercontent.com/letheanVPN/lvpn/main/install-all-in-one.ps1 | iex
```

**Linux/macOS:**
```bash
curl -sSL https://lethean.io/install.sh | bash
```

### 7.2 Basic Usage

```bash
# Run with GUI
python client.py --run-gui=1

# Run headless (CLI only)
python client.py --run-gui=0

# Auto-connect to specific provider
python client.py --auto-connect=lthn://provider-id/1A
```

### 7.3 Key Command Options

| Option | Description |
|--------|-------------|
| `--run-gui={0,1}` | Enable/disable GUI |
| `--run-proxy={0,1}` | Run local proxy server |
| `--run-wallet={0,1}` | Run embedded wallet |
| `--run-daemon={0,1}` | Run embedded daemon |
| `--log-level={DEBUG,INFO,WARNING,ERROR}` | Logging verbosity |
| `--config=<path>` | Config file location |

### 7.4 Directory Options

| Option | Description |
|--------|-------------|
| `--var-dir` | Variable data directory |
| `--cfg-dir` | Configuration directory |
| `--app-dir` | Application directory |
| `--tmp-dir` | Temporary files |
| `--providers-dir` | Provider VDP cache |
| `--sessions-dir` | Session data storage |

---

## 8. WireGuard Integration

### 8.1 WireGuard Configuration

```ini
# WireGuard-specific settings
enable-wg=1
wg-map-device=gate1,wg0
wg-map-privkey=gate1,<base64-private-key>
wg-cmd-prefix=sudo
wg-shutdown-on-disconnect=1
```

### 8.2 WireGuard Commands

| Option | Description |
|--------|-------------|
| `--wg-cmd-create-interface` | Command to create WG interface |
| `--wg-cmd-delete-interface` | Command to delete WG interface |
| `--wg-cmd-set-ip` | Command to assign IP to interface |
| `--wg-cmd-set-interface-up` | Command to bring interface up |
| `--wg-cmd-route` | Command to add routes |

### 8.3 WireGuard Connection Flow

```
Client                                 Exit Node
  │                                       │
  │  1. Generate WG keypair               │
  │                                       │
  │  2. Request WG config from provider   │
  │──────────────────────────────────────>│
  │                                       │
  │  3. Receive peer config               │
  │<──────────────────────────────────────│
  │                                       │
  │  4. Create WG interface               │
  │  5. Configure peer                    │
  │  6. Establish tunnel                  │
  │══════════════════════════════════════>│
  │                                       │
  │  7. Route traffic through WG          │
  │◄═════════════════════════════════════►│
```

---

## 9. GUI Interface

### 9.1 Main Views

| View | Purpose |
|------|---------|
| **Provider List** | Browse and select Exit Nodes |
| **Connection** | Active connection status |
| **Wallet** | Balance and transaction history |
| **Settings** | Configuration options |

### 9.2 Connection Workflow

1. **Browse** - View available providers from SDP
2. **Select** - Choose provider based on criteria
3. **Pay** - Send LTHN to provider wallet
4. **Connect** - Establish tunnel after payment confirms
5. **Use** - Traffic routes through Exit Node

---

## 10. Security Considerations

### 10.1 Transport Security

- SSH tunnels use provider's CA for authentication
- WireGuard provides modern cryptographic security
- TLS 1.3 for HTTPS connections

### 10.2 Payment Security

- Wallet RPC uses authentication
- Transactions are cryptographically signed
- Ring signatures provide sender privacy

### 10.3 Session Security

- Session binding to payment transaction
- No session sharing across clients
- Automatic session cleanup on disconnect

---

## 11. Configuration File

### 11.1 Default Locations

| Platform | Path |
|----------|------|
| Windows | `%USERPROFILE%\lvpn\client.ini` |
| Linux | `/etc/lvpn/client.ini` or `~/.config/lvpn/client.ini` |
| macOS | `~/Library/Application Support/lvpn/client.ini` |

### 11.2 Example Configuration

```ini
[client]
log-level=INFO
run-gui=1
run-proxy=1
run-wallet=1

[wallet]
wallet-rpc-url=http://127.0.0.1:14660/json_rpc
wallet-name=default

[connection]
auto-reconnect=30
default-pay-days=1

[wireguard]
enable-wg=0
wg-shutdown-on-disconnect=1
```

---

## 12. Lethernet Access

### 12.1 Beyond Standard VPN

Connected clients gain access to Lethernet services:

| Service Type | Description |
|--------------|-------------|
| **Web Servers** | Privately hosted web content |
| **File Sharing** | Distributed storage solutions |
| **Name Services** | `.lthn` domain namespace |
| **Social Networks** | Decentralized communication |
| **Custom BYOA** | Bring Your Own Application |

### 12.2 Service Access

```
Client ──► Exit Node ──► Internet (standard VPN)
                    └──► Lethernet Services (private network)
```

---

## 13. References

- RFC-0001: Lethean Network Overview
- RFC-0002: Service Descriptor Protocol (SDP)
- RFC-0003: Exit Node Architecture
- RFC-0004: Payment & Dispatcher Protocol
- WireGuard Protocol: https://www.wireguard.com/protocol/

---

## Appendix A: Full Command Options

```
usage: client.py [-h] [-c CONFIG] [-l {DEBUG,INFO,WARNING,ERROR}]
                 [--log-file LOG_FILE] [--audit-file AUDIT_FILE]
                 [--http-port HTTP_PORT] [--var-dir VAR_DIR]
                 [--cfg-dir CFG_DIR] [--app-dir APP_DIR]
                 [--tmp-dir TMP_DIR] [--daemon-host DAEMON_HOST]
                 [--daemon-bin DAEMON_BIN] [--daemon-rpc-url DAEMON_RPC_URL]
                 [--daemon-p2p-port DAEMON_P2P_PORT]
                 [--wallet-rpc-bin WALLET_RPC_BIN]
                 [--wallet-cli-bin WALLET_CLI_BIN]
                 [--wallet-rpc-url WALLET_RPC_URL]
                 [--wallet-rpc-port WALLET_RPC_PORT]
                 [--wallet-rpc-user WALLET_RPC_USER]
                 [--wallet-rpc-password WALLET_RPC_PASSWORD]
                 [--wallet-address WALLET_ADDRESS]
                 [--spaces-dir SPACES_DIR] [--gates-dir GATES_DIR]
                 [--providers-dir PROVIDERS_DIR]
                 [--sessions-dir SESSIONS_DIR]
                 [--coin-type {lethean}] [--coin-unit COIN_UNIT]
                 [--lthn-price LTHN_PRICE]
                 [--default-pay-days DEFAULT_PAY_DAYS]
                 [--unpaid-expiry UNPAID_EXPIRY]
                 [--use-tx-pool USE_TX_POOL]
                 [--enable-wg {0,1}]
                 [--run-gui {0,1}] [--run-proxy {0,1}]
                 [--run-wallet {0,1}] [--run-daemon {0,1}]
                 [--auto-connect AUTO_CONNECT]
                 [--auto-reconnect AUTO_RECONNECT]
                 [--auto-pay-days AUTO_PAY_DAYS]
                 [--free-session-days FREE_SESSION_DAYS]
```

---

## Appendix B: Connection State Machine

```
                         ┌─────────────────────┐
                         │       IDLE          │
                         └──────────┬──────────┘
                                    │ User selects provider
                                    ▼
                         ┌─────────────────────┐
                         │   DISCOVERING       │
                         │  (fetch VDP)        │
                         └──────────┬──────────┘
                                    │ VDP received
                                    ▼
                         ┌─────────────────────┐
                         │    UNPAID           │◄────────────┐
                         └──────────┬──────────┘             │
                                    │ Payment sent           │
                                    ▼                        │
                         ┌─────────────────────┐             │
                         │    CONFIRMING       │             │
                         │  (await blocks)     │             │
                         └──────────┬──────────┘             │
                                    │ Confirmed              │
                                    ▼                        │
                         ┌─────────────────────┐             │
                         │   CONNECTING        │             │
                         │  (transport setup)  │             │
                         └──────────┬──────────┘             │
                                    │ Connected              │
                                    ▼                        │
                         ┌─────────────────────┐             │
                         │     ACTIVE          │─────────────┘
                         │   (tunneled)        │    Renew
                         └──────────┬──────────┘
                                    │ Disconnect/Expire
                                    ▼
                         ┌─────────────────────┐
                         │   DISCONNECTED      │
                         └─────────────────────┘
```
