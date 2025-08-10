package scales

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMandolin_DrawChord(t *testing.T) {
	t.Parallel()
	g := NewMandolinWithStandardTuning()

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

func TestMandolin_DrawMajChord(t *testing.T) {
	t.Parallel()
	g := NewMandolinWithStandardTuning()

	amScale, err := NewNaturalMinorScale("A")
	require.NoError(t, err)
	cChord := amScale.GetChords()[2]

	b := &strings.Builder{}

	for _, shape := range g.getChordShapes(cChord) {
		err = g.drawChord(shape, cChord, b)
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
