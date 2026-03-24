# flux:separator

Visual divider between content sections.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `vertical` | boolean | false | Vertical orientation |
| `variant` | string | standard | `standard`, `subtle` |
| `text` | string | - | Center text label |
| `orientation` | string | horizontal | `horizontal`, `vertical` |

## Basic Usage

```blade
<flux:separator />
```

## With Text

```blade
<flux:separator text="or" />
```

## Vertical

```blade
<div class="flex items-center gap-4">
    <flux:button>Option A</flux:button>
    <flux:separator vertical />
    <flux:button>Option B</flux:button>
</div>
```

## Vertical with Height

```blade
<flux:separator vertical class="my-2" />
```

## Subtle Variant

Low-contrast separator that blends with background:

```blade
<flux:separator variant="subtle" />
```

## In Form

```blade
<form>
    <flux:input label="Email" wire:model="email" />
    <flux:input label="Password" wire:model="password" />

    <flux:button type="submit" class="w-full">Sign in</flux:button>

    <flux:separator text="or continue with" class="my-4" />

    <div class="flex gap-2">
        <flux:button variant="ghost" class="flex-1">
            <flux:icon.google variant="mini" /> Google
        </flux:button>
        <flux:button variant="ghost" class="flex-1">
            <flux:icon.github variant="mini" /> GitHub
        </flux:button>
    </div>
</form>
```

## In Menu

```blade
<flux:menu>
    <flux:menu.item>Profile</flux:menu.item>
    <flux:menu.item>Settings</flux:menu.item>
    <flux:separator />
    <flux:menu.item variant="danger">Logout</flux:menu.item>
</flux:menu>
```
