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

	chords := []string{
		"dim",
		"mb5",
		"dim",
		"mb5",
		"dim",
		"mb5",
		"dim",
		"mb5",
	}

	return newScale(key, wholeHalfDiminishedScaleStructure, chords)
}
