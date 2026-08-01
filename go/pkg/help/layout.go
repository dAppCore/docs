// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	"html"

	core "dappco.re/go"
)

// pageCSS is the shared dark-theme stylesheet embedded in every page.
const pageCSS = `
:root {
	--bg: #0d1117;
	--bg-secondary: #161b22;
	--bg-tertiary: #21262d;
	--fg: #c9d1d9;
	--fg-muted: #8b949e;
	--fg-subtle: #6e7681;
	--accent: #58a6ff;
	--accent-hover: #79c0ff;
	--border: #30363d;
}
* { margin: 0; padding: 0; box-sizing: border-box; }
body {
	font-family: ui-monospace, 'Cascadia Code', 'Source Code Pro', Menlo, Consolas, 'DejaVu Sans Mono', monospace;
	background: var(--bg);
	color: var(--fg);
	line-height: 1.6;
	font-size: 14px;
}
a { color: var(--accent); text-decoration: none; }
a:hover { color: var(--accent-hover); text-decoration: underline; }
.container { max-width: 960px; margin: 0 auto; padding: 0 1.5rem; }
nav {
	background: var(--bg-secondary);
	border-bottom: 1px solid var(--border);
	padding: 0.75rem 0;
}
nav .container {
	display: flex;
	align-items: center;
	justify-content: space-between;
	gap: 1rem;
}
nav .brand {
	font-weight: bold;
	font-size: 1rem;
	color: var(--fg);
	white-space: nowrap;
}
nav .brand:hover { color: var(--accent); text-decoration: none; }
.search-form { display: flex; flex: 1; max-width: 400px; }
.search-form input {
	flex: 1;
	background: var(--bg-tertiary);
	border: 1px solid var(--border);
	border-radius: 6px;
	color: var(--fg);
	padding: 0.4rem 0.75rem;
	font-family: inherit;
	font-size: 0.85rem;
}
.search-form input::placeholder { color: var(--fg-subtle); }
.search-form input:focus { outline: none; border-color: var(--accent); }
.page-body { display: flex; gap: 2rem; padding: 2rem 0; min-height: calc(100vh - 120px); }
.page-body > main { flex: 1; min-width: 0; }
.page-body > aside { width: 220px; flex-shrink: 0; }
.sidebar-nav { position: sticky; top: 1rem; }
.sidebar-nav h3 { font-size: 0.85rem; color: var(--fg-muted); margin-top: 0; }
.sidebar-nav ul { list-style: none; padding-left: 0; font-size: 0.8rem; }
.sidebar-nav li { margin: 0.3rem 0; }
.sidebar-nav a { color: var(--fg-muted); }
.sidebar-nav a:hover { color: var(--accent); }
.sidebar-group { margin-bottom: 1.25rem; }
.sidebar-group-title { font-size: 0.7rem; text-transform: uppercase; letter-spacing: 0.05em; color: var(--fg-subtle); margin-bottom: 0.25rem; }
footer {
	border-top: 1px solid var(--border);
	padding: 1rem 0;
	text-align: center;
	color: var(--fg-muted);
	font-size: 0.8rem;
}
h1, h2, h3, h4, h5, h6 {
	color: var(--fg);
	margin: 1.5rem 0 0.75rem;
	line-height: 1.3;
}
h1 { font-size: 1.5rem; border-bottom: 1px solid var(--border); padding-bottom: 0.5rem; }
h2 { font-size: 1.25rem; }
h3 { font-size: 1.1rem; }
p { margin: 0.5rem 0; }
pre {
	background: var(--bg-secondary);
	border: 1px solid var(--border);
	border-radius: 6px;
	padding: 1rem;
	overflow-x: auto;
	margin: 0.75rem 0;
}
code {
	background: var(--bg-tertiary);
	padding: 0.15rem 0.3rem;
	border-radius: 3px;
	font-size: 0.9em;
}
pre code { background: none; padding: 0; }
table {
	border-collapse: collapse;
	width: 100%;
	margin: 0.75rem 0;
}
th, td {
	border: 1px solid var(--border);
	padding: 0.5rem 0.75rem;
	text-align: left;
}
th { background: var(--bg-secondary); font-weight: 600; }
ul, ol { padding-left: 1.5rem; margin: 0.5rem 0; }
li { margin: 0.25rem 0; }
.tag {
	display: inline-block;
	background: var(--bg-tertiary);
	border: 1px solid var(--border);
	border-radius: 12px;
	padding: 0.1rem 0.5rem;
	font-size: 0.75rem;
	color: var(--accent);
	margin: 0.15rem 0.15rem;
}
.card {
	background: var(--bg-secondary);
	border: 1px solid var(--border);
	border-radius: 6px;
	padding: 1rem 1.25rem;
	margin: 0.75rem 0;
	transition: border-color 0.15s;
}
.card:hover { border-color: var(--accent); }
.card h3 { margin-top: 0; font-size: 1rem; }
.card p { color: var(--fg-muted); font-size: 0.85rem; }
.badge {
	display: inline-block;
	background: var(--bg-tertiary);
	border-radius: 10px;
	padding: 0.1rem 0.5rem;
	font-size: 0.7rem;
	color: var(--fg-muted);
}
.toc ul { list-style: none; padding-left: 0; font-size: 0.85rem; }
.toc li { margin: 0.3rem 0; }
.toc a { color: var(--fg-muted); }
.toc a:hover { color: var(--accent); }
.centre { text-align: center; }
`

// wrapPage wraps semantic page regions in a complete HTML document.
func wrapPage(title string, body string) string {
	return core.Sprintf(
		`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8">`+
			`<meta name="viewport" content="width=device-width, initial-scale=1.0">`+
			`<title>%s</title><style>%s</style></head><body>`,
		html.EscapeString(title), pageCSS,
	) + body + `</body></html>`
}

func renderHCF(title, header, content, footer string) string {
	return wrapPage(title,
		`<header role="banner" data-block="H">`+header+`</header>`+
			`<main role="main" data-block="C">`+content+`</main>`+
			`<footer role="contentinfo" data-block="F">`+footer+`</footer>`,
	)
}

func renderHLCF(title, header, sidebar, content, footer string) string {
	return wrapPage(title,
		`<header role="banner" data-block="H">`+header+`</header>`+
			`<aside role="complementary" data-block="L">`+sidebar+`</aside>`+
			`<main role="main" data-block="C">`+content+`</main>`+
			`<footer role="contentinfo" data-block="F">`+footer+`</footer>`,
	)
}

// headerNav returns the HLCRF Header node: nav bar with brand + search form.
func headerNav(searchValue string) string {
	escapedValue := html.EscapeString(searchValue)
	return core.Sprintf(
		`<nav><div class="container">`+
			`<a href="/" class="brand">core.help</a>`+
			`<form class="search-form" action="/search" method="get">`+
			`<input type="text" name="q" placeholder="Search topics..." value="%s" autocomplete="off">`+
			`</form>`+
			`</div></nav>`,
		escapedValue,
	)
}

// footerContent returns the HLCRF Footer node: licence + source link.
func footerContent() string {
	return `<div class="container">` +
		`EUPL-1.2 &middot; <a href="https://forge.lthn.ai/core/docs">forge.lthn.ai/core/docs</a>` +
		`</div>`
}

// sidebarTopicTree returns a sidebar node showing topics grouped by tag.
func sidebarTopicTree(topics []*Topic) string {
	if len(topics) == 0 {
		return ""
	}

	groups := groupTopicsByTag(topics)
	var b core.Builder
	b.WriteString(`<div class="sidebar-nav">`)
	b.WriteString(`<h3>Topics</h3>`)

	for _, g := range groups {
		b.WriteString(`<div class="sidebar-group">`)
		b.WriteString(core.Sprintf(`<div class="sidebar-group-title">%s</div>`, html.EscapeString(g.Tag)))
		b.WriteString(`<ul>`)
		for _, t := range g.Topics {
			b.WriteString(core.Sprintf(
				`<li><a href="/topics/%s">%s</a></li>`,
				html.EscapeString(t.ID),
				html.EscapeString(t.Title),
			))
		}
		b.WriteString(`</ul>`)
		b.WriteString(`</div>`)
	}

	b.WriteString(`</div>`)
	return b.String()
}

// topicTableOfContents returns a ToC node with section anchors.
func topicTableOfContents(sections []Section) string {
	if len(sections) == 0 {
		return ""
	}

	var b core.Builder
	b.WriteString(`<div class="toc"><h3>On this page</h3><ul>`)
	for _, s := range sections {
		indent := (s.Level - 1) * 12
		b.WriteString(core.Sprintf(
			`<li style="padding-left: %dpx;"><a href="#%s">%s</a></li>`,
			indent,
			html.EscapeString(s.ID),
			html.EscapeString(s.Title),
		))
	}
	b.WriteString(`</ul></div>`)
	return b.String()
}

// truncateContent returns a plain-text preview of markdown content.
func truncateContent(content string, maxLen int) string {
	lines := core.Split(content, "\n")
	var clean []string
	for _, line := range lines {
		trimmed := core.Trim(line)
		if trimmed == "" || core.HasPrefix(trimmed, "#") {
			continue
		}
		clean = append(clean, trimmed)
	}
	text := core.Join(" ", clean...)
	runes := []rune(text)
	if len(runes) <= maxLen {
		return text
	}
	return string(runes[:maxLen]) + "..."
}

// RenderIndexPage renders a full HTML page listing topics grouped by tag.
func RenderIndexPage(topics []*Topic) string {
	var content core.Builder
	content.WriteString(`<div class="container">`)

	count := len(topics)
	noun := "topics"
	if count == 1 {
		noun = "topic"
	}
	content.WriteString(core.Sprintf(
		`<h1>Help Topics <span class="badge">%d %s</span></h1>`,
		count, noun,
	))

	if count > 0 {
		groups := groupTopicsByTag(topics)
		for _, g := range groups {
			content.WriteString(core.Sprintf(
				`<h2><span class="tag">%s</span></h2>`,
				html.EscapeString(g.Tag),
			))
			for _, t := range g.Topics {
				content.WriteString(`<div class="card">`)
				content.WriteString(core.Sprintf(
					`<h3><a href="/topics/%s">%s</a></h3>`,
					html.EscapeString(t.ID),
					html.EscapeString(t.Title),
				))
				if len(t.Tags) > 0 {
					content.WriteString(`<div>`)
					for _, tag := range t.Tags {
						content.WriteString(core.Sprintf(
							`<span class="tag">%s</span>`,
							html.EscapeString(tag),
						))
					}
					content.WriteString(`</div>`)
				}
				if t.Content != "" {
					content.WriteString(core.Sprintf(
						`<p>%s</p>`,
						html.EscapeString(truncateContent(t.Content, 120)),
					))
				}
				content.WriteString(`</div>`)
			}
		}
	} else {
		content.WriteString(`<p style="color: var(--fg-muted);">No topics available.</p>`)
	}

	content.WriteString(`</div>`)

	return renderHCF("Help Topics", headerNav(""), content.String(), footerContent())
}

// RenderTopicPage renders a full HTML page for a single topic with ToC and sidebar.
func RenderTopicPage(topic *Topic, sidebar []*Topic) string {
	hasSidebar := len(sidebar) > 0

	var content core.Builder
	content.WriteString(`<div class="container">`)

	// Tags
	if len(topic.Tags) > 0 {
		content.WriteString(`<div style="margin-bottom: 1rem;">`)
		for _, tag := range topic.Tags {
			content.WriteString(core.Sprintf(
				`<span class="tag">%s</span>`,
				html.EscapeString(tag),
			))
		}
		content.WriteString(`</div>`)
	}

	// Table of contents (inline, before content)
	if len(topic.Sections) > 0 {
		content.WriteString(topicTableOfContents(topic.Sections))
	}

	// Rendered markdown body
	rendered := "<p>Error rendering content.</p>"
	if r := RenderMarkdown(topic.Content); r.OK {
		rendered = r.Value.(string)
	}
	content.WriteString(`<article class="topic-body">`)
	content.WriteString(rendered)
	content.WriteString(`</article>`)

	// Related topics
	if len(topic.Related) > 0 {
		content.WriteString(`<div style="margin-top: 1.5rem;">`)
		content.WriteString(`<h3 style="font-size: 0.85rem; color: var(--fg-muted);">Related</h3>`)
		content.WriteString(`<ul style="list-style: none; padding-left: 0; font-size: 0.8rem;">`)
		for _, rel := range topic.Related {
			content.WriteString(core.Sprintf(
				`<li style="margin: 0.3rem 0;"><a href="/topics/%s">%s</a></li>`,
				html.EscapeString(rel),
				html.EscapeString(rel),
			))
		}
		content.WriteString(`</ul></div>`)
	}

	content.WriteString(`</div>`)

	title := html.EscapeString(topic.Title) + " - Help"
	if hasSidebar {
		return renderHLCF(title, headerNav(""), sidebarTopicTree(sidebar), content.String(), footerContent())
	}
	return renderHCF(title, headerNav(""), content.String(), footerContent())
}

// RenderSearchPage renders a full HTML page showing search results.
func RenderSearchPage(query string, results []*SearchResult) string {
	var content core.Builder
	content.WriteString(`<div class="container">`)
	content.WriteString(`<h1>Search Results</h1>`)

	escapedQuery := html.EscapeString(query)
	if len(results) > 0 {
		noun := "results"
		if len(results) == 1 {
			noun = "result"
		}
		content.WriteString(core.Sprintf(
			`<p style="color: var(--fg-muted);">Found %d %s for &ldquo;%s&rdquo;</p>`,
			len(results), noun, escapedQuery,
		))

		for _, r := range results {
			content.WriteString(`<div class="card">`)
			content.WriteString(core.Sprintf(
				`<h3><a href="/topics/%s">%s</a> <span class="badge">%.1f</span></h3>`,
				html.EscapeString(r.Topic.ID),
				html.EscapeString(r.Topic.Title),
				r.Score,
			))
			if r.Snippet != "" {
				content.WriteString(core.Sprintf(`<p>%s</p>`, r.Snippet))
			}
			if len(r.Topic.Tags) > 0 {
				content.WriteString(`<div>`)
				for _, tag := range r.Topic.Tags {
					content.WriteString(core.Sprintf(
						`<span class="tag">%s</span>`,
						html.EscapeString(tag),
					))
				}
				content.WriteString(`</div>`)
			}
			content.WriteString(`</div>`)
		}
	} else {
		content.WriteString(core.Sprintf(
			`<p style="color: var(--fg-muted);">No results for &ldquo;%s&rdquo;</p>`,
			escapedQuery,
		))
		content.WriteString(`<div style="margin-top: 2rem; text-align: center; color: var(--fg-muted);">`)
		content.WriteString(`<p>Try a different search term or browse <a href="/">all topics</a>.</p>`)
		content.WriteString(`</div>`)
	}

	content.WriteString(`</div>`)

	return renderHCF("Search: "+html.EscapeString(query)+" - Help", headerNav(query), content.String(), footerContent())
}

// Render404Page renders a full HTML "Not Found" page.
func Render404Page() string {
	content := `<div class="container centre" style="margin-top: 4rem;">` +
		`<h1 style="font-size: 3rem; border: none; color: var(--fg-muted);">404</h1>` +
		`<p style="font-size: 1.1rem; color: var(--fg-muted);">Not Found</p>` +
		`<p style="margin-top: 1.5rem;">` +
		`<a href="/">Browse all topics</a> or try searching above.` +
		`</p></div>`

	return renderHCF("Not Found - Help", headerNav(""), content, footerContent())
}
