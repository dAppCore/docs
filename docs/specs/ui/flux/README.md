# Flux UI Documentation

> **License:** Flux Pro - We have a licensed Flux Pro subscription

This directory contains local documentation for Flux UI, a composable component system for Livewire 3 applications. Flux prioritises simplicity through intuitive component design, whilst offering composability for complex use cases.

## What is Flux?

Flux is "not a set of UI Blade components, it's a system of them." It's designed to work seamlessly with Livewire 3 and Tailwind CSS 4, providing a foundation for building modern web applications with minimal friction.

## Key Features

- **Livewire 3 native** - Built specifically for Livewire 3
- **Tailwind CSS 4** - Modern utility-first styling
- **Composable** - Build complex components from simple building blocks
- **Dark mode ready** - Built-in support for light/dark themes
- **Accessible** - WCAG compliant, keyboard navigation, semantic HTML
- **CSS-first approach** - Uses modern browser APIs like popover and dialog elements
- **Type-safe** - Full type hints for IDE autocomplete

## Design Philosophy

Flux follows seven core principles:

1. **Simplicity** - Straightforward syntax and component usage
2. **Complexity (Composability)** - Flexible alternatives for advanced use cases
3. **Friendliness** - Familiar terminology over technical jargon
4. **Composition** - Mix-and-match core components for flexible solutions
5. **Consistency** - Predictable naming conventions and patterns
6. **Brevity** - Short, clear names without excessive hyphenation
7. **Modern Browser Leverage** - Native browser APIs over custom JavaScript

## Styling Philosophy

Flux uses the approach: **"We style, you space"**

- Flux components handle all component styling and internal padding
- You manage margins and spacing in your layout
- This maintains flexibility and context-appropriate spacing

## Getting Started

Start with the [Principles](./principles.md) to understand the core design patterns, then reference specific components as needed.

## Documentation Sections

| Document | Purpose |
|----------|---------|
| [principles.md](./principles.md) | Core concepts and design patterns |
| [components.md](./components.md) | Complete component reference |
| [navbar.md](./navbar.md) | Navbar, Navlist, and Navmenu components |
| [layouts.md](./layouts.md) | Sidebar and Header layout patterns |
| [styling.md](./styling.md) | Theming, customization, and dark mode |
| [forms.md](./forms.md) | Form components (input, select, etc.) |
| [button.md](./button.md) | Button component and variants |

## Core Components Summary

### Navigation
- `flux:navbar` - Horizontal navigation
- `flux:navlist` - Vertical sidebar navigation
- `flux:navmenu` - Dropdown menus for navigation
- `flux:dropdown` - Expandable dropdown menus
- `flux:menu` - Complex action menus

### Forms & Inputs
- `flux:input` - Text inputs with variants
- `flux:select` - Dropdown select fields
- `flux:checkbox` - Checkbox inputs
- `flux:radio` - Radio button inputs
- `flux:switch` - Toggle switches
- `flux:textarea` - Multi-line text inputs
- `flux:field` - Form field wrapper

### Layout
- `flux:sidebar` - Persistent sidebar layout
- `flux:header` - Top navigation header
- `flux:card` - Content container

### Data Display
- `flux:table` - Structured data with sorting/pagination
- `flux:badge` - Status and category highlighting
- `flux:avatar` - User profile images/initials
- `flux:heading` - Hierarchical headings
- `flux:text` - Formatted text

### Other
- `flux:button` - Buttons with variants and icons
- `flux:icon` - SVG icons
- `flux:modal` - Dialog boxes
- `flux:toast` - Notifications
- `flux:tooltip` - Hover help text

## Quick Reference

### Basic Button
```blade
<flux:button>Click me</flux:button>
<flux:button variant="primary" icon="check">Submit</flux:button>
<flux:button size="sm" variant="ghost">Small</flux:button>
```

### Basic Input
```blade
<flux:input wire:model="email" type="email" label="Email" />
<flux:input label="Search" icon="magnifying-glass" placeholder="Search..." />
```

### Navigation
```blade
<flux:navbar>
    <flux:navbar.item href="/" current icon="home">Home</flux:navbar.item>
    <flux:navbar.item href="/about" icon="information-circle">About</flux:navbar.item>
</flux:navbar>
```

### Sidebar Layout
```blade
<div class="flex h-screen">
    <flux:sidebar sticky collapsible="mobile">
        <flux:sidebar.header>
            <flux:sidebar.brand>App Name</flux:sidebar.brand>
        </flux:sidebar.header>

        <flux:sidebar.nav>
            <flux:sidebar.item href="/" icon="home" current>Dashboard</flux:sidebar.item>
        </flux:sidebar.nav>
    </flux:sidebar>

    <main class="flex-1 overflow-auto">
        <!-- Page content -->
    </main>
</div>
```

## Important Props Patterns

### Styling Props
- `class` - Additional Tailwind classes (merged with component styles)
- `variant` - Visual style variations (default, primary, filled, danger, ghost, subtle)
- `size` - Sizing options (sm, default, lg)
- `color` - Colour choice (zinc, red, blue, green, etc.)

### Icon Props
- `icon` - Leading icon name (e.g., "home", "check")
- `icon:trailing` - Trailing icon name
- `icon:variant` - Icon style (outline, solid, mini, micro)

### State Props
- `disabled` - Disables interaction
- `readonly` - Locks content (input only)
- `invalid` - Error state
- `current` - Marks as active (navigation)

### Layout Props
- `inset` - Negative margin adjustment
- `sticky` - Fixed positioning
- `collapsible` - Collapse/expand capability

## Dark Mode

Flux handles dark mode automatically through CSS classes. Access dark mode utilities via:

```javascript
$flux.appearance  // User preference: 'light', 'dark', 'system'
$flux.dark        // Boolean: true/false
```

Toggle dark mode in your app:
```blade
<flux:button x-data x-on:click="$flux.dark = ! $flux.dark" icon="moon" />
```

## Customisation

Three approaches to customise Flux components:

### 1. Tailwind Classes (Simplest)
```blade
<flux:button class="max-w-sm">Click me</flux:button>
```

### 2. Publish Components Locally
```bash
php artisan flux:publish         # Interactive selection
php artisan flux:publish --all   # Publish everything
```

This puts components in `resources/views/flux/` for full control.

### 3. Global CSS Overrides
```css
[data-flux-button] {
    @apply bg-zinc-800 dark:bg-zinc-400 hover:bg-zinc-700;
}
```

## Theming

Flux uses CSS variables for theming. Customize through `@theme` directive:

```css
@theme {
    --color-accent: var(--color-blue-500);
    --color-accent-content: var(--color-blue-600);
    --color-accent-foreground: var(--color-white);
}

@layer theme {
    .dark {
        --color-accent: var(--color-blue-500);
        --color-accent-content: var(--color-blue-400);
    }
}
```

See [styling.md](./styling.md) for complete theming guide.

## Links

- [Official Flux Documentation](https://fluxui.dev)
- [GitHub Repository](https://github.com/livewire/flux)
- [Tailwind CSS](https://tailwindcss.com)
- [Livewire Documentation](https://livewire.laravel.com)

## Common Gotchas

1. **Blade Conditionals** - Can't use `@if` in component opening tags. Use dynamic attributes instead:
   ```blade
   <!-- Wrong -->
   <flux:button @if($show) variant="primary" @endif></flux:button>

   <!-- Right -->
   <flux:button :variant="$show ? 'primary' : 'default'"></flux:button>
   ```

2. **Class Conflicts** - When user classes conflict with Flux styles, use `!` modifier:
   ```blade
   <flux:button class="bg-zinc-800! hover:bg-zinc-700!">Custom</flux:button>
   ```

3. **Component Opening Tags** - Dynamic expressions don't work. Use attributes with dynamic syntax.

4. **Data Attributes** - Components emit `data-flux-*` attributes for styling hooks, not customisation.

---

Last updated: January 2026 | Flux Pro License
