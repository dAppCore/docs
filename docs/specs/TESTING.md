# Testing Guide

## Overview

Host Hub uses [Pest PHP](https://pestphp.com) for unit and feature testing, with Playwright for browser tests.

**Current Coverage:** ~23% (Target: 80%+)

## Quick Start

```bash
# Run all tests
./vendor/bin/pest

# Run specific test file
./vendor/bin/pest tests/Feature/BioLink/BioLinkTest.php

# Run with coverage report
make coverage
# or
export XDEBUG_MODE=coverage
./vendor/bin/pest --coverage --min=0

# View coverage report
open coverage/html/index.html
```

## Test Structure

```
tests/
├── Unit/               # Pure unit tests (no database, fast)
│   ├── HadesEncryptTest.php
│   └── Services/
│       └── BunnyCdnServiceTest.php
├── Feature/            # Integration tests (database, HTTP)
│   ├── Analytics/
│   ├── BioLink/
│   ├── Social/
│   ├── Support/
│   └── ...
└── Browser/            # Playwright browser tests
    ├── SmokeTest.php
    └── MarketingPagesTest.php
```

## Writing Tests

### Basic Test

```php
<?php

use Mod\Tenant\Models\User;
use Mod\Tenant\Models\Workspace;

test('user can create workspace', function () {
    $user = User::factory()->create();

    actingAs($user)
        ->post('/workspaces', [
            'name' => 'Test Workspace',
        ])
        ->assertRedirect('/workspaces');

    expect(Workspace::count())->toBe(1);
});
```

### Using Datasets

```php
dataset('cases', [
    'lowercase' => ['hello', 'HELLO'],
    'uppercase' => ['HELLO', 'hello'],
    'mixed' => ['HeLLo', 'HeLLo'],
]);

test('converts case correctly', function ($input, $expected) {
    $result = convertCase($input);
    expect($result)->toBe($expected);
})->with('cases');
```

### Grouped Tests (Describe)

```php
describe('BioLink Analytics', function () {
    beforeEach(function () {
        $this->bioLink = BioLink::factory()->create();
    });

    test('tracks pageview');
    test('tracks click');
    test('calculates conversion rate');
});
```

### Testing Livewire Components

```php
use Livewire\Livewire;
use App\Livewire\Admin\Analytics\Dashboard;

test('analytics dashboard loads', function () {
    $workspace = Workspace::factory()->create();

    Livewire::actingAs($workspace->owner)
        ->test(Dashboard::class)
        ->assertOk()
        ->assertSee('Analytics Dashboard');
});
```

## Coverage Reports

### Generate Coverage

```bash
# Quick: Makefile command
make coverage

# Manual: with Xdebug
export XDEBUG_MODE=coverage
./vendor/bin/pest --coverage --min=80

# Parallel (faster, but less accurate coverage)
./vendor/bin/pest --parallel --coverage
```

### View Coverage

**HTML Report (Best):**
```bash
open coverage/html/index.html
```

**Terminal Output:**
```bash
./vendor/bin/pest --coverage --min=0 | grep -A 50 "Code Coverage"
```

**Clover XML (CI/CD):**
```xml
<!-- coverage/clover.xml -->
```

## Current Coverage Gaps

See [TASK-016-TEST-COVERAGE-IMPROVEMENT.md](../released/jan/TASK-016-TEST-COVERAGE-IMPROVEMENT.md) for detailed improvement plan.

### Critical Gaps (Priority Order)

1. **Tools Services** (2% coverage)
   - 42 utility tools with minimal tests
   - Example: `tests/Feature/Services/Tools/TextToolsTest.php`

2. **Analytics Models** (30% coverage)
   - Goal, GoalConversion, Heatmap, SessionReplay
   - Missing: conversion tracking, aggregation logic

3. **Support Services** (60% coverage)
   - Missing: EmailParserService, SearchService

4. **Analytics Services** (50% coverage)
   - Missing: GeoIpService, HeatmapAggregationService

## Test Categories

### Unit Tests

Pure unit tests with no external dependencies:
```php
// tests/Unit/Services/BunnyCdnServiceTest.php
test('BunnyCdnService reports configured when api key present', function () {
    config(['services.bunnycdn.api_key' => 'test-key']);
    config(['services.bunnycdn.pull_zone_id' => '12345']);

    $service = app(BunnyCdnService::class);

    expect($service->isConfigured())->toBeTrue();
});
```

### Feature Tests

Integration tests with database, HTTP, queues:
```php
// tests/Feature/Analytics/PageviewProcessingTest.php
test('pageview creates session and visitor', function () {
    $website = Website::factory()->create();

    $response = $this->postJson('/api/v1/analytics/track', [
        'website_id' => $website->id,
        'url' => 'https://example.com/page',
        'referrer' => 'https://google.com',
    ]);

    $response->assertOk();
    expect(AnalyticsSession::count())->toBe(1)
        ->and(AnalyticsVisitor::count())->toBe(1);
});
```

### Browser Tests (Playwright)

End-to-end tests in real browsers:
```typescript
// tests/Browser/SmokeTest.php
test('homepage loads', async ({ page }) => {
  await page.goto('https://host.uk.com');
  await expect(page).toHaveTitle(/Host UK/);
});
```

## Best Practices

### DO

✅ Use factories for test data
```php
$user = User::factory()->create();
$workspace = Workspace::factory()->create();
```

✅ Test one thing per test
```php
test('validates email format');
test('validates email uniqueness');
// NOT: test('validates email')
```

✅ Use descriptive test names
```php
test('user cannot delete workspace with active subscription');
```

✅ Mock external services
```php
Http::fake([
    'api.twitter.com/*' => Http::response(['status' => 'ok'], 200),
]);
```

✅ Clean up in beforeEach/afterEach
```php
beforeEach(fn() => $this->seed());
afterEach(fn() => Cache::flush());
```

### DON'T

❌ Share state between tests
```php
// BAD
$this->user = User::factory()->create(); // in beforeEach
test('deletes user', function() {
    $this->user->delete(); // breaks next test!
});
```

❌ Test framework behaviour
```php
// BAD - Laravel already tests this
test('user model has email attribute');
```

❌ Skip arranging test data
```php
// BAD
test('creates post', function() {
    $this->post('/posts', []); // what data?
});

// GOOD
test('creates post', function() {
    $data = ['title' => 'Test', 'content' => 'Content'];
    $this->post('/posts', $data)->assertCreated();
});
```

## CI/CD Integration

### GitHub Actions (Future)

```yaml
name: Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: shivammathur/setup-php@v2
        with:
          php-version: 8.5
          extensions: xdebug
          coverage: xdebug

      - name: Install dependencies
        run: composer install --no-interaction

      - name: Run tests with coverage
        run: |
          export XDEBUG_MODE=coverage
          ./vendor/bin/pest --coverage --min=80

      - name: Upload coverage
        uses: codecov/codecov-action@v3
        with:
          files: ./coverage/clover.xml
```

## Debugging Tests

### Run specific test
```bash
./vendor/bin/pest --filter="user can create workspace"
```

### Stop on first failure
```bash
./vendor/bin/pest --stop-on-failure
```

### Show output during tests
```php
test('debug test', function () {
    dump($someVariable); // Shows in terminal
    ray($someVariable);  // Shows in Ray app
});
```

### Use Ray for debugging
```bash
composer require spatie/laravel-ray --dev
```

```php
ray($user)->blue();
ray()->table($data);
```

## Performance

### Parallel Execution

```bash
# Faster test runs (requires ParaTest)
./vendor/bin/pest --parallel
```

### Skip Slow Tests

```php
test('slow integration test', function () {
    // ...
})->group('slow');

// Run without slow tests
./vendor/bin/pest --exclude-group=slow
```

## Resources

- **Pest Docs:** https://pestphp.com
- **Pest Expectations:** https://pestphp.com/docs/expectations
- **Laravel Testing:** https://laravel.com/docs/testing
- **Playwright:** https://playwright.dev
- **Coverage Plan:** [TASK-016](../released/jan/TASK-016-TEST-COVERAGE-IMPROVEMENT.md)
- **Example Test:** `tests/Feature/Services/Tools/TextToolsTest.php`

## Getting Help

Run into issues? Check:
1. Test logs: `./vendor/bin/pest --verbose`
2. Laravel logs: `storage/logs/laravel.log`
3. Coverage report: `open coverage/html/index.html`
4. This guide: `doc/TESTING.md`

---

**Updated:** 4 Jan 2026
