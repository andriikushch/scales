package scales

import "github.com/andriikushch/scales/pkg/internal"

func NewUkuleleWithStandardTuning() *Guitar {
	return &Guitar{
		tuning: []Note{
			NewNote(internal.A),
			NewNote(internal.E),
			NewNote(internal.C),
			NewNote(internal.G),
		},
	}
}
