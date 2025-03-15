package scales

import (
	"errors"
	"github.com/andriikushch/scales/pkg/internal"
	"slices"
	"strings"
)

var errBassNoteIsEmpty = errors.New("bass note is empty")

type chord struct {
	Name           string
	ChordBasicType string
	Root           Note
	Notes          []Note
	Structure      []int
	Type           []string
	BassNote       Note
}

func (c *chord) setName(name string) {
	c.Name = name
}

func (c *chord) addType(t string) {
	c.Type = append(c.Type, t)
}

func (c *chord) addRoot(note Note) {
	c.Root = note
	c.Structure = append(c.Structure, internal.IUnison)
}

func (c *chord) add(interval int) error {
	if len(c.Root.Name) == 0 {
		return errors.New("set the root first")
	}

	if len(c.ChordBasicType) == 0 {
		return errors.New("set the chord type first")
	}

	c.Structure = append(c.Structure, interval)

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
	c.ChordBasicType = s
}

func (c *chord) flatFirst(interval int) {
	for i, v := range c.Structure {
		if v == interval {
			c.Structure[i] -= internal.HalfStep

			break
		}
	}
}

func (c *chord) sharpFirst(interval int) {
	for i, v := range c.Structure {
		if v == interval {
			c.Structure[i] += internal.HalfStep

			break
		}
	}
}

func (c *chord) finish() error {
	for intervalIndex, interval := range c.Structure {
		nextPossibleNotes := defaultChromaticScale.next(c.Root, interval)

		// not reliable, trying to guess sign from the previous Type value
		var isFlat, isSharp bool
		if intervalIndex < len(c.Type)-1 {
			isFlat = c.Type[intervalIndex] == internal.Flat
			isSharp = c.Type[intervalIndex] == internal.Sharp
		}

		nextNote := c.findOptimalNoteForTheChord(nextPossibleNotes, interval, c.Root.Name, c.ChordBasicType, isFlat, isSharp)
		c.Notes = append(c.Notes, nextNote)
	}

	if c.BassNote.Name != "" {
		if baseNoteIndex := slices.Index(c.Notes, c.BassNote); baseNoteIndex != -1 {
			baseNoteInterval := c.Structure[baseNoteIndex]
			bassNote := c.Notes[baseNoteIndex]

			// remove elements
			c.Structure = append(c.Structure[:baseNoteIndex], c.Structure[baseNoteIndex+1:]...)
			c.Notes = append(c.Notes[:baseNoteIndex], c.Notes[baseNoteIndex+1:]...)

			c.Structure = append([]int{baseNoteInterval}, c.Structure...)
			c.Notes = append([]Note{bassNote}, c.Notes...)
		} else {
			// todo: is it valid case? Bass note is not a part of the chord.
			return errBassNoteIsEmpty
		}
	}

	return nil
}

func (c *chord) setBase(note Note) {
	c.BassNote = note
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
