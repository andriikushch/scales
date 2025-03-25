package scales

import (
	"errors"
	"github.com/andriikushch/scales/pkg/internal"
	"slices"
	"strings"
)

var errBassNoteIsEmpty = errors.New("bass note is empty")

type chord struct {
	name           string
	chordBasicType string
	root           Note
	notes          []Note
	structure      []int
	cType          []string
	bassNote       Note
}

func (c *chord) setName(name string) {
	c.name = name
}

func (c *chord) addType(t string) {
	c.cType = append(c.cType, t)
}

func (c *chord) addRoot(note Note) {
	c.root = note
	c.structure = append(c.structure, internal.IUnison)
}

func (c *chord) add(interval int) error {
	if len(c.root.Name) == 0 {
		return errors.New("set the root first")
	}

	if len(c.chordBasicType) == 0 {
		return errors.New("set the chord type first")
	}

	c.structure = append(c.structure, interval)

	return nil
}

func (c *chord) addIntervals(interval ...int) error {
	for _, v := range interval {
		err := c.add(v)
		if err != nil {
			return err
		}
	}

	return nil
}

// findOptimalNoteForTheChord is opinionated selection between multiple annotations of the same note. TODO: improve this logic if necessary.
func (c *chord) findOptimalNoteForTheChord(notes []Note, step int, key string, chordType string, isFlat, isSharp bool) Note {
	if isSharp {
		n := c.getWithSharps(notes, 1)
		if len(n.Name) > 0 {
			return n
		}
	}

	if isFlat {
		n := c.getWithFlats(notes, 1)
		if len(n.Name) > 0 {
			return n
		}
	}

	var scaleNotes []Note
	// try to understand if it is a minor or major chord, it is not
	if chordType == internal.Minor {
		// it is minor, let's build the scale and see if any of the notes annotations are in the scale
		sc, _ := NewNaturalMinorScale(key)
		scaleNotes = sc.GetNotes()
	} else {
		// maybe it is major, let's build the scale and see if any of the notes annotations are in the scale
		sc, _ := NewMajorScale(key)
		scaleNotes = sc.GetNotes()
	}

	for _, note := range notes {
		if slices.Index(scaleNotes, note) != -1 {
			return note
		}
	}

	switch step {
	case internal.Im7, internal.ID5, internal.Im3, internal.Im9:
		// let's search for a flat for minor 7
		n := c.getWithFlats(notes, 1)
		if len(n.Name) > 0 {
			return n
		}
	case internal.IA5:
		n := c.getWithSharps(notes, 1)
		if len(n.Name) > 0 {
			return n
		}
	}

	return notes[0]
}

func (c *chord) setChordBasicType(s string) {
	c.chordBasicType = s
}

func (c *chord) flatFirst(interval int) {
	for i, v := range c.structure {
		if v == interval {
			c.structure[i] -= internal.HalfStep

			break
		}
	}
}

func (c *chord) sharpFirst(interval int) {
	for i, v := range c.structure {
		if v == interval {
			c.structure[i] += internal.HalfStep

			break
		}
	}
}

func (c *chord) finish() error {
	for intervalIndex, interval := range c.structure {
		nextPossibleNotes := defaultChromaticScale.next(c.root, interval)

		// not reliable, trying to guess sign from the previous Type value
		var isFlat, isSharp bool
		if intervalIndex < len(c.cType)-1 {
			isFlat = c.cType[intervalIndex] == internal.Flat
			isSharp = c.cType[intervalIndex] == internal.Sharp
		}

		nextNote := c.findOptimalNoteForTheChord(nextPossibleNotes, interval, c.root.Name, c.chordBasicType, isFlat, isSharp)
		c.notes = append(c.notes, nextNote)
	}

	if c.bassNote.Name != "" {
		if baseNoteIndex := slices.Index(c.notes, c.bassNote); baseNoteIndex != -1 {
			baseNoteInterval := c.structure[baseNoteIndex]
			bassNote := c.notes[baseNoteIndex]

			// remove elements
			c.structure = append(c.structure[:baseNoteIndex], c.structure[baseNoteIndex+1:]...)
			c.notes = append(c.notes[:baseNoteIndex], c.notes[baseNoteIndex+1:]...)

			c.structure = append([]int{baseNoteInterval}, c.structure...)
			c.notes = append([]Note{bassNote}, c.notes...)
		} else {
			// todo: is it valid case? Bass note is not a part of the chord.
			return errBassNoteIsEmpty
		}
	}

	return nil
}

func (c *chord) setBase(note Note) {
	c.bassNote = note
}

func (c *chord) getWithFlats(notes []Note, n int) Note {
	for _, note := range notes {
		if strings.Count(note.Name, "b") == n {
			return note
		}
	}

	return Note{}
}

func (c *chord) getWithSharps(notes []Note, n int) Note {
	for _, note := range notes {
		if strings.Count(note.Name, "#") == n {
			return note
		}
	}

	return Note{}
}
