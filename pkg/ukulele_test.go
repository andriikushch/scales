package scales

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUkulele_DrawChord(t *testing.T) {
	t.Parallel()
	g := NewUkuleleWithStandardTuning()

	amChord, err := NewChord("Am")
	require.NoError(t, err)
	t.Log(amChord.Description())
	for _, shape := range g.getChordShapes(amChord) {
		err = g.drawChord(shape, amChord, os.Stdout)
		require.NoError(t, err)
	}

	cChord, err := NewChord("Cmaj")
	require.NoError(t, err)
	t.Log(cChord.Description())
	for _, shape := range g.getChordShapes(cChord) {
		err = g.drawChord(shape, cChord, os.Stdout)
		require.NoError(t, err)
	}
}
