package scales

import "github.com/andriikushch/scales/pkg/internal"

// NewNeapolitanMinorScale builds the Neapolitan Minor scale.
func NewNeapolitanMinorScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewNeapolitanMajorScale builds the Neapolitan Major scale.
func NewNeapolitanMajorScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewLeadingWholeToneScale builds the Leading Whole Tone scale.
func NewLeadingWholeToneScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewMajorLocrianScale builds the Major Locrian scale.
func NewMajorLocrianScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewGypsyScale builds the Gypsy scale.
func NewGypsyScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewHungarianMajorScale builds the Hungarian Major scale.
func NewHungarianMajorScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewPersianScale builds the Persian scale.
func NewPersianScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewEnigmaticScale builds the Enigmatic scale.
func NewEnigmaticScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}
