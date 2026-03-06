// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	"cmp"
	"embed"
	"html/template"
	"io"
	"slices"
	"strings"
)

//go:embed templates/*.html
var templateFS embed.FS

// templateFuncs returns the function map for help templates.
func templateFuncs() template.FuncMap {
	return template.FuncMap{
		"renderMarkdown": func(content string) template.HTML {
			html, err := RenderMarkdown(content)
			if err != nil {
				return template.HTML("<p>Error rendering content.</p>")
			}
			return template.HTML(html) //nolint:gosec // trusted content from catalog
		},
		"truncate": func(s string, n int) string {
			// Strip markdown headings for truncation preview
			lines := strings.SplitSeq(s, "\n")
			var clean []string
			for line := range lines {
				trimmed := strings.TrimSpace(line)
				if trimmed == "" || strings.HasPrefix(trimmed, "#") {
					continue
				}
				clean = append(clean, trimmed)
			}
			text := strings.Join(clean, " ")
			runes := []rune(text)
			if len(runes) <= n {
				return text
			}
			return string(runes[:n]) + "..."
		},
		"pluralise": func(count int, singular, plural string) string {
			if count == 1 {
				return singular
			}
			return plural
		},
		"multiply": func(a, b int) int {
			return a * b
		},
		"sub": func(a, b int) int {
			return a - b
		},
	}
}

// parseTemplates parses the base layout together with a page template.
func parseTemplates(page string) (*template.Template, error) {
	return template.New("base.html").Funcs(templateFuncs()).ParseFS(
		templateFS, "templates/base.html", "templates/"+page,
	)
}

// topicGroup groups topics under a tag for the index page.
type topicGroup struct {
	Tag    string
	Topics []*Topic
}

// indexData holds template data for the index page.
type indexData struct {
	Topics []*Topic
	Groups []topicGroup
}

// topicData holds template data for a single topic page.
type topicData struct {
	Topic *Topic
}

// searchData holds template data for the search results page.
type searchData struct {
	Query   string
	Results []*SearchResult
}

// groupTopicsByTag groups topics by their first tag.
// Topics without tags are grouped under "other".
// Groups are sorted alphabetically by tag name.
func groupTopicsByTag(topics []*Topic) []topicGroup {
	// Sort topics first by Order then Title
	sorted := slices.Clone(topics)
	slices.SortFunc(sorted, func(a, b *Topic) int {
		if a.Order != b.Order {
			return cmp.Compare(a.Order, b.Order)
		}
		return cmp.Compare(a.Title, b.Title)
	})

	groups := make(map[string][]*Topic)
	for _, t := range sorted {
		tag := "other"
		if len(t.Tags) > 0 {
			tag = t.Tags[0]
		}
		groups[tag] = append(groups[tag], t)
	}

	var result []topicGroup
	for tag, topics := range groups {
		result = append(result, topicGroup{Tag: tag, Topics: topics})
	}
	slices.SortFunc(result, func(a, b topicGroup) int {
		return cmp.Compare(a.Tag, b.Tag)
	})

	return result
}

// renderPage renders a named page template into the writer.
func renderPage(w io.Writer, page string, data any) error {
	tmpl, err := parseTemplates(page)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}
