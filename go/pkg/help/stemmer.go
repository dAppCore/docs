// SPDX-Licence-Identifier: EUPL-1.2
package help

import core "dappco.re/go"

// stem performs lightweight Porter-style suffix stripping on an English word.
// Words shorter than 4 characters are returned unchanged. The result is
// guaranteed to be at least 2 characters long.
//
// This is intentionally NOT the full Porter algorithm — it covers only the
// most impactful suffix rules for a help-catalog search context.
func stem(word string) string {
	if len(word) < 4 {
		return word
	}

	s := word

	// Step 1: plurals and verb inflections.
	s = stemInflectional(s)

	// Step 2: derivational suffixes (longest match first).
	s = stemDerivational(s)

	// Guard: result must be at least 2 characters.
	if len(s) < 2 {
		return word
	}

	return s
}

// stemInflectional handles plurals and -ed/-ing verb forms.
func stemInflectional(s string) string {
	switch {
	case core.HasSuffix(s, "sses"):
		return s[:len(s)-2] // -sses → -ss
	case core.HasSuffix(s, "ies"):
		return s[:len(s)-2] // -ies → -i
	case core.HasSuffix(s, "eed"):
		return s[:len(s)-1] // -eed → -ee
	case core.HasSuffix(s, "ing"):
		r := s[:len(s)-3]
		if len(r) >= 2 {
			return r
		}
	case core.HasSuffix(s, "ed"):
		r := s[:len(s)-2]
		if len(r) >= 2 {
			return r
		}
	case core.HasSuffix(s, "s") && !core.HasSuffix(s, "ss"):
		return s[:len(s)-1] // -s → "" (but not -ss)
	}
	return s
}

// stemDerivational strips common derivational suffixes.
// Ordered longest-first so we match the most specific rule.
func stemDerivational(s string) string {
	// Longest suffixes first (8+ chars).
	type rule struct {
		suffix      string
		replacement string
	}

	rules := []rule{
		{"fulness", "ful"}, // -fulness → -ful
		{"ational", "ate"}, // -ational → -ate
		{"tional", "tion"}, // -tional → -tion
		{"ously", "ous"},   // -ously → -ous
		{"ively", "ive"},   // -ively → -ive
		{"ingly", "ing"},   // -ingly → -ing
		{"ation", "ate"},   // -ation → -ate
		{"ness", ""},       // -ness → ""
		{"ment", ""},       // -ment → ""
		{"ably", "able"},   // -ably → -able
		{"ally", "al"},     // -ally → -al
		{"izer", "ize"},    // -izer → -ize
	}

	for _, r := range rules {
		if core.HasSuffix(s, r.suffix) {
			result := s[:len(s)-len(r.suffix)] + r.replacement
			if len(result) >= 2 {
				return result
			}
			return s // Guard: don't over-strip
		}
	}

	return s
}
