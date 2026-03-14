# Studio Multimedia Pipeline

Studio is a CorePHP module that orchestrates video remixing, transcription, voice synthesis, and image generation by dispatching GPU work to remote services. It separates creative decisions (LEM/Ollama) from mechanical execution (ffmpeg, Whisper, TTS, ComfyUI).

## Architecture

Studio is a job orchestrator, not a renderer. All GPU-intensive work runs on remote Docker services accessed over HTTP.

```
Studio Module (CorePHP)
  ├── Livewire UI (asset browser, remix form, voice, thumbnails)
  ├── Artisan Commands (CLI)
  └── API Routes (/api/studio/*)
        │
  Actions (CatalogueAsset, GenerateManifest, RenderManifest, etc.)
        │
  Redis Job Queue
        │
        ├── Ollama (LEM) ─────── Creative decisions, scripts, manifests
        ├── Whisper ───────────── Speech-to-text transcription
        ├── Kokoro TTS ────────── Voiceover generation
        ├── ffmpeg Worker ─────── Video rendering from manifests
        └── ComfyUI ──────────── Image generation, thumbnails
```

### Smart/Dumb Separation

LEM produces JSON manifests (the creative layer). ffmpeg and GPU services consume them mechanically (the execution layer). Neither side knows about the other's internals — the manifest format is the contract.

## Module Structure

The Studio module lives at `app/Mod/Studio/` and follows standard CorePHP patterns:

```
app/Mod/Studio/
├── Boot.php                    # Lifecycle events (API, Console, Web)
├── Actions/
│   ├── CatalogueAsset.php      # Ingest files, extract metadata
│   ├── TranscribeAsset.php     # Send to Whisper, store transcript
│   ├── GenerateManifest.php    # Brief + library → LEM → manifest JSON
│   ├── RenderManifest.php      # Dispatch manifest to ffmpeg worker
│   ├── SynthesiseSpeech.php    # Text → TTS → audio file
│   ├── GenerateVoiceover.php   # Script → voiced audio for remix
│   ├── GenerateImage.php       # Prompt → ComfyUI → image
│   ├── GenerateThumbnail.php   # Asset → thumbnail image
│   └── BatchRemix.php          # Queue multiple remix jobs
├── Console/
│   ├── Catalogue.php           # studio:catalogue — batch ingest
│   ├── Transcribe.php          # studio:transcribe — batch transcription
│   ├── Remix.php               # studio:remix — brief in, video out
│   ├── Voice.php               # studio:voice — text-to-speech
│   ├── Thumbnail.php           # studio:thumbnail — generate thumbnails
│   └── BatchRemixCommand.php   # studio:batch-remix — queue batch jobs
├── Controllers/Api/
│   ├── AssetController.php     # GET/POST /api/studio/assets
│   ├── RemixController.php     # POST /api/studio/remix
│   ├── VoiceController.php     # POST /api/studio/voice
│   └── ImageController.php     # POST /api/studio/images/thumbnail
├── Models/
│   ├── StudioAsset.php         # Multimedia asset with metadata
│   └── StudioJob.php           # Job tracking (status, manifest, output)
├── Livewire/
│   ├── AssetBrowserPage.php    # Browse/search/tag assets
│   ├── RemixPage.php           # Remix form + job status
│   ├── VoicePage.php           # Voice synthesis interface
│   └── ThumbnailPage.php       # Thumbnail generator
└── Routes/
    ├── api.php                 # REST API endpoints
    └── web.php                 # Livewire page routes
```

## Asset Cataloguing

Assets are multimedia files (video, image, audio) tracked in the `studio_assets` table with metadata including duration, resolution, tags, and transcripts.

### Ingesting Assets

```php
use Mod\Studio\Actions\CatalogueAsset;

// From an uploaded file
$asset = CatalogueAsset::run($uploadedFile, ['summer', 'beach']);

// From an existing storage path
$asset = CatalogueAsset::run('studio/raw/clip-001.mp4', ['interview']);
```

Only `video/*`, `image/*`, and `audio/*` MIME types are accepted.

### CLI Batch Ingest

```bash
php artisan studio:catalogue /path/to/media --tags=summer,promo
```

### Querying Assets

```php
use Mod\Studio\Models\StudioAsset;

// By type
$videos = StudioAsset::videos()->get();
$images = StudioAsset::images()->get();
$audio = StudioAsset::audio()->get();

// By tag
$summer = StudioAsset::tagged('summer')->get();
```

## Transcription

Transcription sends assets to a Whisper service and stores the returned text and detected language.

```php
use Mod\Studio\Actions\TranscribeAsset;

$asset = TranscribeAsset::run($asset);

echo $asset->transcript;           // "Hello and welcome..."
echo $asset->transcript_language;  // "en"
```

The action handles missing files and API failures gracefully — it returns the asset unchanged without throwing.

### CLI Batch Transcription

```bash
php artisan studio:transcribe
```

## Manifest-Driven Remixing

The remix pipeline has two stages: manifest generation (creative) and rendering (mechanical).

### Generating Manifests

```php
use Mod\Studio\Actions\GenerateManifest;

$job = GenerateManifest::run(
    brief: 'Create a 15-second upbeat TikTok from the summer footage',
    template: 'tiktok-15s',
);

// $job->manifest contains the JSON manifest
```

The action collects all video assets from the library, sends them as context to Ollama along with the brief, and parses the returned JSON manifest.

### Manifest Format

```json
{
  "clips": [
    {"asset_id": 42, "start_ms": 3200, "end_ms": 8100, "effects": ["fade_in"]},
    {"asset_id": 17, "start_ms": 0, "end_ms": 5500, "effects": ["crossfade"]}
  ],
  "audio": {"track": "original"},
  "voiceover": {"script": "Summer vibes only", "voice": "default", "volume": 0.8},
  "overlays": [
    {"type": "image", "asset_id": 5, "at": 0.5, "duration": 3.0, "position": "bottom-right", "opacity": 0.8}
  ]
}
```

### Rendering

```php
use Mod\Studio\Actions\RenderManifest;

$job = RenderManifest::run($job);
```

This dispatches the manifest to the ffmpeg worker service, which renders the video and calls back when complete.

### CLI Remix

```bash
php artisan studio:remix "Create a relaxing travel montage" --template=tiktok-30s
```

## Voice & TTS

```php
use Mod\Studio\Actions\SynthesiseSpeech;

$audio = SynthesiseSpeech::run(
    text: 'Welcome to our channel',
    voice: 'default',
);
```

### CLI

```bash
php artisan studio:voice "Welcome to our channel" --voice=default
```

## Image Generation

Thumbnails and image overlays use ComfyUI:

```php
use Mod\Studio\Actions\GenerateThumbnail;

$thumbnail = GenerateThumbnail::run($asset);
```

### CLI

```bash
php artisan studio:thumbnail --asset=42
```

## API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/api/studio/assets` | List assets |
| `GET` | `/api/studio/assets/{id}` | Show asset details |
| `POST` | `/api/studio/assets` | Upload/catalogue asset |
| `POST` | `/api/studio/remix` | Submit remix brief |
| `GET` | `/api/studio/remix/{id}` | Poll job status |
| `POST` | `/api/studio/remix/{id}/callback` | Worker completion callback |
| `POST` | `/api/studio/voice` | Submit voice synthesis |
| `GET` | `/api/studio/voice/{id}` | Poll voice job status |
| `POST` | `/api/studio/images/thumbnail` | Generate thumbnail |

## GPU Services

All GPU services run as Docker containers, accessed over HTTP. Configuration is in `config/studio.php`:

| Service | Default Endpoint | Purpose |
|---|---|---|
| Ollama | `http://studio-ollama:11434` | Creative decisions via LEM |
| Whisper | `http://studio-whisper:9100` | Speech-to-text |
| Kokoro TTS | `http://studio-tts:9200` | Text-to-speech |
| ffmpeg Worker | `http://studio-worker:9300` | Video rendering |
| ComfyUI | `http://studio-comfyui:8188` | Image generation |

## Configuration

```php
// config/studio.php
return [
    'ollama' => [
        'url' => env('STUDIO_OLLAMA_URL', 'http://studio-ollama:11434'),
        'model' => env('STUDIO_OLLAMA_MODEL', 'lem-4b'),
        'timeout' => 60,
    ],
    'whisper' => [
        'url' => env('STUDIO_WHISPER_URL', 'http://studio-whisper:9100'),
        'model' => 'large-v3-turbo',
        'timeout' => 120,
    ],
    'worker' => [
        'url' => env('STUDIO_WORKER_URL', 'http://studio-worker:9300'),
        'timeout' => 300,
    ],
    'storage' => [
        'disk' => 'local',
        'assets_path' => 'studio/assets',
    ],
    'templates' => [
        'tiktok-15s' => ['duration' => 15, 'width' => 1080, 'height' => 1920, 'fps' => 30],
        'tiktok-30s' => ['duration' => 30, 'width' => 1080, 'height' => 1920, 'fps' => 30],
        'youtube-60s' => ['duration' => 60, 'width' => 1920, 'height' => 1080, 'fps' => 30],
    ],
];
```

## Livewire UI

Studio provides four Livewire page components:

- **Asset Browser** — browse, search, and tag multimedia assets
- **Remix Page** — enter a creative brief, select template, view job progress
- **Voice Page** — text-to-speech interface
- **Thumbnail Page** — generate thumbnails from assets

Components are registered via the module's Boot class and available under `mod.studio.livewire.*`.

## Testing

All actions are testable with `Http::fake()`:

```php
use Illuminate\Support\Facades\Http;
use Mod\Studio\Actions\TranscribeAsset;
use Mod\Studio\Models\StudioAsset;

it('transcribes an asset via Whisper', function () {
    Storage::fake('local');
    Storage::disk('local')->put('studio/test.mp4', 'fake-video');

    Http::fake([
        '*/transcribe' => Http::response([
            'text' => 'Hello world',
            'language' => 'en',
        ]),
    ]);

    $asset = StudioAsset::factory()->create(['path' => 'studio/test.mp4']);
    $result = TranscribeAsset::run($asset);

    expect($result->transcript)->toBe('Hello world');
    expect($result->transcript_language)->toBe('en');
});
```

## Learn More

- [Actions Pattern](actions.md)
- [Lifecycle Events](/php/framework/events)
