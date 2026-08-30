package scales

import (
	"github.com/andriikushch/scales/pkg/internal"
)

func NewMajorScale(key string) (*Scale, error) {
	majorScaleStructure := []int{
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.HalfStep,
	}

	return newScale(key, majorScaleStructure)
}
