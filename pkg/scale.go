package scales

import (
	"fmt"
	"slices"
	"strings"
)

type Scale struct {
	notes []Note
}

func (s Scale) GetNotes() []Note {
	return s.notes
}

func (s Scale) String() string {
	buf := new(strings.Builder)

	for i, note := range s.notes {
		buf.WriteString(note.Name)

		if i != len(s.notes)-1 {
			buf.WriteString(", ")
		}
	}

	return buf.String()
}

func newScale(key string, scaleStructure []int) (*Scale, error) {
	scale := &Scale{}

	chromaticScale := defaultChromaticScale
	currentNote := NewNote(key)

	scale.notes = append(scale.notes, currentNote)

	// to track letters that are already used
	var usedPrefixes []string

	// last note is the same as a first one, we can skip it with "len(scaleStructure)-1" condition
	for i := 0; i < len(scaleStructure)-1; i++ {
		step := scaleStructure[i]
		nextNotes := chromaticScale.next(currentNote, step)

		var nextNote Note

		if len(nextNotes) == 1 {
			nextNote = nextNotes[0]
		} else {
			firstLetterInCurrentNote := string(currentNote.Name[0])
			usedPrefixes = append(usedPrefixes, firstLetterInCurrentNote)

			// use sharp, flats, double sharps to keep letter uniq in the scale
			for _, nextNote = range nextNotes {
				if slices.Index(usedPrefixes, string(nextNote.Name[0])) == -1 {
					break
				}
			}
		}

		if len(nextNote.Name) == 0 {
			return scale, fmt.Errorf("error while building scale %s, %v, note after : %s", key, scaleStructure, currentNote.Name)
		}

		scale.notes = append(scale.notes, nextNote)

		currentNote = nextNote
	}

	return scale, nil
}
