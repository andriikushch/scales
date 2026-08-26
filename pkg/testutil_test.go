package scales_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	scales "github.com/andriikushch/scales/pkg"
)

// wantNotes builds a []scales.Note from note-name strings, for comparing against
// scale.GetNotes() in table-driven tests.
func wantNotes(names ...string) []scales.Note {
	notes := make([]scales.Note, len(names))
	for i, n := range names {
		notes[i] = scales.NewNote(n)
	}

	return notes
}

type scaleCase struct {
	name string
	fn   func(string) (*scales.Scale, error)
	want []scales.Note
}

// runScaleCases builds each case's scale in the key of C and asserts its notes match, one
// subtest per scale.
func runScaleCases(t *testing.T, cases []scaleCase) {
	t.Helper()

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			scale, err := tt.fn("C")

			require.NoError(t, err)
			require.Equal(t, tt.want, scale.GetNotes())
		})
	}
}
