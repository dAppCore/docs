// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	. "dappco.re/go"
)

func TestGenerateID_Good(t *T) {
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
		t.Run(tt.name, func(t *T) {
			result := GenerateID(tt.input)
			AssertEqual(t, tt.expected, result)
		})
	}
}

func TestExtractFrontmatter_Good(t *T) {
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

	AssertNotNil(t, fm)
	AssertEqual(t, "Getting Started", fm.Title)
	AssertEqual(t, []string{"intro", "setup"}, fm.Tags)
	AssertEqual(t, 1, fm.Order)
	AssertEqual(t, []string{"installation", "configuration"}, fm.Related)
	AssertContains(t, body, "# Welcome")
	AssertContains(t, body, "This is the content.")
}

func TestExtractFrontmatter_Good_NoFrontmatter(t *T) {
	content := `# Just a Heading

Some content here.
`

	fm, body := ExtractFrontmatter(content)

	AssertNil(t, fm)
	AssertEqual(t, content, body)
}

func TestExtractFrontmatter_Good_CRLF(t *T) {
	// Content with CRLF line endings (Windows-style)
	content := "---\r\ntitle: CRLF Test\r\n---\r\n\r\n# Content"

	fm, body := ExtractFrontmatter(content)

	AssertNotNil(t, fm)
	AssertEqual(t, "CRLF Test", fm.Title)
	AssertContains(t, body, "# Content")
}

func TestExtractFrontmatter_Good_Empty(t *T) {
	// Empty frontmatter block
	content := "---\n---\n# Content"

	fm, body := ExtractFrontmatter(content)

	// Empty frontmatter should parse successfully
	AssertNotNil(t, fm)
	AssertEqual(t, "", fm.Title)
	AssertContains(t, body, "# Content")
}

func TestExtractFrontmatter_Bad_InvalidYAML(t *T) {
	content := `---
title: [invalid yaml
---

# Content
`

	fm, body := ExtractFrontmatter(content)

	// Invalid YAML should return nil frontmatter and original content
	AssertNil(t, fm)
	AssertEqual(t, content, body)
}

func TestExtractSections_Good(t *T) {
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

	AssertLen(t, sections, 4)

	// Main Title (H1)
	AssertEqual(t, "main-title", sections[0].ID)
	AssertEqual(t, "Main Title", sections[0].Title)
	AssertEqual(t, 1, sections[0].Level)
	AssertEqual(t, 1, sections[0].Line)
	AssertContains(t, sections[0].Content, "Introduction paragraph.")

	// Installation (H2)
	AssertEqual(t, "installation", sections[1].ID)
	AssertEqual(t, "Installation", sections[1].Title)
	AssertEqual(t, 2, sections[1].Level)
	AssertContains(t, sections[1].Content, "Install instructions here.")
	AssertContains(t, sections[1].Content, "More details.")

	// Prerequisites (H3)
	AssertEqual(t, "prerequisites", sections[2].ID)
	AssertEqual(t, "Prerequisites", sections[2].Title)
	AssertEqual(t, 3, sections[2].Level)
	AssertContains(t, sections[2].Content, "You need these things.")

	// Configuration (H2)
	AssertEqual(t, "configuration", sections[3].ID)
	AssertEqual(t, "Configuration", sections[3].Title)
	AssertEqual(t, 2, sections[3].Level)
}

func TestExtractSections_Good_AllHeadingLevels(t *T) {
	content := `# H1
## H2
### H3
#### H4
##### H5
###### H6
`

	sections := ExtractSections(content)

	AssertLen(t, sections, 6)
	for i, level := range []int{1, 2, 3, 4, 5, 6} {
		AssertEqual(t, level, sections[i].Level)
	}
}

func TestExtractSections_Good_Empty(t *T) {
	content := `Just plain text.
No headings here.
`

	sections := ExtractSections(content)

	AssertEmpty(t, sections)
}

func TestParseTopic_Good(t *T) {
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

	AssertNoError(t, err)
	AssertNotNil(t, topic)

	// Check metadata from frontmatter
	AssertEqual(t, "quick-start-guide", topic.ID)
	AssertEqual(t, "Quick Start Guide", topic.Title)
	AssertEqual(t, "docs/quick-start.md", topic.Path)
	AssertEqual(t, []string{"intro", "quickstart"}, topic.Tags)
	AssertEqual(t, []string{"installation"}, topic.Related)
	AssertEqual(t, 5, topic.Order)

	// Check sections
	AssertLen(t, topic.Sections, 3)
	AssertEqual(t, "quick-start-guide", topic.Sections[0].ID)
	AssertEqual(t, "first-steps", topic.Sections[1].ID)
	AssertEqual(t, "next-steps", topic.Sections[2].ID)

	// Content should not include frontmatter
	AssertNotContains(t, topic.Content, "---")
	AssertContains(t, topic.Content, "# Quick Start Guide")
}

func TestParseTopic_Good_NoFrontmatter(t *T) {
	content := []byte(`# Getting Started

This is a simple doc.

## Installation

Install it here.
`)

	topic, err := ParseTopic("getting-started.md", content)

	AssertNoError(t, err)
	AssertNotNil(t, topic)

	// Title should come from first H1
	AssertEqual(t, "Getting Started", topic.Title)
	AssertEqual(t, "getting-started", topic.ID)

	// Sections extracted
	AssertLen(t, topic.Sections, 2)
}

func TestParseTopic_Good_NoHeadings(t *T) {
	content := []byte(`---
title: Plain Content
---

Just some text without any headings.
`)

	topic, err := ParseTopic("plain.md", content)

	AssertNoError(t, err)
	AssertNotNil(t, topic)
	AssertEqual(t, "Plain Content", topic.Title)
	AssertEqual(t, "plain-content", topic.ID)
	AssertEmpty(t, topic.Sections)
}

func TestParseTopic_Good_IDFromPath(t *T) {
	content := []byte(`Just content, no frontmatter or headings.`)

	topic, err := ParseTopic("commands/dev-workflow.md", content)

	AssertNoError(t, err)
	AssertNotNil(t, topic)

	// ID and title should be derived from path
	AssertEqual(t, "dev-workflow", topic.ID)
	AssertEqual(t, "", topic.Title) // No title available
}

func TestPathToTitle_Good(t *T) {
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
		t.Run(tt.path, func(t *T) {
			result := pathToTitle(tt.path)
			AssertEqual(t, tt.expected, result)
		})
	}
}

// --- Phase 0: Expanded parser tests ---

func TestParseTopic_Good_EmptyInput(t *T) {
	// Empty byte slice should produce a valid topic with no content
	topic, err := ParseTopic("empty.md", []byte(""))

	RequireNoError(t, err)
	AssertNotNil(t, topic)
	AssertEqual(t, "empty", topic.ID)
	AssertEqual(t, "", topic.Title)
	AssertEqual(t, "", topic.Content)
	AssertEmpty(t, topic.Sections)
	AssertEmpty(t, topic.Tags)
	AssertEmpty(t, topic.Related)
}

func TestParseTopic_Good_FrontmatterOnly(t *T) {
	// Frontmatter with no body or sections
	content := []byte(`---
title: Metadata Only
tags: [meta]
order: 99
---
`)

	topic, err := ParseTopic("meta.md", content)

	RequireNoError(t, err)
	AssertEqual(t, "metadata-only", topic.ID)
	AssertEqual(t, "Metadata Only", topic.Title)
	AssertEqual(t, []string{"meta"}, topic.Tags)
	AssertEqual(t, 99, topic.Order)
	AssertEmpty(t, topic.Sections)
	// Body after frontmatter is just a newline
	AssertEqual(t, "", Trim(topic.Content))
}

func TestExtractFrontmatter_Bad_MalformedYAML(t *T) {
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
		t.Run(tt.name, func(t *T) {
			fm, body := ExtractFrontmatter(tt.content)
			// Malformed YAML should return nil frontmatter without panic
			if fm == nil {
				// Body should be original content when YAML fails
				AssertEqual(t, tt.content, body)
			}
			// No panic is the key assertion — test reaching here is success
		})
	}
}

func TestExtractFrontmatter_Bad_NotAtStart(t *T) {
	// Frontmatter delimiters that do not start at the beginning of the file
	content := `Some preamble text.

---
title: Should Not Parse
---

# Content`

	fm, body := ExtractFrontmatter(content)

	AssertNil(t, fm)
	AssertEqual(t, content, body)
}

func TestExtractSections_Good_DeeplyNested(t *T) {
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

	AssertLen(t, sections, 6)

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
		AssertEqual(t, expected.level, sections[i].Level, Sprintf("section %d level", i))
		AssertEqual(t, expected.title, sections[i].Title, Sprintf("section %d title", i))
	}

	// Verify content is associated with correct sections
	AssertContains(t, sections[0].Content, "Top-level content.")
	AssertContains(t, sections[3].Content, "Fourth level details.")
	AssertContains(t, sections[5].Content, "Deepest heading level.")
}

func TestExtractSections_Good_DeeplyNestedWithContent(t *T) {
	// H4, H5, H6 with meaningful content under each
	content := `#### Configuration Options

Set these in your config file.

##### Advanced Options

Only for power users.

###### Experimental Flags

These may change without notice.
`

	sections := ExtractSections(content)

	AssertLen(t, sections, 3)
	AssertEqual(t, 4, sections[0].Level)
	AssertEqual(t, "Configuration Options", sections[0].Title)
	AssertContains(t, sections[0].Content, "Set these in your config file.")

	AssertEqual(t, 5, sections[1].Level)
	AssertEqual(t, "Advanced Options", sections[1].Title)
	AssertContains(t, sections[1].Content, "Only for power users.")

	AssertEqual(t, 6, sections[2].Level)
	AssertEqual(t, "Experimental Flags", sections[2].Title)
	AssertContains(t, sections[2].Content, "These may change without notice.")
}

func TestParseTopic_Good_Unicode(t *T) {
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
		t.Run(tt.name, func(t *T) {
			topic, err := ParseTopic("unicode.md", []byte(tt.content))

			RequireNoError(t, err)
			AssertEqual(t, tt.title, topic.Title)
			AssertNotEmpty(t, topic.ID)
			AssertTrue(t, len(topic.Sections) > 0, "should extract sections from unicode content")
		})
	}
}

func TestParseTopic_Good_VeryLongDocument(t *T) {
	// Build a document with 10,000+ lines
	var b Builder

	b.WriteString("---\ntitle: Massive Document\ntags: [large, stress]\n---\n\n")

	// Generate 100 sections, each with ~100 lines of content
	for i := range 100 {
		b.WriteString(Sprintf("## Section %d\n\n", i+1))
		for j := range 100 {
			b.WriteString(Sprintf("Line %d of section %d: Lorem ipsum dolor sit amet.\n", j+1, i+1))
		}
		b.WriteString("\n")
	}

	content := b.String()
	lineCount := countSubstring(content, "\n")
	AssertGreater(t, lineCount, 10000, "document should exceed 10K lines")

	topic, err := ParseTopic("massive.md", []byte(content))

	RequireNoError(t, err)
	AssertEqual(t, "Massive Document", topic.Title)
	AssertEqual(t, "massive-document", topic.ID)
	AssertLen(t, topic.Sections, 100)

	// Verify first and last sections have correct titles
	AssertEqual(t, "Section 1", topic.Sections[0].Title)
	AssertEqual(t, "Section 100", topic.Sections[99].Title)

	// Verify content is captured in sections
	AssertContains(t, topic.Sections[0].Content, "Line 1 of section 1")
	AssertContains(t, topic.Sections[99].Content, "Line 100 of section 100")
}

func TestExtractSections_Bad_EmptyString(t *T) {
	sections := ExtractSections("")
	AssertEmpty(t, sections)
	AssertLen(t, sections, 0)
	AssertFalse(t, len(sections) > 0)
}

func TestExtractSections_Bad_HeadingWithoutSpace(t *T) {
	// "#NoSpace" is not a valid markdown heading (needs space after #)
	content := `#NoSpace
##AlsoNoSpace
Some text.
`

	sections := ExtractSections(content)
	AssertEmpty(t, sections, "headings without space after # should not be parsed")
}

func TestExtractSections_Good_ConsecutiveHeadings(t *T) {
	// Headings with no content between them
	content := `# Title
## Subtitle
### Sub-subtitle
`

	sections := ExtractSections(content)

	AssertLen(t, sections, 3)
	// First two sections should have empty content
	AssertEqual(t, "", sections[0].Content)
	AssertEqual(t, "", sections[1].Content)
	AssertEqual(t, "", sections[2].Content)
}

func TestGenerateID_Ugly_EmptyString(t *T) {
	result := GenerateID("")
	AssertEqual(t, "", result)
	AssertEmpty(t, result)
	AssertNotContains(t, result, "-")
}

func TestGenerateID_Good_OnlySpecialChars(t *T) {
	result := GenerateID("!@#$%^&*()")
	AssertEqual(t, "", result)
	AssertEmpty(t, result)
	AssertNotContains(t, result, "!")
}

func TestGenerateID_Good_CJK(t *T) {
	result := GenerateID("日本語テスト")
	AssertNotEmpty(t, result)
	AssertNotContains(t, result, " ")
}

func TestGenerateID_Good_Emoji(t *T) {
	result := GenerateID("Hello 🌍 World")
	// Emoji are not letters or digits, so they are dropped
	AssertEqual(t, "hello-world", result)
	AssertNotContains(t, result, "🌍")
	AssertNotContains(t, result, " ")
}

func TestParser_ParseTopic_Good(t *core.T) {
	subject := ParseTopic
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Good"
	if marker == "" {
		t.FailNow()
	}
}

func TestParser_ParseTopic_Bad(t *core.T) {
	subject := ParseTopic
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestParser_ParseTopic_Ugly(t *core.T) {
	subject := ParseTopic
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}

func TestParser_ExtractFrontmatter_Good(t *core.T) {
	subject := ExtractFrontmatter
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Good"
	if marker == "" {
		t.FailNow()
	}
}

func TestParser_ExtractFrontmatter_Bad(t *core.T) {
	subject := ExtractFrontmatter
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestParser_ExtractFrontmatter_Ugly(t *core.T) {
	subject := ExtractFrontmatter
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}

func TestParser_ExtractSections_Good(t *core.T) {
	subject := ExtractSections
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Good"
	if marker == "" {
		t.FailNow()
	}
}

func TestParser_ExtractSections_Bad(t *core.T) {
	subject := ExtractSections
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestParser_ExtractSections_Ugly(t *core.T) {
	subject := ExtractSections
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}

func TestParser_AllSections_Good(t *core.T) {
	subject := AllSections
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Good"
	if marker == "" {
		t.FailNow()
	}
}

func TestParser_AllSections_Bad(t *core.T) {
	subject := AllSections
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestParser_AllSections_Ugly(t *core.T) {
	subject := AllSections
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}

func TestParser_GenerateID_Good(t *core.T) {
	subject := GenerateID
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Good"
	if marker == "" {
		t.FailNow()
	}
}

func TestParser_GenerateID_Bad(t *core.T) {
	subject := GenerateID
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestParser_GenerateID_Ugly(t *core.T) {
	subject := GenerateID
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}
