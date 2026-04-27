# Flux Design Principles

Core philosophy behind Flux component design.

## 1. Simplicity

Straightforward syntax over verbose nested structures. Simple names and props reduce cognitive overhead.

```blade
{{-- Simple --}}
<flux:input label="Email" wire:model="email" />

{{-- Not this --}}
<flux:form-control>
    <flux:form-control.label>Email</flux:form-control.label>
    <flux:form-control.input wire:model="email" />
</flux:form-control>
```

## 2. Composability

Balance simplicity with control through composable alternatives:

```blade
{{-- Simple for common case --}}
<flux:input label="Email" />

{{-- Composable for control --}}
<flux:field>
    <flux:label badge="Required">Email</flux:label>
    <flux:description>We'll never share your email.</flux:description>
    <flux:input />
    <flux:error name="email" />
</flux:field>
```

## 3. Friendliness

Accessible terminology over technical jargon:

| Flux | Not This |
|------|----------|
| input | form-control |
| accordion | disclosure |
| dropdown | popover-menu |
| heading | typography-h |

Language reflects how developers actually speak.

## 4. Composition

Foundational components mix and match to create complex interfaces:

```blade
{{-- Same button works everywhere --}}
<flux:button>Standalone</flux:button>

<flux:dropdown>
    <flux:button>In Dropdown</flux:button>
    <flux:menu>...</flux:menu>
</flux:dropdown>

<flux:modal.trigger name="confirm">
    <flux:button>Opens Modal</flux:button>
</flux:modal.trigger>
```

## 5. Consistency

Repeated patterns across components reduce cognitive load:

- `variant` for visual styles
- `size` for scaling
- `icon` for leading icons
- `icon:trailing` for trailing icons
- `wire:model` for data binding

## 6. Brevity

Single-word names preferred:

| Good | Avoid |
|------|-------|
| input | text-input |
| dropdown | dropdown-menu |
| button | action-button |

Avoid hyphenated compounds and excessive dot-nesting.

## 7. Browser-Native

Leverage modern browser capabilities:

- Popover API for dropdowns
- `<dialog>` element for modals
- Native form validation

Benefits: reliable behaviour, no extra JavaScript, better accessibility.

## CSS-First Solutions

Modern CSS solves problems historically requiring JavaScript:

```css
/* Search icon changes colour when input focused */
/* Zero JavaScript required */
[data-flux-command-input]:has(+ input:focus) .icon {
    color: var(--color-accent);
}
```

Using `:has()`, `:not()`, `:where()` selectors.

## "We Style, You Space"

**Flux provides:** padding, colours, borders, internal styling

**You manage:** margins, layout spacing, positioning

```blade
{{-- Flux handles button styling --}}
{{-- You handle the margin --}}
<flux:button class="mt-4">Save</flux:button>
```

Rationale: styling is universal, spacing is contextual. Pre-baked spacing creates constant overrides.

## Summary

1. **Simplicity** - Intuitive, minimal syntax
2. **Composability** - Simple and complex patterns with same components
3. **Friendliness** - Human terminology
4. **Composition** - Mix and match freely
5. **Consistency** - Predictable props across components
6. **Brevity** - Short, clear names
7. **Browser-Native** - Modern web APIs first
