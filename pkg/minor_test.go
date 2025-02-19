package scales_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	scales "github.com/andriikushch/scales/pkg"
	"github.com/andriikushch/scales/pkg/internal/chromatic"
)

func Test_NewNaturalMinorScale(t *testing.T) {
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
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
			},
		},
		{
			name: chromatic.E,
			key:  chromatic.E,
			want: []scales.Note{
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.D),
			},
		},
		{
			name: chromatic.B,
			key:  chromatic.B,
			want: []scales.Note{
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
			},
		},
		{
			name: chromatic.FSharp,
			key:  chromatic.FSharp,
			want: []scales.Note{
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.E),
			},
		},
		{
			name: chromatic.CSharp,
			key:  chromatic.CSharp,
			want: []scales.Note{
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.DSharp),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.B),
			},
		},
		{
			name: chromatic.GSharp,
			key:  chromatic.GSharp,
			want: []scales.Note{
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.ASharp),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.DSharp),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
			},
		},
		{
			name: chromatic.DSharp,
			key:  chromatic.DSharp,
			want: []scales.Note{
				scales.NewNote(chromatic.DSharp),
				scales.NewNote(chromatic.ESharp),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.ASharp),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
			},
		},
		{
			name: chromatic.BFlat,
			key:  chromatic.BFlat,
			want: []scales.Note{
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.DFlat),
				scales.NewNote(chromatic.EFlat),
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.GFlat),
				scales.NewNote(chromatic.AFlat),
			},
		},
		{
			name: chromatic.F,
			key:  chromatic.F,
			want: []scales.Note{
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.AFlat),
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.DFlat),
				scales.NewNote(chromatic.EFlat),
			},
		},
		{
			name: chromatic.C,
			key:  chromatic.C,
			want: []scales.Note{
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.EFlat),
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.AFlat),
				scales.NewNote(chromatic.BFlat),
			},
		},
		{
			name: chromatic.G,
			key:  chromatic.G,
			want: []scales.Note{
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.EFlat),
				scales.NewNote(chromatic.F),
			},
		},
		{
			name: chromatic.D,
			key:  chromatic.D,
			want: []scales.Note{
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.C),
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
			name: chromatic.A,
			key:  chromatic.A,
			want: []scales.Note{
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.GSharp),
			},
		},
		{
			name: chromatic.E,
			key:  chromatic.E,
			want: []scales.Note{
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.DSharp),
			},
		},
		{
			name: chromatic.B,
			key:  chromatic.B,
			want: []scales.Note{
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.ASharp),
			},
		},
		{
			name: chromatic.FSharp,
			key:  chromatic.FSharp,
			want: []scales.Note{
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.ESharp),
			},
		},
		{
			name: chromatic.CSharp,
			key:  chromatic.CSharp,
			want: []scales.Note{
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.DSharp),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.BSharp),
			},
		},
		{
			name: chromatic.GSharp,
			key:  chromatic.GSharp,
			want: []scales.Note{
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.ASharp),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.DSharp),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FDoubleSharp),
			},
		},
		{
			name: chromatic.DSharp,
			key:  chromatic.DSharp,
			want: []scales.Note{
				scales.NewNote(chromatic.DSharp),
				scales.NewNote(chromatic.ESharp),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.ASharp),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CDoubleSharp),
			},
		},
		{
			name: chromatic.BFlat,
			key:  chromatic.BFlat,
			want: []scales.Note{
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.DFlat),
				scales.NewNote(chromatic.EFlat),
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.GFlat),
				scales.NewNote(chromatic.A),
			},
		},
		{
			name: chromatic.F,
			key:  chromatic.F,
			want: []scales.Note{
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.AFlat),
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.DFlat),
				scales.NewNote(chromatic.E),
			},
		},
		{
			name: chromatic.C,
			key:  chromatic.C,
			want: []scales.Note{
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.EFlat),
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.AFlat),
				scales.NewNote(chromatic.B),
			},
		},
		{
			name: chromatic.G,
			key:  chromatic.G,
			want: []scales.Note{
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.EFlat),
				scales.NewNote(chromatic.FSharp),
			},
		},
		{
			name: chromatic.D,
			key:  chromatic.D,
			want: []scales.Note{
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.CSharp),
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
			name: chromatic.A,
			key:  chromatic.A,
			want: []scales.Note{
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.GSharp),
			},
		},
		{
			name: chromatic.E,
			key:  chromatic.E,
			want: []scales.Note{
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.DSharp),
			},
		},
		{
			name: chromatic.B,
			key:  chromatic.B,
			want: []scales.Note{
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.D),
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
				scales.NewNote(chromatic.A),
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
				scales.NewNote(chromatic.E),
				scales.NewNote(chromatic.FSharp),
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.ASharp),
				scales.NewNote(chromatic.BSharp),
			},
		},
		{
			name: chromatic.GSharp,
			key:  chromatic.GSharp,
			want: []scales.Note{
				scales.NewNote(chromatic.GSharp),
				scales.NewNote(chromatic.ASharp),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
				scales.NewNote(chromatic.DSharp),
				scales.NewNote(chromatic.ESharp),
				scales.NewNote(chromatic.FDoubleSharp), ///
			},
		},
		{
			name: chromatic.BFlat,
			key:  chromatic.BFlat,
			want: []scales.Note{
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.DFlat),
				scales.NewNote(chromatic.EFlat),
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
			},
		},
		{
			name: chromatic.F,
			key:  chromatic.F,
			want: []scales.Note{
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.AFlat),
				scales.NewNote(chromatic.BFlat),
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.E),
			},
		},
		{
			name: chromatic.C,
			key:  chromatic.C,
			want: []scales.Note{
				scales.NewNote(chromatic.C),
				scales.NewNote(chromatic.D),
				scales.NewNote(chromatic.EFlat),
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
				scales.NewNote(chromatic.BFlat),
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
				scales.NewNote(chromatic.F),
				scales.NewNote(chromatic.G),
				scales.NewNote(chromatic.A),
				scales.NewNote(chromatic.B),
				scales.NewNote(chromatic.CSharp),
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
