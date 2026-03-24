# Flux UI - Forms & Input Components

Complete guide to building forms with Flux components, including all input types, validation, and styling.

## Form Structure

### Basic Form with Field Wrapper

The `flux:field` component wraps inputs with labels and descriptions:

```blade
<form wire:submit="submit" class="space-y-6">
    <flux:field>
        <flux:label>Email</flux:label>
        <flux:input
            type="email"
            wire:model="email"
            placeholder="you@example.com"
        />
        <flux:description>We'll never share your email.</flux:description>
    </flux:field>

    <flux:button type="submit">Send</flux:button>
</form>
```

### Shorthand Form Input

Many inputs support shorthand syntax combining label and description:

```blade
<form wire:submit="submit" class="space-y-6">
    <!-- Shorthand: label is built-in -->
    <flux:input
        type="email"
        wire:model="email"
        label="Email"
        description="We'll never share your email."
        placeholder="you@example.com"
    />

    <!-- Equivalent to: -->
    <flux:field>
        <flux:label>Email</flux:label>
        <flux:input type="email" wire:model="email" />
        <flux:description>We'll never share your email.</flux:description>
    </flux:field>

    <flux:button type="submit">Send</flux:button>
</form>
```

---

## Input Component

The fundamental form input component supporting multiple types and variants.

### flux:input Props

| Prop | Values | Default | Description |
|------|--------|---------|-------------|
| `type` | text, email, password, date, time, number, file, search, tel, url, etc. | text | HTML input type |
| `wire:model` | Livewire property | - | Two-way data binding |
| `label` | string | - | Field label (shorthand) |
| `description` | string | - | Help text above input |
| `description:trailing` | string | - | Help text below input |
| `placeholder` | string | - | Placeholder text |
| `disabled` | boolean | false | Disable interaction |
| `readonly` | boolean | false | Lock input (submission allowed) |
| `invalid` | boolean | false | Error state |
| `size` | sm, xs | default | Sizing |
| `variant` | filled, outline | outline | Visual style |
| `icon` | Heroicon name | - | Leading icon |
| `icon:trailing` | Heroicon name | - | Trailing icon |
| `clearable` | boolean | false | Show clear button |
| `copyable` | boolean | false | Show copy button (HTTPS only) |
| `viewable` | boolean | false | Toggle password visibility |
| `multiple` | boolean | false | Select multiple files (file input) |
| `mask` | Alpine mask pattern | - | Input masking |
| `class` | string | - | Additional CSS classes |

### Text Input

**Basic text input:**
```blade
<flux:input
    wire:model="name"
    label="Full Name"
    placeholder="John Doe"
    icon="user"
/>
```

**Text input with trailing icon:**
```blade
<flux:input
    wire:model="verified"
    icon="shield-check"
    icon:trailing="check-circle"
    label="Verification"
/>
```

### Email Input

```blade
<flux:input
    type="email"
    wire:model="email"
    label="Email Address"
    placeholder="you@example.com"
    icon="envelope"
    description="Used for account recovery"
/>
```

### Password Input

**Basic password field:**
```blade
<flux:input
    type="password"
    wire:model="password"
    label="Password"
    icon="lock-closed"
/>
```

**Password with visibility toggle:**
```blade
<flux:input
    type="password"
    wire:model="password"
    label="Password"
    icon="lock-closed"
    viewable
    description="At least 8 characters, including uppercase and symbols"
/>
```

### Number Input

```blade
<flux:input
    type="number"
    wire:model="age"
    label="Age"
    min="18"
    max="120"
/>
```

### Date & Time Inputs

**Date picker:**
```blade
<flux:input
    type="date"
    wire:model="birth_date"
    label="Date of Birth"
/>
```

**Time input:**
```blade
<flux:input
    type="time"
    wire:model="appointment_time"
    label="Appointment Time"
/>
```

### File Input

**Single file:**
```blade
<flux:input
    type="file"
    wire:model="avatar"
    label="Profile Photo"
    accept="image/*"
/>
```

**Multiple files:**
```blade
<flux:input
    type="file"
    wire:model="documents"
    label="Upload Documents"
    multiple
    accept=".pdf,.doc,.docx"
/>
```

### Search Input

```blade
<flux:input
    type="search"
    wire:model.live="query"
    icon="magnifying-glass"
    placeholder="Search..."
    label="Search Users"
/>
```

### Input with Clearable Button

```blade
<flux:input
    type="text"
    wire:model="search"
    placeholder="Type to search..."
    clearable
    icon="magnifying-glass"
/>
```

### Input with Copy Button

```blade
<flux:input
    type="text"
    value="https://invite.example.com/abc123"
    label="Invite Link"
    copyable
    readonly
    icon="link"
/>
```

### Input Masking

Uses Alpine's mask plugin for formatted input:

```blade
<!-- Phone number mask -->
<flux:input
    type="tel"
    wire:model="phone"
    label="Phone Number"
    mask="(999) 999-9999"
    placeholder="(555) 123-4567"
/>

<!-- Credit card mask -->
<flux:input
    type="text"
    wire:model="card"
    label="Card Number"
    mask="9999 9999 9999 9999"
    placeholder="1234 5678 9012 3456"
/>

<!-- Dynamic mask for currency -->
<flux:input
    type="text"
    wire:model="amount"
    label="Amount"
    mask:dynamic="$money($input)"
    placeholder="$0.00"
/>
```

### Input with Errors

```blade
<flux:input
    type="email"
    wire:model="email"
    label="Email"
    :invalid="$errors->has('email')"
/>

@error('email')
    <flux:description class="text-red-600">
        {{ $message }}
    </flux:description>
@enderror
```

### Disabled & Readonly

```blade
<!-- Disabled: can't interact -->
<flux:input
    type="text"
    value="Locked"
    disabled
    label="Disabled Input"
/>

<!-- Readonly: can select/copy but not edit -->
<flux:input
    type="text"
    value="Read-only"
    readonly
    label="Read-only Input"
/>
```

---

## Textarea Component

Multi-line text input for longer content.

### flux:textarea Props

| Prop | Values | Default | Description |
|------|--------|---------|-------------|
| `wire:model` | Livewire property | - | Data binding |
| `label` | string | - | Field label |
| `description` | string | - | Help text |
| `placeholder` | string | - | Placeholder text |
| `rows` | number | 3 | Number of visible rows |
| `disabled` | boolean | false | Disable interaction |
| `readonly` | boolean | false | Prevent editing |
| `invalid` | boolean | false | Error state |
| `class` | string | - | Additional CSS classes |

### Usage Examples

**Basic textarea:**
```blade
<flux:textarea
    wire:model="bio"
    label="Bio"
    rows="4"
    placeholder="Tell us about yourself..."
/>
```

**Textarea with description:**
```blade
<flux:textarea
    wire:model="message"
    label="Message"
    description="Maximum 500 characters"
    rows="6"
    placeholder="Your message here..."
/>
```

**Textarea with error:**
```blade
<flux:textarea
    wire:model="content"
    label="Post Content"
    :invalid="$errors->has('content')"
    rows="8"
/>

@error('content')
    <flux:description class="text-red-600">{{ $message }}</flux:description>
@enderror
```

---

## Select Component

Dropdown select field for choosing from options.

### flux:select Props

| Prop | Values | Default | Description |
|------|--------|---------|-------------|
| `wire:model` | Livewire property | - | Data binding |
| `label` | string | - | Field label |
| `description` | string | - | Help text |
| `disabled` | boolean | false | Disable interaction |
| `invalid` | boolean | false | Error state |
| `size` | sm | default | Sizing |
| `variant` | filled, outline | outline | Visual style |
| `multiple` | boolean | false | Allow multiple selections |
| `searchable` | boolean | false | Filter options by typing |
| `class` | string | - | Additional CSS classes |

### Basic Select

```blade
<flux:select wire:model="status" label="Status">
    <option value="active">Active</option>
    <option value="inactive">Inactive</option>
    <option value="pending">Pending</option>
</flux:select>
```

### Select with Grouped Options

```blade
<flux:select wire:model="country" label="Country">
    <optgroup label="Europe">
        <option value="uk">United Kingdom</option>
        <option value="de">Germany</option>
        <option value="fr">France</option>
    </optgroup>
    <optgroup label="Americas">
        <option value="us">United States</option>
        <option value="ca">Canada</option>
    </optgroup>
</flux:select>
```

### Select with Description

```blade
<flux:select
    wire:model="timezone"
    label="Timezone"
    description="Select your local timezone"
>
    <option value="">-- Choose --</option>
    <option value="utc">UTC</option>
    <option value="gmt-1">GMT-1</option>
    <option value="gmt+1">GMT+1</option>
</flux:select>
```

### Multiple Select

```blade
<flux:select
    wire:model="tags"
    label="Tags"
    multiple
>
    <option value="php">PHP</option>
    <option value="laravel">Laravel</option>
    <option value="livewire">Livewire</option>
    <option value="tailwind">Tailwind</option>
</flux:select>
```

---

## Checkbox Component

Toggle selection for multiple options.

### flux:checkbox Props

| Prop | Values | Default | Description |
|------|--------|---------|-------------|
| `wire:model` | Livewire property | - | Data binding |
| `value` | string, boolean | - | Checkbox value |
| `disabled` | boolean | false | Disable interaction |
| `class` | string | - | Additional CSS classes |

### Single Checkbox

```blade
<flux:checkbox
    wire:model="agree_to_terms"
    value="true"
/>
<flux:label class="ml-2">I agree to the terms and conditions</flux:label>
```

### Checkbox with Label

**Using field wrapper:**
```blade
<flux:field>
    <flux:checkbox wire:model="newsletter" />
    <flux:label>Subscribe to our newsletter</flux:label>
</flux:field>
```

### Multiple Checkboxes (Group)

```blade
<flux:field>
    <flux:label>Interests</flux:label>
    <flux:checkbox.group wire:model="interests">
        <flux:checkbox value="php">PHP</flux:checkbox>
        <flux:checkbox value="laravel">Laravel</flux:checkbox>
        <flux:checkbox value="javascript">JavaScript</flux:checkbox>
        <flux:checkbox value="react">React</flux:checkbox>
    </flux:checkbox.group>
</flux:field>
```

---

## Radio Component

Single selection from multiple options.

### flux:radio Props

| Prop | Values | Default | Description |
|------|--------|---------|-------------|
| `wire:model` | Livewire property | - | Data binding |
| `value` | string | - | Radio value |
| `disabled` | boolean | false | Disable interaction |
| `class` | string | - | Additional CSS classes |

### Single Radio

```blade
<flux:radio
    wire:model="visibility"
    value="public"
/>
<flux:label class="ml-2">Public</flux:label>
```

### Radio Group

```blade
<flux:field>
    <flux:label>Visibility</flux:label>
    <flux:radio.group wire:model="visibility">
        <flux:radio value="public">Public</flux:radio>
        <flux:radio value="private">Private</flux:radio>
        <flux:radio value="unlisted">Unlisted</flux:radio>
    </flux:radio.group>
</flux:field>
```

### Radio Group with Descriptions

```blade
<flux:field>
    <flux:label>Notification Level</flux:label>
    <div class="space-y-3">
        <div class="flex items-start gap-3">
            <flux:radio wire:model="notifications" value="all" />
            <div>
                <flux:label class="mb-0">All notifications</flux:label>
                <flux:description>Get notified about everything</flux:description>
            </div>
        </div>
        <div class="flex items-start gap-3">
            <flux:radio wire:model="notifications" value="important" />
            <div>
                <flux:label class="mb-0">Important only</flux:label>
                <flux:description>Only critical updates</flux:description>
            </div>
        </div>
        <div class="flex items-start gap-3">
            <flux:radio wire:model="notifications" value="none" />
            <div>
                <flux:label class="mb-0">No notifications</flux:label>
                <flux:description>Disable all notifications</flux:description>
            </div>
        </div>
    </div>
</flux:field>
```

---

## Switch Component

Toggle switch for boolean values.

### flux:switch Props

| Prop | Values | Default | Description |
|------|--------|---------|-------------|
| `wire:model` | Livewire property | - | Data binding |
| `disabled` | boolean | false | Disable interaction |
| `class` | string | - | Additional CSS classes |

### Basic Switch

```blade
<div class="flex items-center gap-3">
    <flux:switch wire:model="isActive" />
    <flux:label>Enable feature</flux:label>
</div>
```

### Switch with Description

```blade
<flux:field>
    <div class="flex items-center justify-between">
        <div>
            <flux:label>Dark Mode</flux:label>
            <flux:description>Use dark theme for the interface</flux:description>
        </div>
        <flux:switch wire:model="darkMode" />
    </div>
</flux:field>
```

---

## Complete Form Examples

### Contact Form

```blade
<form wire:submit="submit" class="space-y-6 max-w-md">
    <flux:input
        type="text"
        wire:model="name"
        label="Name"
        icon="user"
        placeholder="Your name"
        :invalid="$errors->has('name')"
    />
    @error('name')
        <flux:description class="text-red-600">{{ $message }}</flux:description>
    @enderror

    <flux:input
        type="email"
        wire:model="email"
        label="Email"
        icon="envelope"
        placeholder="you@example.com"
        :invalid="$errors->has('email')"
    />
    @error('email')
        <flux:description class="text-red-600">{{ $message }}</flux:description>
    @enderror

    <flux:textarea
        wire:model="message"
        label="Message"
        placeholder="Your message here..."
        rows="5"
        :invalid="$errors->has('message')"
    />
    @error('message')
        <flux:description class="text-red-600">{{ $message }}</flux:description>
    @enderror

    <flux:field>
        <flux:checkbox wire:model="consent" />
        <flux:label>I agree to be contacted about this inquiry</flux:label>
    </flux:field>

    <div class="flex gap-3">
        <flux:button type="submit" variant="primary" icon="paper-airplane">
            Send Message
        </flux:button>
        <flux:button type="reset" variant="ghost">Clear</flux:button>
    </div>
</form>
```

### Login Form

```blade
<form wire:submit="login" class="space-y-6 max-w-md">
    <flux:heading level="1">Sign in to your account</flux:heading>

    <flux:input
        type="email"
        wire:model="email"
        label="Email address"
        icon="envelope"
        placeholder="you@example.com"
        autofocus
        :invalid="$errors->has('email')"
    />
    @error('email')
        <flux:description class="text-red-600">{{ $message }}</flux:description>
    @enderror

    <flux:input
        type="password"
        wire:model="password"
        label="Password"
        icon="lock-closed"
        viewable
        :invalid="$errors->has('password')"
    />
    @error('password')
        <flux:description class="text-red-600">{{ $message }}</flux:description>
    @enderror

    <div class="flex items-center justify-between">
        <flux:field class="mb-0">
            <flux:checkbox wire:model="remember" />
            <flux:label>Remember me</flux:label>
        </flux:field>
        <a href="/password/forgot" class="text-sm text-accent hover:underline">
            Forgot password?
        </a>
    </div>

    <flux:button type="submit" variant="primary" class="w-full">
        Sign In
    </flux:button>

    <div class="text-center text-sm">
        Don't have an account?
        <a href="/register" class="text-accent hover:underline">Sign up</a>
    </div>
</form>
```

### Settings Form

```blade
<form wire:submit="updateSettings" class="max-w-2xl space-y-8">
    <!-- Personal Settings -->
    <div>
        <flux:heading level="2">Personal Information</flux:heading>
        <div class="mt-4 space-y-6">
            <flux:input
                type="text"
                wire:model="fullName"
                label="Full Name"
                icon="user"
            />

            <flux:input
                type="email"
                wire:model="email"
                label="Email Address"
                icon="envelope"
            />

            <flux:select wire:model="timezone" label="Timezone">
                <option value="UTC">UTC</option>
                <option value="GMT-5">Eastern</option>
                <option value="GMT-6">Central</option>
            </flux:select>
        </div>
    </div>

    <!-- Notification Settings -->
    <div class="border-t pt-8">
        <flux:heading level="2">Notifications</flux:heading>
        <div class="mt-4 space-y-4">
            <div class="flex items-center justify-between">
                <div>
                    <flux:label class="mb-0">Email Notifications</flux:label>
                    <flux:description>Receive updates via email</flux:description>
                </div>
                <flux:switch wire:model="emailNotifications" />
            </div>

            <div class="flex items-center justify-between">
                <div>
                    <flux:label class="mb-0">Marketing Emails</flux:label>
                    <flux:description>Promotional offers and updates</flux:description>
                </div>
                <flux:switch wire:model="marketingEmails" />
            </div>
        </div>
    </div>

    <!-- Privacy Settings -->
    <div class="border-t pt-8">
        <flux:heading level="2">Privacy</flux:heading>
        <div class="mt-4 space-y-6">
            <flux:field>
                <flux:label>Profile Visibility</flux:label>
                <flux:radio.group wire:model="profileVisibility">
                    <flux:radio value="public">Public</flux:radio>
                    <flux:radio value="private">Private</flux:radio>
                </flux:radio.group>
            </flux:field>
        </div>
    </div>

    <div class="border-t pt-8 flex gap-3">
        <flux:button type="submit" variant="primary" icon="check">
            Save Changes
        </flux:button>
        <flux:button type="reset" variant="ghost">Reset</flux:button>
    </div>
</form>
```

---

## Form Field Props Patterns

All input components support these common props:

| Category | Props |
|----------|-------|
| **Data** | `wire:model`, `value`, `name` |
| **Labels** | `label`, `description`, `description:trailing` |
| **Display** | `placeholder`, `icon`, `icon:trailing` |
| **State** | `disabled`, `readonly`, `invalid` |
| **Styling** | `size`, `variant`, `class` |
| **Behaviour** | `required`, `autofocus`, `autocomplete` |

---

## Validation & Error Handling

### Laravel Validation

```php
// In your Livewire component
#[Validate('email|required|unique:users')]
public string $email = '';

#[Validate('min:8|regex:/[A-Z]/')]
public string $password = '';
```

### Displaying Errors

```blade
<flux:input
    wire:model="email"
    label="Email"
    :invalid="$errors->has('email')"
/>

@error('email')
    <flux:description class="text-red-600 text-sm">
        {{ $message }}
    </flux:description>
@enderror
```

### Multiple Error Messages

```blade
<flux:field>
    <flux:label>Password</flux:label>
    <flux:input
        type="password"
        wire:model="password"
        :invalid="$errors->has('password')"
    />

    @if ($errors->has('password'))
        <div class="space-y-1 mt-2">
            @foreach ($errors->get('password') as $error)
                <flux:description class="text-red-600 text-sm">
                    • {{ $error }}
                </flux:description>
            @endforeach
        </div>
    @endif
</flux:field>
```

---

Last updated: January 2026
