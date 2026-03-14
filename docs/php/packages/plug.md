# Plug Packages

The Plug system provides a unified interface for integrating with external platforms — social media, messaging, content publishing, CDN, storage, and stock media services. It is split into a framework layer (contracts, registry, shared traits) in `core/php` and nine domain-specific packages.

## Architecture

```
core/php (framework)
├── src/Plug/Contract/*        ← 8 shared interfaces
├── src/Plug/Registry.php      ← provider registry
├── src/Plug/Response.php      ← standardised operation response
├── src/Plug/Concern/*         ← shared traits (UsesHttp, ManagesTokens, BuildsResponse)
├── src/Plug/Enum/Status.php   ← OK, UNAUTHORIZED, RATE_LIMITED, etc.
└── src/Plug/Boot.php          ← registry singleton registration

core/php-plug-social    ─┐
core/php-plug-web3      ─┤
core/php-plug-content   ─┤
core/php-plug-chat      ─┤
core/php-plug-business  ─┤  all depend on core/php
core/php-plug-cdn       ─┤
core/php-plug-storage   ─┤
core/php-plug-stock     ─┤
core/php-plug-altum     ─┘
```

All provider namespaces are under `Core\Plug\*`, matching the framework convention.

## Shared Contracts

Eight interfaces in `Core\Plug\Contract\` define the capabilities a provider can implement:

| Contract | Description |
|---|---|
| `Authenticable` | OAuth/token-based authentication |
| `Postable` | Create posts or content |
| `Readable` | Read posts, profiles, feeds |
| `Deletable` | Delete posts or content |
| `Commentable` | Read/write comments |
| `Listable` | List pages, boards, groups |
| `MediaUploadable` | Upload images, videos, media |
| `Refreshable` | Refresh OAuth tokens |

## Registry

The `Registry` class manages provider discovery and capability checking:

```php
use Core\Plug\Registry;

$registry = app(Registry::class);

// Check if a provider exists
$registry->has('twitter');         // true

// Check capabilities
$registry->supports('twitter', 'post');    // true
$registry->supports('twitter', 'comment'); // true

// Get an operation class
$postClass = $registry->operation('twitter', 'post');
// Returns: Core\Plug\Social\Twitter\Post

// Browse by category
$social = $registry->byCategory('Social'); // Collection of identifiers

// Find providers with a capability
$postable = $registry->withCapability('post'); // All providers that support posting
```

### Programmatic Registration

Packages self-register their providers via `Registry::register()`:

```php
$registry->register(
    identifier: 'twitter',
    category: 'Social',
    name: 'Twitter',
    namespace: 'Core\Plug\Social\Twitter',
);
```

## Shared Traits

### `UsesHttp`

Provides a pre-configured HTTP client with JSON acceptance:

```php
use Core\Plug\Concern\UsesHttp;

class Post
{
    use UsesHttp;

    public function create(array $data): Response
    {
        $response = $this->http()
            ->withToken($this->accessToken)
            ->post('https://api.twitter.com/2/tweets', $data);

        // ...
    }
}
```

### `ManagesTokens`

OAuth token storage and refresh logic.

### `BuildsResponse`

Fluent builder for `Core\Plug\Response` objects with status, data, and error handling.

## Standardised Response

All Plug operations return a `Core\Plug\Response`:

```php
use Core\Plug\Response;
use Core\Plug\Enum\Status;

$response = new Response(
    status: Status::OK,
    data: ['id' => '123456', 'url' => 'https://twitter.com/...'],
);

$response->successful(); // true
$response->data();       // ['id' => '123456', ...]
$response->status();     // Status::OK
```

Status values: `OK`, `UNAUTHORIZED`, `RATE_LIMITED`, `NOT_FOUND`, `ERROR`, `VALIDATION_ERROR`.

## Packages

### Social (`core/php-plug-social`)

8 social media providers for posting, reading, and managing content.

| Provider | Operations |
|---|---|
| Twitter | Auth, Post, Delete, Media, Read |
| Meta (Facebook/Instagram) | Auth, Post, Delete, Media, Pages, Read |
| LinkedIn | Auth, Post, Delete, Media, Pages, Read |
| Pinterest | Auth, Post, Delete, Media, Boards, Read |
| Reddit | Auth, Post, Delete, Media, Read, Subreddits |
| TikTok | Auth, Post, Read |
| VK | Auth, Post, Delete, Media, Groups, Read |
| YouTube | Auth, Post, Delete, Comment, Read |

### Web3 (`core/php-plug-web3`)

6 decentralised/federated platform providers.

| Provider | Operations |
|---|---|
| Bluesky | Auth, Post, Delete, Read |
| Farcaster | Auth, Post, Read |
| Lemmy | Auth, Post, Delete, Comment, Read |
| Mastodon | Auth, Post, Delete, Media, Read |
| Nostr | Auth, Post, Read |
| Threads | Auth, Post, Read |

### Content (`core/php-plug-content`)

4 content publishing platform providers.

| Provider | Operations |
|---|---|
| Dev.to | Auth, Post, Read |
| Hashnode | Auth, Post, Read |
| Medium | Auth, Post, Read |
| WordPress | Auth, Post, Delete, Read |

### Chat (`core/php-plug-chat`)

3 messaging platform providers.

| Provider | Operations |
|---|---|
| Discord | Auth, Post |
| Slack | Auth, Post |
| Telegram | Auth, Post |

### Business (`core/php-plug-business`)

| Provider | Operations |
|---|---|
| Google My Business | Auth, Post, Read |

### CDN (`core/php-plug-cdn`)

CDN management with domain-specific contracts (`Core\Plug\Cdn\Contract\Purgeable`, `HasStats`).

| Provider | Operations |
|---|---|
| Bunny CDN | Purge, Stats |

### Storage (`core/php-plug-storage`)

Object storage with domain-specific contracts (`Core\Plug\Storage\Contract\Browseable`, `Uploadable`, `Downloadable`, `Deletable`).

| Provider | Operations |
|---|---|
| Bunny Storage | Browse, Delete, Download, Upload, VBucket |

### Stock (`core/php-plug-stock`)

Stock media integrations.

| Provider | Operations |
|---|---|
| Unsplash | Search, Photo, Collection, Download |

### AltumCode (`core/php-plug-altum`)

Integration with AltumCode SaaS products (66analytics, 66biolinks, 66pusher, 66socialproof).

| Component | Description |
|---|---|
| `AltumClient` | HTTP client for AltumCode APIs |
| `AltumManager` | Multi-product management |
| `AltumServiceProvider` | Service registration |
| `AltumWebhookVerifier` | Webhook signature verification |

## Adding a Provider

Each provider is a set of operation classes in a subdirectory:

```php
<?php

declare(strict_types=1);

namespace Core\Plug\Social\Twitter;

use Core\Plug\Concern\BuildsResponse;
use Core\Plug\Concern\UsesHttp;
use Core\Plug\Contract\Postable;
use Core\Plug\Response;

class Post implements Postable
{
    use BuildsResponse, UsesHttp;

    public function create(string $accessToken, array $data): Response
    {
        $response = $this->http()
            ->withToken($accessToken)
            ->post('https://api.twitter.com/2/tweets', [
                'text' => $data['text'],
            ]);

        if (! $response->successful()) {
            return $this->error($response->status(), $response->body());
        }

        return $this->ok($response->json());
    }
}
```

## Composer Setup

Each plug package requires `core/php` and uses VCS repositories on Forge:

```json
{
    "require": {
        "core/php-plug-social": "dev-main"
    },
    "repositories": [
        {
            "type": "vcs",
            "url": "ssh://git@forge.lthn.ai:2223/core/php-plug-social.git"
        }
    ]
}
```

## Learn More

- [Actions Pattern](/php/features/actions)
- [API Package](/php/packages/api/)
