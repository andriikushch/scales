package scales

import (
	"strconv"
	"strings"
)

type BassGuitar struct {
	*stringInstrumentWithFrets
}

func (bg *BassGuitar) GetStringsAmount() int {
	return len(bg.stringInstrumentWithFrets.tuning)
}

func (bg *BassGuitar) getChordShapes(chord Chord) []ChordShape {
	structure := make([]string, len(chord.structure))
	for k, v := range chord.structure {
		structure[k] = strconv.Itoa(v)
	}

	return BassGuitarChordShapes[strings.Join(structure, "-")]
}
