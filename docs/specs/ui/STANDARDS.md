# Host Hub UI Standards

> Standardised patterns for Flux Pro components. Reference this when building or reviewing UI.

## Toast Notifications

Use `Flux::toast()` in Livewire components instead of session flash messages.

```php
use Flux\Flux;

// Success
Flux::toast(text: 'Changes saved.', variant: 'success');

// Error
Flux::toast(text: 'Failed to save changes.', variant: 'danger');

// Warning
Flux::toast(text: 'Your session will expire soon.', variant: 'warning');

// Info (default)
Flux::toast('Processing your request...');
```

**Do NOT use:**
- `session()->flash('message', '...')`
- `session()->flash('success', '...')`
- `session()->flash('error', '...')`

---

## Badge Colour Semantics

Consistent colour meanings across the application:

| Colour | Meaning | Examples |
|--------|---------|----------|
| `green` | Active, success, healthy, completed | Active subscription, Published, Online |
| `amber` / `yellow` | Warning, pending, needs attention | Pending approval, Expiring soon |
| `red` | Error, danger, inactive, failed | Failed, Cancelled, Offline |
| `blue` | Info, in progress | Processing, In review |
| `violet` | Branded, Host UK accent | Premium features, Pro |
| `lime` | New, fresh | New feature, Just added |
| `zinc` / `gray` | Neutral, secondary, default | Draft, Unknown, N/A |

### Status Badge Patterns

```blade
{{-- Active/Inactive --}}
<flux:badge color="green">Active</flux:badge>
<flux:badge color="red">Inactive</flux:badge>

{{-- Progress states --}}
<flux:badge color="zinc">Draft</flux:badge>
<flux:badge color="yellow">Pending</flux:badge>
<flux:badge color="blue">In Progress</flux:badge>
<flux:badge color="green">Completed</flux:badge>
<flux:badge color="red">Failed</flux:badge>

{{-- With icons --}}
<flux:badge color="green" icon="check-circle">Published</flux:badge>
<flux:badge color="yellow" icon="clock">Scheduled</flux:badge>
<flux:badge color="red" icon="x-circle">Rejected</flux:badge>

{{-- Pill variant for counts/tags --}}
<flux:badge variant="pill" color="blue">12</flux:badge>
<flux:badge variant="pill" color="violet">Pro</flux:badge>
```

### Helper for Dynamic Status

```php
// In Livewire component
public function statusColor(string $status): string
{
    return match($status) {
        'active', 'published', 'completed', 'success' => 'green',
        'pending', 'scheduled', 'warning' => 'yellow',
        'processing', 'in_progress', 'info' => 'blue',
        'failed', 'cancelled', 'error', 'inactive' => 'red',
        'draft', 'unknown' => 'zinc',
        default => 'zinc',
    };
}
```

---

## Button Variants

Consistent button usage based on action importance:

| Variant | Usage | Example |
|---------|-------|---------|
| `primary` | Main CTA, primary action | Save, Create, Submit |
| (default/outline) | Secondary action | Cancel, Back, View |
| `ghost` | Tertiary action, low emphasis | Close, Skip, Learn more |
| `danger` | Destructive action | Delete, Remove, Disconnect |
| `subtle` | Very low emphasis | Dismiss, Hide |
| `filled` | Alternative emphasis | Special actions |

### Button Patterns

```blade
{{-- Primary action --}}
<flux:button variant="primary">Save Changes</flux:button>
<flux:button variant="primary" icon="plus">Create New</flux:button>

{{-- Secondary action --}}
<flux:button>Cancel</flux:button>
<flux:button icon="arrow-left">Back</flux:button>

{{-- Tertiary/ghost --}}
<flux:button variant="ghost">Skip</flux:button>
<flux:button variant="ghost" icon="x-mark">Close</flux:button>

{{-- Danger --}}
<flux:button variant="danger">Delete</flux:button>
<flux:button variant="danger" icon="trash">Remove</flux:button>

{{-- Icon-only with tooltip --}}
<flux:button icon="pencil" tooltip="Edit" />
<flux:button icon="trash" variant="danger" tooltip="Delete" />
<flux:button icon="cog-6-tooth" variant="ghost" tooltip="Settings" />

{{-- Save/Cancel pair --}}
<div class="flex gap-2">
    <flux:button variant="ghost">Cancel</flux:button>
    <flux:button variant="primary">Save</flux:button>
</div>

{{-- With loading (automatic on wire:click) --}}
<flux:button variant="primary" wire:click="save">Save</flux:button>
```

### Migration from Manual Buttons

Replace manual Tailwind buttons:

```blade
{{-- OLD: Manual button --}}
<a href="..." class="btn bg-violet-500 hover:bg-violet-600 text-white">
    Create
</a>

{{-- NEW: Flux button --}}
<flux:button href="..." variant="primary">Create</flux:button>
```

```blade
{{-- OLD: Manual danger button --}}
<button class="btn bg-red-500 hover:bg-red-600 text-white">
    Delete
</button>

{{-- NEW: Flux button --}}
<flux:button variant="danger">Delete</flux:button>
```

---

## Button Groups

Use for related actions:

```blade
{{-- Toggle buttons --}}
<flux:button.group>
    <flux:button icon="list-bullet">List</flux:button>
    <flux:button icon="squares-2x2">Grid</flux:button>
</flux:button.group>

{{-- Segmented control --}}
<flux:button.group>
    <flux:button>Day</flux:button>
    <flux:button>Week</flux:button>
    <flux:button>Month</flux:button>
</flux:button.group>

{{-- Action with dropdown --}}
<flux:button.group>
    <flux:button variant="primary">Save</flux:button>
    <flux:dropdown>
        <flux:button variant="primary" icon="chevron-down" />
        <flux:menu>
            <flux:menu.item>Save as draft</flux:menu.item>
            <flux:menu.item>Save and publish</flux:menu.item>
        </flux:menu>
    </flux:dropdown>
</flux:button.group>
```

---

## Icon Guidelines

We use Font Awesome Pro. Prefer:

- **Solid icons** for primary actions and filled states
- **Regular icons** for secondary elements
- **Light icons** for subtle/decorative use
- **Brand icons** for social platforms (fa-twitter, fa-instagram, etc.)

With Flux (Heroicons):
```blade
<flux:button icon="plus">Add</flux:button>
<flux:button icon="pencil">Edit</flux:button>
<flux:button icon="trash">Delete</flux:button>
```

With Font Awesome (in custom components):
```blade
<i class="fa-solid fa-plus"></i>
<i class="fa-brands fa-twitter"></i>
```

---

## Quick Reference

### Status Colours
- ✅ Green = Active/Success
- ⚠️ Yellow/Amber = Warning/Pending
- ❌ Red = Error/Danger
- ℹ️ Blue = Info/Processing
- 💜 Violet = Branded/Premium
- ⚪ Zinc = Neutral/Default

### Button Hierarchy
1. `variant="primary"` - Main action
2. (default) - Secondary
3. `variant="ghost"` - Tertiary
4. `variant="danger"` - Destructive
