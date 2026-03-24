# Storage Offload

Offload uploads to S3-compatible storage (AWS S3, Hetzner Object Storage, etc.) with transparent URL rewriting and CDN integration.

## Features

- Upload files to S3-compatible storage
- Transparent URL rewriting in API responses
- CDN integration (BunnyCDN, CloudFlare, etc.)
- File integrity verification with SHA-256 hashing
- Configurable retention (keep or delete local copies)
- Migration command for existing files
- Caching for improved performance

## Configuration

### Environment Variables

```bash
# Enable storage offload
STORAGE_OFFLOAD_ENABLED=true

# Storage disk (hetzner, s3, or custom)
STORAGE_OFFLOAD_DISK=hetzner

# CDN URL (optional, uses disk URL if not set)
STORAGE_OFFLOAD_CDN_URL=https://cdn.host.uk.com

# Hetzner Object Storage
HETZNER_ENDPOINT=https://fsn1.your-objectstorage.com
HETZNER_REGION=fsn1
HETZNER_BUCKET=host-uk-uploads
HETZNER_ACCESS_KEY=your_access_key
HETZNER_SECRET_KEY=your_secret_key

# Or AWS S3
AWS_ACCESS_KEY_ID=your_key
AWS_SECRET_ACCESS_KEY=your_secret
AWS_DEFAULT_REGION=eu-west-2
AWS_BUCKET=host-uk-uploads

# Optional settings
STORAGE_OFFLOAD_MAX_SIZE=104857600  # 100MB in bytes
STORAGE_OFFLOAD_AUTO=true           # Auto-offload on upload
STORAGE_OFFLOAD_KEEP_LOCAL=false    # Delete local copy after upload
STORAGE_OFFLOAD_QUEUE=true          # Queue long operations
```

### Configuration File

See `/config/offload.php` for full configuration options including:

- Allowed file extensions
- Path organisation by category
- Queue settings
- Cache settings

## Usage

### Programmatic Upload

```php
use App\Services\Storage\StorageOffload;

$offloadService = app(StorageOffload::class);

// Upload a file
$localPath = storage_path('app/public/uploads/photo.jpg');
$result = $offloadService->upload(
    localPath: $localPath,
    remotePath: null,              // Auto-generated if null
    category: 'biolink',           // For path organisation
    metadata: [                    // Optional metadata
        'user_id' => auth()->id(),
        'original_name' => 'photo.jpg',
    ]
);

// Get CDN URL for offloaded file
$url = $offloadService->url($localPath);
// Returns: https://cdn.host.uk.com/biolinks/abc123/photo.jpg

// Check if file is offloaded
if ($offloadService->isOffloaded($localPath)) {
    // File is on remote storage
}

// Delete from remote storage
$offloadService->delete($localPath);

// Verify file integrity
$valid = $offloadService->verifyIntegrity($localPath);

// Get statistics
$stats = $offloadService->getStats();
```

### Artisan Command

Migrate existing local files to remote storage:

```bash
# Migrate all files in storage/app/public
php artisan offload:migrate

# Migrate specific directory
php artisan offload:migrate /path/to/directory

# Dry run (preview without uploading)
php artisan offload:migrate --dry-run

# Set category for organisation
php artisan offload:migrate --category=biolink

# Only migrate files not already offloaded
php artisan offload:migrate --only-missing

# Skip confirmation prompt
php artisan offload:migrate --force
```

### URL Rewriting Middleware

Apply to API routes for transparent URL rewriting:

```php
use App\Http\Middleware\RewriteOffloadedUrls;

// In routes/api.php or route groups
Route::middleware([RewriteOffloadedUrls::class])->group(function () {
    Route::get('/biolinks', [BioLinkController::class, 'index']);
});
```

The middleware automatically rewrites local storage URLs to CDN URLs in JSON responses:

```json
{
  "avatar": "https://cdn.host.uk.com/avatars/abc123/photo.jpg",
  "background": "https://cdn.host.uk.com/biolinks/def456/bg.jpg"
}
```

## Database

The `storage_offloads` table tracks offloaded files:

| Column | Description |
|--------|-------------|
| `disk` | Storage disk name (hetzner, s3, etc.) |
| `local_path` | Original local file path |
| `remote_path` | Remote storage path |
| `file_size` | File size in bytes |
| `mime_type` | MIME type |
| `hash` | SHA-256 hash for integrity checking |
| `category` | Category for organisation |
| `metadata` | JSON metadata |
| `offloaded_at` | Upload timestamp |

## Model

Query offloaded files:

```php
use App\Models\StorageOffload;

// Find by local path
$record = StorageOffload::where('local_path', $path)->first();

// Filter by category
$biolinks = StorageOffload::inCategory('biolink')->get();

// Filter by disk
$hetznerFiles = StorageOffload::forDisk('hetzner')->get();

// Check file type
if ($record->isImage()) {
    // Handle image
}

// Human-readable file size
echo $record->file_size_human; // "2.5 MB"
```

## Path Organisation

Files are organised by category in remote storage:

| Category | Remote Path Pattern |
|----------|-------------------|
| `biolink` | `biolinks/{hash}/{filename}` |
| `avatar` | `avatars/{hash}/{filename}` |
| `media` | `media/{hash}/{filename}` |
| `static` | `static/{hash}/{filename}` |

Configure in `config/offload.php` under `paths`.

## CDN Integration

### BunnyCDN (Recommended)

1. Create a pull zone pointing to your origin server
2. Set `STORAGE_OFFLOAD_CDN_URL` to your pull zone URL
3. Files are automatically pulled from remote storage and cached at edge

### CloudFlare

1. Add CNAME record pointing to your storage endpoint
2. Enable caching in CloudFlare settings
3. Set `STORAGE_OFFLOAD_CDN_URL` to your CloudFlare URL

### Direct Access

If no CDN URL is configured, the middleware uses the storage disk's URL configuration.

## Performance

- **Caching**: URLs are cached for 1 hour by default (configurable)
- **Queue**: Long-running migrations can be queued
- **Integrity**: SHA-256 hashing ensures file integrity without re-reading files

## Monitoring

Check offload statistics:

```php
$stats = $offloadService->getStats();

// Output:
// [
//   'total_files' => 1234,
//   'total_size' => 5368709120,
//   'total_size_human' => '5 GB',
//   'by_category' => [
//     ['category' => 'biolink', 'file_count' => 500, 'total_size' => ...],
//     ['category' => 'avatar', 'file_count' => 734, 'total_size' => ...],
//   ]
// ]
```

## Troubleshooting

### Files not uploading

Check logs for errors:
```bash
tail -f storage/logs/laravel.log | grep offload
```

Common issues:
- Incorrect credentials in `.env`
- File size exceeds `STORAGE_OFFLOAD_MAX_SIZE`
- File extension not in allowed list

### URLs not rewriting

1. Ensure middleware is applied to route
2. Check response is JSON (middleware only processes JSON)
3. Verify file is actually offloaded: `php artisan tinker` → `StorageOffload::count()`

### Integrity verification fails

File may be corrupted during upload. Re-upload:

```php
$offloadService->delete($localPath);
$offloadService->upload($localPath, null, $category);
```

## Migration Guide

### Existing Laravel Storage

```bash
# 1. Configure environment variables
nano .env

# 2. Test with dry run
php artisan offload:migrate --dry-run

# 3. Migrate with local backup
STORAGE_OFFLOAD_KEEP_LOCAL=true php artisan offload:migrate

# 4. Verify uploads
php artisan tinker
>>> StorageOffload::count()

# 5. Remove local copies once verified
STORAGE_OFFLOAD_KEEP_LOCAL=false php artisan offload:migrate
```

### Gradual Migration

Enable auto-offload for new uploads only:

```bash
STORAGE_OFFLOAD_AUTO=true
STORAGE_OFFLOAD_ENABLED=true
```

Migrate existing files in batches:

```bash
php artisan offload:migrate storage/app/public/2024
php artisan offload:migrate storage/app/public/2025
```

## Security

- Files maintain original permissions (public/private based on disk config)
- SHA-256 hashing prevents tampering
- Credentials stored in `.env` (never committed to git)
- Soft deletes allow recovery of accidentally deleted records

## Best Practices

1. Use CDN for public assets (images, documents)
2. Keep local copies during initial migration
3. Monitor storage costs (especially with BunnyCDN per-connection pricing for WebSockets)
4. Set appropriate `max_file_size` to avoid huge uploads
5. Use categories for organised storage structure
6. Enable caching for frequently accessed URLs
7. Queue large migration operations

## Testing

```bash
# Run all storage offload tests
./vendor/bin/pest tests/Feature/StorageOffloadTest.php
./vendor/bin/pest tests/Feature/RewriteOffloadedUrlsTest.php
./vendor/bin/pest tests/Feature/OffloadMigrateCommandTest.php
```

## See Also

- `/config/offload.php` - Full configuration options
- `/config/filesystems.php` - Disk configurations
- `BunnyCdnService` - Existing BunnyCDN integration for cache purging
