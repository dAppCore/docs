// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRenderMarkdown_Good(t *testing.T) {
	t.Run("heading hierarchy H1-H6", func(t *testing.T) {
		input := "# H1\n## H2\n### H3\n#### H4\n##### H5\n###### H6\n"
		html, err := RenderMarkdown(input)
		require.NoError(t, err)
		assert.Contains(t, html, "<h1>H1</h1>")
		assert.Contains(t, html, "<h2>H2</h2>")
		assert.Contains(t, html, "<h3>H3</h3>")
		assert.Contains(t, html, "<h4>H4</h4>")
		assert.Contains(t, html, "<h5>H5</h5>")
		assert.Contains(t, html, "<h6>H6</h6>")
	})

	t.Run("fenced code blocks with language", func(t *testing.T) {
		input := "```go\nfmt.Println(\"hello\")\n```\n"
		html, err := RenderMarkdown(input)
		require.NoError(t, err)
		assert.Contains(t, html, "<pre><code class=\"language-go\">")
		assert.Contains(t, html, "fmt.Println")
		assert.Contains(t, html, "</code></pre>")
	})

	t.Run("inline code backticks", func(t *testing.T) {
		input := "Use `go test` to run tests.\n"
		html, err := RenderMarkdown(input)
		require.NoError(t, err)
		assert.Contains(t, html, "<code>go test</code>")
	})

	t.Run("unordered lists", func(t *testing.T) {
		input := "- Alpha\n- Beta\n- Gamma\n"
		html, err := RenderMarkdown(input)
		require.NoError(t, err)
		assert.Contains(t, html, "<ul>")
		assert.Contains(t, html, "<li>Alpha</li>")
		assert.Contains(t, html, "<li>Beta</li>")
		assert.Contains(t, html, "<li>Gamma</li>")
		assert.Contains(t, html, "</ul>")
	})

	t.Run("ordered lists", func(t *testing.T) {
		input := "1. First\n2. Second\n3. Third\n"
		html, err := RenderMarkdown(input)
		require.NoError(t, err)
		assert.Contains(t, html, "<ol>")
		assert.Contains(t, html, "<li>First</li>")
		assert.Contains(t, html, "<li>Second</li>")
		assert.Contains(t, html, "<li>Third</li>")
		assert.Contains(t, html, "</ol>")
	})

	t.Run("links and images", func(t *testing.T) {
		input := "[Example](https://example.com)\n\n![Alt text](image.png)\n"
		html, err := RenderMarkdown(input)
		require.NoError(t, err)
		assert.Contains(t, html, `<a href="https://example.com">Example</a>`)
		assert.Contains(t, html, `<img src="image.png" alt="Alt text"`)
	})

	t.Run("GFM tables", func(t *testing.T) {
		input := "| Name | Value |\n|------|-------|\n| foo  | 42    |\n| bar  | 99    |\n"
		html, err := RenderMarkdown(input)
		require.NoError(t, err)
		assert.Contains(t, html, "<table>")
		assert.Contains(t, html, "<th>Name</th>")
		assert.Contains(t, html, "<th>Value</th>")
		assert.Contains(t, html, "<td>foo</td>")
		assert.Contains(t, html, "<td>42</td>")
		assert.Contains(t, html, "</table>")
	})

	t.Run("empty input returns empty string", func(t *testing.T) {
		html, err := RenderMarkdown("")
		require.NoError(t, err)
		assert.Equal(t, "", html)
	})

	t.Run("special characters escaped in text", func(t *testing.T) {
		input := "Use `<div>` tags & \"quotes\".\n"
		html, err := RenderMarkdown(input)
		require.NoError(t, err)
		// The & in prose should be escaped
		assert.Contains(t, html, "&amp;")
		// Angle brackets in code should be escaped
		assert.Contains(t, html, "&lt;div&gt;")
	})
}

func TestRenderMarkdown_Good_RawHTML(t *testing.T) {
	// html.WithUnsafe() should allow raw HTML pass-through
	input := "<div class=\"custom\">raw html</div>\n"
	html, err := RenderMarkdown(input)
	require.NoError(t, err)
	assert.Contains(t, html, `<div class="custom">raw html</div>`)
}

func TestRenderMarkdown_Good_GFMExtras(t *testing.T) {
	t.Run("strikethrough", func(t *testing.T) {
		input := "~~deleted~~\n"
		html, err := RenderMarkdown(input)
		require.NoError(t, err)
		assert.Contains(t, html, "<del>deleted</del>")
	})

	t.Run("autolinks", func(t *testing.T) {
		input := "Visit https://example.com for details.\n"
		html, err := RenderMarkdown(input)
		require.NoError(t, err)
		assert.Contains(t, html, `<a href="https://example.com">`)
	})
}

func TestRenderMarkdown_Good_Typographer(t *testing.T) {
	// Typographer extension converts straight quotes to smart quotes
	// and -- to en-dash, --- to em-dash.
	input := "She said -- \"hello\" --- and left.\n"
	html, err := RenderMarkdown(input)
	require.NoError(t, err)
	// Check that dashes are converted (en-dash or em-dash entities)
	assert.NotContains(t, html, " -- ")
	assert.NotContains(t, html, " --- ")
}
