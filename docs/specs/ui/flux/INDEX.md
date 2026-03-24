# Flux UI Documentation

Complete documentation for Flux Pro components (Livewire team).

**License:** Flux Pro
**Source:** https://fluxui.dev

---

## Component Reference

All components have individual documentation files in `components/`.

### Layouts

| Component | File | Description |
|-----------|------|-------------|
| Header | [header.md](components/header.md) | Full-width top navigation |
| Sidebar | [sidebar.md](components/sidebar.md) | Vertical navigation with groups |

### Components A-C

| Component | File | Pro |
|-----------|------|-----|
| Accordion | [accordion.md](components/accordion.md) | |
| Autocomplete | [autocomplete.md](components/autocomplete.md) | ✓ |
| Avatar | [avatar.md](components/avatar.md) | |
| Badge | [badge.md](components/badge.md) | |
| Brand | [brand.md](components/brand.md) | |
| Breadcrumbs | [breadcrumbs.md](components/breadcrumbs.md) | |
| Button | [button.md](components/button.md) | |
| Calendar | [calendar.md](components/calendar.md) | ✓ |
| Callout | [callout.md](components/callout.md) | |
| Card | [card.md](components/card.md) | |
| Chart | [chart.md](components/chart.md) | ✓ |
| Checkbox | [checkbox.md](components/checkbox.md) | Partial |
| Command | [command.md](components/command.md) | |
| Composer | [composer.md](components/composer.md) | ✓ |
| Context | [context.md](components/context.md) | |

### Components D-I

| Component | File | Pro |
|-----------|------|-----|
| Date Picker | [date-picker.md](components/date-picker.md) | ✓ |
| Dropdown | [dropdown.md](components/dropdown.md) | |
| Editor | [editor.md](components/editor.md) | ✓ |
| Field | [field.md](components/field.md) | |
| File Upload | [file-upload.md](components/file-upload.md) | |
| Heading | [heading.md](components/heading.md) | |
| Icon | [icon.md](components/icon.md) | |
| Input | [input.md](components/input.md) | |

### Components K-P

| Component | File | Pro |
|-----------|------|-----|
| Kanban | [kanban.md](components/kanban.md) | ✓ |
| Modal | [modal.md](components/modal.md) | |
| Navbar | [navbar.md](components/navbar.md) | |
| OTP Input | [otp-input.md](components/otp-input.md) | |
| Pagination | [pagination.md](components/pagination.md) | |
| Pillbox | [pillbox.md](components/pillbox.md) | ✓ |
| Popover | [popover.md](components/popover.md) | |
| Profile | [profile.md](components/profile.md) | |

### Components R-T

| Component | File | Pro |
|-----------|------|-----|
| Radio | [radio.md](components/radio.md) | Partial |
| Select | [select.md](components/select.md) | Partial |
| Separator | [separator.md](components/separator.md) | |
| Skeleton | [skeleton.md](components/skeleton.md) | |
| Slider | [slider.md](components/slider.md) | |
| Switch | [switch.md](components/switch.md) | |
| Table | [table.md](components/table.md) | |
| Tabs | [tabs.md](components/tabs.md) | |
| Text | [text.md](components/text.md) | |
| Textarea | [textarea.md](components/textarea.md) | |
| Time Picker | [time-picker.md](components/time-picker.md) | ✓ |
| Toast | [toast.md](components/toast.md) | |
| Tooltip | [tooltip.md](components/tooltip.md) | |

---

## Quick Patterns

### wire:model Binding

```blade
<flux:input wire:model="name" />
<flux:input wire:model.live="search" />
<flux:input wire:model.blur="email" />
```

### Shorthand Labels

```blade
{{-- Verbose --}}
<flux:field>
    <flux:label>Email</flux:label>
    <flux:input wire:model="email" />
</flux:field>

{{-- Shorthand (equivalent) --}}
<flux:input label="Email" wire:model="email" />
```

### Size Variants

Most components support: `xs`, `sm`, (default), `lg`, `xl`

### Icon Variants

```blade
<flux:icon.bolt />                  {{-- outline 24px --}}
<flux:icon.bolt variant="solid" />  {{-- solid 24px --}}
<flux:icon.bolt variant="mini" />   {{-- 20px --}}
<flux:icon.bolt variant="micro" />  {{-- 16px --}}
```

### Modal Control

```php
// Livewire
Flux::modal('name')->show();
Flux::modal('name')->close();
```

```javascript
// Alpine.js
$flux.modal('name').show()

// JavaScript
Flux.modal('name').show()
```

### Toast Notifications

```php
Flux::toast('Message saved.');
Flux::toast(text: 'Success!', variant: 'success');
```

---

## Core Principles

1. **We Style, You Space** - Flux handles styling; you manage layout spacing
2. **Props vs Attributes** - Props configure; attributes forward to DOM
3. **Composable** - Simple or complex patterns with same components
4. **Livewire Native** - Deep wire:model integration
5. **Dark Mode** - Automatic handling

---

## Theming

### Dark Mode Toggle

```javascript
$flux.appearance = 'light' | 'dark' | 'system'
$flux.dark = true | false
```

### Accent Colour

```css
@theme {
    --color-accent: var(--color-blue-600);
}
```

---

## Resources

- [Flux UI Docs](https://fluxui.dev)
- [Heroicons](https://heroicons.com)
- [Livewire 3](https://livewire.laravel.com)
- [Tailwind CSS](https://tailwindcss.com)

---

**Total Components Documented:** 38
**Last Updated:** January 2026
