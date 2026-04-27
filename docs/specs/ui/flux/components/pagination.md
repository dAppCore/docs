# flux:pagination

Navigate through paginated data with Previous/Next or page number buttons.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `paginator` | object | - | Laravel paginator instance |

## Basic Usage

With standard pagination (shows page numbers and total count):

```blade
{{-- In Livewire component --}}
public function render()
{
    return view('orders', [
        'orders' => Order::paginate(10),
    ]);
}
```

```blade
{{-- In Blade template --}}
<flux:table>
    {{-- Table content --}}
</flux:table>

<flux:pagination :paginator="$orders" />
```

## Simple Pagination

For large datasets where counting total records is expensive:

```php
$orders = Order::simplePaginate(10);
```

```blade
<flux:pagination :paginator="$orders" />
```

Shows only "Previous" and "Next" buttons without page numbers or total count.

## Features

- **Responsive design** - Automatically adjusts page links based on dataset size
- **Large result sets** - Shows first/last pages with ellipses for gaps
- **Status display** - Shows "Showing X to Y of Z results"
- **Simple mode** - Lightweight navigation for expensive counting operations

## With Table

```blade
<flux:table :paginate="$orders">
    <flux:table.columns>
        <flux:table.column>Order ID</flux:table.column>
        <flux:table.column>Customer</flux:table.column>
        <flux:table.column>Total</flux:table.column>
    </flux:table.columns>

    <flux:table.rows>
        @foreach ($orders as $order)
            <flux:table.row>
                <flux:table.cell>{{ $order->id }}</flux:table.cell>
                <flux:table.cell>{{ $order->customer->name }}</flux:table.cell>
                <flux:table.cell>{{ $order->formatted_total }}</flux:table.cell>
            </flux:table.row>
        @endforeach
    </flux:table.rows>
</flux:table>
```

## Livewire Integration

Pagination works automatically with Livewire using the `WithPagination` trait:

```php
use Livewire\WithPagination;

class OrderList extends Component
{
    use WithPagination;

    public function render()
    {
        return view('orders', [
            'orders' => Order::paginate(15),
        ]);
    }
}
```
