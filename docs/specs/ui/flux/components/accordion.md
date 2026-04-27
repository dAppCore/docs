# flux:accordion

Collapsible content sections with smooth transitions and exclusive mode.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `variant` | string | - | Set to `reverse` to display icon before heading |
| `transition` | boolean | false | Enable expanding transitions for smoother interactions |
| `exclusive` | boolean | false | Only one accordion item can expand simultaneously |

## Child Components

### flux:accordion.item

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `heading` | string | - | Shorthand for heading content |
| `expanded` | boolean | false | Expand item by default |
| `disabled` | boolean | false | Prevent expansion/collapse |

### flux:accordion.heading

| Slot | Description |
|------|-------------|
| `default` | Heading text content |

### flux:accordion.content

| Slot | Description |
|------|-------------|
| `default` | Content displayed when expanded |

---

## Basic Example

```blade
<flux:accordion>
    <flux:accordion.item>
        <flux:accordion.heading>What's your refund policy?</flux:accordion.heading>
        <flux:accordion.content>
            If you are not satisfied with your purchase, we offer a 30-day money-back guarantee. Please contact our support team for assistance.
        </flux:accordion.content>
    </flux:accordion.item>
    <flux:accordion.item>
        <flux:accordion.heading>Do you offer any discounts for bulk purchases?</flux:accordion.heading>
        <flux:accordion.content>
            Yes, we offer special discounts for bulk orders. Please reach out to our sales team with your requirements.
        </flux:accordion.content>
    </flux:accordion.item>
    <flux:accordion.item>
        <flux:accordion.heading>How do I track my order?</flux:accordion.heading>
        <flux:accordion.content>
            Once your order is shipped, you will receive an email with a tracking number. Use this number to track your order on our website.
        </flux:accordion.content>
    </flux:accordion.item>
</flux:accordion>
```

## Shorthand Syntax

Use the `heading` prop instead of nested components:

```blade
<flux:accordion.item heading="What's your refund policy?">
    If you are not satisfied with your purchase, we offer a 30-day money-back guarantee. Please contact our support team for assistance.
</flux:accordion.item>
```

## Exclusive Mode

Only one section open at a time:

```blade
<flux:accordion exclusive>
    <flux:accordion.item heading="Section 1">Content 1</flux:accordion.item>
    <flux:accordion.item heading="Section 2">Content 2</flux:accordion.item>
    <flux:accordion.item heading="Section 3">Content 3</flux:accordion.item>
</flux:accordion>
```

## With Transitions

Smooth expand/collapse animations:

```blade
<flux:accordion transition>
    <flux:accordion.item heading="Smooth animation">
        This content expands and collapses smoothly.
    </flux:accordion.item>
</flux:accordion>
```

## Reverse Icon Position

Icon on the left side:

```blade
<flux:accordion variant="reverse">
    <flux:accordion.item heading="Icon on left">Content here</flux:accordion.item>
</flux:accordion>
```

## Pre-expanded Items

```blade
<flux:accordion>
    <flux:accordion.item heading="Open by default" expanded>
        This section starts expanded.
    </flux:accordion.item>
</flux:accordion>
```

## Disabled Items

```blade
<flux:accordion>
    <flux:accordion.item heading="Cannot be opened" disabled>
        This section cannot be expanded or collapsed.
    </flux:accordion.item>
</flux:accordion>
```

## Combined Example

```blade
<flux:accordion exclusive transition>
    <flux:accordion.item heading="First section" expanded>
        This starts open and others close when you open them.
    </flux:accordion.item>
    <flux:accordion.item heading="Second section">
        Click to open, first will close.
    </flux:accordion.item>
    <flux:accordion.item heading="Disabled section" disabled>
        Cannot interact with this.
    </flux:accordion.item>
</flux:accordion>
```
