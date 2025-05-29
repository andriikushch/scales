package scales

import (
	"strconv"
	"strings"

	"github.com/andriikushch/scales/pkg/internal"
)

type Ukulele struct {
	*stringInstrumentWithFrets
}

func (u *Ukulele) getChordShapes(chord Chord) []chordShape {
	structure := make([]string, len(chord.structure))
	for k, v := range chord.structure {
		structure[k] = strconv.Itoa(v)
	}

	return ukuleleChordShapes[strings.Join(structure, "-")]
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
