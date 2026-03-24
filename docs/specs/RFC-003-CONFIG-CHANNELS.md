# RFC: Config Channels

**Status:** Implemented
**Created:** 2026-01-15
**Authors:** Host UK Engineering

---

## Abstract

Config Channels add a voice/context dimension to configuration resolution. Where scopes (workspace, org, system) determine *who* a setting applies to, channels determine *where* or *how* it applies.

A workspace might have one Twitter handle but different posting styles for different contexts. Channels let you define `social.posting.style = "casual"` for Instagram while keeping `social.posting.style = "professional"` for LinkedIn—same workspace, same key, different channel.

The system resolves values through a two-dimensional matrix: scope chain (workspace → org → system) crossed with channel chain (specific → parent → null). Most specific wins, unless a parent declares FINAL.

---

## Motivation

Traditional configuration systems work on a single dimension: scope hierarchy. You set a value at system level, override it at workspace level. Simple.

But some configuration varies by context within a single workspace:

- **Technical channels:** web vs API vs mobile (different rate limits, caching, auth)
- **Social channels:** Instagram vs Twitter vs TikTok (different post lengths, hashtags, tone)
- **Voice channels:** formal vs casual vs support (different language, greeting styles)

Without channels, you either:
1. Create separate config keys for each context (`twitter.style`, `instagram.style`, etc.)
2. Store JSON blobs and parse them at runtime
3. Build custom logic for each use case

Channels generalise this pattern. One key, multiple channel-specific values, clean resolution.

---

## Core Concepts

### Channel

A named context for configuration. Channels have:

| Property | Purpose |
|----------|---------|
| `code` | Unique identifier (e.g., `instagram`, `api`, `support`) |
| `name` | Human-readable label |
| `parent_id` | Optional parent for inheritance |
| `workspace_id` | Owner workspace (null = system channel) |
| `metadata` | Arbitrary JSON for channel-specific data |

### Channel Inheritance

Channels form inheritance trees. A specific channel inherits from its parent:

```
                ┌─────────┐
                │  null   │  ← All channels (fallback)
                └────┬────┘
                     │
                ┌────┴────┐
                │ social  │  ← Social media defaults
                └────┬────┘
           ┌─────────┼─────────┐
           │         │         │
      ┌────┴────┐ ┌──┴───┐ ┌───┴───┐
      │instagram│ │twitter│ │tiktok │
      └─────────┘ └───────┘ └───────┘
```

When resolving `social.posting.style` for the `instagram` channel:
1. Check instagram-specific value
2. Check social (parent) value
3. Check null (all channels) value

### System vs Workspace Channels

**System channels** (`workspace_id = null`) are available to all workspaces. Platform-level contexts like `web`, `api`, `mobile`.

**Workspace channels** are private to a workspace. Custom contexts like `vip_support`, `internal_comms`, or workspace-specific social accounts.

When looking up a channel by code, workspace channels take precedence over system channels with the same code. This allows workspaces to override system channel behaviour.

### Resolution Matrix

Config resolution operates on a matrix of scope × channel:

```
                        ┌──────────────────────────────────────────┐
                        │              Channel Chain               │
                        │  instagram → social → null               │
                        └──────────────────────────────────────────┘
┌───────────────────┐   ┌──────────┬──────────┬──────────┐
│                   │   │          │          │          │
│  Scope Chain      │   │ instagram│  social  │   null   │
│                   │   │          │          │          │
├───────────────────┼───┼──────────┼──────────┼──────────┤
│ workspace         │   │    1     │    2     │    3     │
├───────────────────┼───┼──────────┼──────────┼──────────┤
│ org               │   │    4     │    5     │    6     │
├───────────────────┼───┼──────────┼──────────┼──────────┤
│ system            │   │    7     │    8     │    9     │
└───────────────────┴───┴──────────┴──────────┴──────────┘

Resolution order: 1 → 2 → 3 → 4 → 5 → 6 → 7 → 8 → 9
(Most specific scope + most specific channel first)
```

The first non-null value wins—unless a less-specific combination has `locked = true` (FINAL), which blocks all more-specific values.

### FINAL (Locked Values)

A value marked as `locked` cannot be overridden by more specific scopes or channels. This implements the FINAL pattern from Java/OOP:

```php
// System admin sets rate limit and locks it
$config->set('api.rate_limit', 1000, $systemProfile, locked: true, channel: 'api');

// Workspace cannot override - locked value always wins
$config->set('api.rate_limit', 5000, $workspaceProfile, channel: 'api');
// ↑ This value exists but is never returned
```

Lock checks traverse from least specific (system + null channel) to most specific. Any lock encountered blocks all more-specific values.

---

## How It Works

### Read Path

```
┌─────────────────────────────────────────────────────────────────┐
│  $config->get('social.posting.style')                           │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│  1. Hash lookup (O(1))                                          │
│     ConfigResolver::$values['social.posting.style']             │
│     → Found? Return immediately                                 │
└─────────────────────────────────────────────────────────────────┘
                              │ Miss
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│  2. Lazy load scope (1 query)                                   │
│     Load all resolved values for workspace+channel into hash    │
│     → Check hash again                                          │
└─────────────────────────────────────────────────────────────────┘
                              │ Still miss
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│  3. Lazy prime (N queries)                                      │
│     Build profile chain (workspace → org → system)              │
│     Build channel chain (specific → parent → null)              │
│     Batch load all values for key                               │
│     Walk resolution matrix until value found                    │
│     Store in hash + database                                    │
└─────────────────────────────────────────────────────────────────┘
```

Most reads hit step 1 (hash lookup). The heavy resolution only runs once per key per scope+channel combination, then gets cached.

### Write Path

```php
$config->set(
    keyCode: 'social.posting.style',
    value: 'casual',
    profile: $workspaceProfile,
    locked: false,
    channel: 'instagram',
);
```

1. Update `config_values` (source of truth)
2. Clear affected entries from hash and `config_resolved`
3. Re-prime the key for affected scope+channel
4. Fire `ConfigChanged` event

### Prime Operation

The prime operation pre-computes resolved values:

```php
// Prime entire workspace
$config->prime($workspace, channel: 'instagram');

// Prime all workspaces (scheduled job)
$config->primeAll();
```

This runs full matrix resolution for every key and stores results in `config_resolved`. Subsequent reads become single indexed lookups.

---

## API Reference

### Channel Model

**Namespace:** `Core\Config\Models\Channel`

#### Properties

| Property | Type | Description |
|----------|------|-------------|
| `code` | string | Unique identifier |
| `name` | string | Human-readable label |
| `parent_id` | int\|null | Parent channel for inheritance |
| `workspace_id` | int\|null | Owner (null = system) |
| `metadata` | array\|null | Arbitrary JSON data |

#### Methods

```php
// Find by code (prefers workspace-specific over system)
Channel::byCode('instagram', $workspaceId): ?Channel

// Get inheritance chain (most specific first)
$channel->inheritanceChain(): Collection

// Get all codes in chain
$channel->inheritanceCodes(): array  // ['instagram', 'social']

// Check inheritance
$channel->inheritsFrom('social'): bool

// Is system channel?
$channel->isSystem(): bool

// Get metadata value
$channel->meta('platform_id'): mixed

// Ensure channel exists
Channel::ensure(
    code: 'instagram',
    name: 'Instagram',
    parentCode: 'social',
    workspaceId: null,
    metadata: ['platform_id' => 'ig'],
): Channel
```

### ConfigService with Channels

```php
$config = app(ConfigService::class);

// Set context (typically in middleware)
$config->setContext($workspace, $channel);

// Get value using context
$value = $config->get('social.posting.style');

// Explicit channel override
$result = $config->resolve('social.posting.style', $workspace, 'instagram');

// Set channel-specific value
$config->set(
    keyCode: 'social.posting.style',
    value: 'casual',
    profile: $profile,
    locked: false,
    channel: 'instagram',
);

// Lock a channel-specific value
$config->lock('social.posting.style', $profile, 'instagram');

// Prime for specific channel
$config->prime($workspace, 'instagram');
```

### ConfigValue with Channels

```php
// Find value for profile + key + channel
ConfigValue::findValue($profileId, $keyId, $channelId): ?ConfigValue

// Set value with channel
ConfigValue::setValue(
    profileId: $profileId,
    keyId: $keyId,
    value: 'casual',
    locked: false,
    inheritedFrom: null,
    channelId: $channelId,
): ConfigValue

// Get all values for key across profiles and channels
ConfigValue::forKeyInProfiles($keyId, $profileIds, $channelIds): Collection
```

---

## Database Schema

### config_channels

```sql
CREATE TABLE config_channels (
    id BIGINT PRIMARY KEY,
    code VARCHAR(255),
    name VARCHAR(255),
    parent_id BIGINT REFERENCES config_channels(id),
    workspace_id BIGINT REFERENCES workspaces(id),
    metadata JSON,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    UNIQUE (code, workspace_id),
    INDEX (parent_id)
);
```

### config_values (extended)

```sql
ALTER TABLE config_values ADD COLUMN
    channel_id BIGINT REFERENCES config_channels(id);

-- Updated unique constraint
UNIQUE (profile_id, key_id, channel_id)
```

### config_resolved (extended)

```sql
-- Channel dimension in resolved cache
channel_id BIGINT,
source_channel_id BIGINT,

-- Composite lookup
INDEX (workspace_id, channel_id, key_code)
```

---

## Examples

### Multi-platform social posting

```php
// System defaults (all channels)
$config->set('social.posting.max_length', 280, $systemProfile);
$config->set('social.posting.style', 'professional', $systemProfile);

// Channel-specific overrides
$config->set('social.posting.max_length', 2200, $systemProfile, channel: 'instagram');
$config->set('social.posting.max_length', 100000, $systemProfile, channel: 'linkedin');
$config->set('social.posting.style', 'casual', $workspaceProfile, channel: 'tiktok');

// Resolution
$config->resolve('social.posting.max_length', $workspace, 'twitter');  // 280 (default)
$config->resolve('social.posting.max_length', $workspace, 'instagram'); // 2200
$config->resolve('social.posting.style', $workspace, 'tiktok');        // 'casual'
```

### API rate limiting with FINAL

```php
// System admin sets hard limit for API channel
$config->set('api.rate_limit.requests', 1000, $systemProfile, locked: true, channel: 'api');
$config->set('api.rate_limit.window', 60, $systemProfile, locked: true, channel: 'api');

// Workspaces cannot exceed this
$config->set('api.rate_limit.requests', 5000, $workspaceProfile, channel: 'api');
// ↑ Stored but never returned - locked value wins

$config->resolve('api.rate_limit.requests', $workspace, 'api'); // Always 1000
```

### Voice/tone channels

```php
// Define voice channels
Channel::ensure('support', 'Customer Support', parentCode: null);
Channel::ensure('vi', 'Virtual Intelligence', parentCode: null);
Channel::ensure('formal', 'Formal Communications', parentCode: null);

// Configure per voice
$config->set('comms.greeting', 'Hello', $workspaceProfile, channel: null);
$config->set('comms.greeting', 'Hey there!', $workspaceProfile, channel: 'support');
$config->set('comms.greeting', 'Greetings', $workspaceProfile, channel: 'formal');
$config->set('comms.greeting', 'Hi, I\'m your AI assistant', $workspaceProfile, channel: 'vi');
```

### Channel inheritance

```php
// Create hierarchy
Channel::ensure('social', 'Social Media');
Channel::ensure('instagram', 'Instagram', parentCode: 'social');
Channel::ensure('instagram_stories', 'Instagram Stories', parentCode: 'instagram');

// Set at parent level
$config->set('social.hashtags.enabled', true, $profile, channel: 'social');
$config->set('social.hashtags.max', 30, $profile, channel: 'instagram');

// Child inherits from parent
$config->resolve('social.hashtags.enabled', $workspace, 'instagram_stories');
// → true (inherited from 'social')

$config->resolve('social.hashtags.max', $workspace, 'instagram_stories');
// → 30 (inherited from 'instagram')
```

### Workspace-specific channel override

```php
// System channel
Channel::ensure('premium', 'Premium Features', workspaceId: null);

// Workspace overrides system channel
Channel::ensure('premium', 'VIP Premium', workspaceId: $workspace->id, metadata: [
    'features' => ['priority_support', 'custom_branding'],
]);

// Lookup prefers workspace channel
$channel = Channel::byCode('premium', $workspace->id);
// → Workspace's 'VIP Premium' channel, not system 'Premium Features'
```

---

## Implementation Notes

### Performance Considerations

The channel system adds a dimension to resolution, but performance impact is minimal:

1. **Read path unchanged** — Most reads hit the hash (O(1))
2. **Batch loading** — Resolution loads all channel values in one query
3. **Cached resolution** — `config_resolved` stores pre-computed values per workspace+channel
4. **Lazy priming** — Only computes on first access, not on every request

### Cycle Detection

Channel inheritance includes cycle detection to handle data corruption:

```php
public function inheritanceChain(): Collection
{
    $seen = [$this->id => true];

    while ($current->parent_id !== null) {
        if (isset($seen[$current->parent_id])) {
            Log::error('Circular reference in channel inheritance');
            break;
        }
        // ...
    }
}
```

### MariaDB NULL Handling

The `config_resolved` table uses `0` instead of `NULL` for system scope and all-channels:

```php
// MariaDB composite unique constraints don't handle NULL well
// workspace_id = 0 means system scope
// channel_id = 0 means all channels
```

This is an implementation detail—the API accepts and returns `null` as expected.

---

## Related Files

- `app/Core/Config/Models/Channel.php` — Channel model
- `app/Core/Config/Models/ConfigValue.php` — Value storage with channel support
- `app/Core/Config/ConfigResolver.php` — Resolution engine
- `app/Core/Config/ConfigService.php` — Main API
- `app/Core/Config/Migrations/2026_01_09_100001_add_config_channels.php` — Schema

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0 | 2026-01-15 | Initial RFC |
