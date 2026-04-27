# Host UK Deployment Guide

## Container Architecture

Single unified container serving both applications:

```
┌─────────────────────────────────────────────────────────┐
│  Port 80 (nginx)                                        │
├─────────────────────────────────────────────────────────┤
│  host.uk.com        → Laravel Host Hub (/app/public)    │
│  *.host.uk.com      → WordPress (/var/www/html)         │
├─────────────────────────────────────────────────────────┤
│  Services (supervisord):                                │
│  ├── nginx          - reverse proxy                     │
│  ├── php-fpm84      - serves both apps                  │
│  ├── queue-worker   - Laravel queues (x2)               │
│  └── scheduler      - Laravel cron                      │
└─────────────────────────────────────────────────────────┘
```

## Health Check

The container uses a simple nginx-level health check that **does not depend on PHP, database, or Redis**:

```
GET /healthz → 200 OK "ok\n"
```

This endpoint is handled directly by nginx and returns instantly. The container will pass health checks even if:
- Database is still connecting
- Redis is unavailable
- Laravel is caching configs
- Queue workers are restarting

### Coolify Configuration

Set the health check path to `/healthz` in Coolify's health check settings if customizable.

The Dockerfile defines:
```dockerfile
HEALTHCHECK --interval=10s --timeout=5s --start-period=30s --retries=3 \
    CMD curl -sf http://localhost/healthz || exit 1
```

## Environment Variables

### Required (Runtime)

Set these in Coolify as **Runtime** environment variables:

| Variable | Example | Description |
|----------|---------|-------------|
| `APP_KEY` | `base64:...` | Laravel encryption key |
| `APP_ENV` | `production` | Environment name |
| `APP_DEBUG` | `false` | Show detailed errors |
| `DB_HOST` | `mariadb` | Database hostname |
| `DB_DATABASE` | `host_hub` | Database name |
| `DB_USERNAME` | `root` | Database user |
| `DB_PASSWORD` | `secret` | Database password |
| `REDIS_HOST` | `redis` | Redis hostname |
| `REDIS_PASSWORD` | `null` | Redis password (if any) |

### WordPress Variables

| Variable | Example | Description |
|----------|---------|-------------|
| `WORDPRESS_DB_HOST` | `mariadb` | WordPress DB host |
| `WORDPRESS_DB_NAME` | `wordpress` | WordPress DB name |
| `WORDPRESS_DB_USER` | `root` | WordPress DB user |
| `WORDPRESS_DB_PASSWORD` | `secret` | WordPress DB password |

### Redis Variables

Redis runs inside the container by default (standalone, no auth). For multi-container replication:

| Variable | Default | Description |
|----------|---------|-------------|
| `REDIS_HOST` | `127.0.0.1` | Redis host (in-container by default) |
| `REDIS_PORT` | `6379` | Redis port |
| `REDIS_PASSWORD` | `null` | Redis auth password |

### Redis Replication (Multi-Container)

For high availability across multiple containers, set these to enable master/replica with Sentinel failover:

| Variable | Example | Description |
|----------|---------|-------------|
| `REDIS_NODES` | `10.0.0.1,10.0.0.2,10.0.0.3` | Comma-separated container IPs. First = master, rest = replicas |
| `REDIS_REPLICATION_KEY` | `secret-key` | Shared auth password for replication |
| `REDIS_SENTINEL_PORT` | `26379` | Sentinel port for failover |
| `REDIS_MASTER_NAME` | `hosthub` | Master name for Sentinel |

**Replication modes:**
- **Standalone (default)**: `REDIS_NODES` empty, `REDIS_REPLICATION_KEY` empty → No auth, single instance
- **Standalone with auth**: `REDIS_NODES` empty, `REDIS_REPLICATION_KEY` set → Auth enabled, single instance
- **Replicated**: `REDIS_NODES` set → Master/replica with Sentinel failover

### Build-time Variables

These can be set but are **optional**:

| Variable | Purpose |
|----------|---------|
| `HOST_IS_BUILDING=true` | Indicates build phase (for debugging) |

**Note:** The `APP_ENV=production` warning from Coolify can be ignored. We hardcode `--no-dev` in the Dockerfile, so build-time APP_ENV doesn't affect dependency installation.

### Hades Mode (God-Mode Debug Access)

For production debugging without exposing errors to users:

| Variable | Example | Description |
|----------|---------|-------------|
| `HADES_TOKEN` | `my-secret-debug-token-2024` | Enables debug mode for cookie holders |

**How it works:**
1. Set `HADES_TOKEN` to any secret string in Coolify
2. Log in to the app - a `hades` cookie is set (encrypted, 1 year lifetime)
3. When errors occur, you see the full debug page instead of 503
4. Other users without the cookie see the friendly 503 page

**To revoke access:**
- Change `HADES_TOKEN` to a different value
- All existing cookies become invalid immediately

**Security notes:**
- Cookie is HTTP-only and encrypted with APP_KEY
- Only works if you've logged in after HADES_TOKEN was set
- Completely disabled if HADES_TOKEN is empty/unset

## Startup Sequence

When the container starts, the entrypoint script runs:

1. **Banner** - Shows environment info (APP_ENV, DB_HOST, REDIS_HOST)
2. **WordPress Setup** - Copies core files if first run, applies patches
3. **Laravel Setup** - Creates storage directories, sets permissions
4. **Migrations** - Runs `php artisan migrate --force`
5. **Config Cache** - Caches config/routes/views (production only)
6. **Supervisord** - Starts nginx, php-fpm, queue workers, scheduler

### Startup Logs

You'll see this on startup:
```
============================================
  Host UK Unified Container
============================================

  Environment: production
  Debug:       false
  DB Host:     mariadb
  Redis Host:  redis

  Laravel:     host.uk.com
  WordPress:   *.host.uk.com
============================================
```

## Troubleshooting

### Health Check Failing

If health checks fail, check in order:

1. **Is nginx running?**
   ```bash
   docker exec <container> nginx -t
   ```

2. **Can you reach /healthz internally?**
   ```bash
   docker exec <container> curl -s http://localhost/healthz
   ```

3. **Check supervisord status:**
   ```bash
   docker exec <container> supervisorctl status
   ```

### Queue Workers Crashing

Common causes:
- **Permission denied on logs**: Fixed by `chmod -R 775 storage`
- **SQLite driver missing**: Fixed by adding `php84-pdo_sqlite`
- **Redis unavailable**: Check `REDIS_HOST` environment variable

### 503 Errors

Laravel returning 503 usually means:
- Database connection failed - check `DB_HOST`, credentials
- Redis connection failed - check `REDIS_HOST`
- Config cached with wrong values - clear cache and redeploy

To debug, check Laravel logs:
```bash
docker exec <container> tail -50 /app/storage/logs/laravel.log
```

### Permission Errors

The entrypoint sets permissions at startup:
```bash
chown -R nobody:nobody storage bootstrap/cache
chmod -R 775 storage bootstrap/cache
```

If issues persist, check that storage isn't a read-only volume.

## Updating Paid Libraries

Flux and Mixpost are vendored locally to avoid auth requirements during CI/CD builds.

To update:
```bash
make update-paid-libs
```

This will:
1. Temporarily enable remote Flux repository
2. Run composer update
3. Copy packages to `product/host-hub/packages/`
4. Restore local-only composer.json
5. Reinstall from local packages

Then commit:
```bash
git add product/host-hub/packages/
git commit -m "Update Flux to vX.X.X"
```

## Local Development

```bash
# Start all services
docker-compose -f docker-compose.dev.yml up -d

# Check status
docker-compose -f docker-compose.dev.yml ps

# View logs
docker-compose -f docker-compose.dev.yml logs -f app

# Shell into container
docker-compose -f docker-compose.dev.yml exec app sh

# Rebuild after changes
docker-compose -f docker-compose.dev.yml build app
docker-compose -f docker-compose.dev.yml up -d app
```

Access locally:
- http://host.uk.com → Laravel (requires /etc/hosts entry)
- http://hestia.host.uk.com → WordPress (requires /etc/hosts entry)

## Domain Routing

Nginx routes by `Host` header:

| Domain | Destination |
|--------|-------------|
| `host.uk.com` | Laravel `/app/public` |
| `www.host.uk.com` | Laravel `/app/public` |
| `*.host.uk.com` | WordPress `/var/www/html` |

## File Locations

| Path | Contents |
|------|----------|
| `/app` | Laravel Host Hub |
| `/app/storage/logs` | Laravel logs |
| `/var/www/html` | WordPress (runtime) |
| `/usr/src/wordpress` | WordPress (source) |
| `/usr/src/wordpress-patch` | Custom themes/plugins |
| `/etc/nginx/conf.d/default.conf` | Nginx config |
| `/etc/supervisor/conf.d/supervisord.conf` | Process manager config |
