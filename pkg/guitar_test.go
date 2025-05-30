package scales

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGuitar_Draw(t *testing.T) {
	t.Parallel()

	harmonicMinorScale, _ := NewHarmonicMinorScale("G#")

	tests := []struct {
		name        string
		notesToDraw []Note
	}{
		{
			name: "Am",
			notesToDraw: []Note{
				NewNote("A"),
				NewNote("C"),
				NewNote("E"),
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
			g := NewGuitarWithStandardTuning()

			require.NoError(t, g.Draw(tt.notesToDraw, io.Discard))
		})
	}
}

func TestGuitar_DrawChord(t *testing.T) {
	t.Parallel()
	g := NewGuitarWithStandardTuning()

	amScale, err := NewNaturalMinorScale("A")
	require.NoError(t, err)
	amChord := amScale.GetChords()[0]

	for _, shape := range g.getChordShapes(amChord) {
		err = g.drawChord(shape, amChord, os.Stdout)
		require.NoError(t, err)
	}

	cScale, err := NewMajorScale("C")
	require.NoError(t, err)
	cChord := cScale.GetChords()[0]
	for _, shape := range g.getChordShapes(cChord) {
		err = g.drawChord(shape, cChord, os.Stdout)
		require.NoError(t, err)
	}
}
