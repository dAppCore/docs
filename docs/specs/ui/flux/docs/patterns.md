# Flux Design Patterns

Understanding patterns allows guessing component usage without prior experience.

## Props vs Attributes

| Type | Purpose | Example |
|------|---------|---------|
| Props | Flux-provided, customise internally | `variant="primary"` |
| Attributes | HTML attributes, forwarded to DOM | `x-on:change.prevent` |

```blade
<flux:button variant="primary" x-on:click="save">
    Save
</flux:button>
```

## Class Merging

Custom classes merge with Flux classes:

```blade
<flux:button class="w-full">
{{-- Output: class="w-full border border-zinc-200 ..." --}}
```

### Resolving Conflicts

Use `!` modifier when classes conflict:

```blade
<flux:button class="bg-zinc-800! hover:bg-zinc-700!">
    Custom
</flux:button>
```

**Alternatives:**
- Publish custom component variants
- Style via data attributes globally
- Create custom component

## Split Attribute Forwarding

Multi-element components distribute attributes intelligently:

```blade
<flux:input class="w-full" autofocus />
{{-- class → wrapper div --}}
{{-- autofocus → input element --}}
```

## Common Props

### Variant

```blade
<flux:button variant="primary" />
<flux:button variant="danger" />
<flux:badge variant="solid" />
```

### Icon

Uses Heroicons:

```blade
<flux:button icon="magnifying-glass" />
<flux:button icon:trailing="chevron-down" />
```

### Size

```blade
<flux:button size="sm" />
<flux:button size="lg" />
<flux:heading size="xl" />
```

### Keyboard Hints

```blade
<flux:button kbd="⌘S">Save</flux:button>
<flux:input kbd="⌘K" placeholder="Search..." />
```

### Inset

Negative margin for inline elements:

```blade
<flux:badge inset="top bottom" />
```

## Data Binding

Works like native Livewire:

```blade
<flux:input wire:model="email" />
<flux:checkbox wire:model="terms" />
<flux:tabs wire:model="activeTab" />
```

Also supports Alpine:

```blade
<flux:input x-model="email" />
<flux:input x-on:change="validate" />
```

## Component Groups

### Standalone with Grouping

`.group` suffix:

```blade
<flux:button.group>
    <flux:button>One</flux:button>
    <flux:button>Two</flux:button>
</flux:button.group>
```

### Child Components

`.item` suffix for children:

```blade
<flux:accordion>
    <flux:accordion.item heading="First">Content</flux:accordion.item>
    <flux:accordion.item heading="Second">Content</flux:accordion.item>
</flux:accordion>
```

## Root Components

Larger primitives use bare names:

```blade
<flux:field>
    <flux:label>Email</flux:label>
    <flux:input />
    <flux:error name="email" />
</flux:field>
```

Not: `flux:field.label` or `flux:field.error`

## Slots

Composition preferred, but slots when necessary:

```blade
<flux:input>
    <x-slot name="iconTrailing">
        <flux:button icon="x-mark" variant="subtle" size="sm" />
    </x-slot>
</flux:input>
```

## Shorthand Props

Verbose arrangements have shortcuts:

```blade
{{-- Long form --}}
<flux:field>
    <flux:label>Email</flux:label>
    <flux:input wire:model="email" />
    <flux:error name="email" />
</flux:field>

{{-- Shorthand --}}
<flux:input wire:model="email" label="Email" />
```

## Gotcha: Dynamic Attributes

Blade conditionals don't work in component tags:

```blade
{{-- Wrong --}}
<flux:input @if($disabled) disabled @endif />

{{-- Right --}}
<flux:input :disabled="$disabled" />
```
