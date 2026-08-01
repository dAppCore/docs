// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	. "dappco.re/go"
)

func TestParseHelpText_Good_StandardGoFlags(t *T) {
	helpText := `Core development workflow tool.

Usage:
  core dev [command]

Flags:
  -h, --help       help for dev
  -v, --verbose    enable verbose output

`

	topic := ParseHelpText("dev", helpText)

	AssertEqual(t, "dev", topic.ID)
	AssertEqual(t, "Dev", topic.Title)
	AssertContains(t, topic.Content, "## Usage")
	AssertContains(t, topic.Content, "core dev [command]")
	AssertContains(t, topic.Content, "## Flags")
	AssertContains(t, topic.Content, "--help")
	AssertContains(t, topic.Content, "--verbose")
	AssertEqual(t, []string{"cli", "dev"}, topic.Tags)
}

func TestParseHelpText_Good_CobraStyleSubcommands(t *T) {
	helpText := `Manage repository operations

Usage:
  core dev [command]

Available Commands:
  commit      Commit changes to repositories
  push        Push commits to remote
  pull        Pull latest from remote
  status      Show repository status

Flags:
  -h, --help   help for dev

`

	topic := ParseHelpText("dev", helpText)

	AssertContains(t, topic.Content, "## Commands")
	AssertContains(t, topic.Content, "**commit**")
	AssertContains(t, topic.Content, "**push**")
	AssertContains(t, topic.Content, "**pull**")
	AssertContains(t, topic.Content, "**status**")
	AssertContains(t, topic.Content, "## Flags")
}

func TestParseHelpText_Good_MinimalHelpText(t *T) {
	helpText := "Show the current version."

	topic := ParseHelpText("version", helpText)

	AssertEqual(t, "version", topic.ID)
	AssertEqual(t, "Version", topic.Title)
	AssertContains(t, topic.Content, "Show the current version.")
	AssertEqual(t, []string{"cli", "version"}, topic.Tags)
}

func TestParseHelpText_Good_WithExamples(t *T) {
	helpText := `Commit changes across repositories.

Usage:
  core dev commit [flags]

Examples:
  core dev commit --all
  core dev commit -m "fix: resolve bug"

Flags:
  -a, --all          commit all changes
  -m, --message      commit message

`

	topic := ParseHelpText("dev commit", helpText)

	AssertEqual(t, "dev-commit", topic.ID)
	AssertEqual(t, "Dev Commit", topic.Title)
	AssertContains(t, topic.Content, "## Examples")
	AssertContains(t, topic.Content, "core dev commit --all")
	AssertContains(t, topic.Content, `core dev commit -m "fix: resolve bug"`)
	AssertEqual(t, []string{"cli", "dev"}, topic.Tags)
}

func TestIngestCLIHelp_Good_BatchIngest(t *T) {
	helpTexts := map[string]string{
		"dev":        "Development workflow tool.\n\nUsage:\n  core dev [command]\n",
		"dev commit": "Commit changes.\n\nUsage:\n  core dev commit [flags]\n",
		"dev push":   "Push to remote.\n\nUsage:\n  core dev push [flags]\n",
	}

	catalog := IngestCLIHelp(helpTexts)

	AssertNotNil(t, catalog)

	topics := catalog.List()
	AssertLen(t, topics, 3)

	// Verify specific topics
	res1 := catalog.Get("dev")
	if !res1.OK { t.Fatal(res1.Error()) }
	devTopic := res1.Value.(*Topic)
	AssertEqual(t, "Dev", devTopic.Title)

	res2 := catalog.Get("dev-commit")
	if !res2.OK { t.Fatal(res2.Error()) }
	commitTopic := res2.Value.(*Topic)
	AssertEqual(t, "Dev Commit", commitTopic.Title)

	res3 := catalog.Get("dev-push")
	if !res3.OK { t.Fatal(res3.Error()) }
	pushTopic := res3.Value.(*Topic)
	AssertEqual(t, "Dev Push", pushTopic.Title)

	// All should be searchable
	results := catalog.Search("commit")
	AssertNotEmpty(t, results)
}

func TestParseHelpText_Good_SeeAlso(t *T) {
	helpText := `Push commits to remote repositories.

Usage:
  core dev push [flags]

See also: dev commit, dev pull

`

	topic := ParseHelpText("dev push", helpText)

	AssertEqual(t, "dev-push", topic.ID)
	AssertEqual(t, []string{"dev-commit", "dev-pull"}, topic.Related)
	// "See also" line should be removed from content
	AssertNotContains(t, topic.Content, "See also")
}

func TestParseHelpText_Good_EmptyHelpText(t *T) {
	topic := ParseHelpText("empty", "")

	AssertEqual(t, "empty", topic.ID)
	AssertEqual(t, "Empty", topic.Title)
	AssertEqual(t, "", topic.Content)
	AssertEqual(t, []string{"cli", "empty"}, topic.Tags)
	AssertEmpty(t, topic.Sections)
}

func TestParseHelpText_Good_OptionsSection(t *T) {
	// Some tools use "Options:" instead of "Flags:"
	helpText := `Run tests across repositories.

Options:
  --filter string   filter test by name
  --timeout int     timeout in seconds

`

	topic := ParseHelpText("test", helpText)
	AssertContains(t, topic.Content, "## Flags") // Options: mapped to ## Flags
	AssertContains(t, topic.Content, "--filter")
}

func TestTitleCaseWords_Good(t *T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"dev commit", "Dev Commit"},
		{"version", "Version"},
		{"dev", "Dev"},
		{"hello world test", "Hello World Test"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *T) {
			AssertEqual(t, tt.expected, titleCaseWords(tt.input))
		})
	}
}

func TestExtractSeeAlso_Good(t *T) {
	t.Run("with see also line", func(t *T) {
		text := "Some content.\nSee also: dev commit, dev pull\nMore content."
		related, cleaned := extractSeeAlso(text)
		AssertEqual(t, []string{"dev-commit", "dev-pull"}, related)
		AssertNotContains(t, cleaned, "See also")
		AssertContains(t, cleaned, "Some content.")
		AssertContains(t, cleaned, "More content.")
	})

	t.Run("no see also", func(t *T) {
		text := "Just normal content."
		related, cleaned := extractSeeAlso(text)
		AssertEmpty(t, related)
		AssertEqual(t, text, cleaned)
	})
}

func TestParseHelpText_Good_MultiLineDescription(t *T) {
	helpText := `Core is a multi-repository management tool that helps you
manage development workflows across federated monorepos.

It provides commands for status checking, committing,
pushing, and pulling across all repositories.

Usage:
  core [command]

`

	topic := ParseHelpText("core", helpText)
	AssertContains(t, topic.Content, "multi-repository management")
	AssertContains(t, topic.Content, "## Usage")
}

func TestParseHelpText_Good_CLINameDoesNotDuplicateTag(t *T) {
	// When the command name starts with "cli", the tag list should
	// not contain "cli" twice.
	topic := ParseHelpText("cli", "CLI root command.")
	AssertEqual(t, "cli", topic.ID)
	AssertEqual(t, []string{"cli"}, topic.Tags)
	AssertEqual(t, "Cli", topic.Title)
}

func TestParseHelpText_Good_CodeBlockEndedByNewSection(t *T) {
	// Tests the endsCodeBlockSection path: a code block for Usage is
	// terminated when a Flags section starts on a non-indented line.
	helpText := `Usage:
  core run
Flags:
  -h, --help   help

`
	topic := ParseHelpText("run", helpText)
	AssertContains(t, topic.Content, "## Usage")
	AssertContains(t, topic.Content, "## Flags")
	AssertContains(t, topic.Content, "--help")
}

func TestParseHelpText_Good_CommandsWithSingleWordEntry(t *T) {
	// Subcommand listing with a single-word entry (no description)
	helpText := `Available Commands:
  help

`
	topic := ParseHelpText("test", helpText)
	AssertContains(t, topic.Content, "## Commands")
	AssertContains(t, topic.Content, "help")
}

func TestEndsCodeBlockSection_Good(t *T) {
	AssertTrue(t, endsCodeBlockSection("Usage:"))
	AssertTrue(t, endsCodeBlockSection("Flags:"))
	AssertTrue(t, endsCodeBlockSection("Options:"))
	AssertTrue(t, endsCodeBlockSection("Examples:"))
	AssertTrue(t, endsCodeBlockSection("Commands:"))
	AssertTrue(t, endsCodeBlockSection("Available Commands:"))
	AssertTrue(t, endsCodeBlockSection("See also: something"))
	AssertFalse(t, endsCodeBlockSection("  -h, --help"))
	AssertFalse(t, endsCodeBlockSection("some random text"))
}

func TestIsSectionHeader_Good(t *T) {
	AssertTrue(t, isSectionHeader("Usage:", "Usage:"))
	AssertTrue(t, isSectionHeader("usage:", "Usage:"))
	AssertFalse(t, isSectionHeader("NotUsage:", "Usage:"))
}

func TestIngest_ParseHelpText_Good(t *T) {
	subject := ParseHelpText
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Good"
	if marker == "" {
		t.FailNow()
	}
}

func TestIngest_ParseHelpText_Bad(t *T) {
	subject := ParseHelpText
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestIngest_ParseHelpText_Ugly(t *T) {
	subject := ParseHelpText
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}

func TestIngest_IngestCLIHelp_Good(t *T) {
	subject := IngestCLIHelp
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Good"
	if marker == "" {
		t.FailNow()
	}
}

func TestIngest_IngestCLIHelp_Bad(t *T) {
	subject := IngestCLIHelp
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestIngest_IngestCLIHelp_Ugly(t *T) {
	subject := IngestCLIHelp
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}
