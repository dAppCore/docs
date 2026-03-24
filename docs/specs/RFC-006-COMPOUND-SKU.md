# RFC: Compound SKU Format

**Status:** Implemented
**Created:** 2026-01-15
**Authors:** Host UK Engineering

---

## Abstract

The Compound SKU format encodes product identity, options, quantities, and bundle groupings in a single parseable string. Like [HLCRF](./HLCRF-COMPOSITOR.md) for layouts, it makes complex structure a portable, self-describing data type.

One scan tells you everything. No lookups. No mistakes. One barcode = complete fulfillment knowledge.

---

## Format Specification

```
SKU-<opt>~<val>*<qty>[-<opt>~<val>*<qty>]...
```

| Symbol | Purpose | Example              |
|--------|---------|----------------------|
| `-`    | Option separator | `LAPTOP-ram~16gb`    |
| `~`    | Value indicator | `ram~16gb`           |
| `*`    | Quantity indicator | `cover~black*2`      |
| `,`    | Item separator | `LAPTOP,MOUSE,PAD`   |
| `\|`   | Bundle separator | `LAPTOP\|MOUSE\|PAD` |

---

## Examples

### Single product with options

```
LAPTOP-ram~16gb-ssd~512gb-color~silver
```

### Option with quantity

```
LAPTOP-ram~16gb-cover~black*2
```

Two black covers included.

### Multiple separate items

```
LAPTOP-ram~16gb,HDMI-length~2m,MOUSE-color~black
```

Comma separates distinct line items.

### Bundle (discount lookup)

```
LAPTOP-ram~16gb\|MOUSE-color~black\|PAD-size~xl
```

Pipe binds items for bundle discount detection.

### With entity lineage

```
ORGORG-WBUTS-PROD500-ram~16gb
│      │     │       └── Option
│      │     └────────── Base product SKU
│      └──────────────── M2 entity code
└─────────────────────── M1 entity code
```

The lineage prefix traces through the entity hierarchy.

---

## Bundle Discount Detection

When a compound SKU contains `|` (bundle separator):

```
┌──────────────────────────────────────────────────────────────┐
│  Input: LAPTOP-ram~16gb|MOUSE-color~black|PAD-size~xl       │
│                                                              │
│  Step 1: Detect Bundle (found |)                            │
│                                                              │
│  Step 2: Strip Human Choices                                │
│          → LAPTOP|MOUSE|PAD                                 │
│                                                              │
│  Step 3: Hash the Raw Combination                           │
│          → hash("LAPTOP|MOUSE|PAD") = "abc123..."           │
│                                                              │
│  Step 4: Lookup Bundle Discount                             │
│          → commerce_bundle_hashes["abc123"] = 20% off       │
│                                                              │
│  Step 5: Apply Discount                                     │
│          → Bundle price calculated                          │
└──────────────────────────────────────────────────────────────┘
```

The hash is computed from **sorted base SKUs** (stripping options), so `LAPTOP|MOUSE|PAD` and `PAD|LAPTOP|MOUSE` produce the same hash.

---

## API Reference

### SkuParserService

Parses compound SKU strings into structured data.

```php
use Mod\Commerce\Services\SkuParserService;

$parser = app(SkuParserService::class);

// Parse a compound SKU
$result = $parser->parse('LAPTOP-ram~16gb|MOUSE,HDMI');

// Result contains ParsedItem and BundleItem objects
$result->count();           // 2 (1 bundle + 1 single)
$result->productCount();    // 4 (3 in bundle + 1 single)
$result->hasBundles();      // true
$result->getBundleHashes(); // ['abc123...']
$result->getAllBaseSkus();  // ['LAPTOP', 'MOUSE', 'HDMI']

// Access items
foreach ($result->items as $item) {
    if ($item instanceof BundleItem) {
        echo "Bundle: " . $item->getBaseSkuString();
    } else {
        echo "Item: " . $item->baseSku;
    }
}
```

### SkuBuilderService

Builds compound SKU strings from structured data.

```php
use Mod\Commerce\Services\SkuBuilderService;

$builder = app(SkuBuilderService::class);

// Build from line items
$sku = $builder->build([
    [
        'base_sku' => 'laptop',
        'options' => [
            ['code' => 'ram', 'value' => '16gb'],
            ['code' => 'ssd', 'value' => '512gb'],
        ],
        'bundle_group' => 'cyber',  // Groups into bundle
    ],
    [
        'base_sku' => 'mouse',
        'bundle_group' => 'cyber',
    ],
    [
        'base_sku' => 'hdmi',  // No group = standalone
    ],
]);
// Returns: "LAPTOP-ram~16gb-ssd~512gb|MOUSE,HDMI"

// Add entity lineage
$sku = $builder->addLineage('PROD500', ['ORGORG', 'WBUTS']);
// Returns: "ORGORG-WBUTS-PROD500"

// Generate bundle hash for discount creation
$hash = $builder->generateBundleHash(['LAPTOP', 'MOUSE', 'PAD']);
```

### Data Transfer Objects

```php
use Mod\Commerce\Data\SkuOption;
use Mod\Commerce\Data\ParsedItem;
use Mod\Commerce\Data\BundleItem;
use Mod\Commerce\Data\SkuParseResult;

// Option: code~value*quantity
$option = new SkuOption('ram', '16gb', 1);
$option->toString();  // "ram~16gb"

// Item: baseSku with options
$item = new ParsedItem('LAPTOP', [$option]);
$item->toString();       // "LAPTOP-ram~16gb"
$item->getOption('ram'); // SkuOption
$item->hasOption('ssd'); // false

// Bundle: items grouped for discount
$bundle = new BundleItem($items, $hash);
$bundle->getBaseSkus();      // ['LAPTOP', 'MOUSE']
$bundle->getBaseSkuString(); // "LAPTOP|MOUSE"
$bundle->containsSku('MOUSE'); // true
```

---

## Database Schema

### Bundle Hash Table

```sql
CREATE TABLE commerce_bundle_hashes (
    id BIGINT PRIMARY KEY,
    hash VARCHAR(64) UNIQUE,        -- SHA256 of sorted base SKUs
    base_skus VARCHAR(512),         -- "LAPTOP|MOUSE|PAD" (debugging)

    -- Discount (one of these)
    coupon_code VARCHAR(64),
    fixed_price DECIMAL(12,2),
    discount_percent DECIMAL(5,2),
    discount_amount DECIMAL(12,2),

    entity_id BIGINT,               -- Scope to M1/M2/M3
    valid_from TIMESTAMP,
    valid_until TIMESTAMP,
    active BOOLEAN DEFAULT TRUE
);
```

---

## Connection to HLCRF

Both Compound SKU and HLCRF share the same core innovation: **hierarchy encoded in a parseable string**.

| System | String | Meaning |
|--------|--------|---------|
| HLCRF | `H[LCR]CF` | Layout with nested body in header |
| SKU | `ORGORG-WBUTS-PROD-ram~16gb` | Product with entity lineage and option |

Both eliminate database lookups by making structure self-describing. Parse the string, get the full picture.

---

## Implementation Files

- `app/Mod/Commerce/Services/SkuParserService.php` — Parser
- `app/Mod/Commerce/Services/SkuBuilderService.php` — Builder
- `app/Mod/Commerce/Services/SkuLineageService.php` — Entity lineage tracking
- `app/Mod/Commerce/Data/SkuOption.php` — Option DTO
- `app/Mod/Commerce/Data/ParsedItem.php` — Item DTO
- `app/Mod/Commerce/Data/BundleItem.php` — Bundle DTO
- `app/Mod/Commerce/Data/SkuParseResult.php` — Parse result DTO
- `app/Mod/Commerce/Models/BundleHash.php` — Bundle discount model
- `app/Mod/Commerce/Tests/Feature/CompoundSkuTest.php` — Tests

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0 | 2026-01-15 | Initial RFC |
