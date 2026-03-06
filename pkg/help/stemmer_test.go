// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ---------------------------------------------------------------------------
// stem() unit tests
// ---------------------------------------------------------------------------

func TestStem_Good(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Step 1: plurals and verb inflections
		{name: "sses to ss", input: "addresses", expected: "address"},
		{name: "ies to i", input: "eries", expected: "eri"},
		{name: "ies to i (ponies)", input: "ponies", expected: "poni"},
		{name: "eed to ee", input: "agreed", expected: "agree"},
		{name: "ed removed", input: "configured", expected: "configur"},
		{name: "ing removed", input: "running", expected: "runn"},
		{name: "ing removed (testing)", input: "testing", expected: "test"},
		{name: "s removed (servers)", input: "servers", expected: "server"},
		{name: "s removed then derivational (configurations)", input: "configurations", expected: "configurate"},
		{name: "ss unchanged", input: "boss", expected: "boss"},

		// Step 2: derivational suffixes
		{name: "ational to ate", input: "configurational", expected: "configurate"},
		{name: "tional to tion", input: "nutritional", expected: "nutrition"},
		{name: "fulness to ful", input: "cheerfulness", expected: "cheerful"},
		{name: "ness removed", input: "darkness", expected: "dark"},
		{name: "ment removed", input: "deployment", expected: "deploy"},
		{name: "ation to ate", input: "configuration", expected: "configurate"},
		{name: "ously to ous", input: "dangerously", expected: "dangerous"},
		{name: "ively to ive", input: "effectively", expected: "effective"},
		{name: "ably to able", input: "comfortably", expected: "comfortable"},
		{name: "ally to al", input: "manually", expected: "manual"},
		{name: "izer to ize", input: "organizer", expected: "organize"},
		{name: "ingly to ing", input: "surprisingly", expected: "surprising"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stem(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestStem_ShortWordsUnchanged(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "single char", input: "a"},
		{name: "two chars", input: "go"},
		{name: "three chars", input: "run"},
		{name: "three chars (the)", input: "the"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.input, stem(tt.input), "words under 4 chars should be unchanged")
		})
	}
}

func TestStem_GuardMinLength(t *testing.T) {
	// The stem function must never reduce a word below 2 characters.
	// "ed" removal from a 4-char word like "abed" would leave "ab" (ok).
	// We test that it doesn't return a single-char result.
	result := stem("abed")
	assert.GreaterOrEqual(t, len(result), 2, "result must be at least 2 chars")
}

// ---------------------------------------------------------------------------
// Search integration tests — stemming recall
// ---------------------------------------------------------------------------

func TestSearch_StemRunningMatchesRun(t *testing.T) {
	idx := newSearchIndex()
	idx.Add(&Topic{
		ID:      "topic-run",
		Title:   "How to Run Commands",
		Content: "You can run any command from the terminal.",
	})

	results := idx.Search("running")
	require.NotEmpty(t, results, "searching 'running' should match topic containing 'run'")
	assert.Equal(t, "topic-run", results[0].Topic.ID)
}

func TestSearch_StemConfigurationsMatchesConfigure(t *testing.T) {
	idx := newSearchIndex()
	idx.Add(&Topic{
		ID:      "topic-configure",
		Title:   "Configure Your Application",
		Content: "Learn how to configure settings for your application.",
	})

	results := idx.Search("configurations")
	require.NotEmpty(t, results, "searching 'configurations' should match topic containing 'configure'")
	assert.Equal(t, "topic-configure", results[0].Topic.ID)
}

func TestSearch_StemPluralServersMatchesServer(t *testing.T) {
	idx := newSearchIndex()
	idx.Add(&Topic{
		ID:      "topic-server",
		Title:   "Server Management",
		Content: "Manage your server with these tools.",
	})

	results := idx.Search("servers")
	require.NotEmpty(t, results, "searching 'servers' should match topic containing 'server'")
	assert.Equal(t, "topic-server", results[0].Topic.ID)
}

func TestSearch_StemScoringLowerThanExact(t *testing.T) {
	idx := newSearchIndex()
	idx.Add(&Topic{
		ID:      "exact-match",
		Title:   "Running Guide",
		Content: "Guide to running applications.",
	})
	idx.Add(&Topic{
		ID:      "stem-match",
		Title:   "How to Run",
		Content: "Run your application.",
	})

	results := idx.Search("running")
	require.Len(t, results, 2, "should match both topics")

	// The topic containing the exact word "running" should score higher
	// than the one matched only via the stem "run" (all else being equal,
	// scoreExactWord > scoreStemWord).
	var exactScore, stemScore float64
	for _, r := range results {
		if r.Topic.ID == "exact-match" {
			exactScore = r.Score
		}
		if r.Topic.ID == "stem-match" {
			stemScore = r.Score
		}
	}
	assert.Greater(t, exactScore, stemScore,
		"exact word match should score higher than stem-only match")
}

func TestSearch_ExistingExactMatchUnaffected(t *testing.T) {
	// Ensure stemming doesn't break exact-match searches.
	idx := newSearchIndex()
	idx.Add(&Topic{
		ID:      "topic-deploy",
		Title:   "Deploy Guide",
		Content: "How to deploy your application step by step.",
	})

	results := idx.Search("deploy")
	require.NotEmpty(t, results)
	assert.Equal(t, "topic-deploy", results[0].Topic.ID)
}

func TestTokenize_IncludesStemmedVariants(t *testing.T) {
	words := tokenize("running configurations servers")

	// Should contain originals
	assert.Contains(t, words, "running")
	assert.Contains(t, words, "configurations")
	assert.Contains(t, words, "servers")

	// Should also contain stems
	assert.Contains(t, words, "runn")        // stem of running (ing removed)
	assert.Contains(t, words, "configurate") // stem of configurations (s->configuration->ation->ate)
	assert.Contains(t, words, "server")      // stem of servers (s removed)
}

// ---------------------------------------------------------------------------
// Benchmark
// ---------------------------------------------------------------------------

func BenchmarkStem(b *testing.B) {
	words := []string{
		"running", "configurations", "servers", "deployment", "testing",
		"addresses", "agreed", "configured", "operational", "cheerfulness",
		"darkness", "dangerously", "effectively", "comfortably", "manually",
		"organizer", "surprisingly", "configuration", "authentication",
		"authorisation", "networking", "monitoring", "scheduling", "routing",
		"migration", "encryption", "compression", "validation", "serialisation",
		"templating", "distributed", "federated", "graceful", "hybrid",
		"incremental", "advanced", "basic", "custom", "encrypted", "install",
		"configure", "deploy", "monitor", "debug", "authenticate", "authorise",
		"connect", "store", "analyse", "cache", "schedule", "route", "migrate",
		"restore", "help", "guide", "overview", "setup", "troubleshooting",
		"performance", "benchmark", "analysis", "documentation", "reference",
		"tutorial", "quickstart", "installation", "requirements", "dependencies",
		"modules", "packages", "services", "workers", "processes", "threads",
		"connections", "sessions", "transactions", "queries", "responses",
		"requests", "handlers", "middleware", "controllers", "models",
		"views", "templates", "layouts", "components", "widgets", "plugins",
		"extensions", "integrations", "providers", "factories", "builders",
		"adapters", "decorators", "observers", "listeners", "subscribers",
		"publishers", "dispatchers", "resolvers", "transformers", "formatters",
		"validators", "sanitizers", "parsers", "compilers", "interpreters",
	}

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		for _, w := range words {
			stem(w)
		}
	}
}
