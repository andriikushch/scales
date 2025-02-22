package scales

import (
	"github.com/andriikushch/scales/pkg/internal"
)

func NewMajorScale(key string) (*Scale, error) {
	chords := []string{
		"maj",
		"m",
		"m",
		"maj",
		"maj",
		"min",
		"dim",
	}

	majorScaleStructure := []int{
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.HalfStep,
	}

	return newScale(key, majorScaleStructure, chords)
}
