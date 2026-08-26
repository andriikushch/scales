package scales_test

import (
	"testing"

	scales "github.com/andriikushch/scales/pkg"
)

func Test_ModesOfHarmonicMajorScale(t *testing.T) {
	t.Parallel()

	runScaleCases(t, []scaleCase{
		{"HarmonicMajor", scales.NewHarmonicMajorScale, wantNotes("C", "D", "E", "F", "G", "Ab", "B")},
		{"DorianFlat5", scales.NewDorianFlat5Scale, wantNotes("C", "D", "Eb", "F", "Gb", "A", "Bb")},
		{"PhrygianFlat4", scales.NewPhrygianFlat4Scale, wantNotes("C", "Db", "Eb", "Fb", "G", "Ab", "Bb")},
		{"LydianFlat3", scales.NewLydianFlat3Scale, wantNotes("C", "D", "Eb", "F#", "G", "A", "B")},
		{"MixolydianFlat2", scales.NewMixolydianFlat2Scale, wantNotes("C", "Db", "E", "F", "G", "A", "Bb")},
		{"LydianAugmentedSharp2", scales.NewLydianAugmentedSharp2Scale, wantNotes("C", "D#", "E", "F#", "G#", "A", "B")},
		{"LocrianDoubleFlat7", scales.NewLocrianDoubleFlat7Scale, wantNotes("C", "Db", "Eb", "F", "Gb", "Ab", "Bbb")},
	})
}
