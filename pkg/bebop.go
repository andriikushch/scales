package scales

import "github.com/andriikushch/scales/pkg/internal"

// NewBebopDominantScale builds the Bebop Dominant scale (Mixolydian plus a passing ♮7).
func NewBebopDominantScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 3, 4, 5, 6, 7, 7})
}

// NewBebopMajorScale builds the Bebop Major scale (Ionian plus a passing ♯5).
func NewBebopMajorScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 3, 4, 5, 5, 6, 7})
}

// NewBebopDorianScale builds the Bebop Dorian scale (Dorian plus a passing ♮3).
func NewBebopDorianScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 3, 3, 4, 5, 6, 7})
}

// NewBebopMelodicMinorScale builds the Bebop Melodic Minor scale (melodic minor plus a passing ♮6).
func NewBebopMelodicMinorScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 3, 4, 5, 6, 6, 7})
}

// NewBebopHarmonicMinorScale builds the Bebop Harmonic Minor scale (harmonic minor plus a passing ♮7).
func NewBebopHarmonicMinorScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 3, 4, 5, 6, 7, 7})
}
