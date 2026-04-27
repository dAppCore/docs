# flux:time-picker (Pro)

Time selection for scheduling with multiple formats and intervals.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Binds to Livewire property |
| `value` | string | - | Selected time (H:i format) |
| `type` | string | button | `button`, `input` |
| `multiple` | boolean | false | Allow multiple selections |
| `time-format` | string | auto | `auto`, `12-hour`, `24-hour` |
| `interval` | number | 30 | Minutes between options |
| `min` | string | - | Earliest time or "now" |
| `max` | string | - | Latest time or "now" |
| `unavailable` | string | - | Disabled times (comma-separated) |
| `open-to` | string | - | Default open time |
| `label` | string | - | Label text |
| `description` | string | - | Help text |
| `description:trailing` | boolean | - | Description below |
| `badge` | string | - | Badge in label |
| `placeholder` | string | - | Empty state text |
| `size` | string | - | `sm`, `xs` |
| `clearable` | boolean | false | Show clear button |
| `disabled` | boolean | false | Disable interaction |
| `invalid` | boolean | false | Error styling |
| `locale` | string | - | e.g., `fr`, `en-US`, `ja-JP` |
| `dropdown` | boolean | true | Show dropdown (for input type) |

## Basic Usage

```blade
<flux:time-picker wire:model="time" />
```

## With Default Value

```blade
<flux:time-picker wire:model="time" value="11:30" />
```

## Input Trigger

```blade
<flux:time-picker wire:model="time" type="input" />
```

## Without Dropdown (Manual Input)

```blade
<flux:time-picker wire:model="time" type="input" :dropdown="false" />
```

## Multiple Selection

```blade
<flux:time-picker wire:model="availableTimes" multiple />
```

## 12-hour Format

```blade
<flux:time-picker wire:model="time" time-format="12-hour" />
```

## 24-hour Format

```blade
<flux:time-picker wire:model="time" time-format="24-hour" />
```

## Custom Interval

```blade
{{-- Every 15 minutes --}}
<flux:time-picker wire:model="time" :interval="15" />

{{-- Every hour --}}
<flux:time-picker wire:model="time" :interval="60" />
```

## Min/Max Restrictions

```blade
{{-- Business hours only --}}
<flux:time-picker wire:model="time" min="09:00" max="17:00" />

{{-- From now onwards --}}
<flux:time-picker wire:model="time" min="now" />
```

## Unavailable Times

```blade
{{-- Block specific times --}}
<flux:time-picker wire:model="time" unavailable="12:00,12:30,13:00" />

{{-- Block time ranges --}}
<flux:time-picker wire:model="time" unavailable="03:00,04:00,05:30-07:29" />
```

## With Label

```blade
<flux:time-picker
    wire:model="meetingTime"
    label="Meeting time"
    placeholder="Select time..."
/>
```

## Clearable

```blade
<flux:time-picker wire:model="time" clearable />
```

## Size Variants

```blade
<flux:time-picker wire:model="time" size="sm" />
<flux:time-picker wire:model="time" size="xs" />
```

## Locale

```blade
<flux:time-picker wire:model="time" locale="fr" />
<flux:time-picker wire:model="time" locale="ja-JP" />
```

## Default Open Position

```blade
<flux:time-picker wire:model="time" open-to="14:00" />
```
