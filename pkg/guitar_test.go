package scales_test

import (
	scales "github.com/andriikushch/scales/pkg"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGuitar_Draw(t *testing.T) {
	t.Parallel()

	harmonicMinorScale, _ := scales.NewHarmonicMinorScale("G#")

	tests := []struct {
		name        string
		notesToDraw []scales.Note
	}{
		{
			name: "Am",
			notesToDraw: []scales.Note{
				scales.NewNote("A"),
				scales.NewNote("C"),
				scales.NewNote("E"),
			},
		},
		{
			name:        "G# harmonic minor scale",
			notesToDraw: harmonicMinorScale.GetNotes(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := scales.NewGuitarWithStandardTuning()

			require.NoError(t, g.Draw(tt.notesToDraw))
		})
	}
}
