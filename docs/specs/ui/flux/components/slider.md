# flux:slider

Range slider for selecting numeric values.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Binds to Livewire property |
| `value` | string/array | - | Initial value(s) |
| `range` | boolean | false | Enable range selection |
| `min` | number | - | Minimum value |
| `max` | number | - | Maximum value |
| `step` | number | - | Step increment |
| `big-step` | number | - | Step when holding Shift |
| `min-steps-between` | number | - | Min distance between thumbs |
| `track:class` | string | - | CSS for track |
| `thumb:class` | string | - | CSS for thumb |

## Child Components

### flux:slider.tick

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `value` | number | - | Position for tick |

**Slot:** Label text (shows line if empty)

## Basic Usage

```blade
<flux:slider wire:model="amount" />
```

## With Min/Max

```blade
<flux:slider wire:model="price" min="0" max="1000" />
```

## With Step

```blade
<flux:slider wire:model="quantity" min="0" max="100" step="5" />
```

## Big Step (Shift Key)

```blade
<flux:slider wire:model="volume" min="0" max="100" step="1" big-step="10" />
```

## Display Value

```blade
<div class="flex items-center gap-4">
    <flux:slider wire:model="amount" class="flex-1" />
    <span wire:text="amount" class="tabular-nums w-12 text-right"></span>
</div>
```

## Range Selection

```blade
<flux:slider wire:model="priceRange" range min="0" max="1000" />
```

```php
public array $priceRange = [100, 500];
```

## With Ticks

```blade
<flux:slider wire:model="rating" min="1" max="5">
    <flux:slider.tick value="1">Poor</flux:slider.tick>
    <flux:slider.tick value="2" />
    <flux:slider.tick value="3">Average</flux:slider.tick>
    <flux:slider.tick value="4" />
    <flux:slider.tick value="5">Excellent</flux:slider.tick>
</flux:slider>
```

## Custom Styling

```blade
<flux:slider
    wire:model="value"
    track:class="h-2"
    thumb:class="size-6"
/>
```

## With Field

```blade
<flux:field>
    <flux:label>Volume</flux:label>
    <flux:slider wire:model="volume" min="0" max="100" />
    <flux:description>Adjust playback volume</flux:description>
</flux:field>
```

## Price Range Filter

```blade
<div class="space-y-2">
    <div class="flex justify-between text-sm">
        <span>£<span wire:text="priceRange.0">0</span></span>
        <span>£<span wire:text="priceRange.1">1000</span></span>
    </div>
    <flux:slider
        wire:model.live="priceRange"
        range
        min="0"
        max="1000"
        step="50"
        min-steps-between="2"
    />
</div>
```
