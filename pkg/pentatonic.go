package scales

import (
	"slices"
	"strings"

	"github.com/andriikushch/scales/pkg/internal"
)

type Pentatonic struct {
	notes []Note
}

func (p *Pentatonic) GetNotes() []Note {
	return p.notes
}

func (p *Pentatonic) String() string {
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

// NewSuspendedPentatonicScale builds the Suspended Pentatonic scale.
func NewSuspendedPentatonicScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.Step + internal.HalfStep,
		internal.Step,
		internal.Step + internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 4, 5, 7})
}

// NewBluesMinorPentatonicScale builds the Blues Minor Pentatonic scale.
func NewBluesMinorPentatonicScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step + internal.HalfStep,
		internal.Step,
		internal.Step + internal.HalfStep,
		internal.Step,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, []int{1, 3, 4, 6, 7})
}

// NewBluesMajorPentatonicScale builds the Blues Major Pentatonic scale.
func NewBluesMajorPentatonicScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.Step + internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.Step + internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 4, 5, 6})
}

// NewIwatoScale builds the Iwato scale.
func NewIwatoScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step + internal.Step,
		internal.HalfStep,
		internal.Step + internal.Step,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 4, 5, 7})
}

// NewMiyakoBushiScale builds the Miyako-bushi scale.
func NewMiyakoBushiScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step + internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.Step,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 4, 5, 6})
}

// NewKumoiScale builds the Kumoi scale.
func NewKumoiScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.Step,
		internal.Step,
		internal.Step + internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 3, 5, 6})
}

// NewInsenScale builds the Insen scale.
func NewInsenScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step + internal.Step,
		internal.Step,
		internal.Step + internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 4, 5, 7})
}

// NewChinesePentatonicScale builds the Chinese pentatonic scale (C E F♯ G B when rooted on C).
//
// This is distinct from the "Chinese" alias found in some sources for the major pentatonic
// scale (see NewMajorPentatonicScale) — the two names collide, but describe different note
// collections.
func NewChinesePentatonicScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step + internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.Step,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, []int{1, 3, 4, 5, 7})
}

// NewBalinesePelogScale builds the Balinese Pelog scale (a 12-tone-equal-temperament approximation).
func NewBalinesePelogScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step,
		internal.Step + internal.Step,
		internal.HalfStep,
		internal.Step + internal.Step,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 3, 5, 6})
}
