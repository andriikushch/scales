package scales

import "github.com/andriikushch/scales/pkg/internal"

type Ukulele struct {
	*stringInstrumentWithFrets
}

func NewUkuleleWithStandardTuning() *Ukulele {
	return &Ukulele{&stringInstrumentWithFrets{
		tuning: []Note{
			NewNote(internal.A),
			NewNote(internal.E),
			NewNote(internal.C),
			NewNote(internal.G),
		}},
	}
}
