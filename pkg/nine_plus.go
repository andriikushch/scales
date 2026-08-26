package scales

import "github.com/andriikushch/scales/pkg/internal"

// NewMessiaenMode3Scale builds Messiaen Mode 3 (9 notes).
func NewMessiaenMode3Scale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 3, 3, 4, 5, 6, 7, 7})
}

// NewMessiaenMode7Scale builds Messiaen Mode 7 (10 notes).
func NewMessiaenMode7Scale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 2, 3, 3, 4, 5, 6, 6, 7})
}

// NewChromaticScale builds the full 12-note chromatic scale.
func NewChromaticScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 2, 3, 3, 4, 5, 5, 6, 6, 7, 7})
}
