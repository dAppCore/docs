# Flux Theming

Customise colours through CSS variables while maintaining aesthetic consistency.

## Colour System

**Two Primary Colours:**
- **Base Colour** - Text, backgrounds, borders (default: zinc)
- **Accent Colour** - Primary buttons and interactive elements

## Base Colour

Replace default zinc with another grey shade:

```css
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

## Accent Colour

Define three variables for light and dark modes:

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

| Variable | Purpose |
|----------|---------|
| `--color-accent` | Primary button backgrounds |
| `--color-accent-content` | Text content (darker, readable) |
| `--color-accent-foreground` | Text on accent backgrounds |

## Disable Accent on Components

```blade
<flux:link :accent="false">Profile</flux:link>

<flux:tabs>
    <flux:tab :accent="false">Profile</flux:tab>
</flux:tabs>

<flux:navbar>
    <flux:navbar.item :accent="false">Menu</flux:navbar.item>
</flux:navbar>
```

## Theme Builder

Interactive theme builder available at https://fluxui.dev/themes
