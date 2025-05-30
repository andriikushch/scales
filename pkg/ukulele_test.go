package scales

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUkulele_DrawChord(t *testing.T) {
	t.Parallel()
	g := NewUkuleleWithStandardTuning()

	amScale, err := NewNaturalMinorScale("A")

	require.NoError(t, err)
	amChord := amScale.GetChords()[0]
	t.Log(amChord.Description())
	for _, shape := range g.getChordShapes(amChord) {
		err = g.drawChord(shape, amChord, os.Stdout)
		require.NoError(t, err)
	}

	cScale, err := NewMajorScale("C")
	require.NoError(t, err)
	cChord := cScale.GetChords()[0]
	t.Log(cChord.Description())
	for _, shape := range g.getChordShapes(cChord) {
		err = g.drawChord(shape, cChord, os.Stdout)
		require.NoError(t, err)
	}
}
