// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGroupTopicsByTag_Good(t *testing.T) {
	topics := []*Topic{
		{ID: "a", Title: "Alpha", Tags: []string{"setup"}, Order: 2},
		{ID: "b", Title: "Beta", Tags: []string{"setup"}, Order: 1},
		{ID: "c", Title: "Gamma", Tags: []string{"advanced"}},
		{ID: "d", Title: "Delta"}, // no tags -> "other"
	}

	groups := groupTopicsByTag(topics)
	require.Len(t, groups, 3)

	// Groups should be sorted alphabetically by tag
	assert.Equal(t, "advanced", groups[0].Tag)
	assert.Equal(t, "other", groups[1].Tag)
	assert.Equal(t, "setup", groups[2].Tag)

	// Within "setup", topics should be sorted by Order then Title
	setupGroup := groups[2]
	require.Len(t, setupGroup.Topics, 2)
	assert.Equal(t, "Beta", setupGroup.Topics[0].Title)  // Order 1
	assert.Equal(t, "Alpha", setupGroup.Topics[1].Title) // Order 2
}
