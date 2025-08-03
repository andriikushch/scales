package scales

import (
	"fmt"
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

func TestCmMaj7(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("Cm(maj7)")
	require.NoError(t, err)
	err = g.drawChord(guitar_m_maj7_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_m_maj7_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_m_maj7_2_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestCmaj7(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("Cmaj7")
	require.NoError(t, err)
	err = g.drawChord(guitar_maj7_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_maj7_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_maj7_2_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestC7(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("C7")
	require.NoError(t, err)
	err = g.drawChord(guitar_7_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_7_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_7_2_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestCm7(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("Cm7")
	require.NoError(t, err)
	err = g.drawChord(guitar_m7_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_m7_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_m7_2_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestXX(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("Cm7")
	require.NoError(t, err)
	shapes := g.getChordShapes(chord)
	for k, v := range shapes {
		fmt.Printf("shape %d: %v\n", k, v)
	}
}
