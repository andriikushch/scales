package scales

import "github.com/andriikushch/scales/pkg/internal"

// NewDorianScale builds the Dorian mode of the major scale (2nd mode).
func NewDorianScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewPhrygianScale builds the Phrygian mode of the major scale (3rd mode).
func NewPhrygianScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewLydianScale builds the Lydian mode of the major scale (4th mode).
func NewLydianScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewMixolydianScale builds the Mixolydian mode of the major scale (5th mode).
func NewMixolydianScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewLocrianScale builds the Locrian mode of the major scale (7th mode).
func NewLocrianScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}
