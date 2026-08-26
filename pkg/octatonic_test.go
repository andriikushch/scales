package scales_test

import (
	"testing"

	scales "github.com/andriikushch/scales/pkg"
)

func Test_OctatonicScales(t *testing.T) {
	t.Parallel()

	runScaleCases(t, []scaleCase{
		{"HalfWholeDiminished", scales.NewHalfWholeDiminishedScale, wantNotes("C", "Db", "D#", "E", "F#", "G", "A", "Bb")},
		{"MessiaenMode4", scales.NewMessiaenMode4Scale, wantNotes("C", "Db", "D", "F", "Gb", "G", "Ab", "B")},
		{"MessiaenMode6", scales.NewMessiaenMode6Scale, wantNotes("C", "D", "E", "F", "F#", "G#", "A#", "B")},
		{"SpanishEightTone", scales.NewSpanishEightToneScale, wantNotes("C", "Db", "Eb", "E", "F", "Gb", "Ab", "Bb")},
	})
}
