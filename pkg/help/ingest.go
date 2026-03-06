// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	"strings"
)

// ParseHelpText parses standard Go CLI help output into a Topic.
//
// The name parameter is the command name (e.g. "dev commit") and text
// is the raw help output. The function converts structured sections
// (Usage, Flags, Options, Examples) into Markdown, and extracts "See also"
// lines into the Related field.
func ParseHelpText(name string, text string) *Topic {
	title := titleCaseWords(name)
	id := GenerateID(name)

	// Extract related topics from "See also:" lines.
	related, cleanedText := extractSeeAlso(text)

	// Convert help text sections to Markdown.
	content := convertHelpToMarkdown(cleanedText)

	// Derive tags: "cli" + first word of the command name.
	tags := []string{"cli"}
	parts := strings.Fields(name)
	if len(parts) > 0 {
		first := strings.ToLower(parts[0])
		if first != "cli" { // avoid duplicate
			tags = append(tags, first)
		}
	}

	topic := &Topic{
		ID:       id,
		Title:    title,
		Content:  content,
		Tags:     tags,
		Related:  related,
		Sections: ExtractSections(content),
	}

	return topic
}

// IngestCLIHelp batch-ingests CLI help texts into a new Catalog.
// Keys are command names (e.g. "dev commit"), values are raw help output.
func IngestCLIHelp(helpTexts map[string]string) *Catalog {
	c := &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}

	for name, text := range helpTexts {
		topic := ParseHelpText(name, text)
		c.Add(topic)
	}

	return c
}

// titleCaseWords converts "dev commit" to "Dev Commit".
func titleCaseWords(name string) string {
	words := strings.Fields(name)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + word[1:]
		}
	}
	return strings.Join(words, " ")
}

// extractSeeAlso extracts related topic references from "See also:" lines.
// Returns the extracted references and the text with the "See also" line removed.
func extractSeeAlso(text string) ([]string, string) {
	var related []string
	var cleanedLines []string

	lines := strings.SplitSeq(text, "\n")
	for line := range lines {
		trimmed := strings.TrimSpace(line)
		lower := strings.ToLower(trimmed)

		if after, ok := strings.CutPrefix(lower, "see also:"); ok {
			// Extract references after "See also:"
			rest := after
			// The original casing version
			restOrig := trimmed[len("See also:"):]
			if len(restOrig) == 0 {
				restOrig = rest
			}

			refs := strings.SplitSeq(restOrig, ",")
			for ref := range refs {
				ref = strings.TrimSpace(ref)
				if ref != "" {
					related = append(related, GenerateID(ref))
				}
			}
			continue
		}
		cleanedLines = append(cleanedLines, line)
	}

	return related, strings.Join(cleanedLines, "\n")
}

// convertHelpToMarkdown converts structured CLI help text to Markdown.
// It identifies common sections (Usage, Flags, Options, Examples) and
// wraps them in appropriate Markdown headings and code blocks.
func convertHelpToMarkdown(text string) string {
	if strings.TrimSpace(text) == "" {
		return ""
	}

	lines := strings.Split(text, "\n")
	var result strings.Builder
	inCodeBlock := false
	var descLines []string

	flushDesc := func() {
		if len(descLines) > 0 {
			for _, dl := range descLines {
				result.WriteString(dl)
				result.WriteByte('\n')
			}
			descLines = nil
		}
	}

	closeCodeBlock := func() {
		if inCodeBlock {
			result.WriteString("```\n\n")
			inCodeBlock = false
		}
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		trimmed := strings.TrimSpace(line)

		// Detect section headers
		switch {
		case isSectionHeader(trimmed, "Usage:"):
			flushDesc()
			closeCodeBlock()
			result.WriteString("## Usage\n\n```\n")
			inCodeBlock = true
			// Include the rest of the line after "Usage:" if present
			rest := strings.TrimSpace(strings.TrimPrefix(trimmed, "Usage:"))
			if rest != "" {
				result.WriteString(rest)
				result.WriteByte('\n')
			}

		case isSectionHeader(trimmed, "Flags:") || isSectionHeader(trimmed, "Options:"):
			flushDesc()
			closeCodeBlock()
			result.WriteString("## Flags\n\n```\n")
			inCodeBlock = true

		case isSectionHeader(trimmed, "Examples:"):
			flushDesc()
			closeCodeBlock()
			result.WriteString("## Examples\n\n```\n")
			inCodeBlock = true

		case isSectionHeader(trimmed, "Commands:") || isSectionHeader(trimmed, "Available Commands:"):
			flushDesc()
			closeCodeBlock()
			result.WriteString("## Commands\n\n")
			// Parse subsequent indented lines as subcommand list
			for i+1 < len(lines) {
				nextLine := lines[i+1]
				nextTrimmed := strings.TrimSpace(nextLine)
				if nextTrimmed == "" {
					i++
					continue
				}
				// Indented lines are subcommands
				if strings.HasPrefix(nextLine, "  ") || strings.HasPrefix(nextLine, "\t") {
					parts := strings.Fields(nextTrimmed)
					if len(parts) >= 2 {
						cmd := parts[0]
						desc := strings.Join(parts[1:], " ")
						result.WriteString("- **")
						result.WriteString(cmd)
						result.WriteString("** -- ")
						result.WriteString(desc)
						result.WriteByte('\n')
					} else {
						result.WriteString("- ")
						result.WriteString(nextTrimmed)
						result.WriteByte('\n')
					}
					i++
				} else {
					break
				}
			}
			result.WriteByte('\n')

		default:
			if inCodeBlock {
				// Check if this line starts a new section (end code block)
				if trimmed != "" && !strings.HasPrefix(line, " ") && !strings.HasPrefix(line, "\t") && endsCodeBlockSection(trimmed) {
					closeCodeBlock()
					// Re-process this line
					i--
					continue
				}
				result.WriteString(line)
				result.WriteByte('\n')
			} else {
				// Descriptive paragraph text
				descLines = append(descLines, line)
			}
		}
	}

	flushDesc()
	closeCodeBlock()

	return strings.TrimSpace(result.String())
}

// isSectionHeader checks if a line starts with the given section prefix.
func isSectionHeader(line, prefix string) bool {
	return strings.HasPrefix(line, prefix) || strings.HasPrefix(strings.ToLower(line), strings.ToLower(prefix))
}

// endsCodeBlockSection detects if a line indicates a new section that should
// end an open code block.
func endsCodeBlockSection(trimmed string) bool {
	lower := strings.ToLower(trimmed)
	prefixes := []string{"usage:", "flags:", "options:", "examples:", "commands:", "available commands:", "see also:"}
	for _, p := range prefixes {
		if strings.HasPrefix(lower, p) {
			return true
		}
	}
	return false
}
