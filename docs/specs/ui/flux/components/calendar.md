# flux:calendar

Flexible calendar component for date selection supporting single dates, multiple dates, and date ranges.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Binds calendar to Livewire property |
| `value` | string | - | Selected date(s): `Y-m-d`, comma-separated, or `Y-m-d/Y-m-d` |
| `mode` | string | single | `single`, `multiple`, `range` |
| `min` | string | - | Earliest selectable date or `today` |
| `max` | string | - | Latest selectable date or `today` |
| `size` | string | base | `xs`, `sm`, `base`, `lg`, `xl`, `2xl` |
| `start-day` | integer | locale-based | Week start (0-6, Sunday-Saturday) |
| `months` | integer | 1 | Number of months displayed (1 or 2) |
| `min-range` | integer | - | Minimum days for range selection |
| `max-range` | integer | - | Maximum days for range selection |
| `open-to` | string | - | Initial calendar month (`Y-m-d` format) |
| `force-open-to` | boolean | false | Always open to specified date |
| `navigation` | boolean | true | Show month navigation controls |
| `static` | boolean | false | Non-interactive display mode |
| `multiple` | boolean | false | Enable multiple date selection |
| `week-numbers` | boolean | false | Display week numbers |
| `selectable-header` | boolean | false | Enable month/year dropdowns |
| `with-today` | boolean | false | Show today navigation button |
| `with-inputs` | boolean | false | Show date input fields |
| `locale` | string | browser | Language/locale (e.g., `fr`, `en-US`, `ja-JP`) |

---

## Single Date Selection

```blade
<flux:calendar value="2026-01-06" />

{{-- With Livewire binding --}}
<flux:calendar wire:model="date" />
```

## Multiple Date Selection

```blade
<flux:calendar multiple value="2026-01-02,2026-01-05,2026-01-15" />

{{-- With Livewire binding --}}
<flux:calendar multiple wire:model="dates" />
```

## Range Selection

```blade
<flux:calendar mode="range" value="2026-01-02/2026-01-06" />

{{-- With Livewire binding --}}
<flux:calendar mode="range" wire:model="range" />
```

## Date Restrictions

```blade
{{-- Minimum date --}}
<flux:calendar min="2026-01-01" />

{{-- Maximum date --}}
<flux:calendar max="2026-12-31" />

{{-- Today as boundary --}}
<flux:calendar min="today" />

{{-- Range limits --}}
<flux:calendar mode="range" min-range="3" max-range="14" />
```

## Display Options

```blade
{{-- With week numbers --}}
<flux:calendar week-numbers />

{{-- Selectable month/year dropdowns --}}
<flux:calendar selectable-header />

{{-- Today button --}}
<flux:calendar with-today />

{{-- With input fields --}}
<flux:calendar with-inputs />

{{-- Multiple months --}}
<flux:calendar months="2" />
```

## Sizes

```blade
<flux:calendar size="xs" />
<flux:calendar size="sm" />
<flux:calendar />           {{-- base (default) --}}
<flux:calendar size="lg" />
<flux:calendar size="xl" />
<flux:calendar size="2xl" />
```

## Static Display

Non-interactive calendar (just displays dates):

```blade
<flux:calendar static />
```

## Locale

```blade
<flux:calendar locale="fr" />
<flux:calendar locale="ja-JP" />
<flux:calendar locale="en-GB" />
```

## Week Start Day

```blade
{{-- Start on Monday --}}
<flux:calendar start-day="1" />

{{-- Start on Sunday --}}
<flux:calendar start-day="0" />
```

---

## DateRange Object

For range mode, use the `DateRange` object in your Livewire component:

```php
use Flux\DateRange;

class MyComponent extends Component
{
    public DateRange $range;

    public function mount()
    {
        $this->range = new DateRange(now(), now()->addDays(7));
    }
}
```

### Static Constructors

```php
DateRange::today()
DateRange::yesterday()
DateRange::thisWeek()
DateRange::lastWeek()
DateRange::last7Days()
DateRange::thisMonth()
DateRange::lastMonth()
DateRange::thisYear()
DateRange::lastYear()
DateRange::yearToDate()
```

### Methods

```php
$range->start()         // Get start as Carbon instance
$range->end()           // Get end as Carbon instance
$range->days()          // Count days in range
$range->contains($date) // Check if date is in range
$range->length()        // Range length in days
$range->toArray()       // Array representation
$range->preset()        // Current preset enum
```

### Eloquent Integration

```php
Model::whereBetween('date', [$range->start(), $range->end()])->get();
```

## Session Persistence

Use `#[Session]` to persist selection:

```php
use Livewire\Attributes\Session;

#[Session]
public DateRange $range;
```

---

## Complete Example

```blade
<flux:calendar
    wire:model="dateRange"
    mode="range"
    months="2"
    min="today"
    max-range="30"
    selectable-header
    with-today
    week-numbers
/>
```
