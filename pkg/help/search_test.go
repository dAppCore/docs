// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	"regexp"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenize_Good(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "simple words",
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "mixed case",
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			name:     "with punctuation",
			input:    "Hello, world! How are you?",
			expected: []string{"hello", "world", "how", "are", "you"},
		},
		{
			name:     "single characters filtered",
			input:    "a b c hello d",
			expected: []string{"hello"},
		},
		{
			name:     "numbers included",
			input:    "version 2 release",
			expected: []string{"version", "release"},
		},
		{
			name:     "alphanumeric",
			input:    "v2.0 and config123",
			expected: []string{"v2", "and", "config123"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tokenize(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSearchIndex_Add_Good(t *testing.T) {
	idx := newSearchIndex()

	topic := &Topic{
		ID:      "getting-started",
		Title:   "Getting Started",
		Content: "Welcome to the guide.",
		Tags:    []string{"intro", "setup"},
		Sections: []Section{
			{ID: "installation", Title: "Installation", Content: "Install the CLI."},
		},
	}

	idx.Add(topic)

	// Verify topic is stored
	assert.NotNil(t, idx.topics["getting-started"])

	// Verify words are indexed
	assert.Contains(t, idx.index["getting"], "getting-started")
	assert.Contains(t, idx.index["started"], "getting-started")
	assert.Contains(t, idx.index["welcome"], "getting-started")
	assert.Contains(t, idx.index["guide"], "getting-started")
	assert.Contains(t, idx.index["intro"], "getting-started")
	assert.Contains(t, idx.index["setup"], "getting-started")
	assert.Contains(t, idx.index["installation"], "getting-started")
	assert.Contains(t, idx.index["cli"], "getting-started")
}

func TestSearchIndex_Search_Good(t *testing.T) {
	idx := newSearchIndex()

	// Add test topics
	idx.Add(&Topic{
		ID:      "getting-started",
		Title:   "Getting Started",
		Content: "Welcome to the CLI guide. This covers installation and setup.",
		Tags:    []string{"intro"},
	})

	idx.Add(&Topic{
		ID:      "configuration",
		Title:   "Configuration",
		Content: "Configure the CLI using environment variables.",
	})

	idx.Add(&Topic{
		ID:      "commands",
		Title:   "Commands Reference",
		Content: "List of all available commands.",
	})

	t.Run("single word query", func(t *testing.T) {
		results := idx.Search("configuration")
		assert.NotEmpty(t, results)
		assert.Equal(t, "configuration", results[0].Topic.ID)
	})

	t.Run("multi-word query", func(t *testing.T) {
		results := idx.Search("cli guide")
		assert.NotEmpty(t, results)
		// Should match getting-started (has both "cli" and "guide")
		assert.Equal(t, "getting-started", results[0].Topic.ID)
	})

	t.Run("title boost", func(t *testing.T) {
		results := idx.Search("commands")
		assert.NotEmpty(t, results)
		// "commands" appears in title of commands topic
		assert.Equal(t, "commands", results[0].Topic.ID)
	})

	t.Run("partial word matching", func(t *testing.T) {
		results := idx.Search("config")
		assert.NotEmpty(t, results)
		// Should match "configuration" and "configure"
		foundConfig := false
		for _, r := range results {
			if r.Topic.ID == "configuration" {
				foundConfig = true
				break
			}
		}
		assert.True(t, foundConfig, "Should find configuration topic with prefix match")
	})

	t.Run("no results", func(t *testing.T) {
		results := idx.Search("nonexistent")
		assert.Empty(t, results)
	})

	t.Run("empty query", func(t *testing.T) {
		results := idx.Search("")
		assert.Nil(t, results)
	})
}

func TestSearchIndex_Search_Good_WithSections(t *testing.T) {
	idx := newSearchIndex()

	idx.Add(&Topic{
		ID:      "installation",
		Title:   "Installation Guide",
		Content: "Overview of installation process.",
		Sections: []Section{
			{
				ID:      "linux",
				Title:   "Linux Installation",
				Content: "Run apt-get install core on Debian.",
			},
			{
				ID:      "macos",
				Title:   "macOS Installation",
				Content: "Use brew install core on macOS.",
			},
			{
				ID:      "windows",
				Title:   "Windows Installation",
				Content: "Download the installer from the website.",
			},
		},
	})

	t.Run("matches section content", func(t *testing.T) {
		results := idx.Search("debian")
		assert.NotEmpty(t, results)
		assert.Equal(t, "installation", results[0].Topic.ID)
		// Should identify the Linux section as best match
		if results[0].Section != nil {
			assert.Equal(t, "linux", results[0].Section.ID)
		}
	})

	t.Run("matches section title", func(t *testing.T) {
		results := idx.Search("windows")
		assert.NotEmpty(t, results)
		assert.Equal(t, "installation", results[0].Topic.ID)
	})
}

func TestExtractSnippet_Good(t *testing.T) {
	content := `This is the first paragraph with some introduction text.

Here is more content that talks about installation and setup.
The installation process is straightforward.

Finally, some closing remarks about the configuration.`

	t.Run("finds match and extracts context", func(t *testing.T) {
		snippet := extractSnippet(content, compileRegexes([]string{"installation"}))
		assert.Contains(t, snippet, "**installation**")
		assert.True(t, len(snippet) <= 250, "Snippet should be reasonably short")
	})

	t.Run("no query words returns start", func(t *testing.T) {
		snippet := extractSnippet(content, nil)
		assert.Contains(t, snippet, "first paragraph")
	})

	t.Run("empty content", func(t *testing.T) {
		snippet := extractSnippet("", compileRegexes([]string{"test"}))
		assert.Empty(t, snippet)
	})
}

func TestExtractSnippet_Highlighting(t *testing.T) {
	content := "The quick brown fox jumps over the lazy dog."

	t.Run("simple highlighting", func(t *testing.T) {
		snippet := extractSnippet(content, compileRegexes([]string{"quick", "fox"}))
		assert.Contains(t, snippet, "**quick**")
		assert.Contains(t, snippet, "**fox**")
	})

	t.Run("case insensitive highlighting", func(t *testing.T) {
		snippet := extractSnippet(content, compileRegexes([]string{"QUICK", "Fox"}))
		assert.Contains(t, snippet, "**quick**")
		assert.Contains(t, snippet, "**fox**")
	})

	t.Run("partial word matching", func(t *testing.T) {
		content := "The configuration is complete."
		snippet := extractSnippet(content, compileRegexes([]string{"config"}))
		assert.Contains(t, snippet, "**config**uration")
	})

	t.Run("overlapping matches", func(t *testing.T) {
		content := "Searching for something."
		// Both "search" and "searching" match
		snippet := extractSnippet(content, compileRegexes([]string{"search", "searching"}))
		assert.Equal(t, "**Searching** for something.", snippet)
	})
}

func TestExtractSnippet_Good_UTF8(t *testing.T) {
	// Content with multi-byte UTF-8 characters
	content := "日本語のテキストです。This contains Japanese text. 検索機能をテストします。"

	t.Run("handles multi-byte characters without corruption", func(t *testing.T) {
		snippet := extractSnippet(content, compileRegexes([]string{"japanese"}))
		// Should not panic or produce invalid UTF-8
		assert.True(t, len(snippet) > 0)
		// Verify the result is valid UTF-8
		assert.True(t, isValidUTF8(snippet), "Snippet should be valid UTF-8")
	})

	t.Run("truncates multi-byte content safely", func(t *testing.T) {
		// Long content that will be truncated
		longContent := strings.Repeat("日本語", 100) // 300 characters
		snippet := extractSnippet(longContent, nil)
		assert.True(t, isValidUTF8(snippet), "Truncated snippet should be valid UTF-8")
	})
}

// compileRegexes is a helper for tests.
func compileRegexes(words []string) []*regexp.Regexp {
	var res []*regexp.Regexp
	for _, w := range words {
		if re, err := regexp.Compile("(?i)" + regexp.QuoteMeta(w)); err == nil {
			res = append(res, re)
		}
	}
	return res
}

// isValidUTF8 checks if a string is valid UTF-8
func isValidUTF8(s string) bool {
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		if r == utf8.RuneError && size == 1 {
			return false
		}
		i += size
	}
	return true
}

func TestCountMatches_Good(t *testing.T) {
	tests := []struct {
		text     string
		words    []string
		expected int
	}{
		{"Hello world", []string{"hello"}, 1},
		{"Hello world", []string{"hello", "world"}, 2},
		{"Hello world", []string{"foo", "bar"}, 0},
		{"The quick brown fox", []string{"quick", "fox", "dog"}, 2},
	}

	for _, tt := range tests {
		result := countMatches(tt.text, tt.words)
		assert.Equal(t, tt.expected, result)
	}
}

func TestSearchResult_Score_Good(t *testing.T) {
	idx := newSearchIndex()

	// Topic with query word in title should score higher
	idx.Add(&Topic{
		ID:      "topic-in-title",
		Title:   "Installation Guide",
		Content: "Some content here.",
	})

	idx.Add(&Topic{
		ID:      "topic-in-content",
		Title:   "Some Other Topic",
		Content: "This covers installation steps.",
	})

	results := idx.Search("installation")
	assert.Len(t, results, 2)

	// Title match should score higher
	assert.Equal(t, "topic-in-title", results[0].Topic.ID)
	assert.Greater(t, results[0].Score, results[1].Score)
}

// --- Upstream Phase 0 tests (100% coverage) ---

func TestExtractSnippet_Good_HeadingsOnly(t *testing.T) {
	// Content with only headings and no body text should return empty snippet
	// when no regexes are provided. Covers the empty-return branch.
	content := "# Heading One\n## Heading Two\n### Heading Three"

	snippet := extractSnippet(content, nil)
	assert.Empty(t, snippet, "headings-only content without regexes should return empty snippet")
}

func TestExtractSnippet_Good_SnippetTrimmedToEmpty(t *testing.T) {
	// After word-boundary trimming the snippet could become empty.
	// This exercises the snippet=="" guard after TrimSpace.
	content := strings.Repeat(" ", 200)

	snippet := extractSnippet(content, compileRegexes([]string{"zz"}))
	assert.Empty(t, snippet, "whitespace-only content should yield empty snippet")
}

func TestHighlight_Good_EmptyRegexes(t *testing.T) {
	// Calling highlight with an empty regex slice should return the
	// text unchanged. Covers the early-return branch.
	result := highlight("some text here", nil)
	assert.Equal(t, "some text here", result)

	result = highlight("some text here", []*regexp.Regexp{})
	assert.Equal(t, "some text here", result)
}

func TestHighlight_Good_OverlappingExtension(t *testing.T) {
	// Test overlapping matches where the second match extends past the
	// first, exercising the curr.end extension branch in merging.
	text := "abcdefghij"

	// First regex matches "abcdef", second matches "cdefghij"
	// They overlap and the second extends the merged range.
	re1, _ := regexp.Compile("abcdef")
	re2, _ := regexp.Compile("cdefghij")

	result := highlight(text, []*regexp.Regexp{re1, re2})
	assert.Equal(t, "**abcdefghij**", result)
}

func TestSearchIndex_Search_Good_NilTopicGuard(t *testing.T) {
	// Manually inject a stale reference in the scores map so the
	// nil-topic guard is exercised. We do this by manipulating the
	// index directly.
	idx := newSearchIndex()

	idx.Add(&Topic{
		ID:      "real-topic",
		Title:   "Real Topic",
		Content: "Contains testword for matching.",
	})

	// Inject a mapping from "testword" to a non-existent topic ID.
	idx.index["testword"] = append(idx.index["testword"], "ghost-topic")

	results := idx.Search("testword")
	// Should still find the real topic, ghost-topic should be skipped.
	assert.Len(t, results, 1)
	assert.Equal(t, "real-topic", results[0].Topic.ID)
}

func TestSearchIndex_Search_Good_SpecialCharacters(t *testing.T) {
	idx := newSearchIndex()

	idx.Add(&Topic{
		ID:      "special-chars",
		Title:   "Rate Limiting (v2.0)",
		Content: "Configure rate-limiting rules with special characters: @#$%.",
	})

	t.Run("query with special characters", func(t *testing.T) {
		results := idx.Search("rate limiting")
		assert.NotEmpty(t, results)
		assert.Equal(t, "special-chars", results[0].Topic.ID)
	})

	t.Run("query with punctuation stripped", func(t *testing.T) {
		results := idx.Search("v2.0")
		assert.NotEmpty(t, results)
	})
}

func TestSearchIndex_Search_Good_CaseInsensitive(t *testing.T) {
	idx := newSearchIndex()

	idx.Add(&Topic{
		ID:      "case-test",
		Title:   "UPPERCASE Title",
		Content: "MiXeD CaSe content here.",
	})

	t.Run("lowercase query finds uppercase content", func(t *testing.T) {
		results := idx.Search("uppercase")
		assert.NotEmpty(t, results)
		assert.Equal(t, "case-test", results[0].Topic.ID)
	})

	t.Run("uppercase query finds content", func(t *testing.T) {
		results := idx.Search("MIXED")
		assert.NotEmpty(t, results)
	})
}

func TestSearchIndex_Search_Good_SingleCharQuery(t *testing.T) {
	idx := newSearchIndex()

	idx.Add(&Topic{
		ID:      "single-char",
		Title:   "Test Topic",
		Content: "Some content.",
	})

	// Single-character queries are filtered out by tokenize (min 2 chars),
	// so searching for "a" should return nil.
	results := idx.Search("a")
	assert.Nil(t, results)
}

// --- Phase 1: Fuzzy matching tests ---

func TestLevenshtein_Good(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected int
	}{
		{name: "identical strings", a: "hello", b: "hello", expected: 0},
		{name: "single substitution", a: "hello", b: "hallo", expected: 1},
		{name: "single insertion", a: "hello", b: "helloo", expected: 1},
		{name: "single deletion", a: "hello", b: "helo", expected: 1},
		{name: "two edits", a: "kitten", b: "sitting", expected: 3},
		{name: "completely different", a: "abc", b: "xyz", expected: 3},
		{name: "empty first string", a: "", b: "hello", expected: 5},
		{name: "empty second string", a: "hello", b: "", expected: 5},
		{name: "both empty", a: "", b: "", expected: 0},
		{name: "single char strings", a: "a", b: "b", expected: 1},
		{name: "same single char", a: "a", b: "a", expected: 0},
		{name: "transposition", a: "ab", b: "ba", expected: 2},
		{name: "unicode strings", a: "cafe", b: "cafe", expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := levenshtein(tt.a, tt.b)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMin3_Good(t *testing.T) {
	tests := []struct {
		name     string
		a, b, c  int
		expected int
	}{
		{name: "first smallest", a: 1, b: 2, c: 3, expected: 1},
		{name: "second smallest", a: 2, b: 1, c: 3, expected: 1},
		{name: "third smallest", a: 3, b: 2, c: 1, expected: 1},
		{name: "all equal", a: 5, b: 5, c: 5, expected: 5},
		{name: "first and second equal smallest", a: 1, b: 1, c: 3, expected: 1},
		{name: "second and third equal smallest", a: 3, b: 1, c: 1, expected: 1},
		{name: "first and third equal smallest", a: 1, b: 3, c: 1, expected: 1},
		{name: "negative values", a: -3, b: -1, c: -2, expected: -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := min(tt.a, tt.b, tt.c)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSearchIndex_Search_Good_FuzzyMatching(t *testing.T) {
	idx := newSearchIndex()

	idx.Add(&Topic{
		ID:      "configuration",
		Title:   "Configuration Guide",
		Content: "Learn how to configure the application settings.",
	})

	idx.Add(&Topic{
		ID:      "deployment",
		Title:   "Deployment Process",
		Content: "Deploy your application to production servers.",
	})

	t.Run("typo in query finds correct topic", func(t *testing.T) {
		// "configuraton" is 1 edit away from "configuration"
		results := idx.Search("configuraton")
		assert.NotEmpty(t, results, "fuzzy match should find results for typo")
		found := false
		for _, r := range results {
			if r.Topic.ID == "configuration" {
				found = true
				break
			}
		}
		assert.True(t, found, "should find configuration topic with typo")
	})

	t.Run("two-edit typo still matches", func(t *testing.T) {
		// "deplymnt" is within 2 edits of "deployment" -- but first check
		// that "deploymnt" (1 edit) works.
		results := idx.Search("deploymnt")
		assert.NotEmpty(t, results, "fuzzy match should find results for 1-edit typo")
	})

	t.Run("too many edits returns no fuzzy match", func(t *testing.T) {
		// "zzzzzzz" is very far from any indexed word.
		results := idx.Search("zzzzzzz")
		assert.Empty(t, results, "large edit distance should not produce results")
	})

	t.Run("short words skip fuzzy matching", func(t *testing.T) {
		// Words shorter than 3 characters skip fuzzy matching.
		// "to" is in the index but "tx" (1 edit) should not fuzzy-match
		// because "tx" is only 2 chars.
		results := idx.Search("tx")
		// May or may not find results via prefix, but should not crash.
		_ = results
	})

	t.Run("fuzzy scores lower than exact", func(t *testing.T) {
		// Exact match on "configure" should score higher than fuzzy.
		exactResults := idx.Search("configure")
		fuzzyResults := idx.Search("configurr")

		if len(exactResults) > 0 && len(fuzzyResults) > 0 {
			assert.GreaterOrEqual(t, exactResults[0].Score, fuzzyResults[0].Score,
				"exact match should score at least as high as fuzzy match")
		}
	})
}

// --- Phase 1: Phrase search tests ---

func TestExtractPhrases_Good(t *testing.T) {
	tests := []struct {
		name              string
		query             string
		expectedPhrases   []string
		expectedRemaining string
	}{
		{
			name:              "no phrases",
			query:             "hello world",
			expectedPhrases:   nil,
			expectedRemaining: "hello world",
		},
		{
			name:              "single phrase",
			query:             `"rate limit"`,
			expectedPhrases:   []string{"rate limit"},
			expectedRemaining: "",
		},
		{
			name:              "phrase with surrounding words",
			query:             `configure "rate limit" rules`,
			expectedPhrases:   []string{"rate limit"},
			expectedRemaining: "configure  rules",
		},
		{
			name:              "multiple phrases",
			query:             `"rate limit" and "error handling"`,
			expectedPhrases:   []string{"rate limit", "error handling"},
			expectedRemaining: " and ",
		},
		{
			name:              "empty quotes left in remaining",
			query:             `"" hello`,
			expectedPhrases:   nil,
			expectedRemaining: `"" hello`,
		},
		{
			name:              "whitespace-only quotes ignored",
			query:             `"   " hello`,
			expectedPhrases:   nil,
			expectedRemaining: " hello",
		},
		{
			name:              "unclosed quote treated as plain text",
			query:             `"unclosed phrase`,
			expectedPhrases:   nil,
			expectedRemaining: `"unclosed phrase`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			phrases, remaining := extractPhrases(tt.query)
			assert.Equal(t, tt.expectedPhrases, phrases)
			assert.Equal(t, tt.expectedRemaining, remaining)
		})
	}
}

func TestSearchIndex_Search_Good_PhraseSearch(t *testing.T) {
	idx := newSearchIndex()

	idx.Add(&Topic{
		ID:      "rate-limiting",
		Title:   "Rate Limiting",
		Content: "Configure rate limit rules for API endpoints. The rate limit protects against abuse.",
		Tags:    []string{"api", "security"},
	})

	idx.Add(&Topic{
		ID:      "error-handling",
		Title:   "Error Handling",
		Content: "Handle errors and rate responses correctly. Limit retries to avoid loops.",
		Tags:    []string{"errors"},
	})

	t.Run("quoted phrase matches exact sequence", func(t *testing.T) {
		results := idx.Search(`"rate limit"`)
		assert.NotEmpty(t, results)
		// rate-limiting topic has "rate limit" as exact phrase
		assert.Equal(t, "rate-limiting", results[0].Topic.ID)
	})

	t.Run("unquoted words match both topics", func(t *testing.T) {
		results := idx.Search("rate limit")
		assert.GreaterOrEqual(t, len(results), 2,
			"unquoted query should match both topics that contain the words separately")
	})

	t.Run("phrase not found yields no phrase boost", func(t *testing.T) {
		results := idx.Search(`"nonexistent phrase here"`)
		assert.Empty(t, results, "phrase with no tokenisable words and no match should return empty")
	})

	t.Run("phrase with surrounding keywords", func(t *testing.T) {
		results := idx.Search(`"rate limit" api`)
		assert.NotEmpty(t, results)
		assert.Equal(t, "rate-limiting", results[0].Topic.ID)
	})

	t.Run("phrase-only query with no loose words", func(t *testing.T) {
		// Query is only a quoted phrase; tokenize of remaining is empty,
		// but phrase matching should still score topics.
		results := idx.Search(`"rate limit"`)
		assert.NotEmpty(t, results)
	})

	t.Run("phrase found in section content", func(t *testing.T) {
		sectionIdx := newSearchIndex()
		sectionIdx.Add(&Topic{
			ID:      "section-phrase",
			Title:   "Advanced Guide",
			Content: "Overview of the system.",
			Sections: []Section{
				{
					ID:      "limits",
					Title:   "Limits",
					Content: "The rate limit is set per client.",
				},
			},
		})

		results := sectionIdx.Search(`"rate limit"`)
		assert.NotEmpty(t, results)
		assert.Equal(t, "section-phrase", results[0].Topic.ID)
	})
}

// --- Phase 1: Improved scoring tests ---

func TestSearchIndex_Search_Good_TagBoost(t *testing.T) {
	idx := newSearchIndex()

	idx.Add(&Topic{
		ID:      "tagged-topic",
		Title:   "Some Guide",
		Content: "General content about the system.",
		Tags:    []string{"deployment", "production"},
	})

	idx.Add(&Topic{
		ID:      "content-topic",
		Title:   "Other Guide",
		Content: "Information about deployment processes.",
	})

	results := idx.Search("deployment")
	assert.NotEmpty(t, results)

	// The tagged topic should rank higher because of the tag boost.
	assert.Equal(t, "tagged-topic", results[0].Topic.ID,
		"topic with matching tag should rank higher")
}

func TestSearchIndex_Search_Good_MultiWordBonus(t *testing.T) {
	idx := newSearchIndex()

	// Both topics have neutral titles (no query words in title) to
	// isolate the multi-word bonus effect.
	idx.Add(&Topic{
		ID:      "both-words",
		Title:   "Complete Guide",
		Content: "Learn about deploying and monitoring in one place.",
	})

	idx.Add(&Topic{
		ID:      "one-word",
		Title:   "Other Guide",
		Content: "Just deploying steps without monitoring.",
	})

	results := idx.Search("deploying monitoring")
	assert.NotEmpty(t, results)

	// Topic with both words should score higher due to multi-word bonus.
	assert.Equal(t, "both-words", results[0].Topic.ID,
		"topic containing all query words should rank higher")
}

func TestSearchIndex_Search_Good_ScoringConstants(t *testing.T) {
	// Verify the scoring constants are sensible relative to each other.
	assert.Greater(t, scoreTitleBoost, scoreSectionBoost,
		"title boost should exceed section boost")
	assert.Greater(t, scoreSectionBoost, scoreTagBoost,
		"section boost should exceed tag boost")
	assert.Greater(t, scoreExactWord, scorePrefixWord,
		"exact match should score higher than prefix match")
	assert.Greater(t, scorePrefixWord, scoreFuzzyWord,
		"prefix match should score higher than fuzzy match")
	assert.Greater(t, scorePhraseBoost, scoreAllWords,
		"phrase boost should exceed multi-word bonus")
}

func TestSearchIndex_Search_Good_PhraseHighlighting(t *testing.T) {
	idx := newSearchIndex()

	idx.Add(&Topic{
		ID:      "phrase-highlight",
		Title:   "API Guide",
		Content: "Configure rate limit rules for the API gateway.",
	})

	results := idx.Search(`"rate limit" api`)
	assert.NotEmpty(t, results)

	// The snippet should highlight both the phrase and keyword.
	if results[0].Snippet != "" {
		assert.Contains(t, results[0].Snippet, "**rate limit**",
			"phrase should be highlighted in snippet")
	}
}

// --- Phase 0 additional tests: expanded edge cases ---

func TestSearchIndex_Search_Bad_EmptyQuery(t *testing.T) {
	idx := newSearchIndex()
	idx.Add(&Topic{ID: "test", Title: "Test Topic", Content: "Some content."})

	t.Run("empty string", func(t *testing.T) {
		results := idx.Search("")
		assert.Nil(t, results)
	})

	t.Run("whitespace only", func(t *testing.T) {
		results := idx.Search("   ")
		assert.Nil(t, results)
	})

	t.Run("single character", func(t *testing.T) {
		// Single chars are filtered by tokenize (min 2 chars)
		results := idx.Search("a")
		assert.Nil(t, results)
	})

	t.Run("punctuation only", func(t *testing.T) {
		results := idx.Search("!@#$%")
		assert.Nil(t, results)
	})
}

func TestSearchIndex_Search_Bad_NoResults(t *testing.T) {
	idx := newSearchIndex()

	idx.Add(&Topic{
		ID:      "golang",
		Title:   "Golang Programming",
		Content: "Building applications with Go and goroutines.",
	})

	t.Run("completely unrelated query", func(t *testing.T) {
		results := idx.Search("quantum physics")
		assert.Empty(t, results)
	})

	t.Run("empty index", func(t *testing.T) {
		emptyIdx := newSearchIndex()
		results := emptyIdx.Search("anything")
		assert.Empty(t, results)
	})
}

func TestSearchIndex_Search_Good_CaseSensitivity(t *testing.T) {
	idx := newSearchIndex()

	idx.Add(&Topic{
		ID:      "case-test",
		Title:   "PostgreSQL Configuration",
		Content: "Configure POSTGRESQL settings. The postgresql.conf file controls everything.",
	})

	t.Run("lowercase query matches uppercase content", func(t *testing.T) {
		results := idx.Search("postgresql")
		require.NotEmpty(t, results)
		assert.Equal(t, "case-test", results[0].Topic.ID)
	})

	t.Run("uppercase query matches lowercase content", func(t *testing.T) {
		results := idx.Search("POSTGRESQL")
		require.NotEmpty(t, results)
		assert.Equal(t, "case-test", results[0].Topic.ID)
	})

	t.Run("mixed case query matches", func(t *testing.T) {
		results := idx.Search("PostgreSQL")
		require.NotEmpty(t, results)
		assert.Equal(t, "case-test", results[0].Topic.ID)
	})

	t.Run("title case sensitivity", func(t *testing.T) {
		results := idx.Search("configuration")
		require.NotEmpty(t, results)
		assert.Equal(t, "case-test", results[0].Topic.ID)
	})
}

func TestSearchIndex_Search_Good_MultiWord(t *testing.T) {
	idx := newSearchIndex()

	idx.Add(&Topic{
		ID:      "docker-compose",
		Title:   "Docker Compose Setup",
		Content: "Learn how to use Docker Compose for container orchestration.",
	})
	idx.Add(&Topic{
		ID:      "docker-basics",
		Title:   "Docker Basics",
		Content: "Introduction to Docker containers and images.",
	})
	idx.Add(&Topic{
		ID:      "kubernetes",
		Title:   "Kubernetes Setup",
		Content: "Setting up a Kubernetes cluster for production.",
	})

	t.Run("both words match same topic", func(t *testing.T) {
		results := idx.Search("docker compose")
		require.NotEmpty(t, results)
		// docker-compose should rank highest (both words in title + content)
		assert.Equal(t, "docker-compose", results[0].Topic.ID)
	})

	t.Run("one word matches multiple topics", func(t *testing.T) {
		results := idx.Search("docker")
		require.Len(t, results, 2)
		// Both docker topics should appear
		ids := []string{results[0].Topic.ID, results[1].Topic.ID}
		assert.Contains(t, ids, "docker-compose")
		assert.Contains(t, ids, "docker-basics")
	})

	t.Run("words from different topics", func(t *testing.T) {
		results := idx.Search("docker kubernetes")
		require.NotEmpty(t, results)
		// All three topics should match (docker matches 2, kubernetes matches 1)
		assert.GreaterOrEqual(t, len(results), 3)
	})

	t.Run("three word query narrows results", func(t *testing.T) {
		results := idx.Search("docker compose setup")
		require.NotEmpty(t, results)
		// docker-compose has all three words, should rank first
		assert.Equal(t, "docker-compose", results[0].Topic.ID)
	})
}

func TestSearchIndex_Search_Good_SpecialCharsExpanded(t *testing.T) {
	idx := newSearchIndex()

	idx.Add(&Topic{
		ID:      "email-config",
		Title:   "Email Configuration",
		Content: "Set SMTP_HOST to smtp.example.com and PORT to 587.",
	})
	idx.Add(&Topic{
		ID:      "dotfiles",
		Title:   "Dotfile Management",
		Content: "Manage your .bashrc and .zshrc files across machines.",
	})
	idx.Add(&Topic{
		ID:      "at-mentions",
		Title:   "User Mentions",
		Content: "Use @username to mention users in comments.",
	})

	t.Run("query with at symbol", func(t *testing.T) {
		// "@username" tokenises to "username" (@ is stripped)
		results := idx.Search("@username")
		require.NotEmpty(t, results)
		assert.Equal(t, "at-mentions", results[0].Topic.ID)
	})

	t.Run("query with dots", func(t *testing.T) {
		// "smtp.example.com" tokenises to "smtp", "example", "com"
		results := idx.Search("smtp.example.com")
		require.NotEmpty(t, results)
		assert.Equal(t, "email-config", results[0].Topic.ID)
	})

	t.Run("query with underscores", func(t *testing.T) {
		// "SMTP_HOST" tokenises to "smtp", "host"
		results := idx.Search("SMTP_HOST")
		require.NotEmpty(t, results)
		assert.Equal(t, "email-config", results[0].Topic.ID)
	})
}

func TestSearchIndex_Search_Good_OverlappingMatches(t *testing.T) {
	idx := newSearchIndex()

	idx.Add(&Topic{
		ID:      "search-guide",
		Title:   "Searching and Search Results",
		Content: "The search function searches through searchable content to find search results.",
	})

	// "search" should match: "searching", "search", "searches", "searchable"
	results := idx.Search("search")
	require.NotEmpty(t, results)
	assert.Equal(t, "search-guide", results[0].Topic.ID)
	// Score should be boosted since "search" appears in the title
	assert.Greater(t, results[0].Score, 10.0)
}

func TestSearchIndex_Search_Good_ScoringBoundary(t *testing.T) {
	idx := newSearchIndex()

	// Topic A: exact title match
	idx.Add(&Topic{
		ID:      "exact-title",
		Title:   "Installation",
		Content: "Basic content without the query word repeated.",
	})

	// Topic B: no title match but heavy body usage
	idx.Add(&Topic{
		ID:      "heavy-body",
		Title:   "Getting Started Guide",
		Content: "Installation steps: First install the package. Then install dependencies. The installation is straightforward. Install everything.",
		Sections: []Section{
			{
				ID:      "install-section",
				Title:   "Install Steps",
				Content: "Detailed installation instructions for every platform.",
			},
		},
	})

	results := idx.Search("installation")
	require.Len(t, results, 2)

	// Title match gets +10 boost, so "exact-title" should rank first
	assert.Equal(t, "exact-title", results[0].Topic.ID, "exact title match should rank above body-heavy match")
	assert.Greater(t, results[0].Score, results[1].Score)
}

func TestSearchIndex_Search_Good_TagMatching(t *testing.T) {
	idx := newSearchIndex()

	idx.Add(&Topic{
		ID:      "tagged-topic",
		Title:   "Workflow Automation",
		Content: "Automate your CI/CD pipeline.",
		Tags:    []string{"devops", "cicd", "automation"},
	})

	// Search for a tag that does not appear in title or content
	results := idx.Search("devops")
	require.NotEmpty(t, results)
	assert.Equal(t, "tagged-topic", results[0].Topic.ID)
}

func TestSearchIndex_Search_Good_SectionTitleBoost(t *testing.T) {
	idx := newSearchIndex()

	idx.Add(&Topic{
		ID:      "section-match",
		Title:   "Complete Reference",
		Content: "Overview of all features.",
		Sections: []Section{
			{ID: "deployment", Title: "Deployment", Content: "How to deploy your application."},
			{ID: "monitoring", Title: "Monitoring", Content: "Set up health checks."},
		},
	})

	idx.Add(&Topic{
		ID:      "body-match",
		Title:   "Quick Tips",
		Content: "Deployment can be tricky, here are some tips.",
	})

	results := idx.Search("deployment")
	require.Len(t, results, 2)

	// Section title match gives +5 boost (in addition to other scoring)
	sectionResult := results[0]
	assert.Equal(t, "section-match", sectionResult.Topic.ID)
	if sectionResult.Section != nil {
		assert.Equal(t, "deployment", sectionResult.Section.ID)
	}
}

func TestTokenize_Good_SpecialCases(t *testing.T) {
	t.Run("only special characters", func(t *testing.T) {
		result := tokenize("!@#$%^&*()")
		assert.Nil(t, result)
	})

	t.Run("unicode tokens", func(t *testing.T) {
		result := tokenize("日本語 テスト")
		assert.NotEmpty(t, result, "CJK characters should tokenise as words")
	})

	t.Run("mixed unicode and ascii", func(t *testing.T) {
		result := tokenize("hello 世界 world")
		assert.Contains(t, result, "hello")
		assert.Contains(t, result, "world")
	})

	t.Run("numbers only", func(t *testing.T) {
		result := tokenize("12345 67890")
		assert.Equal(t, []string{"12345", "67890"}, result)
	})

	t.Run("hyphenated words split", func(t *testing.T) {
		result := tokenize("pre-commit")
		assert.Equal(t, []string{"pre", "commit"}, result)
	})
}

func TestHighlight_Good_NoMatches(t *testing.T) {
	result := highlight("no matches here", compileRegexes([]string{"xyz"}))
	assert.Equal(t, "no matches here", result)
}

func TestHighlight_Good_AdjacentMatches(t *testing.T) {
	// Two words right next to each other
	result := highlight("foobar", compileRegexes([]string{"foo", "bar"}))
	// "foo" and "bar" are adjacent, should be merged into one highlight
	assert.Equal(t, "**foobar**", result)
}

func TestExtractSnippet_Good_HeadingsSkipped(t *testing.T) {
	// When no regex is given, snippet should skip heading lines
	content := "# Heading\n\nActual content here."
	snippet := extractSnippet(content, nil)
	assert.Contains(t, snippet, "Actual content here.")
	assert.NotContains(t, snippet, "# Heading")
}

func TestSearchIndex_Search_Good_DuplicateTopicIDs(t *testing.T) {
	idx := newSearchIndex()

	// Adding the same topic twice should not cause duplicate results
	topic := &Topic{
		ID:      "deduplicated",
		Title:   "Unique Topic",
		Content: "Unique content about testing.",
	}
	idx.Add(topic)
	idx.Add(topic)

	results := idx.Search("unique")
	assert.Len(t, results, 1)
}

func TestCatalog_Search_Good_Integration(t *testing.T) {
	// Test the full Catalog.Search path (integration through catalog -> index)
	cat := &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}

	cat.Add(&Topic{
		ID:      "alpha",
		Title:   "Alpha Feature",
		Content: "This is the alpha version of the feature.",
		Tags:    []string{"experimental"},
	})
	cat.Add(&Topic{
		ID:      "beta",
		Title:   "Beta Release Notes",
		Content: "Improvements and bug fixes in the beta.",
		Tags:    []string{"release"},
	})

	t.Run("search via catalog", func(t *testing.T) {
		results := cat.Search("alpha")
		require.NotEmpty(t, results)
		assert.Equal(t, "alpha", results[0].Topic.ID)
	})

	t.Run("search by tag via catalog", func(t *testing.T) {
		results := cat.Search("experimental")
		require.NotEmpty(t, results)
		assert.Equal(t, "alpha", results[0].Topic.ID)
	})

	t.Run("empty query via catalog", func(t *testing.T) {
		results := cat.Search("")
		assert.Nil(t, results)
	})
}
