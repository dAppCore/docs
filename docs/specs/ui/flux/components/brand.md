# flux:brand

Display a company or application logo and name consistently across interfaces.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `name` | string | - | Company or application name to display |
| `logo` | string | - | URL to image for logo |
| `alt` | string | - | Alternative text for the logo |
| `href` | string | / | URL to navigate to when clicked |

## Slots

| Slot | Description |
|------|-------------|
| `logo` | Custom content for logo section (image, SVG, or HTML) |

---

## Basic Usage

```blade
<flux:brand href="/" logo="/img/logo.png" name="Acme Inc." />
```

## Logo Only (No Name)

```blade
<flux:brand href="/" logo="/img/logo.png" />
```

## Custom Logo Slot

```blade
<flux:brand href="/" name="Acme Inc.">
    <x-slot name="logo">
        <div class="size-6 rounded bg-accent text-accent-foreground flex items-center justify-center">
            <i class="font-serif font-bold">A</i>
        </div>
    </x-slot>
</flux:brand>
```

## Icon-Based Logo

```blade
<flux:brand href="/" name="Launchpad">
    <x-slot name="logo" class="size-6 rounded-full bg-cyan-500 text-white text-xs font-bold">
        <flux:icon name="rocket-launch" variant="micro" />
    </x-slot>
</flux:brand>
```

## SVG Logo

```blade
<flux:brand href="/" name="My App">
    <x-slot name="logo">
        <svg class="size-6" viewBox="0 0 24 24" fill="currentColor">
            <!-- SVG content -->
        </svg>
    </x-slot>
</flux:brand>
```

## In Header Layout

```blade
<flux:header class="px-4! w-full bg-zinc-50 dark:bg-zinc-800 rounded-lg border">
    <flux:brand href="/" name="Acme Inc.">
        <x-slot name="logo" class="bg-accent text-accent-foreground">
            <i class="font-serif font-bold">A</i>
        </x-slot>
    </flux:brand>

    <flux:spacer />

    <flux:navbar>
        <flux:navbar.item href="/dashboard">Dashboard</flux:navbar.item>
        <flux:navbar.item href="/settings">Settings</flux:navbar.item>
    </flux:navbar>
</flux:header>
```

## In Sidebar

```blade
<flux:sidebar>
    <flux:sidebar.header>
        <flux:sidebar.brand name="Acme Inc." logo="/img/logo.png" href="/" />
        <flux:sidebar.collapse />
    </flux:sidebar.header>

    <flux:sidebar.nav>
        <!-- Navigation items -->
    </flux:sidebar.nav>
</flux:sidebar>
```

## Dark Mode Logo

For different logos in light/dark mode, use CSS or separate elements:

```blade
<flux:brand href="/" name="My App">
    <x-slot name="logo">
        <img src="/logo-light.svg" class="size-6 dark:hidden" alt="Logo" />
        <img src="/logo-dark.svg" class="size-6 hidden dark:block" alt="Logo" />
    </x-slot>
</flux:brand>
```

---

## Related Components

- [Header](./header.md) - Top navigation headers
- [Sidebar](./sidebar.md) - Sidebar navigation
