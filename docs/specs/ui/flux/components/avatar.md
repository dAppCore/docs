# flux:avatar

Display an image, initials, or icon as a user avatar.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `name` | string | - | User's name; auto-generates initials if no `initials` provided |
| `src` | string | - | Image URL for avatar display |
| `initials` | string | - | Custom initials; overrides `name` if provided |
| `initials:single` | boolean | false | Use single initial from name |
| `alt` | string | name | Alternative text for image |
| `size` | string | 40px | `xs` (24px), `sm` (32px), default (40px), `lg` (48px), `xl` (64px) |
| `color` | string | - | Background colour for initials/icon avatars |
| `color:seed` | mixed | - | Deterministic colour generation with `color="auto"` |
| `circle` | boolean | false | Makes avatar fully circular |
| `icon` | string | - | Icon name instead of image/initials |
| `icon:variant` | string | solid | `outline` or `solid` |
| `as` | string | div | Render as: `button`, `div` |
| `href` | string | - | Makes avatar a link |

### Colour Options

`zinc`, `red`, `orange`, `amber`, `yellow`, `lime`, `green`, `emerald`, `teal`, `cyan`, `sky`, `blue`, `indigo`, `violet`, `purple`, `fuchsia`, `pink`, `rose`, `auto`

## Badge Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `badge` | string\|boolean\|slot | - | Badge content |
| `badge:color` | string | - | Same colour options as `color` prop |
| `badge:circle` | boolean | false | Fully circular badge |
| `badge:position` | string | bottom right | `top left`, `top right`, `bottom left`, `bottom right` |
| `badge:variant` | string | solid | `solid` or `outline` |

## Tooltip Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `tooltip` | string\|boolean | - | Hover tooltip text; `true` uses `name` |
| `tooltip:position` | string | top | `top`, `right`, `bottom`, `left` |

## Slots

| Slot | Description |
|------|-------------|
| `default` | Custom content overriding initials |
| `badge` | Complex badge content |

---

## Basic Image

```blade
<flux:avatar src="/path/to/avatar.jpg" />
```

## With Name (Auto Initials)

```blade
<flux:avatar name="Caleb Porzio" />
{{-- Generates "CP" --}}
```

## Single Initial

```blade
<flux:avatar name="calebporzio" initials:single />
{{-- Generates "C" --}}
```

## Custom Initials

```blade
<flux:avatar initials="JD" />
```

## Icon Avatar

```blade
<flux:avatar icon="user" />
```

## Sizes

```blade
<flux:avatar src="/avatar.jpg" size="xs" />  {{-- 24px --}}
<flux:avatar src="/avatar.jpg" size="sm" />  {{-- 32px --}}
<flux:avatar src="/avatar.jpg" />            {{-- 40px default --}}
<flux:avatar src="/avatar.jpg" size="lg" />  {{-- 48px --}}
<flux:avatar src="/avatar.jpg" size="xl" />  {{-- 64px --}}
```

## Colours

```blade
<flux:avatar name="John Doe" color="blue" />
<flux:avatar name="Jane Smith" color="green" />
<flux:avatar name="Bob Wilson" color="auto" />
{{-- auto assigns consistent colour based on name --}}
```

## Deterministic Colour

```blade
<flux:avatar name="User" color="auto" color:seed="{{ $user->id }}" />
{{-- Same user ID always gets same colour --}}
```

## Circular

```blade
<flux:avatar src="/avatar.jpg" circle />
```

## With Tooltip

```blade
<flux:avatar tooltip="John Doe" src="/avatar.jpg" />

{{-- Or use name as tooltip --}}
<flux:avatar tooltip name="John Doe" src="/avatar.jpg" />

{{-- Positioned tooltip --}}
<flux:avatar tooltip="John Doe" tooltip:position="right" src="/avatar.jpg" />
```

## With Badge

```blade
{{-- Dot indicator --}}
<flux:avatar src="/avatar.jpg" badge />

{{-- Coloured dot --}}
<flux:avatar src="/avatar.jpg" badge badge:color="green" />

{{-- Numeric badge --}}
<flux:avatar src="/avatar.jpg" badge="5" />

{{-- Positioned badge --}}
<flux:avatar src="/avatar.jpg" badge badge:position="top right" />

{{-- Emoji badge --}}
<flux:avatar src="/avatar.jpg">
    <x-slot:badge>
        <span class="text-xs">🎉</span>
    </x-slot:badge>
</flux:avatar>
```

## As Link

```blade
<flux:avatar href="/profile" src="/avatar.jpg" />
```

## As Button

```blade
<flux:avatar as="button" src="/avatar.jpg" wire:click="openProfile" />
```

---

## Avatar Group

Group multiple avatars with overlap styling:

```blade
<flux:avatar.group>
    <flux:avatar src="/avatar1.jpg" />
    <flux:avatar src="/avatar2.jpg" />
    <flux:avatar src="/avatar3.jpg" />
    <flux:avatar initials="+5" />
</flux:avatar.group>
```

Customise ring colour:

```blade
<flux:avatar.group class="*:ring-zinc-100">
    <flux:avatar src="/avatar1.jpg" />
    <flux:avatar src="/avatar2.jpg" />
</flux:avatar.group>
```
