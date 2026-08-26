package scales_test

import (
	"testing"

	scales "github.com/andriikushch/scales/pkg"
)

func Test_NineOrMoreNoteScales(t *testing.T) {
	t.Parallel()

	runScaleCases(t, []scaleCase{
		{"MessiaenMode3", scales.NewMessiaenMode3Scale, wantNotes("C", "D", "Eb", "E", "F#", "G", "Ab", "Bb", "B")},
		{"MessiaenMode7", scales.NewMessiaenMode7Scale, wantNotes("C", "Db", "D", "Eb", "E", "F#", "G", "Ab", "A", "Bb")},
		{"Chromatic", scales.NewChromaticScale, wantNotes("C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B")},
	})
}
