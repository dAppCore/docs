# flux:callout

Highlight important information and guide users toward key actions.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `icon` | string | - | Icon name displayed next to heading |
| `icon:variant` | string | - | Icon variant (e.g., `outline`) |
| `variant` | string | secondary | `secondary`, `success`, `warning`, `danger` |
| `color` | string | - | Custom Tailwind colour |
| `inline` | boolean | false | Actions appear inline with content |
| `heading` | string | - | Shorthand for `flux:callout.heading` |
| `text` | string | - | Shorthand for `flux:callout.text` |

### Colour Options

`zinc`, `red`, `orange`, `amber`, `yellow`, `lime`, `green`, `emerald`, `teal`, `cyan`, `sky`, `blue`, `indigo`, `violet`, `purple`, `fuchsia`, `pink`, `rose`

## Slots

| Slot | Description |
|------|-------------|
| `icon` | Custom icon SVG |
| `actions` | Buttons/links (typically `flux:button`) |
| `controls` | UI elements at top right (e.g., close button) |

## Child Components

### flux:callout.heading

| Prop | Type | Description |
|------|------|-------------|
| `icon` | string | Icon inside heading |
| `icon:variant` | string | Icon variant |

### flux:callout.text

Default slot for text content.

### flux:callout.link

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `href` | string | - | URL destination |
| `external` | boolean | false | Opens in new tab |

---

## Basic Usage

```blade
<flux:callout icon="clock">
    <flux:callout.heading>Upcoming maintenance</flux:callout.heading>
    <flux:callout.text>
        We will be performing scheduled maintenance on Saturday.
        <flux:callout.link href="/status">Learn more</flux:callout.link>
    </flux:callout.text>
</flux:callout>
```

## Shorthand Syntax

```blade
<flux:callout
    icon="information-circle"
    heading="Quick tip"
    text="You can use keyboard shortcuts to navigate faster."
/>
```

## Icon Inside Heading

```blade
<flux:callout>
    <flux:callout.heading icon="newspaper">Policy update</flux:callout.heading>
    <flux:callout.text>Our terms of service have been updated.</flux:callout.text>
</flux:callout>
```

## Variants

```blade
<flux:callout variant="secondary" icon="information-circle">
    <flux:callout.heading>Information</flux:callout.heading>
</flux:callout>

<flux:callout variant="success" icon="check-circle">
    <flux:callout.heading>Success</flux:callout.heading>
</flux:callout>

<flux:callout variant="warning" icon="exclamation-circle">
    <flux:callout.heading>Warning</flux:callout.heading>
</flux:callout>

<flux:callout variant="danger" icon="x-circle">
    <flux:callout.heading>Error</flux:callout.heading>
</flux:callout>
```

## With Actions

```blade
<flux:callout icon="credit-card">
    <flux:callout.heading>Subscription expiring soon</flux:callout.heading>
    <flux:callout.text>Your plan expires in 3 days. Renew to avoid interruption.</flux:callout.text>
    <x-slot name="actions">
        <flux:button variant="primary">Renew now</flux:button>
        <flux:button variant="ghost" href="/pricing">View plans</flux:button>
    </x-slot>
</flux:callout>
```

## Inline Actions

Actions appear beside content instead of below:

```blade
<flux:callout icon="cube" variant="secondary" inline>
    <flux:callout.heading>Your package is delayed</flux:callout.heading>
    <x-slot name="actions">
        <flux:button>Track order</flux:button>
    </x-slot>
</flux:callout>
```

## Dismissible

Using Alpine.js for dismiss functionality:

```blade
<flux:callout
    icon="bell"
    variant="secondary"
    inline
    x-data="{ visible: true }"
    x-show="visible"
>
    <flux:callout.heading>Upcoming meeting in 15 minutes</flux:callout.heading>
    <x-slot name="controls">
        <flux:button icon="x-mark" variant="ghost" x-on:click="visible = false" />
    </x-slot>
</flux:callout>
```

## Custom Colour

```blade
<flux:callout icon="sparkles" color="purple">
    <flux:callout.heading>Have a question?</flux:callout.heading>
    <flux:callout.text>Try our AI assistant for instant answers.</flux:callout.text>
</flux:callout>
```

## Custom Icon (SVG Slot)

```blade
<flux:callout>
    <x-slot name="icon">
        <svg class="size-6" viewBox="0 0 24 24" fill="currentColor">
            <!-- Custom SVG content -->
        </svg>
    </x-slot>
    <flux:callout.heading>Custom notification</flux:callout.heading>
</flux:callout>
```

## External Link

```blade
<flux:callout icon="document-text">
    <flux:callout.heading>Documentation available</flux:callout.heading>
    <flux:callout.text>
        Read our comprehensive guide.
        <flux:callout.link href="https://docs.example.com" external>
            View documentation
        </flux:callout.link>
    </flux:callout.text>
</flux:callout>
```
