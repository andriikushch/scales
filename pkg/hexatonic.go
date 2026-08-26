package scales

import "github.com/andriikushch/scales/pkg/internal"

// NewWholeToneScale builds the Whole Tone scale.
func NewWholeToneScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 3, 4, 5, 6})
}

// NewAugmentedScale builds the Augmented scale.
func NewAugmentedScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, []int{1, 3, 3, 5, 6, 7})
}

// NewBluesScale builds the Blues scale.
func NewBluesScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step + internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, []int{1, 3, 4, 5, 5, 7})
}

// NewMajorBluesScale builds the Major Blues scale.
func NewMajorBluesScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.Step,
		internal.Step + internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 3, 3, 5, 6})
}

// NewMajorHexatonicScale builds the Major Hexatonic scale (the major scale without its 7th degree).
func NewMajorHexatonicScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.Step + internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 3, 4, 5, 6})
}

// NewMinorHexatonicScale builds the Minor Hexatonic scale (the natural minor scale without its 6th degree).
func NewMinorHexatonicScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.Step + internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 3, 4, 5, 7})
}

// NewPrometheusScale builds the Prometheus scale.
func NewPrometheusScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.Step,
		internal.Step,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 3, 4, 6, 7})
}

// NewPrometheusNeapolitanScale builds the Prometheus Neapolitan scale.
func NewPrometheusNeapolitanScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.Step,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 3, 4, 6, 7})
}

// NewTritoneScale builds the Tritone scale.
func NewTritoneScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 3, 5, 5, 7})
}

// NewIstrianScale builds the Istrian scale.
func NewIstrianScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.Step + internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 3, 4, 5, 5})
}

// NewScaleOfHarmonicsScale builds the Scale of Harmonics.
func NewScaleOfHarmonicsScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.Step + internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, []int{1, 3, 3, 4, 5, 6})
}

// NewTwoSemitoneTritoneScale builds the Two-Semitone Tritone scale.
func NewTwoSemitoneTritoneScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.HalfStep,
		internal.Step + internal.Step,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step + internal.Step,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 2, 4, 5, 6})
}

// NewMessiaenMode5Scale builds Messiaen Mode 5.
func NewMessiaenMode5Scale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step + internal.Step,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step + internal.Step,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, []int{1, 2, 4, 5, 5, 7})
}
