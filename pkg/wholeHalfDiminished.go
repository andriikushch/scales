package scales

import "github.com/andriikushch/scales/pkg/internal"

func NewWholeHalfDiminishedScale(key string) (*Scale, error) {
	wholeHalfDiminishedScaleStructure := []int{
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
	}

	return newScale(key, wholeHalfDiminishedScaleStructure)
}
