// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// searchIndexEntry represents a single topic in the client-side search index.
type searchIndexEntry struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Tags    []string `json:"tags"`
	Content string   `json:"content"`
}

// Generate writes a complete static help site to outputDir.
// It creates:
//   - index.html            -- topic listing grouped by tags
//   - topics/{id}.html      -- one page per topic
//   - search.html           -- client-side search with inline JS
//   - search-index.json     -- JSON index for client-side search
//   - 404.html              -- not found page
//
// All CSS is inlined; no external stylesheets are needed.
func Generate(catalog *Catalog, outputDir string) error {
	topics := catalog.List()

	// Ensure output directories exist.
	topicsDir := filepath.Join(outputDir, "topics")
	if err := os.MkdirAll(topicsDir, 0o755); err != nil {
		return err
	}

	// 1. index.html
	if err := writeStaticPage(outputDir, "index.html", "index.html", indexData{
		Topics: topics,
		Groups: groupTopicsByTag(topics),
	}); err != nil {
		return err
	}

	// 2. topics/{id}.html -- one per topic
	for _, t := range topics {
		if err := writeStaticPage(topicsDir, t.ID+".html", "topic.html", topicData{Topic: t}); err != nil {
			return err
		}
	}

	// 3. search.html -- client-side search page
	if err := writeSearchPage(outputDir); err != nil {
		return err
	}

	// 4. search-index.json
	if err := writeSearchIndex(outputDir, topics); err != nil {
		return err
	}

	// 5. 404.html
	if err := writeStaticPage(outputDir, "404.html", "404.html", nil); err != nil {
		return err
	}

	return nil
}

// writeStaticPage renders a template page to a file.
func writeStaticPage(dir, filename, templatePage string, data any) error {
	path := filepath.Join(dir, filename)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return renderPage(f, templatePage, data)
}

// writeSearchIndex writes the JSON search index for client-side search.
func writeSearchIndex(outputDir string, topics []*Topic) error {
	entries := make([]searchIndexEntry, 0, len(topics))
	for _, t := range topics {
		// Truncate content for the index to keep file size reasonable.
		content := t.Content
		runes := []rune(content)
		if len(runes) > 500 {
			content = string(runes[:500])
		}
		entries = append(entries, searchIndexEntry{
			ID:      t.ID,
			Title:   t.Title,
			Tags:    t.Tags,
			Content: content,
		})
	}

	path := filepath.Join(outputDir, "search-index.json")
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(entries)
}

// writeSearchPage generates search.html with inline client-side JS search.
// The JS uses escapeHTML() on all data before DOM insertion to prevent XSS.
// Data comes from our own search-index.json, not external user input.
func writeSearchPage(outputDir string) error {
	path := filepath.Join(outputDir, "search.html")
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Render via a search template with empty results + inject client-side JS.
	data := searchData{Query: "", Results: nil}
	if err := renderPage(f, "search.html", data); err != nil {
		return err
	}

	// Append inline script for client-side search.
	_, err = f.WriteString(clientSearchScript)
	return err
}

// clientSearchScript is the inline JS for static-site client-side search.
// All values are escaped via escapeHTML() before DOM insertion.
// The search index is generated from our own catalog data, not user input.
const clientSearchScript = `
<script>
(function() {
    let index = [];

    fetch('search-index.json')
        .then(r => r.json())
        .then(data => { index = data; })
        .catch(() => {});

    const form = document.querySelector('.search-form');
    const input = form ? form.querySelector('input[name="q"]') : null;
    const main = document.querySelector('main');
    const container = main ? main.querySelector('.container') : null;

    if (form && input) {
        form.addEventListener('submit', function(e) {
            e.preventDefault();
            doSearch(input.value.trim());
        });
    }

    function escapeHTML(s) {
        const div = document.createElement('div');
        div.textContent = s;
        return div.innerHTML;
    }

    function doSearch(query) {
        if (!query || !container) return;
        const lower = query.toLowerCase();
        const words = lower.split(/\s+/);

        const results = index
            .map(function(entry) {
                let score = 0;
                const title = (entry.title || '').toLowerCase();
                const content = (entry.content || '').toLowerCase();
                const tags = (entry.tags || []).map(function(t) { return t.toLowerCase(); });

                words.forEach(function(w) {
                    if (title.indexOf(w) !== -1) score += 10;
                    if (content.indexOf(w) !== -1) score += 1;
                    tags.forEach(function(tag) { if (tag.indexOf(w) !== -1) score += 3; });
                });
                return { entry: entry, score: score };
            })
            .filter(function(r) { return r.score > 0; })
            .sort(function(a, b) { return b.score - a.score; });

        // Build result using safe DOM methods
        while (container.firstChild) container.removeChild(container.firstChild);

        var h1 = document.createElement('h1');
        h1.textContent = 'Search Results';
        container.appendChild(h1);

        var summary = document.createElement('p');
        summary.style.color = 'var(--fg-muted)';
        if (results.length > 0) {
            summary.textContent = 'Found ' + results.length + ' result' + (results.length !== 1 ? 's' : '') + ' for \u201c' + query + '\u201d';
        } else {
            summary.textContent = 'No results for \u201c' + query + '\u201d';
        }
        container.appendChild(summary);

        results.forEach(function(r) {
            var e = r.entry;
            var card = document.createElement('div');
            card.className = 'card';

            var heading = document.createElement('h3');
            var link = document.createElement('a');
            link.href = 'topics/' + encodeURIComponent(e.id) + '.html';
            link.textContent = e.title;
            heading.appendChild(link);
            card.appendChild(heading);

            if (e.content) {
                var p = document.createElement('p');
                var snippet = e.content.substring(0, 150);
                p.textContent = snippet + (e.content.length > 150 ? '...' : '');
                card.appendChild(p);
            }

            if (e.tags && e.tags.length) {
                var tagDiv = document.createElement('div');
                e.tags.forEach(function(t) {
                    var span = document.createElement('span');
                    span.className = 'tag';
                    span.textContent = t;
                    tagDiv.appendChild(span);
                });
                card.appendChild(tagDiv);
            }

            container.appendChild(card);
        });

        if (results.length === 0) {
            var noResults = document.createElement('div');
            noResults.style.cssText = 'margin-top:2rem;text-align:center;color:var(--fg-muted);';
            var tip = document.createElement('p');
            tip.textContent = 'Try a different search term or browse ';
            var browseLink = document.createElement('a');
            browseLink.href = 'index.html';
            browseLink.textContent = 'all topics';
            tip.appendChild(browseLink);
            tip.appendChild(document.createTextNode('.'));
            noResults.appendChild(tip);
            container.appendChild(noResults);
        }
    }
})();
</script>
`
