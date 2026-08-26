package scales_test

import (
	"testing"

	scales "github.com/andriikushch/scales/pkg"
)

func Test_BebopScales(t *testing.T) {
	t.Parallel()

	runScaleCases(t, []scaleCase{
		{"BebopDominant", scales.NewBebopDominantScale, wantNotes("C", "D", "E", "F", "G", "A", "Bb", "B")},
		{"BebopMajor", scales.NewBebopMajorScale, wantNotes("C", "D", "E", "F", "G", "G#", "A", "B")},
		{"BebopDorian", scales.NewBebopDorianScale, wantNotes("C", "D", "Eb", "E", "F", "G", "A", "Bb")},
		{"BebopMelodicMinor", scales.NewBebopMelodicMinorScale, wantNotes("C", "D", "Eb", "F", "G", "Ab", "A", "B")},
		{"BebopHarmonicMinor", scales.NewBebopHarmonicMinorScale, wantNotes("C", "D", "Eb", "F", "G", "Ab", "Bb", "B")},
	})
}
