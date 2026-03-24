# flux:radio

Single selection from mutually exclusive options with multiple visual variants.

## Props

### flux:radio.group

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Binds selection to Livewire property |
| `label` | string | - | Label above group |
| `description` | string | - | Help text below group |
| `variant` | string | default | `default`, `segmented`, `cards`, `pills`, `buttons` |
| `invalid` | boolean | false | Error styling |
| `size` | string | - | `sm` (segmented only) |
| `:indicator` | boolean | true | Show/hide radio indicator |

### flux:radio

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `label` | string | - | Radio label text |
| `description` | string | - | Help text for option |
| `value` | string | - | Value when selected |
| `checked` | boolean | false | Default selected state |
| `disabled` | boolean | false | Prevents interaction |
| `icon` | string | - | Icon (segmented variant) |
| `name` | string | - | Form field name |

### flux:radio.indicator

Used for custom layouts in card variants.

## Basic Usage

```blade
<flux:radio.group wire:model="payment" label="Payment method">
    <flux:radio value="card" label="Credit Card" checked />
    <flux:radio value="paypal" label="PayPal" />
    <flux:radio value="bank" label="Bank Transfer" />
</flux:radio.group>
```

## With Descriptions

```blade
<flux:radio.group wire:model="role" label="Role">
    <flux:radio value="admin" label="Administrator"
        description="Can perform any action." checked />
    <flux:radio value="editor" label="Editor"
        description="Can edit content but not settings." />
    <flux:radio value="viewer" label="Viewer"
        description="Read-only access." />
</flux:radio.group>
```

## Segmented Variant

```blade
<flux:radio.group wire:model="view" variant="segmented">
    <flux:radio label="List" value="list" />
    <flux:radio label="Grid" value="grid" />
    <flux:radio label="Board" value="board" />
</flux:radio.group>
```

## Segmented with Icons

```blade
<flux:radio.group wire:model="view" variant="segmented">
    <flux:radio icon="list-bullet" label="List" value="list" />
    <flux:radio icon="squares-2x2" label="Grid" value="grid" />
    <flux:radio icon="view-columns" label="Board" value="board" />
</flux:radio.group>
```

## Small Segmented

```blade
<flux:radio.group wire:model="view" variant="segmented" size="sm">
    <flux:radio label="List" />
    <flux:radio label="Grid" />
</flux:radio.group>
```

## Cards Variant (Pro)

```blade
<flux:radio.group wire:model="plan" variant="cards" label="Select plan">
    <flux:radio value="starter" label="Starter" description="For individuals" />
    <flux:radio value="pro" label="Pro" description="For teams" />
    <flux:radio value="enterprise" label="Enterprise" description="For organisations" />
</flux:radio.group>
```

## Cards with Custom Content (Pro)

```blade
<flux:radio.group wire:model="shipping" variant="cards" label="Shipping">
    <flux:radio value="standard" checked>
        <flux:radio.indicator />
        <div class="flex-1">
            <flux:heading>Standard</flux:heading>
            <flux:text size="sm">3-5 business days</flux:text>
        </div>
        <flux:text>Free</flux:text>
    </flux:radio>
    <flux:radio value="express">
        <flux:radio.indicator />
        <div class="flex-1">
            <flux:heading>Express</flux:heading>
            <flux:text size="sm">1-2 business days</flux:text>
        </div>
        <flux:text>£9.99</flux:text>
    </flux:radio>
</flux:radio.group>
```

## Cards without Indicators (Pro)

```blade
<flux:radio.group wire:model="plan" variant="cards" :indicator="false">
    <flux:radio value="monthly" label="Monthly" description="£10/month" />
    <flux:radio value="yearly" label="Yearly" description="£100/year" />
</flux:radio.group>
```

## Pills Variant (Pro)

```blade
<flux:radio.group wire:model="size" variant="pills">
    <flux:radio value="xs">XS</flux:radio>
    <flux:radio value="s">S</flux:radio>
    <flux:radio value="m">M</flux:radio>
    <flux:radio value="l">L</flux:radio>
    <flux:radio value="xl">XL</flux:radio>
</flux:radio.group>
```

## Buttons Variant (Pro)

```blade
<flux:radio.group wire:model="align" variant="buttons">
    <flux:radio value="left" icon="bars-3-bottom-left" />
    <flux:radio value="center" icon="bars-3" />
    <flux:radio value="right" icon="bars-3-bottom-right" />
</flux:radio.group>
```
