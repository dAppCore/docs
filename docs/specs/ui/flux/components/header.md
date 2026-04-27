# flux:header

Full-width top navigation with branding, navigation items, and user profile management.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `sticky` | boolean | false | Makes header remain fixed during scrolling |
| `container` | boolean | false | Constrains content to container width |

## Slots

| Slot | Description |
|------|-------------|
| `default` | Primary content area (branding, navigation, profile elements) |

## Related Components

- `flux:brand` - Logo and company name display
- `flux:navbar` - Navigation bar container
- `flux:navbar.item` - Individual navigation items
- `flux:dropdown` - Dropdown menu wrapper
- `flux:profile` - User profile avatar display
- `flux:separator` - Visual dividers

## Basic Example

```blade
<flux:header sticky container class="bg-zinc-50 border-b dark:bg-zinc-900 dark:border-zinc-700">
    <flux:sidebar.toggle icon="bars-2" inset="left" class="lg:hidden" />

    <flux:brand href="/" logo="/logo.svg" name="Acme Inc." class="max-lg:hidden" />

    <flux:navbar class="-mb-px max-lg:hidden">
        <flux:navbar.item href="/" current>Dashboard</flux:navbar.item>
        <flux:navbar.item href="/orders">Orders</flux:navbar.item>
        <flux:navbar.item href="/products">Products</flux:navbar.item>
    </flux:navbar>

    <flux:spacer />

    <flux:dropdown position="bottom" align="end">
        <flux:profile avatar="/avatar.jpg" />

        <flux:menu>
            <flux:menu.item href="/settings">Settings</flux:menu.item>
            <flux:menu.separator />
            <flux:menu.item href="/logout">Logout</flux:menu.item>
        </flux:menu>
    </flux:dropdown>
</flux:header>

<flux:main container>
    {{ $slot }}
</flux:main>
```

## Header with Mobile Sidebar

```blade
<flux:sidebar stashable sticky class="lg:hidden bg-zinc-50 border-r dark:bg-zinc-900 dark:border-zinc-700">
    <flux:sidebar.toggle class="lg:hidden" icon="x-mark" />

    <flux:navlist variant="outline">
        <flux:navlist.item icon="home" href="/" current>Dashboard</flux:navlist.item>
        <flux:navlist.item icon="shopping-bag" href="/orders">Orders</flux:navlist.item>
        <flux:navlist.item icon="cube" href="/products">Products</flux:navlist.item>
    </flux:navlist>
</flux:sidebar>

<flux:header sticky container>
    <flux:sidebar.toggle icon="bars-2" inset="left" class="lg:hidden" />

    <flux:brand href="/" logo="/logo.svg" name="Acme Inc." />

    <flux:navbar class="-mb-px max-lg:hidden">
        <flux:navbar.item href="/" current>Dashboard</flux:navbar.item>
        <flux:navbar.item href="/orders">Orders</flux:navbar.item>
        <flux:navbar.item href="/products">Products</flux:navbar.item>
    </flux:navbar>

    <flux:spacer />

    <flux:dropdown position="bottom" align="end">
        <flux:profile avatar="/avatar.jpg" />

        <flux:menu>
            <flux:menu.item href="/settings">Settings</flux:menu.item>
            <flux:menu.separator />
            <flux:menu.item href="/logout">Logout</flux:menu.item>
        </flux:menu>
    </flux:dropdown>
</flux:header>

<flux:main container>
    {{ $slot }}
</flux:main>
```

## Header with Secondary Sidebar

```blade
<flux:header sticky container class="border-b bg-zinc-50 dark:bg-zinc-900 dark:border-zinc-700">
    <flux:brand href="/" logo="/logo.svg" name="Acme Inc." />

    <flux:navbar class="-mb-px">
        <flux:navbar.item href="/" current>Dashboard</flux:navbar.item>
        <flux:navbar.item href="/settings">Settings</flux:navbar.item>
    </flux:navbar>

    <flux:spacer />

    <flux:dropdown position="bottom" align="end">
        <flux:profile avatar="/avatar.jpg" />

        <flux:menu>
            <flux:menu.item href="/logout">Logout</flux:menu.item>
        </flux:menu>
    </flux:dropdown>
</flux:header>

<div class="flex min-h-screen">
    <flux:sidebar sticky class="border-r bg-zinc-50 dark:bg-zinc-900 dark:border-zinc-700">
        <flux:navlist>
            <flux:navlist.item icon="user" href="/settings/profile" current>Profile</flux:navlist.item>
            <flux:navlist.item icon="lock-closed" href="/settings/password">Password</flux:navlist.item>
            <flux:navlist.item icon="bell" href="/settings/notifications">Notifications</flux:navlist.item>
        </flux:navlist>
    </flux:sidebar>

    <flux:main container>
        {{ $slot }}
    </flux:main>
</div>
```

## Dark Mode

Header supports dark mode via Tailwind dark: classes:

```blade
<flux:header class="bg-zinc-50 border-b dark:bg-zinc-900 dark:border-zinc-700">
    ...
</flux:header>
```

## CSS Classes

Common classes for headers:
- `bg-zinc-50` / `dark:bg-zinc-900` - Background colour
- `border-b` / `dark:border-zinc-700` - Border
- `-mb-px` - Negative margin for navbar alignment
