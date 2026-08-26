package scales_test

import (
	"testing"

	scales "github.com/andriikushch/scales/pkg"
)

func Test_ModesOfMelodicMinorScale(t *testing.T) {
	t.Parallel()

	runScaleCases(t, []scaleCase{
		{"DorianFlat2", scales.NewDorianFlat2Scale, wantNotes("C", "Db", "Eb", "F", "G", "A", "Bb")},
		{"LydianAugmented", scales.NewLydianAugmentedScale, wantNotes("C", "D", "E", "F#", "G#", "A", "B")},
		{"LydianDominant", scales.NewLydianDominantScale, wantNotes("C", "D", "E", "F#", "G", "A", "Bb")},
		{"MixolydianFlat6", scales.NewMixolydianFlat6Scale, wantNotes("C", "D", "E", "F", "G", "Ab", "Bb")},
		{"LocrianNatural2", scales.NewLocrianNatural2Scale, wantNotes("C", "D", "Eb", "F", "Gb", "Ab", "Bb")},
		{"Altered", scales.NewAlteredScale, wantNotes("C", "Db", "Eb", "Fb", "Gb", "Ab", "Bb")},
	})
}
