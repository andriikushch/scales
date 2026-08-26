package scales

import (
	"fmt"
	"slices"
)

// letters is the natural note-letter cycle used to spell scale degrees.
var letters = []string{"C", "D", "E", "F", "G", "A", "B"}

// naturalSemitone is the chromatic position (0-11, C=0) of each letter with no accidental.
var naturalSemitone = map[string]int{
	"C": 0,
	"D": 2,
	"E": 4,
	"F": 5,
	"G": 7,
	"A": 9,
	"B": 11,
}

// heptatonicDegrees is the degree sequence shared by every 7-note scale: each of the 7
// scale degrees appears exactly once, only the accidentals (encoded in each scale's
// structure) differ.
var heptatonicDegrees = []int{1, 2, 3, 4, 5, 6, 7}

// newScaleFromDegrees builds a scale from a semitone structure and a parallel list of
// scale-degree numbers (1-7, taken from a scale's Degrees column, ignoring accidentals).
//
// Unlike newScale, the letter for each note is derived directly from its degree number
// rather than by avoiding a previously used letter, so it correctly handles scales that
// intentionally reuse a letter (e.g. bebop passing tones) or need double-flat spellings
// (e.g. F♭, B♭♭) that aren't present in the chromatic lookup table newScale relies on.
func newScaleFromDegrees(key string, structure []int, degrees []int) (*Scale, error) {
	rootLetter := string(key[0])
	rootLetterIndex := slices.Index(letters, rootLetter)
	if rootLetterIndex == -1 {
		return nil, fmt.Errorf("error while building scale %s: unrecognized root letter in key %q", key, key)
	}

	rootAccidental := 0
	for _, r := range key[1:] {
		switch r {
		case '#':
			rootAccidental++
		case 'b':
			rootAccidental--
		}
	}
	rootPosition := mod12(naturalSemitone[rootLetter] + rootAccidental)

	notes := make([]Note, len(degrees))
	cumulative := 0

	for i, degree := range degrees {
		letter := letters[mod7(rootLetterIndex+degree-1)]
		targetPosition := mod12(rootPosition + cumulative)

		diff := mod12(targetPosition - naturalSemitone[letter])
		if diff > 6 {
			diff -= 12
		}

		var accidental string
		switch diff {
		case 0:
			accidental = ""
		case 1:
			accidental = "#"
		case 2:
			accidental = "##"
		case -1:
			accidental = "b"
		case -2:
			accidental = "bb"
		default:
			return nil, fmt.Errorf("error while building scale %s, %v, %v: degree %d needs unsupported accidental offset %d", key, structure, degrees, degree, diff)
		}

		notes[i] = NewNote(letter + accidental)

		if i < len(structure) {
			cumulative += structure[i]
		}
	}

	return &Scale{notes: notes}, nil
}

func mod12(n int) int {
	return ((n % 12) + 12) % 12
}

func mod7(n int) int {
	return ((n % 7) + 7) % 7
}
