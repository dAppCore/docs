# RFC-0003: Exit Node Architecture

```
RFC:            0003
Title:          Exit Node Architecture
Status:         Standard
Category:       Standards Track
Authors:        Darbs, Snider
License:        EUPL-1.2
Created:        2026-02-01
Requires:       RFC-0001, RFC-0002
```

---

## Abstract

This document specifies the architecture and operation of Lethean Exit Nodes—the infrastructure components that provide VPN and proxy services to network users. It covers deployment, configuration, service components, and operational requirements.

---

## 1. Introduction

### 1.1 Purpose

Exit Nodes are the service delivery points of the Lethean network. They receive user traffic, validate payments, and route connections to the internet. This RFC standardizes Exit Node implementation to ensure interoperability and consistent user experience.

### 1.2 Scope

This document covers:
- Exit Node component architecture
- Deployment and configuration
- Service types and protocols
- Operational requirements

---

## 2. Architecture Overview

### 2.1 Component Stack

```
┌─────────────────────────────────────────────────────────────┐
│                    EXTERNAL INTERFACE                       │
│                  (Internet-facing ports)                    │
├──────────────────────┬──────────────────────────────────────┤
│     HAProxy          │          TCP 8880 (HTTP)             │
│   (Load Balancer)    │          TCP 8881 (HTTPS/TLS)        │
├──────────────────────┼──────────────────────────────────────┤
│                      │                                      │
│   ┌──────────────────┴───────────────────────────┐         │
│   │              SERVICE LAYER                    │         │
│   │  ┌─────────────┐  ┌─────────────┐            │         │
│   │  │  TinyProxy  │  │   OpenVPN   │            │         │
│   │  │  TCP 8888   │  │  UDP 18080  │            │         │
│   │  └─────────────┘  └─────────────┘            │         │
│   └──────────────────────────────────────────────┘         │
│                                                             │
├─────────────────────────────────────────────────────────────┤
│                    DISPATCHER (lthnvpnd)                    │
│              Payment validation & routing                   │
├──────────────────────┬──────────────────────────────────────┤
│                      │                                      │
│   ┌──────────────────┴───────────────────────────┐         │
│   │              BLOCKCHAIN LAYER                 │         │
│   │  ┌─────────────┐  ┌─────────────┐            │         │
│   │  │  letheand   │  │   Wallet    │            │         │
│   │  │TCP 48782/72 │  │ TCP 1444/45 │            │         │
│   │  └─────────────┘  └─────────────┘            │         │
│   └──────────────────────────────────────────────┘         │
│                                                             │
├─────────────────────────────────────────────────────────────┤
│                    MANAGEMENT LAYER                         │
│     lvmgmt, lvpnc-client-man (TCP 8124), Nginx             │
└─────────────────────────────────────────────────────────────┘
```

### 2.2 Component Summary

| Component | Binary | Ports | Purpose |
|-----------|--------|-------|---------|
| HAProxy | haproxy | 8880, 8881 | Load balancing, TLS termination |
| TinyProxy | tinyproxy | 8888 | HTTP/HTTPS proxy service |
| OpenVPN | openvpn | 18080/UDP | VPN tunnel service |
| Dispatcher | lthnvpnd | - | Payment validation, routing |
| Daemon | letheand | 48782, 48772 | Blockchain node |
| Wallet | lethean-wallet-vpn-rpc | 1444, 1445 | Payment handling |
| Client Manager | lvpnc-client-man | 8124 | Client session management |
| Management | lvmgmt | - | Configuration, SDP generation |

---

## 3. Deployment

### 3.1 Docker Deployment (Recommended)

Exit Nodes are deployed as Docker containers for consistency:

```bash
docker run -d \
  --name lethean-node \
  -p 8880:8880 \
  -p 8881:8881 \
  -p 8888:8888 \
  -p 18080:18080/udp \
  -v /opt/lthn:/opt/lthn \
  letheanvpn/lvpn:latest
```

### 3.2 Startup Sequence

When the container starts in **node mode**:

1. **WireGuard Connection**
   - Connects to Lethean space via WireGuard
   - Endpoint: `vpn2.lethean.space:8774`

2. **Blockchain Sync**
   - letheand starts and syncs blockchain
   - Connects to configured seed peers
   - Private address: `172.31.129.19` (via tunnel)

3. **Easy-Deploy**
   - Automated infrastructure setup
   - Certificate generation
   - Service configuration

4. **Wallet RPC**
   - lethean-wallet-vpn-rpc starts
   - Listens on `127.0.0.1:14660`

5. **Proxy Services**
   - TinyProxy starts on port 8888
   - OpenVPN starts on port 18080 (if enabled)

6. **VDP Registration**
   - Pushes VDP to VDP Manager
   - URL: `https://mgr.lethean.space`

7. **VDP Sync Loop**
   - Fetches provider list from VDP Manager
   - Refreshes every 3500 seconds (TTL: 3600s)

---

## 4. Configuration

### 4.1 dispatcher.ini

The primary configuration file at `/opt/lthn/etc/dispatcher.ini`:

#### 4.1.1 Global Section

```ini
[global]
; Debug level: DEBUG, INFO, WARN, ERROR
debug=INFO

; CA certificate for provider identity
ca=/opt/lthn/etc/ca/certs/ca.cert.pem

; Provider identification
provider-id=efaa812b358956f93a0e324385c8b44469a99e5a82f2de327297b25d8c2ee288
provider-key=<secret-key>
provider-name=MyExitNode
provider-type=commercial

; Terms of service
provider-terms=https://example.com/terms
; Or from file: provider-terms=@/opt/lthn/etc/terms.txt
```

#### 4.1.2 Wallet Section

```ini
; Wallet configuration
wallet-address=iz5HSgJUW0max8Hs2TEAacKhKA9LXLLDvc4u7yCV7Lm4iwkgFXTMFBAdtj2mqMpqy7T4BNveDQdW8LBPVxWqy94B2A6sKJXQ7
wallet-rpc-uri=http://127.0.0.1:14660/json_rpc
wallet-username=dispatcher
wallet-password=<secure-password>
```

#### 4.1.3 Proxy Service (1A-1Z)

```ini
[service-1A]
name=AU_Melbourne_Proxy1
backend_proxy_server=localhost:3128
crt=/opt/lthn/etc/ca/certs/ha.cert.pem
key=/opt/lthn/etc/ca/private/ha.key.pem
crtkey=/opt/lthn/etc/ca/certs/ha.both.pem
```

#### 4.1.4 VPN Service (2A-2Z)

```ini
[service-1B]
crt=/opt/lthn/etc/ca/certs/openvpn.cert.pem
key=/opt/lthn/etc/ca/private/openvpn.key.pem
crtkey=/opt/lthn/etc/ca/certs/openvpn.both.pem
reneg=600
enabled=true
iprange=10.8.0.0
ipmask=255.255.255.0
ip6range=fd00::/64
dns=1.1.1.1
mgmtport=11123
```

### 4.2 Configuration Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| provider-id | hex(64) | Yes | Unique provider identifier |
| provider-key | string | Yes | Authentication secret |
| provider-name | string(16) | Yes | Display name |
| wallet-address | string | Yes | LTHN payment address |
| wallet-rpc-uri | URL | Yes | Wallet RPC endpoint |
| backend_proxy_server | host:port | Per-service | Proxy backend |
| crt | path | Per-service | Service certificate |
| key | path | Per-service | Private key |
| reneg | integer | VPN only | Renegotiation interval (seconds) |
| iprange | CIDR | VPN only | Client IP pool |

---

## 5. Service Types

### 5.1 HTTP Proxy (Type: proxy)

**Implementation**: TinyProxy or Squid

**Ports**:
- External: 8080 (via HAProxy 8880)
- Internal: 8888, 3128

**Features**:
- HTTP/HTTPS forwarding
- Optional authentication
- Access logging

**Use Cases**:
- Browser-based privacy
- Application proxy configuration
- Lightweight traffic routing

### 5.2 VPN Tunnel (Type: vpn)

**Implementation**: OpenVPN

**Ports**:
- UDP 18080 (primary)
- TCP 443 (fallback, optional)

**Features**:
- Full tunnel encryption
- All traffic routing
- DNS leak protection

**Configuration Options**:
```ini
iprange=10.8.0.0          ; Client IP pool start
ipmask=255.255.255.0      ; Subnet mask
ip6range=fd00::/64        ; IPv6 pool
dns=1.1.1.1               ; DNS server for clients
reneg=600                 ; Rekey interval
```

### 5.3 Future Service Types

The architecture supports additional service types:

- **SOCKS5 Proxy** (planned)
- **WireGuard VPN** (planned)
- **Custom BYOA** (Bring Your Own Application)

---

## 6. Security

### 6.1 Certificate Architecture

```
Root CA (provider)
    │
    ├── HAProxy Certificate (TLS termination)
    │
    ├── Proxy Service Certificate
    │
    └── VPN Service Certificate
```

### 6.2 Certificate Generation

```bash
# Generate CA
openssl genrsa -out ca.key 4096
openssl req -new -x509 -days 3650 -key ca.key -out ca.cert.pem

# Generate service certificate
openssl genrsa -out service.key 2048
openssl req -new -key service.key -out service.csr
openssl x509 -req -in service.csr -CA ca.cert.pem -CAkey ca.key \
  -CAcreateserial -out service.cert.pem -days 365
```

### 6.3 Access Control

The Dispatcher validates:

1. **Payment Verification** - Check blockchain for valid transaction
2. **Session Management** - Track active sessions
3. **Rate Limiting** - Prevent abuse
4. **Geographic Restrictions** - Optional country blocking

---

## 7. Monitoring

### 7.1 Health Checks

Exit Nodes should implement health endpoints:

| Endpoint | Purpose |
|----------|---------|
| /health | Basic liveness check |
| /status | Detailed component status |
| /metrics | Prometheus-compatible metrics |

### 7.2 Zabbix Integration

For daemon reward eligibility (see RFC-0004):

```
Zabbix Agent → letheand port check
            → Blockchain height sync
            → Service availability
```

### 7.3 Logging

Recommended log locations:

```
/var/log/lthn/dispatcher.log
/var/log/lthn/haproxy.log
/var/log/lthn/tinyproxy.log
/var/log/lthn/openvpn.log
```

---

## 8. Network Requirements

### 8.1 Bandwidth

| Tier | Minimum | Recommended |
|------|---------|-------------|
| Basic | 10 Mbps | 100 Mbps |
| Standard | 100 Mbps | 1 Gbps |
| Premium | 1 Gbps | 10 Gbps |

### 8.2 Ports

| Port | Protocol | Direction | Purpose |
|------|----------|-----------|---------|
| 8880 | TCP | Inbound | Client HTTP |
| 8881 | TCP | Inbound | Client HTTPS |
| 8888 | TCP | Internal | Proxy service |
| 18080 | UDP | Inbound | VPN service |
| 48782 | TCP | Outbound | Daemon P2P |
| 48772 | TCP | Outbound | Daemon RPC |

### 8.3 Firewall Rules

```bash
# Allow client connections
iptables -A INPUT -p tcp --dport 8880 -j ACCEPT
iptables -A INPUT -p tcp --dport 8881 -j ACCEPT
iptables -A INPUT -p udp --dport 18080 -j ACCEPT

# Allow blockchain P2P
iptables -A OUTPUT -p tcp --dport 48782 -j ACCEPT
```

---

## 9. Operational Procedures

### 9.1 Initial Setup

```bash
# 1. Install Docker
curl -fsSL https://get.docker.com | sh

# 2. Create directories
mkdir -p /opt/lthn/etc/ca/{certs,private}

# 3. Generate provider identity
lvmgmt --generate-provider-id

# 4. Generate SDP configuration
lvmgmt --generate-sdp --wallet-address <YOUR_WALLET>

# 5. Start node
docker-compose up -d
```

### 9.2 Updating

```bash
# Pull latest image
docker pull letheanvpn/lvpn:latest

# Restart with new image
docker-compose down
docker-compose up -d
```

### 9.3 Backup

Critical files to backup:
- `/opt/lthn/etc/dispatcher.ini`
- `/opt/lthn/etc/ca/` (certificates)
- `/opt/lthn/etc/sdp.json`
- Wallet files

---

## 10. References

- RFC-0001: Lethean Network Overview
- RFC-0002: Service Descriptor Protocol (SDP)
- RFC-0004: Payment & Dispatcher Protocol
- OpenVPN Documentation: https://openvpn.net/
- HAProxy Documentation: https://www.haproxy.org/

---

## Appendix A: Docker Compose Example

```yaml
version: '3.8'

services:
  lethean-node:
    image: letheanvpn/lvpn:latest
    container_name: lethean-exit-node
    restart: unless-stopped
    ports:
      - "8880:8880"
      - "8881:8881"
      - "18080:18080/udp"
    volumes:
      - ./config:/opt/lthn/etc
      - ./data:/opt/lthn/data
    environment:
      - LETHEAN_MODE=node
      - WALLET_ADDRESS=${WALLET_ADDRESS}
    cap_add:
      - NET_ADMIN
    sysctls:
      - net.ipv4.ip_forward=1
```

---

## Appendix B: Component Ports Reference

```
┌────────────────────────────────────────────────────────────────┐
│                        EXIT NODE                               │
│                                                                │
│  EXTERNAL                    INTERNAL                          │
│  ────────                    ────────                          │
│  8880 ──────► HAProxy                                          │
│  8881 ──────►    │                                             │
│                  ├────────► 8888 (tinyproxy)                   │
│                  ├────────► 3128 (squid, optional)             │
│                  └────────► 8124 (lvpnc-client-man)            │
│                                                                │
│  18080/UDP ──────────────► OpenVPN                             │
│                                                                │
│                             letheand ──────► 48782, 48772      │
│                             wallet ────────► 1444, 1445        │
│                             wallet-rpc ────► 14660             │
│                                                                │
└────────────────────────────────────────────────────────────────┘
```
