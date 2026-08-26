package scales_test

import (
	"testing"

	scales "github.com/andriikushch/scales/pkg"
)

func Test_OtherHeptatonicScales(t *testing.T) {
	t.Parallel()

	runScaleCases(t, []scaleCase{
		{"NeapolitanMinor", scales.NewNeapolitanMinorScale, wantNotes("C", "Db", "Eb", "F", "G", "Ab", "B")},
		{"NeapolitanMajor", scales.NewNeapolitanMajorScale, wantNotes("C", "Db", "Eb", "F", "G", "A", "B")},
		{"LeadingWholeTone", scales.NewLeadingWholeToneScale, wantNotes("C", "D", "E", "F#", "G#", "A#", "B")},
		{"MajorLocrian", scales.NewMajorLocrianScale, wantNotes("C", "D", "E", "F", "Gb", "Ab", "Bb")},
		{"Gypsy", scales.NewGypsyScale, wantNotes("C", "D", "Eb", "F#", "G", "Ab", "Bb")},
		{"HungarianMajor", scales.NewHungarianMajorScale, wantNotes("C", "D#", "E", "F#", "G", "A", "Bb")},
		{"Persian", scales.NewPersianScale, wantNotes("C", "Db", "E", "F", "Gb", "Ab", "B")},
		{"Enigmatic", scales.NewEnigmaticScale, wantNotes("C", "Db", "E", "F#", "G#", "A#", "B")},
	})
}
