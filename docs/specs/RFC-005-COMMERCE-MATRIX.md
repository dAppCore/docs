# RFC: Commerce Entity Matrix

**Status:** Implemented
**Created:** 2026-01-15
**Authors:** Host UK Engineering

---

## Abstract

The Commerce Entity Matrix is a hierarchical permission and content system for multi-channel commerce. It enables master companies (M1) to control product catalogues, storefronts (M2) to select and white-label products, and dropshippers (M3) to inherit complete stores with zero management overhead.

The core innovation is **top-down immutable permissions**: if a parent says "NO", every descendant is locked to "NO". Children can only restrict further, never expand. Combined with sparse content overrides and a self-learning training mode, the system provides complete audit trails and deterministic behaviour.

Like [HLCRF](./HLCRF-COMPOSITOR.md) for layouts and [Compound SKU](./COMPOUND-SKU.md) for product identity, the Matrix eliminates complexity through composable primitives rather than configuration sprawl.

---

## Motivation

Traditional multi-tenant commerce systems copy data between entities, leading to synchronisation nightmares, inconsistent pricing, and broken audit trails. When Original Organics ran four websites, telephone orders, mail orders, and garden centre voucher schemes in 2008, they needed a system where:

1. **M1 owns truth** — Products exist in one place; everything else references them
2. **M2 selects and customises** — Storefronts choose products and can override presentation
3. **M3 inherits completely** — Dropshippers get fully functional stores without management burden
4. **Permissions cascade down** — A restriction at the top is immutable below
5. **Every action is gated** — No default-allow; if it wasn't trained, it doesn't work

The Matrix addresses this through hierarchical entities, sparse overrides, and request-level permission enforcement.

---

## Terminology

### Entity Types

| Code | Type | Role |
|------|------|------|
| **M1** | Master Company | Source of truth. Owns the product catalogue, sets base pricing, controls what's possible. |
| **M2** | Facade/Storefront | Selects from M1's catalogue. Can override content, adjust pricing within bounds, operate independent sales channels. |
| **M3** | Dropshipper | Full inheritance with zero management. Sees everything, reports everything, manages nothing. Can create their own M2s. |

### Entity Hierarchy

```
M1 - Master Company (Source of Truth)
│
├── Master Product Catalogue
│   └── Products live here, nowhere else
│
├── M2 - Storefronts (Select from M1)
│   ├── waterbutts.com
│   ├── originalorganics.co.uk
│   ├── telephone-orders (internal)
│   └── garden-vouchers (B2B)
│
└── M3 - Dropshippers (Full Inheritance)
    ├── External company selling our products
    └── Can have their own M2s
        ├── dropshipper.com
        └── dropshipper-wholesale.com
```

### Materialised Path

Each entity stores its position in the hierarchy as a path string:

| Entity | Path | Depth |
|--------|------|-------|
| ORGORG (M1) | `ORGORG` | 0 |
| WBUTS (M2) | `ORGORG/WBUTS` | 1 |
| DRPSHP (M3) | `ORGORG/WBUTS/DRPSHP` | 2 |

The path enables ancestor lookups without recursive queries.

---

## Permission Matrix

### The Core Rules

```
If M1 says "NO" → Everything below is "NO"
If M1 says "YES" → M2 can say "NO" for itself
If M2 says "YES" → M3 can say "NO" for itself

Permissions cascade DOWN. Restrictions are IMMUTABLE from above.
```

### Visual Model

```
                    M1 (Master)
                    ├── can_sell_alcohol: NO ──────────────┐
                    ├── can_discount: YES                  │
                    └── can_export: YES                    │
                         │                                 │
            ┌────────────┼────────────┐                    │
            ▼            ▼            ▼                    │
         M2-Web      M2-Phone     M2-Voucher              │
         ├── can_sell_alcohol: [LOCKED NO] ◄──────────────┘
         ├── can_discount: NO (restricted self)
         └── can_export: YES (inherited)
              │
              ▼
           M3-Dropshipper
           ├── can_sell_alcohol: [LOCKED NO] (from M1)
           ├── can_discount: [LOCKED NO] (from M2)
           └── can_export: YES (can restrict to NO)
```

### The Three Dimensions

```
Dimension 1: Entity Hierarchy (M1 → M2 → M3)
Dimension 2: Permission Keys (can_sell, can_discount, can_view_cost...)
Dimension 3: Resource Scope (products, orders, customers, reports...)

Permission = Matrix[Entity][Key][Scope]
```

### Permission Entry Schema

```sql
CREATE TABLE permission_matrix (
    id BIGINT PRIMARY KEY,
    entity_id BIGINT NOT NULL,

    -- What permission
    key VARCHAR(128),              -- product.create, order.refund
    scope VARCHAR(128),            -- Resource type or specific ID

    -- The value
    allowed BOOLEAN DEFAULT FALSE,
    locked BOOLEAN DEFAULT FALSE,  -- Set by parent, cannot override

    -- Audit
    source VARCHAR(32),            -- inherited, explicit, trained
    set_by_entity_id BIGINT,       -- Who locked it
    trained_at TIMESTAMP,          -- When it was learned
    trained_route VARCHAR(255),    -- Which route triggered training

    UNIQUE (entity_id, key, scope)
);
```

### Source Types

| Source | Meaning |
|--------|---------|
| `inherited` | Cascaded from parent entity's lock |
| `explicit` | Manually set by administrator |
| `trained` | Learned through training mode |

---

## Permission Cascade Algorithm

When checking if an entity can perform an action:

```
1. Build hierarchy path (root M1 → parent M2 → current entity)
2. For each ancestor, top-down:
   - Find permission for (entity, key, scope)
   - If locked AND denied → RETURN DENIED (immutable)
   - If denied (not locked) → RETURN DENIED
3. Check entity's own permission:
   - If exists → RETURN allowed/denied
4. Permission undefined → handle based on mode
```

### Lock Cascade

When an entity locks a permission, all descendants receive an inherited lock:

```php
public function lock(Entity $entity, string $key, bool $allowed): void
{
    // Set on this entity
    PermissionMatrix::updateOrCreate(
        ['entity_id' => $entity->id, 'key' => $key],
        ['allowed' => $allowed, 'locked' => true, 'source' => 'explicit']
    );

    // Cascade to all descendants
    $descendants = Entity::where('path', 'like', $entity->path . '/%')->get();

    foreach ($descendants as $descendant) {
        PermissionMatrix::updateOrCreate(
            ['entity_id' => $descendant->id, 'key' => $key],
            [
                'allowed' => $allowed,
                'locked' => true,
                'source' => 'inherited',
                'set_by_entity_id' => $entity->id,
            ]
        );
    }
}
```

---

## Training Mode

### The Problem

Building a complete permission matrix upfront is impractical. You don't know every action until you build the system.

### The Solution

Training mode learns permissions by observing real usage:

```
1. Developer navigates to /admin/products
2. Clicks "Create Product"
3. System: "BLOCKED - No permission defined for:"
   - Entity: M1-Admin
   - Action: product.create
   - Route: POST /admin/products

4. Developer clicks [Allow for M1-Admin]
5. Permission recorded in matrix with source='trained'
6. Continue working

Result: Complete map of every action in the system
```

### Configuration

```php
// config/commerce.php
'matrix' => [
    // Training mode - undefined permissions prompt for approval
    'training_mode' => env('COMMERCE_MATRIX_TRAINING', false),

    // Production mode - undefined = denied
    'strict_mode' => env('COMMERCE_MATRIX_STRICT', true),

    // Log all permission checks (for audit)
    'log_all_checks' => env('COMMERCE_MATRIX_LOG_ALL', false),

    // Log denied requests
    'log_denials' => true,

    // Default action when permission undefined (only if strict=false)
    'default_allow' => false,
],
```

### Permission Request Logging

```sql
CREATE TABLE permission_requests (
    id BIGINT PRIMARY KEY,
    entity_id BIGINT NOT NULL,

    -- Request details
    method VARCHAR(10),            -- GET, POST, PUT, DELETE
    route VARCHAR(255),            -- /admin/products
    action VARCHAR(128),           -- product.create
    scope VARCHAR(128),

    -- Context
    request_data JSON,             -- Sanitised request params
    user_agent VARCHAR(255),
    ip_address VARCHAR(45),

    -- Result
    status VARCHAR(32),            -- allowed, denied, pending
    was_trained BOOLEAN DEFAULT FALSE,
    trained_at TIMESTAMP,

    created_at TIMESTAMP
);
```

### Production Mode

```
If permission not in matrix → 403 Forbidden
No exceptions. No fallbacks. No "default allow".

If it wasn't trained, it doesn't exist.
```

---

## Product Assignment

### How Products Flow Through the Hierarchy

M1 owns the master catalogue. M2/M3 entities don't copy products; they create **assignments** that reference the master and optionally override specific fields.

```sql
CREATE TABLE commerce_product_assignments (
    id BIGINT PRIMARY KEY,
    entity_id BIGINT NOT NULL,     -- M2 or M3
    product_id BIGINT NOT NULL,    -- Reference to master

    -- SKU customisation
    sku_suffix VARCHAR(64),        -- Custom suffix for this entity

    -- Price overrides (if allowed by matrix)
    price_override INT,            -- Override base price
    price_tier_overrides JSON,     -- Override tier pricing
    margin_percent DECIMAL(5,2),   -- Percentage margin
    fixed_margin INT,              -- Fixed margin amount

    -- Content overrides
    name_override VARCHAR(255),
    description_override TEXT,
    image_override VARCHAR(512),

    -- Control
    is_active BOOLEAN DEFAULT TRUE,
    is_featured BOOLEAN DEFAULT FALSE,
    sort_order INT DEFAULT 0,
    allocated_stock INT,           -- Entity-specific allocation
    can_discount BOOLEAN DEFAULT TRUE,
    min_price INT,                 -- Floor price
    max_price INT,                 -- Ceiling price

    UNIQUE (entity_id, product_id)
);
```

### Effective Values

The assignment provides effective value getters that fall back to the master product:

```php
public function getEffectivePrice(): int
{
    return $this->price_override ?? $this->product->price;
}

public function getEffectiveName(): string
{
    return $this->name_override ?? $this->product->name;
}
```

### SKU Lineage

Full SKUs encode the entity path:

```
ORGORG-WBUTS-WB500L    # Original Organics → Waterbutts → 500L Water Butt
ORGORG-PHONE-WB500L    # Same product, telephone channel
DRPSHP-THEIR1-WB500L   # Dropshipper's storefront selling our product
```

This tracks:
- Where the sale originated
- Which facade/channel
- Back to master SKU

---

## Content Overrides

### The Core Insight

**Don't copy data. Create sparse overrides. Resolve at runtime.**

```
M1 (Master) has content
    │
    │ (M2 sees M1's content by default)
    ▼
M2 customises product name
    │
    │ Override entry: (M2, product:123, name, "Custom Name")
    │ Everything else still inherits from M1
    ▼
M3 (Dropshipper) inherits M2's view
    │
    │ (Sees M2's custom name, M1's everything else)
    ▼
M3 customises description
    │
    │ Override entry: (M3, product:123, description, "Their description")
    │ Still has M2's name, M1's other fields
    ▼
Resolution: M3 sees merged content from all levels
```

### Override Table Schema

```sql
CREATE TABLE commerce_content_overrides (
    id BIGINT PRIMARY KEY,
    entity_id BIGINT NOT NULL,

    -- What's being overridden (polymorphic)
    overrideable_type VARCHAR(128),  -- Product, Category, Page, etc.
    overrideable_id BIGINT,
    field VARCHAR(64),               -- name, description, image, price

    -- The override value
    value TEXT,
    value_type VARCHAR(32),          -- string, json, html, decimal, boolean

    -- Audit
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    UNIQUE (entity_id, overrideable_type, overrideable_id, field)
);
```

### Value Types

| Type | Storage | Use Case |
|------|---------|----------|
| `string` | Raw text | Names, short descriptions |
| `json` | JSON-encoded | Structured data, arrays |
| `html` | Raw HTML | Rich content |
| `integer` | String → int | Counts, quantities |
| `decimal` | String → float | Prices, percentages |
| `boolean` | `1`/`0` | Flags, toggles |

### Resolution Algorithm

```
Query: "What is product 123's name for M3-ACME?"

Step 1: Check M3-ACME overrides
        → NULL (no override)

Step 2: Check M2-WATERBUTTS overrides (parent)
        → "Premium 500L Water Butt" ✓

Step 3: Return "Premium 500L Water Butt"
        (M3-ACME sees M2's override, not M1's original)
```

If M3-ACME later customises the name, their override takes precedence for themselves and their descendants.

---

## API Reference

### PermissionMatrixService

The service handles all permission checks and training.

```php
use Mod\Commerce\Services\PermissionMatrixService;

$matrix = app(PermissionMatrixService::class);

// Check permission
$result = $matrix->can($entity, 'product.create', $scope);

if ($result->isAllowed()) {
    // Proceed
} elseif ($result->isDenied()) {
    // Handle denial: $result->reason
} elseif ($result->isUndefined()) {
    // No permission defined
}

// Gate a request (handles training mode)
$result = $matrix->gateRequest($request, $entity, 'order.refund');

// Set permission explicitly
$matrix->setPermission($entity, 'product.create', true);

// Lock permission (cascades to descendants)
$matrix->lock($entity, 'product.view_cost', false);

// Unlock (removes inherited locks)
$matrix->unlock($entity, 'product.view_cost');

// Train permission (dev mode)
$matrix->train($entity, 'product.create', $scope, true, $route);
```

### PermissionResult

```php
use Mod\Commerce\Services\PermissionResult;

// Factory methods
PermissionResult::allowed();
PermissionResult::denied(reason: 'Locked by M1', lockedBy: $entity);
PermissionResult::undefined(key: 'action', scope: 'resource');
PermissionResult::pending(key: 'action', trainingUrl: '/train/...');

// Status checks
$result->isAllowed();
$result->isDenied();
$result->isUndefined();
$result->isPending();
```

### Entity Model

```php
use Mod\Commerce\Models\Entity;

// Create master
$m1 = Entity::createMaster('ORGORG', 'Original Organics');

// Create facade under master
$m2 = $m1->createFacade('WBUTS', 'Waterbutts.com', [
    'domain' => 'waterbutts.com',
    'currency' => 'GBP',
]);

// Create dropshipper under facade
$m3 = $m2->createDropshipper('ACME', 'ACME Supplies');

// Hierarchy helpers
$m3->getAncestors();    // [M1, M2]
$m3->getHierarchy();    // [M1, M2, M3]
$m3->getRoot();         // M1
$m3->getDescendants();  // Children, grandchildren, etc.

// Type checks
$entity->isMaster();      // or isM1()
$entity->isFacade();      // or isM2()
$entity->isDropshipper(); // or isM3()

// SKU building
$entity->buildSku('WB500L'); // "ORGORG-WBUTS-WB500L"
```

---

## Standard Permission Keys

```php
// Product permissions
'product.list'              // View product list
'product.view'              // View product detail
'product.view_cost'         // See cost price (M1 only usually)
'product.create'            // Create new product (M1 only)
'product.update'            // Update product
'product.delete'            // Delete product
'product.price_override'    // Override price on facade

// Order permissions
'order.list'                // View orders
'order.view'                // View order detail
'order.create'              // Create order
'order.update'              // Update order
'order.cancel'              // Cancel order
'order.refund'              // Process refund
'order.export'              // Export order data

// Customer permissions
'customer.list'
'customer.view'
'customer.view_email'       // See customer email
'customer.view_phone'       // See customer phone
'customer.export'           // Export customer data (GDPR)

// Report permissions
'report.sales'              // Sales reports
'report.revenue'            // Revenue (might hide from M3)
'report.cost'               // Cost reports (M1 only)
'report.margin'             // Margin reports (M1 only)

// System permissions
'settings.view'
'settings.update'
'entity.create'             // Create child entities
'entity.manage'             // Manage entity settings
```

---

## Middleware Integration

### CommerceMatrixGate

```php
// app/Http/Middleware/CommerceMatrixGate.php

public function handle(Request $request, Closure $next)
{
    $entity = $this->resolveEntity($request);
    $action = $this->resolveAction($request);

    if (!$entity || !$action) {
        return $next($request); // Not a commerce route
    }

    $result = $this->matrix->gateRequest($request, $entity, $action);

    if ($result->isDenied()) {
        return response()->json([
            'error' => 'permission_denied',
            'message' => $result->reason,
        ], 403);
    }

    if ($result->isPending()) {
        // Training mode - show prompt
        return response()->view('commerce.matrix.train-prompt', [
            'result' => $result,
            'entity' => $entity,
        ], 428); // Precondition Required
    }

    return $next($request);
}
```

### Route Definition

```php
// Explicit action mapping
Route::post('/products', [ProductController::class, 'store'])
    ->matrixAction('product.create');

Route::post('/orders/{order}/refund', [OrderController::class, 'refund'])
    ->matrixAction('order.refund');
```

---

## Order Flow Through the Matrix

```
Customer places order on waterbutts.com (M2)
    │
    ▼
┌─────────────────────────────────────────┐
│ Order Created                            │
│ - entity_id: M2-WBUTS                   │
│ - sku: ORGORG-WBUTS-WB500L              │
│ - customer sees: M2 branding            │
└────────────────┬────────────────────────┘
                 │
                 ▼
┌─────────────────────────────────────────┐
│ M1 Fulfillment Queue                     │
│ - M1 sees all orders from all M2s       │
│ - Can filter by facade                  │
│ - Ships with M2 branding (or neutral)   │
└────────────────┬────────────────────────┘
                 │
                 ▼
┌─────────────────────────────────────────┐
│ Reporting                                │
│ - M1: Sees all, costs, margins          │
│ - M2: Sees own orders, no cost data     │
│ - M3: Sees own orders, wholesale price  │
└─────────────────────────────────────────┘
```

---

## Pricing

Pricing is not a separate system. It emerges from:

1. **Permission Matrix** — `can_discount`, `max_discount_percent`, `can_sell_below_wholesale`
2. **Product Assignments** — `price_override`, `min_price`, `max_price`, `margin_percent`
3. **Content Overrides** — Sparse price adjustments per entity
4. **SKU System** — Bundle hashes, option modifiers, volume rules

No separate pricing engine needed. Primitives compose.

---

## Implementation Files

### Models

- `app/Mod/Commerce/Models/Entity.php` — Entity hierarchy
- `app/Mod/Commerce/Models/PermissionMatrix.php` — Permission entries
- `app/Mod/Commerce/Models/PermissionRequest.php` — Request logging
- `app/Mod/Commerce/Models/ContentOverride.php` — Sparse overrides
- `app/Mod/Commerce/Models/ProductAssignment.php` — M2/M3 product links

### Services

- `app/Mod/Commerce/Services/PermissionMatrixService.php` — Permission logic
- `app/Mod/Commerce/Services/ContentOverrideService.php` — Override resolution

### Configuration

- `app/Mod/Commerce/config.php` — Matrix configuration

---

## Related RFCs

- [HLCRF Compositor](./HLCRF-COMPOSITOR.md) — Same philosophy applied to layouts
- [Compound SKU](./COMPOUND-SKU.md) — Same philosophy applied to product identity

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0 | 2026-01-15 | Initial RFC |
