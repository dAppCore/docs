# Flux Pro Components Reference

Documentation for Flux Pro components extracted from fluxui.dev.

---

## Autocomplete

A searchable input with suggestion items.

**Component:** `flux:autocomplete`

**Props:**
| Prop | Description |
|------|-------------|
| `wire:model` | Binds input value to Livewire property |
| `label` | Label displayed above input |
| `description` | Descriptive text below label |
| `placeholder` | Text when input empty |
| `size` | sm, xs |
| `variant` | filled (default: outline) |
| `disabled` | Prevents interaction |
| `readonly` | Makes input read-only |
| `invalid` | Applies error styling |
| `icon` | Icon at input start |
| `icon:trailing` | Icon at input end |
| `kbd` | Keyboard shortcut hint |
| `clearable` | Shows clear button |
| `copyable` | Shows copy button |
| `container:class` | Classes for container |
| `class:input` | Classes for input element |

**Child Components:**
- `flux:autocomplete.item` - Individual suggestion item (supports `disabled`)

**Usage:**
```blade
<flux:autocomplete wire:model="state" label="State">
    <flux:autocomplete.item>Alabama</flux:autocomplete.item>
    <flux:autocomplete.item>California</flux:autocomplete.item>
</flux:autocomplete>
```

---

## Calendar

Date selection with single, multiple, or range modes.

**Component:** `flux:calendar`

**Props:**
| Prop | Description |
|------|-------------|
| `wire:model` | Binds to Livewire property |
| `value` | Selected date(s) in Y-m-d format |
| `mode` | single, multiple, range |
| `min` | Earliest date (or "today") |
| `max` | Latest date (or "today") |
| `size` | base, xs, sm, lg, xl, 2xl |
| `start-day` | Week start (0-6) |
| `months` | Number of months displayed |
| `min-range` | Minimum days for range |
| `max-range` | Maximum days for range |
| `open-to` | Initial calendar date |
| `force-open-to` | Force to open-to date |
| `navigation` | Show/hide month navigation |
| `static` | Non-interactive display |
| `multiple` | Enable multiple selection |
| `week-numbers` | Display week numbers |
| `selectable-header` | Month/year dropdowns |
| `with-today` | Today navigation button |
| `with-inputs` | Date input fields |
| `locale` | Locale string (fr, en-US, ja-JP) |

**Usage:**
```blade
<flux:calendar wire:model="date" />
<flux:calendar multiple wire:model="dates" />
<flux:calendar mode="range" wire:model="range" />
```

---

## Chart

Lightweight charting for line and area charts.

**Component:** `flux:chart`

**Props:**
| Prop | Description |
|------|-------------|
| `wire:model` | Binds to Livewire property with data |
| `value` | Array of data points |
| `curve` | smooth (default), none |
| `class` | Container classes (aspect-3/1, h-64) |

**Child Components:**
- `flux:chart.svg` - Container for SVG elements
- `flux:chart.line` - Renders line (`field` prop)
- `flux:chart.area` - Filled area beneath line
- `flux:chart.point` - Dots marking data points
- `flux:chart.axis` - Configure axes (x/y)
- `flux:chart.cursor` - Interactive vertical guide
- `flux:chart.tooltip` - Contextual data display
- `flux:chart.legend` - Multiple series identifier
- `flux:chart.summary` - Key metrics display
- `flux:chart.viewport` - Wraps SVG for siblings
- `flux:chart.axis.line` - Baseline
- `flux:chart.axis.grid` - Gridlines
- `flux:chart.axis.tick` - Tick marks/labels
- `flux:chart.axis.mark` - Tick marks only

**Usage:**
```blade
<flux:chart :value="$data" class="aspect-3/1">
    <flux:chart.svg>
        <flux:chart.line field="visitors" />
        <flux:chart.point field="visitors" />
    </flux:chart.svg>
</flux:chart>
```

---

## Command

Searchable command palette interface.

**Component:** `flux:command`

**Child Components:**
- `flux:command.input` - Search input (`clearable`, `closable`, `icon`, `placeholder`)
- `flux:command.items` - Container for items
- `flux:command.item` - Individual item (`icon`, `icon:variant`, `kbd`)

**Usage:**
```blade
<flux:command>
    <flux:command.input placeholder="Search..." />
    <flux:command.items>
        <flux:command.item icon="user-plus" kbd="⌘A">Assign to...</flux:command.item>
    </flux:command.items>
</flux:command>
```

---

## Composer

Message input for chat interfaces and AI prompts.

**Component:** `flux:composer`

**Props:**
| Prop | Description |
|------|-------------|
| `wire:model` | Binds to Livewire property |
| `name` | Name for validation |
| `placeholder` | Text when empty |
| `label` | Label text |
| `label:sr-only` | Screen-reader only label |
| `description` | Help text |
| `rows` | Visible text lines (default: 2) |
| `max-rows` | Maximum expandable rows |
| `inline` | Actions alongside input |
| `submit` | cmd+enter (default), enter |
| `disabled` | Prevents interaction |
| `invalid` | Error styling |
| `variant` | input (matches form inputs) |

**Slots:**
- `input` - Replace textarea with custom editor
- `header` - Content above input (file previews)
- `footer` - Content below input
- `actionsLeading` - Buttons at start
- `actionsTrailing` - Buttons at end (submit)

---

## Context Menu

Right-click context menu functionality.

**Component:** `flux:context`

**Props:**
| Prop | Description |
|------|-------------|
| `wire:model` | Binds menu state |
| `position` | Placement (top/bottom start/center/end) |
| `gap` | Distance from click (default: 4) |
| `offset` | Additional offset [x] [y] |
| `target` | ID of external menu element |
| `disabled` | Prevents context menu |

**Usage:**
```blade
<flux:context>
    <div>Right-click here</div>
    <flux:menu>
        <flux:menu.item>Action 1</flux:menu.item>
    </flux:menu>
</flux:context>
```

---

## Date Picker

Input-triggered calendar for date selection.

**Component:** `flux:date-picker`

**Props:**
| Prop | Description |
|------|-------------|
| `wire:model` | Binds to Livewire property |
| `value` | Selected date(s) |
| `mode` | single (default), range |
| `min-range` | Minimum days in range |
| `max-range` | Maximum days in range |
| `min` | Earliest date |
| `max` | Latest date |
| `open-to` | Default opening date |
| `force-open-to` | Force to date |
| `months` | Number of months |
| `label` | Label text |
| `description` | Help text |
| `badge` | Badge on label |
| `placeholder` | Text when empty |
| `size` | sm, lg, xl, 2xl |
| `start-day` | Week start (0-6) |
| `week-numbers` | Show week numbers |
| `selectable-header` | Month/year dropdowns |
| `with-today` | Today button |
| `with-presets` | Display preset ranges |
| `presets` | Space-separated presets |
| `clearable` | Show clear button |
| `disabled` | Prevents interaction |
| `invalid` | Error styling |
| `locale` | Locale string |
| `unavailable` | Unavailable dates |

**Child Components:**
- `flux:date-picker.input` - Custom trigger input
- `flux:date-picker.button` - Custom trigger button

---

## Editor

Rich text editor built on ProseMirror/Tiptap.

**Component:** `flux:editor`

**Props:**
| Prop | Description |
|------|-------------|
| `wire:model` | Binds content |
| `value` | Initial content |
| `label` | Label above editor |
| `description` | Help text |
| `description:trailing` | Help text below |
| `badge` | Badge on label |
| `placeholder` | Text when empty |
| `toolbar` | Space-separated items |
| `disabled` | Prevents interaction |
| `invalid` | Error styling |

**Child Components:**
- `flux:editor.toolbar` - Container (`items` prop)
- `flux:editor.button` - Toolbar button (`icon`, `iconVariant`, `tooltip`, `disabled`)
- `flux:editor.content` - Editable content area

**Toolbar Items:** heading, bold, italic, strike, underline, bullet, ordered, blockquote, code, link, align, undo, redo

**Usage:**
```blade
<flux:editor wire:model="content" toolbar="heading | bold italic | undo redo" />
```

---

## File Upload

Drag-and-drop file upload with preview.

**Component:** `flux:file-upload`

**Props:**
| Prop | Description |
|------|-------------|
| `wire:model` | Binds to Livewire property |
| `name` | Input name |
| `multiple` | Multiple files |
| `label` | Field label |
| `description` | Helper text |
| `error` | Validation error |
| `disabled` | Prevents interaction |

**Child Components:**
- `flux:file-upload.dropzone` - Drop zone (`heading`, `text`, `icon`, `inline`, `with-progress`)
- `flux:file-item` - Uploaded file display (`heading`, `text`, `image`, `size`, `icon`, `invalid`)
- `flux:file-item.remove` - Removal button

---

## Kanban

Workflow visualisation with draggable cards in columns.

**Component:** `flux:kanban`

**Child Components:**
- `flux:kanban.column` - Individual column
- `flux:kanban.column.header` - Column title (`heading`, `subheading`, `count`, `badge`)
- `flux:kanban.column.cards` - Card container
- `flux:kanban.card` - Individual card (`heading`, `as`)
- `flux:kanban.column.footer` - Optional footer

**Slots:** header, footer, actions, default

---

## Pillbox

Multi-select with pills/tags display.

**Component:** `flux:pillbox`

**Props:**
| Prop | Description |
|------|-------------|
| `wire:model` | Binds to Livewire property |
| `placeholder` | Text when empty |
| `label` | Label text |
| `description` | Help text |
| `size` | sm |
| `searchable` | Add search input |
| `search:placeholder` | Search placeholder |
| `filter` | false for server-side |
| `disabled` | Prevents interaction |
| `invalid` | Error styling |
| `variant` | combobox |
| `multiple` | Multiple selections |

**Child Components:**
- `flux:pillbox.option` - Selectable option
- `flux:pillbox.option.create` - Create new option
- `flux:pillbox.option.empty` - Empty state
- `flux:pillbox.search` - Custom search
- `flux:pillbox.trigger` - Custom trigger
- `flux:pillbox.input` - Custom input

---

## Slider

Range selection control.

**Component:** `flux:slider`

**Props:**
| Prop | Description |
|------|-------------|
| `wire:model` | Binds to Livewire property |
| `range` | Two thumbs for range |
| `min` | Minimum value |
| `max` | Maximum value |
| `step` | Step increment |
| `big-step` | Step with shift key |
| `min-steps-between` | Minimum thumb distance |
| `track:class` | Track CSS classes |
| `thumb:class` | Thumb CSS classes |

**Child Components:**
- `flux:slider.tick` - Marks at values (`value` prop)

---

## Time Picker

Time selection input.

**Component:** `flux:time-picker`

**Props:**
| Prop | Description |
|------|-------------|
| `wire:model` | Binds to Livewire property |
| `value` | Selected time(s) H:i |
| `type` | input, button (default) |
| `multiple` | Multiple times |
| `time-format` | auto, 12-hour, 24-hour |
| `interval` | Minutes between options (30) |
| `min` | Earliest time or "now" |
| `max` | Latest time or "now" |
| `unavailable` | Unavailable times |
| `open-to` | Opening time |
| `label` | Label text |
| `description` | Help text |
| `badge` | Label badge |
| `placeholder` | Empty text |
| `size` | sm, xs |
| `clearable` | Clear button |
| `disabled` | Prevents interaction |
| `invalid` | Error styling |
| `locale` | Locale string |

---

## Select Pro Features

### Listbox Variant (Custom Select)
```blade
<flux:select variant="listbox" clearable>
    <flux:option value="1" icon="star">Premium</flux:option>
</flux:select>
```

### Searchable Select
```blade
<flux:select searchable empty="No results">
    <flux:option>...</flux:option>
</flux:select>
```

### Multiple Select
```blade
<flux:select variant="listbox" multiple selected-suffix="items">
    <flux:option>...</flux:option>
</flux:select>
```

### Combobox
```blade
<flux:select variant="combobox" :filter="false">
    <flux:option>...</flux:option>
    <flux:select.option.create wire:click="create">Create new</flux:select.option.create>
</flux:select>
```

---

## Checkbox Cards

**Component:** `flux:checkbox.group` with `variant="cards"`

```blade
<flux:checkbox.group variant="cards" wire:model="selected">
    <flux:checkbox value="opt1" label="Option 1" description="Description" />
</flux:checkbox.group>
```

---

## Radio Cards

**Component:** `flux:radio.group` with `variant="cards"`

```blade
<flux:radio.group variant="cards" wire:model="selected">
    <flux:radio value="opt1" label="Option 1" description="Description" icon="star" />
</flux:radio.group>
```

---

## Table Pro Features

### Sortable Columns
```blade
<flux:table.column sortable :sorted="$sortBy === 'name'" :direction="$sortDirection" wire:click="sort('name')">
    Name
</flux:table.column>
```

### Sticky Headers/Columns
```blade
<flux:table.columns sticky>...</flux:table.columns>
<flux:table.column sticky>...</flux:table.column>
```

### Pagination
```blade
<flux:table :paginate="$users">...</flux:table>
```

---

## Tabs Pro Features

### Variants
- `variant="segmented"` - Segmented control style
- `variant="pills"` - Pill/button style

### Scrollable
```blade
<flux:tabs scrollable scrollable:fade>...</flux:tabs>
```
