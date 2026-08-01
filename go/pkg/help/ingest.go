// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	core "dappco.re/go"
)

// cutPrefix is equivalent of strings.CutPrefix without importing strings;
// no core wrapper exists in dappco.re/go.
func cutPrefix(s, prefix string) (after string, found bool) {
	if !core.HasPrefix(s, prefix) {
		return s, false
	}
	return s[len(prefix):], true
}

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
	parts := fields(name)
	if len(parts) > 0 {
		first := core.Lower(parts[0])
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
	words := fields(name)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = core.Upper(word[:1]) + word[1:]
		}
	}
	return core.Join(" ", words...)
}

// extractSeeAlso extracts related topic references from "See also:" lines.
// Returns the extracted references and the text with the "See also" line removed.
func extractSeeAlso(text string) ([]string, string) {
	var related []string
	var cleanedLines []string

	lines := core.Split(text, "\n")
	for _, line := range lines {
		trimmed := core.Trim(line)
		lower := core.Lower(trimmed)

		if after, ok := cutPrefix(lower, "see also:"); ok {
			// Extract references after "See also:"
			rest := after
			// The original casing version
			restOrig := trimmed[len("See also:"):]
			if len(restOrig) == 0 {
				restOrig = rest
			}

			refs := core.Split(restOrig, ",")
			for _, ref := range refs {
				ref = core.Trim(ref)
				if ref != "" {
					related = append(related, GenerateID(ref))
				}
			}
			continue
		}
		cleanedLines = append(cleanedLines, line)
	}

	return related, core.Join("\n", cleanedLines...)
}

// convertHelpToMarkdown converts structured CLI help text to Markdown.
// It identifies common sections (Usage, Flags, Options, Examples) and
// wraps them in appropriate Markdown headings and code blocks.
func convertHelpToMarkdown(text string) string {
	if core.Trim(text) == "" {
		return ""
	}

	lines := core.Split(text, "\n")
	var result core.Builder
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
		trimmed := core.Trim(line)

		// Detect section headers
		switch {
		case isSectionHeader(trimmed, "Usage:"):
			flushDesc()
			closeCodeBlock()
			result.WriteString("## Usage\n\n```\n")
			inCodeBlock = true
			// Include the rest of the line after "Usage:" if present
			rest := core.Trim(core.TrimPrefix(trimmed, "Usage:"))
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
				nextTrimmed := core.Trim(nextLine)
				if nextTrimmed == "" {
					i++
					continue
				}
				// Indented lines are subcommands
				if core.HasPrefix(nextLine, "  ") || core.HasPrefix(nextLine, "\t") {
					parts := fields(nextTrimmed)
					if len(parts) >= 2 {
						cmd := parts[0]
						desc := core.Join(" ", parts[1:]...)
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
				if trimmed != "" && !core.HasPrefix(line, " ") && !core.HasPrefix(line, "\t") && endsCodeBlockSection(trimmed) {
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

	return core.Trim(result.String())
}

// isSectionHeader checks if a line starts with the given section prefix.
func isSectionHeader(line, prefix string) bool {
	return core.HasPrefix(line, prefix) || core.HasPrefix(core.Lower(line), core.Lower(prefix))
}

// endsCodeBlockSection detects if a line indicates a new section that should
// end an open code block.
func endsCodeBlockSection(trimmed string) bool {
	lower := core.Lower(trimmed)
	prefixes := []string{"usage:", "flags:", "options:", "examples:", "commands:", "available commands:", "see also:"}
	for _, p := range prefixes {
		if core.HasPrefix(lower, p) {
			return true
		}
	}
	return false
}
