// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	"cmp"
	"slices"
)

// topicGroup groups topics under a tag for the index page.
type topicGroup struct {
	Tag    string
	Topics []*Topic
}

// groupTopicsByTag groups topics by their first tag.
// Topics without tags are grouped under "other".
// Groups are sorted alphabetically by tag name.
func groupTopicsByTag(topics []*Topic) []topicGroup {
	// Sort topics first by Order then Title
	sorted := slices.Clone(topics)
	slices.SortFunc(sorted, func(a, b *Topic) int {
		if a.Order != b.Order {
			return cmp.Compare(a.Order, b.Order)
		}
		return cmp.Compare(a.Title, b.Title)
	})

	groups := make(map[string][]*Topic)
	for _, t := range sorted {
		tag := "other"
		if len(t.Tags) > 0 {
			tag = t.Tags[0]
		}
		groups[tag] = append(groups[tag], t)
	}

	var result []topicGroup
	for tag, topics := range groups {
		result = append(result, topicGroup{Tag: tag, Topics: topics})
	}
	slices.SortFunc(result, func(a, b topicGroup) int {
		return cmp.Compare(a.Tag, b.Tag)
	})

	return result
}
