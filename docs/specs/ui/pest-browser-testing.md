# Pest PHP Testing Framework Documentation

## Introduction

Pest is an elegant and sophisticated testing framework for PHP that prioritizes developer experience and code readability. Built on top of PHPUnit, Pest provides a modern, expressive syntax that resembles natural human language, making tests easier to write, read, and maintain. The framework transforms traditional PHPUnit class-based testing into a functional, concise style using global functions like `test()`, `it()`, and `expect()`.

The framework offers comprehensive features including parallel test execution, browser testing capabilities, architectural testing to enforce code structure rules, mutation testing for test suite quality evaluation, and extensive dataset support for parameterized testing. With its beautiful terminal output, intuitive expectation API, and rich plugin ecosystem, Pest is designed to make testing enjoyable while maintaining compatibility with PHPUnit's robust foundation. Whether building small personal projects or large enterprise applications, Pest provides the tools needed for effective test-driven development.

## API Reference and Examples

### Installation and Setup

Initialize Pest in your PHP project by installing it via Composer and configuring your test suite.

```bash
# Remove existing PHPUnit installation
composer remove phpunit/phpunit

# Install Pest as a development dependency
composer require pestphp/pest --dev --with-all-dependencies

# Initialize Pest configuration
./vendor/bin/pest --init

# Run your test suite
./vendor/bin/pest
```

### Basic Test Definition with test() and it()

Write simple, expressive tests using the `test()` or `it()` function with closures.

```php
<?php

// Using test() function - simple description
test('sum', function () {
    $result = sum(1, 2);

    expect($result)->toBe(3);
});

// Using it() function - automatically prefixes with "it"
it('performs sums', function () {
    $result = sum(1, 2);

    expect($result)->toBe(3);
});

// Using describe() to group related tests
describe('sum', function () {
    it('may sum integers', function () {
        $result = sum(1, 2);

        expect($result)->toBe(3);
    });

    it('may sum floats', function () {
        $result = sum(1.5, 2.5);

        expect($result)->toBe(4.0);
    });
});
```

### Expectation API - Type Checking

Verify data types and values using Pest's fluent expectation API.

```php
<?php

test('type checking expectations', function () {
    // Strict equality (same type and value)
    expect(1)->toBe(1);
    expect('1')->not->toBe(1);

    // Type assertions
    expect(123)->toBeInt();
    expect(3.14)->toBeFloat();
    expect('hello')->toBeString();
    expect(true)->toBeBool();
    expect(['a', 'b'])->toBeArray();
    expect(new stdClass())->toBeObject();

    // Loose equality (same value, different type allowed)
    expect('1')->toEqual(1);
    expect(new StdClass())->toEqual(new StdClass());

    // Numeric checks
    expect('10')->toBeNumeric();
    expect(10)->toBeDigits();
    expect('0.123')->not->toBeDigits();

    // Empty and null checks
    expect('')->toBeEmpty();
    expect([])->toBeEmpty();
    expect(null)->toBeNull();
});
```

### Expectation API - Collections and Arrays

Work with arrays, collections, and iterables using specialized expectations.

```php
<?php

test('array and collection expectations', function () {
    $users = ['id' => 1, 'name' => 'Nuno', 'email' => 'enunomaduro@gmail.com'];

    // Count and length
    expect(['Nuno', 'Luke', 'Alex', 'Dan'])->toHaveCount(4);
    expect('Pest')->toHaveLength(4);
    expect(['Nuno', 'Maduro'])->toHaveLength(2);

    // Key existence
    expect($users)->toHaveKey('name');
    expect($users)->toHaveKey('name', 'Nuno');
    expect(['id' => 1, 'name' => 'Nuno'])->toHaveKeys(['id', 'name']);

    // Nested key access with dot notation
    expect(['user' => ['name' => 'Nuno']])->toHaveKey('user.name');
    expect(['user' => ['name' => 'Nuno']])->toHaveKey('user.name', 'Nuno');

    // Contains checks
    expect('Hello World')->toContain('Hello');
    expect([1, 2, 3, 4])->toContain(2, 4);
    expect([1, 2, 3])->toContainEqual('1'); // Loose equality

    // Match subset
    expect($users)->toMatchArray([
        'email' => 'enunomaduro@gmail.com',
        'name' => 'Nuno'
    ]);

    // Each modifier - apply expectation to all items
    expect([1, 2, 3])->each->toBeInt();
    expect([1, 2, 3])->each(fn ($number) => $number->toBeLessThan(4));
});
```

### Expectation API - String Operations

Validate string content, format, and patterns with string-specific expectations.

```php
<?php

test('string expectations', function () {
    // Start and end checks
    expect('Hello World')->toStartWith('Hello');
    expect('Hello World')->toEndWith('World');

    // Pattern matching
    expect('Hello World')->toMatch('/^hello wo.*$/i');

    // Case validation
    expect('PESTPHP')->toBeUppercase();
    expect('pestphp')->toBeLowercase();

    // Character type checks
    expect('pestphp')->toBeAlpha();
    expect('pestPHP123')->toBeAlphaNumeric();

    // Naming convention validation
    expect('snake_case')->toBeSnakeCase();
    expect('kebab-case')->toBeKebabCase();
    expect('camelCase')->toBeCamelCase();
    expect('StudlyCase')->toBeStudlyCase();

    // Array keys naming conventions
    expect(['snake_case' => 'abc123'])->toHaveSnakeCaseKeys();
    expect(['camelCase' => 'abc123'])->toHaveCamelCaseKeys();

    // Format validation
    expect('{"hello":"world"}')->toBeJson();
    expect('https://pestphp.com/')->toBeUrl();
    expect('ca0a8228-cdf6-41db-b34b-c2f31485796c')->toBeUuid();
});
```

### Expectation API - Comparison and Range

Perform numeric comparisons and range validations.

```php
<?php

test('comparison expectations', function () {
    $count = 25;

    // Greater than comparisons
    expect($count)->toBeGreaterThan(20);
    expect($count)->toBeGreaterThanOrEqual(25);

    // Less than comparisons
    expect($count)->toBeLessThan(30);
    expect($count)->toBeLessThanOrEqual(25);

    // Between range (works with int, float, DateTime)
    expect(2)->toBeBetween(1, 3);
    expect(1.5)->toBeBetween(1, 2);

    $date = new DateTime('2023-09-22');
    $oldest = new DateTime('2023-09-21');
    $latest = new DateTime('2023-09-23');
    expect($date)->toBeBetween($oldest, $latest);

    // In set validation
    expect($newUser->status)->toBeIn(['pending', 'new', 'active']);

    // Delta comparison (useful for floats)
    expect(14)->toEqualWithDelta(10, 5); // Pass: difference is 4
    expect(14)->not->toEqualWithDelta(10, 0.1); // Fail: difference > 0.1
});
```

### Expectation API - Object and Class Testing

Test object properties, instances, and class relationships.

```php
<?php

test('object and class expectations', function () {
    $user = new User();
    $user->name = 'Nuno';
    $user->email = 'enunomaduro@gmail.com';

    // Instance checking
    expect($user)->toBeInstanceOf(User::class);
    expect($user)->toBeObject();

    // Property existence and values
    expect($user)->toHaveProperty('name');
    expect($user)->toHaveProperty('name', 'Nuno');
    expect($user)->toHaveProperties(['name', 'email']);
    expect($user)->toHaveProperties([
        'name' => 'Nuno',
        'email' => 'enunomaduro@gmail.com'
    ]);

    // Match object subset
    expect($user)->toMatchObject([
        'email' => 'enunomaduro@gmail.com',
        'name' => 'Nuno'
    ]);

    // Array of instances
    $dates = [new DateTime(), new DateTime()];
    expect($dates)->toContainOnlyInstancesOf(DateTime::class);

    // Callable and resource checks
    $myFunction = function () {};
    expect($myFunction)->toBeCallable();

    $handle = fopen('php://memory', 'r+');
    expect($handle)->toBeResource();
});
```

### Expectation API - Exception Testing

Verify that code throws expected exceptions with proper messages.

```php
<?php

test('exception expectations', function () {
    // Expect specific exception class
    expect(fn() => throw new Exception('Something happened.'))
        ->toThrow(Exception::class);

    // Expect exception message
    expect(fn() => throw new Exception('Something happened.'))
        ->toThrow('Something happened.');

    // Expect both class and message
    expect(fn() => throw new Exception('Something happened.'))
        ->toThrow(Exception::class, 'Something happened.');

    // Expect specific exception instance
    expect(fn() => throw new Exception('Something happened.'))
        ->toThrow(new Exception('Something happened.'));

    // Combined with other expectations
    expect(fn() => throw new InvalidArgumentException('Invalid input'))
        ->toThrow(InvalidArgumentException::class);
});
```

### Expectation API - Advanced Modifiers

Chain multiple expectations and use conditional logic with modifiers.

```php
<?php

test('expectation modifiers', function () {
    // and() - test multiple values
    $id = 14;
    $name = 'Nuno';
    expect($id)->toBe(14)
        ->and($name)->toBe('Nuno');

    // each() - apply expectation to all items
    expect([1, 2, 3])->each->toBeInt();
    expect([1, 2, 3])->each->not->toBeString();
    expect([1, 2, 3])->each(fn ($number, $key) => $number->toEqual($key + 1));

    // sequence() - ordered expectations for iterable
    expect([1, 2, 3])->sequence(
        fn ($number) => $number->toBe(1),
        fn ($number) => $number->toBe(2),
        fn ($number) => $number->toBe(3),
    );

    // Shorthand sequence
    expect(['foo', 'bar', 'baz'])->sequence('foo', 'bar', 'baz');

    // json() - decode JSON and continue chaining
    expect('{"name":"Nuno","credit":1000.00}')
        ->json()
        ->toHaveCount(2)
        ->name->toBe('Nuno')
        ->credit->toBeFloat();

    // when() - conditional expectations
    expect($user)
        ->when($user->is_verified === true,
            fn ($user) => $user->daily_limit->toBeGreaterThan(10))
        ->email->not->toBeEmpty();

    // unless() - inverse conditional
    expect($user)
        ->unless($user->is_verified === true,
            fn ($user) => $user->daily_limit->toBe(10))
        ->email->not->toBeEmpty();

    // match() - pattern matching on value
    expect($user->miles)
        ->match($user->status, [
            'new' => fn ($miles) => $miles->toBe(0),
            'gold' => fn ($miles) => $miles->toBeGreaterThan(500),
            'platinum' => fn ($miles) => $miles->toBeGreaterThan(1000),
        ]);
});
```

### Hooks - Test Setup and Teardown

Execute code before and after tests using lifecycle hooks.

```php
<?php

// File: tests/Feature/UserTest.php

beforeEach(function () {
    // Runs before every test in this file
    $this->userRepository = new UserRepository();
    $this->userRepository->truncate();
});

afterEach(function () {
    // Runs after every test in this file
    $this->userRepository->reset();
});

beforeAll(function () {
    // Runs once before any test in this file
    // Note: $this is not available here
    DB::beginTransaction();
});

afterAll(function () {
    // Runs once after all tests in this file
    // Note: $this is not available here
    DB::rollBack();
});

it('may create a user', function () {
    $user = $this->userRepository->create(['name' => 'John']);

    expect($user)->toBeInstanceOf(User::class);
});

// Per-test cleanup with after() method
it('may delete a user', function () {
    $user = $this->userRepository->create(['name' => 'Jane']);
    $this->userRepository->delete($user->id);

    expect($this->userRepository->find($user->id))->toBeNull();
})->after(function () {
    // Cleanup specific to this test only
    $this->userRepository->cleanup();
});

// Nested hooks with describe()
describe('user authentication', function () {
    beforeEach(function () {
        // Runs only for tests in this describe block
        $this->authService = new AuthService();
    });

    it('may authenticate valid credentials', function () {
        $result = $this->authService->login('user@example.com', 'password');
        expect($result)->toBeTrue();
    });
});
```

### Datasets - Parameterized Testing

Run the same test with multiple input variations using datasets.

```php
<?php

// Inline datasets - simple arrays
it('has emails', function (string $email) {
    expect($email)->not->toBeEmpty();
})->with(['enunomaduro@gmail.com', 'other@example.com']);

// Multiple arguments
it('can sum numbers', function (int $a, int $b, int $expected) {
    expect(sum($a, $b))->toBe($expected);
})->with([
    [1, 2, 3],
    [5, 5, 10],
    [-1, 1, 0],
]);

// Named datasets for better output
it('validates email format', function (string $email) {
    expect($email)->toMatch('/@/');
})->with([
    'james' => 'james@laravel.com',
    'taylor' => 'taylor@laravel.com',
]);

// Datasets with closures for computed values
test('numbers are integers', function ($i) {
    expect($i)->toBeInt();
})->with(fn (): array => range(1, 99));

// Generator for large datasets (memory efficient)
test('large dataset test', function ($i) {
    expect($i)->toBeInt();
})->with(function (): Generator {
    for ($i = 1; $i < 100_000_000; $i++) {
        yield $i;
    }
});

// Shared datasets - stored in tests/Datasets/Emails.php
dataset('emails', [
    'enunomaduro@gmail.com',
    'other@example.com'
]);

// Using shared dataset in tests
it('has emails', function (string $email) {
    expect($email)->not->toBeEmpty();
})->with('emails');

// Combining datasets (cartesian product)
dataset('days_of_the_week', ['Saturday', 'Sunday']);

test('business is closed on day', function(string $business, string $day) {
    expect(new $business)->isClosed($day)->toBeTrue();
})->with([
    Office::class,
    Bank::class,
    School::class
])->with('days_of_the_week');

// Bound datasets - executed after beforeEach()
it('can generate full name', function (User $user) {
    expect($user->full_name)->toBe("{$user->first_name} {$user->last_name}");
})->with([
    fn() => User::factory()->create(['first_name' => 'Nuno', 'last_name' => 'Maduro']),
    fn() => User::factory()->create(['first_name' => 'Luke', 'last_name' => 'Downing']),
]);

// Repeat tests for stability checking
it('random operation is stable', function () {
    $result = performRandomOperation();
    expect($result)->toBeTrue();
})->repeat(100);
```

### Configuration - Base Test Classes and Traits

Configure test suite behavior using Pest.php configuration file.

```php
<?php

// File: tests/Pest.php

use Tests\TestCase;
use Illuminate\Foundation\Testing\RefreshDatabase;

// Apply TestCase to all Feature tests
pest()->extend(TestCase::class)->in('Feature');

// Apply multiple traits
pest()
    ->extend(TestCase::class)
    ->use(RefreshDatabase::class)
    ->in('Feature');

// Glob patterns for selective application
pest()->extend(TestCase::class)->in('Feature/*Job*.php');

// Multiple directories with pattern matching
pest()
    ->extend(DuskTestCase::class)
    ->use(DatabaseMigrations::class)
    ->in('../Modules/*/Tests/Browser');

// Per-file configuration (in test file itself)
pest()->extend(Tests\MySpecificTestCase::class);

it('uses custom test case', function () {
    echo get_class($this); // \Tests\MySpecificTestCase

    // Access public methods from base test class
    $this->performCustomSetup();
});

// Example TestCase with helper methods
// File: tests/TestCase.php
class TestCase extends PHPUnit\Framework\TestCase
{
    protected function createUser(array $attributes = []): User
    {
        return User::factory()->create($attributes);
    }

    protected function actingAs(User $user): self
    {
        $this->user = $user;
        return $this;
    }
}
```

### Mocking with Mockery

Create test doubles to isolate code and control dependencies.

```php
<?php

use App\Repositories\BookRepository;
use App\Services\PaymentClient;
use Mockery;

test('may purchase book without actual payment', function () {
    // Create mock object
    $client = Mockery::mock(PaymentClient::class);

    // Set method expectations
    $client->shouldReceive('post')
        ->with('/api/payments', Mockery::type('array'))
        ->once()
        ->andReturn(['status' => 'success', 'transaction_id' => '12345']);

    $client->shouldReceive('getStatus')
        ->with('12345')
        ->andReturn('completed');

    // Use mock in code under test
    $repository = new BookRepository($client);
    $result = $repository->purchase('PHP Book', 29.99);

    expect($result)->toBe('completed');
});

// Argument matchers
test('mock with flexible arguments', function () {
    $mock = Mockery::mock(ApiClient::class);

    // Match any argument
    $mock->shouldReceive('post')->with(Mockery::any());

    // Match specific types
    $mock->shouldReceive('send')->with(Mockery::type('string'));

    // Custom argument validation with closure
    $mock->shouldReceive('process')
        ->withArgs(function ($arg) {
            return $arg > 0 && $arg < 100;
        })
        ->andReturn('valid');

    expect($mock->process(50))->toBe('valid');
});

// Return value sequences
test('mock with changing return values', function () {
    $mock = Mockery::mock(DataService::class);

    // Return different values on successive calls
    $mock->shouldReceive('fetch')->andReturn(1, 2, 3);

    expect($mock->fetch())->toBe(1);
    expect($mock->fetch())->toBe(2);
    expect($mock->fetch())->toBe(3);
});

// Computed return values
test('mock with computed returns', function () {
    $mock = Mockery::mock(Calculator::class);

    $mock->shouldReceive('multiply')
        ->andReturnUsing(fn ($a, $b) => $a * $b);

    expect($mock->multiply(3, 4))->toBe(12);
});

// Exception throwing
test('mock throws exception', function () {
    $mock = Mockery::mock(FileService::class);

    $mock->shouldReceive('read')
        ->with('missing.txt')
        ->andThrow(new FileNotFoundException('File not found'));

    expect(fn() => $mock->read('missing.txt'))
        ->toThrow(FileNotFoundException::class);
});

// Call count expectations
test('verify method call counts', function () {
    $mock = Mockery::mock(Logger::class);

    $mock->shouldReceive('log')->once();
    $mock->shouldReceive('warn')->twice();
    $mock->shouldReceive('error')->times(3);
    $mock->shouldReceive('debug')->atLeast()->times(2);
    $mock->shouldReceive('info')->atMost()->times(5);

    // Execute code that should trigger these calls
    $service = new MyService($mock);
    $service->process();
});
```

### Architecture Testing - Enforce Code Structure Rules

Define and validate architectural boundaries and coding standards.

```php
<?php

// File: tests/Architecture/ArchTest.php

// Namespace-level rules
arch()
    ->expect('App')
    ->toUseStrictTypes()
    ->not->toUse(['die', 'dd', 'dump']);

// Model layer constraints
arch()
    ->expect('App\Models')
    ->toBeClasses()
    ->toExtend('Illuminate\Database\Eloquent\Model')
    ->toOnlyBeUsedIn('App\Repositories')
    ->ignoring('App\Models\User');

// Layer isolation
arch()
    ->expect('App\Http')
    ->toOnlyBeUsedIn('App\Http');

arch()
    ->expect('App\Domain')
    ->not->toUse(['App\Http', 'App\Infrastructure']);

// File organization patterns
arch()
    ->expect('App\*\Traits')
    ->toBeTraits();

arch()
    ->expect('App\*\Interfaces')
    ->toBeInterfaces();

// Class structure rules
arch()
    ->expect('App\Http\Controllers')
    ->toHaveSuffix('Controller')
    ->toHaveMethod('__invoke')
    ->toExtend('App\Http\Controllers\Controller');

// Method visibility constraints
arch()
    ->expect('App\Services')
    ->toHavePublicMethodsBesides(['__construct', '__destruct']);

// Documentation requirements
arch()
    ->expect('App\Contracts')
    ->toHaveMethodsDocumented()
    ->toHavePropertiesDocumented();

// Naming conventions
arch()
    ->expect('App\Events')
    ->toHaveSuffix('Event')
    ->toHaveAttribute('Illuminate\Foundation\Events\Dispatchable');

// File permissions and security
arch()
    ->expect('App')
    ->toHaveFileSystemPermissions(0644)
    ->not->toHaveSuspiciousCharacters();

// Code quality rules
arch()
    ->expect('App\Services')
    ->toHaveLineCountLessThan(200)
    ->toUseStrictEquality();

// Preset rules for common patterns
arch()->preset()->php(); // PHP best practices
arch()->preset()->security()->ignoring('md5'); // Security rules

// Custom rule with specific files
arch('controllers must be invokable')
    ->expect('App\Http\Controllers')
    ->toBeInvokable()
    ->toExtend('Illuminate\Routing\Controller');
```

### Browser Testing - Configuration

Configure browser testing behavior in `tests/Pest.php`.

```php
<?php

// File: tests/Pest.php

// Global host configuration for subdomain/domain-based routing
// Useful for multi-tenant apps, workspace-based apps, Laravel Sail
pest()->browser()->withHost('host.test');

// Browser type configuration
pest()->browser()->inChrome();   // Default
pest()->browser()->inFirefox();
pest()->browser()->inSafari();

// Theme configuration
pest()->browser()->inLightMode();
pest()->browser()->inDarkMode();

// Headed mode (visible browser window for debugging)
pest()->browser()->headed();

// Assertion timeout in milliseconds
pest()->browser()->timeout(10000);
```

Per-test host configuration is also available:

```php
<?php

// Override host for a specific test
it('tests subdomain routing', function () {
    $page = visit('/dashboard')->withHost('tenant1.host.test');

    $page->assertSee('Tenant 1 Dashboard');
});
```

### Browser Testing - E2E Testing

Test your application in real browsers with automated interactions.

```php
<?php

// File: tests/Browser/AuthenticationTest.php

use App\Models\User;
use Illuminate\Support\Facades\Event;

it('may visit homepage', function () {
    $page = visit('/');

    $page->assertSee('Welcome')
         ->assertSee('Get Started');
});

it('may sign in user', function () {
    Event::fake();

    // Setup test data
    User::factory()->create([
        'email' => 'nuno@laravel.com',
        'password' => bcrypt('password'),
    ]);

    // Perform browser test with specific browser and device
    $page = visit('/')->on()->mobile()->firefox();

    $page->click('Sign In')
         ->assertUrlIs('/login')
         ->assertSee('Sign In to Your Account')
         ->fill('email', 'nuno@laravel.com')
         ->fill('password', 'password')
         ->click('Submit')
         ->assertSee('Dashboard')
         ->assertDontSee('Sign In');

    // Laravel-specific assertions still work
    $this->assertAuthenticated();
    Event::assertDispatched(UserLoggedIn::class);
});

it('displays validation errors', function () {
    $page = visit('/register');

    $page->fill('name', '')
         ->fill('email', 'invalid-email')
         ->click('Register')
         ->assertSee('The name field is required')
         ->assertSee('The email must be a valid email address');
});

it('navigates between pages', function () {
    $page = visit('/');

    $page->click('About')
         ->assertUrlIs('/about')
         ->assertSee('About Us')
         ->click('Contact')
         ->assertUrlIs('/contact')
         ->assertSee('Contact Form');
});

it('handles javascript interactions', function () {
    $page = visit('/dashboard');

    $page->click('#dropdown-menu')
         ->waitFor('.dropdown-items')
         ->click('.logout-button')
         ->assertUrlIs('/');
});
```

## Summary and Integration Patterns

Pest PHP revolutionizes PHP testing by combining PHPUnit's robust foundation with an elegant, modern syntax that emphasizes readability and developer experience. The framework's core strength lies in its functional approach using `test()` and `it()` functions paired with the fluent `expect()` API, eliminating verbose class structures while maintaining full PHPUnit compatibility. Its extensive feature set includes parameterized testing through datasets, lifecycle management via hooks, powerful mocking integration with Mockery, architectural rule enforcement, and real browser testing capabilities—all orchestrated through simple, chainable function calls.

Integration patterns with Pest focus on flexibility and gradual adoption. Teams can configure test suites via the `Pest.php` file to extend base test classes and apply traits selectively to different directories using glob patterns, enabling framework-specific features like Laravel's `RefreshDatabase` trait. The framework excels in CI/CD pipelines with built-in parallel execution support, supports mixed PHPUnit/Pest codebases for incremental migration, and offers extensive customization through plugins, custom expectations, and global hooks. Whether testing standalone PHP libraries, Laravel applications, or complex enterprise systems, Pest's consistent API and beautiful terminal output create a unified, enjoyable testing experience that encourages comprehensive test coverage and test-driven development practices.
