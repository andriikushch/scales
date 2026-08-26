package scales

import "github.com/andriikushch/scales/pkg/internal"

// NewLocrianNatural6Scale builds the Locrian ♮6 mode of the harmonic minor scale (2nd mode).
func NewLocrianNatural6Scale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewIonianSharp5Scale builds the Ionian ♯5 mode of the harmonic minor scale (3rd mode).
func NewIonianSharp5Scale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewDorianSharp4Scale builds the Dorian ♯4 mode of the harmonic minor scale (4th mode).
func NewDorianSharp4Scale(key string) (*Scale, error) {
	structure := []int{
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewPhrygianDominantScale builds the Phrygian Dominant mode of the harmonic minor scale (5th mode).
func NewPhrygianDominantScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewLydianSharp2Scale builds the Lydian ♯2 mode of the harmonic minor scale (6th mode).
func NewLydianSharp2Scale(key string) (*Scale, error) {
	structure := []int{
		internal.Step + internal.HalfStep,
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}

// NewUltralocrianScale builds the Ultralocrian mode of the harmonic minor scale (7th mode).
func NewUltralocrianScale(key string) (*Scale, error) {
	structure := []int{
		internal.HalfStep,
		internal.Step,
		internal.HalfStep,
		internal.Step,
		internal.Step,
		internal.HalfStep,
		internal.Step + internal.HalfStep,
	}

	return newScaleFromDegrees(key, structure, heptatonicDegrees)
}
