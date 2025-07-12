package scales_test

import (
	"fmt"
	"testing"

	"github.com/andriikushch/scales/pkg/internal"

	"github.com/stretchr/testify/require"

	scales "github.com/andriikushch/scales/pkg"
)

func Test_NewMajorPentatonicScale(t *testing.T) {
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
				scales.NewNote(internal.G),
				scales.NewNote(internal.A),
			},
		},
		{
			name: internal.G,
			key:  internal.G,
			want: []scales.Note{
				scales.NewNote(internal.G),
				scales.NewNote(internal.A),
				scales.NewNote(internal.B),
				scales.NewNote(internal.D),
				scales.NewNote(internal.E),
			},
		},
		{
			name: internal.D,
			key:  internal.D,
			want: []scales.Note{
				scales.NewNote(internal.D),
				scales.NewNote(internal.E),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.A),
				scales.NewNote(internal.B),
			},
		},
		{
			name: internal.A,
			key:  internal.A,
			want: []scales.Note{
				scales.NewNote(internal.A),
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.E),
				scales.NewNote(internal.FSharp),
			},
		},
		{
			name: internal.E,
			key:  internal.E,
			want: []scales.Note{
				scales.NewNote(internal.E),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
			},
		},
		{
			name: internal.B,
			key:  internal.B,
			want: []scales.Note{
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.DSharp),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.GSharp),
			},
		},
		{
			name: internal.FSharp,
			key:  internal.FSharp,
			want: []scales.Note{
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.ASharp),
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.DSharp),
			},
		},
		{
			name: internal.CSharp,
			key:  internal.CSharp,
			want: []scales.Note{
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.DSharp),
				scales.NewNote(internal.ESharp), // E# = F
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.ASharp),
			},
		},
		{
			name: internal.F,
			key:  internal.F,
			want: []scales.Note{
				scales.NewNote(internal.F),
				scales.NewNote(internal.G),
				scales.NewNote(internal.A),
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
			},
		},
		{
			name: internal.BFlat,
			key:  internal.BFlat,
			want: []scales.Note{
				scales.NewNote(internal.BFlat),
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
				scales.NewNote(internal.F),
				scales.NewNote(internal.G),
			},
		},
		{
			name: internal.EFlat,
			key:  internal.EFlat,
			want: []scales.Note{
				scales.NewNote(internal.EFlat),
				scales.NewNote(internal.F),
				scales.NewNote(internal.G),
				scales.NewNote(internal.BFlat),
				scales.NewNote(internal.C),
			},
		},
		{
			name: internal.AFlat,
			key:  internal.AFlat,
			want: []scales.Note{
				scales.NewNote(internal.AFlat),
				scales.NewNote(internal.BFlat),
				scales.NewNote(internal.C),
				scales.NewNote(internal.EFlat),
				scales.NewNote(internal.F),
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

// NewMajorPentatonicScale demonstrates the use of NewMajorPentatonicScale function.
func ExampleNewMajorPentatonicScale() {
	scale, err := scales.NewMajorPentatonicScale("A")
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Println(scale.GetNotes())
	// Output: [A B C# E F#]
}

func Test_NewMinorPentatonicScale(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		key  string
		want []scales.Note
	}{
		{
			name: internal.A,
			key:  internal.A,
			want: []scales.Note{
				scales.NewNote(internal.A),
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
				scales.NewNote(internal.E),
				scales.NewNote(internal.G),
			},
		},
		{
			name: internal.E,
			key:  internal.E,
			want: []scales.Note{
				scales.NewNote(internal.E),
				scales.NewNote(internal.G),
				scales.NewNote(internal.A),
				scales.NewNote(internal.B),
				scales.NewNote(internal.D),
			},
		},
		{
			name: internal.B,
			key:  internal.B,
			want: []scales.Note{
				scales.NewNote(internal.B),
				scales.NewNote(internal.D),
				scales.NewNote(internal.E),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.A),
			},
		},
		{
			name: internal.FSharp,
			key:  internal.FSharp,
			want: []scales.Note{
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.A),
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.E),
			},
		},
		{
			name: internal.CSharp,
			key:  internal.CSharp,
			want: []scales.Note{
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.E),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.B),
			},
		},
		{
			name: internal.GSharp,
			key:  internal.GSharp,
			want: []scales.Note{
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.DSharp),
				scales.NewNote(internal.FSharp),
			},
		},
		{
			name: internal.DSharp,
			key:  internal.DSharp,
			want: []scales.Note{
				scales.NewNote(internal.DSharp),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.ASharp),
				scales.NewNote(internal.CSharp),
			},
		},
		{
			name: internal.BFlat,
			key:  internal.BFlat,
			want: []scales.Note{
				scales.NewNote(internal.BFlat),
				scales.NewNote(internal.DFlat),
				scales.NewNote(internal.EFlat),
				scales.NewNote(internal.F),
				scales.NewNote(internal.AFlat),
			},
		},
		{
			name: internal.F,
			key:  internal.F,
			want: []scales.Note{
				scales.NewNote(internal.F),
				scales.NewNote(internal.AFlat),
				scales.NewNote(internal.BFlat),
				scales.NewNote(internal.C),
				scales.NewNote(internal.EFlat),
			},
		},
		{
			name: internal.C,
			key:  internal.C,
			want: []scales.Note{
				scales.NewNote(internal.C),
				scales.NewNote(internal.EFlat),
				scales.NewNote(internal.F),
				scales.NewNote(internal.G),
				scales.NewNote(internal.BFlat),
			},
		},
		{
			name: internal.G,
			key:  internal.G,
			want: []scales.Note{
				scales.NewNote(internal.G),
				scales.NewNote(internal.BFlat),
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
				scales.NewNote(internal.F),
			},
		},
		{
			name: internal.D,
			key:  internal.D,
			want: []scales.Note{
				scales.NewNote(internal.D),
				scales.NewNote(internal.F),
				scales.NewNote(internal.G),
				scales.NewNote(internal.A),
				scales.NewNote(internal.C),
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

// NewMinorPentatonicScale demonstrates the use of NewMinorPentatonicScale function.
func ExampleNewMinorPentatonicScale() {
	scale, err := scales.NewMinorPentatonicScale("A")
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Println(scale.GetNotes())
	// Output: [A C D E G]
}
