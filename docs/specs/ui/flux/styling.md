# Flux UI - Theming, Customisation & Dark Mode

Complete guide to styling Flux components properly, including theming, dark mode, and customisation approaches.

## Theming Overview

Flux is designed to look polished immediately, but offers extensive customisation through CSS variables. The system uses two primary colour concepts:

1. **Base colour** - Neutral colour for most components (default: zinc)
2. **Accent colour** - Primary action/highlight colour (customisable)

---

## Base Colour Customisation

The default base colour is "zinc," embedded throughout Flux's source. To change it, override in your CSS using the `@theme` directive.

### Switching from Zinc to Slate

Since zinc is baked into Flux, you must map all zinc shades to your chosen colour:

```css
/* resources/css/app.css */
@import "tailwindcss";
@import '../../vendor/livewire/flux/dist/flux.css';

@theme {
    --color-zinc-50: var(--color-slate-50);
    --color-zinc-100: var(--color-slate-100);
    --color-zinc-200: var(--color-slate-200);
    --color-zinc-300: var(--color-slate-300);
    --color-zinc-400: var(--color-slate-400);
    --color-zinc-500: var(--color-slate-500);
    --color-zinc-600: var(--color-slate-600);
    --color-zinc-700: var(--color-slate-700);
    --color-zinc-800: var(--color-slate-800);
    --color-zinc-900: var(--color-slate-900);
    --color-zinc-950: var(--color-slate-950);
}
```

Then use slate utilities normally:
```blade
<div class="text-slate-800 dark:text-white">Content</div>
```

### Common Base Colour Swaps

**Slate**
```css
@theme {
    --color-zinc-50: var(--color-slate-50);
    --color-zinc-100: var(--color-slate-100);
    /* ... all shades ... */
}
```

**Gray**
```css
@theme {
    --color-zinc-50: var(--color-gray-50);
    --color-zinc-100: var(--color-gray-100);
    /* ... all shades ... */
}
```

**Stone**
```css
@theme {
    --color-zinc-50: var(--color-stone-50);
    --color-zinc-100: var(--color-stone-100);
    /* ... all shades ... */
}
```

---

## Accent Colour System

Flux uses CSS variables for accent colours, allowing complete customisation across light and dark modes.

### Accent Colour Variables

| Variable | Purpose | Usage |
|----------|---------|-------|
| `--color-accent` | Primary accent background | Buttons, highlights, links |
| `--color-accent-content` | Darker variant for text contrast | Secondary text on accent |
| `--color-accent-foreground` | Text colour on accent background | Button text, link text |

### Customising Accent Colour

**Example: Change from default (blue) to red**

```css
@theme {
    --color-accent: var(--color-red-500);
    --color-accent-content: var(--color-red-600);
    --color-accent-foreground: var(--color-white);
}

@layer theme {
    .dark {
        --color-accent: var(--color-red-500);
        --color-accent-content: var(--color-red-400);
        --color-accent-foreground: var(--color-white);
    }
}
```

### Using Accent Variables in Components

Components automatically use accent colours:

```blade
<!-- These use --color-accent by default -->
<flux:button variant="primary">Primary action</flux:button>
<flux:tabs>
    <flux:tab>Tab with accent</flux:tab>
</flux:tabs>
```

Disable accent colour use:
```blade
<!-- Use default styling without accent -->
<flux:tabs>
    <flux:tab :accent="false">Regular Tab</flux:tab>
</flux:tabs>
```

### Common Accent Colour Schemes

**Red Accent (Alert/Danger)**
```css
@theme {
    --color-accent: var(--color-red-500);
    --color-accent-content: var(--color-red-600);
    --color-accent-foreground: var(--color-white);
}

@layer theme {
    .dark {
        --color-accent: var(--color-red-500);
        --color-accent-content: var(--color-red-400);
    }
}
```

**Green Accent (Success)**
```css
@theme {
    --color-accent: var(--color-green-500);
    --color-accent-content: var(--color-green-600);
    --color-accent-foreground: var(--color-white);
}

@layer theme {
    .dark {
        --color-accent: var(--color-green-500);
        --color-accent-content: var(--color-green-400);
    }
}
```

**Purple Accent (Premium)**
```css
@theme {
    --color-accent: var(--color-purple-600);
    --color-accent-content: var(--color-purple-700);
    --color-accent-foreground: var(--color-white);
}

@layer theme {
    .dark {
        --color-accent: var(--color-purple-500);
        --color-accent-content: var(--color-purple-400);
    }
}
```

### Using Accent Variables in CSS

Reference variables directly:
```css
.custom-element {
    background-color: var(--color-accent);
    color: var(--color-accent-foreground);
    border-color: var(--color-accent-content);
}
```

In Tailwind classes:
```blade
<div class="bg-[var(--color-accent)] text-[var(--color-accent-foreground)]">
    Custom element
</div>
```

---

## Component-Level Styling

### Using Utility Classes (Recommended)

The simplest approach—pass Tailwind classes directly:

```blade
<flux:button class="max-w-sm">Wide Button</flux:button>
<flux:input class="text-lg placeholder:text-gray-500" />
<flux:card class="border-2 border-blue-500">Content</flux:card>
```

### Handling Class Conflicts

When your classes conflict with Flux's built-in styles, use Tailwind's important modifier:

```blade
<!-- Your bg-zinc-800 overrides Flux defaults -->
<flux:button class="bg-zinc-800! hover:bg-zinc-700! text-white">
    Custom button
</flux:button>
```

**Better approach:** Publish components locally and modify directly (see below).

### Global Style Overrides

Target component elements using `data-flux-*` attributes with CSS:

```css
/* Override all buttons globally */
[data-flux-button] {
    @apply bg-zinc-800 dark:bg-zinc-400 hover:bg-zinc-700;
}

/* Override input styling */
[data-flux-input] {
    @apply border-2 border-blue-500;
}

/* Override navbar items */
[data-flux-navbar] [data-current] {
    @apply bg-accent text-accent-foreground;
}

/* Override table cells */
[data-flux-table] [data-flux-table-cell] {
    @apply border-blue-200;
}
```

### Publishing Components Locally

Take full ownership of any component for unlimited customisation:

**Command:**
```bash
php artisan flux:publish
```

**Interactive mode** - Select specific components, or use `--all`:
```bash
php artisan flux:publish --all
```

**Files are created in:**
```
resources/views/flux/
├── button.blade.php
├── input.blade.php
├── card.blade.php
└── ... (all published components)
```

**After publishing:**
- You control 100% of the component's Blade file
- Modify classes, slots, and structure freely
- Updates to Flux don't affect your versions
- Can evolve components independently

**Example: Customising button.blade.php**
```blade
<!-- Before: Default Flux styling -->
<button class="px-4 py-2 rounded-md bg-zinc-600 text-white">{{ $slot }}</button>

<!-- After: Your custom styling -->
<button class="px-6 py-3 rounded-lg bg-gradient-to-r from-blue-500 to-purple-600 text-white shadow-lg hover:shadow-xl">
    {{ $slot }}
</button>
```

---

## Dark Mode

Flux provides built-in dark mode support that respects user preferences and system settings.

### Configuration

Enable dark mode in your CSS:

```css
@import "tailwindcss";
@import '../../vendor/livewire/flux/dist/flux.css';

@custom-variant dark (&:where(.dark, .dark *));
```

This allows Flux to toggle dark mode by adding/removing a `.dark` class on the `<html>` element.

### Automatic Handling

By default, Flux manages appearance automatically:
- Detects system preference (light/dark)
- Persists user choice in localStorage
- Adds `.dark` class to `<html>` when dark mode is active

Disable automatic handling:
```blade
<!-- In your layout file, remove: -->
@fluxAppearance

<!-- Then manage appearance manually -->
```

### JavaScript Utilities

Flux provides two key utilities for dark mode management:

**`$flux.appearance`** - User preference string
```javascript
$flux.appearance = 'light'   // Force light mode
$flux.appearance = 'dark'    // Force dark mode
$flux.appearance = 'system'  // Follow system preference
```

**`$flux.dark`** - Boolean for current dark mode state
```javascript
if ($flux.dark) {
    // Dark mode is active
}
```

Access via Alpine directives:
```blade
<button x-data x-on:click="$flux.dark = ! $flux.dark" icon="moon">
    Toggle Dark Mode
</button>
```

### Dark Mode Toggle Component

**Simple toggle button:**
```blade
<button
    x-data
    x-on:click="$flux.dark = ! $flux.dark"
    :class="$flux.dark ? 'bg-yellow-400' : 'bg-zinc-800'"
>
    @if ($flux.dark)
        ☀️ Light Mode
    @else
        🌙 Dark Mode
    @endif
</button>
```

**Radio group (Light/Dark/System):**
```blade
<div x-data x-model="$flux.appearance" class="space-y-2">
    <label>
        <input type="radio" value="light" />
        Light Mode
    </label>
    <label>
        <input type="radio" value="dark" />
        Dark Mode
    </label>
    <label>
        <input type="radio" value="system" />
        System Preference
    </label>
</div>
```

**Flux radio group:**
```blade
<flux:radio.group x-data x-model="$flux.appearance">
    <flux:radio value="light">Light Mode</flux:radio>
    <flux:radio value="dark">Dark Mode</flux:radio>
    <flux:radio value="system">System Preference</flux:radio>
</flux:radio.group>
```

### Styling for Dark Mode

All Flux components automatically support dark mode through CSS classes. Use `dark:` prefix for dark mode styles:

```blade
<div class="bg-white dark:bg-zinc-900 text-black dark:text-white">
    Content that adapts to dark mode
</div>

<flux:card class="bg-white dark:bg-zinc-800 border-zinc-200 dark:border-zinc-700">
    Dark mode aware card
</flux:card>
```

### Custom Dark Mode Styles

Define custom dark mode variables:

```css
@layer theme {
    .dark {
        --color-accent: var(--color-blue-500);
        --color-accent-content: var(--color-blue-400);
        --color-accent-foreground: var(--color-white);

        /* Custom theme variables */
        --background: var(--color-zinc-950);
        --foreground: var(--color-zinc-50);
    }
}
```

Use in components:
```blade
<div class="bg-[var(--background)] text-[var(--foreground)]">
    Custom dark mode styling
</div>
```

---

## Theme Builder

Flux provides an interactive theme builder at `https://fluxui.dev/themes` for previewing hand-picked colour schemes before implementation.

Steps:
1. Visit theme builder on Flux docs
2. Preview various colour combinations
3. Copy the CSS variables
4. Paste into your `resources/css/app.css`

---

## Complete Theming Example

**Host UK Theme - Blue/Purple Accent**

```css
/* resources/css/app.css */
@import "tailwindcss";
@import '../../vendor/livewire/flux/dist/flux.css';
@custom-variant dark (&:where(.dark, .dark *));

/* Change base from zinc to slate */
@theme {
    --color-zinc-50: var(--color-slate-50);
    --color-zinc-100: var(--color-slate-100);
    --color-zinc-200: var(--color-slate-200);
    --color-zinc-300: var(--color-slate-300);
    --color-zinc-400: var(--color-slate-400);
    --color-zinc-500: var(--color-slate-500);
    --color-zinc-600: var(--color-slate-600);
    --color-zinc-700: var(--color-slate-700);
    --color-zinc-800: var(--color-slate-800);
    --color-zinc-900: var(--color-slate-900);
    --color-zinc-950: var(--color-slate-950);

    /* Set accent to Host UK brand purple */
    --color-accent: var(--color-violet-600);
    --color-accent-content: var(--color-violet-700);
    --color-accent-foreground: var(--color-white);
}

/* Dark mode adjustments */
@layer theme {
    .dark {
        --color-accent: var(--color-violet-500);
        --color-accent-content: var(--color-violet-400);
        --color-accent-foreground: var(--color-white);
    }
}

/* Global component overrides */
[data-flux-navbar] {
    @apply border-b border-slate-200 dark:border-slate-800;
}

[data-flux-sidebar] {
    @apply bg-slate-50 dark:bg-slate-950;
}
```

---

## Customisation Priority

1. **Component props** - Highest priority
   ```blade
   <flux:button variant="primary" size="lg" color="red">Button</flux:button>
   ```

2. **Inline classes** - Override component defaults
   ```blade
   <flux:button class="bg-custom!">Button</flux:button>
   ```

3. **Published components** - Full control
   ```bash
   php artisan flux:publish
   ```

4. **Global CSS** - Affects all instances
   ```css
   [data-flux-button] { @apply bg-custom; }
   ```

5. **CSS variables** - Base theme
   ```css
   --color-accent: var(--color-blue-600);
   ```

---

## Common Customisation Patterns

### Creating Custom Button Variant

**Option 1: With classes**
```blade
<flux:button class="bg-gradient-to-r from-blue-500 to-purple-600 text-white hover:shadow-lg">
    Gradient Button
</flux:button>
```

**Option 2: Publish and create variant**
```bash
php artisan flux:publish
```

Edit `resources/views/flux/button.blade.php` to add custom variant:
```blade
@if ($variant === 'gradient')
    <button class="bg-gradient-to-r from-blue-500 to-purple-600 text-white ...">
        {{ $slot }}
    </button>
@endif
```

Use it:
```blade
<flux:button variant="gradient">Gradient Button</flux:button>
```

### Custom Input Styling

```blade
<!-- Add custom classes -->
<flux:input
    wire:model="email"
    type="email"
    label="Email"
    class="border-2 border-blue-500 rounded-lg"
    placeholder="your@email.com"
/>
```

Or globally override:
```css
[data-flux-input] {
    @apply border-2 border-blue-500 rounded-lg;
}
```

### Custom Card Styling

```blade
<flux:card class="bg-gradient-to-br from-blue-50 to-indigo-50 dark:from-indigo-900 dark:to-blue-900 border-2 border-indigo-200 dark:border-indigo-700 shadow-lg">
    <flux:heading>Premium Card</flux:heading>
    <flux:text>With custom styling</flux:text>
</flux:card>
```

---

## Testing Responsive & Dark Modes

### Chrome DevTools

1. Open Chrome DevTools (F12)
2. Press `Ctrl+Shift+P` (or `Cmd+Shift+P`)
3. Type "Dark Mode" or "Emulate CSS media"
4. Select `prefers-color-scheme: dark`

### Local Testing Script

```html
<!-- Toggle dark mode for testing -->
<button onclick="document.documentElement.classList.toggle('dark')">
    Toggle Dark Mode
</button>
```

---

## Important Reminders

1. **"We style, you space"** - Flux handles component styling; you manage layout spacing
2. **Use CSS variables** - Most flexible for theming
3. **Publish components** - For complex customisations
4. **Avoid excessive `!` modifiers** - Indicates need for better approach
5. **Test dark mode** - Ensure accent colours work in both modes
6. **Use `data-flux-*` attributes** - For global style overrides

---

Last updated: January 2026
