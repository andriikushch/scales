package scales

import (
	"slices"
)

func NewMajorPentatonicScale(key string) (*Scale, error) {
	majorScale, err := NewMajorScale(key)
	if err != nil {
		return nil, err
	}

	// remove 4th and 7th degrees
	majorScale.notes = slices.Delete(majorScale.notes, 6, 7)
	majorScale.notes = slices.Delete(majorScale.notes, 3, 4)

	return majorScale, nil
}

func NewMinorPentatonicScale(key string) (*Scale, error) {
	naturalMinorScale, err := NewNaturalMinorScale(key)
	if err != nil {
		return nil, err
	}

	// remove 2d and 6th degrees
	naturalMinorScale.notes = slices.Delete(naturalMinorScale.notes, 5, 6)
	naturalMinorScale.notes = slices.Delete(naturalMinorScale.notes, 1, 2)

	return naturalMinorScale, nil
}
