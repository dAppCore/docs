// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseTemplates_Good(t *testing.T) {
	pages := []string{"index.html", "topic.html", "search.html", "404.html"}
	for _, page := range pages {
		t.Run(page, func(t *testing.T) {
			tmpl, err := parseTemplates(page)
			require.NoError(t, err, "template %s should parse without error", page)
			assert.NotNil(t, tmpl)
		})
	}
}

func TestRenderPage_Good_Index(t *testing.T) {
	topics := []*Topic{
		{ID: "getting-started", Title: "Getting Started", Tags: []string{"intro"}, Content: "Welcome."},
		{ID: "config", Title: "Configuration", Tags: []string{"setup"}, Content: "Config options."},
	}

	data := indexData{
		Topics: topics,
		Groups: groupTopicsByTag(topics),
	}

	var buf bytes.Buffer
	err := renderPage(&buf, "index.html", data)
	require.NoError(t, err)

	html := buf.String()
	assert.Contains(t, html, "Getting Started")
	assert.Contains(t, html, "Configuration")
	assert.Contains(t, html, "2 topics")
	assert.Contains(t, html, "core help")
}

func TestRenderPage_Good_Topic(t *testing.T) {
	topic := &Topic{
		ID:      "getting-started",
		Title:   "Getting Started",
		Content: "# Getting Started\n\nWelcome to the **guide**.\n",
		Tags:    []string{"intro"},
		Sections: []Section{
			{ID: "overview", Title: "Overview", Level: 2},
		},
		Related: []string{"config"},
	}

	data := topicData{Topic: topic}

	var buf bytes.Buffer
	err := renderPage(&buf, "topic.html", data)
	require.NoError(t, err)

	html := buf.String()
	assert.Contains(t, html, "Getting Started")
	assert.Contains(t, html, "<strong>guide</strong>")
	assert.Contains(t, html, "Overview")
	assert.Contains(t, html, "config")
}

func TestRenderPage_Good_Search(t *testing.T) {
	data := searchData{
		Query: "install",
		Results: []*SearchResult{
			{
				Topic:   &Topic{ID: "install", Title: "Installation", Tags: []string{"setup"}},
				Score:   12.5,
				Snippet: "How to **install** the tool.",
			},
		},
	}

	var buf bytes.Buffer
	err := renderPage(&buf, "search.html", data)
	require.NoError(t, err)

	html := buf.String()
	assert.Contains(t, html, "install")
	assert.Contains(t, html, "Installation")
	assert.Contains(t, html, "1 result")
	assert.Contains(t, html, "12.5")
}

func TestRenderPage_Good_404(t *testing.T) {
	var buf bytes.Buffer
	err := renderPage(&buf, "404.html", nil)
	require.NoError(t, err)

	html := buf.String()
	assert.Contains(t, html, "not found")
	assert.Contains(t, html, "404")
}

func TestGroupTopicsByTag_Good(t *testing.T) {
	topics := []*Topic{
		{ID: "a", Title: "Alpha", Tags: []string{"setup"}, Order: 2},
		{ID: "b", Title: "Beta", Tags: []string{"setup"}, Order: 1},
		{ID: "c", Title: "Gamma", Tags: []string{"advanced"}},
		{ID: "d", Title: "Delta"}, // no tags -> "other"
	}

	groups := groupTopicsByTag(topics)
	require.Len(t, groups, 3)

	// Groups should be sorted alphabetically by tag
	assert.Equal(t, "advanced", groups[0].Tag)
	assert.Equal(t, "other", groups[1].Tag)
	assert.Equal(t, "setup", groups[2].Tag)

	// Within "setup", topics should be sorted by Order then Title
	setupGroup := groups[2]
	require.Len(t, setupGroup.Topics, 2)
	assert.Equal(t, "Beta", setupGroup.Topics[0].Title)  // Order 1
	assert.Equal(t, "Alpha", setupGroup.Topics[1].Title) // Order 2
}

func TestTemplateFuncs_Good(t *testing.T) {
	fns := templateFuncs()

	t.Run("truncate short string", func(t *testing.T) {
		fn := fns["truncate"].(func(string, int) string)
		assert.Equal(t, "hello", fn("hello", 10))
	})

	t.Run("truncate long string", func(t *testing.T) {
		fn := fns["truncate"].(func(string, int) string)
		result := fn("hello world this is long", 11)
		assert.Equal(t, "hello world...", result)
	})

	t.Run("truncate strips headings", func(t *testing.T) {
		fn := fns["truncate"].(func(string, int) string)
		result := fn("# Title\n\nSome content here.", 100)
		assert.Equal(t, "Some content here.", result)
		assert.NotContains(t, result, "#")
	})

	t.Run("pluralise singular", func(t *testing.T) {
		fn := fns["pluralise"].(func(int, string, string) string)
		assert.Equal(t, "topic", fn(1, "topic", "topics"))
	})

	t.Run("pluralise plural", func(t *testing.T) {
		fn := fns["pluralise"].(func(int, string, string) string)
		assert.Equal(t, "topics", fn(0, "topic", "topics"))
		assert.Equal(t, "topics", fn(5, "topic", "topics"))
	})

	t.Run("multiply", func(t *testing.T) {
		fn := fns["multiply"].(func(int, int) int)
		assert.Equal(t, 24, fn(4, 6))
	})

	t.Run("sub", func(t *testing.T) {
		fn := fns["sub"].(func(int, int) int)
		assert.Equal(t, 2, fn(5, 3))
	})
}
