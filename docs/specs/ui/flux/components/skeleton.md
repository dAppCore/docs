# flux:skeleton

Loading placeholder with animated shimmer or pulse effects.

## Props

### flux:skeleton

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `animate` | string | - | `shimmer`, `pulse` |

### flux:skeleton.line

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `size` | string | base | `base`, `lg` |
| `animate` | string | - | `shimmer`, `pulse` (inherits from parent) |

### flux:skeleton.group

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `animate` | string | - | Applied to all children |

## CSS Variables

- `--flux-shimmer-color` - Shimmer background (defaults to white/zinc-900)

## Basic Usage

```blade
<flux:skeleton class="h-4 w-32 rounded" />
```

## With Animation

```blade
<flux:skeleton animate="shimmer" class="h-4 w-32 rounded" />
<flux:skeleton animate="pulse" class="h-4 w-32 rounded" />
```

## Text Lines

```blade
<flux:skeleton.group animate="shimmer">
    <flux:skeleton.line class="mb-2 w-1/4" />
    <flux:skeleton.line />
    <flux:skeleton.line />
    <flux:skeleton.line class="w-3/4" />
</flux:skeleton.group>
```

## Profile Placeholder

```blade
<flux:skeleton.group animate="shimmer" class="flex items-center gap-4">
    <flux:skeleton class="size-10 rounded-full" />
    <div class="flex-1">
        <flux:skeleton.line class="mb-1" />
        <flux:skeleton.line class="w-1/2" />
    </div>
</flux:skeleton.group>
```

## Card Placeholder

```blade
<flux:card>
    <flux:skeleton.group animate="shimmer">
        <flux:skeleton class="h-48 w-full rounded mb-4" />
        <flux:skeleton.line size="lg" class="w-3/4 mb-2" />
        <flux:skeleton.line class="w-full" />
        <flux:skeleton.line class="w-2/3" />
    </flux:skeleton.group>
</flux:card>
```

## Table Placeholder

```blade
<flux:table>
    <flux:table.columns>
        <flux:table.column>Name</flux:table.column>
        <flux:table.column>Email</flux:table.column>
        <flux:table.column>Role</flux:table.column>
    </flux:table.columns>

    <flux:table.rows>
        @for ($i = 0; $i < 5; $i++)
            <flux:table.row>
                <flux:table.cell>
                    <flux:skeleton.line animate="shimmer" class="w-24" />
                </flux:table.cell>
                <flux:table.cell>
                    <flux:skeleton.line animate="shimmer" class="w-32" />
                </flux:table.cell>
                <flux:table.cell>
                    <flux:skeleton.line animate="shimmer" class="w-16" />
                </flux:table.cell>
            </flux:table.row>
        @endfor
    </flux:table.rows>
</flux:table>
```

## Chart Placeholder

```blade
<flux:skeleton.group animate="shimmer" class="space-y-2">
    <flux:skeleton class="h-64 w-full rounded" />
    <div class="flex justify-between">
        <flux:skeleton class="h-4 w-12 rounded" />
        <flux:skeleton class="h-4 w-12 rounded" />
        <flux:skeleton class="h-4 w-12 rounded" />
    </div>
</flux:skeleton.group>
```
