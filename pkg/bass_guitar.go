package scales

import (
	"strconv"
	"strings"
)

type BassGuitar struct {
	*stringInstrumentWithFrets
}

func (bg *BassGuitar) getChordShapes(chord Chord) []chordShape {
	structure := make([]string, len(chord.structure))
	for k, v := range chord.structure {
		structure[k] = strconv.Itoa(v)
	}

	return bassGuitarChordShapes[strings.Join(structure, "-")]
}
