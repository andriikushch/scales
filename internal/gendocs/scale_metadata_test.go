package main

import (
	"testing"

	"github.com/stretchr/testify/require"

	scales "github.com/andriikushch/scales/pkg"
)

// Test_ScaleCategoriesMatchRegistry guards against pkg/registry.go gaining/losing a canonical
// scale without scaleCategories (docs/scales.md's section data) being updated to match.
func Test_ScaleCategoriesMatchRegistry(t *testing.T) {
	t.Parallel()

	registryNames := map[string]bool{}
	for _, name := range scales.ScaleNames() {
		registryNames[name] = true
	}
	for _, alias := range scaleAliases {
		delete(registryNames, alias.Key)
	}

	categorized := map[string]bool{}
	for _, cat := range scaleCategories {
		for _, entry := range cat.Scales {
			require.Falsef(t, categorized[entry.Key], "scale %q listed in more than one category", entry.Key)
			categorized[entry.Key] = true
		}
	}

	require.Equal(t, registryNames, categorized, "scaleCategories (internal/gendocs) is out of sync with pkg/registry.go's canonical scales")
}

// Test_ScaleAliasesMatchRegistry guards against pkg/registry.go's alias section gaining/losing an
// entry without scaleAliases (docs/scales.md's "Other names" table data) being updated to match.
func Test_ScaleAliasesMatchRegistry(t *testing.T) {
	t.Parallel()

	canonicalKeys := map[string]bool{}
	for _, cat := range scaleCategories {
		for _, entry := range cat.Scales {
			canonicalKeys[entry.Key] = true
		}
	}

	registryAliasNames := map[string]bool{}
	for _, name := range scales.ScaleNames() {
		if !canonicalKeys[name] {
			registryAliasNames[name] = true
		}
	}

	aliased := map[string]bool{}
	for _, alias := range scaleAliases {
		require.Falsef(t, aliased[alias.Key], "alias %q listed more than once", alias.Key)
		aliased[alias.Key] = true
	}

	require.Equal(t, registryAliasNames, aliased, "scaleAliases (internal/gendocs) is out of sync with pkg/registry.go's alias section")
}

// Test_GenerateScalesMarkdownBuilds is a smoke test that every scale referenced in
// scaleCategories/scaleAliases actually builds via scales.NewScaleByName, so a stale/typo'd key
// fails fast instead of silently producing a malformed docs/scales.md.
func Test_GenerateScalesMarkdownBuilds(t *testing.T) {
	t.Parallel()

	content, err := GenerateScalesMarkdown()
	require.NoError(t, err)
	require.NotEmpty(t, content)
}
