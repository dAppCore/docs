# Flux UI - Core Principles & Patterns

## Seven Core Design Principles

Flux is built on seven intentional design decisions that guide every component:

### 1. Simplicity

The foundation of Flux's design philosophy. Components use straightforward, intuitive syntax.

**Example:**
```blade
<!-- Simple and direct -->
<flux:input wire:model="email" label="Email" />
```

Instead of nested, verbose structures with separate label components, the input handles labelling internally.

### 2. Complexity (Composability)

Acknowledges that simplicity has trade-offs. For complex use cases, components offer composable alternatives without forcing rigid hierarchies.

**Example:**
A button becomes a dropdown by wrapping it:
```blade
<!-- Simple button -->
<flux:button>Actions</flux:button>

<!-- Complex dropdown -->
<flux:dropdown>
    <flux:button>Actions</flux:button>
    <flux:menu>
        <flux:menu.item>Edit</flux:menu.item>
        <flux:menu.item>Delete</flux:menu.item>
    </flux:menu>
</flux:dropdown>
```

### 3. Friendliness

Uses familiar terminology over technical jargon. Developers can intuitively understand components.

| Instead of | Use |
|-----------|-----|
| Form controls | Form inputs |
| Disclosure | Accordion |
| Floating menu | Dropdown |

### 4. Composition

Treats learning as a mix-and-match process. Core components combine into more robust solutions. No need to learn 50 variations—learn the base components and compose them.

**Examples:**
- `flux:button` + `flux:dropdown` = action menu
- `flux:input` + `flux:field` = form field with label and description
- `flux:table` + `flux:badge` = data table with status indicators

### 5. Consistency

Employs repeated syntax patterns throughout the component system. Consistent naming and structures reduce guesswork and enhance intuitiveness.

**Patterns:**
- All inputs support `wire:model`, `label`, `description`, `invalid`
- All buttons support `variant`, `size`, `icon`, `icon:trailing`
- Navigation items all support `href`, `current`, `icon`, `badge`

### 6. Brevity

Minimises compound hyphenated words and nested dot-separated names. Uses simple, single words.

| Instead of | Use |
|-----------|-----|
| text-input | input |
| form-control | input (or field) |
| action-button | button |

### 7. Modern Browser Leverage

Utilises native browser features like the popover API and `<dialog>` element for reliable, accessible components without excessive custom JavaScript.

**Examples:**
- Dropdowns use `popover` API
- Modals use `<dialog>` element
- Tooltips use native HTML attributes

---

## Styling Philosophy: "We Style, You Space"

This principle defines the division of responsibility:

**Flux handles:**
- Component styling (colours, fonts, borders)
- Internal padding and spacing within components
- Dark mode variants
- Hover, focus, active states
- Accessibility features

**You handle:**
- Margins around components
- Layout spacing (gaps between items)
- Contextual spacing based on your design

**Example:**
```blade
<!-- Flux handles internal spacing of the button -->
<flux:button>Click me</flux:button>

<!-- You add margins to create layout spacing -->
<div class="flex gap-4">
    <flux:button>Save</flux:button>
    <flux:button variant="ghost">Cancel</flux:button>
</div>
```

This maintains flexibility because what's appropriate margin in one context differs in another.

---

## Common Prop Patterns

### Styling Props

**`variant`** - Visual style variations
```blade
<flux:button variant="default">Default</flux:button>
<flux:button variant="primary">Primary</flux:button>
<flux:button variant="filled">Filled</flux:button>
<flux:button variant="danger">Danger</flux:button>
<flux:button variant="ghost">Ghost</flux:button>
<flux:button variant="subtle">Subtle</flux:button>
```

**`size`** - Component sizing
```blade
<flux:button size="xs">Extra small</flux:button>
<flux:button>Default</flux:button>
<flux:button size="sm">Small</flux:button>
<flux:button size="lg">Large</flux:button>
```

**`color`** - Color palette (18 options)
```blade
<flux:badge color="red">Alert</flux:badge>
<flux:badge color="green">Success</flux:badge>
<flux:badge color="blue">Info</flux:badge>
```

**`class`** - Additional Tailwind utilities
```blade
<flux:button class="max-w-sm">Constrained width</flux:button>
```

### Icon Props

**`icon`** - Leading icon (Heroicons)
```blade
<flux:button icon="check">Confirm</flux:button>
<flux:input icon="magnifying-glass" placeholder="Search..." />
```

**`icon:trailing`** - Trailing icon
```blade
<flux:button icon:trailing="arrow-right">Next</flux:button>
```

**`icon:variant`** - Icon style
```blade
<flux:button icon="check" icon:variant="solid">Solid icon</flux:button>
```

Values: `outline` (default), `solid`, `mini`, `micro`

### State Props

**`disabled`** - Disable interaction
```blade
<flux:button disabled>Disabled</flux:button>
<flux:input disabled />
```

**`readonly`** - Lock content (inputs only)
```blade
<flux:input readonly />
```

**`invalid`** - Error state
```blade
<flux:input invalid />
```

**`current`** - Mark as active (navigation items)
```blade
<flux:navbar.item href="/" current>Home</flux:navbar.item>
```

### Layout Props

**`inset`** - Negative margin adjustment for inline spacing
```blade
<flux:badge inset="start">Badge</flux:badge>
```

**`sticky`** - Fixed positioning
```blade
<flux:sidebar sticky>
    <!-- Navigation stays fixed during scroll -->
</flux:sidebar>
```

**`collapsible`** - Collapse/expand capability
```blade
<flux:sidebar collapsible="mobile">
    <!-- Collapses on mobile, expands on desktop -->
</flux:sidebar>
```

---

## Component Structures

### Root Components

Primitive, simple components use bare names:
- `flux:label`
- `flux:text`
- `flux:heading`
- `flux:icon`
- `flux:button`
- `flux:input`

### Parent-Child Components

Complex structures use dot notation:
```blade
<flux:accordion>
    <flux:accordion.item>
        <flux:accordion.heading>Title</flux:accordion.heading>
        Content
    </flux:accordion.item>
</flux:accordion>
```

### Grouping Components

Related items can be grouped:
```blade
<flux:button.group>
    <flux:button>Option 1</flux:button>
    <flux:button>Option 2</flux:button>
</flux:button.group>
```

---

## Props vs Attributes

Flux distinguishes between two types of parameters:

### Props
Component-specific properties that control internal behaviour and styling.

```blade
<flux:button variant="primary" size="sm" icon="check">
```

These modify CSS classes and component structure.

### Attributes
Custom HTML attributes that forward directly to underlying DOM elements.

```blade
<flux:input x-on:change.prevent="handleChange" />
<flux:button onclick="handleClick()" />
```

### Split Attribute Forwarding

Complex components intelligently distribute attributes to appropriate elements:

```blade
<!-- class applies to wrapper, autofocus applies to actual input -->
<flux:input class="max-w-sm" autofocus />

<!-- Same pattern with form fields -->
<flux:field>
    <flux:label>Email</flux:label>
    <flux:input wire:model="email" autofocus />
</flux:field>
```

---

## Class Merging

Components automatically merge user-provided Tailwind classes with framework styles:

```blade
<flux:button class="w-full">Full width button</flux:button>
```

**When conflicts occur**, use the `!` modifier to force your styles:

```blade
<!-- Your bg-zinc-800 will override Flux defaults -->
<flux:button class="bg-zinc-800! hover:bg-zinc-700!">Custom</flux:button>
```

However, the documentation recommends alternatives like:
1. Using a published custom variant
2. Publishing the component locally to modify it
3. Using global CSS overrides with data attributes

---

## Data Binding with Wire:model

Livewire data binding integrates directly with form components:

```blade
<!-- Text inputs -->
<flux:input wire:model="email" />

<!-- Checkboxes -->
<flux:checkbox wire:model="agreeToTerms" />

<!-- Switches -->
<flux:switch wire:model="isActive" />

<!-- Textareas -->
<flux:textarea wire:model="bio" />

<!-- Grouped checkboxes -->
<flux:checkbox.group wire:model="tags">
    <flux:checkbox value="php">PHP</flux:checkbox>
    <flux:checkbox value="laravel">Laravel</flux:checkbox>
</flux:checkbox.group>
```

---

## Important Gotchas

### 1. No Conditionals in Component Opening Tags

**This doesn't work:**
```blade
<!-- WRONG -->
<flux:button @if($show) variant="primary" @endif>
    Click me
</flux:button>
```

**Use dynamic attributes instead:**
```blade
<!-- RIGHT -->
<flux:button :variant="$show ? 'primary' : 'default'">
    Click me
</flux:button>
```

### 2. Dynamic Expressions in Attributes

Blade component opening tags have limited expression support. Use Alpine or Livewire properties for complex logic:

```blade
<!-- Limited -->
<flux:input :placeholder="$placeholder" />

<!-- Complex: use Livewire property -->
<flux:input :placeholder="$this->getPlaceholder()" />
```

### 3. Class Specificity

When your Tailwind classes conflict with Flux's built-in styles, conflicts can occur. Use the `!` modifier or publish components locally:

```blade
<!-- Using ! modifier -->
<flux:button class="bg-custom! border-2!">Custom</flux:button>

<!-- OR: Publish and modify locally -->
php artisan flux:publish
```

---

## Design Pattern Summary

| Pattern | Purpose | Example |
|---------|---------|---------|
| Simplicity first | Easy defaults | `<flux:input label="Name" />` |
| Composable alternatives | Complex cases | Nest components for flexibility |
| Consistent naming | Reduced learning curve | All inputs follow same pattern |
| Brevity | Short, clear names | `input` not `text-input` |
| Icon support | Visual enhancement | `icon="check"`, `icon:trailing="arrow"` |
| Variant system | Visual variations | `variant="primary"`, `variant="danger"` |
| Data binding | Livewire integration | `wire:model="property"` |
| "We style, you space" | Clear responsibility | Flux handles styling, you handle layout |

---

Last updated: January 2026
