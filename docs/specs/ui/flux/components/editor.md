# flux:editor

Rich text editor built on ProseMirror and Tiptap with markdown support.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Binds editor content to Livewire property |
| `value` | string | - | Initial content when not using wire:model |
| `label` | string | - | Wraps editor in flux:field with label |
| `description` | string | - | Help text between label and editor |
| `description:trailing` | boolean | false | Displays description below editor |
| `badge` | string | - | Badge text in label component |
| `placeholder` | string | - | Text shown when editor is empty |
| `toolbar` | string | default | Space-separated items, `|` separator, `~` spacer |
| `disabled` | boolean | false | Prevents user interaction |
| `invalid` | boolean | false | Applies error styling |

## Default Toolbar

`heading`, `bold`, `italic`, `strike`, `bullet`, `ordered`, `blockquote`, `link`, `align`

## All Toolbar Items

- `heading` - Heading levels
- `bold` - Bold text
- `italic` - Italic text
- `strike` - Strikethrough
- `underline` - Underline text
- `bullet` - Bullet list
- `ordered` - Numbered list
- `blockquote` - Block quote
- `subscript` - Subscript text
- `superscript` - Superscript text
- `highlight` - Highlighted text
- `link` - Hyperlink
- `code` - Inline code
- `undo` - Undo action
- `redo` - Redo action

## Child Components

### flux:editor.toolbar

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `items` | string | - | Space-separated toolbar configuration |

### flux:editor.button

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `icon` | string | - | Icon name |
| `iconVariant` | string | mini | `mini`, `micro`, `outline` |
| `tooltip` | string | - | Hover text |
| `disabled` | boolean | false | Disable button |

### flux:editor.content

Container for initial HTML content.

### Toolbar Item Components

- `flux:editor.heading`
- `flux:editor.bold`
- `flux:editor.italic`
- `flux:editor.strike`
- `flux:editor.underline`
- `flux:editor.bullet`
- `flux:editor.ordered`
- `flux:editor.blockquote`
- `flux:editor.code`
- `flux:editor.link`
- `flux:editor.align`
- `flux:editor.undo`
- `flux:editor.redo`
- `flux:editor.separator`
- `flux:editor.spacer`

## Keyboard Shortcuts

| Action | Mac | Windows |
|--------|-----|---------|
| Bold | ⌘B | Ctrl+B |
| Italic | ⌘I | Ctrl+I |
| Underline | ⌘U | Ctrl+U |
| Link | ⌘K | Ctrl+K |
| Blockquote | ⌘⇧B | Ctrl+Shift+B |
| Undo | ⌘Z | Ctrl+Z |
| Redo | ⌘⇧Z | Ctrl+Shift+Z |

## Markdown Shortcuts

| Type | Shortcut |
|------|----------|
| H1 | `#` |
| Bold | `**text**` |
| Italic | `*text*` |
| Strike | `~~text~~` |
| Bullet list | `-` |
| Ordered list | `1.` |
| Blockquote | `>` |
| Inline code | `` ` `` |
| Code block | `` ``` `` |

## Basic Usage

```blade
<flux:editor wire:model="content" />
```

## With Label and Placeholder

```blade
<flux:editor
    wire:model="content"
    label="Content"
    placeholder="Start typing..."
/>
```

## Custom Toolbar

```blade
<flux:editor
    wire:model="content"
    toolbar="heading | bold italic | bullet ordered | link ~ undo redo"
/>
```

## Custom Toolbar Components

```blade
<flux:editor>
    <flux:editor.toolbar>
        <flux:editor.heading />
        <flux:editor.separator />
        <flux:editor.bold />
        <flux:editor.italic />
        <flux:editor.spacer />
        <flux:dropdown position="bottom end">
            <flux:editor.button icon="ellipsis-horizontal" />
            <flux:menu>
                <flux:menu.item>More options...</flux:menu.item>
            </flux:menu>
        </flux:dropdown>
    </flux:editor.toolbar>
    <flux:editor.content />
</flux:editor>
```

## Height Configuration

```blade
{{-- Adjust minimum height --}}
<flux:editor class="**:data-[slot=content]:min-h-[100px]!" />

{{-- Adjust maximum height --}}
<flux:editor class="**:data-[slot=content]:max-h-[300px]!" />
```

## Custom Extensions

```blade
<flux:editor x-on:flux:editor="e.detail.registerExtensions([
    // Your custom Tiptap extensions
])" />
```

## Pre-installed Extensions

- Highlight
- Link
- Placeholder
- StarterKit
- Superscript
- Subscript
- TextAlign
- Underline
