package scales_test

import (
	"testing"

	scales "github.com/andriikushch/scales/pkg"
)

func Test_ModesOfHarmonicMinorScale(t *testing.T) {
	t.Parallel()

	runScaleCases(t, []scaleCase{
		{"LocrianNatural6", scales.NewLocrianNatural6Scale, wantNotes("C", "Db", "Eb", "F", "Gb", "A", "Bb")},
		{"IonianSharp5", scales.NewIonianSharp5Scale, wantNotes("C", "D", "E", "F", "G#", "A", "B")},
		{"DorianSharp4", scales.NewDorianSharp4Scale, wantNotes("C", "D", "Eb", "F#", "G", "A", "Bb")},
		{"PhrygianDominant", scales.NewPhrygianDominantScale, wantNotes("C", "Db", "E", "F", "G", "Ab", "Bb")},
		{"LydianSharp2", scales.NewLydianSharp2Scale, wantNotes("C", "D#", "E", "F#", "G", "A", "B")},
		{"Ultralocrian", scales.NewUltralocrianScale, wantNotes("C", "Db", "Eb", "Fb", "Gb", "Ab", "Bbb")},
	})
}
