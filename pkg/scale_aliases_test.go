package scales_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	scales "github.com/andriikushch/scales/pkg"
)

// Test_ScaleAliasesMatchCanonicalScale spot-checks that a handful of alias names from
// docs/scales.md's "Other names" table produce identical notes to the canonical scale they
// name, guarding against future drift between an alias wrapper and its target.
func Test_ScaleAliasesMatchCanonicalScale(t *testing.T) {
	t.Parallel()

	tests := []struct {
		alias     string
		canonical string
	}{
		{"freygish", "phrygianDominant"},
		{"jazzMinor", "melodicMinor"},
		{"naturalMinor", "aeolian"},
		{"major", "ionian"},
		{"manGong", "bluesMinorPentatonic"},
		{"chinese", "majorPentatonic"},
	}

	for _, tt := range tests {
		t.Run(tt.alias, func(t *testing.T) {
			t.Parallel()

			aliasScale, err := scales.NewScaleByName(tt.alias, "C")
			require.NoError(t, err)

			canonicalScale, err := scales.NewScaleByName(tt.canonical, "C")
			require.NoError(t, err)

			require.Equal(t, canonicalScale.GetNotes(), aliasScale.GetNotes())
		})
	}
}
