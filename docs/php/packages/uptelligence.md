# Uptelligence

Uptelligence is a vendor dependency monitoring package that tracks upstream releases, analyses diffs, generates upgrade plans, and sends digest notifications. It supports GitHub, Gitea, and AltumCode platforms.

## Vendor Update Checking

The `VendorUpdateCheckerService` checks all active vendors for new releases by querying their upstream sources.

### Supported Platforms

| Platform | Source Type | Check Method |
|---|---|---|
| GitHub | OSS repos | GitHub Releases API |
| Gitea/Forgejo | OSS repos | Gitea Releases API |
| AltumCode | Licensed products | Public `info.php` endpoints |
| AltumCode | Plugins | `dev.altumcode.com/plugins-versions` |

### Running Checks

```bash
# Check all active vendors
php artisan uptelligence:check-updates

# Output:
# Product          Deployed    Latest     Status
# ──────────────────────────────────────────────
# 66analytics      65.0.0      66.0.0     UPDATE AVAILABLE
# 66biolinks       65.0.0      66.0.0     UPDATE AVAILABLE
# 66pusher         65.0.0      65.0.0     ✓ current
# laravel/framework 12.0.1     12.1.0     UPDATE AVAILABLE
```

### AltumCode Version Detection

AltumCode products expose public version endpoints that require no authentication:

| Endpoint | Returns |
|---|---|
| `https://66analytics.com/info.php` | `{"latest_release_version": "66.0.0", ...}` |
| `https://66biolinks.com/info.php` | Same format |
| `https://66pusher.com/info.php` | Same format |
| `https://66socialproof.com/info.php` | Same format |
| `https://dev.altumcode.com/plugins-versions` | All plugin versions in one response |

Plugin versions are cached in memory during a check run to avoid redundant HTTP calls.

### Syncing Deployed Versions

The `SyncAltumVersionsCommand` reads actual deployed versions from source files on disk:

```bash
# Show what would change
php artisan uptelligence:sync-altum-versions --dry-run

# Sync from default path
php artisan uptelligence:sync-altum-versions

# Sync from custom path
php artisan uptelligence:sync-altum-versions --path=/path/to/saas/services
```

The command reads:

- **Product versions** from `PRODUCT_VERSION` defines in each product's `app/init.php`
- **Plugin versions** from `'version'` entries in each plugin's `config.php`

Output is a table showing old version, new version, and status (UPDATED, current, or SKIPPED).

## Vendor Model

Vendors are tracked in the `uptelligence_vendors` table:

```php
use Core\Mod\Uptelligence\Models\Vendor;

// Source types
Vendor::SOURCE_OSS;       // Open source (GitHub/Gitea)
Vendor::SOURCE_LICENSED;  // Licensed products (AltumCode)
Vendor::SOURCE_PLUGIN;    // Plugins (AltumCode)

// Platform types
Vendor::PLATFORM_ALTUM;   // AltumCode products/plugins

// Query active vendors
$active = Vendor::active()->get();
```

## Services

| Service | Description |
|---|---|
| `VendorUpdateCheckerService` | Checks upstream sources for new releases |
| `DiffAnalyzerService` | Analyses differences between versions |
| `AIAnalyzerService` | AI-powered analysis of changes |
| `IssueGeneratorService` | Creates issues for available updates |
| `UpstreamPlanGeneratorService` | Generates upgrade plans |
| `UptelligenceDigestService` | Compiles and sends update digests |
| `VendorStorageService` | Manages downloaded vendor files |
| `WebhookReceiverService` | Receives webhook notifications |
| `AssetTrackerService` | Tracks vendor assets |

## Commands

| Command | Description |
|---|---|
| `uptelligence:check-updates` | Check all vendors for new releases |
| `uptelligence:sync-altum-versions` | Sync deployed AltumCode versions from source |
| `uptelligence:sync-forge` | Sync vendors from Forge repositories |
| `uptelligence:analyze` | Run AI analysis on pending updates |
| `uptelligence:issues` | Generate issues for available updates |
| `uptelligence:send-digests` | Send update digest notifications |

## AltumCode Vendor Seeder

Seed the vendors table with all AltumCode products and plugins:

```bash
php artisan db:seed --class="Core\Mod\Uptelligence\Database\Seeders\AltumCodeVendorSeeder"
```

This creates 4 product entries and 13 plugin entries. The seeder is idempotent — it uses `updateOrCreate` so it can be run repeatedly without creating duplicates.

## Learn More

- [Developer Package](/php/packages/developer/)
