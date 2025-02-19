package scales_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	scales "github.com/andriikushch/scales/pkg"
	"github.com/andriikushch/scales/pkg/internal/chromatic"
)

func Test_NewMajorScale(t *testing.T) {
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
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.B),
			},
		},
		{
			name: chromatic.G,
			key:  chromatic.G,
			want: []scales.Note{
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
			},
		},
		{
			name: chromatic.D,
			key:  chromatic.D,
			want: []scales.Note{
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
			},
		},
		{
			name: chromatic.A,
			key:  chromatic.A,
			want: []scales.Note{
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.GSharp),
			},
		},
		{
			name: string(chromatic.E),
			key:  chromatic.E,
			want: []scales.Note{
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.DSharp),
			},
		},
		{
			name: string(chromatic.B),
			key:  chromatic.B,
			want: []scales.Note{
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.DSharp),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.ASharp),
			},
		},
		{
			name: chromatic.FSharp,
			key:  chromatic.FSharp,
			want: []scales.Note{
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.ASharp),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.DSharp),
				scales.NewNote(chromatic.ESharp),
			},
		},
		{
			name: chromatic.CSharp,
			key:  chromatic.CSharp,
			want: []scales.Note{
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.DSharp),
				scales.NewNote(chromatic.ESharp),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.ASharp),
				scales.NewNote(chromatic.BSharp),
			},
		},
		{
			name: chromatic.F,
			key:  chromatic.F,
			want: []scales.Note{
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.E),
			},
		},
		{
			name: chromatic.BFlat,
			key:  chromatic.BFlat,
			want: []scales.Note{
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.EFlat),
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
			},
		},
		{
			name: chromatic.EFlat,
			key:  chromatic.EFlat,
			want: []scales.Note{
				scales.NewNote(chromatic.EFlat),
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.AFlat),
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.D),
			},
		},
		{
			name: chromatic.AFlat,
			key:  chromatic.AFlat,
			want: []scales.Note{
				scales.NewNote(chromatic.AFlat),
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.DFlat),
				scales.NewNote(chromatic.EFlat),
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			scale, err := scales.NewMajorScale(tt.key)

			require.NoError(t, err)
			require.Equal(t, tt.want, scale.GetNotes())
		})
	}
}
