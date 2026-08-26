package scales

import "github.com/andriikushch/scales/pkg/internal"

// NewDoubleHarmonicMajorScale builds the double harmonic major scale (1st mode).
func NewDoubleHarmonicMajorScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewLydianSharp2Sharp6Scale builds the Lydian ♯2 ♯6 mode of the double harmonic major scale (2nd mode).
func NewLydianSharp2Sharp6Scale(key string) (*Scale, error) {
	structure := []int{
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewUltraphrygianScale builds the Ultraphrygian mode of the double harmonic major scale (3rd mode).
func NewUltraphrygianScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewHungarianMinorScale builds the Hungarian Minor mode of the double harmonic major scale (4th mode).
func NewHungarianMinorScale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewOrientalScale builds the Oriental mode of the double harmonic major scale (5th mode).
func NewOrientalScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewIonianSharp2Sharp5Scale builds the Ionian ♯2 ♯5 mode of the double harmonic major scale (6th mode).
func NewIonianSharp2Sharp5Scale(key string) (*Scale, error) {
	structure := []int{
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewLocrianDoubleFlat3DoubleFlat7Scale builds the Locrian ♭♭3 ♭♭7 mode of the double harmonic major scale (7th mode).
func NewLocrianDoubleFlat3DoubleFlat7Scale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}
