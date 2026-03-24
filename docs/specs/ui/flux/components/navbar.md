# flux:navbar

Horizontal and vertical navigation with icons, badges, and dropdowns.

## Components

### flux:navbar (Horizontal)

Container for horizontal navigation.

### flux:navbar.item

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `href` | string | - | Link URL |
| `current` | boolean | auto | Active styling (auto-detected from URL) |
| `icon` | string | - | Leading icon |
| `icon:trailing` | string | - | Trailing icon |
| `badge` | string/boolean | - | Badge content |
| `badge:color` | string | - | Badge colour |
| `badge:variant` | string | solid | `solid`, `outline` |

### flux:navlist (Vertical)

Container for vertical navigation (sidebar).

### flux:navlist.item

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `href` | string | - | Link URL |
| `current` | boolean | auto | Active styling |
| `icon` | string | - | Leading icon |
| `badge` | string/boolean | - | Badge content |
| `badge:color` | string | - | Badge colour |
| `badge:variant` | string | solid | `solid`, `outline` |

### flux:navlist.group

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `heading` | string | - | Group heading |
| `expandable` | boolean | false | Makes group collapsible |
| `expanded` | boolean | true | Initial expanded state |

## Basic Navbar

```blade
<flux:navbar>
    <flux:navbar.item href="/">Home</flux:navbar.item>
    <flux:navbar.item href="/features">Features</flux:navbar.item>
    <flux:navbar.item href="/pricing">Pricing</flux:navbar.item>
    <flux:navbar.item href="/about">About</flux:navbar.item>
</flux:navbar>
```

## With Icons

```blade
<flux:navbar>
    <flux:navbar.item href="/" icon="home">Home</flux:navbar.item>
    <flux:navbar.item href="/users" icon="users">Users</flux:navbar.item>
    <flux:navbar.item href="/settings" icon="cog-6-tooth">Settings</flux:navbar.item>
</flux:navbar>
```

## With Badges

```blade
<flux:navbar>
    <flux:navbar.item href="/inbox" badge="12">Inbox</flux:navbar.item>
    <flux:navbar.item href="/calendar" badge="Pro" badge:color="lime">Calendar</flux:navbar.item>
</flux:navbar>
```

## Dropdown Navigation

```blade
<flux:navbar>
    <flux:navbar.item href="/">Home</flux:navbar.item>
    <flux:dropdown>
        <flux:navbar.item icon:trailing="chevron-down">Account</flux:navbar.item>
        <flux:navmenu>
            <flux:navmenu.item href="/profile">Profile</flux:navmenu.item>
            <flux:navmenu.item href="/settings">Settings</flux:navmenu.item>
            <flux:navmenu.separator />
            <flux:navmenu.item href="/logout">Logout</flux:navmenu.item>
        </flux:navmenu>
    </flux:dropdown>
</flux:navbar>
```

## Vertical Navigation (Navlist)

```blade
<flux:navlist class="w-64">
    <flux:navlist.item href="/" icon="home">Home</flux:navlist.item>
    <flux:navlist.item href="/dashboard" icon="chart-bar">Dashboard</flux:navlist.item>
    <flux:navlist.item href="/projects" icon="folder">Projects</flux:navlist.item>
    <flux:navlist.item href="/team" icon="users">Team</flux:navlist.item>
</flux:navlist>
```

## Collapsible Groups

```blade
<flux:navlist class="w-64">
    <flux:navlist.group heading="Platform" expandable>
        <flux:navlist.item href="/dashboard" icon="home">Dashboard</flux:navlist.item>
        <flux:navlist.item href="/analytics" icon="chart-bar">Analytics</flux:navlist.item>
    </flux:navlist.group>

    <flux:navlist.group heading="Account" expandable :expanded="false">
        <flux:navlist.item href="/profile" icon="user">Profile</flux:navlist.item>
        <flux:navlist.item href="/settings" icon="cog">Settings</flux:navlist.item>
    </flux:navlist.group>
</flux:navlist>
```

## Manual Current State

```blade
<flux:navbar.item href="/dashboard" :current="request()->is('dashboard*')">
    Dashboard
</flux:navbar.item>
```
