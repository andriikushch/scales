package internal

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	t.Parallel()

	type args struct {
		chord string
	}

	tests := []struct {
		name string
		args args
		want []Token
	}{
		{
			name: "Am",
			args: args{
				chord: "Am",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "A",
				},
				{
					Type: MINOR,
				},
			},
		},
		{
			name: "C6/9",
			args: args{
				chord: "C6/9",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
				{
					Type:  NUMBER,
					Value: "6",
				},
				{
					Type:  ADD,
					Value: "9",
				},
			},
		},
		{
			name: "C6add9",
			args: args{
				chord: "C6add9",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
				{
					Type:  NUMBER,
					Value: "6",
				},
				{
					Type:  ADD,
					Value: "9",
				},
			},
		},
		{
			name: "C6(add9)",
			args: args{
				chord: "C6(add9)",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
				{
					Type:  NUMBER,
					Value: "6",
				},
				{
					Type:  ADD,
					Value: "9",
				},
			},
		},
		{
			name: "A#m",
			args: args{
				chord: "A#m",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "A#",
				},
				{
					Type: MINOR,
				},
			},
		},
		{
			name: "A#m#5",
			args: args{
				chord: "A#m#5",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "A#",
				},
				{
					Type: MINOR,
				},
				{
					Type:  SHARP,
					Value: "5",
				},
			},
		},
		{
			name: "Am7",
			args: args{
				chord: "Am7",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "A",
				},
				{
					Type:  MINOR,
					Value: "7",
				},
			},
		},
		{
			name: "C",
			args: args{
				chord: "C",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
			},
		},
		{
			name: "Csus2",
			args: args{
				chord: "Csus2",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
				{
					Type:  SUS,
					Value: "2",
				},
			},
		},
		{
			name: "C#",
			args: args{
				chord: "C#",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C#",
				},
			},
		},
		{
			name: "Cb",
			args: args{
				chord: "Cb",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "Cb",
				},
			},
		},
		{
			name: "Cmaj7",
			args: args{
				chord: "Cmaj7",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
				{
					Type:  MAJ,
					Value: "7",
				},
			},
		},
		{
			name: "C7",
			args: args{
				chord: "C7",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
				{
					Type:  NUMBER,
					Value: "7",
				},
			},
		},
		{
			name: "Cm7",
			args: args{
				chord: "Cm7",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
				{
					Type:  MINOR,
					Value: "7",
				},
			},
		},
		{
			name: "Cm7b5",
			args: args{
				chord: "Cm7b5",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
				{
					Type:  MINOR,
					Value: "7",
				},
				{
					Type:  FLAT,
					Value: "5",
				},
			},
		},
		{
			name: "Cdim7",
			args: args{
				chord: "Cdim7",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
				{
					Type:  DIM,
					Value: "7",
				},
			},
		},
		{
			name: "Cdim7add13",
			args: args{
				chord: "Cdim7add13",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
				{
					Type:  DIM,
					Value: "7",
				},
				{
					Type:  ADD,
					Value: "13",
				},
			},
		},
		{
			name: "Caug7",
			args: args{
				chord: "Caug7",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
				{
					Type:  AUG,
					Value: "7",
				},
			},
		},
		{
			name: "C7b5",
			args: args{
				chord: "C7b5",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
				{
					Type:  NUMBER,
					Value: "7",
				},
				{
					Type:  FLAT,
					Value: "5",
				},
			},
		},
		{
			name: "Cm(maj7)",
			args: args{
				chord: "Cm(maj7)",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
				{
					Type: MINOR,
				},
				{
					Type:  MAJ,
					Value: "7",
				},
			},
		},
		{
			name: "Cmaj9",
			args: args{
				chord: "Cmaj9",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
				{
					Type:  MAJ,
					Value: "9",
				},
			},
		},
		{
			name: "Cmaj13",
			args: args{
				chord: "Cmaj13",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
				{
					Type:  MAJ,
					Value: "13",
				},
			},
		},
		{
			name: "F##dim7",
			args: args{
				chord: "F##dim7",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "F##",
				},
				{
					Type:  DIM,
					Value: "7",
				},
			},
		},
		{
			name: "C7/Bb",
			args: args{
				chord: "C7/Bb",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
				{
					Type:  NUMBER,
					Value: "7",
				},
				{
					Type:  BASS,
					Value: "Bb",
				},
			},
		},
		{
			name: "Cm7/B#",
			args: args{
				chord: "Cm7/B#",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
				{
					Type:  MINOR,
					Value: "7",
				},
				{
					Type:  BASS,
					Value: "B#",
				},
			},
		},
		{
			name: "C7/G",
			args: args{
				chord: "C7/G",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
				{
					Type:  NUMBER,
					Value: "7",
				},
				{
					Type:  BASS,
					Value: "G",
				},
			},
		},
		{
			name: "Csus2",
			args: args{
				chord: "Csus2",
			},
			want: []Token{
				{
					Type:  ROOT,
					Value: "C",
				},
				{
					Type:  SUS,
					Value: "2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			l := Lexer{}

			got, err := l.Tokenize(tt.args.chord)
			if err != nil {
				t.Errorf("unexpected err %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
