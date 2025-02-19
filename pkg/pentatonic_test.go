package scales_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	scales "github.com/andriikushch/scales/pkg"
	"github.com/andriikushch/scales/pkg/internal/chromatic"
)

func Test_NewMajorPentatonicScale(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		key  string
		want []scales.Note
	}{
		{
			name: chromatic.C,
			key:  chromatic.C,
			want: []scales.Note{
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
			},
		},
		{
			name: chromatic.G,
			key:  chromatic.G,
			want: []scales.Note{
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.E),
			},
		},
		{
			name: chromatic.D,
			key:  chromatic.D,
			want: []scales.Note{
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.B),
			},
		},
		{
			name: chromatic.A,
			key:  chromatic.A,
			want: []scales.Note{
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
			},
		},
		{
			name: chromatic.E,
			key:  chromatic.E,
			want: []scales.Note{
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
			},
		},
		{
			name: chromatic.B,
			key:  chromatic.B,
			want: []scales.Note{
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.DSharp),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.GSharp),
			},
		},
		{
			name: chromatic.FSharp,
			key:  chromatic.FSharp,
			want: []scales.Note{
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.ASharp),
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.DSharp),
			},
		},
		{
			name: chromatic.CSharp,
			key:  chromatic.CSharp,
			want: []scales.Note{
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.DSharp),
				scales.NewNote(chromatic.ESharp), // E# = F
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.ASharp),
			},
		},
		{
			name: chromatic.F,
			key:  chromatic.F,
			want: []scales.Note{
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.D),
			},
		},
		{
			name: chromatic.BFlat,
			key:  chromatic.BFlat,
			want: []scales.Note{
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
			},
		},
		{
			name: chromatic.EFlat,
			key:  chromatic.EFlat,
			want: []scales.Note{
				scales.NewNote(chromatic.EFlat),
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.C),
			},
		},
		{
			name: chromatic.AFlat,
			key:  chromatic.AFlat,
			want: []scales.Note{
				scales.NewNote(chromatic.AFlat),
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.EFlat),
				scales.NewNote(chromatic.F),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			scale, err := scales.NewMajorPentatonicScale(tt.key)

			require.NoError(t, err)
			require.Equal(t, tt.want, scale.GetNotes())
		})
	}
}

//nolint:dupl
func Test_NewMinorPentatonicScale(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		key  string
		want []scales.Note
	}{
		{
			name: chromatic.A,
			key:  chromatic.A,
			want: []scales.Note{
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.G),
			},
		},
		{
			name: chromatic.E,
			key:  chromatic.E,
			want: []scales.Note{
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.D),
			},
		},
		{
			name: chromatic.B,
			key:  chromatic.B,
			want: []scales.Note{
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.A),
			},
		},
		{
			name: chromatic.FSharp,
			key:  chromatic.FSharp,
			want: []scales.Note{
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.E),
			},
		},
		{
			name: chromatic.CSharp,
			key:  chromatic.CSharp,
			want: []scales.Note{
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.B),
			},
		},
		{
			name: chromatic.GSharp,
			key:  chromatic.GSharp,
			want: []scales.Note{
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.DSharp),
				scales.NewNote(chromatic.FSharp),
			},
		},
		{
			name: chromatic.DSharp,
			key:  chromatic.DSharp,
			want: []scales.Note{
				scales.NewNote(chromatic.DSharp),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.ASharp),
				scales.NewNote(chromatic.CSharp),
			},
		},
		{
			name: chromatic.BFlat,
			key:  chromatic.BFlat,
			want: []scales.Note{
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.DFlat),
				scales.NewNote(chromatic.EFlat),
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.AFlat),
			},
		},
		{
			name: chromatic.F,
			key:  chromatic.F,
			want: []scales.Note{
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.AFlat),
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.EFlat),
			},
		},
		{
			name: chromatic.C,
			key:  chromatic.C,
			want: []scales.Note{
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.EFlat),
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.BFlat),
			},
		},
		{
			name: chromatic.G,
			key:  chromatic.G,
			want: []scales.Note{
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.F),
			},
		},
		{
			name: chromatic.D,
			key:  chromatic.D,
			want: []scales.Note{
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.C),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			scale, err := scales.NewMinorPentatonicScale(tt.key)

			require.NoError(t, err)
			require.Equal(t, tt.want, scale.GetNotes())
		})
	}
}
