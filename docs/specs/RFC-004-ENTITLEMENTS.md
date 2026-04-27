# RFC: Entitlements and Feature System

**Status:** Implemented
**Created:** 2026-01-15
**Authors:** Host UK Engineering

---

## Abstract

The Entitlement System controls feature access, usage limits, and tier gating across all Host services. It answers one question: "Can this workspace do this action?"

Workspaces subscribe to **Packages** that bundle **Features**. Features are either boolean flags (access gates) or numeric limits (usage caps). **Boosts** provide temporary or permanent additions to base limits. Usage is tracked, cached, and enforced in real-time.

The system integrates with Commerce for subscription lifecycle and exposes an API for cross-service entitlement checks.

---

## Core Model

### Entity Relationships

```
┌──────────────────────────────────────────────────────────────────┐
│                                                                  │
│   Workspace ───┬─── WorkspacePackage ─── Package ─── Features    │
│                │                                                 │
│                ├─── Boosts (temporary limit additions)           │
│                │                                                 │
│                ├─── UsageRecords (consumption tracking)          │
│                │                                                 │
│                └─── EntitlementLogs (audit trail)                │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘
```

### Workspace

The tenant unit. All entitlement checks happen against a workspace, not a user. Users belong to workspaces; workspaces own entitlements.

```php
// Check if workspace can use a feature
$workspace->can('social.accounts', quantity: 3);

// Record usage
$workspace->recordUsage('ai.credits', quantity: 10);

// Get usage summary
$workspace->getUsageSummary();
```

### Package

A bundle of features with defined limits. Two types:

| Type | Behaviour |
|------|-----------|
| **Base Package** | Only one active per workspace. Upgrading replaces the previous base package. |
| **Add-on Package** | Stackable. Multiple can be active simultaneously. Limits accumulate. |

**Database:** `entitlement_packages`

```php
// Package fields
'code'              // Unique identifier (e.g., 'social-creator')
'name'              // Display name
'is_base_package'   // true = only one allowed
'is_stackable'      // true = limits add to base
'monthly_price'     // Pricing
'yearly_price'
'stripe_price_id_monthly'
'stripe_price_id_yearly'
```

### Feature

A capability or limit that can be granted. Three types:

| Type | Behaviour | Example |
|------|-----------|---------|
| **Boolean** | On/off access gate | `tier.apollo`, `host.social` |
| **Limit** | Numeric cap on usage | `social.accounts` (5), `ai.credits` (100) |
| **Unlimited** | No cap (special limit value) | Agency tier social posts |

**Database:** `entitlement_features`

```php
// Feature fields
'code'              // Unique identifier (e.g., 'social.accounts')
'name'              // Display name
'type'              // boolean, limit, unlimited
'reset_type'        // none, monthly, rolling
'rolling_window_days' // For rolling reset (e.g., 30)
'parent_feature_id' // For global pools (see Storage Pool below)
```

#### Reset Types

| Reset Type | Behaviour |
|------------|-----------|
| **None** | Usage accumulates forever (e.g., account limits) |
| **Monthly** | Resets at billing cycle start |
| **Rolling** | Rolling window (e.g., last 30 days) |

#### Hierarchical Features (Global Pools)

Child features share a parent's limit pool. Used for storage allocation across services:

```
host.storage.total (1000 MB)
├── host.cdn (draws from parent pool)
├── bio.cdn (draws from parent pool)
└── social.cdn (draws from parent pool)
```

### WorkspacePackage

The pivot linking workspaces to packages. Tracks subscription state.

**Database:** `entitlement_workspace_packages`

```php
// Status constants
STATUS_ACTIVE     // Package in effect
STATUS_SUSPENDED  // Temporarily disabled (e.g., payment failure)
STATUS_CANCELLED  // Marked for removal
STATUS_EXPIRED    // Past expiry date

// Key fields
'starts_at'             // When package becomes active
'expires_at'            // When package ends
'billing_cycle_anchor'  // For monthly reset calculations
'blesta_service_id'     // External billing system reference
```

### Boost

Temporary or permanent additions to feature limits. Use cases:
- One-time credit top-ups
- Promotional extras
- Cycle-bound bonuses that expire at billing renewal

**Database:** `entitlement_boosts`

```php
// Boost types
BOOST_TYPE_ADD_LIMIT  // Add to existing limit
BOOST_TYPE_ENABLE     // Enable a boolean feature
BOOST_TYPE_UNLIMITED  // Grant unlimited access

// Duration types
DURATION_CYCLE_BOUND  // Expires at billing cycle end
DURATION_DURATION     // Expires after set time
DURATION_PERMANENT    // Never expires

// Key fields
'limit_value'        // Amount to add
'consumed_quantity'  // How much has been used
'status'             // active, exhausted, expired, cancelled
```

---

## How Checking Works

### The `can()` Method

All access checks flow through `EntitlementService::can()`.

```php
public function can(Workspace $workspace, string $featureCode, int $quantity = 1): EntitlementResult
```

**Algorithm:**

```
1. Look up feature by code
2. If feature has parent, use parent's code for pool lookup
3. Sum limits from all active packages + boosts
4. If any source grants unlimited → return allowed (unlimited)
5. Get current usage (respecting reset type)
6. If usage + quantity > limit → deny
7. Otherwise → allow
```

**Example:**

```php
// Check before creating social account
$result = $workspace->can('social.accounts');

if ($result->isDenied()) {
    throw new EntitlementException($result->getMessage());
}

// Proceed with creation...

// Record the usage
$workspace->recordUsage('social.accounts');
```

### EntitlementResult

The return value from `can()`. Provides all context needed for UI feedback.

```php
$result = $workspace->can('ai.credits', quantity: 10);

$result->isAllowed();        // bool
$result->isDenied();         // bool
$result->isUnlimited();      // bool
$result->getMessage();       // Denial reason

$result->limit;              // Total limit (100)
$result->used;               // Current usage (75)
$result->remaining;          // Remaining (25)
$result->getUsagePercentage(); // 75.0
$result->isNearLimit();      // true if > 80%
```

### Caching

Limits and usage are cached for 5 minutes to avoid repeated database queries.

```php
// Cache keys
"entitlement:{workspace_id}:limit:{feature_code}"
"entitlement:{workspace_id}:usage:{feature_code}"
```

Cache is invalidated when:
- Package is provisioned, suspended, cancelled, or reactivated
- Boost is provisioned or expires
- Usage is recorded

---

## Usage Tracking

### Recording Usage

After a gated action succeeds, record the consumption:

```php
$workspace->recordUsage(
    featureCode: 'ai.credits',
    quantity: 10,
    user: $user,           // Optional: who triggered it
    metadata: [            // Optional: context
        'model' => 'claude-3',
        'tokens' => 1500,
    ]
);
```

**Database:** `entitlement_usage_records`

### Usage Calculation

Usage is calculated based on the feature's reset type:

| Reset Type | Query |
|------------|-------|
| None | All records ever |
| Monthly | Records since billing cycle start |
| Rolling | Records in last N days |

```php
// Monthly: Get current cycle start from primary package
$cycleStart = $workspace->workspacePackages()
    ->whereHas('package', fn($q) => $q->where('is_base_package', true))
    ->first()
    ->getCurrentCycleStart();

UsageRecord::getTotalUsage($workspaceId, $featureCode, $cycleStart);

// Rolling: Last 30 days
UsageRecord::getRollingUsage($workspaceId, $featureCode, days: 30);
```

### Usage Summary

For dashboards, get all features with their current state:

```php
$summary = $workspace->getUsageSummary();

// Returns Collection grouped by category:
[
    'social' => [
        ['code' => 'social.accounts', 'limit' => 5, 'used' => 3, ...],
        ['code' => 'social.posts.scheduled', 'limit' => 100, 'used' => 45, ...],
    ],
    'ai' => [
        ['code' => 'ai.credits', 'limit' => 100, 'used' => 75, ...],
    ],
]
```

---

## Integration Points

### Commerce Integration

Subscriptions from Commerce automatically provision/revoke entitlement packages.

**Event Flow:**

```
SubscriptionCreated → ProvisionSocialHostSubscription listener
                      → EntitlementService::provisionPackage()

SubscriptionCancelled → Revoke package (immediate or at period end)

SubscriptionRenewed → Update expires_at
                      → Expire cycle-bound boosts
                      → Reset monthly usage (via cycle anchor)
```

**Plan Changes:**

```php
$subscriptionService->changePlan(
    $subscription,
    $newPackage,
    prorate: true,      // Calculate credit/charge
    immediate: true     // Apply now vs. period end
);
```

### External Billing (Blesta)

The API supports external billing systems via webhook-style endpoints:

```
POST   /api/v1/entitlements       → Provision package
POST   /api/v1/entitlements/{id}/suspend
POST   /api/v1/entitlements/{id}/unsuspend
POST   /api/v1/entitlements/{id}/cancel
POST   /api/v1/entitlements/{id}/renew
GET    /api/v1/entitlements/{id}  → Get status
```

### Cross-Service API

External services (BioHost, etc.) check entitlements via API:

```
GET  /api/v1/entitlements/check
     ?email=user@example.com
     &feature=bio.pages
     &quantity=1

POST /api/v1/entitlements/usage
     { email, feature, quantity, metadata }

GET  /api/v1/entitlements/summary
GET  /api/v1/entitlements/summary/{workspace}
```

---

## Feature Categories

Features are organised by category for display grouping:

| Category | Features |
|----------|----------|
| **tier** | `tier.apollo`, `tier.hades`, `tier.nyx`, `tier.stygian` |
| **service** | `host.social`, `host.bio`, `host.analytics`, `host.trust` |
| **social** | `social.accounts`, `social.posts.scheduled`, `social.workspaces` |
| **ai** | `ai.credits`, `ai.providers.claude`, `ai.providers.gemini` |
| **biolink** | `bio.pages`, `bio.shortlinks`, `bio.domains` |
| **analytics** | `analytics.sites`, `analytics.pageviews` |
| **storage** | `host.storage.total`, `host.cdn`, `bio.cdn`, `social.cdn` |
| **team** | `team.members` |
| **api** | `api.requests` |
| **support** | `support.mailboxes`, `support.agents`, `support.conversations` |
| **tools** | `tool.url_shortener`, `tool.qr_generator`, `tool.dns_lookup` |

---

## Audit Logging

All entitlement changes are logged for compliance and debugging.

**Database:** `entitlement_logs`

```php
// Log actions
ACTION_PACKAGE_PROVISIONED
ACTION_PACKAGE_SUSPENDED
ACTION_PACKAGE_CANCELLED
ACTION_PACKAGE_REACTIVATED
ACTION_PACKAGE_RENEWED
ACTION_PACKAGE_EXPIRED
ACTION_BOOST_PROVISIONED
ACTION_BOOST_CONSUMED
ACTION_BOOST_EXHAUSTED
ACTION_BOOST_EXPIRED
ACTION_BOOST_CANCELLED
ACTION_USAGE_RECORDED
ACTION_USAGE_DENIED

// Log sources
SOURCE_BLESTA      // External billing
SOURCE_COMMERCE    // Internal commerce
SOURCE_ADMIN       // Manual admin action
SOURCE_SYSTEM      // Automated (e.g., expiry)
SOURCE_API         // API call
```

---

## Implementation Files

### Models
- `app/Mod/Tenant/Models/Feature.php`
- `app/Mod/Tenant/Models/Package.php`
- `app/Mod/Tenant/Models/WorkspacePackage.php`
- `app/Mod/Tenant/Models/Boost.php`
- `app/Mod/Tenant/Models/UsageRecord.php`
- `app/Mod/Tenant/Models/EntitlementLog.php`

### Services
- `app/Mod/Tenant/Services/EntitlementService.php` - Core logic
- `app/Mod/Tenant/Services/EntitlementResult.php` - Result DTO

### API
- `app/Mod/Api/Controllers/EntitlementApiController.php`

### Commerce Integration
- `app/Mod/Commerce/Listeners/ProvisionSocialHostSubscription.php`
- `app/Mod/Commerce/Services/SubscriptionService.php`

### Database
- `entitlement_features` - Feature definitions
- `entitlement_packages` - Package definitions
- `entitlement_package_features` - Package/feature pivot with limits
- `entitlement_workspace_packages` - Workspace subscriptions
- `entitlement_boosts` - Temporary additions
- `entitlement_usage_records` - Consumption tracking
- `entitlement_logs` - Audit trail

### Seeders
- `app/Mod/Tenant/Database/Seeders/FeatureSeeder.php`

### Tests
- `app/Mod/Tenant/Tests/Feature/EntitlementServiceTest.php`
- `app/Mod/Tenant/Tests/Feature/EntitlementApiTest.php`

---

## Usage Examples

### Basic Access Check

```php
// In controller or service
$result = $workspace->can('social.accounts');

if ($result->isDenied()) {
    return back()->with('error', $result->getMessage());
}

// Perform action...
$workspace->recordUsage('social.accounts');
```

### With Quantity

```php
// Before bulk import
$result = $workspace->can('social.posts.scheduled', quantity: 50);

if ($result->isDenied()) {
    return "Cannot schedule {$quantity} posts. " .
           "Remaining: {$result->remaining}";
}
```

### Tier Check

```php
// Gate premium features
if ($workspace->isApollo()) {
    // Show Apollo-tier features
}

// Or directly
if ($workspace->can('tier.apollo')->isAllowed()) {
    // ...
}
```

### Usage Dashboard Data

```php
// For billing/usage page
$summary = $workspace->getUsageSummary();
$packages = $entitlements->getActivePackages($workspace);
$boosts = $entitlements->getActiveBoosts($workspace);
```

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0 | 2026-01-15 | Initial RFC |
