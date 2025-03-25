package scales

type Note struct {
	Name string
}

func (n Note) Equal(other Note) bool {
	if n.Name == other.Name {
		return true
	}

	normalizedNote := defaultChromaticScale.normalize(n)
	normalizedOtherNote := defaultChromaticScale.normalize(other)

	notePosition := defaultChromaticScale.position(normalizedNote)
	otherNotePosition := defaultChromaticScale.position(normalizedOtherNote)

	if notePosition == -1 {
		return false
	}

	return notePosition == otherNotePosition
}

func (n Note) String() string {
	return n.Name
}

func NewNote(name string) Note {
	return Note{
		Name: name,
	}
}
