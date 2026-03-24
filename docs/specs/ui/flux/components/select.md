# flux:select

Dropdown selection with native, listbox, and combobox variants.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Binds to Livewire property |
| `placeholder` | string | - | Text when no option selected |
| `label` | string | - | Label above select |
| `description` | string | - | Help text below select |
| `description:trailing` | boolean | false | Description below select |
| `badge` | string | - | Badge in label |
| `size` | string | - | `sm`, `xs` |
| `variant` | string | default | `default`, `listbox`, `combobox` |
| `multiple` | boolean | false | Multiple selections (listbox) |
| `filter` | boolean | true | Client-side filtering |
| `searchable` | boolean | false | Add search input |
| `empty` | string | No results found | Empty message |
| `clearable` | boolean | false | Show clear button |
| `selected-suffix` | string | selected | Text after count |
| `clear` | string | select | When to clear: `select`, `close` |
| `disabled` | boolean | false | Disable interaction |
| `invalid` | boolean | false | Error styling |

## Child Components

### flux:select.option

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `value` | mixed | - | Option value |
| `label` | string | - | Display text |
| `selected-label` | string | - | Text when selected |
| `disabled` | boolean | false | Prevent selection |

### flux:select.option.create (Pro)

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `min-length` | number | - | Chars before showing |
| `modal` | string | - | Modal to open |
| `wire:click` | string | - | Livewire action |

### flux:select.option.empty

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `when-loading` | string | - | Loading message |

### flux:select.button (listbox)

### flux:select.input (combobox)

### flux:select.search (listbox)

## Slots

| Slot | Description |
|------|-------------|
| `default` | Options |
| `trigger` | Custom trigger |
| `button` | Custom button (listbox) |
| `search` | Custom search (listbox) |
| `input` | Custom input (combobox) |
| `empty` | Custom empty state |

## Basic Usage (Native)

```blade
<flux:select wire:model="country" placeholder="Select country...">
    <flux:select.option value="uk">United Kingdom</flux:select.option>
    <flux:select.option value="us">United States</flux:select.option>
    <flux:select.option value="ca">Canada</flux:select.option>
</flux:select>
```

## With Label

```blade
<flux:select wire:model="role" label="Role" placeholder="Choose role...">
    <flux:select.option value="admin">Administrator</flux:select.option>
    <flux:select.option value="editor">Editor</flux:select.option>
    <flux:select.option value="viewer">Viewer</flux:select.option>
</flux:select>
```

## Listbox Variant (Pro)

Custom styled dropdown with icons and images:

```blade
<flux:select wire:model="status" variant="listbox" placeholder="Select status...">
    <flux:select.option value="active">
        <flux:icon.check-circle variant="mini" class="text-green-500" />
        Active
    </flux:select.option>
    <flux:select.option value="pending">
        <flux:icon.clock variant="mini" class="text-yellow-500" />
        Pending
    </flux:select.option>
    <flux:select.option value="inactive">
        <flux:icon.x-circle variant="mini" class="text-red-500" />
        Inactive
    </flux:select.option>
</flux:select>
```

## Searchable Listbox (Pro)

```blade
<flux:select wire:model="country" variant="listbox" searchable placeholder="Select country...">
    <flux:select.option value="uk">United Kingdom</flux:select.option>
    <flux:select.option value="us">United States</flux:select.option>
    {{-- Many more options --}}
</flux:select>
```

## Multiple Selection (Pro)

```blade
<flux:select wire:model="tags" variant="listbox" multiple placeholder="Select tags...">
    <flux:select.option value="php">PHP</flux:select.option>
    <flux:select.option value="laravel">Laravel</flux:select.option>
    <flux:select.option value="livewire">Livewire</flux:select.option>
</flux:select>
```

## Combobox Variant (Pro)

```blade
<flux:select wire:model="user" variant="combobox" placeholder="Search users...">
    @foreach ($users as $user)
        <flux:select.option :value="$user->id">{{ $user->name }}</flux:select.option>
    @endforeach
</flux:select>
```

## Backend Search (Pro)

```blade
<flux:select wire:model.live="userId" variant="combobox" :filter="false">
    <x-slot name="input">
        <flux:select.input wire:model.live="search" placeholder="Search..." />
    </x-slot>

    @foreach ($this->users as $user)
        <flux:select.option :value="$user->id">{{ $user->name }}</flux:select.option>
    @endforeach

    <flux:select.option.empty>No users found</flux:select.option.empty>
</flux:select>
```

## Create Option (Pro)

```blade
<flux:select wire:model="tagId" variant="listbox" searchable>
    @foreach ($tags as $tag)
        <flux:select.option :value="$tag->id">{{ $tag->name }}</flux:select.option>
    @endforeach

    <flux:select.option.create wire:click="createTag" min-length="2">
        Create "<span wire:text="search"></span>"
    </flux:select.option.create>
</flux:select>
```

## Clearable

```blade
<flux:select wire:model="filter" variant="listbox" clearable placeholder="Filter by...">
    {{-- Options --}}
</flux:select>
```
