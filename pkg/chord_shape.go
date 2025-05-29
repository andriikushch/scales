package scales

import "math"

const MutedNote = -math.MaxInt

type chordShape struct {
	Instrument       string
	ChordID          []int
	Schema           []int
	RootNotePosition int
}

func newChordShape(instrument string, chordID []int, schema []int, rootNotePosition int) chordShape {
	return chordShape{
		Instrument:       instrument,
		ChordID:          chordID,
		Schema:           schema,
		RootNotePosition: rootNotePosition,
	}
}
