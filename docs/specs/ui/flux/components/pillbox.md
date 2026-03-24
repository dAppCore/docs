# flux:pillbox (Pro)

Multi-select with removable pill tags, search, and dynamic option creation.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Binds to Livewire property (array) |
| `placeholder` | string | - | Text when no pills selected |
| `label` | string | - | Label above pillbox |
| `description` | string | - | Help text below |
| `size` | string | md | `sm` |
| `searchable` | boolean | false | Adds search input |
| `search:placeholder` | string | Search... | Search placeholder |
| `filter` | boolean | true | Client-side filtering (false for server-side) |
| `disabled` | boolean | false | Prevents interaction |
| `invalid` | boolean | false | Error styling |
| `variant` | string | default | `combobox` |

## Slots

| Slot | Description |
|------|-------------|
| `default` | pillbox.option components |
| `trigger` | Custom dropdown trigger |
| `search` | Custom search input |
| `empty` | No options content |

## Child Components

### flux:pillbox.option

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `value` | mixed | - | Value when selected |
| `label` | string | - | Display text |
| `selected-label` | string | - | Text when selected |
| `disabled` | boolean | false | Prevents selection |
| `filterable` | boolean | true | Hidden by search if false |

### flux:pillbox.input

Custom input for combobox variant.

### flux:pillbox.option.create

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `min-length` | number | - | Min chars before showing |
| `modal` | string | - | Modal to open on select |
| `wire:click` | string | - | Livewire action on select |

### flux:pillbox.option.empty

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `when-loading` | string | - | Message during loading |

### flux:pillbox.search

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `placeholder` | string | Search... | Placeholder text |
| `icon` | string | magnifying-glass | Search icon |
| `clearable` | boolean | true | Show clear button |

### flux:pillbox.trigger

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `placeholder` | string | - | Placeholder text |
| `invalid` | boolean | false | Error styling |
| `size` | string | - | `sm` |
| `clearable` | boolean | false | Clear all button |

## Basic Usage

```blade
<flux:pillbox wire:model="selectedTags" placeholder="Choose tags...">
    <flux:pillbox.option value="design">Design</flux:pillbox.option>
    <flux:pillbox.option value="development">Development</flux:pillbox.option>
    <flux:pillbox.option value="marketing">Marketing</flux:pillbox.option>
</flux:pillbox>
```

## Searchable

```blade
<flux:pillbox wire:model="skills" searchable placeholder="Choose skills...">
    <flux:pillbox.option value="php">PHP</flux:pillbox.option>
    <flux:pillbox.option value="javascript">JavaScript</flux:pillbox.option>
    <flux:pillbox.option value="python">Python</flux:pillbox.option>
    <flux:pillbox.option value="rust">Rust</flux:pillbox.option>
</flux:pillbox>
```

## With Icons

```blade
<flux:pillbox placeholder="Choose platforms...">
    <flux:pillbox.option value="github">
        <div class="flex items-center gap-2">
            <flux:icon.code-bracket variant="mini" /> GitHub
        </div>
    </flux:pillbox.option>
    <flux:pillbox.option value="gitlab">
        <div class="flex items-center gap-2">
            <flux:icon.cube variant="mini" /> GitLab
        </div>
    </flux:pillbox.option>
</flux:pillbox>
```

## Combobox with Create

```blade
<flux:pillbox wire:model="selectedTags" variant="combobox">
    <x-slot name="input">
        <flux:pillbox.input wire:model.live="search" placeholder="Choose tags..." />
    </x-slot>

    @foreach ($this->tags as $tag)
        <flux:pillbox.option :value="$tag->id">{{ $tag->name }}</flux:pillbox.option>
    @endforeach

    <flux:pillbox.option.create wire:click="createTag" min-length="2">
        Create "<span wire:text="search"></span>"
    </flux:pillbox.option.create>

    <flux:pillbox.option.empty>No tags found</flux:pillbox.option.empty>
</flux:pillbox>
```

```php
public string $search = '';

#[Computed]
public function tags()
{
    return Tag::where('name', 'like', "%{$this->search}%")->get();
}

public function createTag()
{
    $tag = Tag::create(['name' => $this->search]);
    $this->selectedTags[] = $tag->id;
    $this->search = '';
}
```

## Modal Integration

```blade
<flux:pillbox.option.create modal="create-tag">
    Create new tag
</flux:pillbox.option.create>

<flux:modal name="create-tag" class="md:w-96">
    <flux:heading>Create Tag</flux:heading>
    <flux:input wire:model="newTagName" label="Tag name" />
    <flux:button wire:click="saveTag">Create</flux:button>
</flux:modal>
```

## Size Variant

```blade
<flux:pillbox size="sm" placeholder="Choose tags...">
    {{-- Options --}}
</flux:pillbox>
```

## Server-side Search

Disable client filtering for backend search:

```blade
<flux:pillbox wire:model.live="selected" :filter="false">
    <x-slot name="input">
        <flux:pillbox.input wire:model.live="search" />
    </x-slot>

    @foreach ($this->results as $result)
        <flux:pillbox.option :value="$result->id">{{ $result->name }}</flux:pillbox.option>
    @endforeach
</flux:pillbox>
```
