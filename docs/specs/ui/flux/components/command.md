# flux:command

Searchable command palette for quick access to actions, often displayed as a modal.

## Child Components

### flux:command.input

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `placeholder` | string | - | Input placeholder |
| `clearable` | boolean | false | Show clear button |
| `closable` | boolean | false | Show close button |
| `icon` | string | magnifying-glass | Leading icon |

### flux:command.items

Container for command items.

### flux:command.item

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `icon` | string | - | Item icon |
| `icon:variant` | string | - | `outline`, `solid`, `mini`, `micro` |
| `kbd` | string | - | Keyboard shortcut hint |

## Basic Usage

```blade
<flux:command>
    <flux:command.input placeholder="Search commands..." />
    <flux:command.items>
        <flux:command.item wire:click="assign" icon="user-plus" kbd="⌘A">
            Assign to...
        </flux:command.item>
        <flux:command.item icon="document-text" kbd="⌘D">
            Documents
        </flux:command.item>
        <flux:command.item icon="cog-6-tooth" kbd="⌘,">
            Settings
        </flux:command.item>
    </flux:command.items>
</flux:command>
```

## In Modal

```blade
<flux:modal name="command" variant="bare">
    <flux:command class="w-full max-w-lg">
        <flux:command.input placeholder="Search..." closable />
        <flux:command.items>
            <flux:command.item icon="home">Dashboard</flux:command.item>
            <flux:command.item icon="users">Users</flux:command.item>
        </flux:command.items>
    </flux:command>
</flux:modal>
```

Trigger with keyboard: `x-on:keydown.cmd.k.window="$flux.modal('command').show()"`
