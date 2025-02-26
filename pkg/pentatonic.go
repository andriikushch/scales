package scales

import (
	"slices"
	"strings"
)

type Pentatonic struct {
	notes []Note
}

func (p Pentatonic) GetNotes() []Note {
	return p.notes
}

func (p Pentatonic) String() string {
	buf := new(strings.Builder)

	for i, note := range p.notes {
		buf.WriteString(note.Name)

		if i != len(p.notes)-1 {
			buf.WriteString(", ")
		}
	}

	return buf.String()
}

func NewMajorPentatonicScale(key string) (*Pentatonic, error) {
	majorScale, err := NewMajorScale(key)
	if err != nil {
		return nil, err
	}

	// remove 4th and 7th degrees
	majorScale.notes = slices.Delete(majorScale.notes, 6, 7)
	majorScale.notes = slices.Delete(majorScale.notes, 3, 4)

	return &Pentatonic{
		notes: majorScale.notes,
	}, nil
}

func NewMinorPentatonicScale(key string) (*Pentatonic, error) {
	naturalMinorScale, err := NewNaturalMinorScale(key)
	if err != nil {
		return nil, err
	}

	// remove 2d and 6th degrees
	naturalMinorScale.notes = slices.Delete(naturalMinorScale.notes, 5, 6)
	naturalMinorScale.notes = slices.Delete(naturalMinorScale.notes, 1, 2)

	return &Pentatonic{
		notes: naturalMinorScale.notes,
	}, nil
}
