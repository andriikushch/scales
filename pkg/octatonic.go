package scales

import "github.com/andriikushch/scales/pkg/internal"

// NewHalfWholeDiminishedScale builds the Half-Whole Diminished scale.
func NewHalfWholeDiminishedScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 2, 3, 4, 5, 6, 7})
}

// NewMessiaenMode4Scale builds Messiaen Mode 4.
func NewMessiaenMode4Scale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 2, 4, 5, 5, 6, 7})
}

// NewMessiaenMode6Scale builds Messiaen Mode 6.
func NewMessiaenMode6Scale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 3, 4, 4, 5, 6, 7})
}

// NewSpanishEightToneScale builds the Spanish Eight-Tone scale.
func NewSpanishEightToneScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 3, 3, 4, 5, 6, 7})
}
