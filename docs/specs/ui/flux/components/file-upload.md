# flux:file-upload

File upload with drag-and-drop, previews, and Livewire integration.

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `wire:model` | string | - | Binds upload to Livewire property |
| `name` | string | - | Input name for form submissions |
| `multiple` | boolean | false | Allow multiple file selection |
| `label` | string | - | Field label above upload area |
| `description` | string | - | Helper text below field |
| `error` | string | - | Validation error message |
| `disabled` | boolean | false | Prevents interaction |

**Data Attributes:**
- `data-dragging` - Added when files dragged over
- `data-loading` - Added during upload

## Child Components

### flux:file-upload.dropzone

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `heading` | string | - | Main dropzone text |
| `text` | string | - | Supporting text (file restrictions) |
| `icon` | string | cloud-arrow-up | Dropzone icon |
| `inline` | boolean | false | Compact horizontal layout |
| `with-progress` | boolean | false | Show progress bar during upload |

**CSS Variables:**
- `--flux-file-upload-progress` - Upload percentage (e.g., "42%")
- `--flux-file-upload-progress-as-string` - Quoted percentage

### flux:file-item

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `heading` | string | - | File name/title |
| `text` | string | auto | Additional text |
| `image` | string | - | Preview image URL |
| `size` | number | - | File size in bytes (auto-formatted) |
| `icon` | string | document | Icon when no image |
| `invalid` | boolean | false | Error state styling |

**Slots:** `actions` (action buttons)

### flux:file-item.remove

Pre-styled remove button for file items.

## Basic Usage

```blade
<flux:file-upload wire:model="file">
    <flux:file-upload.dropzone
        heading="Upload a file"
        text="PNG, JPG, PDF up to 10MB"
    />
</flux:file-upload>
```

## Multiple Files

```blade
<flux:file-upload wire:model="files" multiple>
    <flux:file-upload.dropzone
        heading="Upload files"
        text="Select multiple files"
    />
</flux:file-upload>
```

## With Progress

```blade
<flux:file-upload wire:model="file">
    <flux:file-upload.dropzone
        heading="Upload a file"
        text="Uploading..."
        with-progress
    />
</flux:file-upload>
```

## Inline Layout

```blade
<flux:file-upload wire:model="file">
    <flux:file-upload.dropzone inline heading="Choose file" />
</flux:file-upload>
```

## Displaying Uploaded Files

```blade
<flux:file-upload wire:model="files" multiple>
    <flux:file-upload.dropzone heading="Upload files" />

    @foreach ($files as $index => $file)
        <flux:file-item
            :heading="$file->getClientOriginalName()"
            :size="$file->getSize()"
        >
            <x-slot name="actions">
                <flux:file-item.remove wire:click="removeFile({{ $index }})" />
            </x-slot>
        </flux:file-item>
    @endforeach
</flux:file-upload>
```

## With Image Preview

```blade
<flux:file-item
    heading="photo.jpg"
    :image="$file->temporaryUrl()"
    :size="$file->getSize()"
>
    <x-slot name="actions">
        <flux:file-item.remove wire:click="removeFile" />
    </x-slot>
</flux:file-item>
```

## Livewire Integration

```php
use Livewire\WithFileUploads;

class FileUploader extends Component
{
    use WithFileUploads;

    public $file;
    public $files = [];

    public function save()
    {
        $this->validate([
            'file' => 'required|file|max:10240', // 10MB
        ]);

        $path = $this->file->store('uploads');
    }

    public function removeFile($index)
    {
        unset($this->files[$index]);
        $this->files = array_values($this->files);
    }
}
```

## Custom Dropzone

```blade
<flux:file-upload wire:model="file">
    <div
        class="border-2 border-dashed p-8 text-center cursor-pointer"
        :class="{ 'border-blue-500': $el.closest('[data-dragging]') }"
    >
        Click or drag to upload
    </div>
</flux:file-upload>
```
