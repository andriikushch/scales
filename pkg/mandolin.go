package scales

import (
	"strconv"
	"strings"

	"github.com/andriikushch/scales/pkg/internal"
)

type Mandolin struct {
	*stringInstrumentWithFrets
}

func (m *Mandolin) GetStringsAmount() int {
	return len(m.stringInstrumentWithFrets.tuning)
}

func (m *Mandolin) getChordShapes(chord Chord) []ChordShape {
	structure := make([]string, len(chord.structure))
	for k, v := range chord.structure {
		structure[k] = strconv.Itoa(v)
	}

	return MandolinChordShapes[strings.Join(structure, "-")]
}

func NewMandolinWithStandardTuning() *Mandolin {
	return &Mandolin{
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
