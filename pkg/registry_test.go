package scales_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	scales "github.com/andriikushch/scales/pkg"
)

// Test_ScaleRegistry is a blanket regression guard across every scale in the registry: each
// one must build without error in the key of C and produce at least one note.
func Test_ScaleRegistry(t *testing.T) {
	t.Parallel()

	names := scales.ScaleNames()
	require.Len(t, names, 80)

	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			scale, err := scales.NewScaleByName(name, "C")

			require.NoError(t, err)
			require.NotEmpty(t, scale.GetNotes())
		})
	}
}

func Test_NewScaleByNameUnknownScale(t *testing.T) {
	t.Parallel()

	_, err := scales.NewScaleByName("notARealScale", "C")

	require.Error(t, err)
}
