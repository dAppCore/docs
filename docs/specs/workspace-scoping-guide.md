# Workspace Scoping Guide for Developers

## Quick Start

When creating a new model that belongs to a workspace, add the `BelongsToWorkspace` trait:

```php
<?php

declare(strict_types=1);

namespace App\Models\Social;

use Mod\Tenant\Models\Workspace;
use App\Traits\BelongsToWorkspace;
use Illuminate\Database\Eloquent\Model;

class YourModel extends Model
{
    use BelongsToWorkspace;

    protected $fillable = [
        'workspace_id',
        // ... other fields
    ];
}
```

That's it! The trait automatically handles:
- ✅ Auto-assigning `workspace_id` on creation
- ✅ Cache invalidation when models change
- ✅ Query scopes for workspace filtering
- ✅ Helper methods for ownership checks

## What You Get

### 1. Automatic Workspace Assignment

```php
// If user is authenticated with a workspace
actingAs($user);

$account = Account::create([
    'name' => 'My Account',
    // workspace_id automatically set to $user->defaultHostWorkspace()->id
]);
```

### 2. Query Scopes

```php
// Get all records for current user's workspace
$accounts = Account::ownedByCurrentWorkspace()->get();

// Get all records for a specific workspace
$accounts = Account::forWorkspace($workspace)->get();
$accounts = Account::forWorkspace($workspaceId)->get();

// With additional filters
$activeAccounts = Account::ownedByCurrentWorkspace()
    ->where('status', 'active')
    ->get();
```

### 3. Cached Queries

```php
// Get cached collection for current workspace (5 min TTL)
$accounts = Account::ownedByCurrentWorkspaceCached();

// Custom TTL (in seconds)
$accounts = Account::ownedByCurrentWorkspaceCached(600); // 10 minutes

// Get cached collection for specific workspace
$accounts = Account::forWorkspaceCached($workspace, 300); // 5 min
```

Cache is automatically invalidated when any model in that workspace is saved or deleted.

### 4. Ownership Helpers

```php
// Check if model belongs to a specific workspace
if ($account->belongsToWorkspace($workspace)) {
    // Allow access
}

// Check if model belongs to current user's workspace
if ($account->belongsToCurrentWorkspace()) {
    // Allow edit
}
```

### 5. Relationships

The trait provides a `workspace()` relationship:

```php
$account = Account::find(1);
$workspace = $account->workspace; // BelongsTo relationship
$workspaceName = $account->workspace->name;
```

## When NOT to Use

Don't use `BelongsToWorkspace` for:

### Pivot Tables
```php
// Don't use on pivots
class PostAccount extends Pivot
{
    // No BelongsToWorkspace here
}
```

### Shared Resources
```php
// Media might be shared across workspaces
class Media extends Model
{
    // Decide based on your sharing model
    // If media is workspace-specific, use the trait
    // If media is shared globally, don't use it
}
```

### System/Admin Models
```php
// Global system configuration
class SystemConfig extends Model
{
    // No workspace scoping
}
```

## Advanced Usage

### Manual Cache Control

```php
// Clear cache for a specific workspace
Account::clearWorkspaceCache($workspaceId);

// Clear cache for all workspaces (expensive!)
Account::clearAllWorkspaceCache();
```

### Working Without Current User

```php
// In queues, commands, or system processes
$workspace = Workspace::find($workspaceId);

$accounts = Account::forWorkspace($workspace)->get();
$accounts = Account::forWorkspaceCached($workspace);
```

### Global Scope Alternative

If you prefer Laravel's global scopes instead of the trait approach:

```php
use App\Scopes\WorkspaceScope;

class Account extends Model
{
    protected static function booted(): void
    {
        static::addGlobalScope(new WorkspaceScope);
    }

    // Note: You lose the caching features with this approach
}
```

Then in queries:

```php
// Query specific workspace
Account::forWorkspace($workspace)->get();

// Query across all workspaces (admin only!)
Account::acrossWorkspaces()->get();

// Remove scope for one query
Account::withoutGlobalScope(WorkspaceScope::class)->get();
```

## Testing

### With Factories

```php
use Mod\Tenant\Models\Workspace;
use Mod\Social\Models\Account;

it('scopes accounts to workspace', function () {
    $workspace1 = Workspace::factory()->create();
    $workspace2 = Workspace::factory()->create();

    $account1 = Account::factory()->create(['workspace_id' => $workspace1->id]);
    $account2 = Account::factory()->create(['workspace_id' => $workspace2->id]);

    // Test scoping
    $accounts = Account::forWorkspace($workspace1)->get();

    expect($accounts)->toHaveCount(1);
    expect($accounts->first()->id)->toBe($account1->id);
});
```

### With Authentication

```php
it('auto-assigns workspace from authenticated user', function () {
    $user = User::factory()->create();
    $workspace = Workspace::factory()->create();
    $user->hostWorkspaces()->attach($workspace, ['is_default' => true]);

    actingAs($user);

    $account = Account::factory()->create();

    expect($account->workspace_id)->toBe($workspace->id);
});
```

### Testing Cache

```php
it('caches workspace queries', function () {
    $workspace = Workspace::factory()->create();
    Account::factory()->count(3)->create(['workspace_id' => $workspace->id]);

    // First call - hits database
    $accounts1 = Account::forWorkspaceCached($workspace);

    // Second call - hits cache (faster)
    $accounts2 = Account::forWorkspaceCached($workspace);

    expect($accounts1)->toEqual($accounts2);
    expect($accounts1)->toHaveCount(3);
});

it('invalidates cache on save', function () {
    $workspace = Workspace::factory()->create();
    $account = Account::factory()->create(['workspace_id' => $workspace->id]);

    $cached1 = Account::forWorkspaceCached($workspace);
    expect($cached1)->toHaveCount(1);

    // Create another account - cache should invalidate
    Account::factory()->create(['workspace_id' => $workspace->id]);

    $cached2 = Account::forWorkspaceCached($workspace);
    expect($cached2)->toHaveCount(2); // Fresh data
});
```

## Common Patterns

### Controller Usage

```php
class AccountController extends Controller
{
    public function index()
    {
        // Simple cached list for current workspace
        $accounts = Account::ownedByCurrentWorkspaceCached();

        return view('accounts.index', ['accounts' => $accounts]);
    }

    public function store(Request $request)
    {
        $validated = $request->validate([...]);

        // workspace_id automatically assigned
        $account = Account::create($validated);

        return redirect()->route('accounts.show', $account);
    }

    public function show(Account $account)
    {
        // Verify ownership before showing
        abort_unless($account->belongsToCurrentWorkspace(), 403);

        return view('accounts.show', ['account' => $account]);
    }
}
```

### API Usage

```php
class AccountApiController extends Controller
{
    public function index(Request $request)
    {
        $workspace = $request->user()->defaultHostWorkspace();

        // Don't cache API responses (clients should cache)
        $accounts = Account::forWorkspace($workspace)
            ->when($request->status, fn($q) => $q->where('status', $request->status))
            ->paginate(20);

        return AccountResource::collection($accounts);
    }
}
```

### Livewire Components

```php
use Livewire\Component;
use Mod\Social\Models\Account;

class AccountList extends Component
{
    public function mount()
    {
        // Cached for performance
        $this->accounts = Account::ownedByCurrentWorkspaceCached();
    }

    public function delete(Account $account)
    {
        // Security check
        if (!$account->belongsToCurrentWorkspace()) {
            $this->addError('general', 'Unauthorised');
            return;
        }

        $account->delete();
        // Cache automatically invalidated

        $this->accounts = Account::ownedByCurrentWorkspaceCached();
    }

    public function render()
    {
        return view('admin.account-list');
    }
}
```

### Queue Jobs

```php
use Mod\Social\Models\Account;
use Illuminate\Bus\Queueable;

class SyncAccountJob implements ShouldQueue
{
    use Queueable;

    public function __construct(
        public int $accountId,
        public int $workspaceId
    ) {}

    public function handle()
    {
        // No current user in queue context, use explicit workspace
        $account = Account::forWorkspace($this->workspaceId)
            ->findOrFail($this->accountId);

        // Do sync work...
    }
}
```

## Migration Checklist

When migrating a model from MixPost or adding workspace scoping:

- [ ] Add `workspace_id` column to migration
- [ ] Add `workspace_id` to `$fillable` array
- [ ] Add `use BelongsToWorkspace;` trait
- [ ] Remove any MixPost trait imports
- [ ] Update foreign key constraints if needed
- [ ] Update factory to include `workspace_id`
- [ ] Add tests for workspace scoping
- [ ] Update controllers to use workspace helpers
- [ ] Check for N+1 queries and add caching where beneficial

## Performance Tips

### Cache Aggressively
```php
// Good for frequently accessed, rarely changing data
$settings = Config::ownedByCurrentWorkspaceCached(3600); // 1 hour
```

### Don't Cache User-Specific Queries
```php
// Bad - per-user cache keys cause cache bloat
$myPosts = Post::where('user_id', $userId)
    ->ownedByCurrentWorkspace()
    ->get(); // Don't cache this

// Good - workspace-wide cache benefits all users
$workspaceAccounts = Account::ownedByCurrentWorkspaceCached();
```

### Use Eager Loading
```php
// Good - one query for workspace, one for accounts
$workspace = auth()->user()->defaultHostWorkspace();
$accounts = Account::forWorkspace($workspace)
    ->with('metrics', 'analytics')
    ->get();

// Bad - N+1 query problem
$accounts = Account::ownedByCurrentWorkspace()->get();
foreach ($accounts as $account) {
    $metrics = $account->metrics; // Separate query each time
}
```

## Troubleshooting

### "workspace_id cannot be null"
```php
// Problem: No authenticated user or user has no default workspace
Account::create([...]);

// Solution 1: Ensure user is authenticated
actingAs($user);
Account::create([...]);

// Solution 2: Explicitly set workspace_id
Account::create([
    'workspace_id' => $workspace->id,
    ...
]);
```

### Cache not invalidating
```php
// Make sure you're using save(), not update() on query builder
$account->status = 'active';
$account->save(); // ✅ Invalidates cache

// This won't trigger cache invalidation:
Account::where('id', $id)->update(['status' => 'active']); // ❌
```

### Getting empty results in tests
```php
// Problem: No workspace context in test
it('lists accounts', function () {
    $account = Account::factory()->create();
    $accounts = Account::ownedByCurrentWorkspace()->get();
    expect($accounts)->toHaveCount(1); // ❌ Fails - no current user
});

// Solution: Set up proper context
it('lists accounts', function () {
    $user = User::factory()->create();
    $workspace = Workspace::factory()->create();
    $user->hostWorkspaces()->attach($workspace, ['is_default' => true]);

    actingAs($user);

    $account = Account::factory()->create();
    $accounts = Account::ownedByCurrentWorkspace()->get();
    expect($accounts)->toHaveCount(1); // ✅ Works
});
```

## Further Reading

- [Workspace Model](app/Mod/Tenant/Models/Workspace.php) - Core workspace model
- [BelongsToWorkspace Trait](app/Traits/BelongsToWorkspace.php) - Trait source code
- [WorkspaceScope](app/Scopes/WorkspaceScope.php) - Global scope alternative

---

**Last Updated:** 2026-01-01
**Maintained By:** Host UK Development Team
