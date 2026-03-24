# Flux UI - Component Reference

Complete list of all available Flux UI components organised by category.

## Navigation Components

### flux:navbar
Horizontal navigation bar with automatic active page detection.

| Prop | Values | Default |
|------|--------|---------|
| - | - | - |

**Child Components:**
- `flux:navbar.item` - Individual navigation link

**Props for navbar.item:**
| Prop | Values | Default |
|------|--------|---------|
| `href` | URL string | - |
| `current` | boolean | auto-detected |
| `icon` | Heroicon name | - |
| `icon:trailing` | Heroicon name | - |
| `badge` | string, boolean, slot | - |
| `badge:color` | colour name | - |
| `badge:variant` | solid, outline | solid |

```blade
<flux:navbar>
    <flux:navbar.item href="/" icon="home" current>Home</flux:navbar.item>
    <flux:navbar.item href="/about" icon="information-circle">About</flux:navbar.item>
</flux:navbar>
```

### flux:navlist
Vertical sidebar navigation with collapsible groups.

**Child Components:**
- `flux:navlist.item` - Navigation link
- `flux:navlist.group` - Expandable section

**Props for navlist.item:**
| Prop | Values | Default |
|------|--------|---------|
| `href` | URL string | - |
| `current` | boolean | auto-detected |
| `icon` | Heroicon name | - |
| `badge` | string, boolean | - |
| `badge:color` | colour name | - |

**Props for navlist.group:**
| Prop | Values | Default |
|------|--------|---------|
| `heading` | string | - |
| `expandable` | boolean | false |
| `expanded` | boolean | false |

```blade
<flux:navlist>
    <flux:navlist.item href="/" icon="home">Dashboard</flux:navlist.item>

    <flux:navlist.group heading="Content" expandable>
        <flux:navlist.item href="/posts" icon="document">Posts</flux:navlist.item>
        <flux:navlist.item href="/pages" icon="document-duplicate">Pages</flux:navlist.item>
    </flux:navlist.group>
</flux:navlist>
```

### flux:navmenu
Dropdown menus within navigation items.

Combines with `flux:navbar.item` or `flux:navlist.item`:

```blade
<flux:navbar.item>
    <flux:navmenu>
        <flux:button>Menu</flux:button>
        <flux:menu>
            <flux:menu.item href="/option1">Option 1</flux:menu.item>
            <flux:menu.item href="/option2">Option 2</flux:menu.item>
        </flux:menu>
    </flux:navmenu>
</flux:navbar.item>
```

### flux:dropdown
Generic expandable menu for various use cases.

| Prop | Values | Default |
|------|--------|---------|
| `position` | top, right, bottom, left | bottom |
| `align` | start, center, end | start |
| `offset` | pixels | 0 |
| `gap` | pixels | 4 |

**Child Components:**
- `flux:menu` - Menu container
- `flux:menu.item` - Individual menu item
- `flux:menu.separator` - Visual divider

```blade
<flux:dropdown position="bottom" align="end">
    <flux:button>Actions</flux:button>

    <flux:menu>
        <flux:menu.item icon="pencil">Edit</flux:menu.item>
        <flux:menu.item icon="trash" variant="danger">Delete</flux:menu.item>
    </flux:menu>
</flux:dropdown>
```

### flux:menu
Complex action menus with keyboard navigation, submenus, and checkboxes.

**Props:**
| Prop | Values | Default |
|------|--------|---------|
| `keep-open` | boolean | false |

**Child Components:**
- `flux:menu.item` - Menu item
- `flux:menu.checkbox` - Checkbox option
- `flux:menu.radio` - Radio option
- `flux:menu.submenu` - Submenu section
- `flux:menu.separator` - Divider
- `flux:menu.group` - Grouped items

**Props for menu.item:**
| Prop | Values | Default |
|------|--------|---------|
| `icon` | Heroicon name | - |
| `icon:trailing` | Heroicon name | - |
| `icon:variant` | outline, solid, mini, micro | outline |
| `kbd` | string | - |
| `suffix` | string | - |
| `variant` | default, danger | default |
| `disabled` | boolean | false |
| `keep-open` | boolean | false |

---

## Form Components

### flux:input
Text input field with variants and icons.

| Prop | Values | Default |
|------|--------|---------|
| `wire:model` | Livewire property | - |
| `label` | string | - |
| `description` | string | - |
| `description:trailing` | string | - |
| `type` | text, email, password, date, file, etc. | text |
| `placeholder` | string | - |
| `disabled` | boolean | false |
| `readonly` | boolean | false |
| `invalid` | boolean | false |
| `size` | sm, xs | default |
| `variant` | filled, outline | outline |
| `icon` | Heroicon name | - |
| `icon:trailing` | Heroicon name | - |
| `clearable` | boolean | false |
| `copyable` | boolean | false |
| `viewable` | boolean | false (password toggle) |
| `multiple` | boolean | false (file input) |
| `mask` | Alpine mask pattern | - |

```blade
<flux:input wire:model="email" type="email" label="Email" />
<flux:input icon="magnifying-glass" placeholder="Search..." />
<flux:input type="password" viewable label="Password" />
```

### flux:select
Dropdown select field.

| Prop | Values | Default |
|------|--------|---------|
| `wire:model` | Livewire property | - |
| `label` | string | - |
| `description` | string | - |
| `disabled` | boolean | false |
| `invalid` | boolean | false |
| `size` | sm | default |
| `variant` | filled | outline |
| `multiple` | boolean | false |
| `searchable` | boolean | false |

**Slot Options:**
```blade
<flux:select wire:model="status" label="Status">
    <option value="active">Active</option>
    <option value="inactive">Inactive</option>
</flux:select>
```

### flux:checkbox
Individual checkbox input.

| Prop | Values | Default |
|------|--------|---------|
| `wire:model` | Livewire property | - |
| `value` | string/boolean | - |
| `disabled` | boolean | false |

**Group variant: flux:checkbox.group**
```blade
<flux:checkbox.group wire:model="tags">
    <flux:checkbox value="php">PHP</flux:checkbox>
    <flux:checkbox value="laravel">Laravel</flux:checkbox>
</flux:checkbox.group>
```

### flux:radio
Individual radio button.

| Prop | Values | Default |
|------|--------|---------|
| `wire:model` | Livewire property | - |
| `value` | string | - |
| `disabled` | boolean | false |

**Group variant: flux:radio.group**
```blade
<flux:radio.group wire:model="visibility">
    <flux:radio value="public">Public</flux:radio>
    <flux:radio value="private">Private</flux:radio>
</flux:radio.group>
```

### flux:switch
Toggle switch/checkbox.

| Prop | Values | Default |
|------|--------|---------|
| `wire:model` | Livewire property | - |
| `disabled` | boolean | false |

```blade
<flux:switch wire:model="isActive" />
```

### flux:textarea
Multi-line text input.

| Prop | Values | Default |
|------|--------|---------|
| `wire:model` | Livewire property | - |
| `label` | string | - |
| `description` | string | - |
| `placeholder` | string | - |
| `disabled` | boolean | false |
| `readonly` | boolean | false |
| `invalid` | boolean | false |
| `rows` | number | 3 |

```blade
<flux:textarea wire:model="bio" label="Bio" rows="5" />
```

### flux:field
Form field wrapper combining label, input, and descriptions.

| Prop | Values | Default |
|------|--------|---------|
| - | - | - |

**Child Components:**
- `flux:label` - Field label
- `flux:description` - Help text

```blade
<flux:field>
    <flux:label>Email</flux:label>
    <flux:input type="email" wire:model="email" />
    <flux:description>We'll never share your email.</flux:description>
</flux:field>
```

---

## Button Components

### flux:button
Powerful, composable button component.

| Prop | Values | Default |
|------|--------|---------|
| `variant` | default, primary, filled, danger, ghost, subtle | default |
| `size` | xs, sm, default, lg | default |
| `color` | zinc, red, orange, amber, yellow, lime, green, emerald, teal, cyan, sky, blue, indigo, violet, purple, fuchsia, pink, rose | zinc |
| `icon` | Heroicon name | - |
| `icon:trailing` | Heroicon name | - |
| `icon:variant` | outline, solid, mini, micro | outline |
| `href` | URL (renders as link) | - |
| `disabled` | boolean | false |
| `loading` | boolean | auto-detected with wire:click |
| `square` | boolean | false |
| `inset` | string (top, bottom, start, end) | - |
| `as` | button, a, div | auto |
| `tooltip` | string | - |
| `tooltip:position` | top, right, bottom, left | - |

**Group variant: flux:button.group**
```blade
<flux:button.group>
    <flux:button>One</flux:button>
    <flux:button>Two</flux:button>
</flux:button.group>
```

```blade
<flux:button>Default</flux:button>
<flux:button variant="primary" icon="check">Submit</flux:button>
<flux:button size="sm" variant="ghost">Small Ghost</flux:button>
<flux:button href="/page">Link Button</flux:button>
```

---

## Data Display Components

### flux:table
Structured data display with sorting and pagination.

| Prop | Values | Default |
|------|--------|---------|
| `paginate` | Laravel paginator | - |
| `container:class` | CSS classes | - |

**Child Components:**
- `flux:table.columns` - Header row
- `flux:table.column` - Column header
- `flux:table.rows` - Body wrapper
- `flux:table.row` - Data row
- `flux:table.cell` - Individual cell

**Props for table.column:**
| Prop | Values | Default |
|------|--------|---------|
| `align` | start, center, end | start |
| `sortable` | boolean | false |
| `sorted` | boolean | false |
| `direction` | asc, desc | - |
| `sticky` | boolean | false |

**Props for table.cell:**
| Prop | Values | Default |
|------|--------|---------|
| `align` | start, center, end | start |
| `variant` | default, strong | default |
| `sticky` | boolean | false |

```blade
<flux:table>
    <flux:table.columns>
        <flux:table.column>Name</flux:table.column>
        <flux:table.column align="center" sortable>Status</flux:table.column>
    </flux:table.columns>

    <flux:table.rows>
        @foreach ($users as $user)
            <flux:table.row>
                <flux:table.cell>{{ $user->name }}</flux:table.cell>
                <flux:table.cell align="center">
                    <flux:badge color="green">Active</flux:badge>
                </flux:table.cell>
            </flux:table.row>
        @endforeach
    </flux:table.rows>
</flux:table>
```

### flux:badge
Status and category highlighting.

| Prop | Values | Default |
|------|--------|---------|
| `color` | zinc, red, orange, amber, yellow, lime, green, emerald, teal, cyan, sky, blue, indigo, violet, purple, fuchsia, pink, rose | zinc |
| `size` | sm, default, lg | default |
| `variant` | default, pill, solid | default |
| `icon` | Heroicon name | - |
| `icon:trailing` | Heroicon name | - |
| `icon:variant` | outline, solid, mini, micro | mini |
| `inset` | string (top, bottom, start, end) | - |
| `as` | button, div | div |

**Closeable variant: flux:badge.close**
```blade
<flux:badge color="blue">Active
    <flux:badge.close />
</flux:badge>

<flux:badge variant="pill" color="green" size="lg">Success</flux:badge>
<flux:badge variant="solid" color="red" icon="exclamation-triangle">Alert</flux:badge>
```

### flux:avatar
User profile images or initials.

| Prop | Values | Default |
|------|--------|---------|
| `src` | image URL | - |
| `initials` | string (1-2 chars) | - |
| `alt` | text | - |
| `size` | xs, sm, default, lg, xl | default |

```blade
<flux:avatar src="/path/to/avatar.jpg" alt="User name" />
<flux:avatar initials="JD" />
```

### flux:heading
Hierarchical text headings.

| Prop | Values | Default |
|------|--------|---------|
| `level` | 1-6 | 2 |
| `size` | default, lg, xl | default |

```blade
<flux:heading level="1">Page Title</flux:heading>
<flux:heading level="2" size="lg">Section Heading</flux:heading>
```

### flux:text
Formatted text content.

| Prop | Values | Default |
|------|--------|---------|
| `size` | xs, sm, default, lg | default |
| `variant` | default, muted | default |
| `weight` | normal, medium, semibold, bold | normal |

```blade
<flux:text size="sm" variant="muted">Helper text</flux:text>
```

---

## Layout Components

### flux:card
Content container.

| Prop | Values | Default |
|------|--------|---------|
| `size` | sm | default |
| `class` | CSS classes | - |

```blade
<flux:card class="space-y-6">
    <flux:heading>Title</flux:heading>
    <p>Content goes here</p>
</flux:card>
```

### flux:sidebar
Persistent sidebar layout.

| Prop | Values | Default |
|------|--------|---------|
| `sticky` | boolean | false |
| `collapsible` | boolean, mobile | false |
| `breakpoint` | pixels | 1024 |
| `persist` | boolean | true |

**Child Components:**
- `flux:sidebar.header` - Top section
- `flux:sidebar.brand` - Logo/branding
- `flux:sidebar.collapse` - Collapse toggle
- `flux:sidebar.search` - Search input
- `flux:sidebar.nav` - Navigation container
- `flux:sidebar.item` - Navigation link
- `flux:sidebar.group` - Grouped items
- `flux:sidebar.spacer` - Vertical spacer
- `flux:sidebar.profile` - User profile section
- `flux:sidebar.toggle` - Mobile toggle button

```blade
<flux:sidebar sticky collapsible="mobile">
    <flux:sidebar.header>
        <flux:sidebar.brand>AppName</flux:sidebar.brand>
        <flux:sidebar.collapse />
    </flux:sidebar.header>

    <flux:sidebar.nav>
        <flux:sidebar.item href="/" icon="home" current>Dashboard</flux:sidebar.item>
    </flux:sidebar.nav>
</flux:sidebar>
```

### flux:header
Top navigation bar.

Can be used alongside sidebar for secondary navigation.

---

## Overlay Components

### flux:modal
Dialog box overlay.

| Prop | Values | Default |
|------|--------|---------|
| `name` | string | - |
| `maxWidth` | sm, md, lg, xl, 2xl | xl |

```blade
<flux:modal name="confirm-delete" maxWidth="sm">
    <flux:heading>Confirm Deletion</flux:heading>
    <p>Are you sure?</p>

    <div class="flex gap-2">
        <flux:button variant="danger" wire:click="delete">Delete</flux:button>
        <flux:button variant="ghost" x-on:click="close()">Cancel</flux:button>
    </div>
</flux:modal>
```

### flux:popover
Small overlay panel.

| Prop | Values | Default |
|------|--------|---------|
| `position` | top, right, bottom, left | bottom |
| `align` | start, center, end | start |

```blade
<flux:popover>
    <flux:button>Show Info</flux:button>
    <div>Information content</div>
</flux:popover>
```

### flux:tooltip
Hover help text.

| Prop | Values | Default |
|------|--------|---------|
| `text` | string | - |
| `position` | top, right, bottom, left | top |

```blade
<flux:button tooltip="Save changes" tooltip:position="top">Save</flux:button>
```

### flux:toast
Notification messages.

| Prop | Values | Default |
|------|--------|---------|
| `title` | string | - |
| `description` | string | - |
| `variant` | default, success, warning, danger | default |
| `icon` | Heroicon name | - |
| `actions` | array | - |

---

## Other Components

### flux:icon
SVG icon display.

| Prop | Values | Default |
|------|--------|---------|
| `name` | Heroicon name | - |
| `variant` | outline, solid, mini, micro | outline |
| `size` | xs, sm, default, lg, xl, 2xl | default |

```blade
<flux:icon name="check-circle" variant="solid" size="lg" />
```

### flux:accordion
Collapsible content sections with smooth transitions and exclusive mode.

| Prop | Values | Default |
|------|--------|---------|
| `variant` | reverse | - (icon after heading) |
| `transition` | boolean | false |
| `exclusive` | boolean | false |

**Child Components:**
- `flux:accordion.item` - Accordion section
- `flux:accordion.heading` - Item heading (or use `heading` prop shorthand)
- `flux:accordion.content` - Item content

**Props for accordion.item:**
| Prop | Values | Default |
|------|--------|---------|
| `heading` | string | - (shorthand for heading element) |
| `expanded` | boolean | false |
| `disabled` | boolean | false |

**Basic example:**
```blade
<flux:accordion>
    <flux:accordion.item>
        <flux:accordion.heading>What's your refund policy?</flux:accordion.heading>
        <flux:accordion.content>
            If you are not satisfied with your purchase, we offer a 30-day money-back guarantee.
        </flux:accordion.content>
    </flux:accordion.item>
    <flux:accordion.item>
        <flux:accordion.heading>Do you offer bulk discounts?</flux:accordion.heading>
        <flux:accordion.content>
            Yes, we offer special discounts for bulk orders. Contact our sales team.
        </flux:accordion.content>
    </flux:accordion.item>
</flux:accordion>
```

**Shorthand syntax (heading prop):**
```blade
<flux:accordion.item heading="What's your refund policy?">
    If you are not satisfied with your purchase, we offer a 30-day money-back guarantee.
</flux:accordion.item>
```

**Exclusive mode (only one open at a time):**
```blade
<flux:accordion exclusive>
    <flux:accordion.item heading="Section 1">Content 1</flux:accordion.item>
    <flux:accordion.item heading="Section 2">Content 2</flux:accordion.item>
</flux:accordion>
```

**With transitions:**
```blade
<flux:accordion transition>
    <flux:accordion.item heading="Smooth animation">
        This content expands and collapses smoothly.
    </flux:accordion.item>
</flux:accordion>
```

**Reverse icon position:**
```blade
<flux:accordion variant="reverse">
    <flux:accordion.item heading="Icon on left">Content here</flux:accordion.item>
</flux:accordion>
```

**Pre-expanded and disabled items:**
```blade
<flux:accordion>
    <flux:accordion.item heading="Open by default" expanded>
        This section starts expanded.
    </flux:accordion.item>
    <flux:accordion.item heading="Cannot be opened" disabled>
        This section cannot be expanded or collapsed.
    </flux:accordion.item>
</flux:accordion>
```

### flux:separator
Visual divider line.

```blade
<div>Section 1</div>
<flux:separator />
<div>Section 2</div>
```

---

## New Components (Recent Additions)

- **flux:composer** - Rich text editor
- **flux:kanban** - Kanban board layout
- **flux:otp-input** - One-time password input
- **flux:pillbox** - Pill-shaped input tags
- **flux:skeleton** - Loading placeholder
- **flux:slider** - Range slider input

---

## Component Naming Convention

| Component Type | Naming Pattern | Examples |
|----------------|----------------|----------|
| Standalone | Single word | input, button, card |
| Parent-child | Dot notation | accordion.item, table.cell |
| Groups | Component.group | button.group, checkbox.group |
| Subcomponents | dot notation | navbar.item, navlist.group |

---

## Props Categories

| Category | Applies To | Examples |
|----------|-----------|----------|
| **Styling** | Most components | variant, size, color, class |
| **Icons** | Navigation, buttons, badges | icon, icon:trailing, icon:variant |
| **Form** | Input components | wire:model, label, description, invalid |
| **State** | Interactive components | disabled, readonly, current, active |
| **Layout** | Layout components | sticky, collapsible, breakpoint |
| **Positioning** | Dropdowns, popovers | position, align, offset, gap |

---

Last updated: January 2026
