package scales

import (
	"strconv"
	"strings"

	"github.com/andriikushch/scales/pkg/internal"
)

type Guitar struct {
	*stringInstrumentWithFrets
}

func (g *Guitar) getChordShapes(chord Chord) []chordShape {
	structure := make([]string, len(chord.structure))
	for k, v := range chord.structure {
		structure[k] = strconv.Itoa(v)
	}

	return guitarChordShapes[strings.Join(structure, "-")]
}

func NewGuitarWithStandardTuning() *Guitar {
	return &Guitar{
		&stringInstrumentWithFrets{
			tuning: []Note{
				NewNote(internal.E),
				NewNote(internal.B),
				NewNote(internal.G),
				NewNote(internal.D),
				NewNote(internal.A),
				NewNote(internal.E),
			},
		},
	}
}

func NewBassGuitarWithStandardTuning() *BassGuitar {
	return &BassGuitar{&stringInstrumentWithFrets{
		tuning: []Note{
			NewNote(internal.G),
			NewNote(internal.D),
			NewNote(internal.A),
			NewNote(internal.E),
		},
	}}
}
