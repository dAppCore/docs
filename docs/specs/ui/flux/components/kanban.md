# flux:kanban (Pro)

Workflow visualisation with draggable cards organised into columns.

## Child Components

### flux:kanban

Container for all columns.

### flux:kanban.column

Individual workflow stage column.

### flux:kanban.column.header

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `heading` | string | - | Column title |
| `subheading` | string | - | Secondary text |
| `count` | number | - | Card count display |
| `badge` | string | - | Badge content |

**Slots:** `default` (overrides heading/count), `actions` (right-aligned buttons)

### flux:kanban.column.cards

Container for cards within a column.

### flux:kanban.column.footer

Bottom section for "Add card" forms/buttons.

### flux:kanban.card

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `heading` | string | - | Card title |
| `as` | string | div | `button`, `div` (button enables click) |

**Slots:** `default` (content), `header` (badges/tags), `footer` (avatars/metadata)

## Basic Usage

```blade
<flux:kanban>
    <flux:kanban.column>
        <flux:kanban.column.header heading="To Do" :count="3" />
        <flux:kanban.column.cards>
            <flux:kanban.card heading="Design homepage" />
            <flux:kanban.card heading="Write copy" />
            <flux:kanban.card heading="Review assets" />
        </flux:kanban.column.cards>
    </flux:kanban.column>

    <flux:kanban.column>
        <flux:kanban.column.header heading="In Progress" :count="2" />
        <flux:kanban.column.cards>
            <flux:kanban.card heading="Build components" />
            <flux:kanban.card heading="API integration" />
        </flux:kanban.column.cards>
    </flux:kanban.column>

    <flux:kanban.column>
        <flux:kanban.column.header heading="Done" :count="1" />
        <flux:kanban.column.cards>
            <flux:kanban.card heading="Project setup" />
        </flux:kanban.column.cards>
    </flux:kanban.column>
</flux:kanban>
```

## With Header Actions

```blade
<flux:kanban.column.header heading="To Do" :count="5">
    <x-slot name="actions">
        <flux:dropdown>
            <flux:button icon="ellipsis-horizontal" variant="ghost" size="sm" />
            <flux:menu>
                <flux:menu.item icon="plus">Add card</flux:menu.item>
                <flux:menu.item icon="archive-box">Archive column</flux:menu.item>
            </flux:menu>
        </flux:dropdown>
    </x-slot>
</flux:kanban.column.header>
```

## Cards with Metadata

```blade
<flux:kanban.card heading="Design system">
    <x-slot name="header">
        <flux:badge color="blue" size="sm">Design</flux:badge>
    </x-slot>
    <x-slot name="footer">
        <flux:avatar.group>
            <flux:avatar size="xs" src="/avatars/1.jpg" />
            <flux:avatar size="xs" src="/avatars/2.jpg" />
        </flux:avatar.group>
        <flux:text size="sm" class="text-zinc-500">Due Dec 15</flux:text>
    </x-slot>
</flux:kanban.card>
```

## Interactive Cards

```blade
<flux:kanban.card as="button" wire:click="openCard({{ $card->id }})" heading="{{ $card->title }}" />
```

## With Footer Actions

```blade
<flux:kanban.column>
    <flux:kanban.column.header heading="Backlog" />
    <flux:kanban.column.cards>
        {{-- Cards --}}
    </flux:kanban.column.cards>
    <flux:kanban.column.footer>
        <flux:button icon="plus" variant="ghost" class="w-full">Add card</flux:button>
    </flux:kanban.column.footer>
</flux:kanban.column>
```
