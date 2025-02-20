package scales

import (
	"testing"

	"github.com/andriikushch/scales/pkg/internal"

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
				previousNote: NewNote(internal.C),
				steps:        0,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(internal.BSharp), NewNote(internal.C)},
		},
		{
			name: "no steps, from C#",
			args: args{
				previousNote: NewNote(internal.CSharp),
				steps:        0,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(internal.CSharp), NewNote(internal.DFlat)},
		},
		{
			name: "no steps, from B",
			args: args{
				previousNote: NewNote(internal.B),
				steps:        0,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(internal.B)},
		},
		{
			name: "half step, from C",
			args: args{
				previousNote: NewNote(internal.C),
				steps:        internal.HalfStep,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(internal.CSharp), NewNote(internal.DFlat)},
		},
		{
			name: "half step, from C#",
			args: args{
				previousNote: NewNote(internal.CSharp),
				steps:        internal.HalfStep,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(internal.CDoubleSharp), NewNote(internal.D)},
		},
		{
			name: "half step, from Db",
			args: args{
				previousNote: NewNote(internal.DFlat),
				steps:        internal.HalfStep,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(internal.CDoubleSharp), NewNote(internal.D)},
		},
		{
			name: "half step, from B",
			args: args{
				previousNote: NewNote(internal.B),
				steps:        internal.HalfStep,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(internal.BSharp), NewNote(internal.C)},
		},
		{
			name: "step, from C",
			args: args{
				previousNote: NewNote(internal.C),
				steps:        internal.Step,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(internal.CDoubleSharp), NewNote(internal.D)},
		},
		{
			name: "step, from C#",
			args: args{
				previousNote: NewNote(internal.CSharp),
				steps:        internal.Step,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(internal.DSharp), NewNote(internal.EFlat)},
		},
		{
			name: "step, from Db",
			args: args{
				previousNote: NewNote(internal.DFlat),
				steps:        internal.Step,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(internal.DSharp), NewNote(internal.EFlat)},
		},
		{
			name: "step, from B",
			args: args{
				previousNote: NewNote(internal.B),
				steps:        internal.Step,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(internal.CSharp), NewNote(internal.DFlat)},
		},
		{
			name: "-1 step, from C",
			args: args{
				previousNote: NewNote(internal.C),
				steps:        -internal.Step,
			},
			s:    defaultChromaticScale,
			want: []Note{NewNote(internal.ASharp), NewNote(internal.BFlat)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, tt.s.next(tt.args.previousNote, tt.args.steps))
		})
	}
}

func TestScale_normalize(t *testing.T) {
	t.Parallel()

	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   Note
	}{
		{
			name: "F##",
			fields: fields{
				Name: "F##",
			},
			want: Note{
				Name: "F##",
			},
		},
		{
			name: "F###",
			fields: fields{
				Name: "F###",
			},
			want: Note{
				Name: "G#",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			n := Note{
				Name: tt.fields.Name,
			}

			got := defaultChromaticScale.normalize(n)
			require.Equal(t, tt.want, got)
		})
	}
}
