# Authentication Guards

Host UK's native authentication guard implementation for API token-based authentication.

## Overview

This is a **native rewrite** of MixPost's authentication guards, rebuilt from the ground up to work with Host UK's Laravel stack. It provides stateful API authentication using personal access tokens.

**Important:** This is NOT an integration or wrapper around MixPost code. We studied their implementation and built our own version with zero dependencies on `inovector/mixpost`.

## Architecture

### Components

| Component | Location | Purpose |
|-----------|----------|---------|
| **UserToken Model** | `/app/Models/UserToken.php` | Eloquent model for API tokens |
| **AccessTokenGuard** | `/app/Guards/AccessTokenGuard.php` | Authentication guard for Bearer tokens |
| **HasApiTokens Trait** | `/app/Traits/HasApiTokens.php` | User model methods for token management |
| **Migration** | `/database/migrations/2026_01_01_170628_create_user_tokens_table.php` | Database schema |
| **Factory** | `/database/factories/UserTokenFactory.php` | Test data generation |

### How It Works

1. **Token Creation**: User creates a token via `$user->createToken('Mobile App')`
2. **Storage**: Token is hashed with SHA-256 and stored in `user_tokens` table
3. **Authentication**: API requests include `Authorization: Bearer {token}` header
4. **Validation**: Guard hashes incoming token, looks up in database, checks expiry
5. **Usage Tracking**: Guard updates `last_used_at` timestamp on successful auth

## Database Schema

```sql
CREATE TABLE user_tokens (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    name VARCHAR(255) NOT NULL,           -- Human-readable token name
    token VARCHAR(64) UNIQUE NOT NULL,    -- SHA-256 hash of actual token
    last_used_at TIMESTAMP NULL,          -- Track usage
    expires_at TIMESTAMP NULL,            -- Optional expiration
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    INDEX (token),
    INDEX (user_id, created_at)
);
```

## Usage

### Creating Tokens

```php
use Mod\Tenant\Models\User;

$user = User::find(1);

// Create a token (never expires)
$result = $user->createToken('Mobile App');
$plainToken = $result['token'];    // Show this to user ONCE
$tokenModel = $result['model'];    // UserToken instance

// Create a token that expires
$result = $user->createToken(
    'Temporary Access',
    now()->addDays(7)
);

// The plain-text token is only available at creation time
// You MUST save it or show it to the user immediately
```

### Managing Tokens

```php
// List all tokens
$tokens = $user->tokens;

// Revoke a specific token
$user->revokeToken($tokenId);

// Revoke all tokens (e.g., on password change)
$user->revokeAllTokens();

// Check token validity
$token = UserToken::findToken($plainTextToken);
if ($token && $token->isValid()) {
    // Token is good
}
```

### Using in API Routes

```php
use Illuminate\Support\Facades\Route;

// Protect routes with the access_token guard
Route::middleware('auth:access_token')->prefix('v1')->group(function () {
    Route::get('/social/posts', [SocialPostController::class, 'index']);
    Route::post('/social/posts', [SocialPostController::class, 'store']);
});

// In your controller
public function index(Request $request)
{
    $user = $request->user(); // Authenticated user
    $workspace = $user->defaultHostWorkspace();

    // Your logic here
}
```

### Making API Requests

```bash
# Include token in Authorization header
curl -X GET https://host.uk.com/api/v1/social/posts \
  -H "Authorization: Bearer your-40-character-token-here" \
  -H "Accept: application/json"
```

## Security Features

### Token Hashing

Tokens are **never stored in plain text**. We use SHA-256 hashing:

```php
// When creating a token
$plainToken = Str::random(40);  // 40 random characters
$hashed = hash('sha256', $plainToken);  // Store this

// When authenticating
$incoming = $request->bearerToken();
$hashed = hash('sha256', $incoming);
$token = UserToken::where('token', $hashed)->first();
```

### Expiry Handling

Tokens can optionally expire:

```php
// Check if expired
if ($token->isExpired()) {
    return response()->json(['error' => 'Token expired'], 401);
}

// Only valid tokens authenticate
if ($token->isValid()) {
    // Not expired (or no expiry set)
}
```

### Usage Tracking

Every successful authentication updates `last_used_at`:

```php
// Preserves hasModifiedRecords state to avoid triggering model events
$token->recordUsage();
```

This is useful for:
- Detecting abandoned/unused tokens
- Security auditing
- Automatic cleanup of stale tokens

## Configuration

### Register the Guard

The guard is registered in `AppServiceProvider::boot()`:

```php
use Mod\Api\Guards\AccessTokenGuard;
use Illuminate\Support\Facades\Auth;

Auth::viaRequest('access_token', new AccessTokenGuard($this->app['auth']));
```

### Auth Config

Guard is defined in `config/auth.php`:

```php
'guards' => [
    'web' => [
        'driver' => 'session',
        'provider' => 'users',
    ],

    'access_token' => [
        'driver' => 'access_token',
        'provider' => 'users',
    ],
],
```

## Testing

### Factory Usage

```php
use Mod\Tenant\Models\User;
use Mod\Tenant\Models\UserToken;

// Create a token for testing
$user = User::factory()->create();
$token = UserToken::factory()
    ->for($user)
    ->withToken('test-token-12345')
    ->create();

// Test with known token
$response = $this->getJson('/api/v1/social/posts', [
    'Authorization' => 'Bearer test-token-12345',
]);

// Create expired token
$expiredToken = UserToken::factory()
    ->expired()
    ->create();

// Create token that expires in 7 days
$futureToken = UserToken::factory()
    ->expiresIn(7)
    ->create();

// Create recently-used token
$usedToken = UserToken::factory()
    ->used()
    ->create();
```

### Test Example

```php
use Mod\Tenant\Models\User;
use Mod\Tenant\Models\UserToken;

test('can authenticate with valid token', function () {
    $user = User::factory()->create();
    $result = $user->createToken('Test Token');

    $response = $this->getJson('/api/v1/social/posts', [
        'Authorization' => "Bearer {$result['token']}",
    ]);

    $response->assertOk();
});

test('cannot authenticate with expired token', function () {
    $user = User::factory()->create();
    $token = UserToken::factory()
        ->for($user)
        ->expired()
        ->withToken('expired-token')
        ->create();

    $response = $this->getJson('/api/v1/social/posts', [
        'Authorization' => 'Bearer expired-token',
    ]);

    $response->assertUnauthorized();
});
```

## Design Decisions

### Why Not Laravel Sanctum?

While Sanctum is excellent, we built our own solution because:

1. **Learning from MixPost**: We're studying their implementation patterns
2. **Full Control**: Can customise every aspect for Host UK's needs
3. **Simplicity**: Only need Bearer token auth, not full SPA + mobile token system
4. **No Extra Dependencies**: One less package to maintain

### Why SHA-256 Over Bcrypt?

- **Performance**: SHA-256 is faster for token lookups (not passwords!)
- **Token Uniqueness**: Tokens are already random 40-character strings
- **MixPost Compatibility**: Matches their approach for easier migration
- **Security**: Still secure since tokens are unguessable random strings

For **passwords**, we still use bcrypt/argon2 via Laravel's hashing.

### Why Separate Guard vs Middleware?

- **Guard**: Handles **authentication** (who is the user?)
- **Middleware**: Can handle **authorisation** (what can they do?)

This separation follows Laravel's architecture patterns.

## Migration from MixPost

If you're migrating existing MixPost token data:

```php
use Inovector\Mixpost\Models\UserToken as MixpostToken;
use Mod\Tenant\Models\UserToken as HostToken;

// Migrate tokens from MixPost to Host UK
MixpostToken::all()->each(function ($mixpostToken) {
    HostToken::create([
        'user_id' => $mixpostToken->user_id,
        'name' => $mixpostToken->name,
        'token' => $mixpostToken->token,  // Already hashed
        'last_used_at' => $mixpostToken->last_used_at,
        'expires_at' => $mixpostToken->expires_at,
        'created_at' => $mixpostToken->created_at,
        'updated_at' => $mixpostToken->updated_at,
    ]);
});
```

## Best Practices

### Token Naming

Use descriptive names so users can identify tokens:

```php
// Good
$user->createToken('iPhone 15 Pro');
$user->createToken('GitHub Actions CI');
$user->createToken('Zapier Integration');

// Bad
$user->createToken('Token 1');
$user->createToken('My Token');
```

### Token Rotation

Rotate tokens periodically for security:

```php
// Revoke old token and create new one
$user->revokeToken($oldTokenId);
$newToken = $user->createToken('Rotated Mobile Token', now()->addDays(90));
```

### Workspace Context

Always resolve the workspace for API requests:

```php
public function index(Request $request)
{
    $user = $request->user();
    $workspace = $user->defaultHostWorkspace();

    // Scope queries to workspace
    $posts = $workspace->socialPosts()->get();

    return response()->json($posts);
}
```

### Rate Limiting

Always rate limit API routes:

```php
Route::middleware(['auth:access_token', 'throttle:api'])
    ->prefix('v1')
    ->group(function () {
        // Protected routes
    });
```

## Troubleshooting

### "Unauthenticated" Response

Check:
1. Token is included in `Authorization: Bearer {token}` header
2. Token hasn't expired: `$token->expires_at`
3. Token exists in database (check hash matches)
4. Guard is registered in `AppServiceProvider`
5. Route uses `auth:access_token` middleware

### Token Not Found

```php
$token = UserToken::findToken($plainToken);
if (!$token) {
    // Either:
    // - Token was revoked
    // - Token value is incorrect
    // - Database not migrated
}
```

### Performance Issues

If you have millions of tokens:
1. Add index on `last_used_at` for cleanup queries
2. Regularly prune expired/unused tokens:

```php
// Artisan command to clean up tokens
UserToken::where('expires_at', '<', now())
    ->orWhere('last_used_at', '<', now()->subMonths(6))
    ->delete();
```

## Further Reading

- [Laravel Authentication](https://laravel.com/docs/authentication)
- [Custom Guards](https://laravel.com/docs/authentication#adding-custom-guards)
- [API Authentication](https://laravel.com/docs/sanctum#how-it-works)
- MixPost Reference: `/packages/mixpost-pro-team/src/Guards/AccessTokenGuard.php`
