# flux:dropdown

Composable dropdown menus with navigation, actions, checkboxes, radios, and submenus.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `position` | string | bottom | `top`, `right`, `bottom`, `left` |
| `align` | string | start | `start`, `center`, `end` |
| `offset` | number | 0 | Pixel offset from trigger |
| `gap` | number | 4 | Pixel gap between trigger and menu |

## Child Components

### flux:menu

Complex menu with keyboard navigation, submenus, checkboxes, radios.

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `keep-open` | boolean | false | Prevents closure on item click |

### flux:menu.item

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `icon` | string | - | Leading icon name |
| `icon:trailing` | string | - | Trailing icon name |
| `icon:variant` | string | - | `outline`, `solid`, `mini`, `micro` |
| `kbd` | string | - | Keyboard shortcut hint |
| `suffix` | string | - | Trailing text |
| `variant` | string | default | `default`, `danger` |
| `disabled` | boolean | false | Disables interaction |
| `keep-open` | boolean | false | Prevents menu closure |

### flux:menu.submenu

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `heading` | string | - | Submenu label |
| `icon` | string | - | Leading icon |
| `icon:trailing` | string | - | Trailing icon |
| `icon:variant` | string | - | Icon variant |
| `keep-open` | boolean | false | Prevents closure |

### flux:menu.separator

Horizontal line separating menu items. No props.

### flux:menu.checkbox

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Livewire binding |
| `checked` | boolean | false | Default checked state |
| `disabled` | boolean | false | Disables interaction |
| `keep-open` | boolean | false | Prevents menu closure |

### flux:menu.checkbox-group

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Livewire binding for group |

### flux:menu.radio

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `checked` | boolean | false | Default selected state |
| `disabled` | boolean | false | Disables interaction |
| `keep-open` | boolean | false | Prevents menu closure |

### flux:menu.radio.group

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Livewire binding for group |
| `keep-open` | boolean | false | Prevents closure |

## Basic Usage

```blade
<flux:dropdown>
    <flux:button icon:trailing="chevron-down">Options</flux:button>
    <flux:menu>
        <flux:menu.item icon="pencil">Edit</flux:menu.item>
        <flux:menu.item icon="trash" variant="danger">Delete</flux:menu.item>
    </flux:menu>
</flux:dropdown>
```

## With Keyboard Shortcuts

```blade
<flux:dropdown>
    <flux:button>Actions</flux:button>
    <flux:menu>
        <flux:menu.item icon="clipboard" kbd="⌘C">Copy</flux:menu.item>
        <flux:menu.item icon="clipboard-document" kbd="⌘V">Paste</flux:menu.item>
    </flux:menu>
</flux:dropdown>
```

## With Submenu

```blade
<flux:dropdown position="bottom" align="end">
    <flux:button icon:trailing="chevron-down">Options</flux:button>
    <flux:menu>
        <flux:menu.item icon="plus">New post</flux:menu.item>
        <flux:menu.separator />
        <flux:menu.submenu heading="Sort by">
            <flux:menu.radio.group>
                <flux:menu.radio checked>Name</flux:menu.radio>
                <flux:menu.radio>Date</flux:menu.radio>
            </flux:menu.radio.group>
        </flux:menu.submenu>
    </flux:menu>
</flux:dropdown>
```

## With Checkboxes

```blade
<flux:dropdown>
    <flux:button>Filters</flux:button>
    <flux:menu>
        <flux:menu.checkbox wire:model="showActive">Active</flux:menu.checkbox>
        <flux:menu.checkbox wire:model="showArchived">Archived</flux:menu.checkbox>
    </flux:menu>
</flux:dropdown>
```

## Grouped Radios

```blade
<flux:dropdown>
    <flux:button>Sort</flux:button>
    <flux:menu>
        <flux:menu.radio.group wire:model="sortBy">
            <flux:menu.radio value="name">Name</flux:menu.radio>
            <flux:menu.radio value="date">Date</flux:menu.radio>
            <flux:menu.radio value="size">Size</flux:menu.radio>
        </flux:menu.radio.group>
    </flux:menu>
</flux:dropdown>
```
