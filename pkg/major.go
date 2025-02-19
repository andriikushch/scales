package scales

import "github.com/andriikushch/scales/pkg/internal/chromatic"

func NewMajorScale(key string) (*Scale, error) {
	majorScaleStructure := []int{
		chromatic.Step,
		chromatic.Step,
		chromatic.HalfStep,
		chromatic.Step,
		chromatic.Step,
		chromatic.Step,
		chromatic.HalfStep,
	}

	return newScale(key, majorScaleStructure)
}
