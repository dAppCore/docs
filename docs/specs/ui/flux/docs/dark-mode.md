# Flux Dark Mode

Automatic dark mode with user preference persistence.

## Setup

Configure Tailwind in `resources/css/app.css`:

```css
@import "tailwindcss";
@import '../../vendor/livewire/flux/dist/flux.css';
@custom-variant dark (&:where(.dark, .dark *));
```

Flux toggles dark mode by adding/removing `.dark` class on `<html>`.

## Disable Default Behaviour

Remove `@fluxAppearance` directive from layout to handle manually.

## JavaScript Utilities

### Alpine.js

```javascript
$flux.appearance = 'light' | 'dark' | 'system'
$flux.dark = true | false
```

### Vanilla JavaScript

```javascript
Flux.dark = !Flux.dark
Flux.appearance = 'light'
```

| Property | Purpose |
|----------|---------|
| `appearance` | User preference (persisted) |
| `dark` | Current theme state |

## Toggle Button

```blade
<flux:button
    x-data
    x-on:click="$flux.dark = ! $flux.dark"
    icon="moon"
    variant="subtle"
/>
```

## Radio Group

```blade
<flux:radio.group x-data x-model="$flux.appearance">
    <flux:radio value="light">Light</flux:radio>
    <flux:radio value="dark">Dark</flux:radio>
    <flux:radio value="system">System</flux:radio>
</flux:radio.group>
```

## Switch Control

```blade
<flux:switch x-data x-model="$flux.dark" label="Dark mode" />
```

## Features

- Stores preferences in localStorage
- Respects system colour scheme
- Monitors system preference changes
- Manages dark mode in iframes
