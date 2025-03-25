package scales

import (
	"github.com/andriikushch/scales/pkg/internal"
	"reflect"
	"testing"
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
				name:           "Am",
				chordBasicType: internal.Minor,
				root:           NewNote("A"),
				notes: []Note{
					NewNote(internal.A),
					NewNote(internal.C),
					NewNote(internal.E),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.IP5},
				cType:     []string{internal.Minor},
			},
		},
		{
			name: "Am/C",
			args: args{
				description: "Am/C",
			},
			want: Chord{
				name:           "Am/C",
				chordBasicType: internal.Minor,
				bassNote:       NewNote("C"),
				root:           NewNote("A"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.A),
					NewNote(internal.E),
				},
				structure: []int{internal.Im3, internal.IUnison, internal.IP5},
				cType:     []string{internal.Minor, internal.Over, "C"},
			},
		},
		{
			name: "Am7",
			args: args{
				description: "Am7",
			},
			want: Chord{
				name:           "Am7",
				chordBasicType: internal.Minor,
				root:           NewNote("A"),
				notes: []Note{
					NewNote(internal.A),
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.IP5, internal.Im7},
				cType:     []string{internal.Minor, internal.Seventh},
			},
		},
		{
			name: "C",
			args: args{
				description: "C",
			},
			want: Chord{
				name:           "C",
				chordBasicType: internal.Major,
				root:           NewNote(internal.C),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5},
				cType:     []string{internal.Major},
			},
		},
		{
			name: "Cmaj",
			args: args{
				description: "Cmaj",
			},
			want: Chord{
				name:           "Cmaj",
				chordBasicType: internal.Major,
				root:           NewNote(internal.C),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5},
				cType:     []string{internal.Major},
			},
		},
		{
			name: "Cmaj7",
			args: args{
				description: "Cmaj7",
			},
			want: Chord{
				name:           "Cmaj7",
				chordBasicType: internal.Major,
				root:           NewNote(internal.C),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.B),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM7},
				cType:     []string{internal.Major, internal.Seventh},
			},
		},
		{
			name: "C7",
			args: args{
				description: "C7",
			},
			want: Chord{
				name:           "C7",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7},
				cType:     []string{internal.Dominant, internal.Seventh},
			},
		},
		{
			name: "Cm7",
			args: args{
				description: "Cm7",
			},
			want: Chord{
				name:           "Cm7",
				chordBasicType: internal.Minor,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.G),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.IP5, internal.Im7},
				cType:     []string{internal.Minor, internal.Seventh},
			},
		},
		{
			name: "Cm7b5",
			args: args{
				description: "Cm7b5",
			},
			want: Chord{
				name:           "Cm7b5",
				chordBasicType: internal.Minor,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.ID5, internal.Im7},
				cType:     []string{internal.Minor, internal.Seven, internal.Flat, internal.Five},
			},
		},
		{
			name: "Cdim7",
			args: args{
				description: "Cdim7",
			},
			want: Chord{
				name:           "Cdim7",
				chordBasicType: internal.Diminished,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GFlat),
					NewNote(internal.A),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.ID5, internal.IM6},
				cType:     []string{internal.Diminished, internal.Seventh},
			},
		},
		{
			name: "Caug7",
			args: args{
				description: "Caug7",
			},
			want: Chord{
				name:           "Caug7",
				chordBasicType: internal.Augmented,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.Im7},
				cType:     []string{internal.Augmented, internal.Seventh},
			},
		},
		{
			name: "Caug9",
			args: args{
				description: "Caug9",
			},
			want: Chord{
				name:           "Caug9",
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
				cType:     []string{internal.Augmented, internal.Ninth},
			},
		},
		{
			name: "Caug11",
			args: args{
				description: "Caug11",
			},
			want: Chord{
				name:           "Caug11",
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
				cType:     []string{internal.Augmented, internal.Eleventh},
			},
		},
		{
			name: "Caug6",
			args: args{
				description: "Caug6",
			},
			want: Chord{
				name:           "Caug6",
				chordBasicType: internal.Augmented,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.A),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.IM6},
				cType:     []string{internal.Augmented, internal.Sixth},
			},
		},
		{
			name: "C7b5",
			args: args{
				description: "C7b5",
			},
			want: Chord{
				name:           "C7b5",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.ID5, internal.Im7},
				cType:     []string{internal.Dominant, internal.Seven, internal.Flat, internal.Five},
			},
		},
		{
			name: "Cm(maj7)",
			args: args{
				description: "Cm(maj7)",
			},
			want: Chord{
				name:           "Cm(maj7)",
				chordBasicType: internal.Minor,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.G),
					NewNote(internal.B),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.IP5, internal.IM7},
				cType:     []string{internal.Minor, internal.Major, internal.Seventh},
			},
		},
		{
			name: "Cmaj9",
			args: args{
				description: "Cmaj9",
			},
			want: Chord{
				name:           "Cmaj9",
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
				cType:     []string{internal.Major, internal.Ninth},
			},
		},
		{
			name: "C9",
			args: args{
				description: "C9",
			},
			want: Chord{
				name:           "C9",
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
				cType:     []string{internal.Dominant, internal.Ninth},
			},
		},
		{
			name: "Cm9",
			args: args{
				description: "Cm9",
			},
			want: Chord{
				name:           "Cm9",
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
				cType:     []string{internal.Minor, internal.Ninth},
			},
		},
		{
			name: "C11",
			args: args{
				description: "C11",
			},
			want: Chord{
				name:           "C11",
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
				cType:     []string{internal.Dominant, internal.Eleventh},
			},
		},
		{
			name: "Cm11",
			args: args{
				description: "Cm11",
			},
			want: Chord{
				name:           "Cm11",
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
				cType:     []string{internal.Minor, internal.Eleventh},
			},
		},
		{
			name: "Cmaj13",
			args: args{
				description: "Cmaj13",
			},
			want: Chord{
				name:           "Cmaj13",
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
				cType:     []string{internal.Major, internal.Thirteenth},
			},
		},
		{
			name: "C13",
			args: args{
				description: "C13",
			},
			want: Chord{
				name:           "C13",
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
				cType:     []string{internal.Dominant, internal.Thirteenth},
			},
		},
		{
			name: "Cm13",
			args: args{
				description: "Cm13",
			},
			want: Chord{
				name:           "Cm13",
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
				cType:     []string{internal.Minor, internal.Thirteenth},
			},
		},
		{
			name: "C7b5",
			args: args{
				description: "C7b5",
			},
			want: Chord{
				name:           "C7b5",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.ID5, internal.Im7},
				cType:     []string{internal.Dominant, internal.Seven, internal.Flat, internal.Five},
			},
		},
		{
			name: "C7b9",
			args: args{
				description: "C7b9",
			},
			want: Chord{
				name:           "C7b9",
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
				cType:     []string{internal.Dominant, internal.Seven, internal.Flat, internal.Nine},
			},
		},
		{
			name: "C7#9",
			args: args{
				description: "C7#9",
			},
			want: Chord{
				name:           "C7#9",
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
				cType:     []string{internal.Dominant, internal.Seven, internal.Sharp, internal.Nine},
			},
		},
		{
			name: "C7b5b9",
			args: args{
				description: "C7b5b9",
			},
			want: Chord{
				name:           "C7b5b9",
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
				cType:     []string{internal.Dominant, internal.Seven, internal.Flat, internal.Five, internal.Flat, internal.Nine},
			},
		},
		{
			name: "C7b5(b9)",
			args: args{
				description: "C7b5(b9)",
			},
			want: Chord{
				name:           "C7b5(b9)",
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
				cType:     []string{internal.Dominant, internal.Seven, internal.Flat, internal.Five, internal.Flat, internal.Nine},
			},
		},
		{
			name: "C7#5#9",
			args: args{
				description: "C7#5#9",
			},
			want: Chord{
				name:           "C7#5#9",
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
				cType:     []string{internal.Dominant, internal.Seven, internal.Sharp, internal.Five, internal.Sharp, internal.Nine},
			},
		},
		{
			name: "C7#5(#9)",
			args: args{
				description: "C7#5(#9)",
			},
			want: Chord{
				name:           "C7#5(#9)",
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
				cType:     []string{internal.Dominant, internal.Seven, internal.Sharp, internal.Five, internal.Sharp, internal.Nine},
			},
		},
		{
			name: "C7b5#9",
			args: args{
				description: "C7b5#9",
			},
			want: Chord{
				name:           "C7b5#9",
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
				cType:     []string{internal.Dominant, internal.Seven, internal.Flat, internal.Five, internal.Sharp, internal.Nine},
			},
		},
		{
			name: "C7b5(#9)",
			args: args{
				description: "C7b5(#9)",
			},
			want: Chord{
				name:           "C7b5(#9)",
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
				cType:     []string{internal.Dominant, internal.Seven, internal.Flat, internal.Five, internal.Sharp, internal.Nine},
			},
		},
		{
			name: "C7#5(b9)",
			args: args{
				description: "C7#5(b9)",
			},
			want: Chord{
				name:           "C7#5(b9)",
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
				cType:     []string{internal.Dominant, internal.Seven, internal.Sharp, internal.Five, internal.Flat, internal.Nine},
			},
		},
		{
			name: "C7#11",
			args: args{
				description: "C7#11",
			},
			want: Chord{
				name:           "C7#11",
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
				cType:     []string{internal.Dominant, internal.Seven, internal.Sharp, internal.Eleven},
			},
		},
		{
			name: "C9b5",
			args: args{
				description: "C9b5",
			},
			want: Chord{
				name:           "C9b5",
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
				cType:     []string{internal.Dominant, internal.Nine, internal.Flat, internal.Five},
			},
		},
		{
			name: "C9#5",
			args: args{
				description: "C9#5",
			},
			want: Chord{
				name:           "C9#5",
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
				cType:     []string{internal.Dominant, internal.Nine, internal.Sharp, internal.Five},
			},
		},
		{
			name: "C9#11",
			args: args{
				description: "C9#11",
			},
			want: Chord{
				name:           "C9#11",
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
				cType:     []string{internal.Dominant, internal.Nine, internal.Sharp, internal.Eleven},
			},
		},
		{
			name: "C13b9",
			args: args{
				description: "C13b9",
			},
			want: Chord{
				name:           "C13b9",
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
				cType:     []string{internal.Dominant, internal.Thirteen, internal.Flat, internal.Nine},
			},
		},
		{
			name: "C13#9",
			args: args{
				description: "C13#9",
			},
			want: Chord{
				name:           "C13#9",
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
				cType:     []string{internal.Dominant, internal.Thirteen, internal.Sharp, internal.Nine},
			},
		},
		{
			name: "C13#11",
			args: args{
				description: "C13#11",
			},
			want: Chord{
				name:           "C13#11",
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
				cType:     []string{internal.Dominant, internal.Thirteen, internal.Sharp, internal.Eleven},
			},
		},
		{
			name: "Cm7b5",
			args: args{
				description: "Cm7b5",
			},
			want: Chord{
				name:           "Cm7b5",
				chordBasicType: internal.Minor,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.ID5, internal.Im7},
				cType:     []string{internal.Minor, internal.Seven, internal.Flat, internal.Five},
			},
		},
		{
			name: "Cm7#5",
			args: args{
				description: "Cm7#5",
			},
			want: Chord{
				name:           "Cm7#5",
				chordBasicType: internal.Minor,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GSharp),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.IA5, internal.Im7},
				cType:     []string{internal.Minor, internal.Seven, internal.Sharp, internal.Five},
			},
		},
		{
			name: "Cmaj7b5",
			args: args{
				description: "Cmaj7b5",
			},
			want: Chord{
				name:           "Cmaj7b5",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GFlat),
					NewNote(internal.B),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.ID5, internal.IM7},
				cType:     []string{internal.Major, internal.Seven, internal.Flat, internal.Five},
			},
		},
		{
			name: "Cmaj7#5",
			args: args{
				description: "Cmaj7#5",
			},
			want: Chord{
				name:           "Cmaj7#5",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.B),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.IM7},
				cType:     []string{internal.Major, internal.Seven, internal.Sharp, internal.Five},
			},
		},
		{
			name: "Cmaj7#11",
			args: args{
				description: "Cmaj7#11",
			},
			want: Chord{
				name:           "Cmaj7#11",
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
				cType:     []string{internal.Major, internal.Seven, internal.Sharp, internal.Eleven},
			},
		},
		{
			name: "Cmaj9#11",
			args: args{
				description: "Cmaj9#11",
			},
			want: Chord{
				name:           "Cmaj9#11",
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
				cType:     []string{internal.Major, internal.Nine, internal.Sharp, internal.Eleven},
			},
		},
		{
			name: "C6",
			args: args{
				description: "C6",
			},
			want: Chord{
				name:           "C6",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.A),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM6},
				cType:     []string{internal.Sixth},
			},
		},
		{
			name: "C6/9",
			args: args{
				description: "C6/9",
			},
			want: Chord{
				name:           "C6/9",
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
				cType:     []string{internal.Sixth, internal.Added, internal.Ninth},
			},
		},
		{
			name: "C6add9",
			args: args{
				description: "C6add9",
			},
			want: Chord{
				name:           "C6add9",
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
				cType:     []string{internal.Sixth, internal.Added, internal.Ninth},
			},
		},
		{
			name: "C6(add9)",
			args: args{
				description: "C6(add9)",
			},
			want: Chord{
				name:           "C6(add9)",
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
				cType:     []string{internal.Sixth, internal.Added, internal.Ninth},
			},
		},
		{
			name: "C/G",
			args: args{
				description: "C/G",
			},
			want: Chord{
				name:           "C/G",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				bassNote:       NewNote("G"),
				notes: []Note{
					NewNote(internal.G),
					NewNote(internal.C),
					NewNote(internal.E),
				},
				structure: []int{internal.IP5, internal.IUnison, internal.IM3},
				cType:     []string{internal.Major, internal.Over, "G"},
			},
		},
		{
			name: "C/G#",
			args: args{
				description: "C/G#",
			},
			want: Chord{
				name:           "C/G#",
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
				cType:     []string{internal.Major, internal.Over, "G#"},
			},
		},
		{
			name: "Csus2",
			args: args{
				description: "Csus2",
			},
			want: Chord{
				name:           "Csus2",
				chordBasicType: internal.Suspended,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.D),
					NewNote(internal.G),
				},
				structure: []int{internal.IUnison, internal.IM2, internal.IP5},
				cType:     []string{internal.Suspended, internal.Second},
			},
		},
		{
			name: "Csus4",
			args: args{
				description: "Csus4",
			},
			want: Chord{
				name:           "Csus4",
				chordBasicType: internal.Suspended,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.F),
					NewNote(internal.G),
				},
				structure: []int{internal.IUnison, internal.IP4, internal.IP5},
				cType:     []string{internal.Suspended, internal.Fourth},
			},
		},
		{
			name: "Csus",
			args: args{
				description: "Csus",
			},
			want: Chord{
				name:           "Csus",
				chordBasicType: internal.Suspended,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.F),
					NewNote(internal.G),
				},
				structure: []int{internal.IUnison, internal.IP4, internal.IP5},
				cType:     []string{internal.Suspended, internal.Fourth},
			},
		},
		{
			name: "C7sus4",
			args: args{
				description: "C7sus4",
			},
			want: Chord{
				name:           "C7sus4",
				chordBasicType: internal.Suspended,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.F),
					NewNote(internal.G),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.IP4, internal.IP5, internal.Im7},
				cType:     []string{internal.Dominant, internal.Seventh, internal.Suspended, internal.Fourth},
			},
		},
		{
			name: "C7sus",
			args: args{
				description: "C7sus",
			},
			want: Chord{
				name:           "C7sus",
				chordBasicType: internal.Suspended,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.F),
					NewNote(internal.G),
					NewNote(internal.BFlat),
				},
				structure: []int{internal.IUnison, internal.IP4, internal.IP5, internal.Im7},
				cType:     []string{internal.Dominant, internal.Seventh, internal.Suspended, internal.Fourth},
			},
		},
		{
			name: "Cadd9",
			args: args{
				description: "Cadd9",
			},
			want: Chord{
				name:           "Cadd9",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.D),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM9},
				cType:     []string{internal.Added, internal.Ninth},
			},
		},
		{
			name: "C(add9)",
			args: args{
				description: "C(add9)",
			},
			want: Chord{
				name:           "C(add9)",
				chordBasicType: internal.Major,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.D),
				},
				structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM9},
				cType:     []string{internal.Added, internal.Ninth},
			},
		},
		{
			name: "Cm(add9)",
			args: args{
				description: "Cm(add9)",
			},
			want: Chord{
				name:           "Cm(add9)",
				chordBasicType: internal.Minor,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.G),
					NewNote(internal.D),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.IP5, internal.IM9},
				cType:     []string{internal.Minor, internal.Added, internal.Ninth},
			},
		},
		{
			name: "Cdim6",
			args: args{
				description: "Cdim6",
			},
			want: Chord{
				name:           "Cdim6",
				chordBasicType: internal.Diminished,
				root:           NewNote("C"),
				notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GFlat),
					NewNote(internal.A),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.ID5, internal.IM6},
				cType:     []string{internal.Diminished, internal.Sixth},
			},
		},
		{
			name: "Cdim9",
			args: args{
				description: "Cdim9",
			},
			want: Chord{
				name:           "Cdim9",
				chordBasicType: internal.Diminished,
				root:           NewNote("C"),
				notes: []Note{
					//C, E♭, G♭, B♭♭, D
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GFlat),
					NewNote(internal.A),
					NewNote(internal.D),
				},
				structure: []int{internal.IUnison, internal.Im3, internal.ID5, internal.IM6, internal.IM9},
				cType:     []string{internal.Diminished, internal.Ninth},
			},
		},
		{
			name: "Cdim11",
			args: args{
				description: "Cdim11",
			},
			want: Chord{
				name:           "Cdim11",
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
				cType:     []string{internal.Diminished, internal.Eleventh},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := parseChord(tt.args.description)
			if err != nil {
				t.Errorf("parseChord() error = %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseChord() = %v, want %v", got, tt.want)
			}
		})
	}
}
