# flux:autocomplete

Text input with suggestions that insert directly into the field as users type.

> **Note:** Does not support a `value` attribute. For scenarios requiring label display with stored values (e.g., user names with IDs), use `flux:select` with `searchable` instead.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Livewire property binding for input value |
| `type` | string | text | HTML input type (text, email, password, file, date) |
| `label` | string | - | Label text above input |
| `description` | string | - | Descriptive text below label |
| `placeholder` | string | - | Text shown when empty |
| `size` | string | - | Input sizing: `sm`, `xs` |
| `variant` | string | outline | Visual style: `outline`, `filled` |
| `disabled` | boolean | false | Disables user interaction |
| `readonly` | boolean | false | Makes input read-only |
| `invalid` | boolean | false | Applies error styling |
| `multiple` | boolean | false | Allows multiple file selection |
| `mask` | string | - | Alpine mask plugin pattern |
| `icon` | string | - | Leading icon name |
| `icon:trailing` | string | - | Trailing icon name |
| `kbd` | string | - | Keyboard shortcut hint |
| `clearable` | boolean | false | Shows clear button when filled |
| `copyable` | boolean | false | Displays copy button |
| `viewable` | boolean | false | Password toggle (password inputs) |
| `as` | string | input | Render element: `input`, `button` |
| `container:class` | string | - | CSS classes on container |
| `class:input` | string | - | CSS classes on input element |

## Child Component: flux:autocomplete.item

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `disabled` | boolean | false | Prevents item selection |

## Slots

| Slot | Description |
|------|-------------|
| `icon` / `icon:leading` | Custom leading content |
| `icon:trailing` | Custom trailing content |

---

## Basic Example

```blade
<flux:autocomplete wire:model="state" label="State of residence">
    <flux:autocomplete.item>Alabama</flux:autocomplete.item>
    <flux:autocomplete.item>Arkansas</flux:autocomplete.item>
    <flux:autocomplete.item>California</flux:autocomplete.item>
    <flux:autocomplete.item>Colorado</flux:autocomplete.item>
    <flux:autocomplete.item>Connecticut</flux:autocomplete.item>
</flux:autocomplete>
```

## With Placeholder

```blade
<flux:autocomplete wire:model="city" placeholder="Start typing a city...">
    <flux:autocomplete.item>London</flux:autocomplete.item>
    <flux:autocomplete.item>Manchester</flux:autocomplete.item>
    <flux:autocomplete.item>Birmingham</flux:autocomplete.item>
</flux:autocomplete>
```

## With Icon

```blade
<flux:autocomplete wire:model="search" icon="magnifying-glass" placeholder="Search...">
    <flux:autocomplete.item>Result 1</flux:autocomplete.item>
    <flux:autocomplete.item>Result 2</flux:autocomplete.item>
</flux:autocomplete>
```

## Disabled Items

```blade
<flux:autocomplete wire:model="option">
    <flux:autocomplete.item>Available option</flux:autocomplete.item>
    <flux:autocomplete.item disabled>Unavailable option</flux:autocomplete.item>
    <flux:autocomplete.item>Another option</flux:autocomplete.item>
</flux:autocomplete>
```

## Small Size

```blade
<flux:autocomplete wire:model="query" size="sm" placeholder="Quick search...">
    <flux:autocomplete.item>Option A</flux:autocomplete.item>
    <flux:autocomplete.item>Option B</flux:autocomplete.item>
</flux:autocomplete>
```

## Clearable

```blade
<flux:autocomplete wire:model="filter" clearable>
    <flux:autocomplete.item>Filter 1</flux:autocomplete.item>
    <flux:autocomplete.item>Filter 2</flux:autocomplete.item>
</flux:autocomplete>
```

---

## Related Components

- [Input](./input.md) - Standard text field
- [Select](./select.md) - Dropdown for single/multiple option selection with values
