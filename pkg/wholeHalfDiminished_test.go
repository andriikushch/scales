package scales_test

import (
	"testing"

	scales "github.com/andriikushch/scales/pkg"
	"github.com/andriikushch/scales/pkg/internal"
	"github.com/stretchr/testify/require"
)

func TestNewWholeHalfDiminishedScale(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		key  string
		want []scales.Note
	}{
		{
			name: internal.C,
			key:  internal.C,
			want: []scales.Note{
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
				scales.NewNote(internal.EFlat),
				scales.NewNote(internal.F),
				scales.NewNote(internal.GFlat),
				scales.NewNote(internal.AFlat),
				scales.NewNote(internal.A),
				scales.NewNote(internal.B),
			},
		},
		{
			name: internal.DFlat,
			key:  internal.DFlat,
			want: []scales.Note{
				scales.NewNote(internal.DFlat),
				scales.NewNote(internal.EFlat),
				scales.NewNote(internal.E),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.G),
				scales.NewNote(internal.A),
				scales.NewNote(internal.BFlat),
				scales.NewNote(internal.C),
			},
		},
		{
			name: internal.D,
			key:  internal.D,
			want: []scales.Note{
				scales.NewNote(internal.D),
				scales.NewNote(internal.E),
				scales.NewNote(internal.F),
				scales.NewNote(internal.G),
				scales.NewNote(internal.AFlat),
				scales.NewNote(internal.BFlat),
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
			},
		},
		{
			name: internal.EFlat,
			key:  internal.EFlat,
			want: []scales.Note{
				scales.NewNote(internal.EFlat),
				scales.NewNote(internal.F),
				scales.NewNote(internal.GFlat),
				scales.NewNote(internal.AFlat),
				scales.NewNote(internal.A),
				scales.NewNote(internal.B),
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
			},
		},
		{
			name: internal.E,
			key:  internal.E,
			want: []scales.Note{
				scales.NewNote(internal.E),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.G),
				scales.NewNote(internal.A),
				scales.NewNote(internal.BFlat),
				scales.NewNote(internal.C),
				scales.NewNote(internal.DFlat),
				scales.NewNote(internal.EFlat),
			},
		},
		{
			name: internal.F,
			key:  internal.F,
			want: []scales.Note{
				scales.NewNote(internal.F),
				scales.NewNote(internal.G),
				scales.NewNote(internal.AFlat),
				scales.NewNote(internal.BFlat),
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.D),
				scales.NewNote(internal.E),
			},
		},
		{
			name: internal.GFlat,
			key:  internal.GFlat,
			want: []scales.Note{
				scales.NewNote(internal.GFlat),
				scales.NewNote(internal.AFlat),
				scales.NewNote(internal.A),
				scales.NewNote(internal.B),
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
				scales.NewNote(internal.EFlat),
				scales.NewNote(internal.F),
			},
		},
		{
			name: internal.G,
			key:  internal.G,
			want: []scales.Note{
				scales.NewNote(internal.G),
				scales.NewNote(internal.A),
				scales.NewNote(internal.BFlat),
				scales.NewNote(internal.C),
				scales.NewNote(internal.DFlat),
				scales.NewNote(internal.EFlat),
				scales.NewNote(internal.E),
				scales.NewNote(internal.FSharp),
			},
		},
		{
			name: internal.AFlat,
			key:  internal.AFlat,
			want: []scales.Note{
				scales.NewNote(internal.AFlat),
				scales.NewNote(internal.BFlat),
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.D),
				scales.NewNote(internal.E),
				scales.NewNote(internal.F),
				scales.NewNote(internal.G),
			},
		},
		{
			name: internal.A,
			key:  internal.A,
			want: []scales.Note{
				scales.NewNote(internal.A),
				scales.NewNote(internal.B),
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
				scales.NewNote(internal.EFlat),
				scales.NewNote(internal.F),
				scales.NewNote(internal.GFlat),
				scales.NewNote(internal.AFlat),
			},
		},
		{
			name: internal.BFlat,
			key:  internal.BFlat,
			want: []scales.Note{
				scales.NewNote(internal.BFlat),
				scales.NewNote(internal.C),
				scales.NewNote(internal.DFlat),
				scales.NewNote(internal.EFlat),
				scales.NewNote(internal.E),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.G),
				scales.NewNote(internal.A),
			},
		},
		{
			name: internal.B,
			key:  internal.B,
			want: []scales.Note{
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.D),
				scales.NewNote(internal.E),
				scales.NewNote(internal.F),
				scales.NewNote(internal.G),
				scales.NewNote(internal.AFlat),
				scales.NewNote(internal.BFlat),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			scale, err := scales.NewWholeHalfDiminishedScale(tt.key)

			require.NoError(t, err)
			require.Equal(t, tt.want, scale.GetNotes())
		})
	}
}

func Test_WholeHalfDiminishedScaleChords(t *testing.T) {
	t.Parallel()

	scale, err := scales.NewWholeHalfDiminishedScale("C")
	require.NoError(t, err)

	var results []string
	for _, c := range scale.GetChords() {
		results = append(results, c.Description())
	}

	require.Equal(t, []string{"Cdim", "Dmb5", "Ebdim", "Fmb5", "Gbdim", "Abmb5", "Adim", "Bmb5"}, results)
}
