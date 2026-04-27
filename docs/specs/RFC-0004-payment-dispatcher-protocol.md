# RFC-0004: Payment & Dispatcher Protocol

```
RFC:            0004
Title:          Payment & Dispatcher Protocol
Status:         Standard
Category:       Standards Track
Authors:        Darbs, Snider
License:        EUPL-1.2
Created:        2026-02-01
Requires:       RFC-0001, RFC-0003
```

---

## Abstract

This document specifies the payment flow between clients and Exit Node providers, including the Dispatcher component that validates payments and authorizes service access. It also describes the Daemon Reward Program that incentivizes blockchain node operators.

---

## 1. Introduction

### 1.1 Purpose

The Lethean payment model enables direct peer-to-peer transactions between users and service providers without intermediaries. The Dispatcher component on each Exit Node validates these payments and grants service access.

### 1.2 Design Goals

1. **Direct Payment** - No payment processor or intermediary
2. **Privacy** - Untraceable transactions on blockchain
3. **Automation** - No manual verification required
4. **Fairness** - Providers paid for actual service delivery

---

## 2. Payment Model

### 2.1 Overview

```
┌──────────┐         ┌─────────────┐         ┌─────────────┐
│  Client  │  LTHN   │  Blockchain │  Verify │  Exit Node  │
│          │────────>│             │────────>│ (Dispatcher)│
└──────────┘         └─────────────┘         └─────────────┘
     │                                              │
     │              Service Access                  │
     │<─────────────────────────────────────────────│
```

### 2.2 Payment Flow

1. **Discovery**: Client queries SDP for available providers
2. **Selection**: User selects provider based on price, location, etc.
3. **Payment**: Client sends LTHN to provider's wallet address
4. **Confirmation**: Transaction confirmed on blockchain
5. **Validation**: Dispatcher detects payment, authorizes client
6. **Connection**: Client connects and uses service

### 2.3 Payment Parameters

| Parameter | Description |
|-----------|-------------|
| cost | LTHN amount per service unit (defined by provider) |
| firstVerificationsRequired | Confirmations needed for first payment (0-2) |
| subsequentVerificationsRequired | Confirmations for later payments (0-1) |

---

## 3. Dispatcher Protocol

### 3.1 Component Architecture

```
                    ┌─────────────────────────────────┐
                    │           DISPATCHER            │
                    │           (lthnvpnd)            │
                    │                                 │
  Client ──────────>│  ┌─────────────────────────┐   │
  Connection        │  │   Payment Validator     │   │
                    │  │   (wallet RPC queries)  │   │
                    │  └───────────┬─────────────┘   │
                    │              │                  │
                    │  ┌───────────▼─────────────┐   │
                    │  │   Session Manager       │   │
                    │  │   (active connections)  │   │
                    │  └───────────┬─────────────┘   │
                    │              │                  │
                    │  ┌───────────▼─────────────┐   │
                    │  │   Service Router        │   │──────> Services
                    │  │   (proxy/vpn dispatch)  │   │
                    │  └─────────────────────────┘   │
                    └─────────────────────────────────┘
```

### 3.2 Wallet RPC Interface

The Dispatcher communicates with the local wallet via JSON-RPC:

**Endpoint**: `http://127.0.0.1:14660/json_rpc`

**Key Methods**:

```json
// Check incoming transfers
{
  "jsonrpc": "2.0",
  "method": "get_transfers",
  "params": {"in": true, "pending": true},
  "id": "1"
}

// Verify specific payment
{
  "jsonrpc": "2.0",
  "method": "get_transfer_by_txid",
  "params": {"txid": "<transaction_hash>"},
  "id": "2"
}
```

### 3.3 Payment Validation Logic

```python
def validate_payment(client_info):
    # Query wallet for incoming transfers
    transfers = wallet_rpc.get_transfers(in=True)

    for transfer in transfers:
        # Check if payment matches expected amount
        if transfer.amount >= service.cost:
            # Check confirmation count
            if transfer.confirmations >= service.firstVerificationsRequired:
                # Check if not already used
                if not is_payment_used(transfer.txid):
                    mark_payment_used(transfer.txid)
                    return authorize_client(client_info, transfer)

    return reject_client(client_info, "Payment not found")
```

### 3.4 Session Management

| Session State | Description |
|---------------|-------------|
| PENDING | Awaiting payment confirmation |
| ACTIVE | Payment validated, service access granted |
| EXPIRED | Session time/data limit exceeded |
| TERMINATED | Manually ended or error |

---

## 4. Transaction Format

### 4.1 Payment Transaction

LTHN uses CryptoNote-based transactions:

| Field | Description |
|-------|-------------|
| inputs | Ring signature inputs (privacy) |
| outputs | Destination addresses (stealth) |
| amount | Payment amount (encrypted) |
| payment_id | Optional identifier (deprecated in favor of subaddresses) |

### 4.2 Confirmation Levels

| Confirmations | Security | Wait Time |
|---------------|----------|-----------|
| 0 | Low (double-spend risk) | Immediate |
| 1 | Medium | ~2 minutes |
| 2 | High | ~4 minutes |

### 4.3 Recommended Settings

| Use Case | First Confirmations | Subsequent |
|----------|---------------------|------------|
| Low-value, trusted | 0 | 0 |
| Standard | 1 | 1 |
| High-value | 2 | 1 |

---

## 5. Pricing Model

### 5.1 Cost Specification

Providers define costs in their SDP configuration:

```json
{
  "services": [{
    "id": "1A",
    "cost": "0.1",
    "unit": "session",
    "duration": 3600
  }]
}
```

### 5.2 Pricing Strategies

| Model | Description |
|-------|-------------|
| Per-session | Fixed cost per connection |
| Per-time | Cost per hour/day |
| Per-data | Cost per GB transferred |
| Subscription | Pre-paid time blocks |

### 5.3 Dynamic Pricing (Future)

Smart contracts may enable:
- Demand-based pricing
- Bulk discounts
- Loyalty rewards

---

## 6. Daemon Reward Program

### 6.1 Purpose

The Daemon Reward Program incentivizes running blockchain nodes (letheand) to ensure network decentralization and stability.

### 6.2 Eligibility Flow

```
┌─────────────────┐
│ Daemon          │
│ Commissioned    │
└────────┬────────┘
         │
         ▼
┌─────────────────┐     No     ┌─────────────────┐
│ Zabbix Agent    │───────────>│   Ineligible    │
│ Running?        │            │   for Reward    │
└────────┬────────┘            └─────────────────┘
         │ Yes
         ▼
┌─────────────────┐     No
│ Daemon          │───────────> (Re-poll)
│ Synchronized?   │
└────────┬────────┘
         │ Yes
         ▼
┌─────────────────┐     No
│ Port            │───────────> (Re-poll)
│ Accessible?     │
└────────┬────────┘
         │ Yes
         ▼
┌─────────────────┐
│ Mark as LIVE    │
│ Start SLA Clock │
└─────────────────┘
```

### 6.3 SLA Calculation

**Formula**:
```
if daemon_sla >= sla_threshold:
    reward = escrow / number_of_eligible_daemons
else:
    reward = 0
```

**Parameters**:

| Parameter | Value | Description |
|-----------|-------|-------------|
| sla_threshold | 98% | Minimum uptime requirement |
| escrow | $250 USD/month | Monthly reward pool |
| cycle_length | 30 days | Evaluation period |
| max_daemons | 50 | Cutoff for reward distribution |

### 6.4 SLA Thresholds

| SLA Level | Allowed Downtime (Monthly) |
|-----------|---------------------------|
| 98% | 14h 36m 34s |
| 99% | 7h 18m 17s |

**Recommendation**: 98% is realistic for independent operators who may not notice overnight outages but can restore service the next day.

### 6.5 Reward Distribution

```
┌─────────────────┐
│ 30 Days Elapsed │
└────────┬────────┘
         │
         ▼
┌─────────────────┐     No     ┌─────────────────┐
│ SLA >= 98%?     │───────────>│ No Reward       │
└────────┬────────┘            │ Notify Operator │
         │ Yes                  │ Reset Clock     │
         ▼                      └─────────────────┘
┌─────────────────┐
│ Calculate:      │
│ reward = escrow │
│   / daemons     │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ Auto-payment    │
│ from Escrow     │
│ to Operator     │
└─────────────────┘
```

### 6.6 Anti-Gaming Measures

To prevent abuse:

1. **Wallet Analysis** - Compare addresses across daemons
2. **WHOIS Comparison** - Detect same provider/location
3. **Geographic Distribution** - Encourage spread across regions
4. **Hotspot Prevention** - Auto-exclude clustered daemons

---

## 7. Security Considerations

### 7.1 Double-Spend Protection

- Require confirmations before service access
- Monitor mempool for conflicting transactions
- Rate-limit rapid reconnection attempts

### 7.2 Replay Protection

- Track used payment transaction IDs
- Reject previously-used payments
- Session binding to payment hash

### 7.3 Payment Privacy

- Ring signatures hide sender
- Stealth addresses hide receiver
- Amount encryption

---

## 8. Implementation Notes

### 8.1 Wallet RPC Configuration

```ini
; dispatcher.ini
wallet-rpc-uri=http://127.0.0.1:14660/json_rpc
wallet-username=dispatcher
wallet-password=<secure-random-string>
```

### 8.2 Starting Wallet RPC

```bash
lethean-wallet-vpn-rpc \
  --wallet-file /opt/lthn/wallet \
  --rpc-bind-port 14660 \
  --rpc-login dispatcher:<password> \
  --daemon-address 127.0.0.1:48782
```

### 8.3 Monitoring Payments

```python
# Example: Watch for incoming payments
import requests

def check_payments():
    response = requests.post(
        "http://127.0.0.1:14660/json_rpc",
        json={
            "jsonrpc": "2.0",
            "method": "get_transfers",
            "params": {"in": True, "pending": True},
            "id": "1"
        },
        auth=("dispatcher", password)
    )
    return response.json()["result"]["in"]
```

---

## 9. References

- RFC-0001: Lethean Network Overview
- RFC-0003: Exit Node Architecture
- CryptoNote Protocol: https://cryptonote.org/
- Lethean Daemon Reward Workflow Documentation

---

## Appendix A: Reward Calculation Example

```
Given:
  - escrow = $250 USD
  - eligible_daemons = 25
  - daemon_sla = 99.2%
  - threshold = 98%

Calculation:
  daemon_sla (99.2%) >= threshold (98%) ✓

  reward = escrow / eligible_daemons
  reward = $250 / 25
  reward = $10 USD per daemon

Result:
  Operator receives $10 USD equivalent in LTHN
```

---

## Appendix B: Payment State Machine

```
                    ┌──────────────┐
                    │   INITIAL    │
                    └──────┬───────┘
                           │ Client connects
                           ▼
                    ┌──────────────┐
                    │   PENDING    │◄─────────────┐
                    └──────┬───────┘              │
                           │                      │
            ┌──────────────┼──────────────┐       │
            │              │              │       │
            ▼              ▼              ▼       │
     ┌───────────┐  ┌───────────┐  ┌───────────┐ │
     │  PAYMENT  │  │  TIMEOUT  │  │  REJECTED │ │
     │  FOUND    │  │           │  │           │ │
     └─────┬─────┘  └───────────┘  └───────────┘ │
           │                                      │
           │ Confirmations met                    │
           ▼                                      │
     ┌───────────┐                               │
     │  ACTIVE   │───────────────────────────────┘
     │ (session) │      Session expired/renewed
     └─────┬─────┘
           │ Disconnect/Limit reached
           ▼
     ┌───────────┐
     │TERMINATED │
     └───────────┘
```
