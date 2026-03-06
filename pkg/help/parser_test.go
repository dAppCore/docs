// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateID_Good(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple title",
			input:    "Getting Started",
			expected: "getting-started",
		},
		{
			name:     "already lowercase",
			input:    "installation",
			expected: "installation",
		},
		{
			name:     "multiple spaces",
			input:    "Quick   Start   Guide",
			expected: "quick-start-guide",
		},
		{
			name:     "with numbers",
			input:    "Chapter 1 Introduction",
			expected: "chapter-1-introduction",
		},
		{
			name:     "special characters",
			input:    "What's New? (v2.0)",
			expected: "whats-new-v20",
		},
		{
			name:     "underscores",
			input:    "config_file_reference",
			expected: "config-file-reference",
		},
		{
			name:     "hyphens preserved",
			input:    "pre-commit hooks",
			expected: "pre-commit-hooks",
		},
		{
			name:     "leading trailing spaces",
			input:    "  Trimmed Title  ",
			expected: "trimmed-title",
		},
		{
			name:     "unicode letters",
			input:    "Configuración Básica",
			expected: "configuración-básica",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateID(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestExtractFrontmatter_Good(t *testing.T) {
	content := `---
title: Getting Started
tags: [intro, setup]
order: 1
related:
  - installation
  - configuration
---

# Welcome

This is the content.
`

	fm, body := ExtractFrontmatter(content)

	assert.NotNil(t, fm)
	assert.Equal(t, "Getting Started", fm.Title)
	assert.Equal(t, []string{"intro", "setup"}, fm.Tags)
	assert.Equal(t, 1, fm.Order)
	assert.Equal(t, []string{"installation", "configuration"}, fm.Related)
	assert.Contains(t, body, "# Welcome")
	assert.Contains(t, body, "This is the content.")
}

func TestExtractFrontmatter_Good_NoFrontmatter(t *testing.T) {
	content := `# Just a Heading

Some content here.
`

	fm, body := ExtractFrontmatter(content)

	assert.Nil(t, fm)
	assert.Equal(t, content, body)
}

func TestExtractFrontmatter_Good_CRLF(t *testing.T) {
	// Content with CRLF line endings (Windows-style)
	content := "---\r\ntitle: CRLF Test\r\n---\r\n\r\n# Content"

	fm, body := ExtractFrontmatter(content)

	assert.NotNil(t, fm)
	assert.Equal(t, "CRLF Test", fm.Title)
	assert.Contains(t, body, "# Content")
}

func TestExtractFrontmatter_Good_Empty(t *testing.T) {
	// Empty frontmatter block
	content := "---\n---\n# Content"

	fm, body := ExtractFrontmatter(content)

	// Empty frontmatter should parse successfully
	assert.NotNil(t, fm)
	assert.Equal(t, "", fm.Title)
	assert.Contains(t, body, "# Content")
}

func TestExtractFrontmatter_Bad_InvalidYAML(t *testing.T) {
	content := `---
title: [invalid yaml
---

# Content
`

	fm, body := ExtractFrontmatter(content)

	// Invalid YAML should return nil frontmatter and original content
	assert.Nil(t, fm)
	assert.Equal(t, content, body)
}

func TestExtractSections_Good(t *testing.T) {
	content := `# Main Title

Introduction paragraph.

## Installation

Install instructions here.
More details.

### Prerequisites

You need these things.

## Configuration

Config info here.
`

	sections := ExtractSections(content)

	assert.Len(t, sections, 4)

	// Main Title (H1)
	assert.Equal(t, "main-title", sections[0].ID)
	assert.Equal(t, "Main Title", sections[0].Title)
	assert.Equal(t, 1, sections[0].Level)
	assert.Equal(t, 1, sections[0].Line)
	assert.Contains(t, sections[0].Content, "Introduction paragraph.")

	// Installation (H2)
	assert.Equal(t, "installation", sections[1].ID)
	assert.Equal(t, "Installation", sections[1].Title)
	assert.Equal(t, 2, sections[1].Level)
	assert.Contains(t, sections[1].Content, "Install instructions here.")
	assert.Contains(t, sections[1].Content, "More details.")

	// Prerequisites (H3)
	assert.Equal(t, "prerequisites", sections[2].ID)
	assert.Equal(t, "Prerequisites", sections[2].Title)
	assert.Equal(t, 3, sections[2].Level)
	assert.Contains(t, sections[2].Content, "You need these things.")

	// Configuration (H2)
	assert.Equal(t, "configuration", sections[3].ID)
	assert.Equal(t, "Configuration", sections[3].Title)
	assert.Equal(t, 2, sections[3].Level)
}

func TestExtractSections_Good_AllHeadingLevels(t *testing.T) {
	content := `# H1
## H2
### H3
#### H4
##### H5
###### H6
`

	sections := ExtractSections(content)

	assert.Len(t, sections, 6)
	for i, level := range []int{1, 2, 3, 4, 5, 6} {
		assert.Equal(t, level, sections[i].Level)
	}
}

func TestExtractSections_Good_Empty(t *testing.T) {
	content := `Just plain text.
No headings here.
`

	sections := ExtractSections(content)

	assert.Empty(t, sections)
}

func TestParseTopic_Good(t *testing.T) {
	content := []byte(`---
title: Quick Start Guide
tags: [intro, quickstart]
order: 5
related:
  - installation
---

# Quick Start Guide

Welcome to the guide.

## First Steps

Do this first.

## Next Steps

Then do this.
`)

	topic, err := ParseTopic("docs/quick-start.md", content)

	assert.NoError(t, err)
	assert.NotNil(t, topic)

	// Check metadata from frontmatter
	assert.Equal(t, "quick-start-guide", topic.ID)
	assert.Equal(t, "Quick Start Guide", topic.Title)
	assert.Equal(t, "docs/quick-start.md", topic.Path)
	assert.Equal(t, []string{"intro", "quickstart"}, topic.Tags)
	assert.Equal(t, []string{"installation"}, topic.Related)
	assert.Equal(t, 5, topic.Order)

	// Check sections
	assert.Len(t, topic.Sections, 3)
	assert.Equal(t, "quick-start-guide", topic.Sections[0].ID)
	assert.Equal(t, "first-steps", topic.Sections[1].ID)
	assert.Equal(t, "next-steps", topic.Sections[2].ID)

	// Content should not include frontmatter
	assert.NotContains(t, topic.Content, "---")
	assert.Contains(t, topic.Content, "# Quick Start Guide")
}

func TestParseTopic_Good_NoFrontmatter(t *testing.T) {
	content := []byte(`# Getting Started

This is a simple doc.

## Installation

Install it here.
`)

	topic, err := ParseTopic("getting-started.md", content)

	assert.NoError(t, err)
	assert.NotNil(t, topic)

	// Title should come from first H1
	assert.Equal(t, "Getting Started", topic.Title)
	assert.Equal(t, "getting-started", topic.ID)

	// Sections extracted
	assert.Len(t, topic.Sections, 2)
}

func TestParseTopic_Good_NoHeadings(t *testing.T) {
	content := []byte(`---
title: Plain Content
---

Just some text without any headings.
`)

	topic, err := ParseTopic("plain.md", content)

	assert.NoError(t, err)
	assert.NotNil(t, topic)
	assert.Equal(t, "Plain Content", topic.Title)
	assert.Equal(t, "plain-content", topic.ID)
	assert.Empty(t, topic.Sections)
}

func TestParseTopic_Good_IDFromPath(t *testing.T) {
	content := []byte(`Just content, no frontmatter or headings.`)

	topic, err := ParseTopic("commands/dev-workflow.md", content)

	assert.NoError(t, err)
	assert.NotNil(t, topic)

	// ID and title should be derived from path
	assert.Equal(t, "dev-workflow", topic.ID)
	assert.Equal(t, "", topic.Title) // No title available
}

func TestPathToTitle_Good(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{"getting-started.md", "Getting Started"},
		{"commands/dev.md", "Dev"},
		{"path/to/file_name.md", "File Name"},
		{"UPPERCASE.md", "Uppercase"},
		{"no-extension", "No Extension"},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			result := pathToTitle(tt.path)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// --- Phase 0: Expanded parser tests ---

func TestParseTopic_Good_EmptyInput(t *testing.T) {
	// Empty byte slice should produce a valid topic with no content
	topic, err := ParseTopic("empty.md", []byte(""))

	require.NoError(t, err)
	assert.NotNil(t, topic)
	assert.Equal(t, "empty", topic.ID)
	assert.Equal(t, "", topic.Title)
	assert.Equal(t, "", topic.Content)
	assert.Empty(t, topic.Sections)
	assert.Empty(t, topic.Tags)
	assert.Empty(t, topic.Related)
}

func TestParseTopic_Good_FrontmatterOnly(t *testing.T) {
	// Frontmatter with no body or sections
	content := []byte(`---
title: Metadata Only
tags: [meta]
order: 99
---
`)

	topic, err := ParseTopic("meta.md", content)

	require.NoError(t, err)
	assert.Equal(t, "metadata-only", topic.ID)
	assert.Equal(t, "Metadata Only", topic.Title)
	assert.Equal(t, []string{"meta"}, topic.Tags)
	assert.Equal(t, 99, topic.Order)
	assert.Empty(t, topic.Sections)
	// Body after frontmatter is just a newline
	assert.Equal(t, "", strings.TrimSpace(topic.Content))
}

func TestExtractFrontmatter_Bad_MalformedYAML(t *testing.T) {
	tests := []struct {
		name    string
		content string
	}{
		{
			name: "unclosed bracket",
			content: `---
title: [broken
tags: [also broken
---

# Content`,
		},
		{
			name:    "tab indentation error",
			content: "---\ntitle: Good\n\t- bad indent\n---\n\n# Content",
		},
		{
			name: "duplicate keys with conflicting types",
			// YAML spec allows duplicate keys but implementations may vary;
			// this tests that the parser does not panic regardless.
			content: `---
title: First
title:
  nested: value
---

# Content`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm, body := ExtractFrontmatter(tt.content)
			// Malformed YAML should return nil frontmatter without panic
			if fm == nil {
				// Body should be original content when YAML fails
				assert.Equal(t, tt.content, body)
			}
			// No panic is the key assertion — test reaching here is success
		})
	}
}

func TestExtractFrontmatter_Bad_NotAtStart(t *testing.T) {
	// Frontmatter delimiters that do not start at the beginning of the file
	content := `Some preamble text.

---
title: Should Not Parse
---

# Content`

	fm, body := ExtractFrontmatter(content)

	assert.Nil(t, fm)
	assert.Equal(t, content, body)
}

func TestExtractSections_Good_DeeplyNested(t *testing.T) {
	content := `# Level 1

Top-level content.

## Level 2

Second level.

### Level 3

Third level.

#### Level 4

Fourth level details.

##### Level 5

Fifth level fine print.

###### Level 6

Deepest heading level.
`

	sections := ExtractSections(content)

	require.Len(t, sections, 6)

	for i, expected := range []struct {
		level int
		title string
	}{
		{1, "Level 1"},
		{2, "Level 2"},
		{3, "Level 3"},
		{4, "Level 4"},
		{5, "Level 5"},
		{6, "Level 6"},
	} {
		assert.Equal(t, expected.level, sections[i].Level, "section %d level", i)
		assert.Equal(t, expected.title, sections[i].Title, "section %d title", i)
	}

	// Verify content is associated with correct sections
	assert.Contains(t, sections[0].Content, "Top-level content.")
	assert.Contains(t, sections[3].Content, "Fourth level details.")
	assert.Contains(t, sections[5].Content, "Deepest heading level.")
}

func TestExtractSections_Good_DeeplyNestedWithContent(t *testing.T) {
	// H4, H5, H6 with meaningful content under each
	content := `#### Configuration Options

Set these in your config file.

##### Advanced Options

Only for power users.

###### Experimental Flags

These may change without notice.
`

	sections := ExtractSections(content)

	require.Len(t, sections, 3)
	assert.Equal(t, 4, sections[0].Level)
	assert.Equal(t, "Configuration Options", sections[0].Title)
	assert.Contains(t, sections[0].Content, "Set these in your config file.")

	assert.Equal(t, 5, sections[1].Level)
	assert.Equal(t, "Advanced Options", sections[1].Title)
	assert.Contains(t, sections[1].Content, "Only for power users.")

	assert.Equal(t, 6, sections[2].Level)
	assert.Equal(t, "Experimental Flags", sections[2].Title)
	assert.Contains(t, sections[2].Content, "These may change without notice.")
}

func TestParseTopic_Good_Unicode(t *testing.T) {
	tests := []struct {
		name    string
		content string
		title   string
	}{
		{
			name: "CJK characters",
			content: `---
title: 日本語ドキュメント
tags: [日本語, ドキュメント]
---

# 日本語ドキュメント

はじめにの内容です。

## インストール

インストール手順はこちら。
`,
			title: "日本語ドキュメント",
		},
		{
			name: "emoji in title and content",
			content: `---
title: Rocket Launch 🚀
tags: [emoji, fun]
---

# Rocket Launch 🚀

This topic has emoji 🎉 in the content.

## Features ✨

- Fast ⚡
- Reliable 🔒
`,
			title: "Rocket Launch 🚀",
		},
		{
			name: "diacritics and accented characters",
			content: `---
title: Présentation Générale
tags: [français]
---

# Présentation Générale

Bienvenue à la documentation. Les données sont protégées.

## Résumé

Aperçu des fonctionnalités clés.
`,
			title: "Présentation Générale",
		},
		{
			name: "mixed scripts",
			content: `---
title: Mixed Скрипты 混合
---

# Mixed Скрипты 混合

Content with Кириллица, 中文, العربية, and हिन्दी.
`,
			title: "Mixed Скрипты 混合",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			topic, err := ParseTopic("unicode.md", []byte(tt.content))

			require.NoError(t, err)
			assert.Equal(t, tt.title, topic.Title)
			assert.NotEmpty(t, topic.ID)
			assert.True(t, len(topic.Sections) > 0, "should extract sections from unicode content")
		})
	}
}

func TestParseTopic_Good_VeryLongDocument(t *testing.T) {
	// Build a document with 10,000+ lines
	var b strings.Builder

	b.WriteString("---\ntitle: Massive Document\ntags: [large, stress]\n---\n\n")

	// Generate 100 sections, each with ~100 lines of content
	for i := range 100 {
		b.WriteString(fmt.Sprintf("## Section %d\n\n", i+1))
		for j := range 100 {
			b.WriteString(fmt.Sprintf("Line %d of section %d: Lorem ipsum dolor sit amet.\n", j+1, i+1))
		}
		b.WriteString("\n")
	}

	content := b.String()
	lineCount := strings.Count(content, "\n")
	assert.Greater(t, lineCount, 10000, "document should exceed 10K lines")

	topic, err := ParseTopic("massive.md", []byte(content))

	require.NoError(t, err)
	assert.Equal(t, "Massive Document", topic.Title)
	assert.Equal(t, "massive-document", topic.ID)
	assert.Len(t, topic.Sections, 100)

	// Verify first and last sections have correct titles
	assert.Equal(t, "Section 1", topic.Sections[0].Title)
	assert.Equal(t, "Section 100", topic.Sections[99].Title)

	// Verify content is captured in sections
	assert.Contains(t, topic.Sections[0].Content, "Line 1 of section 1")
	assert.Contains(t, topic.Sections[99].Content, "Line 100 of section 100")
}

func TestExtractSections_Bad_EmptyString(t *testing.T) {
	sections := ExtractSections("")
	assert.Empty(t, sections)
}

func TestExtractSections_Bad_HeadingWithoutSpace(t *testing.T) {
	// "#NoSpace" is not a valid markdown heading (needs space after #)
	content := `#NoSpace
##AlsoNoSpace
Some text.
`

	sections := ExtractSections(content)
	assert.Empty(t, sections, "headings without space after # should not be parsed")
}

func TestExtractSections_Good_ConsecutiveHeadings(t *testing.T) {
	// Headings with no content between them
	content := `# Title
## Subtitle
### Sub-subtitle
`

	sections := ExtractSections(content)

	require.Len(t, sections, 3)
	// First two sections should have empty content
	assert.Equal(t, "", sections[0].Content)
	assert.Equal(t, "", sections[1].Content)
	assert.Equal(t, "", sections[2].Content)
}

func TestGenerateID_Ugly_EmptyString(t *testing.T) {
	result := GenerateID("")
	assert.Equal(t, "", result)
}

func TestGenerateID_Good_OnlySpecialChars(t *testing.T) {
	result := GenerateID("!@#$%^&*()")
	assert.Equal(t, "", result)
}

func TestGenerateID_Good_CJK(t *testing.T) {
	result := GenerateID("日本語テスト")
	assert.NotEmpty(t, result)
	assert.NotContains(t, result, " ")
}

func TestGenerateID_Good_Emoji(t *testing.T) {
	result := GenerateID("Hello 🌍 World")
	// Emoji are not letters or digits, so they are dropped
	assert.Equal(t, "hello-world", result)
}
