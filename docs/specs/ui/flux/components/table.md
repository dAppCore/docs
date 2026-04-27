# flux:table

Structured data display with sorting, pagination, and sticky headers/columns.

## Props

### flux:table

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `paginate` | object | - | Laravel paginator instance |
| `container:class` | string | - | Container CSS (e.g., `max-h-80`) |

### flux:table.columns

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `sticky` | boolean | false | Sticky header row |

### flux:table.column

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `align` | string | - | `start`, `center`, `end` |
| `sortable` | boolean | false | Enable sorting |
| `sorted` | boolean | false | Currently sorted |
| `direction` | string | - | `asc`, `desc` |
| `sticky` | boolean | false | Sticky column |

### flux:table.row

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `key` | string | - | Unique row identifier |
| `sticky` | boolean | false | Sticky row |

### flux:table.cell

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `align` | string | - | `start`, `center`, `end` |
| `variant` | string | default | `default`, `strong` |
| `sticky` | boolean | false | Sticky cell |

## Basic Usage

```blade
<flux:table>
    <flux:table.columns>
        <flux:table.column>Name</flux:table.column>
        <flux:table.column>Email</flux:table.column>
        <flux:table.column>Role</flux:table.column>
    </flux:table.columns>

    <flux:table.rows>
        @foreach ($users as $user)
            <flux:table.row :key="$user->id">
                <flux:table.cell variant="strong">{{ $user->name }}</flux:table.cell>
                <flux:table.cell>{{ $user->email }}</flux:table.cell>
                <flux:table.cell>{{ $user->role }}</flux:table.cell>
            </flux:table.row>
        @endforeach
    </flux:table.rows>
</flux:table>
```

## With Pagination

```blade
<flux:table :paginate="$users">
    {{-- Columns and rows --}}
</flux:table>
```

## Sortable Columns

```blade
<flux:table>
    <flux:table.columns>
        <flux:table.column
            sortable
            :sorted="$sortBy === 'name'"
            :direction="$sortDirection"
            wire:click="sort('name')"
        >
            Name
        </flux:table.column>
        <flux:table.column
            sortable
            :sorted="$sortBy === 'created_at'"
            :direction="$sortDirection"
            wire:click="sort('created_at')"
        >
            Created
        </flux:table.column>
    </flux:table.columns>
    {{-- Rows --}}
</flux:table>
```

```php
public string $sortBy = 'name';
public string $sortDirection = 'asc';

public function sort($column)
{
    if ($this->sortBy === $column) {
        $this->sortDirection = $this->sortDirection === 'asc' ? 'desc' : 'asc';
    } else {
        $this->sortBy = $column;
        $this->sortDirection = 'asc';
    }
}
```

## Sticky Header

```blade
<flux:table container:class="max-h-80">
    <flux:table.columns sticky>
        <flux:table.column>Name</flux:table.column>
        <flux:table.column>Email</flux:table.column>
    </flux:table.columns>
    {{-- Rows --}}
</flux:table>
```

## Sticky Column

```blade
<flux:table>
    <flux:table.columns>
        <flux:table.column sticky>Name</flux:table.column>
        <flux:table.column>Email</flux:table.column>
        <flux:table.column>Phone</flux:table.column>
        {{-- More columns --}}
    </flux:table.columns>

    <flux:table.rows>
        @foreach ($users as $user)
            <flux:table.row>
                <flux:table.cell sticky variant="strong">{{ $user->name }}</flux:table.cell>
                <flux:table.cell>{{ $user->email }}</flux:table.cell>
                <flux:table.cell>{{ $user->phone }}</flux:table.cell>
            </flux:table.row>
        @endforeach
    </flux:table.rows>
</flux:table>
```

## With Actions

```blade
<flux:table.row>
    <flux:table.cell>{{ $user->name }}</flux:table.cell>
    <flux:table.cell>{{ $user->email }}</flux:table.cell>
    <flux:table.cell align="end">
        <flux:dropdown>
            <flux:button icon="ellipsis-horizontal" variant="ghost" size="sm" />
            <flux:menu>
                <flux:menu.item icon="pencil" wire:click="edit({{ $user->id }})">Edit</flux:menu.item>
                <flux:menu.item icon="trash" variant="danger" wire:click="delete({{ $user->id }})">Delete</flux:menu.item>
            </flux:menu>
        </flux:dropdown>
    </flux:table.cell>
</flux:table.row>
```

## With Badges and Avatars

```blade
<flux:table.row>
    <flux:table.cell>
        <div class="flex items-center gap-3">
            <flux:avatar size="sm" :src="$user->avatar_url" />
            {{ $user->name }}
        </div>
    </flux:table.cell>
    <flux:table.cell>
        <flux:badge :color="$user->is_active ? 'lime' : 'zinc'">
            {{ $user->is_active ? 'Active' : 'Inactive' }}
        </flux:badge>
    </flux:table.cell>
</flux:table.row>
```
