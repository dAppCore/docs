// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	"fmt"
	"strings"
	"testing"
)

// titleCase capitalises the first letter of a string.
// Used in benchmarks to avoid deprecated strings.Title.
func titleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// buildLargeCatalog creates a search index with n topics for benchmarking.
// Each topic has a title, content with multiple paragraphs, sections, and tags.
func buildLargeCatalog(n int) *searchIndex {
	idx := newSearchIndex()

	// Word pools for generating varied content
	subjects := []string{
		"configuration", "deployment", "monitoring", "testing", "debugging",
		"authentication", "authorisation", "networking", "storage", "logging",
		"caching", "scheduling", "routing", "migration", "backup",
		"encryption", "compression", "validation", "serialisation", "templating",
	}
	verbs := []string{
		"install", "configure", "deploy", "monitor", "debug",
		"authenticate", "authorise", "connect", "store", "analyse",
		"cache", "schedule", "route", "migrate", "restore",
	}
	adjectives := []string{
		"advanced", "basic", "custom", "distributed", "encrypted",
		"federated", "graceful", "hybrid", "incremental", "just-in-time",
	}

	for i := range n {
		subj := subjects[i%len(subjects)]
		verb := verbs[i%len(verbs)]
		adj := adjectives[i%len(adjectives)]

		title := fmt.Sprintf("%s %s Guide %d", titleCase(adj), titleCase(subj), i)
		content := fmt.Sprintf(
			"This guide covers how to %s %s %s systems. "+
				"It includes step-by-step instructions for setting up %s "+
				"in both development and production environments. "+
				"The %s process requires careful planning and %s tools. "+
				"Make sure to review the prerequisites before starting.",
			verb, adj, subj, subj, subj, adj,
		)

		sections := []Section{
			{
				ID:      fmt.Sprintf("overview-%d", i),
				Title:   "Overview",
				Content: fmt.Sprintf("An overview of %s %s patterns and best practices.", adj, subj),
			},
			{
				ID:      fmt.Sprintf("setup-%d", i),
				Title:   fmt.Sprintf("%s Setup", titleCase(subj)),
				Content: fmt.Sprintf("Detailed setup instructions for %s. Run the %s command to begin.", subj, verb),
			},
			{
				ID:      fmt.Sprintf("troubleshooting-%d", i),
				Title:   "Troubleshooting",
				Content: fmt.Sprintf("Common issues when working with %s and how to resolve them.", subj),
			},
		}

		idx.Add(&Topic{
			ID:       fmt.Sprintf("%s-%s-%d", adj, subj, i),
			Title:    title,
			Content:  content,
			Sections: sections,
			Tags:     []string{subj, adj, verb, "guide"},
		})
	}

	return idx
}

func BenchmarkSearch_SingleWord(b *testing.B) {
	idx := buildLargeCatalog(200)
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		idx.Search("configuration")
	}
}

func BenchmarkSearch_MultiWord(b *testing.B) {
	idx := buildLargeCatalog(200)
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		idx.Search("advanced deployment guide")
	}
}

func BenchmarkSearch_NoResults(b *testing.B) {
	idx := buildLargeCatalog(200)
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		idx.Search("xylophone")
	}
}

func BenchmarkSearch_PartialMatch(b *testing.B) {
	idx := buildLargeCatalog(200)
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		idx.Search("config")
	}
}

func BenchmarkSearch_LargeCatalog500(b *testing.B) {
	idx := buildLargeCatalog(500)
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		idx.Search("deployment monitoring")
	}
}

func BenchmarkSearch_LargeCatalog1000(b *testing.B) {
	idx := buildLargeCatalog(1000)
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		idx.Search("testing guide")
	}
}

func BenchmarkSearchIndex_Add(b *testing.B) {
	// Benchmark the indexing/add path
	topic := &Topic{
		ID:      "bench-topic",
		Title:   "Benchmark Topic Title",
		Content: "This is benchmark content with several words for indexing purposes.",
		Tags:    []string{"bench", "performance"},
		Sections: []Section{
			{ID: "s1", Title: "First Section", Content: "Section content for benchmarking."},
			{ID: "s2", Title: "Second Section", Content: "More section content here."},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		idx := newSearchIndex()
		idx.Add(topic)
	}
}

func BenchmarkTokenize(b *testing.B) {
	text := "The quick brown fox jumps over the lazy dog. Configuration and deployment are covered in detail."
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		tokenize(text)
	}
}
