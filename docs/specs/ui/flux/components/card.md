# flux:card

A container for related content, such as a form, alert, or data list.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `size` | string | default | `sm` (small), default (regular) |
| `class` | string | - | Additional CSS classes |

## Slots

| Slot | Description |
|------|-------------|
| `default` | Content to display within the card |

---

## Basic Usage

```blade
<flux:card>
    <p>Card content goes here.</p>
</flux:card>
```

## With Heading and Text

```blade
<flux:card class="space-y-6">
    <flux:heading size="lg">Card Title</flux:heading>
    <flux:text>Card description or content.</flux:text>
</flux:card>
```

## Login Form

```blade
<flux:card class="space-y-6">
    <div>
        <flux:heading size="lg">Log in to your account</flux:heading>
        <flux:text class="mt-2">Welcome back! Please enter your details.</flux:text>
    </div>

    <flux:input label="Email" type="email" placeholder="you@example.com" />
    <flux:input label="Password" type="password" />

    <div class="flex justify-between">
        <flux:checkbox label="Remember me" />
        <flux:link href="/forgot-password">Forgot password?</flux:link>
    </div>

    <flux:button variant="primary" class="w-full">Log in</flux:button>
</flux:card>
```

## Small Size

Compact content like notifications, alerts, or brief summaries:

```blade
<flux:card size="sm" class="hover:bg-zinc-50 dark:hover:bg-zinc-800">
    <flux:heading>Latest on our blog</flux:heading>
    <flux:text class="mt-1">Stay up to date with our latest insights and updates.</flux:text>
</flux:card>
```

## With Header Actions

```blade
<flux:card class="space-y-6">
    <div class="flex items-center justify-between">
        <flux:heading size="lg">Are you sure?</flux:heading>
        <flux:button variant="ghost" icon="x-mark" />
    </div>

    <flux:text>This action cannot be undone.</flux:text>

    <div class="flex gap-2">
        <flux:button variant="ghost">Cancel</flux:button>
        <flux:button variant="danger">Delete</flux:button>
    </div>
</flux:card>
```

## With Sections

```blade
<flux:card>
    <div class="p-6 border-b dark:border-zinc-700">
        <flux:heading>Settings</flux:heading>
    </div>

    <div class="p-6 space-y-4">
        <flux:input label="Name" />
        <flux:input label="Email" type="email" />
    </div>

    <div class="p-6 border-t dark:border-zinc-700 flex justify-end gap-2">
        <flux:button variant="ghost">Cancel</flux:button>
        <flux:button variant="primary">Save</flux:button>
    </div>
</flux:card>
```

## Grid of Cards

```blade
<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
    <flux:card size="sm">
        <flux:heading>Feature 1</flux:heading>
        <flux:text>Description of feature one.</flux:text>
    </flux:card>

    <flux:card size="sm">
        <flux:heading>Feature 2</flux:heading>
        <flux:text>Description of feature two.</flux:text>
    </flux:card>

    <flux:card size="sm">
        <flux:heading>Feature 3</flux:heading>
        <flux:text>Description of feature three.</flux:text>
    </flux:card>
</div>
```

## Clickable Card

```blade
<a href="/article">
    <flux:card size="sm" class="hover:bg-zinc-50 dark:hover:bg-zinc-800 transition-colors">
        <flux:heading>Read our latest article</flux:heading>
        <flux:text>Click to learn more about our updates.</flux:text>
    </flux:card>
</a>
```

## With Image

```blade
<flux:card class="overflow-hidden">
    <img src="/image.jpg" alt="Cover" class="-mx-6 -mt-6 mb-4" />
    <flux:heading>Article Title</flux:heading>
    <flux:text>Article excerpt goes here...</flux:text>
</flux:card>
```

---

## Related Components

- [Heading](./heading.md)
- [Text](./text.md)
