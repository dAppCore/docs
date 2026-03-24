# flux:breadcrumbs

Help users navigate and understand their location within an application.

## Components

### flux:breadcrumbs

Container for breadcrumb items.

| Slot | Description |
|------|-------------|
| `default` | Contains the breadcrumb items |

### flux:breadcrumbs.item

Individual breadcrumb link or text.

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `href` | string | - | URL the item links to; omit for non-clickable text |
| `icon` | string | - | Icon name to display before text |
| `icon:variant` | string | mini | `outline`, `solid`, `mini`, `micro` |
| `separator` | string | chevron-right | Separator icon: `chevron-right`, `slash` |

---

## Basic Usage

```blade
<flux:breadcrumbs>
    <flux:breadcrumbs.item href="/">Home</flux:breadcrumbs.item>
    <flux:breadcrumbs.item href="/blog">Blog</flux:breadcrumbs.item>
    <flux:breadcrumbs.item>Post</flux:breadcrumbs.item>
</flux:breadcrumbs>
```

## With Slash Separators

```blade
<flux:breadcrumbs>
    <flux:breadcrumbs.item href="/" separator="slash">Home</flux:breadcrumbs.item>
    <flux:breadcrumbs.item href="/blog" separator="slash">Blog</flux:breadcrumbs.item>
    <flux:breadcrumbs.item separator="slash">Post</flux:breadcrumbs.item>
</flux:breadcrumbs>
```

## With Home Icon

```blade
<flux:breadcrumbs>
    <flux:breadcrumbs.item href="/" icon="home" />
    <flux:breadcrumbs.item href="/blog">Blog</flux:breadcrumbs.item>
    <flux:breadcrumbs.item>Post</flux:breadcrumbs.item>
</flux:breadcrumbs>
```

## With Ellipsis

For deep navigation, use ellipsis to collapse middle items:

```blade
<flux:breadcrumbs>
    <flux:breadcrumbs.item href="/" icon="home" />
    <flux:breadcrumbs.item icon="ellipsis-horizontal" />
    <flux:breadcrumbs.item>Post</flux:breadcrumbs.item>
</flux:breadcrumbs>
```

## With Ellipsis Dropdown

Expandable ellipsis showing hidden items:

```blade
<flux:breadcrumbs>
    <flux:breadcrumbs.item href="/" icon="home" />
    <flux:breadcrumbs.item>
        <flux:dropdown>
            <flux:button icon="ellipsis-horizontal" variant="ghost" size="sm" />
            <flux:navmenu>
                <flux:navmenu.item href="/clients">Client</flux:navmenu.item>
                <flux:navmenu.item href="/clients/team" icon="arrow-turn-down-right">Team</flux:navmenu.item>
                <flux:navmenu.item href="/clients/team/user" icon="arrow-turn-down-right">User</flux:navmenu.item>
            </flux:navmenu>
        </flux:dropdown>
    </flux:breadcrumbs.item>
    <flux:breadcrumbs.item>Post</flux:breadcrumbs.item>
</flux:breadcrumbs>
```

## Dynamic Breadcrumbs

```blade
<flux:breadcrumbs>
    <flux:breadcrumbs.item href="/" icon="home" />

    @foreach ($breadcrumbs as $crumb)
        @if ($loop->last)
            <flux:breadcrumbs.item>{{ $crumb['name'] }}</flux:breadcrumbs.item>
        @else
            <flux:breadcrumbs.item href="{{ $crumb['url'] }}">{{ $crumb['name'] }}</flux:breadcrumbs.item>
        @endif
    @endforeach
</flux:breadcrumbs>
```

## In Header

```blade
<flux:header>
    <flux:breadcrumbs>
        <flux:breadcrumbs.item href="/" icon="home" />
        <flux:breadcrumbs.item href="/settings">Settings</flux:breadcrumbs.item>
        <flux:breadcrumbs.item>Profile</flux:breadcrumbs.item>
    </flux:breadcrumbs>

    <flux:spacer />

    <flux:button variant="primary">Save</flux:button>
</flux:header>
```

## Longer Path

```blade
<flux:breadcrumbs>
    <flux:breadcrumbs.item href="/" icon="home" />
    <flux:breadcrumbs.item href="/products">Products</flux:breadcrumbs.item>
    <flux:breadcrumbs.item href="/products/electronics">Electronics</flux:breadcrumbs.item>
    <flux:breadcrumbs.item href="/products/electronics/phones">Phones</flux:breadcrumbs.item>
    <flux:breadcrumbs.item>iPhone 15</flux:breadcrumbs.item>
</flux:breadcrumbs>
```
