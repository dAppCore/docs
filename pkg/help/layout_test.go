// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRenderLayout_Good_IndexPage(t *testing.T) {
	topics := []*Topic{
		{ID: "getting-started", Title: "Getting Started", Tags: []string{"intro"}, Content: "Welcome to the guide."},
		{ID: "config", Title: "Configuration", Tags: []string{"setup"}, Content: "Config options here."},
		{ID: "advanced", Title: "Advanced Usage", Tags: []string{"intro"}, Content: "Power user tips."},
	}

	html := RenderIndexPage(topics)

	// Must contain ARIA roles from HLCRF layout
	assert.Contains(t, html, `role="banner"`, "header should have banner role")
	assert.Contains(t, html, `role="main"`, "content should have main role")
	assert.Contains(t, html, `role="contentinfo"`, "footer should have contentinfo role")

	// Must contain topic titles
	assert.Contains(t, html, "Getting Started")
	assert.Contains(t, html, "Configuration")
	assert.Contains(t, html, "Advanced Usage")

	// Must contain brand
	assert.Contains(t, html, "core.help")

	// Must contain tag group headings
	assert.Contains(t, html, "intro")
	assert.Contains(t, html, "setup")
}

func TestRenderLayout_Good_TopicPage(t *testing.T) {
	topic := &Topic{
		ID:      "getting-started",
		Title:   "Getting Started",
		Content: "# Getting Started\n\nWelcome to the **guide**.\n",
		Tags:    []string{"intro"},
		Sections: []Section{
			{ID: "overview", Title: "Overview", Level: 2},
			{ID: "installation", Title: "Installation", Level: 2},
			{ID: "quick-start", Title: "Quick Start", Level: 3},
		},
		Related: []string{"config"},
	}

	sidebar := []*Topic{
		{ID: "getting-started", Title: "Getting Started", Tags: []string{"intro"}},
		{ID: "config", Title: "Configuration", Tags: []string{"setup"}},
	}

	html := RenderTopicPage(topic, sidebar)

	// Must have table of contents with section anchors
	assert.Contains(t, html, `href="#overview"`, "ToC should link to overview section")
	assert.Contains(t, html, `href="#installation"`, "ToC should link to installation section")
	assert.Contains(t, html, `href="#quick-start"`, "ToC should link to quick-start section")

	// Must contain section titles in the ToC
	assert.Contains(t, html, "Overview")
	assert.Contains(t, html, "Installation")
	assert.Contains(t, html, "Quick Start")

	// Must contain rendered markdown content
	assert.Contains(t, html, "<strong>guide</strong>")

	// Must have sidebar with topic links
	assert.Contains(t, html, "Configuration")

	// Must have ARIA roles
	assert.Contains(t, html, `role="banner"`)
	assert.Contains(t, html, `role="main"`)
	assert.Contains(t, html, `role="complementary"`)
	assert.Contains(t, html, `role="contentinfo"`)
}

func TestRenderLayout_Good_TopicPage_NoSidebar(t *testing.T) {
	topic := &Topic{
		ID:      "solo",
		Title:   "Solo Topic",
		Content: "Just content, no sidebar.",
	}

	html := RenderTopicPage(topic, nil)

	// Should still render content
	assert.Contains(t, html, "Solo Topic")
	assert.Contains(t, html, `role="main"`)

	// Without sidebar topics, left aside should not appear
	assert.NotContains(t, html, `role="complementary"`)
}

func TestRenderLayout_Good_SearchPage(t *testing.T) {
	results := []*SearchResult{
		{
			Topic:   &Topic{ID: "install", Title: "Installation Guide", Tags: []string{"setup"}},
			Score:   12.5,
			Snippet: "How to **install** the tool.",
		},
		{
			Topic:   &Topic{ID: "config", Title: "Configuration", Tags: []string{"setup"}},
			Score:   8.0,
			Snippet: "Set up your **config** file.",
		},
	}

	html := RenderSearchPage("install", results)

	// Must show the query
	assert.Contains(t, html, "install")

	// Must show result titles
	assert.Contains(t, html, "Installation Guide")
	assert.Contains(t, html, "Configuration")

	// Must have ARIA roles
	assert.Contains(t, html, `role="banner"`)
	assert.Contains(t, html, `role="main"`)
}

func TestRenderLayout_Good_SearchPage_NoResults(t *testing.T) {
	html := RenderSearchPage("nonexistent", nil)

	assert.Contains(t, html, "nonexistent")
	assert.Contains(t, html, "No results")
}

func TestRenderLayout_Good_HasDoctype(t *testing.T) {
	tests := []struct {
		name string
		html string
	}{
		{"index", RenderIndexPage(nil)},
		{"topic", RenderTopicPage(&Topic{ID: "t", Title: "T"}, nil)},
		{"search", RenderSearchPage("q", nil)},
		{"404", Render404Page()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.True(t, strings.HasPrefix(tt.html, "<!DOCTYPE html>"),
				"page should start with <!DOCTYPE html>, got: %s", tt.html[:min(50, len(tt.html))])
		})
	}
}

func TestRenderLayout_Good_404Page(t *testing.T) {
	html := Render404Page()

	assert.Contains(t, html, "Not Found")
	assert.Contains(t, html, "404")
	assert.Contains(t, html, `role="banner"`)
	assert.Contains(t, html, `role="main"`)
	assert.Contains(t, html, `role="contentinfo"`)
}

func TestRenderLayout_Good_EscapesHTML(t *testing.T) {
	topic := &Topic{
		ID:    "xss",
		Title: `<script>alert("xss")</script>`,
		Content: "Safe content.",
	}

	html := RenderIndexPage([]*Topic{topic})

	// Title must be escaped
	assert.NotContains(t, html, `<script>alert`)
	assert.Contains(t, html, "&lt;script&gt;")
}
