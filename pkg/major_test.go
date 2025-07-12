package scales_test

import (
	"fmt"
	"testing"

	"github.com/andriikushch/scales/pkg/internal"

	"github.com/stretchr/testify/require"

	scales "github.com/andriikushch/scales/pkg"
)

func Test_NewMajorScale(t *testing.T) {
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
				scales.NewNote(internal.E),
				scales.NewNote(internal.F),
				scales.NewNote(internal.G),
				scales.NewNote(internal.A),
				scales.NewNote(internal.B),
			},
		},
		{
			name: internal.G,
			key:  internal.G,
			want: []scales.Note{
				scales.NewNote(internal.G),
				scales.NewNote(internal.A),
				scales.NewNote(internal.B),
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
				scales.NewNote(internal.E),
				scales.NewNote(internal.FSharp),
			},
		},
		{
			name: internal.D,
			key:  internal.D,
			want: []scales.Note{
				scales.NewNote(internal.D),
				scales.NewNote(internal.E),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.G),
				scales.NewNote(internal.A),
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
			},
		},
		{
			name: internal.A,
			key:  internal.A,
			want: []scales.Note{
				scales.NewNote(internal.A),
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.D),
				scales.NewNote(internal.E),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.GSharp),
			},
		},
		{
			name: string(internal.E),
			key:  internal.E,
			want: []scales.Note{
				scales.NewNote(internal.E),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.A),
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.DSharp),
			},
		},
		{
			name: string(internal.B),
			key:  internal.B,
			want: []scales.Note{
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.DSharp),
				scales.NewNote(internal.E),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.ASharp),
			},
		},
		{
			name: internal.FSharp,
			key:  internal.FSharp,
			want: []scales.Note{
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.ASharp),
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.DSharp),
				scales.NewNote(internal.ESharp),
			},
		},
		{
			name: internal.CSharp,
			key:  internal.CSharp,
			want: []scales.Note{
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.DSharp),
				scales.NewNote(internal.ESharp),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.ASharp),
				scales.NewNote(internal.BSharp),
			},
		},
		{
			name: internal.F,
			key:  internal.F,
			want: []scales.Note{
				scales.NewNote(internal.F),
				scales.NewNote(internal.G),
				scales.NewNote(internal.A),
				scales.NewNote(internal.BFlat),
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
				scales.NewNote(internal.E),
			},
		},
		{
			name: internal.BFlat,
			key:  internal.BFlat,
			want: []scales.Note{
				scales.NewNote(internal.BFlat),
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
				scales.NewNote(internal.EFlat),
				scales.NewNote(internal.F),
				scales.NewNote(internal.G),
				scales.NewNote(internal.A),
			},
		},
		{
			name: internal.EFlat,
			key:  internal.EFlat,
			want: []scales.Note{
				scales.NewNote(internal.EFlat),
				scales.NewNote(internal.F),
				scales.NewNote(internal.G),
				scales.NewNote(internal.AFlat),
				scales.NewNote(internal.BFlat),
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
			},
		},
		{
			name: internal.AFlat,
			key:  internal.AFlat,
			want: []scales.Note{
				scales.NewNote(internal.AFlat),
				scales.NewNote(internal.BFlat),
				scales.NewNote(internal.C),
				scales.NewNote(internal.DFlat),
				scales.NewNote(internal.EFlat),
				scales.NewNote(internal.F),
				scales.NewNote(internal.G),
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

func Test_MajorScaleGetChords(t *testing.T) {
	t.Parallel()

	scale, err := scales.NewMajorScale("C")
	require.NoError(t, err)

	var results []string
	for _, c := range scale.GetChords() {
		results = append(results, c.Description())
	}

	require.Equal(t, []string{"Cmaj", "Dm", "Em", "Fmaj", "Gmaj", "Amin", "Bdim"}, results)
}

// NewMajorScale demonstrates the use of NewMajorScale function.
func ExampleNewMajorScale() {
	scale, err := scales.NewMajorScale("C")
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Println(scale.GetNotes())
	// Output: [C D E F G A B]
}
