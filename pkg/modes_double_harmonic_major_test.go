package scales_test

import (
	"testing"

	scales "github.com/andriikushch/scales/pkg"
)

func Test_ModesOfDoubleHarmonicMajorScale(t *testing.T) {
	t.Parallel()

	runScaleCases(t, []scaleCase{
		{"DoubleHarmonicMajor", scales.NewDoubleHarmonicMajorScale, wantNotes("C", "Db", "E", "F", "G", "Ab", "B")},
		{"LydianSharp2Sharp6", scales.NewLydianSharp2Sharp6Scale, wantNotes("C", "D#", "E", "F#", "G", "A#", "B")},
		{"Ultraphrygian", scales.NewUltraphrygianScale, wantNotes("C", "Db", "Eb", "Fb", "G", "Ab", "Bbb")},
		{"HungarianMinor", scales.NewHungarianMinorScale, wantNotes("C", "D", "Eb", "F#", "G", "Ab", "B")},
		{"Oriental", scales.NewOrientalScale, wantNotes("C", "Db", "E", "F", "Gb", "A", "Bb")},
		{"IonianSharp2Sharp5", scales.NewIonianSharp2Sharp5Scale, wantNotes("C", "D#", "E", "F", "G#", "A", "B")},
		{"LocrianDoubleFlat3DoubleFlat7", scales.NewLocrianDoubleFlat3DoubleFlat7Scale, wantNotes("C", "Db", "Ebb", "F", "Gb", "Ab", "Bbb")},
	})
}
