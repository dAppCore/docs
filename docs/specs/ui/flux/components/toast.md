# flux:toast

Temporary feedback messages with auto-dismiss and stacking.

## Setup

Add to your layout (typically in `layouts/app.blade.php`):

```blade
<body>
    {{ $slot }}
    <flux:toast />
</body>
```

For persistence with `wire:navigate`:

```blade
@persist('toast')
    <flux:toast />
@endpersist
```

## Props

### flux:toast

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `position` | string | bottom end | Position on screen |
| `class` | string | - | Additional CSS (e.g., `pt-24`) |

### flux:toast.group

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `position` | string | bottom end | Stack position |
| `expanded` | boolean | false | Always show expanded |

## Position Options

- `bottom end` (default)
- `bottom center`
- `bottom start`
- `top end`
- `top center`
- `top start`

## Triggering Toasts

### From Livewire (PHP)

```php
use Flux\Flux;

public function save()
{
    // Save logic...
    Flux::toast('Your changes have been saved.');
}
```

### From Alpine.js

```blade
<button x-on:click="$flux.toast('Changes saved.')">
    Save
</button>
```

### From JavaScript

```javascript
window.Flux.toast('Your changes have been saved.')
```

## Toast Options

```php
Flux::toast(
    heading: 'Success',
    text: 'Your changes have been saved.',
    variant: 'success',
    duration: 5000,
);
```

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `heading` | string | - | Optional title |
| `text` | string | - | Message content |
| `variant` | string | - | `success`, `warning`, `danger` |
| `duration` | int | 5000 | Auto-dismiss ms (0 = permanent) |

## Simple Messages

```php
Flux::toast('Changes saved.');
```

## With Heading

```php
Flux::toast(
    heading: 'Profile updated',
    text: 'Your profile changes have been saved.',
);
```

## Variants

```php
// Success (green)
Flux::toast(text: 'File uploaded successfully.', variant: 'success');

// Warning (yellow)
Flux::toast(text: 'Your session will expire soon.', variant: 'warning');

// Danger (red)
Flux::toast(text: 'Failed to save changes.', variant: 'danger');
```

## Custom Duration

```php
// 10 seconds
Flux::toast(text: 'Processing...', duration: 10000);

// Permanent (manual dismiss)
Flux::toast(text: 'Action required.', duration: 0);
```

## Positioning

```blade
<flux:toast position="top center" />
```

## Offset for Navbar

```blade
<flux:toast class="pt-24" />
```

## Stacking Multiple Toasts

```blade
<flux:toast.group position="bottom end">
    <flux:toast />
</flux:toast.group>
```

## Always Expanded Stack

```blade
<flux:toast.group expanded>
    <flux:toast />
</flux:toast.group>
```
