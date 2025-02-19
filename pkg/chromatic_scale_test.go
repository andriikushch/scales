package scales

import (
	"testing"

	"github.com/andriikushch/scales/pkg/internal/chromatic"
	"github.com/stretchr/testify/require"
)

func TestScale_Next(t *testing.T) {
	t.Parallel()

	type args struct {
		previousNote Note
		steps        int
	}

	tests := []struct {
		name string
		s    chromaticScale
		args args
		want []Note
	}{
		{
			name: "no steps, from C",
			args: args{
				previousNote: NewNote(chromatic.C),
				steps:        0,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(chromatic.BSharp), NewNote(chromatic.C)},
		},
		{
			name: "no steps, from C#",
			args: args{
				previousNote: NewNote(chromatic.CSharp),
				steps:        0,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(chromatic.CSharp), NewNote(chromatic.DFlat)},
		},
		{
			name: "no steps, from B",
			args: args{
				previousNote: NewNote(chromatic.B),
				steps:        0,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(chromatic.B)},
		},
		{
			name: "half step, from C",
			args: args{
				previousNote: NewNote(chromatic.C),
				steps:        chromatic.HalfStep,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(chromatic.CSharp), NewNote(chromatic.DFlat)},
		},
		{
			name: "half step, from C#",
			args: args{
				previousNote: NewNote(chromatic.CSharp),
				steps:        chromatic.HalfStep,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(chromatic.CDoubleSharp), NewNote(chromatic.D)},
		},
		{
			name: "half step, from Db",
			args: args{
				previousNote: NewNote(chromatic.DFlat),
				steps:        chromatic.HalfStep,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(chromatic.CDoubleSharp), NewNote(chromatic.D)},
		},
		{
			name: "half step, from B",
			args: args{
				previousNote: NewNote(chromatic.B),
				steps:        chromatic.HalfStep,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(chromatic.BSharp), NewNote(chromatic.C)},
		},
		{
			name: "step, from C",
			args: args{
				previousNote: NewNote(chromatic.C),
				steps:        chromatic.Step,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(chromatic.CDoubleSharp), NewNote(chromatic.D)},
		},
		{
			name: "step, from C#",
			args: args{
				previousNote: NewNote(chromatic.CSharp),
				steps:        chromatic.Step,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(chromatic.DSharp), NewNote(chromatic.EFlat)},
		},
		{
			name: "step, from Db",
			args: args{
				previousNote: NewNote(chromatic.DFlat),
				steps:        chromatic.Step,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(chromatic.DSharp), NewNote(chromatic.EFlat)},
		},
		{
			name: "step, from B",
			args: args{
				previousNote: NewNote(chromatic.B),
				steps:        chromatic.Step,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(chromatic.CSharp), NewNote(chromatic.DFlat)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, tt.s.next(tt.args.previousNote, tt.args.steps))
		})
	}
}
