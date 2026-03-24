# Flux UI - Layouts

Complete documentation for Flux layout components: Sidebar and Header.

## flux:sidebar - Persistent Sidebar Layout

The sidebar layout keeps application content prominent while providing persistent navigation. It supports responsive behaviour, sticky positioning, and mobile collapse.

### Basic Structure

```blade
<flux:sidebar>
    <flux:sidebar.header>
        <flux:sidebar.brand>App Name</flux:sidebar.brand>
    </flux:sidebar.header>

    <flux:sidebar.nav>
        <flux:sidebar.item href="/" icon="home" current>Dashboard</flux:sidebar.item>
    </flux:sidebar.nav>

    <flux:sidebar.profile>
        <flux:avatar initials="JD" />
    </flux:sidebar.profile>
</flux:sidebar>
```

### flux:sidebar Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `sticky` | boolean | false | Makes sidebar fixed during scrolling |
| `collapsible` | boolean, "mobile" | false | Enables collapse: true (all sizes), "mobile" (mobile only), false (disabled) |
| `breakpoint` | number | 1024 | Responsive threshold in pixels (lg breakpoint) |
| `persist` | boolean | true | Save collapsed state to localStorage |
| `class` | string | - | Additional CSS classes |

### Child Components

| Component | Purpose |
|-----------|---------|
| `flux:sidebar.header` | Top section for branding and controls |
| `flux:sidebar.brand` | Logo and company name display |
| `flux:sidebar.collapse` | Toggle button for sidebar state |
| `flux:sidebar.search` | Search input field |
| `flux:sidebar.nav` | Container for navigation items and groups |
| `flux:sidebar.item` | Individual navigation entry |
| `flux:sidebar.group` | Grouped navigation items with expand/collapse |
| `flux:sidebar.spacer` | Vertical spacer (pushes content down) |
| `flux:sidebar.profile` | User avatar and name display |
| `flux:sidebar.toggle` | Mobile header toggle button |

### flux:sidebar.item Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `href` | string | - | Link destination |
| `icon` | string | - | Icon name (Heroicon) |
| `badge` | string, number | - | Numeric indicator |
| `current` | boolean | auto-detected | Mark as active page |
| `tooltip` | string | auto-populated | Hover text (shows when collapsed) |
| `class` | string | - | Additional CSS classes |

### flux:sidebar.group Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `heading` | string | - | Group title/section heading |
| `expandable` | boolean | false | Enable collapse/expand toggle |
| `expanded` | boolean | false | Initial expanded state |
| `icon` | string | - | Group icon (optional) |
| `class` | string | - | Additional CSS classes |

### Basic Sidebar Example

```blade
<div class="flex h-screen">
    <flux:sidebar sticky collapsible="mobile">
        <flux:sidebar.header>
            <flux:sidebar.brand>MyApp</flux:sidebar.brand>
            <flux:sidebar.collapse />
        </flux:sidebar.header>

        <flux:sidebar.nav>
            <flux:sidebar.item href="/" icon="home" current>
                Dashboard
            </flux:sidebar.item>
            <flux:sidebar.item href="/posts" icon="document">
                Posts
            </flux:sidebar.item>
            <flux:sidebar.item href="/settings" icon="cog-6-tooth">
                Settings
            </flux:sidebar.item>
        </flux:sidebar.nav>

        <flux:sidebar.spacer />

        <flux:sidebar.profile>
            <flux:avatar src="/avatar.jpg" />
        </flux:sidebar.profile>
    </flux:sidebar>

    <main class="flex-1 overflow-auto">
        <!-- Page content -->
    </main>
</div>
```

### Sidebar with Groups

```blade
<flux:sidebar sticky collapsible>
    <flux:sidebar.header>
        <flux:sidebar.brand>Host Hub</flux:sidebar.brand>
        <flux:sidebar.collapse />
    </flux:sidebar.header>

    <flux:sidebar.nav>
        <flux:sidebar.item href="/dashboard" icon="home" current>
            Dashboard
        </flux:sidebar.item>

        <flux:sidebar.group heading="Social" expandable expanded>
            <flux:sidebar.item href="/social/accounts" icon="globe-alt">
                Accounts
            </flux:sidebar.item>
            <flux:sidebar.item href="/social/posts" icon="document">
                Posts
            </flux:sidebar.item>
            <flux:sidebar.item href="/social/analytics" icon="chart-bar">
                Analytics
            </flux:sidebar.item>
        </flux:sidebar.group>

        <flux:sidebar.group heading="Tools" expandable>
            <flux:sidebar.item href="/tools/email" icon="envelope">
                Email
            </flux:sidebar.item>
            <flux:sidebar.item href="/tools/sms" icon="chat-bubble-left">
                SMS
            </flux:sidebar.item>
        </flux:sidebar.group>

        <flux:sidebar.group heading="Admin" expandable>
            <flux:sidebar.item href="/admin/users" icon="users">
                Users
            </flux:sidebar.item>
            <flux:sidebar.item href="/admin/settings" icon="cog-6-tooth">
                Settings
            </flux:sidebar.item>
        </flux:sidebar.group>
    </flux:sidebar.nav>

    <flux:sidebar.spacer />

    <flux:sidebar.profile>
        <flux:avatar initials="JD" />
        <div class="flex-1">
            <flux:text weight="semibold">John Doe</flux:text>
            <flux:text size="sm" variant="muted">john@example.com</flux:text>
        </div>
    </flux:sidebar.profile>
</flux:sidebar>
```

### Sidebar with Search

```blade
<flux:sidebar sticky collapsible="mobile">
    <flux:sidebar.header>
        <flux:sidebar.brand>App</flux:sidebar.brand>
        <flux:sidebar.collapse />
    </flux:sidebar.header>

    <flux:sidebar.search placeholder="Search..." />

    <flux:sidebar.nav>
        <!-- Navigation items -->
    </flux:sidebar.nav>
</flux:sidebar>
```

### Responsive Behaviour

The `collapsible` prop controls how the sidebar responds to viewport width:

**Option 1: Mobile-only collapse**
```blade
<!-- Collapses only on mobile (below 1024px), stays open on desktop -->
<flux:sidebar collapsible="mobile">
    <!-- Content -->
</flux:sidebar>
```

**Option 2: Always collapsible**
```blade
<!-- Can be collapsed at any viewport size -->
<flux:sidebar collapsible>
    <!-- Content -->
</flux:sidebar>
```

**Option 3: No collapse**
```blade
<!-- Never collapses, always visible -->
<flux:sidebar collapsible="false">
    <!-- Content -->
</flux:sidebar>
```

### Sticky Positioning

```blade
<!-- Sidebar stays fixed at top when scrolling -->
<flux:sidebar sticky collapsible="mobile">
    <!-- Content scrolls, sidebar stays visible -->
</flux:sidebar>
```

### Persisting Collapse State

```blade
<!-- User's collapse preference saved to localStorage (default) -->
<flux:sidebar collapsible persist>
    <!-- State persists across page refreshes -->
</flux:sidebar>

<!-- Disable persistence -->
<flux:sidebar collapsible :persist="false">
    <!-- State resets on page refresh -->
</flux:sidebar>
```

### Styling the Sidebar

Apply Tailwind classes for customisation:

```blade
<flux:sidebar sticky collapsible="mobile" class="bg-zinc-50 dark:bg-zinc-900 border-r border-zinc-200 dark:border-zinc-800">
    <!-- Content -->
</flux:sidebar>
```

### Dark Mode Support

Sidebar automatically handles dark mode:

```blade
<flux:sidebar class="dark:bg-zinc-900 dark:border-zinc-800">
    <flux:sidebar.item href="/" class="dark:hover:bg-zinc-800">
        Dashboard
    </flux:sidebar.item>
</flux:sidebar>
```

---

## flux:header - Top Navigation Header

The header component provides top navigation, typically paired with a sidebar for secondary navigation options.

### Basic Structure

```blade
<flux:header>
    <div class="flex items-center justify-between px-6 py-4">
        <flux:heading>Page Title</flux:heading>
        <div class="flex gap-2">
            <!-- Right-aligned content -->
        </div>
    </div>
</flux:header>
```

### Usage as Secondary Navigation

```blade
<flux:header class="border-b">
    <flux:navbar>
        <flux:navbar.item href="/dashboard">Dashboard</flux:navbar.item>
        <flux:navbar.item href="/analytics">Analytics</flux:navbar.item>
    </flux:navbar>
</flux:header>
```

### Complete Header with Actions

```blade
<flux:header class="border-b bg-white dark:bg-zinc-950">
    <div class="flex items-center justify-between px-6 py-4 gap-4">
        <!-- Left: Page title -->
        <div>
            <flux:heading level="1" size="lg">Dashboard</flux:heading>
            <flux:text size="sm" variant="muted">Welcome back!</flux:text>
        </div>

        <!-- Right: Actions and controls -->
        <div class="flex items-center gap-2 ml-auto">
            <flux:button variant="ghost" icon="bell">
                Notifications
            </flux:button>

            <flux:dropdown align="end">
                <flux:button variant="ghost" icon="user-circle" square />
                <flux:menu align="end">
                    <flux:menu.item href="/profile" icon="user">
                        My Profile
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
                        Sign Out
                    </flux:menu.item>
                </flux:menu>
            </flux:dropdown>
        </div>
    </div>
</flux:header>
```

---

## Complete Layout Examples

### Example 1: Admin Dashboard Layout

```blade
<div class="flex h-screen bg-zinc-50 dark:bg-zinc-950">
    <!-- Sidebar -->
    <flux:sidebar sticky collapsible="mobile" class="w-64">
        <flux:sidebar.header>
            <flux:sidebar.brand class="font-bold text-lg">
                HostHub Admin
            </flux:sidebar.brand>
            <flux:sidebar.collapse />
        </flux:sidebar.header>

        <flux:sidebar.search placeholder="Search menu..." />

        <flux:sidebar.nav>
            <flux:sidebar.item href="/admin" icon="home" current>
                Dashboard
            </flux:sidebar.item>

            <flux:sidebar.group heading="Management" expandable expanded>
                <flux:sidebar.item href="/admin/users" icon="users">
                    Users
                </flux:sidebar.item>
                <flux:sidebar.item href="/admin/workspaces" icon="building-office-2">
                    Workspaces
                </flux:sidebar.item>
                <flux:sidebar.item href="/admin/billing" icon="credit-card">
                    Billing
                </flux:sidebar.item>
            </flux:sidebar.group>

            <flux:sidebar.group heading="System" expandable>
                <flux:sidebar.item href="/admin/logs" icon="document-text">
                    Logs
                </flux:sidebar.item>
                <flux:sidebar.item href="/admin/settings" icon="cog-6-tooth">
                    Settings
                </flux:sidebar.item>
            </flux:sidebar.group>
        </flux:sidebar.nav>

        <flux:sidebar.spacer />

        <flux:sidebar.profile>
            <flux:avatar initials="AD" />
            <div class="flex-1">
                <flux:text weight="semibold" size="sm">Admin User</flux:text>
                <flux:text size="xs" variant="muted">admin@host.uk.com</flux:text>
            </div>
        </flux:sidebar.profile>
    </flux:sidebar>

    <!-- Main Content -->
    <div class="flex-1 flex flex-col overflow-hidden">
        <!-- Header -->
        <flux:header class="border-b bg-white dark:bg-zinc-900 shadow-sm">
            <div class="flex items-center justify-between px-6 py-4">
                <div>
                    <flux:heading level="1" size="lg">Dashboard</flux:heading>
                    <flux:text size="sm" variant="muted">System overview</flux:text>
                </div>

                <div class="flex items-center gap-3">
                    <flux:button variant="ghost" icon="bell" />
                    <flux:dropdown align="end">
                        <flux:button variant="ghost" icon="user-circle" square />
                        <flux:menu align="end">
                            <flux:menu.item href="/profile" icon="user">
                                Profile
                            </flux:menu.item>
                            <flux:menu.separator />
                            <flux:menu.item href="/logout" variant="danger">
                                Sign Out
                            </flux:menu.item>
                        </flux:menu>
                    </flux:dropdown>
                </div>
            </div>
        </flux:header>

        <!-- Page Content -->
        <main class="flex-1 overflow-auto p-6">
            <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
                <flux:card>
                    <flux:heading level="3" size="sm">Total Users</flux:heading>
                    <flux:text size="2xl" weight="bold" class="mt-2">1,234</flux:text>
                </flux:card>
                <!-- More cards -->
            </div>
        </main>
    </div>
</div>
```

### Example 2: SaaS App Layout

```blade
<div class="flex h-screen">
    <!-- Sidebar Navigation -->
    <flux:sidebar sticky collapsible="mobile">
        <flux:sidebar.header>
            <flux:sidebar.brand class="font-bold">
                <flux:icon name="rocket-launch" class="inline mr-2" />
                SocialHost
            </flux:sidebar.brand>
        </flux:sidebar.header>

        <flux:sidebar.nav>
            <flux:sidebar.item href="/dashboard" icon="home" current>
                Dashboard
            </flux:sidebar.item>

            <flux:sidebar.group heading="Social" expandable expanded icon="globe-alt">
                <flux:sidebar.item href="/accounts" icon="squares-2x2">
                    Connected Accounts
                </flux:sidebar.item>
                <flux:sidebar.item href="/posts" icon="document" badge="3">
                    Posts
                </flux:sidebar.item>
                <flux:sidebar.item href="/calendar" icon="calendar">
                    Calendar
                </flux:sidebar.item>
                <flux:sidebar.item href="/analytics" icon="chart-bar">
                    Analytics
                </flux:sidebar.item>
            </flux:sidebar.group>

            <flux:sidebar.group heading="Team" expandable>
                <flux:sidebar.item href="/members" icon="users">
                    Members
                </flux:sidebar.item>
                <flux:sidebar.item href="/permissions" icon="shield-check">
                    Permissions
                </flux:sidebar.item>
            </flux:sidebar.group>
        </flux:sidebar.nav>

        <flux:sidebar.spacer />

        <flux:sidebar.profile>
            <flux:avatar src="/avatars/user.jpg" />
            <div class="flex-1 min-w-0">
                <flux:text weight="semibold" class="truncate">Sarah Johnson</flux:text>
                <flux:text size="xs" variant="muted" class="truncate">
                    sarah@company.com
                </flux:text>
            </div>
        </flux:sidebar.profile>
    </flux:sidebar>

    <!-- Main Content Area -->
    <div class="flex-1 flex flex-col">
        <!-- Top Bar -->
        <flux:header class="border-b px-6 py-4 flex items-center justify-between">
            <div>
                <flux:heading level="1" size="lg">Social Accounts</flux:heading>
            </div>

            <div class="flex gap-2">
                <flux:button icon="plus">Add Account</flux:button>
                <flux:dropdown>
                    <flux:button variant="ghost" icon="ellipsis-vertical" square />
                    <flux:menu>
                        <flux:menu.item href="/settings" icon="cog-6-tooth">
                            Settings
                        </flux:menu.item>
                    </flux:menu>
                </flux:dropdown>
            </div>
        </flux:header>

        <!-- Page Content -->
        <main class="flex-1 overflow-auto p-6">
            <!-- Content goes here -->
        </main>
    </div>
</div>
```

### Example 3: Minimal Blog Layout

```blade
<div class="flex h-screen flex-col">
    <!-- Header Navigation -->
    <flux:header class="border-b bg-white dark:bg-zinc-900">
        <flux:navbar class="px-6">
            <flux:navbar.item href="/" icon="home" current>Home</flux:navbar.item>
            <flux:navbar.item href="/blog" icon="document">Blog</flux:navbar.item>
            <flux:navbar.item href="/about" icon="information-circle">About</flux:navbar.item>
        </flux:navbar>
    </flux:header>

    <!-- Main Content -->
    <main class="flex-1 overflow-auto p-6">
        <!-- Content -->
    </main>
</div>
```

---

## Layout Props Summary

| Component | Key Props | Usage |
|-----------|-----------|-------|
| `flux:sidebar` | `sticky`, `collapsible`, `breakpoint`, `persist` | Vertical navigation |
| `flux:sidebar.item` | `href`, `icon`, `badge`, `current` | Navigation links |
| `flux:sidebar.group` | `heading`, `expandable`, `expanded` | Grouped sections |
| `flux:header` | (container) | Top bar for content |

---

## Responsive Design

### Mobile-First Approach

```blade
<flux:sidebar collapsible="mobile">
    <!-- Collapsed on mobile (< 1024px) -->
    <!-- Expanded on desktop (>= 1024px) -->
</flux:sidebar>
```

### Custom Breakpoint

```blade
<!-- Use sm breakpoint instead of lg -->
<flux:sidebar collapsible breakpoint="640">
    <!-- Collapses at 640px (sm breakpoint) -->
</flux:sidebar>
```

---

## Styling Patterns

### Apply Classes to Sections

```blade
<!-- Styling the header -->
<flux:sidebar.header class="bg-blue-600 text-white">
    <flux:sidebar.brand class="font-bold text-lg">App</flux:sidebar.brand>
</flux:sidebar.header>

<!-- Styling navigation items -->
<flux:sidebar.item
    href="/"
    class="hover:bg-zinc-100 dark:hover:bg-zinc-800 rounded-lg"
>
    Dashboard
</flux:sidebar.item>
```

### Custom Width

```blade
<!-- Default sidebar width: 16rem (256px) -->
<!-- Adjust with classes or CSS variables -->
<flux:sidebar class="w-80">
    <!-- Wider sidebar -->
</flux:sidebar>
```

---

Last updated: January 2026
