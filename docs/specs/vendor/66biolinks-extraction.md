# 66biolinks Theme & Block Extraction Guide

This document describes how to extract themes and blocks from 66biolinks (AltumCode) for use in Host Hub.

## Source Software

- **Product**: 66biolinks by AltumCode
- **Local Path**: `/Users/snider/Code/Software/66biolinks/product`
- **Plugins Path**: `/Users/snider/Code/Software/altum-plugins`
- **Updates**: Monthly releases via AltumCode

## Database Schema

### biolinks_themes table

```sql
CREATE TABLE biolinks_themes (
    biolink_theme_id INT PRIMARY KEY,
    name VARCHAR(64),
    settings JSON,
    is_enabled TINYINT(1),
    `order` INT,
    datetime DATETIME,
    last_datetime DATETIME
);
```

### Theme Settings JSON Structure

```json
{
    "additional": {
        "custom_css": "string|null",
        "custom_js": "string|null"
    },
    "biolink": {
        "background_type": "preset|gradient|color|image",
        "background": "preset_name|image_filename|null",
        "background_color_one": "#hex",
        "background_color_two": "#hex",
        "font": "font_key",
        "font_size": 16,
        "background_blur": 0,
        "background_brightness": 100,
        "width": 8,
        "block_spacing": 2,
        "hover_animation": "false|smooth|instant"
    },
    "biolink_block": {
        "text_color": "#hex",
        "title_color": "#hex",
        "description_color": "#hex",
        "background_color": "#hex",
        "border_width": "0",
        "border_color": "#hex",
        "border_radius": "rounded|rounded-lg|rounded-pill|straight",
        "border_style": "solid|dashed|dotted",
        "border_shadow_style": "none|small|medium|large",
        "border_shadow_color": "#hex"
    },
    "biolink_block_socials": {
        "color": "#hex",
        "background_color": "#hex",
        "border_radius": "rounded|rounded-lg|rounded-pill|straight"
    },
    "biolink_block_paragraph": {
        "text_color": "#hex",
        "background_color": "#hex",
        "border_radius": "rounded|rounded-lg|rounded-pill|straight",
        "border_shadow_style": "none|small|medium|large",
        "border_shadow_color": "#hex"
    },
    "biolink_block_heading": {
        "text_color": "#hex"
    }
}
```

## Extraction Methods

### Method 1: Admin JSON Export (Recommended)

1. Run 66biolinks locally at `http://66biolinks.test`
2. Log in as admin
3. Go to **Admin > Biolinks Themes**
4. Use the **Export** dropdown and select **JSON**
5. Save the exported file to `storage/app/imports/66biolinks-themes.json`
6. Run the converter: `php artisan bio:import-66biolinks-themes`

### Method 2: Direct Database Query

If you have database access:

```sql
SELECT
    biolink_theme_id,
    name,
    settings,
    is_enabled,
    `order`
FROM biolinks_themes
WHERE is_enabled = 1
ORDER BY `order`;
```

Export as JSON and save to `storage/app/imports/66biolinks-themes.json`.

### Method 3: Browser DevTools

1. Open a biolink page with theme applied
2. Open DevTools > Elements
3. Copy computed styles from `.link-body` and `.link-btn` elements
4. Manually create theme entry

## Background Presets

The software includes preset gradients in `app/includes/biolink_backgrounds.php`:

### Gradient Presets (23 total)

| Key | Description |
|-----|-------------|
| zero | Purple/Pink/Orange diagonal |
| one | Cyan to Purple |
| two | Orange to Red |
| three | Green to Blue |
| four | Pink gradient |
| five | Blue gradient |
| six | Pink to Purple |
| seven | Dark blue gradient |
| eight | Purple radial |
| nine | Light blue to purple |
| ten | Pink to cyan |
| eleven | Teal to purple |
| twelve | Yellow to pink |
| thirteen | Pink to dark purple |
| fourteen | Dark blue to teal radial |
| fifteen | Purple to dark blue |
| sixteen | Light cream to blue |
| seventeen | Off-white gradient |
| eighteen | Pastel multi-color |
| nineteen | Hot pink radial |
| twenty | Blue radial |
| twentyone | Lavender radial |
| twentytwo | Magenta to purple |
| twentythree | Pink radial |

### Abstract Presets (8 total)

Complex multi-layer gradients with blend modes - these are the "premium animated" ones.

## Field Mapping: 66biolinks → Host Hub

| 66biolinks | Host Hub |
|------------|----------|
| `biolink.background_type` | `background.type` |
| `biolink.background_color_one` | `background.color` / `background.gradient_start` |
| `biolink.background_color_two` | `background.gradient_end` |
| `biolink.font` | `font_family` (needs font name lookup) |
| `biolink_block.text_color` | `button.text_color` |
| `biolink_block.background_color` | `button.background_color` |
| `biolink_block.border_radius` | `button.border_radius` (convert to px) |
| `biolink_block.border_width` | `button.border_width` |
| `biolink_block.border_color` | `button.border_color` |
| `biolink_block_heading.text_color` | `text_color` |

### Border Radius Conversion

| 66biolinks | Host Hub |
|------------|----------|
| `straight` | `0px` |
| `rounded` | `4px` |
| `rounded-lg` | `8px` |
| `rounded-pill` | `50px` |

## Monthly Update Process

1. **Download new release** from AltumCode
2. **Extract** to `/Users/snider/Code/Software/66biolinks/product`
3. **Check for new themes**: Look for changes in admin theme management
4. **Export themes** using Method 1
5. **Run converter**: `php artisan bio:import-66biolinks-themes`
6. **Review** imported themes in BioThemeSeeder
7. **Test** on local environment
8. **Commit** changes to seeder

## Files Reference

### 66biolinks Source Files

- `app/includes/biolink_backgrounds.php` - Gradient presets
- `app/includes/admin_primary_themes.php` - Tailwind color palettes
- `app/includes/admin_gray_themes.php` - Neutral color palettes
- `app/models/BiolinksThemes.php` - Theme model
- `app/controllers/admin/AdminBiolinkThemeUpdate.php` - Theme settings schema

### Host Hub Target Files

- `app/Mod/WebPage/Database/Seeders/BioThemeSeeder.php` - Theme definitions
- `app/Mod/WebPage/Models/Theme.php` - Theme model
- `app/Mod/WebPage/Console/Commands/Import66BiolinksThemes.php` - Converter command

## Premium Animated Themes

The "abstract" presets use CSS blend modes for animated/layered effects:

```css
background: linear-gradient(...), radial-gradient(...), ...;
background-blend-mode: overlay, color-dodge, difference, ...;
```

These require special handling - store as `background.type: 'advanced'` with raw CSS in `background.css`.
