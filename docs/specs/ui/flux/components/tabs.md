# flux:tabs

Tabbed content organisation with multiple visual variants.

## Props

### flux:tabs

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Binds active tab to Livewire property |
| `variant` | string | default | `default`, `segmented`, `pills` |
| `size` | string | base | `base`, `sm` |
| `scrollable` | boolean | false | Enable horizontal scrolling |
| `scrollable:scrollbar` | string | - | `hide` |
| `scrollable:fade` | boolean | false | Fade trailing edge |

### flux:tab

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `name` | string | - | Unique identifier (matches panel) |
| `icon` | string | - | Leading icon |
| `icon:trailing` | string | - | Trailing icon |
| `icon:variant` | string | outline | `outline`, `solid`, `mini`, `micro` |
| `selected` | boolean | false | Pre-select tab |
| `action` | boolean | false | Convert to action button |
| `accent` | boolean | false | Accent colour styling |
| `disabled` | boolean | false | Disable interaction |

### flux:tab.panel

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `name` | string | - | Matches tab identifier |
| `selected` | boolean | false | Pre-select panel |

## Basic Usage

```blade
<flux:tab.group>
    <flux:tabs wire:model="tab">
        <flux:tab name="profile">Profile</flux:tab>
        <flux:tab name="account">Account</flux:tab>
        <flux:tab name="billing">Billing</flux:tab>
    </flux:tabs>

    <flux:tab.panel name="profile">
        Profile content...
    </flux:tab.panel>
    <flux:tab.panel name="account">
        Account content...
    </flux:tab.panel>
    <flux:tab.panel name="billing">
        Billing content...
    </flux:tab.panel>
</flux:tab.group>
```

## With Icons

```blade
<flux:tabs wire:model="tab">
    <flux:tab name="profile" icon="user">Profile</flux:tab>
    <flux:tab name="account" icon="cog-6-tooth">Account</flux:tab>
    <flux:tab name="billing" icon="credit-card">Billing</flux:tab>
</flux:tabs>
```

## Segmented Variant

```blade
<flux:tabs wire:model="view" variant="segmented">
    <flux:tab name="list">List</flux:tab>
    <flux:tab name="board">Board</flux:tab>
    <flux:tab name="timeline">Timeline</flux:tab>
</flux:tabs>
```

## Segmented with Icons

```blade
<flux:tabs wire:model="view" variant="segmented" size="sm">
    <flux:tab name="list" icon="list-bullet">List</flux:tab>
    <flux:tab name="grid" icon="squares-2x2">Grid</flux:tab>
</flux:tabs>
```

## Pills Variant

```blade
<flux:tabs wire:model="filter" variant="pills">
    <flux:tab name="all">All</flux:tab>
    <flux:tab name="active">Active</flux:tab>
    <flux:tab name="archived">Archived</flux:tab>
</flux:tabs>
```

## Scrollable

```blade
<flux:tabs wire:model="tab" scrollable scrollable:fade>
    <flux:tab name="overview">Overview</flux:tab>
    <flux:tab name="analytics">Analytics</flux:tab>
    <flux:tab name="reports">Reports</flux:tab>
    <flux:tab name="settings">Settings</flux:tab>
    <flux:tab name="integrations">Integrations</flux:tab>
    <flux:tab name="billing">Billing</flux:tab>
</flux:tabs>
```

## Hide Scrollbar

```blade
<flux:tabs scrollable scrollable:scrollbar="hide">
    {{-- Tabs --}}
</flux:tabs>
```

## Action Tab (Dynamic)

```blade
<flux:tab.group>
    <flux:tabs>
        @foreach ($tabs as $id => $tab)
            <flux:tab :name="$id">{{ $tab }}</flux:tab>
        @endforeach

        <flux:tab icon="plus" wire:click="addTab" action>
            Add tab
        </flux:tab>
    </flux:tabs>

    @foreach ($tabs as $id => $tab)
        <flux:tab.panel :name="$id">
            Content for {{ $tab }}
        </flux:tab.panel>
    @endforeach
</flux:tab.group>
```

## Pre-selected Tab

```blade
<flux:tabs>
    <flux:tab name="profile">Profile</flux:tab>
    <flux:tab name="account" selected>Account</flux:tab>
</flux:tabs>

<flux:tab.panel name="profile">...</flux:tab.panel>
<flux:tab.panel name="account" selected>...</flux:tab.panel>
```

## Disabled Tab

```blade
<flux:tabs>
    <flux:tab name="free">Free</flux:tab>
    <flux:tab name="pro" disabled>Pro (Coming soon)</flux:tab>
</flux:tabs>
```
