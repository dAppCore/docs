# Image Optimizer Integration Example

This file shows how to integrate the ImageOptimizer into your upload pipeline.

## Basic Usage in Livewire Components

```php
<?php

namespace App\Livewire\BioLink;

use Core\Media\Image\ImageOptimizer;
use Illuminate\Support\Facades\Storage;
use Livewire\Component;
use Livewire\WithFileUploads;

class UploadBioImage extends Component
{
    use WithFileUploads;

    public $image;

    public function save()
    {
        $this->validate([
            'image' => 'required|image|max:10240', // 10MB
        ]);

        $user = auth()->user();
        $workspace = $user->defaultHostWorkspace();

        // Store the file first
        $path = $this->image->store('biolinks/images/' . $workspace->id, 'local');
        $absolutePath = Storage::path($path);

        // Optimize the image
        $optimizer = app(ImageOptimizer::class);
        $result = $optimizer->optimize($absolutePath);

        // Record optimization statistics
        if ($result->wasSuccessful()) {
            $optimizer->recordOptimization(
                $result,
                $workspace,
                $this->biolink, // or whatever model
            );
        }

        // Continue with your logic...
        $this->biolink->update([
            'image_path' => $path,
        ]);

        $this->dispatch('notify', message: "Image uploaded and optimised ({$result->getSummary()})", type: 'success');
    }
}
```

## Optimizing Uploaded Files Before Storage

```php
public function saveOptimized()
{
    $this->validate([
        'image' => 'required|image|max:10240',
    ]);

    $optimizer = app(ImageOptimizer::class);

    // Optimize the uploaded file directly
    $result = $optimizer->optimizeUploadedFile($this->image);

    // Now store the optimized file
    $path = $this->image->store('biolinks/images', 'local');

    // Record the optimization
    $optimizer->recordOptimization(
        $result,
        auth()->user()->defaultHostWorkspace()
    );
}
```

## Retrieving Statistics

```php
use Core\Media\Image\ImageOptimizer;

// In a controller or Livewire component
public function getOptimizationStats()
{
    $workspace = auth()->user()->defaultHostWorkspace();
    $optimizer = app(ImageOptimizer::class);

    $stats = $optimizer->getStats($workspace);

    // Returns:
    // [
    //     'count' => 150,
    //     'total_original' => 45000000,  // bytes
    //     'total_optimized' => 28000000,
    //     'total_saved' => 17000000,
    //     'average_percentage' => 37.8,
    //     'total_saved_human' => '16.2MB',
    // ]
}
```

## Admin Dashboard Widget

```php
<?php

namespace App\Livewire\Admin;

use Core\Media\Image\ImageOptimizer;
use Livewire\Component;

class ImageOptimizationStats extends Component
{
    public function render()
    {
        $optimizer = app(ImageOptimizer::class);
        $stats = $optimizer->getStats(); // null = all workspaces

        return view('admin.admin.image-optimization-stats', [
            'stats' => $stats,
        ]);
    }
}
```

## Configuration

Add to your `.env`:

```env
IMAGE_OPTIMIZATION_ENABLED=true
IMAGE_OPTIMIZATION_DRIVER=gd
IMAGE_OPTIMIZATION_QUALITY=80
IMAGE_OPTIMIZATION_PNG_COMPRESSION=6
IMAGE_OPTIMIZATION_MIN_SIZE_KB=10
IMAGE_OPTIMIZATION_MAX_SIZE_MB=10
```

## Disabling for Specific Uploads

```php
// Temporarily disable optimization
config(['images.optimization.enabled' => false]);

$path = $this->file->store('uploads');

// Re-enable
config(['images.optimization.enabled' => true]);
```

## Integration Points

The ImageOptimizer should be integrated at these points:

1. **BioLink image uploads** - Avatar, background, block images
2. **Static page images** - Embedded images in HTML content
3. **Theme preview images** - Template and theme gallery
4. **User profile images** - Avatar uploads
5. **Social media uploads** - Before posting to social networks

## Notes

- Optimization happens **in-place** by default (replaces the original)
- Files smaller than 10KB are skipped (configurable)
- Files larger than 10MB are skipped (configurable)
- Supports JPEG, PNG, WebP formats
- Uses GD by default (Imagick support can be added)
- Gracefully handles errors (returns no-op result instead of throwing)
