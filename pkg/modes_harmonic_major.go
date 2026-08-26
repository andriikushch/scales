package scales

import "github.com/andriikushch/scales/pkg/internal"

// NewHarmonicMajorScale builds the harmonic major scale (1st mode).
func NewHarmonicMajorScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewDorianFlat5Scale builds the Dorian ♭5 mode of the harmonic major scale (2nd mode).
func NewDorianFlat5Scale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewPhrygianFlat4Scale builds the Phrygian ♭4 mode of the harmonic major scale (3rd mode).
func NewPhrygianFlat4Scale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewLydianFlat3Scale builds the Lydian ♭3 mode of the harmonic major scale (4th mode).
func NewLydianFlat3Scale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewMixolydianFlat2Scale builds the Mixolydian ♭2 mode of the harmonic major scale (5th mode).
func NewMixolydianFlat2Scale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewLydianAugmentedSharp2Scale builds the Lydian Augmented ♯2 mode of the harmonic major scale (6th mode).
func NewLydianAugmentedSharp2Scale(key string) (*Scale, error) {
	structure := []int{
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewLocrianDoubleFlat7Scale builds the Locrian ♭♭7 mode of the harmonic major scale (7th mode).
func NewLocrianDoubleFlat7Scale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}
