package internal

import "math"

const mutedNote = -math.MaxInt

type ChordShape struct {
	Instrument       string
	ChordID          string
	Schema           []int
	RootNotePosition int
}

func NewChordShape(instrument string, chordID string, schema []int, rootNotePosition int) ChordShape {
	return ChordShape{
		Instrument:       instrument,
		ChordID:          chordID,
		Schema:           schema,
		RootNotePosition: rootNotePosition,
	}
}
