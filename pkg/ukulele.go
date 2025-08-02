package scales

import (
	"strconv"
	"strings"

	"github.com/andriikushch/scales/pkg/internal"
)

type Ukulele struct {
	*stringInstrumentWithFrets
}

func (u *Ukulele) GetStringsAmount() int {
	return len(u.stringInstrumentWithFrets.tuning)
}

func (u *Ukulele) getChordShapes(chord Chord) []ChordShape {
	structure := make([]string, len(chord.structure))
	for k, v := range chord.structure {
		structure[k] = strconv.Itoa(v)
	}

	return UkuleleChordShapes[strings.Join(structure, "-")]
}

func NewUkuleleWithStandardTuning() *Ukulele {
	return &Ukulele{
		&stringInstrumentWithFrets{
			tuning: []Note{
				NewNote(internal.A),
				NewNote(internal.E),
				NewNote(internal.C),
				NewNote(internal.G),
			},
		},
	}
}
