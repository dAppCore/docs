# flux:profile

User profile display with avatar, name, and dropdown indicator.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `name` | string | - | User name displayed next to avatar |
| `avatar` | string | - | Avatar image URL |
| `avatar:name` | string | - | Name for initial generation |
| `avatar:color` | string | - | Avatar background colour |
| `circle` | boolean | false | Circular avatar shape |
| `initials` | string | - | Custom initials (no avatar) |
| `chevron` | boolean | true | Show dropdown indicator |
| `icon:trailing` | string | - | Custom trailing icon |
| `icon:variant` | string | micro | `micro`, `outline` |

## Slots

| Slot | Description |
|------|-------------|
| `avatar` | Custom avatar content |

## Basic Usage

```blade
<flux:profile avatar="https://example.com/avatar.jpg" />
```

## With Name

```blade
<flux:profile
    name="Jane Smith"
    avatar="https://example.com/avatar.jpg"
/>
```

## Without Chevron

```blade
<flux:profile
    :chevron="false"
    avatar="https://example.com/avatar.jpg"
/>
```

## Circular Avatar

```blade
<flux:profile
    circle
    :chevron="false"
    avatar="https://example.com/avatar.jpg"
/>
```

## Auto-generated Initials

When no avatar provided, generates initials from name:

```blade
<flux:profile name="Jane Smith" />
{{-- Shows "JS" --}}
```

## Custom Initials

```blade
<flux:profile initials="CP" />
```

## Custom Trailing Icon

```blade
<flux:profile
    icon:trailing="chevron-up-down"
    name="Jane Smith"
    avatar="https://example.com/avatar.jpg"
/>
```

## In Header with Dropdown

```blade
<flux:header>
    <flux:brand href="/" logo="/logo.svg" name="App" />
    <flux:spacer />
    <flux:dropdown>
        <flux:profile
            name="{{ auth()->user()->name }}"
            :avatar="auth()->user()->avatar_url"
        />
        <flux:menu>
            <flux:menu.item href="/profile" icon="user">Profile</flux:menu.item>
            <flux:menu.item href="/settings" icon="cog-6-tooth">Settings</flux:menu.item>
            <flux:menu.separator />
            <flux:menu.item href="/logout" icon="arrow-right-on-rectangle">Logout</flux:menu.item>
        </flux:menu>
    </flux:dropdown>
</flux:header>
```

## With Custom Avatar Slot

```blade
<flux:profile name="Jane Smith">
    <x-slot name="avatar">
        <flux:avatar circle size="sm" src="https://example.com/avatar.jpg">
            <flux:avatar.indicator color="lime" />
        </flux:avatar>
    </x-slot>
</flux:profile>
```

## Coloured Avatar

```blade
<flux:profile name="Jane Smith" avatar:color="blue" />
```
