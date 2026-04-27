# flux:chart

Lightweight, zero-dependency charting for Livewire. Supports line and area charts.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Binds to Livewire property containing data |
| `value` | array | - | Direct data array |
| `curve` | string | smooth | `smooth`, `none` |

## Child Components

- `flux:chart.svg` - SVG container (`gutter` prop for padding)
- `flux:chart.line` - Line (`field`, `curve`, `class`)
- `flux:chart.area` - Filled area (`field`, `curve`, `class`)
- `flux:chart.point` - Data points (`field`, `r`, `stroke-width`)
- `flux:chart.axis` - Axes (`axis`, `field`, `scale`, `position`, `tick-*`, `format`)
- `flux:chart.axis.line` - Axis baseline
- `flux:chart.axis.mark` - Tick marks
- `flux:chart.axis.grid` - Grid lines
- `flux:chart.axis.tick` - Tick labels
- `flux:chart.cursor` - Hover guide line
- `flux:chart.zero-line` - Y=0 line
- `flux:chart.tooltip` - Hover tooltip
- `flux:chart.tooltip.heading` - Tooltip header
- `flux:chart.tooltip.value` - Tooltip value
- `flux:chart.summary` - Key metric display
- `flux:chart.summary.value` - Summary value
- `flux:chart.legend` - Legend container
- `flux:chart.legend.indicator` - Colour indicator
- `flux:chart.viewport` - Wrapper for siblings

## Basic Line Chart

```blade
<flux:chart wire:model="data" class="aspect-3/1">
    <flux:chart.svg>
        <flux:chart.line field="visitors" class="text-pink-500" />
        <flux:chart.axis axis="x" field="date">
            <flux:chart.axis.line />
            <flux:chart.axis.tick />
        </flux:chart.axis>
        <flux:chart.axis axis="y">
            <flux:chart.axis.grid />
            <flux:chart.axis.tick />
        </flux:chart.axis>
        <flux:chart.cursor />
    </flux:chart.svg>
    <flux:chart.tooltip>
        <flux:chart.tooltip.heading field="date" />
        <flux:chart.tooltip.value field="visitors" label="Visitors" />
    </flux:chart.tooltip>
</flux:chart>
```

## Data Format

```php
public array $data = [
    ['date' => '2026-01-01', 'visitors' => 267],
    ['date' => '2026-01-02', 'visitors' => 259],
];
```

## Format Options (Intl API)

```blade
{{-- Currency --}}
:format="['style' => 'currency', 'currency' => 'USD']"

{{-- Percent --}}
:format="['style' => 'percent']"

{{-- Compact numbers --}}
:format="['notation' => 'compact']"
```
