# Flux UI - Navigation Components

Complete documentation for navbar, navlist, and navmenu components.

## flux:navbar - Horizontal Navigation

The navbar component arranges navigation links horizontally with automatic detection of the current page.

### Basic Structure

```blade
<flux:navbar>
    <flux:navbar.item href="/">Home</flux:navbar.item>
    <flux:navbar.item href="/about">About</flux:navbar.item>
    <flux:navbar.item href="/contact">Contact</flux:navbar.item>
</flux:navbar>
```

### flux:navbar Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `class` | string | - | Additional CSS classes for styling |

### flux:navbar.item Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `href` | string | - | Destination URL for the link |
| `current` | boolean | auto-detected | Manually mark item as active |
| `icon` | string | - | Leading icon name (Heroicon) |
| `icon:trailing` | string | - | Trailing icon name |
| `badge` | string, boolean, slot | - | Badge content (count, label, etc.) |
| `badge:color` | string | - | Badge colour (zinc, red, blue, green, etc.) |
| `badge:variant` | string | solid | Badge style: solid or outline |
| `class` | string | - | Additional CSS classes |

### Usage Examples

**Basic navbar:**
```blade
<flux:navbar>
    <flux:navbar.item href="/" current>Home</flux:navbar.item>
    <flux:navbar.item href="/about">About</flux:navbar.item>
    <flux:navbar.item href="/services">Services</flux:navbar.item>
    <flux:navbar.item href="/blog">Blog</flux:navbar.item>
</flux:navbar>
```

**Navbar with icons:**
```blade
<flux:navbar>
    <flux:navbar.item href="/" icon="home" current>Home</flux:navbar.item>
    <flux:navbar.item href="/inbox" icon="inbox">Inbox</flux:navbar.item>
    <flux:navbar.item href="/settings" icon="cog-6-tooth">Settings</flux:navbar.item>
</flux:navbar>
```

**Navbar with badges:**
```blade
<flux:navbar>
    <flux:navbar.item href="/" current>Home</flux:navbar.item>
    <flux:navbar.item href="/inbox" badge="3" badge:color="red">
        Inbox
    </flux:navbar.item>
    <flux:navbar.item href="/messages" badge="5" badge:color="blue">
        Messages
    </flux:navbar.item>
</flux:navbar>
```

**Navbar with trailing icons:**
```blade
<flux:navbar>
    <flux:navbar.item href="/" icon="home" current>Home</flux:navbar.item>
    <flux:navbar.item href="/download" icon="arrow-down" icon:trailing="check">
        Download
    </flux:navbar.item>
</flux:navbar>
```

### Current Page Detection

Flux automatically detects the current page based on the `href` attribute:

```blade
<!-- User is on /about - this item will show as current -->
<flux:navbar>
    <flux:navbar.item href="/">Home</flux:navbar.item>
    <flux:navbar.item href="/about" data-current>About</flux:navbar.item>
</flux:navbar>
```

**Manual override:**
```blade
<flux:navbar>
    <flux:navbar.item href="/" :current="request()->is('/')">Home</flux:navbar.item>
</flux:navbar>
```

### Styling

Navbar items apply the `data-current` attribute when active:

```css
[data-flux-navbar] [data-current] {
    /* Styles for active navbar items */
}
```

Customise appearance:
```blade
<flux:navbar class="bg-blue-600">
    <flux:navbar.item href="/" class="text-white">Home</flux:navbar.item>
</flux:navbar>
```

---

## flux:navlist - Vertical Sidebar Navigation

The navlist component provides vertical navigation for sidebars with collapsible groups.

### Basic Structure

```blade
<flux:navlist>
    <flux:navlist.item href="/">Dashboard</flux:navlist.item>
    <flux:navlist.item href="/posts">Posts</flux:navlist.item>
    <flux:navlist.item href="/settings">Settings</flux:navlist.item>
</flux:navlist>
```

### flux:navlist Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `class` | string | - | Additional CSS classes |

### flux:navlist.item Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `href` | string | - | Destination URL |
| `current` | boolean | auto-detected | Mark as active page |
| `icon` | string | - | Leading icon name (Heroicon) |
| `badge` | string, boolean | - | Badge content (count, label, etc.) |
| `badge:color` | string | - | Badge colour |
| `badge:variant` | string | solid | Badge style: solid or outline |
| `class` | string | - | Additional CSS classes |

### flux:navlist.group Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `heading` | string | - | Group title/section heading |
| `expandable` | boolean | false | Enable collapse/expand toggle |
| `expanded` | boolean | false | Initial expanded state |
| `class` | string | - | Additional CSS classes |

### Usage Examples

**Basic vertical navigation:**
```blade
<flux:navlist>
    <flux:navlist.item href="/" icon="home" current>Dashboard</flux:navlist.item>
    <flux:navlist.item href="/posts" icon="document">Posts</flux:navlist.item>
    <flux:navlist.item href="/pages" icon="document-duplicate">Pages</flux:navlist.item>
    <flux:navlist.item href="/settings" icon="cog-6-tooth">Settings</flux:navlist.item>
</flux:navlist>
```

**With collapsible groups:**
```blade
<flux:navlist>
    <flux:navlist.item href="/" icon="home" current>Dashboard</flux:navlist.item>

    <flux:navlist.group heading="Content" expandable>
        <flux:navlist.item href="/posts" icon="document">Posts</flux:navlist.item>
        <flux:navlist.item href="/pages" icon="document-duplicate">Pages</flux:navlist.item>
    </flux:navlist.group>

    <flux:navlist.group heading="Settings" expandable>
        <flux:navlist.item href="/site-settings" icon="cog-6-tooth">Site Settings</flux:navlist.item>
        <flux:navlist.item href="/users" icon="users">Users</flux:navlist.item>
    </flux:navlist.group>
</flux:navlist>
```

**With badges:**
```blade
<flux:navlist>
    <flux:navlist.item href="/" icon="home" current>Dashboard</flux:navlist.item>
    <flux:navlist.item href="/inbox" icon="inbox" badge="3" badge:color="red">
        Inbox
    </flux:navlist.item>
    <flux:navlist.item href="/notifications" icon="bell" badge="5" badge:color="blue">
        Notifications
    </flux:navlist.item>
</flux:navlist>
```

**With icons and advanced grouping:**
```blade
<flux:navlist>
    <flux:navlist.item href="/" icon="home" current>Home</flux:navlist.item>

    <flux:navlist.group heading="User Management" expandable expanded>
        <flux:navlist.item href="/users" icon="users">Users</flux:navlist.item>
        <flux:navlist.item href="/roles" icon="shield-check">Roles</flux:navlist.item>
        <flux:navlist.item href="/permissions" icon="lock-closed">Permissions</flux:navlist.item>
    </flux:navlist.group>

    <flux:navlist.group heading="Content" expandable>
        <flux:navlist.item href="/categories" icon="tag">Categories</flux:navlist.item>
        <flux:navlist.item href="/tags" icon="hashtag">Tags</flux:navlist.item>
    </flux:navlist.group>
</flux:navlist>
```

### Styling

Navlist items use `data-current` attribute when active:

```css
[data-flux-navlist] [data-current] {
    /* Active state styling */
}
```

Apply custom styling:
```blade
<flux:navlist class="bg-zinc-50 dark:bg-zinc-900">
    <flux:navlist.item href="/" class="font-semibold">Home</flux:navlist.item>
</flux:navlist>
```

---

## flux:navmenu - Navigation Dropdown Menus

The navmenu component creates dropdown menus within navbar or navlist items.

### Basic Structure

```blade
<flux:navbar.item>
    <flux:navmenu>
        <flux:button>Menu</flux:button>
        <flux:menu>
            <flux:menu.item href="/option1">Option 1</flux:menu.item>
            <flux:menu.item href="/option2">Option 2</flux:menu.item>
        </flux:menu>
    </flux:navmenu>
</flux:navbar.item>
```

### flux:navmenu Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `position` | string | bottom | Dropdown position: top, right, bottom, left |
| `align` | string | start | Alignment: start, center, end |
| `offset` | number | 0 | Pixel offset from trigger |
| `gap` | number | 4 | Gap between trigger and menu |
| `class` | string | - | Additional CSS classes |

### Child Components

**flux:menu** - Container for menu items
**flux:menu.item** - Individual menu link
**flux:menu.separator** - Visual divider
**flux:menu.submenu** - Grouped items with heading

### flux:menu.item Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `href` | string | - | Link destination |
| `icon` | string | - | Leading icon name |
| `icon:trailing` | string | - | Trailing icon name |
| `icon:variant` | string | outline | Icon style: outline, solid, mini, micro |
| `kbd` | string | - | Keyboard shortcut hint |
| `suffix` | string | - | Trailing text label |
| `variant` | string | default | Style: default or danger |
| `disabled` | boolean | false | Disable interaction |
| `keep-open` | boolean | false | Don't close menu on click |
| `class` | string | - | Additional CSS classes |

### Usage Examples

**Simple navigation menu:**
```blade
<flux:navbar.item>
    <flux:navmenu>
        <flux:button>Help</flux:button>
        <flux:menu>
            <flux:menu.item href="/docs">Documentation</flux:menu.item>
            <flux:menu.item href="/faq">FAQ</flux:menu.item>
            <flux:menu.item href="/contact">Contact Support</flux:menu.item>
        </flux:menu>
    </flux:navmenu>
</flux:navbar.item>
```

**With keyboard shortcuts:**
```blade
<flux:navmenu>
    <flux:button>File</flux:button>
    <flux:menu>
        <flux:menu.item href="/new" kbd="⌘N">New</flux:menu.item>
        <flux:menu.item href="/open" kbd="⌘O">Open</flux:menu.item>
        <flux:menu.item href="/save" kbd="⌘S">Save</flux:menu.item>
        <flux:menu.separator />
        <flux:menu.item href="/quit" kbd="⌘Q" variant="danger">Quit</flux:menu.item>
    </flux:menu>
</flux:navmenu>
```

**With icons:**
```blade
<flux:navmenu>
    <flux:button icon="bars-3">Menu</flux:button>
    <flux:menu>
        <flux:menu.item href="/profile" icon="user">My Profile</flux:menu.item>
        <flux:menu.item href="/settings" icon="cog-6-tooth">Settings</flux:menu.item>
        <flux:menu.separator />
        <flux:menu.item href="/logout" icon="arrow-left-on-rectangle" variant="danger">
            Logout
        </flux:menu.item>
    </flux:menu>
</flux:navmenu>
```

**In sidebar:**
```blade
<flux:navlist.item>
    <flux:navmenu position="right" align="start">
        <flux:button variant="ghost" icon="ellipsis-vertical" square />
        <flux:menu>
            <flux:menu.item href="/edit">Edit</flux:menu.item>
            <flux:menu.item href="/delete" variant="danger">Delete</flux:menu.item>
        </flux:menu>
    </flux:navmenu>
</flux:navlist.item>
```

### Menu Separators

Divide menu sections visually:

```blade
<flux:menu>
    <flux:menu.item href="/new">New</flux:menu.item>
    <flux:menu.item href="/open">Open</flux:menu.item>
    <flux:menu.separator />
    <flux:menu.item href="/save">Save</flux:menu.item>
    <flux:menu.item href="/export">Export</flux:menu.item>
</flux:menu>
```

### Submenu Groups

Organise related items:

```blade
<flux:menu>
    <flux:menu.submenu heading="File">
        <flux:menu.item href="/new">New</flux:menu.item>
        <flux:menu.item href="/open">Open</flux:menu.item>
    </flux:menu.submenu>

    <flux:menu.submenu heading="Edit">
        <flux:menu.item href="/cut">Cut</flux:menu.item>
        <flux:menu.item href="/copy">Copy</flux:menu.item>
        <flux:menu.item href="/paste">Paste</flux:menu.item>
    </flux:menu.submenu>
</flux:menu>
```

---

## Integration with Full Layouts

### Complete Navbar + Sidebar Example

```blade
<div class="flex h-screen">
    <flux:sidebar sticky collapsible="mobile">
        <flux:sidebar.header>
            <flux:sidebar.brand>Host Hub</flux:sidebar.brand>
            <flux:sidebar.collapse />
        </flux:sidebar.header>

        <flux:sidebar.nav>
            <flux:sidebar.item href="/" icon="home" current>
                Dashboard
            </flux:sidebar.item>
            <flux:sidebar.group heading="Social" expandable>
                <flux:sidebar.item href="/social/accounts" icon="globe-alt">
                    Accounts
                </flux:sidebar.item>
                <flux:sidebar.item href="/social/posts" icon="document">
                    Posts
                </flux:sidebar.item>
            </flux:sidebar.group>
        </flux:sidebar.nav>

        <flux:sidebar.spacer />

        <flux:sidebar.profile>
            <flux:avatar initials="JD" />
        </flux:sidebar.profile>
    </flux:sidebar>

    <div class="flex-1 flex flex-col">
        <flux:header class="border-b">
            <div class="flex items-center justify-between px-6 py-4">
                <flux:heading level="1" size="lg">Dashboard</flux:heading>

                <div class="flex gap-2">
                    <flux:navmenu>
                        <flux:button variant="ghost" icon="bell">
                            Notifications
                        </flux:button>
                        <flux:menu>
                            <flux:menu.item href="/notifications">
                                View All
                            </flux:menu.item>
                        </flux:menu>
                    </flux:navmenu>

                    <flux:navmenu>
                        <flux:button variant="ghost" icon="user-circle" square />
                        <flux:menu align="end">
                            <flux:menu.item href="/profile" icon="user">
                                Profile
                            </flux:menu.item>
                            <flux:menu.item href="/settings" icon="cog-6-tooth">
                                Settings
                            </flux:menu.item>
                            <flux:menu.separator />
                            <flux:menu.item
                                href="/logout"
                                icon="arrow-left-on-rectangle"
                                variant="danger"
                            >
                                Logout
                            </flux:menu.item>
                        </flux:menu>
                    </flux:navmenu>
                </div>
            </div>
        </flux:header>

        <main class="flex-1 overflow-auto p-6">
            <!-- Page content -->
        </main>
    </div>
</div>
```

---

## Styling & Customisation

### Applying Tailwind Classes

```blade
<flux:navbar class="bg-blue-600 dark:bg-blue-900">
    <flux:navbar.item href="/" class="text-white" current>Home</flux:navbar.item>
</flux:navbar>
```

### Dark Mode Support

All navigation components automatically support dark mode:

```blade
<flux:navlist class="bg-zinc-50 dark:bg-zinc-900 border-r border-zinc-200 dark:border-zinc-800">
    <flux:navlist.item href="/" icon="home" current>Home</flux:navlist.item>
</flux:navlist>
```

### Active State Styling

Customise active item appearance:

```css
[data-flux-navbar] [data-current],
[data-flux-navlist] [data-current] {
    @apply text-accent font-semibold bg-accent-foreground/10;
}
```

### Positioning Dropdowns

```blade
<!-- Top-aligned dropdown -->
<flux:navmenu position="top" align="center">
    <flux:button>Menu</flux:button>
    <flux:menu>...</flux:menu>
</flux:navmenu>

<!-- Right-aligned dropdown -->
<flux:navmenu position="right" align="end">
    <flux:button>Actions</flux:button>
    <flux:menu>...</flux:menu>
</flux:navmenu>
```

---

## Key Features Summary

| Feature | flux:navbar | flux:navlist | flux:navmenu |
|---------|------------|-------------|------------|
| Auto-detection of current page | Yes | Yes | No |
| Icons support | Yes | Yes | Yes |
| Badges | Yes | Yes | No |
| Collapsible groups | No | Yes | No |
| Keyboard shortcuts display | No | No | Yes |
| Dropdown menus | With navmenu | With navmenu | Yes (native) |
| Dark mode | Yes | Yes | Yes |
| Custom styling | Yes | Yes | Yes |

---

Last updated: January 2026
