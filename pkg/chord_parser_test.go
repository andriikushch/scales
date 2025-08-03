package scales

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/andriikushch/scales/pkg/internal"
)

func TestParseChord(t *testing.T) {
	t.Parallel()

	type args struct {
		description string
	}

	tests := []struct {
		name string
		args args
		want Chord
	}{
		{
			name: "Am",
			args: args{
				description: "Am",
			},
			want: Chord{
				description:    "Am",
				chordBasicType: internal.Minor,
				root:           NewNote("A"),
				notes: []Note{
					NewNote(internal.A),
					NewNote(internal.C),
					NewNote(internal.E),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.IP5},
				quality:   []string{internal.Minor},
			},
		},
		{
			name: "Am/C",
			args: args{
				description: "Am/C",
			},
			want: Chord{
				description:    "Am/C",
				chordBasicType: internal.Minor,
				bassNote:       NewNote("C"),
				root:           NewNote("A"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.A),
					NewNote(internal.E),
				},
				structure: []int{internal.Im3, internal.IUnison, internal.IP5},
				quality:   []string{internal.Minor, internal.Over, "C"},
			},
		},
		{
			name: "Am7",
			args: args{
				description: "Am7",
			},
			want: Chord{
				description:    "Am7",
				chordBasicType: internal.Minor,
				root:           NewNote("A"),
				notes: []Note{
					NewNote(internal.A),
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.IP5, internal.Im7},
				quality:   []string{internal.Minor, internal.Seventh},
			},
		},
		{
			name: "C",
			args: args{
				description: "C",
			},
			want: Chord{
				description:    "C",
				chordBasicType: internal.Major,
				root:           NewNote(internal.C),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5},
				quality:   []string{internal.Major},
			},
		},
		{
			name: "Cmaj",
			args: args{
				description: "Cmaj",
			},
			want: Chord{
				description:    "Cmaj",
				chordBasicType: internal.Major,
				root:           NewNote(internal.C),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5},
				quality:   []string{internal.Major},
			},
		},
		{
			name: "Cmaj7",
			args: args{
				description: "Cmaj7",
			},
			want: Chord{
				description:    "Cmaj7",
				chordBasicType: internal.Major,
				root:           NewNote(internal.C),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.B),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM7},
				quality:   []string{internal.Major, internal.Seventh},
			},
		},
		{
			name: "C7",
			args: args{
				description: "C7",
			},
			want: Chord{
				description:    "C7",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7},
				quality:   []string{internal.Dominant, internal.Seventh},
			},
		},
		{
			name: "Cm7",
			args: args{
				description: "Cm7",
			},
			want: Chord{
				description:    "Cm7",
				chordBasicType: internal.Minor,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.G),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.IP5, internal.Im7},
				quality:   []string{internal.Minor, internal.Seventh},
			},
		},
		{
			name: "Cm7b5",
			args: args{
				description: "Cm7b5",
			},
			want: Chord{
				description:    "Cm7b5",
				chordBasicType: internal.Minor,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.ID5, internal.Im7},
				quality:   []string{internal.Minor, internal.Seven, internal.Flat, internal.Five},
			},
		},
		{
			name: "Cdim7",
			args: args{
				description: "Cdim7",
			},
			want: Chord{
				description:    "Cdim7",
				chordBasicType: internal.Diminished,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GFlat),
					NewNote(internal.A),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.ID5, internal.IM6},
				quality:   []string{internal.Diminished, internal.Seventh},
			},
		},
		{
			name: "Cdim",
			args: args{
				description: "Cdim",
			},
			want: Chord{
				description:    "Cdim",
				chordBasicType: internal.Diminished,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GFlat),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.ID5},
				quality:   []string{internal.Diminished},
			},
		},
		{
			name: "Caug7",
			args: args{
				description: "Caug7",
			},
			want: Chord{
				description:    "Caug7",
				chordBasicType: internal.Augmented,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.Im7},
				quality:   []string{internal.Augmented, internal.Seventh},
			},
		},
		{
			name: "Caug9",
			args: args{
				description: "Caug9",
			},
			want: Chord{
				description:    "Caug9",
				chordBasicType: internal.Augmented,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.BFlat),
					NewNote(internal.D),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.Im7, internal.IM9},
				quality:   []string{internal.Augmented, internal.Ninth},
			},
		},
		{
			name: "Caug11",
			args: args{
				description: "Caug11",
			},
			want: Chord{
				description:    "Caug11",
				chordBasicType: internal.Augmented,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.BFlat),
					NewNote(internal.D),
					NewNote(internal.F),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.Im7, internal.IM9, internal.IP11},
				quality:   []string{internal.Augmented, internal.Eleventh},
			},
		},
		{
			name: "Caug6",
			args: args{
				description: "Caug6",
			},
			want: Chord{
				description:    "Caug6",
				chordBasicType: internal.Augmented,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.A),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.IM6},
				quality:   []string{internal.Augmented, internal.Sixth},
			},
		},
		{
			name: "C7b5",
			args: args{
				description: "C7b5",
			},
			want: Chord{
				description:    "C7b5",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.ID5, internal.Im7},
				quality:   []string{internal.Dominant, internal.Seven, internal.Flat, internal.Five},
			},
		},
		{
			name: "Cm(maj7)",
			args: args{
				description: "Cm(maj7)",
			},
			want: Chord{
				description:    "Cm(maj7)",
				chordBasicType: internal.Minor,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.G),
					NewNote(internal.B),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.IP5, internal.IM7},
				quality:   []string{internal.Minor, internal.Major, internal.Seventh},
			},
		},
		{
			name: "Cmaj9",
			args: args{
				description: "Cmaj9",
			},
			want: Chord{
				description:    "Cmaj9",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.B),
					NewNote(internal.D),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM7, internal.IM9},
				quality:   []string{internal.Major, internal.Ninth},
			},
		},
		{
			name: "C9",
			args: args{
				description: "C9",
			},
			want: Chord{
				description:    "C9",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.D),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.IM9},
				quality:   []string{internal.Dominant, internal.Ninth},
			},
		},
		{
			name: "Cm9",
			args: args{
				description: "Cm9",
			},
			want: Chord{
				description:    "Cm9",
				chordBasicType: internal.Minor,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.D),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.IP5, internal.Im7, internal.IM9},
				quality:   []string{internal.Minor, internal.Ninth},
			},
		},
		{
			name: "C11",
			args: args{
				description: "C11",
			},
			want: Chord{
				description:    "C11",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.D),
					NewNote(internal.F),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.IM9, internal.IP11},
				quality:   []string{internal.Dominant, internal.Eleventh},
			},
		},
		{
			name: "Cm11",
			args: args{
				description: "Cm11",
			},
			want: Chord{
				description:    "Cm11",
				chordBasicType: internal.Minor,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.D),
					NewNote(internal.F),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.IP5, internal.Im7, internal.IM9, internal.IP11},
				quality:   []string{internal.Minor, internal.Eleventh},
			},
		},
		{
			name: "Cmaj13",
			args: args{
				description: "Cmaj13",
			},
			want: Chord{
				description:    "Cmaj13",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.B),
					NewNote(internal.D),
					NewNote(internal.F),
					NewNote(internal.A),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM7, internal.IM9, internal.IP11, internal.IM13},
				quality:   []string{internal.Major, internal.Thirteenth},
			},
		},
		{
			name: "C13",
			args: args{
				description: "C13",
			},
			want: Chord{
				description:    "C13",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.D),
					NewNote(internal.F),
					NewNote(internal.A),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.IM9, internal.IP11, internal.IM13},
				quality:   []string{internal.Dominant, internal.Thirteenth},
			},
		},
		{
			name: "Cm13",
			args: args{
				description: "Cm13",
			},
			want: Chord{
				description:    "Cm13",
				chordBasicType: internal.Minor,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.D),
					NewNote(internal.F),
					NewNote(internal.A),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.IP5, internal.Im7, internal.IM9, internal.IP11, internal.IM13},
				quality:   []string{internal.Minor, internal.Thirteenth},
			},
		},
		{
			name: "C7b5",
			args: args{
				description: "C7b5",
			},
			want: Chord{
				description:    "C7b5",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.ID5, internal.Im7},
				quality:   []string{internal.Dominant, internal.Seven, internal.Flat, internal.Five},
			},
		},
		{
			name: "C7b9",
			args: args{
				description: "C7b9",
			},
			want: Chord{
				description:    "C7b9",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.DFlat),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.Im9},
				quality:   []string{internal.Dominant, internal.Seven, internal.Flat, internal.Nine},
			},
		},
		{
			name: "C7#9",
			args: args{
				description: "C7#9",
			},
			want: Chord{
				description:    "C7#9",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.DSharp),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.Im10},
				quality:   []string{internal.Dominant, internal.Seven, internal.Sharp, internal.Nine},
			},
		},
		{
			name: "C7b5b9",
			args: args{
				description: "C7b5b9",
			},
			want: Chord{
				description:    "C7b5b9",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
					NewNote(internal.DFlat),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.ID5, internal.Im7, internal.Im9},
				quality:   []string{internal.Dominant, internal.Seven, internal.Flat, internal.Five, internal.Flat, internal.Nine},
			},
		},
		{
			name: "C7b5(b9)",
			args: args{
				description: "C7b5(b9)",
			},
			want: Chord{
				description:    "C7b5(b9)",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
					NewNote(internal.DFlat),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.ID5, internal.Im7, internal.Im9},
				quality:   []string{internal.Dominant, internal.Seven, internal.Flat, internal.Five, internal.Flat, internal.Nine},
			},
		},
		{
			name: "C7#5#9",
			args: args{
				description: "C7#5#9",
			},
			want: Chord{
				description:    "C7#5#9",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.BFlat),
					NewNote(internal.DSharp),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.Im7, internal.Im10},
				quality:   []string{internal.Dominant, internal.Seven, internal.Sharp, internal.Five, internal.Sharp, internal.Nine},
			},
		},
		{
			name: "C7#5(#9)",
			args: args{
				description: "C7#5(#9)",
			},
			want: Chord{
				description:    "C7#5(#9)",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.BFlat),
					NewNote(internal.DSharp),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.Im7, internal.Im10},
				quality:   []string{internal.Dominant, internal.Seven, internal.Sharp, internal.Five, internal.Sharp, internal.Nine},
			},
		},
		{
			name: "C7b5#9",
			args: args{
				description: "C7b5#9",
			},
			want: Chord{
				description:    "C7b5#9",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
					NewNote(internal.DSharp),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.ID5, internal.Im7, internal.Im10},
				quality:   []string{internal.Dominant, internal.Seven, internal.Flat, internal.Five, internal.Sharp, internal.Nine},
			},
		},
		{
			name: "C7b5(#9)",
			args: args{
				description: "C7b5(#9)",
			},
			want: Chord{
				description:    "C7b5(#9)",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
					NewNote(internal.DSharp),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.ID5, internal.Im7, internal.Im10},
				quality:   []string{internal.Dominant, internal.Seven, internal.Flat, internal.Five, internal.Sharp, internal.Nine},
			},
		},
		{
			name: "C7#5(b9)",
			args: args{
				description: "C7#5(b9)",
			},
			want: Chord{
				description:    "C7#5(b9)",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.BFlat),
					NewNote(internal.DFlat),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.Im7, internal.Im9},
				quality:   []string{internal.Dominant, internal.Seven, internal.Sharp, internal.Five, internal.Flat, internal.Nine},
			},
		},
		{
			name: "C7#11",
			args: args{
				description: "C7#11",
			},
			want: Chord{
				description:    "C7#11",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.FSharp),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.IA11},
				quality:   []string{internal.Dominant, internal.Seven, internal.Sharp, internal.Eleven},
			},
		},
		{
			name: "C9b5",
			args: args{
				description: "C9b5",
			},
			want: Chord{
				description:    "C9b5",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
					NewNote(internal.D),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.ID5, internal.Im7, internal.IM9},
				quality:   []string{internal.Dominant, internal.Nine, internal.Flat, internal.Five},
			},
		},
		{
			name: "C9#5",
			args: args{
				description: "C9#5",
			},
			want: Chord{
				description:    "C9#5",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.BFlat),
					NewNote(internal.D),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.Im7, internal.IM9},
				quality:   []string{internal.Dominant, internal.Nine, internal.Sharp, internal.Five},
			},
		},
		{
			name: "C9#11",
			args: args{
				description: "C9#11",
			},
			want: Chord{
				description:    "C9#11",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.D),
					NewNote(internal.FSharp),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.IM9, internal.IA11},
				quality:   []string{internal.Dominant, internal.Nine, internal.Sharp, internal.Eleven},
			},
		},
		{
			name: "C13b9",
			args: args{
				description: "C13b9",
			},
			want: Chord{
				description:    "C13b9",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.DFlat),
					NewNote(internal.F),
					NewNote(internal.A),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.Im9, internal.IP11, internal.IM13},
				quality:   []string{internal.Dominant, internal.Thirteen, internal.Flat, internal.Nine},
			},
		},
		{
			name: "C13#9",
			args: args{
				description: "C13#9",
			},
			want: Chord{
				description:    "C13#9",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.DSharp),
					NewNote(internal.F),
					NewNote(internal.A),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.IA9, internal.IP11, internal.IM13},
				quality:   []string{internal.Dominant, internal.Thirteen, internal.Sharp, internal.Nine},
			},
		},
		{
			name: "C13#11",
			args: args{
				description: "C13#11",
			},
			want: Chord{
				description:    "C13#11",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.D),
					NewNote(internal.FSharp),
					NewNote(internal.A),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.IM9, internal.IA11, internal.IM13},
				quality:   []string{internal.Dominant, internal.Thirteen, internal.Sharp, internal.Eleven},
			},
		},
		{
			name: "Cm7b5",
			args: args{
				description: "Cm7b5",
			},
			want: Chord{
				description:    "Cm7b5",
				chordBasicType: internal.Minor,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.ID5, internal.Im7},
				quality:   []string{internal.Minor, internal.Seven, internal.Flat, internal.Five},
			},
		},
		{
			name: "Cm7#5",
			args: args{
				description: "Cm7#5",
			},
			want: Chord{
				description:    "Cm7#5",
				chordBasicType: internal.Minor,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GSharp),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.IA5, internal.Im7},
				quality:   []string{internal.Minor, internal.Seven, internal.Sharp, internal.Five},
			},
		},
		{
			name: "Cmaj7b5",
			args: args{
				description: "Cmaj7b5",
			},
			want: Chord{
				description:    "Cmaj7b5",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GFlat),
					NewNote(internal.B),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.ID5, internal.IM7},
				quality:   []string{internal.Major, internal.Seven, internal.Flat, internal.Five},
			},
		},
		{
			name: "Cmaj7#5",
			args: args{
				description: "Cmaj7#5",
			},
			want: Chord{
				description:    "Cmaj7#5",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.B),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.IM7},
				quality:   []string{internal.Major, internal.Seven, internal.Sharp, internal.Five},
			},
		},
		{
			name: "Cmaj7#11",
			args: args{
				description: "Cmaj7#11",
			},
			want: Chord{
				description:    "Cmaj7#11",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.B),
					NewNote(internal.FSharp),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM7, internal.IA11},
				quality:   []string{internal.Major, internal.Seven, internal.Sharp, internal.Eleven},
			},
		},
		{
			name: "Cmaj9#11",
			args: args{
				description: "Cmaj9#11",
			},
			want: Chord{
				description:    "Cmaj9#11",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.B),
					NewNote(internal.D),
					NewNote(internal.FSharp),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM7, internal.IM9, internal.IA11},
				quality:   []string{internal.Major, internal.Nine, internal.Sharp, internal.Eleven},
			},
		},
		{
			name: "C6",
			args: args{
				description: "C6",
			},
			want: Chord{
				description:    "C6",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.A),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM6},
				quality:   []string{internal.Sixth},
			},
		},
		{
			name: "C6/9",
			args: args{
				description: "C6/9",
			},
			want: Chord{
				description:    "C6/9",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.A),
					NewNote(internal.D),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM6, internal.IM9},
				quality:   []string{internal.Sixth, internal.Added, internal.Ninth},
			},
		},
		{
			name: "C6add9",
			args: args{
				description: "C6add9",
			},
			want: Chord{
				description:    "C6add9",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.A),
					NewNote(internal.D),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM6, internal.IM9},
				quality:   []string{internal.Sixth, internal.Added, internal.Ninth},
			},
		},
		{
			name: "C6(add9)",
			args: args{
				description: "C6(add9)",
			},
			want: Chord{
				description:    "C6(add9)",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.A),
					NewNote(internal.D),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM6, internal.IM9},
				quality:   []string{internal.Sixth, internal.Added, internal.Ninth},
			},
		},
		{
			name: "C/G",
			args: args{
				description: "C/G",
			},
			want: Chord{
				description:    "C/G",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				bassNote:       NewNote("G"),
				notes: []Note{
					NewNote(internal.G),
					NewNote(internal.C),
					NewNote(internal.E),
				},
				structure: []int{internal.IP5, internal.IUnison, internal.IM3},
				quality:   []string{internal.Major, internal.Over, "G"},
			},
		},
		{
			name: "C/G#",
			args: args{
				description: "C/G#",
			},
			want: Chord{
				description:    "C/G#",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				bassNote:       NewNote("G#"),
				notes: []Note{
					NewNote(internal.GSharp),
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
				},
				structure: []int{internal.IA5, internal.IUnison, internal.IM3, internal.IP5},
				quality:   []string{internal.Major, internal.Over, "G#"},
			},
		},
		{
			name: "Csus2",
			args: args{
				description: "Csus2",
			},
			want: Chord{
				description:    "Csus2",
				chordBasicType: internal.Suspended,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.D),
					NewNote(internal.G),
				},
				structure: []int{internal.IUnison, internal.IM2, internal.IP5},
				quality:   []string{internal.Suspended, internal.Second},
			},
		},
		{
			name: "Csus4",
			args: args{
				description: "Csus4",
			},
			want: Chord{
				description:    "Csus4",
				chordBasicType: internal.Suspended,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.F),
					NewNote(internal.G),
				},
				structure: []int{internal.IUnison, internal.IP4, internal.IP5},
				quality:   []string{internal.Suspended, internal.Fourth},
			},
		},
		{
			name: "Csus",
			args: args{
				description: "Csus",
			},
			want: Chord{
				description:    "Csus",
				chordBasicType: internal.Suspended,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.F),
					NewNote(internal.G),
				},
				structure: []int{internal.IUnison, internal.IP4, internal.IP5},
				quality:   []string{internal.Suspended, internal.Fourth},
			},
		},
		{
			name: "C7sus4",
			args: args{
				description: "C7sus4",
			},
			want: Chord{
				description:    "C7sus4",
				chordBasicType: internal.Suspended,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.F),
					NewNote(internal.G),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.IP4, internal.IP5, internal.Im7},
				quality:   []string{internal.Dominant, internal.Seventh, internal.Suspended, internal.Fourth},
			},
		},
		{
			name: "C7sus",
			args: args{
				description: "C7sus",
			},
			want: Chord{
				description:    "C7sus",
				chordBasicType: internal.Suspended,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.F),
					NewNote(internal.G),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.IP4, internal.IP5, internal.Im7},
				quality:   []string{internal.Dominant, internal.Seventh, internal.Suspended, internal.Fourth},
			},
		},
		{
			name: "Cadd9",
			args: args{
				description: "Cadd9",
			},
			want: Chord{
				description:    "Cadd9",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.D),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM9},
				quality:   []string{internal.Added, internal.Ninth},
			},
		},
		{
			name: "C(add9)",
			args: args{
				description: "C(add9)",
			},
			want: Chord{
				description:    "C(add9)",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.D),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM9},
				quality:   []string{internal.Added, internal.Ninth},
			},
		},
		{
			name: "Cm(add9)",
			args: args{
				description: "Cm(add9)",
			},
			want: Chord{
				description:    "Cm(add9)",
				chordBasicType: internal.Minor,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.G),
					NewNote(internal.D),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.IP5, internal.IM9},
				quality:   []string{internal.Minor, internal.Added, internal.Ninth},
			},
		},
		{
			name: "Cdim6",
			args: args{
				description: "Cdim6",
			},
			want: Chord{
				description:    "Cdim6",
				chordBasicType: internal.Diminished,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GFlat),
					NewNote(internal.AFlat),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.ID5, internal.Im6},
				quality:   []string{internal.Diminished, internal.Sixth},
			},
		},
		{
			name: "Cdim9",
			args: args{
				description: "Cdim9",
			},
			want: Chord{
				description:    "Cdim9",
				chordBasicType: internal.Diminished,
				root:           NewNote("C"),
				notes: []Note{
					// C, E♭, G♭, B♭♭, D
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GFlat),
					NewNote(internal.A),
					NewNote(internal.D),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.ID5, internal.IM6, internal.IM9},
				quality:   []string{internal.Diminished, internal.Ninth},
			},
		},
		{
			name: "Cdim11",
			args: args{
				description: "Cdim11",
			},
			want: Chord{
				description:    "Cdim11",
				chordBasicType: internal.Diminished,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GFlat),
					NewNote(internal.A),
					NewNote(internal.D),
					NewNote(internal.F),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.ID5, internal.IM6, internal.IM9, internal.IP11},
				quality:   []string{internal.Diminished, internal.Eleventh},
			},
		},
		{
			name: "G#m",
			args: args{
				description: "G#m",
			},
			want: Chord{
				description:    "G#m",
				chordBasicType: internal.Minor,
				root:           NewNote("G#"),
				notes: []Note{
					NewNote(internal.GSharp),
					NewNote(internal.B),
					NewNote(internal.DSharp),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.IP5},
				quality:   []string{internal.Minor},
			},
		},
		{
			name: "C#dim11",
			args: args{
				description: "C#dim11",
			},
			want: Chord{
				description:    "C#dim11",
				chordBasicType: internal.Diminished,
				root:           NewNote("C#"),
				notes: []Note{
					NewNote(internal.CSharp),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.DSharp),
					NewNote(internal.FSharp),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.ID5, internal.IM6, internal.IM9, internal.IP11},
				quality:   []string{internal.Diminished, internal.Eleventh},
			},
		},
		{
			name: "Ebaug",
			args: args{
				description: "Ebaug",
			},
			want: Chord{
				description:    "Ebaug",
				chordBasicType: internal.Augmented,
				root:           NewNote("Eb"),
				notes: []Note{
					NewNote(internal.EFlat),
					NewNote(internal.G),
					NewNote(internal.B),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IA5},
				quality:   []string{internal.Augmented},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewChord(tt.args.description)
			if err != nil {
				t.Errorf("parseChord() error = %v", err)
			}

			if "true" == os.Getenv("SCALES_TEST_COLLECT_CHORD_QUALITIES") {
				qualityLine := strings.TrimPrefix(got.description, got.root.String())

				extraBassNote, err := regexp.MatchString(`/[a-zA-Z]`, qualityLine)
				require.NoError(t, err)
				if !extraBassNote {
					f, err := os.OpenFile("../build/output.txt", os.O_WRONLY|os.O_SYNC|os.O_CREATE|os.O_APPEND, 0o777)
					require.NoError(t, err)

					structureAsLine := ""
					for i, v := range got.structure {
						if i != 0 {
							structureAsLine += "-"
						}
						structureAsLine += strconv.Itoa(v)
					}

					_, err = f.WriteString(fmt.Sprintf("%s;%s\n", structureAsLine, qualityLine))
					require.NoError(t, err)
					require.NoError(t, f.Close())
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseChord() = %v, want %v", got, tt.want)
			}
		})
	}
}
