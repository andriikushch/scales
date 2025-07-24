package scales

import (
	"strconv"
	"strings"

	"github.com/andriikushch/scales/pkg/internal"
)

type Mandolin struct {
	*stringInstrumentWithFrets
}

func (u *Mandolin) getChordShapes(chord Chord) []chordShape {
	structure := make([]string, len(chord.structure))
	for k, v := range chord.structure {
		structure[k] = strconv.Itoa(v)
	}

	return mandolinChordShapes[strings.Join(structure, "-")]
}

func NewMandolinWithStandardTuning() *Ukulele {
	return &Ukulele{
		&stringInstrumentWithFrets{
			tuning: []Note{
				NewNote(internal.E),
				NewNote(internal.A),
				NewNote(internal.D),
				NewNote(internal.G),
			},
		},
	}
}
