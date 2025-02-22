package scales

import (
	"github.com/andriikushch/scales/pkg/internal"
)

func NewNaturalMinorScale(key string) (*Scale, error) {
	naturalMinorScaleStructure := []int{
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
	}

	chords := []string{
		"m",
		"dim",
		"maj",
		"m",
		"m",
		"maj",
		"maj",
	}

	return newScale(key, naturalMinorScaleStructure, chords)
}

func NewHarmonicMinorScale(key string) (*Scale, error) {
	harmonicMinorScaleStructure := []int{
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
	}

	chords := []string{
		"m",
		"dim",
		"aug",
		"m",
		"maj",
		"maj",
		"dim",
	}

	return newScale(key, harmonicMinorScaleStructure, chords)
}

func NewMelodicMinorScale(key string) (*Scale, error) {
	chords := []string{
		"m",
		"min",
		"aug",
		"maj",
		"maj",
		"dim",
		"dim",
	}

	melodicMinorScaleStructure := []int{
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.HalfStep,
	}

	return newScale(key, melodicMinorScaleStructure, chords)
}
