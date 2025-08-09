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
