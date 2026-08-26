package scales_test

import (
	"testing"

	scales "github.com/andriikushch/scales/pkg"
)

func Test_ModesOfMajorScale(t *testing.T) {
	t.Parallel()

	runScaleCases(t, []scaleCase{
		{"Dorian", scales.NewDorianScale, wantNotes("C", "D", "Eb", "F", "G", "A", "Bb")},
		{"Phrygian", scales.NewPhrygianScale, wantNotes("C", "Db", "Eb", "F", "G", "Ab", "Bb")},
		{"Lydian", scales.NewLydianScale, wantNotes("C", "D", "E", "F#", "G", "A", "B")},
		{"Mixolydian", scales.NewMixolydianScale, wantNotes("C", "D", "E", "F", "G", "A", "Bb")},
		{"Locrian", scales.NewLocrianScale, wantNotes("C", "Db", "Eb", "F", "Gb", "Ab", "Bb")},
	})
}
