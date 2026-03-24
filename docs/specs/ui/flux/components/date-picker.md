# flux:date-picker

Date selection via calendar overlay. Supports single dates, ranges, and presets.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Binds to Livewire property |
| `value` | string | - | Selected date(s): Y-m-d or Y-m-d/Y-m-d for ranges |
| `mode` | string | single | `single`, `range` |
| `min-range` | number | - | Minimum selectable days in range |
| `max-range` | number | - | Maximum selectable days in range |
| `min` | string | - | Earliest selectable date or "today" |
| `max` | string | - | Latest selectable date or "today" |
| `open-to` | string | - | Initial calendar view date |
| `force-open-to` | boolean | false | Always open to specified date |
| `months` | number | 1/2 | Months displayed (1 single, 2 range) |
| `label` | string | - | Wraps in flux:field with flux:label |
| `description` | string | - | Help text above picker |
| `description:trailing` | string | - | Help text below picker |
| `badge` | string | - | Label badge text |
| `placeholder` | string | - | Default text when empty |
| `size` | string | default | `sm`, `default`, `lg`, `xl`, `2xl` |
| `start-day` | number | locale | Week start day (0-6) |
| `week-numbers` | boolean | false | Display week numbers |
| `selectable-header` | boolean | false | Month/year dropdown navigation |
| `with-today` | boolean | false | Quick "today" navigation button |
| `with-inputs` | boolean | false | Manual date entry inputs |
| `with-confirmation` | boolean | false | Require confirmation |
| `with-presets` | boolean | false | Show preset ranges |
| `presets` | string | - | Space-separated preset names |
| `clearable` | boolean | false | Show clear button |
| `disabled` | boolean | false | Disable interaction |
| `invalid` | boolean | false | Error styling |
| `locale` | string | browser | Locale code (e.g., fr, en-US) |
| `unavailable` | string | - | Comma-separated disabled dates |

## Child Components

### flux:date-picker.input

Trigger input for date selection.

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `label` | string | - | Input label |
| `description` | string | - | Help text |
| `placeholder` | string | - | Placeholder text |
| `clearable` | boolean | false | Show clear button |
| `disabled` | boolean | false | Disable input |
| `invalid` | boolean | false | Error styling |

### flux:date-picker.button

Trigger button for date selection.

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `placeholder` | string | - | Button text |
| `size` | string | - | `sm`, `xs` |
| `clearable` | boolean | false | Show clear button |
| `disabled` | boolean | false | Disable button |
| `invalid` | boolean | false | Error styling |

## Slots

| Slot | Description |
|------|-------------|
| `trigger` | Custom element opening picker |

## Available Presets

| Key | Constructor | Range |
|-----|-------------|-------|
| `today` | `DateRange::today()` | Current day |
| `yesterday` | `DateRange::yesterday()` | Previous day |
| `thisWeek` | `DateRange::thisWeek()` | Current week |
| `lastWeek` | `DateRange::lastWeek()` | Previous week |
| `last7Days` | `DateRange::last7Days()` | Previous 7 days |
| `last14Days` | `DateRange::last14Days()` | Previous 14 days |
| `last30Days` | `DateRange::last30Days()` | Previous 30 days |
| `thisMonth` | `DateRange::thisMonth()` | Current month |
| `lastMonth` | `DateRange::lastMonth()` | Previous month |
| `last3Months` | `DateRange::last3Months()` | Previous 3 months |
| `last6Months` | `DateRange::last6Months()` | Previous 6 months |
| `thisQuarter` | `DateRange::thisQuarter()` | Current quarter |
| `lastQuarter` | `DateRange::lastQuarter()` | Previous quarter |
| `thisYear` | `DateRange::thisYear()` | Current year |
| `lastYear` | `DateRange::lastYear()` | Previous year |
| `yearToDate` | `DateRange::yearToDate()` | Jan 1 to today |
| `allTime` | `DateRange::allTime($start)` | Min date to today |
| `custom` | `DateRange::custom()` | User-defined |

## DateRange Object

```php
// Instance methods
$range->start()      // Carbon instance
$range->end()        // Carbon instance
$range->days()       // Integer count
$range->preset()     // DateRangePreset enum
$range->toArray()    // Array [start, end]
$range->contains($date)  // Boolean
$range->length()     // Number of days
$range->isNotAllTime()   // Boolean

// Custom instantiation
new DateRange(now(), now()->addDays(7))
```

## Basic Usage

```blade
<flux:date-picker wire:model="date" />
```

## Range with Presets

```blade
<flux:date-picker mode="range" with-presets />
```

## Input Trigger

```blade
<flux:date-picker wire:model="date">
    <x-slot name="trigger">
        <flux:date-picker.input />
    </x-slot>
</flux:date-picker>
```

## Range with Dual Inputs

```blade
<flux:date-picker mode="range">
    <x-slot name="trigger">
        <div class="flex gap-4">
            <flux:date-picker.input label="Start" />
            <flux:date-picker.input label="End" />
        </div>
    </x-slot>
</flux:date-picker>
```

## Livewire Property

```php
public DateRange $range;

public function mount()
{
    $this->range = new DateRange(
        now()->subDays(1),
        now()->addDays(1)
    );
}
```

## Session Persistence

```php
#[Session]
public DateRange $range;
```

## Eloquent Integration

```php
$query->whereBetween('created_at', $this->range->toArray());
```
