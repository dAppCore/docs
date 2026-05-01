// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

// RenderMarkdown converts Markdown content to an HTML fragment.
// It uses goldmark with GitHub Flavoured Markdown (tables, strikethrough,
// autolinks), smart quotes/dashes (typographer), and allows raw HTML
// in the source for embedded code examples.
//
// The returned string is an HTML fragment without <html>/<body> wrappers;
// the server templates handle the page structure.
func RenderMarkdown(content string) (string, error) {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Typographer,
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)

	var buf bytes.Buffer
	if err := md.Convert([]byte(content), &buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}
