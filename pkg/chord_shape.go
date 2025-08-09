package scales

const MutedNote = -100

type ChordShape struct {
	Instrument        string
	ChordID           []int
	Schema            []int
	RootNotePosition  int
	PossibleNotations []string
}

func newChordShape(instrument string, chordID []int, schema []int, rootNotePosition int) ChordShape {
	return ChordShape{
		Instrument:       instrument,
		ChordID:          chordID,
		Schema:           schema,
		RootNotePosition: rootNotePosition,
	}
}
