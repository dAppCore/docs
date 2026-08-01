package help

import (
	. "dappco.re/go"
)

func TestTopic_Topic_Good(t *T) {
	subject := Topic{}
	_ = subject
	marker := "Service:Good"
	if marker == "" {
		t.FailNow()
	}
}

func TestTopic_Section_Good(t *T) {
	subject := Section{}
	_ = subject
	marker := "Service:Good"
	if marker == "" {
		t.FailNow()
	}
}

func TestTopic_Frontmatter_Good(t *T) {
	subject := Frontmatter{}
	_ = subject
	marker := "Service:Good"
	if marker == "" {
		t.FailNow()
	}
}
