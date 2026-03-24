# flux:text

Typography components for body copy and links.

## Props

### flux:text

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `size` | string | default | `sm`, `default`, `lg`, `xl` |
| `variant` | string | default | `default`, `strong`, `subtle` |
| `color` | string | default | Any Tailwind colour name |
| `inline` | boolean | false | Renders as span instead of p |

### flux:link

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `href` | string | required | Link URL |
| `variant` | string | default | `default`, `ghost`, `subtle` |
| `external` | boolean | false | Opens in new tab |
| `as` | string | a | `a`, `button` |

## Basic Text

```blade
<flux:text>Default paragraph text.</flux:text>
```

## Size Variants

```blade
<flux:text size="sm">Small text</flux:text>
<flux:text>Default text</flux:text>
<flux:text size="lg">Large text</flux:text>
<flux:text size="xl">Extra large text</flux:text>
```

## Text Variants

```blade
<flux:text variant="strong">Strong emphasis text</flux:text>
<flux:text variant="subtle">Subtle muted text</flux:text>
```

## Coloured Text

```blade
<flux:text color="blue">Blue text</flux:text>
<flux:text color="red">Red text</flux:text>
<flux:text color="green">Green text</flux:text>
```

## Inline Text

```blade
<flux:text>
    This is a paragraph with <flux:text inline variant="strong">inline strong</flux:text> text.
</flux:text>
```

## Links

```blade
<flux:link href="/about">About us</flux:link>
```

## Link Variants

```blade
<flux:link href="#" variant="default">Default link</flux:link>
<flux:link href="#" variant="ghost">Ghost link</flux:link>
<flux:link href="#" variant="subtle">Subtle link</flux:link>
```

## External Link

```blade
<flux:link href="https://example.com" external>External site</flux:link>
```

## Link as Button

```blade
<flux:link as="button" wire:click="createAccount">
    Create new account →
</flux:link>
```

## Text with Link

```blade
<flux:text>
    Already have an account? <flux:link href="/login">Sign in</flux:link>
</flux:text>
```

## Heading with Text

```blade
<div>
    <flux:heading size="lg">Welcome back</flux:heading>
    <flux:text class="mt-2">
        Sign in to your account to continue.
    </flux:text>
</div>
```
