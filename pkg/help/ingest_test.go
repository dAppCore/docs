// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseHelpText_Good_StandardGoFlags(t *testing.T) {
	helpText := `Core development workflow tool.

Usage:
  core dev [command]

Flags:
  -h, --help       help for dev
  -v, --verbose    enable verbose output

`

	topic := ParseHelpText("dev", helpText)

	assert.Equal(t, "dev", topic.ID)
	assert.Equal(t, "Dev", topic.Title)
	assert.Contains(t, topic.Content, "## Usage")
	assert.Contains(t, topic.Content, "core dev [command]")
	assert.Contains(t, topic.Content, "## Flags")
	assert.Contains(t, topic.Content, "--help")
	assert.Contains(t, topic.Content, "--verbose")
	assert.Equal(t, []string{"cli", "dev"}, topic.Tags)
}

func TestParseHelpText_Good_CobraStyleSubcommands(t *testing.T) {
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

	assert.Contains(t, topic.Content, "## Commands")
	assert.Contains(t, topic.Content, "**commit**")
	assert.Contains(t, topic.Content, "**push**")
	assert.Contains(t, topic.Content, "**pull**")
	assert.Contains(t, topic.Content, "**status**")
	assert.Contains(t, topic.Content, "## Flags")
}

func TestParseHelpText_Good_MinimalHelpText(t *testing.T) {
	helpText := "Show the current version."

	topic := ParseHelpText("version", helpText)

	assert.Equal(t, "version", topic.ID)
	assert.Equal(t, "Version", topic.Title)
	assert.Contains(t, topic.Content, "Show the current version.")
	assert.Equal(t, []string{"cli", "version"}, topic.Tags)
}

func TestParseHelpText_Good_WithExamples(t *testing.T) {
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

	assert.Equal(t, "dev-commit", topic.ID)
	assert.Equal(t, "Dev Commit", topic.Title)
	assert.Contains(t, topic.Content, "## Examples")
	assert.Contains(t, topic.Content, "core dev commit --all")
	assert.Contains(t, topic.Content, `core dev commit -m "fix: resolve bug"`)
	assert.Equal(t, []string{"cli", "dev"}, topic.Tags)
}

func TestIngestCLIHelp_Good_BatchIngest(t *testing.T) {
	helpTexts := map[string]string{
		"dev":        "Development workflow tool.\n\nUsage:\n  core dev [command]\n",
		"dev commit": "Commit changes.\n\nUsage:\n  core dev commit [flags]\n",
		"dev push":   "Push to remote.\n\nUsage:\n  core dev push [flags]\n",
	}

	catalog := IngestCLIHelp(helpTexts)

	assert.NotNil(t, catalog)

	topics := catalog.List()
	assert.Len(t, topics, 3)

	// Verify specific topics
	devTopic, err := catalog.Get("dev")
	require.NoError(t, err)
	assert.Equal(t, "Dev", devTopic.Title)

	commitTopic, err := catalog.Get("dev-commit")
	require.NoError(t, err)
	assert.Equal(t, "Dev Commit", commitTopic.Title)

	pushTopic, err := catalog.Get("dev-push")
	require.NoError(t, err)
	assert.Equal(t, "Dev Push", pushTopic.Title)

	// All should be searchable
	results := catalog.Search("commit")
	assert.NotEmpty(t, results)
}

func TestParseHelpText_Good_SeeAlso(t *testing.T) {
	helpText := `Push commits to remote repositories.

Usage:
  core dev push [flags]

See also: dev commit, dev pull

`

	topic := ParseHelpText("dev push", helpText)

	assert.Equal(t, "dev-push", topic.ID)
	assert.Equal(t, []string{"dev-commit", "dev-pull"}, topic.Related)
	// "See also" line should be removed from content
	assert.NotContains(t, topic.Content, "See also")
}

func TestParseHelpText_Good_EmptyHelpText(t *testing.T) {
	topic := ParseHelpText("empty", "")

	assert.Equal(t, "empty", topic.ID)
	assert.Equal(t, "Empty", topic.Title)
	assert.Equal(t, "", topic.Content)
	assert.Equal(t, []string{"cli", "empty"}, topic.Tags)
	assert.Empty(t, topic.Sections)
}

func TestParseHelpText_Good_OptionsSection(t *testing.T) {
	// Some tools use "Options:" instead of "Flags:"
	helpText := `Run tests across repositories.

Options:
  --filter string   filter test by name
  --timeout int     timeout in seconds

`

	topic := ParseHelpText("test", helpText)
	assert.Contains(t, topic.Content, "## Flags") // Options: mapped to ## Flags
	assert.Contains(t, topic.Content, "--filter")
}

func TestTitleCaseWords_Good(t *testing.T) {
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
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, titleCaseWords(tt.input))
		})
	}
}

func TestExtractSeeAlso_Good(t *testing.T) {
	t.Run("with see also line", func(t *testing.T) {
		text := "Some content.\nSee also: dev commit, dev pull\nMore content."
		related, cleaned := extractSeeAlso(text)
		assert.Equal(t, []string{"dev-commit", "dev-pull"}, related)
		assert.NotContains(t, cleaned, "See also")
		assert.Contains(t, cleaned, "Some content.")
		assert.Contains(t, cleaned, "More content.")
	})

	t.Run("no see also", func(t *testing.T) {
		text := "Just normal content."
		related, cleaned := extractSeeAlso(text)
		assert.Empty(t, related)
		assert.Equal(t, text, cleaned)
	})
}

func TestParseHelpText_Good_MultiLineDescription(t *testing.T) {
	helpText := `Core is a multi-repository management tool that helps you
manage development workflows across federated monorepos.

It provides commands for status checking, committing,
pushing, and pulling across all repositories.

Usage:
  core [command]

`

	topic := ParseHelpText("core", helpText)
	assert.Contains(t, topic.Content, "multi-repository management")
	assert.Contains(t, topic.Content, "## Usage")
}

func TestParseHelpText_Good_CLINameDoesNotDuplicateTag(t *testing.T) {
	// When the command name starts with "cli", the tag list should
	// not contain "cli" twice.
	topic := ParseHelpText("cli", "CLI root command.")
	assert.Equal(t, []string{"cli"}, topic.Tags)
}

func TestParseHelpText_Good_CodeBlockEndedByNewSection(t *testing.T) {
	// Tests the endsCodeBlockSection path: a code block for Usage is
	// terminated when a Flags section starts on a non-indented line.
	helpText := `Usage:
  core run
Flags:
  -h, --help   help

`
	topic := ParseHelpText("run", helpText)
	assert.Contains(t, topic.Content, "## Usage")
	assert.Contains(t, topic.Content, "## Flags")
	assert.Contains(t, topic.Content, "--help")
}

func TestParseHelpText_Good_CommandsWithSingleWordEntry(t *testing.T) {
	// Subcommand listing with a single-word entry (no description)
	helpText := `Available Commands:
  help

`
	topic := ParseHelpText("test", helpText)
	assert.Contains(t, topic.Content, "## Commands")
	assert.Contains(t, topic.Content, "help")
}

func TestEndsCodeBlockSection_Good(t *testing.T) {
	assert.True(t, endsCodeBlockSection("Usage:"))
	assert.True(t, endsCodeBlockSection("Flags:"))
	assert.True(t, endsCodeBlockSection("Options:"))
	assert.True(t, endsCodeBlockSection("Examples:"))
	assert.True(t, endsCodeBlockSection("Commands:"))
	assert.True(t, endsCodeBlockSection("Available Commands:"))
	assert.True(t, endsCodeBlockSection("See also: something"))
	assert.False(t, endsCodeBlockSection("  -h, --help"))
	assert.False(t, endsCodeBlockSection("some random text"))
}

func TestIsSectionHeader_Good(t *testing.T) {
	assert.True(t, isSectionHeader("Usage:", "Usage:"))
	assert.True(t, isSectionHeader("usage:", "Usage:"))
	assert.False(t, isSectionHeader("NotUsage:", "Usage:"))
}
