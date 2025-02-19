package scales

import "github.com/andriikushch/scales/pkg/internal/chromatic"

func NewNaturalMinorScale(key string) (*Scale, error) {
	naturalMinorScaleStructure := []int{
		chromatic.Step,
		chromatic.HalfStep,
		chromatic.Step,
		chromatic.Step,
		chromatic.HalfStep,
		chromatic.Step,
		chromatic.Step,
	}

	return newScale(key, naturalMinorScaleStructure)
}

func NewHarmonicMinorScale(key string) (*Scale, error) {
	harmonicMinorScaleStructure := []int{
		chromatic.Step,
		chromatic.HalfStep,
		chromatic.Step,
		chromatic.Step,
		chromatic.HalfStep,
		chromatic.Step + chromatic.HalfStep,
		chromatic.HalfStep,
	}

	return newScale(key, harmonicMinorScaleStructure)
}

func NewMelodicMinorScale(key string) (*Scale, error) {
	melodicMinorScaleStructure := []int{
		chromatic.Step,
		chromatic.HalfStep,
		chromatic.Step,
		chromatic.Step,
		chromatic.Step,
		chromatic.Step,
		chromatic.HalfStep,
	}

	return newScale(key, melodicMinorScaleStructure)
}
