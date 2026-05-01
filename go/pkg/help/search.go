package help

import (
	"cmp"
	"iter"
	"regexp"
	"slices"
	"strings"
	"unicode"
)

// Scoring weights for search result ranking.
const (
	scoreExactWord    = 1.0  // Exact word match in the index
	scorePrefixWord   = 0.5  // Prefix/partial word match
	scoreFuzzyWord    = 0.3  // Fuzzy (Levenshtein) match
	scoreStemWord     = 0.7  // Stemmed word match (between exact and prefix)
	scoreTitleBoost   = 10.0 // Query word appears in topic title
	scoreSectionBoost = 5.0  // Query word appears in section title
	scoreTagBoost     = 3.0  // Query word appears in topic tags
	scorePhraseBoost  = 8.0  // Exact phrase match in content
	scoreAllWords     = 2.0  // All query words present (multi-word bonus)
	fuzzyMaxDistance  = 2    // Maximum edit distance for fuzzy matching
)

// SearchResult represents a search match.
type SearchResult struct {
	Topic   *Topic
	Section *Section // nil if topic-level match
	Score   float64
	Snippet string // Context around match
}

// searchIndex provides full-text search.
type searchIndex struct {
	topics map[string]*Topic   // topicID -> Topic
	index  map[string][]string // word -> []topicID
}

// newSearchIndex creates a new empty search index.
func newSearchIndex() *searchIndex {
	return &searchIndex{
		topics: make(map[string]*Topic),
		index:  make(map[string][]string),
	}
}

// Add indexes a topic for searching.
func (i *searchIndex) Add(topic *Topic) {
	i.topics[topic.ID] = topic

	// Index title words with boost
	for _, word := range tokenize(topic.Title) {
		i.addToIndex(word, topic.ID)
	}

	// Index content words
	for _, word := range tokenize(topic.Content) {
		i.addToIndex(word, topic.ID)
	}

	// Index section titles and content
	for _, section := range topic.Sections {
		for _, word := range tokenize(section.Title) {
			i.addToIndex(word, topic.ID)
		}
		for _, word := range tokenize(section.Content) {
			i.addToIndex(word, topic.ID)
		}
	}

	// Index tags
	for _, tag := range topic.Tags {
		for _, word := range tokenize(tag) {
			i.addToIndex(word, topic.ID)
		}
	}
}

// addToIndex adds a word-to-topic mapping.
func (i *searchIndex) addToIndex(word, topicID string) {
	// Avoid duplicates
	if slices.Contains(i.index[word], topicID) {
		return
	}
	i.index[word] = append(i.index[word], topicID)
}

// Search finds topics matching the query. Supports:
//   - Single and multi-word keyword queries
//   - Quoted phrase search (e.g. `"rate limit"`)
//   - Fuzzy matching via Levenshtein distance for typo tolerance
//   - Prefix matching for partial words
func (i *searchIndex) Search(query string) []*SearchResult {
	// Extract quoted phrases before tokenising.
	phrases, stripped := extractPhrases(query)

	queryWords := tokenize(stripped)
	if len(queryWords) == 0 && len(phrases) == 0 {
		return nil
	}

	// Track scores per topic
	scores := make(map[string]float64)

	// Build set of stemmed query variants for stem-aware scoring.
	stemmedWords := make(map[string]bool)
	for _, word := range queryWords {
		if s := stem(word); s != word {
			stemmedWords[s] = true
		}
	}

	for _, word := range queryWords {
		isStem := stemmedWords[word]

		// Exact matches — score stems lower than raw words.
		if topicIDs, ok := i.index[word]; ok {
			sc := scoreExactWord
			if isStem {
				sc = scoreStemWord
			}
			for _, topicID := range topicIDs {
				scores[topicID] += sc
			}
		}

		// Prefix matches (partial word matching)
		for indexWord, topicIDs := range i.index {
			if strings.HasPrefix(indexWord, word) && indexWord != word {
				for _, topicID := range topicIDs {
					scores[topicID] += scorePrefixWord
				}
			}
		}

		// Fuzzy matches (Levenshtein distance)
		if len(word) >= 3 {
			for indexWord, topicIDs := range i.index {
				if indexWord == word {
					continue // Already scored as exact match
				}
				if strings.HasPrefix(indexWord, word) {
					continue // Already scored as prefix match
				}
				dist := levenshtein(word, indexWord)
				if dist > 0 && dist <= fuzzyMaxDistance {
					for _, topicID := range topicIDs {
						scores[topicID] += scoreFuzzyWord
					}
				}
			}
		}
	}

	// Pre-compile regexes for snippets
	var res []*regexp.Regexp
	for _, word := range queryWords {
		if len(word) >= 2 {
			if re, err := regexp.Compile("(?i)" + regexp.QuoteMeta(word)); err == nil {
				res = append(res, re)
			}
		}
	}
	// Also add phrase regexes for highlighting
	for _, phrase := range phrases {
		if re, err := regexp.Compile("(?i)" + regexp.QuoteMeta(phrase)); err == nil {
			res = append(res, re)
		}
	}

	// Phrase matching: boost topics that contain the exact phrase.
	for _, phrase := range phrases {
		phraseLower := strings.ToLower(phrase)
		for topicID, topic := range i.topics {
			var text strings.Builder
			text.WriteString(strings.ToLower(topic.Title + " " + topic.Content))
			for _, section := range topic.Sections {
				text.WriteString(" " + strings.ToLower(section.Title+" "+section.Content))
			}
			if strings.Contains(text.String(), phraseLower) {
				scores[topicID] += scorePhraseBoost
			}
		}
	}

	// Build results with title/section/tag boosts and snippet extraction
	var results []*SearchResult
	for topicID, score := range scores {
		topic := i.topics[topicID]
		if topic == nil {
			continue
		}

		// Title boost: if query words appear in title
		titleLower := strings.ToLower(topic.Title)
		titleMatchCount := 0
		for _, word := range queryWords {
			if strings.Contains(titleLower, word) {
				titleMatchCount++
			}
		}
		if titleMatchCount > 0 {
			score += scoreTitleBoost
		}

		// Tag boost: if query words match tags
		for _, tag := range topic.Tags {
			tagLower := strings.ToLower(tag)
			for _, word := range queryWords {
				if tagLower == word || strings.Contains(tagLower, word) {
					score += scoreTagBoost
					break
				}
			}
		}

		// Multi-word bonus: if all query words are present in the topic
		if len(queryWords) > 1 {
			allPresent := true
			fullText := strings.ToLower(topic.Title + " " + topic.Content)
			for _, word := range queryWords {
				if !strings.Contains(fullText, word) {
					allPresent = false
					break
				}
			}
			if allPresent {
				score += scoreAllWords
			}
		}

		// Find matching section and extract snippet
		section, snippet := i.findBestMatch(topic, queryWords, res)

		// Section title boost
		if section != nil {
			sectionTitleLower := strings.ToLower(section.Title)
			hasSectionTitleMatch := false
			for _, word := range queryWords {
				if strings.Contains(sectionTitleLower, word) {
					hasSectionTitleMatch = true
					break
				}
			}
			if hasSectionTitleMatch {
				score += scoreSectionBoost
			}
		}

		results = append(results, &SearchResult{
			Topic:   topic,
			Section: section,
			Score:   score,
			Snippet: snippet,
		})
	}

	// Sort by score (highest first)
	slices.SortFunc(results, func(a, b *SearchResult) int {
		if a.Score != b.Score {
			return cmp.Compare(b.Score, a.Score) // Reversed for highest first
		}
		return cmp.Compare(a.Topic.Title, b.Topic.Title)
	})

	return results
}

// extractPhrases pulls quoted substrings from the query and returns them
// alongside the remaining query text with quotes removed.
// For example: `hello "rate limit" world` returns
// phrases=["rate limit"], remaining="hello  world".
func extractPhrases(query string) (phrases []string, remaining string) {
	re := regexp.MustCompile(`"([^"]+)"`)
	matches := re.FindAllStringSubmatch(query, -1)
	for _, m := range matches {
		phrase := strings.TrimSpace(m[1])
		if phrase != "" {
			phrases = append(phrases, phrase)
		}
	}
	remaining = re.ReplaceAllString(query, "")
	return phrases, remaining
}

// levenshtein computes the edit distance between two strings.
// Used for fuzzy matching to tolerate typos in search queries.
func levenshtein(a, b string) int {
	aRunes := []rune(a)
	bRunes := []rune(b)
	aLen := len(aRunes)
	bLen := len(bRunes)

	if aLen == 0 {
		return bLen
	}
	if bLen == 0 {
		return aLen
	}

	// Use two rows instead of full matrix to save memory.
	prev := make([]int, bLen+1)
	curr := make([]int, bLen+1)

	for j := range bLen + 1 {
		prev[j] = j
	}

	for i := 1; i <= aLen; i++ {
		curr[0] = i
		for j := 1; j <= bLen; j++ {
			cost := 1
			if aRunes[i-1] == bRunes[j-1] {
				cost = 0
			}
			curr[j] = min(
				prev[j]+1,      // deletion
				curr[j-1]+1,    // insertion
				prev[j-1]+cost, // substitution
			)
		}
		prev, curr = curr, prev
	}

	return prev[bLen]
}

// findBestMatch finds the section with the best match and extracts a snippet.
func (i *searchIndex) findBestMatch(topic *Topic, queryWords []string, res []*regexp.Regexp) (*Section, string) {
	var bestSection *Section
	var bestSnippet string
	bestScore := 0

	// Check topic title
	titleScore := countMatches(topic.Title, queryWords)
	if titleScore > 0 {
		bestSnippet = extractSnippet(topic.Content, res)
	}

	// Check sections
	for idx := range topic.Sections {
		section := &topic.Sections[idx]
		sectionScore := countMatches(section.Title, queryWords)
		contentScore := countMatches(section.Content, queryWords)
		totalScore := sectionScore*2 + contentScore // Title matches worth more

		if totalScore > bestScore {
			bestScore = totalScore
			bestSection = section
			if contentScore > 0 {
				bestSnippet = extractSnippet(section.Content, res)
			} else {
				bestSnippet = extractSnippet(section.Content, nil)
			}
		}
	}

	// If no section matched, use topic content
	if bestSnippet == "" && topic.Content != "" {
		bestSnippet = extractSnippet(topic.Content, res)
	}

	return bestSection, bestSnippet
}

// tokenize splits text into lowercase words for indexing/searching.
// For each word, it also emits the stemmed variant (if different from the
// original) so the index contains both raw and stemmed forms.
func tokenize(text string) []string {
	return slices.Collect(Tokens(text))
}

// Tokens returns an iterator for lowercase words in text.
func Tokens(text string) iter.Seq[string] {
	return func(yield func(string) bool) {
		text = strings.ToLower(text)
		var word strings.Builder

		emit := func(w string) bool {
			if len(w) < 2 {
				return true
			}
			if !yield(w) {
				return false
			}
			if s := stem(w); s != w {
				if !yield(s) {
					return false
				}
			}
			return true
		}

		for _, r := range text {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				word.WriteRune(r)
			} else if word.Len() > 0 {
				if !emit(word.String()) {
					return
				}
				word.Reset()
			}
		}

		// Don't forget the last word
		if word.Len() > 0 {
			emit(word.String())
		}
	}
}

// countMatches counts how many query words appear in the text.
func countMatches(text string, queryWords []string) int {
	textLower := strings.ToLower(text)
	count := 0
	for _, word := range queryWords {
		if strings.Contains(textLower, word) {
			count++
		}
	}
	return count
}

// extractSnippet extracts a short snippet around the first match and highlights matches.
func extractSnippet(content string, res []*regexp.Regexp) string {
	if content == "" {
		return ""
	}

	const snippetLen = 150

	// If no regexes, return start of content without highlighting
	if len(res) == 0 {
		lines := strings.SplitSeq(content, "\n")
		for line := range lines {
			line = strings.TrimSpace(line)
			if line != "" && !strings.HasPrefix(line, "#") {
				runes := []rune(line)
				if len(runes) > snippetLen {
					return string(runes[:snippetLen]) + "..."
				}
				return line
			}
		}
		return ""
	}

	// Find first match position (byte-based)
	matchPos := -1
	for _, re := range res {
		loc := re.FindStringIndex(content)
		if loc != nil && (matchPos == -1 || loc[0] < matchPos) {
			matchPos = loc[0]
		}
	}

	// Convert to runes for safe slicing
	runes := []rune(content)
	runeLen := len(runes)

	var start, end int
	if matchPos == -1 {
		// No match found, use start of content
		start = 0
		end = min(snippetLen, runeLen)
	} else {
		// Convert byte position to rune position
		matchRunePos := len([]rune(content[:matchPos]))

		// Extract snippet around match (rune-based)
		start = max(matchRunePos-50, 0)

		end = min(start+snippetLen, runeLen)
	}

	snippet := string(runes[start:end])

	// Trim to word boundaries
	prefix := ""
	suffix := ""
	if start > 0 {
		if idx := strings.Index(snippet, " "); idx != -1 {
			snippet = snippet[idx+1:]
			prefix = "..."
		}
	}
	if end < runeLen {
		if idx := strings.LastIndex(snippet, " "); idx != -1 {
			snippet = snippet[:idx]
			suffix = "..."
		}
	}

	snippet = strings.TrimSpace(snippet)
	if snippet == "" {
		return ""
	}

	// Apply highlighting
	highlighted := highlight(snippet, res)

	return prefix + highlighted + suffix
}

// highlight wraps matches in **bold**.
func highlight(text string, res []*regexp.Regexp) string {
	if len(res) == 0 {
		return text
	}

	type match struct {
		start, end int
	}
	var matches []match

	for _, re := range res {
		indices := re.FindAllStringIndex(text, -1)
		for _, idx := range indices {
			matches = append(matches, match{idx[0], idx[1]})
		}
	}

	if len(matches) == 0 {
		return text
	}

	// Sort matches by start position
	slices.SortFunc(matches, func(a, b match) int {
		if a.start != b.start {
			return cmp.Compare(a.start, b.start)
		}
		return cmp.Compare(b.end, a.end)
	})

	// Merge overlapping or adjacent matches
	var merged []match
	if len(matches) > 0 {
		curr := matches[0]
		for i := 1; i < len(matches); i++ {
			if matches[i].start <= curr.end {
				if matches[i].end > curr.end {
					curr.end = matches[i].end
				}
			} else {
				merged = append(merged, curr)
				curr = matches[i]
			}
		}
		merged = append(merged, curr)
	}

	// Build highlighted string from back to front to avoid position shifts
	result := text
	for i := len(merged) - 1; i >= 0; i-- {
		m := merged[i]
		result = result[:m.end] + "**" + result[m.end:]
		result = result[:m.start] + "**" + result[m.start:]
	}

	return result
}
