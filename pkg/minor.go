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

	return newScale(key, naturalMinorScaleStructure)
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

	return newScale(key, harmonicMinorScaleStructure)
}

func NewMelodicMinorScale(key string) (*Scale, error) {
	melodicMinorScaleStructure := []int{
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.HalfStep,
	}

	return newScale(key, melodicMinorScaleStructure)
}
