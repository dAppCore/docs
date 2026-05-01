// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	. "dappco.re/go"
)

func TestGroupTopicsByTag_Good(t *T) {
	topics := []*Topic{
		{ID: "a", Title: "Alpha", Tags: []string{"setup"}, Order: 2},
		{ID: "b", Title: "Beta", Tags: []string{"setup"}, Order: 1},
		{ID: "c", Title: "Gamma", Tags: []string{"advanced"}},
		{ID: "d", Title: "Delta"}, // no tags -> "other"
	}

	groups := groupTopicsByTag(topics)
	AssertLen(t, groups, 3)

	// Groups should be sorted alphabetically by tag
	AssertEqual(t, "advanced", groups[0].Tag)
	AssertEqual(t, "other", groups[1].Tag)
	AssertEqual(t, "setup", groups[2].Tag)

	// Within "setup", topics should be sorted by Order then Title
	setupGroup := groups[2]
	AssertLen(t, setupGroup.Topics, 2)
	AssertEqual(t, "Beta", setupGroup.Topics[0].Title)  // Order 1
	AssertEqual(t, "Alpha", setupGroup.Topics[1].Title) // Order 2
}
