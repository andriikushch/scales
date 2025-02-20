package scales_test

import (
	"fmt"
	"testing"

	"github.com/andriikushch/scales/pkg/internal"

	"github.com/stretchr/testify/require"

	scales "github.com/andriikushch/scales/pkg"
)

func Test_NewNaturalMinorScale(t *testing.T) {
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
				scales.NewNote(internal.B),
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
				scales.NewNote(internal.E),
				scales.NewNote(internal.F),
				scales.NewNote(internal.G),
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
				scales.NewNote(internal.B),
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
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
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.G),
				scales.NewNote(internal.A),
			},
		},
		{
			name: internal.FSharp,
			key:  internal.FSharp,
			want: []scales.Note{
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.A),
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.D),
				scales.NewNote(internal.E),
			},
		},
		{
			name: internal.CSharp,
			key:  internal.CSharp,
			want: []scales.Note{
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.DSharp),
				scales.NewNote(internal.E),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.A),
				scales.NewNote(internal.B),
			},
		},
		{
			name: internal.GSharp,
			key:  internal.GSharp,
			want: []scales.Note{
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.ASharp),
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.DSharp),
				scales.NewNote(internal.E),
				scales.NewNote(internal.FSharp),
			},
		},
		{
			name: internal.DSharp,
			key:  internal.DSharp,
			want: []scales.Note{
				scales.NewNote(internal.DSharp),
				scales.NewNote(internal.ESharp),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.ASharp),
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
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
				scales.NewNote(internal.F),
				scales.NewNote(internal.GFlat),
				scales.NewNote(internal.AFlat),
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
				scales.NewNote(internal.C),
				scales.NewNote(internal.DFlat),
				scales.NewNote(internal.EFlat),
			},
		},
		{
			name: internal.C,
			key:  internal.C,
			want: []scales.Note{
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
				scales.NewNote(internal.EFlat),
				scales.NewNote(internal.F),
				scales.NewNote(internal.G),
				scales.NewNote(internal.AFlat),
				scales.NewNote(internal.BFlat),
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
				scales.NewNote(internal.D),
				scales.NewNote(internal.EFlat),
				scales.NewNote(internal.F),
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
				scales.NewNote(internal.A),
				scales.NewNote(internal.BFlat),
				scales.NewNote(internal.C),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			scale, err := scales.NewNaturalMinorScale(tt.key)

			require.NoError(t, err)
			require.Equal(t, tt.want, scale.GetNotes())
		})
	}
}

// NewNaturalMinorScale demonstrates the use of NewNaturalMinorScale function.
func ExampleNewNaturalMinorScale() {
	scale, err := scales.NewNaturalMinorScale("A")
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Println(scale.GetNotes())
	// Output: [{A} {B} {C} {D} {E} {F} {G}]
}

func Test_NewHarmonicMinorScale(t *testing.T) {
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
				scales.NewNote(internal.B),
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
				scales.NewNote(internal.E),
				scales.NewNote(internal.F),
				scales.NewNote(internal.GSharp),
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
				scales.NewNote(internal.B),
				scales.NewNote(internal.C),
				scales.NewNote(internal.DSharp),
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
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.G),
				scales.NewNote(internal.ASharp),
			},
		},
		{
			name: internal.FSharp,
			key:  internal.FSharp,
			want: []scales.Note{
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.A),
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.D),
				scales.NewNote(internal.ESharp),
			},
		},
		{
			name: internal.CSharp,
			key:  internal.CSharp,
			want: []scales.Note{
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.DSharp),
				scales.NewNote(internal.E),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.A),
				scales.NewNote(internal.BSharp),
			},
		},
		{
			name: internal.GSharp,
			key:  internal.GSharp,
			want: []scales.Note{
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.ASharp),
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.DSharp),
				scales.NewNote(internal.E),
				scales.NewNote(internal.FDoubleSharp),
			},
		},
		{
			name: internal.DSharp,
			key:  internal.DSharp,
			want: []scales.Note{
				scales.NewNote(internal.DSharp),
				scales.NewNote(internal.ESharp),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.ASharp),
				scales.NewNote(internal.B),
				scales.NewNote(internal.CDoubleSharp),
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
				scales.NewNote(internal.F),
				scales.NewNote(internal.GFlat),
				scales.NewNote(internal.A),
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
				scales.NewNote(internal.C),
				scales.NewNote(internal.DFlat),
				scales.NewNote(internal.E),
			},
		},
		{
			name: internal.C,
			key:  internal.C,
			want: []scales.Note{
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
				scales.NewNote(internal.EFlat),
				scales.NewNote(internal.F),
				scales.NewNote(internal.G),
				scales.NewNote(internal.AFlat),
				scales.NewNote(internal.B),
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
				scales.NewNote(internal.D),
				scales.NewNote(internal.EFlat),
				scales.NewNote(internal.FSharp),
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
				scales.NewNote(internal.A),
				scales.NewNote(internal.BFlat),
				scales.NewNote(internal.CSharp),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			scale, err := scales.NewHarmonicMinorScale(tt.key)

			require.NoError(t, err)
			require.Equal(t, tt.want, scale.GetNotes())
		})
	}
}

// NewHarmonicMinorScale demonstrates the use of NewHarmonicMinorScale function.
func ExampleNewHarmonicMinorScale() {
	scale, err := scales.NewHarmonicMinorScale("A")
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Println(scale.GetNotes())
	// Output: [{A} {B} {C} {D} {E} {F} {G#}]
}

func Test_NewMelodicMinorScale(t *testing.T) {
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
				scales.NewNote(internal.B),
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
				scales.NewNote(internal.E),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.GSharp),
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
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.DSharp),
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
				scales.NewNote(internal.A),
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
				scales.NewNote(internal.E),
				scales.NewNote(internal.FSharp),
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.ASharp),
				scales.NewNote(internal.BSharp),
			},
		},
		{
			name: internal.GSharp,
			key:  internal.GSharp,
			want: []scales.Note{
				scales.NewNote(internal.GSharp),
				scales.NewNote(internal.ASharp),
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
				scales.NewNote(internal.DSharp),
				scales.NewNote(internal.ESharp),
				scales.NewNote(internal.FDoubleSharp), ///
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
				scales.NewNote(internal.F),
				scales.NewNote(internal.G),
				scales.NewNote(internal.A),
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
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
				scales.NewNote(internal.E),
			},
		},
		{
			name: internal.C,
			key:  internal.C,
			want: []scales.Note{
				scales.NewNote(internal.C),
				scales.NewNote(internal.D),
				scales.NewNote(internal.EFlat),
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
				scales.NewNote(internal.BFlat),
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
				scales.NewNote(internal.F),
				scales.NewNote(internal.G),
				scales.NewNote(internal.A),
				scales.NewNote(internal.B),
				scales.NewNote(internal.CSharp),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			scale, err := scales.NewMelodicMinorScale(tt.key)

			require.NoError(t, err)
			require.Equal(t, tt.want, scale.GetNotes())
		})
	}
}

// NewMelodicMinorScale demonstrates the use of NewMelodicMinorScale function.
func ExampleNewMelodicMinorScale() {
	scale, err := scales.NewMelodicMinorScale("A")
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Println(scale.GetNotes())
	// Output: [{A} {B} {C} {D} {E} {F#} {G#}]
}
