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
	majorScale.Notes = slices.Delete(majorScale.Notes, 6, 7)
	majorScale.Notes = slices.Delete(majorScale.Notes, 3, 4)

	return majorScale, nil
}

func NewMinorPentatonicScale(key string) (*Scale, error) {
	naturalMinorScale, err := NewNaturalMinorScale(key)
	if err != nil {
		return nil, err
	}

	// remove 2d and 6th degrees
	naturalMinorScale.Notes = slices.Delete(naturalMinorScale.Notes, 5, 6)
	naturalMinorScale.Notes = slices.Delete(naturalMinorScale.Notes, 1, 2)

	return naturalMinorScale, nil
}
