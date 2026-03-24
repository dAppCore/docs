# flux:sidebar

Vertical navigation layout that keeps application content prominent. Supports sticky positioning, collapsibility, and responsive behaviour.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `sticky` | boolean | false | Makes sidebar fixed during scrolling |
| `collapsible` | string\|boolean | false | `"mobile"` (mobile-only), `true` (mobile+desktop), `false` (never) |
| `breakpoint` | string\|number | 1024px | Viewport breakpoint. Accepts "1024px" or "64rem" |
| `stashable` | boolean | - | **Deprecated**: use `collapsible="mobile"` instead |
| `persist` | boolean | true | Saves collapsed state to localStorage |

## Slots

| Slot | Description |
|------|-------------|
| `default` | Main sidebar content |

---

## Child Components

### flux:sidebar.header

Container for brand and collapse controls.

```blade
<flux:sidebar.header>
    <flux:sidebar.brand name="Acme Inc." logo="/logo.png" href="#" />
    <flux:sidebar.collapse class="lg:hidden" />
</flux:sidebar.header>
```

### flux:sidebar.brand

Logo and company name display.

| Prop | Type | Description |
|------|------|-------------|
| `href` | string | Navigation URL |
| `logo` | string | Logo image URL |
| `logo:dark` | string | Dark mode logo URL |
| `name` | string | Brand name text |

```blade
<flux:sidebar.brand
    name="Acme Inc."
    logo="/logo.svg"
    logo:dark="/logo-dark.svg"
    href="/"
/>
```

### flux:sidebar.collapse

Toggle button for expanding/collapsing sidebar.

| Prop | Type | Description |
|------|------|-------------|
| `inset` | string | Position: "left", "right", "top", "bottom" or combinations |
| `tooltip` | string | Hover text (default: "Toggle sidebar") |

```blade
<flux:sidebar.collapse inset="left" tooltip="Collapse navigation" />
```

### flux:sidebar.search

Search input field.

| Prop | Type | Description |
|------|------|-------------|
| `placeholder` | string | Input placeholder text |

```blade
<flux:sidebar.search placeholder="Search..." />
```

### flux:sidebar.nav

Navigation container.

```blade
<flux:sidebar.nav>
    <flux:sidebar.item icon="home" href="/" current>Home</flux:sidebar.item>
    <flux:sidebar.item icon="inbox" badge="12" href="/inbox">Inbox</flux:sidebar.item>
</flux:sidebar.nav>
```

### flux:sidebar.item

Individual navigation item.

| Prop | Type | Description |
|------|------|-------------|
| `href` | string | Navigation URL |
| `icon` | string | Heroicon identifier |
| `badge` | string\|number | Badge value |
| `current` | boolean | Marks active item |
| `tooltip` | string | Shown when collapsed (defaults to text content) |

```blade
<flux:sidebar.item
    href="/dashboard"
    icon="home"
    current
    tooltip="Dashboard"
>
    Dashboard
</flux:sidebar.item>
```

### flux:sidebar.group

Collapsible navigation group.

| Prop | Type | Description |
|------|------|-------------|
| `heading` | string | Group title |
| `expandable` | boolean | Allows collapse/expand |
| `icon` | string | Icon before heading (replaces chevron) |
| `expanded` | boolean | Initial state (default: true) |

```blade
<flux:sidebar.group expandable heading="Favorites" class="grid">
    <flux:sidebar.item href="/marketing">Marketing site</flux:sidebar.item>
    <flux:sidebar.item href="/android">Android app</flux:sidebar.item>
</flux:sidebar.group>
```

### flux:sidebar.spacer

Vertical spacing element.

```blade
<flux:sidebar.spacer />
```

### flux:sidebar.profile

User profile display (typically bottom of sidebar).

| Prop | Type | Description |
|------|------|-------------|
| `avatar` | string | User image URL |
| `name` | string | Display name |

```blade
<flux:sidebar.profile avatar="/avatar.jpg" name="John Doe" />
```

### flux:sidebar.toggle

Mobile header toggle button.

| Prop | Type | Description |
|------|------|-------------|
| `icon` | string | Icon name (e.g., "bars-2", "x-mark") |
| `inset` | string | Position (e.g., "left") |

```blade
<flux:sidebar.toggle icon="bars-2" inset="left" class="lg:hidden" />
```

### flux:main

Primary content container.

| Prop | Type | Description |
|------|------|-------------|
| `container` | boolean | Constrains content to container width |

```blade
<flux:main container>
    {{ $slot }}
</flux:main>
```

---

## Complete Examples

### Basic Sidebar

```blade
<flux:sidebar sticky collapsible="mobile">
    <flux:sidebar.header>
        <flux:sidebar.brand name="Acme Inc." logo="/logo.png" href="/" />
        <flux:sidebar.collapse class="lg:hidden" />
    </flux:sidebar.header>

    <flux:sidebar.search placeholder="Search..." />

    <flux:sidebar.nav>
        <flux:sidebar.item icon="home" href="/" current>Home</flux:sidebar.item>
        <flux:sidebar.item icon="inbox" badge="12" href="/inbox">Inbox</flux:sidebar.item>
        <flux:sidebar.item icon="document-text" href="/documents">Documents</flux:sidebar.item>
        <flux:sidebar.item icon="calendar" href="/calendar">Calendar</flux:sidebar.item>
    </flux:sidebar.nav>

    <flux:sidebar.spacer />

    <flux:sidebar.nav>
        <flux:sidebar.item icon="cog-6-tooth" href="/settings">Settings</flux:sidebar.item>
        <flux:sidebar.item icon="question-mark-circle" href="/help">Help</flux:sidebar.item>
    </flux:sidebar.nav>
</flux:sidebar>

<flux:main>
    {{ $slot }}
</flux:main>
```

### With Collapsible Groups

```blade
<flux:sidebar sticky collapsible>
    <flux:sidebar.header>
        <flux:sidebar.brand name="Acme Inc." logo="/logo.png" href="/" />
        <flux:sidebar.collapse />
    </flux:sidebar.header>

    <flux:sidebar.nav>
        <flux:sidebar.item icon="home" href="/" current>Home</flux:sidebar.item>

        <flux:sidebar.group expandable heading="Projects">
            <flux:sidebar.item href="/projects/marketing">Marketing</flux:sidebar.item>
            <flux:sidebar.item href="/projects/engineering">Engineering</flux:sidebar.item>
            <flux:sidebar.item href="/projects/design">Design</flux:sidebar.item>
        </flux:sidebar.group>

        <flux:sidebar.group expandable heading="Favorites" expanded>
            <flux:sidebar.item href="/favorites/reports">Reports</flux:sidebar.item>
            <flux:sidebar.item href="/favorites/analytics">Analytics</flux:sidebar.item>
        </flux:sidebar.group>
    </flux:sidebar.nav>

    <flux:sidebar.spacer />

    <flux:sidebar.profile avatar="/avatar.jpg" name="John Doe" />
</flux:sidebar>

<flux:main container>
    {{ $slot }}
</flux:main>
```

### Desktop and Mobile Collapsible

```blade
<flux:sidebar sticky collapsible class="bg-zinc-50 border-r dark:bg-zinc-900 dark:border-zinc-700">
    <flux:sidebar.header>
        <flux:sidebar.brand name="Acme" logo="/logo.svg" href="/" />
        <flux:sidebar.collapse />
    </flux:sidebar.header>

    <flux:sidebar.nav>
        <flux:sidebar.item icon="home" href="/" current tooltip="Dashboard">
            Dashboard
        </flux:sidebar.item>
        <flux:sidebar.item icon="users" href="/users" tooltip="Users">
            Users
        </flux:sidebar.item>
    </flux:sidebar.nav>
</flux:sidebar>
```

---

## Responsive Behaviour

| Mode | Usage |
|------|-------|
| Mobile-only collapse | `collapsible="mobile"` |
| Both mobile + desktop | `collapsible` or `collapsible="true"` |
| Disable collapse | `collapsible="false"` |
| Show on mobile only | `class="lg:hidden"` |
| Show on desktop only | `class="max-lg:hidden"` |

---

## Tooltips When Collapsed

When sidebar is collapsed, tooltips auto-populate from item text:

```blade
<flux:sidebar.item icon="home" href="/" tooltip="Dashboard">
    Dashboard
</flux:sidebar.item>
```

If `tooltip` prop is omitted, the inner text ("Dashboard") is used.

---

## Dark Mode

```blade
<flux:sidebar class="bg-zinc-50 border-r dark:bg-zinc-900 dark:border-zinc-700">
    ...
</flux:sidebar>
```
