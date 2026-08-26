package scales

import "github.com/andriikushch/scales/pkg/internal"

// NewDorianFlat2Scale builds the Dorian ♭2 mode of the melodic minor scale (2nd mode).
func NewDorianFlat2Scale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewLydianAugmentedScale builds the Lydian Augmented mode of the melodic minor scale (3rd mode).
func NewLydianAugmentedScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewLydianDominantScale builds the Lydian Dominant mode of the melodic minor scale (4th mode).
func NewLydianDominantScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewMixolydianFlat6Scale builds the Mixolydian ♭6 mode of the melodic minor scale (5th mode).
func NewMixolydianFlat6Scale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewLocrianNatural2Scale builds the Locrian ♮2 mode of the melodic minor scale (6th mode).
func NewLocrianNatural2Scale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewAlteredScale builds the Altered mode of the melodic minor scale (7th mode).
func NewAlteredScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.Step,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}
