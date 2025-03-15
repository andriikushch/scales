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
		want chord
	}{
		{
			name: "Am",
			args: args{
				description: "Am",
			},
			want: chord{
				Name:           "Am",
				ChordBasicType: internal.Minor,
				Root:           NewNote("A"),
				Notes: []Note{
					NewNote(internal.A),
					NewNote(internal.C),
					NewNote(internal.E),
				},
				Structure: []int{internal.IUnison, internal.Im3, internal.IP5},
				Type:      []string{internal.Minor},
			},
		},
		{
			name: "Am/C",
			args: args{
				description: "Am/C",
			},
			want: chord{
				Name:           "Am/C",
				ChordBasicType: internal.Minor,
				BassNote:       NewNote("C"),
				Root:           NewNote("A"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.A),
					NewNote(internal.E),
				},
				Structure: []int{internal.Im3, internal.IUnison, internal.IP5},
				Type:      []string{internal.Minor, internal.Over, "C"},
			},
		},
		{
			name: "Am7",
			args: args{
				description: "Am7",
			},
			want: chord{
				Name:           "Am7",
				ChordBasicType: internal.Minor,
				Root:           NewNote("A"),
				Notes: []Note{
					NewNote(internal.A),
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
				},
				Structure: []int{internal.IUnison, internal.Im3, internal.IP5, internal.Im7},
				Type:      []string{internal.Minor, internal.Seventh},
			},
		},
		{
			name: "C",
			args: args{
				description: "C",
			},
			want: chord{
				Name:           "C",
				ChordBasicType: internal.Major,
				Root:           NewNote(internal.C),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5},
				Type:      []string{internal.Major},
			},
		},
		{
			name: "Cmaj",
			args: args{
				description: "Cmaj",
			},
			want: chord{
				Name:           "Cmaj",
				ChordBasicType: internal.Major,
				Root:           NewNote(internal.C),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5},
				Type:      []string{internal.Major},
			},
		},
		{
			name: "Cmaj7",
			args: args{
				description: "Cmaj7",
			},
			want: chord{
				Name:           "Cmaj7",
				ChordBasicType: internal.Major,
				Root:           NewNote(internal.C),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.B),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM7},
				Type:      []string{internal.Major, internal.Seventh},
			},
		},
		{
			name: "C7",
			args: args{
				description: "C7",
			},
			want: chord{
				Name:           "C7",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7},
				Type:      []string{internal.Dominant, internal.Seventh},
			},
		},
		{
			name: "Cm7",
			args: args{
				description: "Cm7",
			},
			want: chord{
				Name:           "Cm7",
				ChordBasicType: internal.Minor,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.G),
					NewNote(internal.BFlat),
				},
				Structure: []int{internal.IUnison, internal.Im3, internal.IP5, internal.Im7},
				Type:      []string{internal.Minor, internal.Seventh},
			},
		},
		{
			name: "Cm7b5",
			args: args{
				description: "Cm7b5",
			},
			want: chord{
				Name:           "Cm7b5",
				ChordBasicType: internal.Minor,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
				},
				Structure: []int{internal.IUnison, internal.Im3, internal.ID5, internal.Im7},
				Type:      []string{internal.Minor, internal.Seven, internal.Flat, internal.Five},
			},
		},
		{
			name: "Cdim7",
			args: args{
				description: "Cdim7",
			},
			want: chord{
				Name:           "Cdim7",
				ChordBasicType: internal.Diminished,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GFlat),
					NewNote(internal.A),
				},
				Structure: []int{internal.IUnison, internal.Im3, internal.ID5, internal.IM6},
				Type:      []string{internal.Diminished, internal.Seventh},
			},
		},
		{
			name: "Caug7",
			args: args{
				description: "Caug7",
			},
			want: chord{
				Name:           "Caug7",
				ChordBasicType: internal.Augmented,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.BFlat),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.Im7},
				Type:      []string{internal.Augmented, internal.Seventh},
			},
		},
		{
			name: "Caug9",
			args: args{
				description: "Caug9",
			},
			want: chord{
				Name:           "Caug9",
				ChordBasicType: internal.Augmented,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.BFlat),
					NewNote(internal.D),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.Im7, internal.IM9},
				Type:      []string{internal.Augmented, internal.Ninth},
			},
		},
		{
			name: "Caug11",
			args: args{
				description: "Caug11",
			},
			want: chord{
				Name:           "Caug11",
				ChordBasicType: internal.Augmented,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.BFlat),
					NewNote(internal.D),
					NewNote(internal.F),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.Im7, internal.IM9, internal.IP11},
				Type:      []string{internal.Augmented, internal.Eleventh},
			},
		},
		{
			name: "Caug6",
			args: args{
				description: "Caug6",
			},
			want: chord{
				Name:           "Caug6",
				ChordBasicType: internal.Augmented,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.A),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.IM6},
				Type:      []string{internal.Augmented, internal.Sixth},
			},
		},
		{
			name: "C7b5",
			args: args{
				description: "C7b5",
			},
			want: chord{
				Name:           "C7b5",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.ID5, internal.Im7},
				Type:      []string{internal.Dominant, internal.Seven, internal.Flat, internal.Five},
			},
		},
		{
			name: "Cm(maj7)",
			args: args{
				description: "Cm(maj7)",
			},
			want: chord{
				Name:           "Cm(maj7)",
				ChordBasicType: internal.Minor,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.G),
					NewNote(internal.B),
				},
				Structure: []int{internal.IUnison, internal.Im3, internal.IP5, internal.IM7},
				Type:      []string{internal.Minor, internal.Major, internal.Seventh},
			},
		},
		{
			name: "Cmaj9",
			args: args{
				description: "Cmaj9",
			},
			want: chord{
				Name:           "Cmaj9",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.B),
					NewNote(internal.D),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM7, internal.IM9},
				Type:      []string{internal.Major, internal.Ninth},
			},
		},
		{
			name: "C9",
			args: args{
				description: "C9",
			},
			want: chord{
				Name:           "C9",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.D),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.IM9},
				Type:      []string{internal.Dominant, internal.Ninth},
			},
		},
		{
			name: "Cm9",
			args: args{
				description: "Cm9",
			},
			want: chord{
				Name:           "Cm9",
				ChordBasicType: internal.Minor,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.D),
				},
				Structure: []int{internal.IUnison, internal.Im3, internal.IP5, internal.Im7, internal.IM9},
				Type:      []string{internal.Minor, internal.Ninth},
			},
		},
		{
			name: "C11",
			args: args{
				description: "C11",
			},
			want: chord{
				Name:           "C11",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.D),
					NewNote(internal.F),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.IM9, internal.IP11},
				Type:      []string{internal.Dominant, internal.Eleventh},
			},
		},
		{
			name: "Cm11",
			args: args{
				description: "Cm11",
			},
			want: chord{
				Name:           "Cm11",
				ChordBasicType: internal.Minor,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.D),
					NewNote(internal.F),
				},
				Structure: []int{internal.IUnison, internal.Im3, internal.IP5, internal.Im7, internal.IM9, internal.IP11},
				Type:      []string{internal.Minor, internal.Eleventh},
			},
		},
		{
			name: "Cmaj13",
			args: args{
				description: "Cmaj13",
			},
			want: chord{
				Name:           "Cmaj13",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.B),
					NewNote(internal.D),
					NewNote(internal.F),
					NewNote(internal.A),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM7, internal.IM9, internal.IP11, internal.IM13},
				Type:      []string{internal.Major, internal.Thirteenth},
			},
		},
		{
			name: "C13",
			args: args{
				description: "C13",
			},
			want: chord{
				Name:           "C13",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.D),
					NewNote(internal.F),
					NewNote(internal.A),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.IM9, internal.IP11, internal.IM13},
				Type:      []string{internal.Dominant, internal.Thirteenth},
			},
		},
		{
			name: "Cm13",
			args: args{
				description: "Cm13",
			},
			want: chord{
				Name:           "Cm13",
				ChordBasicType: internal.Minor,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.D),
					NewNote(internal.F),
					NewNote(internal.A),
				},
				Structure: []int{internal.IUnison, internal.Im3, internal.IP5, internal.Im7, internal.IM9, internal.IP11, internal.IM13},
				Type:      []string{internal.Minor, internal.Thirteenth},
			},
		},
		{
			name: "C7b5",
			args: args{
				description: "C7b5",
			},
			want: chord{
				Name:           "C7b5",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.ID5, internal.Im7},
				Type:      []string{internal.Dominant, internal.Seven, internal.Flat, internal.Five},
			},
		},
		{
			name: "C7b9",
			args: args{
				description: "C7b9",
			},
			want: chord{
				Name:           "C7b9",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.DFlat),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.Im9},
				Type:      []string{internal.Dominant, internal.Seven, internal.Flat, internal.Nine},
			},
		},
		{
			name: "C7#9",
			args: args{
				description: "C7#9",
			},
			want: chord{
				Name:           "C7#9",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.DSharp),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.Im10},
				Type:      []string{internal.Dominant, internal.Seven, internal.Sharp, internal.Nine},
			},
		},
		{
			name: "C7b5b9",
			args: args{
				description: "C7b5b9",
			},
			want: chord{
				Name:           "C7b5b9",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
					NewNote(internal.DFlat),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.ID5, internal.Im7, internal.Im9},
				Type:      []string{internal.Dominant, internal.Seven, internal.Flat, internal.Five, internal.Flat, internal.Nine},
			},
		},
		{
			name: "C7b5(b9)",
			args: args{
				description: "C7b5(b9)",
			},
			want: chord{
				Name:           "C7b5(b9)",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
					NewNote(internal.DFlat),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.ID5, internal.Im7, internal.Im9},
				Type:      []string{internal.Dominant, internal.Seven, internal.Flat, internal.Five, internal.Flat, internal.Nine},
			},
		},
		{
			name: "C7#5#9",
			args: args{
				description: "C7#5#9",
			},
			want: chord{
				Name:           "C7#5#9",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.BFlat),
					NewNote(internal.DSharp),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.Im7, internal.Im10},
				Type:      []string{internal.Dominant, internal.Seven, internal.Sharp, internal.Five, internal.Sharp, internal.Nine},
			},
		},
		{
			name: "C7#5(#9)",
			args: args{
				description: "C7#5(#9)",
			},
			want: chord{
				Name:           "C7#5(#9)",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.BFlat),
					NewNote(internal.DSharp),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.Im7, internal.Im10},
				Type:      []string{internal.Dominant, internal.Seven, internal.Sharp, internal.Five, internal.Sharp, internal.Nine},
			},
		},
		{
			name: "C7b5#9",
			args: args{
				description: "C7b5#9",
			},
			want: chord{
				Name:           "C7b5#9",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
					NewNote(internal.DSharp),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.ID5, internal.Im7, internal.Im10},
				Type:      []string{internal.Dominant, internal.Seven, internal.Flat, internal.Five, internal.Sharp, internal.Nine},
			},
		},
		{
			name: "C7b5(#9)",
			args: args{
				description: "C7b5(#9)",
			},
			want: chord{
				Name:           "C7b5(#9)",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
					NewNote(internal.DSharp),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.ID5, internal.Im7, internal.Im10},
				Type:      []string{internal.Dominant, internal.Seven, internal.Flat, internal.Five, internal.Sharp, internal.Nine},
			},
		},
		{
			name: "C7#5(b9)",
			args: args{
				description: "C7#5(b9)",
			},
			want: chord{
				Name:           "C7#5(b9)",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.BFlat),
					NewNote(internal.DFlat),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.Im7, internal.Im9},
				Type:      []string{internal.Dominant, internal.Seven, internal.Sharp, internal.Five, internal.Flat, internal.Nine},
			},
		},
		{
			name: "C7#11",
			args: args{
				description: "C7#11",
			},
			want: chord{
				Name:           "C7#11",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.FSharp),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.IA11},
				Type:      []string{internal.Dominant, internal.Seven, internal.Sharp, internal.Eleven},
			},
		},
		{
			name: "C9b5",
			args: args{
				description: "C9b5",
			},
			want: chord{
				Name:           "C9b5",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
					NewNote(internal.D),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.ID5, internal.Im7, internal.IM9},
				Type:      []string{internal.Dominant, internal.Nine, internal.Flat, internal.Five},
			},
		},
		{
			name: "C9#5",
			args: args{
				description: "C9#5",
			},
			want: chord{
				Name:           "C9#5",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.BFlat),
					NewNote(internal.D),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.Im7, internal.IM9},
				Type:      []string{internal.Dominant, internal.Nine, internal.Sharp, internal.Five},
			},
		},
		{
			name: "C9#11",
			args: args{
				description: "C9#11",
			},
			want: chord{
				Name:           "C9#11",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.D),
					NewNote(internal.FSharp),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.IM9, internal.IA11},
				Type:      []string{internal.Dominant, internal.Nine, internal.Sharp, internal.Eleven},
			},
		},
		{
			name: "C13b9",
			args: args{
				description: "C13b9",
			},
			want: chord{
				Name:           "C13b9",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.DFlat),
					NewNote(internal.F),
					NewNote(internal.A),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.Im9, internal.IP11, internal.IM13},
				Type:      []string{internal.Dominant, internal.Thirteen, internal.Flat, internal.Nine},
			},
		},
		{
			name: "C13#9",
			args: args{
				description: "C13#9",
			},
			want: chord{
				Name:           "C13#9",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.DSharp),
					NewNote(internal.F),
					NewNote(internal.A),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.IA9, internal.IP11, internal.IM13},
				Type:      []string{internal.Dominant, internal.Thirteen, internal.Sharp, internal.Nine},
			},
		},
		{
			name: "C13#11",
			args: args{
				description: "C13#11",
			},
			want: chord{
				Name:           "C13#11",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.BFlat),
					NewNote(internal.D),
					NewNote(internal.FSharp),
					NewNote(internal.A),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.Im7, internal.IM9, internal.IA11, internal.IM13},
				Type:      []string{internal.Dominant, internal.Thirteen, internal.Sharp, internal.Eleven},
			},
		},
		{
			name: "Cm7b5",
			args: args{
				description: "Cm7b5",
			},
			want: chord{
				Name:           "Cm7b5",
				ChordBasicType: internal.Minor,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GFlat),
					NewNote(internal.BFlat),
				},
				Structure: []int{internal.IUnison, internal.Im3, internal.ID5, internal.Im7},
				Type:      []string{internal.Minor, internal.Seven, internal.Flat, internal.Five},
			},
		},
		{
			name: "Cm7#5",
			args: args{
				description: "Cm7#5",
			},
			want: chord{
				Name:           "Cm7#5",
				ChordBasicType: internal.Minor,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GSharp),
					NewNote(internal.BFlat),
				},
				Structure: []int{internal.IUnison, internal.Im3, internal.IA5, internal.Im7},
				Type:      []string{internal.Minor, internal.Seven, internal.Sharp, internal.Five},
			},
		},
		{
			name: "Cmaj7b5",
			args: args{
				description: "Cmaj7b5",
			},
			want: chord{
				Name:           "Cmaj7b5",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GFlat),
					NewNote(internal.B),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.ID5, internal.IM7},
				Type:      []string{internal.Major, internal.Seven, internal.Flat, internal.Five},
			},
		},
		{
			name: "Cmaj7#5",
			args: args{
				description: "Cmaj7#5",
			},
			want: chord{
				Name:           "Cmaj7#5",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.GSharp),
					NewNote(internal.B),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IA5, internal.IM7},
				Type:      []string{internal.Major, internal.Seven, internal.Sharp, internal.Five},
			},
		},
		{
			name: "Cmaj7#11",
			args: args{
				description: "Cmaj7#11",
			},
			want: chord{
				Name:           "Cmaj7#11",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.B),
					NewNote(internal.FSharp),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM7, internal.IA11},
				Type:      []string{internal.Major, internal.Seven, internal.Sharp, internal.Eleven},
			},
		},
		{
			name: "Cmaj9#11",
			args: args{
				description: "Cmaj9#11",
			},
			want: chord{
				Name:           "Cmaj9#11",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.B),
					NewNote(internal.D),
					NewNote(internal.FSharp),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM7, internal.IM9, internal.IA11},
				Type:      []string{internal.Major, internal.Nine, internal.Sharp, internal.Eleven},
			},
		},
		{
			name: "C6",
			args: args{
				description: "C6",
			},
			want: chord{
				Name:           "C6",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.A),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM6},
				Type:      []string{internal.Sixth},
			},
		},
		{
			name: "C6/9",
			args: args{
				description: "C6/9",
			},
			want: chord{
				Name:           "C6/9",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.A),
					NewNote(internal.D),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM6, internal.IM9},
				Type:      []string{internal.Sixth, internal.Added, internal.Ninth},
			},
		},
		{
			name: "C6add9",
			args: args{
				description: "C6add9",
			},
			want: chord{
				Name:           "C6add9",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.A),
					NewNote(internal.D),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM6, internal.IM9},
				Type:      []string{internal.Sixth, internal.Added, internal.Ninth},
			},
		},
		{
			name: "C6(add9)",
			args: args{
				description: "C6(add9)",
			},
			want: chord{
				Name:           "C6(add9)",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.A),
					NewNote(internal.D),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM6, internal.IM9},
				Type:      []string{internal.Sixth, internal.Added, internal.Ninth},
			},
		},
		{
			name: "C/G",
			args: args{
				description: "C/G",
			},
			want: chord{
				Name:           "C/G",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				BassNote:       NewNote("G"),
				Notes: []Note{
					NewNote(internal.G),
					NewNote(internal.C),
					NewNote(internal.E),
				},
				Structure: []int{internal.IP5, internal.IUnison, internal.IM3},
				Type:      []string{internal.Major, internal.Over, "G"},
			},
		},
		{
			name: "C/G#",
			args: args{
				description: "C/G#",
			},
			want: chord{
				Name:           "C/G#",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				BassNote:       NewNote("G#"),
				Notes: []Note{
					NewNote(internal.GSharp),
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
				},
				Structure: []int{internal.IA5, internal.IUnison, internal.IM3, internal.IP5},
				Type:      []string{internal.Major, internal.Over, "G#"},
			},
		},
		{
			name: "Csus2",
			args: args{
				description: "Csus2",
			},
			want: chord{
				Name:           "Csus2",
				ChordBasicType: internal.Suspended,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.D),
					NewNote(internal.G),
				},
				Structure: []int{internal.IUnison, internal.IM2, internal.IP5},
				Type:      []string{internal.Suspended, internal.Second},
			},
		},
		{
			name: "Csus4",
			args: args{
				description: "Csus4",
			},
			want: chord{
				Name:           "Csus4",
				ChordBasicType: internal.Suspended,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.F),
					NewNote(internal.G),
				},
				Structure: []int{internal.IUnison, internal.IP4, internal.IP5},
				Type:      []string{internal.Suspended, internal.Fourth},
			},
		},
		{
			name: "Csus",
			args: args{
				description: "Csus",
			},
			want: chord{
				Name:           "Csus",
				ChordBasicType: internal.Suspended,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.F),
					NewNote(internal.G),
				},
				Structure: []int{internal.IUnison, internal.IP4, internal.IP5},
				Type:      []string{internal.Suspended, internal.Fourth},
			},
		},
		{
			name: "C7sus4",
			args: args{
				description: "C7sus4",
			},
			want: chord{
				Name:           "C7sus4",
				ChordBasicType: internal.Suspended,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.F),
					NewNote(internal.G),
					NewNote(internal.BFlat),
				},
				Structure: []int{internal.IUnison, internal.IP4, internal.IP5, internal.Im7},
				Type:      []string{internal.Dominant, internal.Seventh, internal.Suspended, internal.Fourth},
			},
		},
		{
			name: "C7sus",
			args: args{
				description: "C7sus",
			},
			want: chord{
				Name:           "C7sus",
				ChordBasicType: internal.Suspended,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.F),
					NewNote(internal.G),
					NewNote(internal.BFlat),
				},
				Structure: []int{internal.IUnison, internal.IP4, internal.IP5, internal.Im7},
				Type:      []string{internal.Dominant, internal.Seventh, internal.Suspended, internal.Fourth},
			},
		},
		{
			name: "Cadd9",
			args: args{
				description: "Cadd9",
			},
			want: chord{
				Name:           "Cadd9",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.D),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM9},
				Type:      []string{internal.Added, internal.Ninth},
			},
		},
		{
			name: "C(add9)",
			args: args{
				description: "C(add9)",
			},
			want: chord{
				Name:           "C(add9)",
				ChordBasicType: internal.Major,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.E),
					NewNote(internal.G),
					NewNote(internal.D),
				},
				Structure: []int{internal.IUnison, internal.IM3, internal.IP5, internal.IM9},
				Type:      []string{internal.Added, internal.Ninth},
			},
		},
		{
			name: "Cm(add9)",
			args: args{
				description: "Cm(add9)",
			},
			want: chord{
				Name:           "Cm(add9)",
				ChordBasicType: internal.Minor,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.G),
					NewNote(internal.D),
				},
				Structure: []int{internal.IUnison, internal.Im3, internal.IP5, internal.IM9},
				Type:      []string{internal.Minor, internal.Added, internal.Ninth},
			},
		},
		{
			name: "Cdim6",
			args: args{
				description: "Cdim6",
			},
			want: chord{
				Name:           "Cdim6",
				ChordBasicType: internal.Diminished,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GFlat),
					NewNote(internal.A),
				},
				Structure: []int{internal.IUnison, internal.Im3, internal.ID5, internal.IM6},
				Type:      []string{internal.Diminished, internal.Sixth},
			},
		},
		{
			name: "Cdim9",
			args: args{
				description: "Cdim9",
			},
			want: chord{
				Name:           "Cdim9",
				ChordBasicType: internal.Diminished,
				Root:           NewNote("C"),
				Notes: []Note{
					//C, E♭, G♭, B♭♭, D
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GFlat),
					NewNote(internal.A),
					NewNote(internal.D),
				},
				Structure: []int{internal.IUnison, internal.Im3, internal.ID5, internal.IM6, internal.IM9},
				Type:      []string{internal.Diminished, internal.Ninth},
			},
		},
		{
			name: "Cdim11",
			args: args{
				description: "Cdim11",
			},
			want: chord{
				Name:           "Cdim11",
				ChordBasicType: internal.Diminished,
				Root:           NewNote("C"),
				Notes: []Note{
					NewNote(internal.C),
					NewNote(internal.EFlat),
					NewNote(internal.GFlat),
					NewNote(internal.A),
					NewNote(internal.D),
					NewNote(internal.F),
				},
				Structure: []int{internal.IUnison, internal.Im3, internal.ID5, internal.IM6, internal.IM9, internal.IP11},
				Type:      []string{internal.Diminished, internal.Eleventh},
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
