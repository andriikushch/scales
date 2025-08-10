package scales

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBassGuitar_DrawChord(t *testing.T) {
	t.Parallel()
	g := NewBassGuitarWithStandardTuning()

	amScale, err := NewNaturalMinorScale("A")
	require.NoError(t, err)
	amChord := amScale.GetChords()[0]

	b := &strings.Builder{}

	for _, shape := range g.getChordShapes(amChord) {
		err = g.drawChord(shape, amChord, b)
		require.NoError(t, err)
	}

	lines := strings.Split(b.String(), "\r\n")

	for i := 0; i < len(g.tuning)+2; i++ {
		for j := i; j < len(lines); j += len(g.tuning) + 2 {
			fmt.Print(lines[j], "    ")
		}
		fmt.Println()
	}
}

func TestBassGuitarAm(t *testing.T) {
	g := NewBassGuitarWithStandardTuning()
	chord, err := NewChord("Am")
	require.NoError(t, err)
	err = g.drawChord(bassGuitar_m_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(bassGuitar_m_1_ChordShape, chord, os.Stdout)
	require.NoError(t, err)

	err = g.drawChord(bassGuitar_m_2_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestBassGuitarCmaj7Sharp5(t *testing.T) {
	g := NewBassGuitarWithStandardTuning()
	chord, err := NewChord("Cmaj7#5")
	require.NoError(t, err)
	err = g.drawChord(bassGuitar_maj7_sharp_5_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestBassGuitarAug(t *testing.T) {
	g := NewBassGuitarWithStandardTuning()
	chord, err := NewChord("Caug")
	require.NoError(t, err)

	err = g.drawChord(bassGuitar_aug_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestBassGuitarCmadd9(t *testing.T) {
	g := NewBassGuitarWithStandardTuning()
	chord, err := NewChord("Cm(add9)")
	require.NoError(t, err)

	err = g.drawChord(bassGuitar_m_add9_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestBassGuitarCm7(t *testing.T) {
	g := NewBassGuitarWithStandardTuning()
	chord, err := NewChord("Cm7")
	require.NoError(t, err)

	err = g.drawChord(bassGuitar_m7_0_ChordShape, chord, os.Stdout)
	require.NoError(t, err)
}

func TestPrint(t *testing.T) {
	g := NewBassGuitarWithStandardTuning()
	chord, err := NewChord("Caug")
	require.NoError(t, err)
	shapes := g.getChordShapes(chord)
	for k, v := range shapes {
		fmt.Printf("shape %d: %v\n", k, v)
	}
}
