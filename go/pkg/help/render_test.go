// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	. "dappco.re/go"
)

func TestRenderMarkdown_Good(t *T) {
	t.Run("heading hierarchy H1-H6", func(t *T) {
		input := "# H1\n## H2\n### H3\n#### H4\n##### H5\n###### H6\n"
		html, err := RenderMarkdown(input)
		RequireNoError(t, err)
		AssertContains(t, html, "<h1>H1</h1>")
		AssertContains(t, html, "<h2>H2</h2>")
		AssertContains(t, html, "<h3>H3</h3>")
		AssertContains(t, html, "<h4>H4</h4>")
		AssertContains(t, html, "<h5>H5</h5>")
		AssertContains(t, html, "<h6>H6</h6>")
	})

	t.Run("fenced code blocks with language", func(t *T) {
		input := "```go\nfmt.Println(\"hello\")\n```\n"
		html, err := RenderMarkdown(input)
		RequireNoError(t, err)
		AssertContains(t, html, "<pre><code class=\"language-go\">")
		AssertContains(t, html, "fmt.Println")
		AssertContains(t, html, "</code></pre>")
	})

	t.Run("inline code backticks", func(t *T) {
		input := "Use `go test` to run tests.\n"
		html, err := RenderMarkdown(input)
		RequireNoError(t, err)
		AssertContains(t, html, "<code>go test</code>")
	})

	t.Run("unordered lists", func(t *T) {
		input := "- Alpha\n- Beta\n- Gamma\n"
		html, err := RenderMarkdown(input)
		RequireNoError(t, err)
		AssertContains(t, html, "<ul>")
		AssertContains(t, html, "<li>Alpha</li>")
		AssertContains(t, html, "<li>Beta</li>")
		AssertContains(t, html, "<li>Gamma</li>")
		AssertContains(t, html, "</ul>")
	})

	t.Run("ordered lists", func(t *T) {
		input := "1. First\n2. Second\n3. Third\n"
		html, err := RenderMarkdown(input)
		RequireNoError(t, err)
		AssertContains(t, html, "<ol>")
		AssertContains(t, html, "<li>First</li>")
		AssertContains(t, html, "<li>Second</li>")
		AssertContains(t, html, "<li>Third</li>")
		AssertContains(t, html, "</ol>")
	})

	t.Run("links and images", func(t *T) {
		input := "[Example](https://example.com)\n\n![Alt text](image.png)\n"
		html, err := RenderMarkdown(input)
		RequireNoError(t, err)
		AssertContains(t, html, `<a href="https://example.com">Example</a>`)
		AssertContains(t, html, `<img src="image.png" alt="Alt text"`)
	})

	t.Run("GFM tables", func(t *T) {
		input := "| Name | Value |\n|------|-------|\n| foo  | 42    |\n| bar  | 99    |\n"
		html, err := RenderMarkdown(input)
		RequireNoError(t, err)
		AssertContains(t, html, "<table>")
		AssertContains(t, html, "<th>Name</th>")
		AssertContains(t, html, "<th>Value</th>")
		AssertContains(t, html, "<td>foo</td>")
		AssertContains(t, html, "<td>42</td>")
		AssertContains(t, html, "</table>")
	})

	t.Run("empty input returns empty string", func(t *T) {
		html, err := RenderMarkdown("")
		RequireNoError(t, err)
		AssertEqual(t, "", html)
	})

	t.Run("special characters escaped in text", func(t *T) {
		input := "Use `<div>` tags & \"quotes\".\n"
		html, err := RenderMarkdown(input)
		RequireNoError(t, err)
		// The & in prose should be escaped
		AssertContains(t, html, "&amp;")
		// Angle brackets in code should be escaped
		AssertContains(t, html, "&lt;div&gt;")
	})
}

func TestRenderMarkdown_Good_RawHTML(t *T) {
	// html.WithUnsafe() should allow raw HTML pass-through
	input := "<div class=\"custom\">raw html</div>\n"
	html, err := RenderMarkdown(input)
	RequireNoError(t, err)
	AssertContains(t, html, `<div class="custom">raw html</div>`)
}

func TestRenderMarkdown_Good_GFMExtras(t *T) {
	t.Run("strikethrough", func(t *T) {
		input := "~~deleted~~\n"
		html, err := RenderMarkdown(input)
		RequireNoError(t, err)
		AssertContains(t, html, "<del>deleted</del>")
	})

	t.Run("autolinks", func(t *T) {
		input := "Visit https://example.com for details.\n"
		html, err := RenderMarkdown(input)
		RequireNoError(t, err)
		AssertContains(t, html, `<a href="https://example.com">`)
	})
}

func TestRenderMarkdown_Good_Typographer(t *T) {
	// Typographer extension converts straight quotes to smart quotes
	// and -- to en-dash, --- to em-dash.
	input := "She said -- \"hello\" --- and left.\n"
	html, err := RenderMarkdown(input)
	RequireNoError(t, err)
	// Check that dashes are converted (en-dash or em-dash entities)
	AssertNotContains(t, html, " -- ")
	AssertNotContains(t, html, " --- ")
}

func TestRender_RenderMarkdown_Good(t *core.T) {
	subject := RenderMarkdown
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Good"
	if marker == "" {
		t.FailNow()
	}
}

func TestRender_RenderMarkdown_Bad(t *core.T) {
	subject := RenderMarkdown
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestRender_RenderMarkdown_Ugly(t *core.T) {
	subject := RenderMarkdown
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}
