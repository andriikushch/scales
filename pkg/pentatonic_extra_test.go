package scales_test

import (
	"testing"

	scales "github.com/andriikushch/scales/pkg"
)

func Test_OtherPentatonicScales(t *testing.T) {
	t.Parallel()

	runScaleCases(t, []scaleCase{
		{"SuspendedPentatonic", scales.NewSuspendedPentatonicScale, wantNotes("C", "D", "F", "G", "Bb")},
		{"BluesMinorPentatonic", scales.NewBluesMinorPentatonicScale, wantNotes("C", "Eb", "F", "Ab", "Bb")},
		{"BluesMajorPentatonic", scales.NewBluesMajorPentatonicScale, wantNotes("C", "D", "F", "G", "A")},
		{"Iwato", scales.NewIwatoScale, wantNotes("C", "Db", "F", "Gb", "Bb")},
		{"MiyakoBushi", scales.NewMiyakoBushiScale, wantNotes("C", "Db", "F", "G", "Ab")},
		{"Kumoi", scales.NewKumoiScale, wantNotes("C", "D", "Eb", "G", "A")},
		{"Insen", scales.NewInsenScale, wantNotes("C", "Db", "F", "G", "Bb")},
		{"ChinesePentatonic", scales.NewChinesePentatonicScale, wantNotes("C", "E", "F#", "G", "B")},
		{"BalinesePelog", scales.NewBalinesePelogScale, wantNotes("C", "Db", "Eb", "G", "Ab")},
	})
}
