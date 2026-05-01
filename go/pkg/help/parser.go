package help

import (
	"iter"
	"regexp"
	"slices"
	"unicode"

	core "dappco.re/go"
	"gopkg.in/yaml.v3"
)

// repeat returns a string consisting of n copies of s. Equivalent of
// strings.Repeat without importing strings; no core wrapper exists in
// dappco.re/go.
func repeat(s string, n int) string {
	if n <= 0 {
		return ""
	}
	out := make([]byte, 0, len(s)*n)
	for i := 0; i < n; i++ {
		out = append(out, s...)
	}
	return string(out)
}

// fields splits s around runs of unicode.IsSpace whitespace, returning a
// slice of non-empty substrings. Equivalent of strings.Fields without
// importing strings; no core wrapper exists in dappco.re/go.
func fields(s string) []string {
	var out []string
	start := -1
	for i, r := range s {
		if unicode.IsSpace(r) {
			if start >= 0 {
				out = append(out, s[start:i])
				start = -1
			}
		} else if start < 0 {
			start = i
		}
	}
	if start >= 0 {
		out = append(out, s[start:])
	}
	return out
}

var (
	// frontmatterRegex matches YAML frontmatter delimited by ---
	// Supports both LF and CRLF line endings, and empty frontmatter blocks
	frontmatterRegex = regexp.MustCompile(`(?s)^---\r?\n(.*?)(?:\r?\n)?---\r?\n?`)

	// headingRegex matches markdown headings (# to ######)
	headingRegex = regexp.MustCompile(`^(#{1,6})\s+(.+)$`)
)

// ParseTopic parses a markdown file into a Topic.
func ParseTopic(path string, content []byte) (*Topic, error) {
	contentStr := string(content)

	topic := &Topic{
		Path:     path,
		ID:       GenerateID(pathToTitle(path)),
		Sections: []Section{},
		Tags:     []string{},
		Related:  []string{},
	}

	// Extract YAML frontmatter if present
	fm, body := ExtractFrontmatter(contentStr)
	if fm != nil {
		topic.Title = fm.Title
		topic.Tags = fm.Tags
		topic.Related = fm.Related
		topic.Order = fm.Order
		if topic.Title != "" {
			topic.ID = GenerateID(topic.Title)
		}
	}

	topic.Content = body

	// Extract sections from headings
	topic.Sections = ExtractSections(body)

	// If no title from frontmatter, try first H1
	if topic.Title == "" && len(topic.Sections) > 0 {
		for _, s := range topic.Sections {
			if s.Level == 1 {
				topic.Title = s.Title
				topic.ID = GenerateID(s.Title)
				break
			}
		}
	}

	return topic, nil
}

// ExtractFrontmatter extracts YAML frontmatter from markdown content.
// Returns the parsed frontmatter and the remaining content.
func ExtractFrontmatter(content string) (*Frontmatter, string) {
	match := frontmatterRegex.FindStringSubmatch(content)
	if match == nil {
		return nil, content
	}

	var fm Frontmatter
	if err := yaml.Unmarshal([]byte(match[1]), &fm); err != nil {
		// Invalid YAML, return content as-is
		return nil, content
	}

	// Return content without frontmatter
	body := content[len(match[0]):]
	return &fm, body
}

// ExtractSections parses markdown and returns sections.
func ExtractSections(content string) []Section {
	return slices.Collect(AllSections(content))
}

// AllSections returns an iterator for markdown sections.
func AllSections(content string) iter.Seq[Section] {
	return func(yield func(Section) bool) {
		lines := core.Split(content, "\n")

		var currentSection *Section
		var contentLines []string

		i := 0
		for _, line := range lines {
			lineNum := i + 1 // 1-indexed

			match := headingRegex.FindStringSubmatch(line)
			if match != nil {
				// Save previous section's content
				if currentSection != nil {
					currentSection.Content = core.Trim(core.Join("\n", contentLines...))
					if !yield(*currentSection) {
						return
					}
				}

				// Start new section
				level := len(match[1])
				title := core.Trim(match[2])

				section := Section{
					ID:    GenerateID(title),
					Title: title,
					Level: level,
					Line:  lineNum,
				}
				// We need to keep a pointer to the current section to update its content
				currentSection = &section
				contentLines = []string{}
			} else if currentSection != nil {
				contentLines = append(contentLines, line)
			}
			i++
		}

		// Save last section's content
		if currentSection != nil {
			currentSection.Content = core.Trim(core.Join("\n", contentLines...))
			yield(*currentSection)
		}
	}
}

// GenerateID creates a URL-safe ID from a title.
// "Getting Started" -> "getting-started"
func GenerateID(title string) string {
	var result core.Builder

	for _, r := range core.Lower(title) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result.WriteRune(r)
		} else if unicode.IsSpace(r) || r == '-' || r == '_' {
			// Only add hyphen if last char isn't already a hyphen
			str := result.String()
			if len(str) > 0 && str[len(str)-1] != '-' {
				result.WriteRune('-')
			}
		}
		// Skip other characters
	}

	// Trim trailing hyphens
	str := result.String()
	return core.TrimCutset(str, "-")
}

// pathToTitle converts a file path to a title.
// "getting-started.md" -> "Getting Started"
func pathToTitle(path string) string {
	// Get filename without directory (cross-platform)
	filename := core.PathBase(path)

	// Remove extension
	if ext := core.PathExt(filename); ext != "" {
		filename = core.TrimSuffix(filename, ext)
	}

	// Replace hyphens/underscores with spaces
	filename = core.Replace(filename, "-", " ")
	filename = core.Replace(filename, "_", " ")

	// Title case
	words := fields(filename)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = core.Upper(string(word[0])) + core.Lower(word[1:])
		}
	}

	return core.Join(" ", words...)
}
