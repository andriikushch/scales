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
	err = g.drawChord(guitar_maj7_sharp_5_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_maj7_sharp_5_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_maj7_sharp_5_2_ChordShape, chord, os.Stdout)
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

func TestCm7b5(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("Cm7b5")
	require.NoError(t, err)
	err = g.drawChord(guitar_m7b5_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_m7b5_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_m7b5_2_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestCdim7(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("Cdim7")
	require.NoError(t, err)
	err = g.drawChord(guitar_dim7_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
	err = g.drawChord(guitar_dim7_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
	err = g.drawChord(guitar_dim7_2_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
	err = g.drawChord(guitar_dim7_3_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
	err = g.drawChord(guitar_dim7_4_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
	err = g.drawChord(guitar_dim7_5_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
	err = g.drawChord(guitar_dim7_6_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
	err = g.drawChord(guitar_dim7_7_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
	err = g.drawChord(guitar_dim7_8_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
	err = g.drawChord(guitar_dim7_9_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
	err = g.drawChord(guitar_dim7_10_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
	err = g.drawChord(guitar_dim7_11_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestCaug7(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("Caug7")
	require.NoError(t, err)
	err = g.drawChord(guitar_aug7_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_aug7_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_aug7_2_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestC7b5(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("C7b5")
	require.NoError(t, err)
	err = g.drawChord(guitar_7b5_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_7b5_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_7b5_2_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestCmaj9(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("Cmaj9")
	require.NoError(t, err)
	err = g.drawChord(guitar_maj9_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_maj9_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestC9(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("C9")
	require.NoError(t, err)
	err = g.drawChord(guitar_9_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_9_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestCm9(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("Cm9")
	require.NoError(t, err)
	err = g.drawChord(guitar_m9_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_m9_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestC11(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("C11")
	require.NoError(t, err)
	err = g.drawChord(guitar_11_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_11_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestCm11(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("Cm11")
	require.NoError(t, err)
	err = g.drawChord(guitar_m11_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_m11_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestC7b9(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("C7b9")
	require.NoError(t, err)
	err = g.drawChord(guitar_7b9_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_7b9_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestC7Sharp9(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("C7#9")
	require.NoError(t, err)
	err = g.drawChord(guitar_7_sharp_9_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_7_sharp_9_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestC7b5b9(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("C7b5b9")
	require.NoError(t, err)
	err = g.drawChord(guitar_7b5_b9_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_7b5_b9_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestC7Sharp5Sharp9(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("C7#5#9")
	require.NoError(t, err)
	err = g.drawChord(guitar_7_sharp_5_sharp_9_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_7_sharp_5_sharp_9_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestC7b5Sharp9(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("C7b5#9")
	require.NoError(t, err)
	err = g.drawChord(guitar_7b5_sharp_9_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_7b5_sharp_9_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestC7Sharp5b9(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("C7#5b9")
	require.NoError(t, err)
	err = g.drawChord(guitar_7_sharp_5_b9_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_7_sharp_5_b9_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestC7Sharp11(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("C7#11")
	require.NoError(t, err)
	err = g.drawChord(guitar_7_sharp_11_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_7_sharp_11_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestC9b5(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("C9b5")
	require.NoError(t, err)
	err = g.drawChord(guitar_9b5_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_9b5_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestC9Sharp5(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("C9#5")
	require.NoError(t, err)
	err = g.drawChord(guitar_9_sharp_5_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_9_sharp_5_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestC9Sharp11(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("C9#11")
	require.NoError(t, err)
	err = g.drawChord(guitar_9_sharp_11_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_9_sharp_11_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestC6(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("C6")
	require.NoError(t, err)
	err = g.drawChord(guitar_6_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_6_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestCm6(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("Cm6")
	require.NoError(t, err)
	err = g.drawChord(guitar_m6_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_m6_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestC6Slash9(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("C6/9")
	require.NoError(t, err)
	err = g.drawChord(guitar_6_add9_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_6_add9_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestCsus2(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("Csus2")
	require.NoError(t, err)
	err = g.drawChord(guitar_sus2_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_sus2_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestCsus(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("Csus")
	require.NoError(t, err)
	err = g.drawChord(guitar_sus_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_sus_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestCsus4(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("Csus4")
	require.NoError(t, err)
	err = g.drawChord(guitar_sus_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_sus_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestC7sus4(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("C7sus4")
	require.NoError(t, err)
	err = g.drawChord(guitar_7sus_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_7sus_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestC7sus(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("C7sus")
	require.NoError(t, err)
	err = g.drawChord(guitar_7sus_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_7sus_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestCadd9(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("Cadd9")
	require.NoError(t, err)
	err = g.drawChord(guitar_add9_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_add9_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestCmadd9(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("Cm(add9)")
	require.NoError(t, err)
	err = g.drawChord(guitar_m_add9_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_m_add9_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestCaug(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("Caug")
	require.NoError(t, err)
	err = g.drawChord(guitar_aug_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_aug_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(guitar_aug_2_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestPrint(t *testing.T) {
	g := NewGuitarWithStandardTuning()
	chord, err := NewChord("Caug")
	require.NoError(t, err)
	shapes := g.getChordShapes(chord)
	for k, v := range shapes {
		fmt.Printf("shape %d: %v\n", k, v)
	}
}
