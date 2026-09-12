package main

import (
	"fmt"
	"strings"

	scales "github.com/andriikushch/scales/pkg"
)

// letterIndex gives each natural letter's position in the musical alphabet (C=0..B=6), used to
// compute scale-degree numbers by letter distance from the root.
var letterIndex = map[byte]int{'C': 0, 'D': 1, 'E': 2, 'F': 3, 'G': 4, 'A': 5, 'B': 6}

// naturalSemitone gives each natural letter's semitone position (C=0..B=11), used to compute the
// Structure column.
var naturalSemitone = map[byte]int{'C': 0, 'D': 2, 'E': 4, 'F': 5, 'G': 7, 'A': 9, 'B': 11}

// parseNote splits a Note.Name (e.g. "F##", "Bbb", "C") into its letter and ASCII accidental.
func parseNote(name string) (letter byte, accidental string) {
	return name[0], name[1:]
}

// accidentalOffset returns the semitone shift for an ASCII accidental string.
func accidentalOffset(accidental string) int {
	switch accidental {
	case "#":
		return 1
	case "##":
		return 2
	case "b":
		return -1
	case "bb":
		return -2
	default:
		return 0
	}
}

// accidentalSymbol translates an ASCII accidental to its Unicode music-notation glyph.
func accidentalSymbol(accidental string) string {
	switch accidental {
	case "#":
		return "♯"
	case "##":
		return "♯♯"
	case "b":
		return "♭"
	case "bb":
		return "♭♭"
	default:
		return ""
	}
}

// mod12 normalizes a semitone difference into [0, 12).
func mod12(n int) int {
	return ((n % 12) + 12) % 12
}

// semitone returns a note's 0-11 pitch-class position.
func semitone(letter byte, accidental string) int {
	return mod12(naturalSemitone[letter] + accidentalOffset(accidental))
}

// unicodeNoteName renders a Note.Name with its accidental translated to Unicode, e.g.
// "Bbb" -> "B♭♭".
func unicodeNoteName(name string) string {
	letter, accidental := parseNote(name)
	return string(letter) + accidentalSymbol(accidental)
}

// structure returns the semitone gap between each consecutive note, wrapping from the last note
// back to the root, e.g. major scale -> [2 2 1 2 2 2 1].
func structure(notes []scales.Note) []int {
	semitones := make([]int, len(notes))
	for i, n := range notes {
		letter, accidental := parseNote(n.Name)
		semitones[i] = semitone(letter, accidental)
	}

	gaps := make([]int, len(notes))
	for i := range notes {
		next := semitones[(i+1)%len(semitones)]
		gaps[i] = mod12(next - semitones[i])
	}

	return gaps
}

// structureString renders structure() as "2-2-1-2-2-2-1".
func structureString(notes []scales.Note) string {
	gaps := structure(notes)
	parts := make([]string, len(gaps))
	for i, g := range gaps {
		parts[i] = fmt.Sprintf("%d", g)
	}

	return strings.Join(parts, "-")
}

// degreesString renders each note as a scale-degree number with its accidental, e.g.
// "1 ♭3 ♮3 5 ♭6 7". The degree number comes from the note's letter distance from the root
// letter; an explicit natural symbol (♮) is added only when a degree number repeats within the
// scale and this occurrence is otherwise unaltered, matching scales like Bebop Dominant's
// "♭7 ♮7" and Augmented's "♭3 ♮3".
func degreesString(notes []scales.Note) string {
	rootLetter, _ := parseNote(notes[0].Name)
	rootIdx := letterIndex[rootLetter]

	degreeNumbers := make([]int, len(notes))
	accidentals := make([]string, len(notes))
	counts := map[int]int{}

	for i, n := range notes {
		letter, accidental := parseNote(n.Name)
		degreeNumbers[i] = ((letterIndex[letter]-rootIdx)%7+7)%7 + 1
		accidentals[i] = accidental
		counts[degreeNumbers[i]]++
	}

	parts := make([]string, len(notes))
	for i := range notes {
		sym := accidentalSymbol(accidentals[i])
		if sym == "" && counts[degreeNumbers[i]] > 1 {
			sym = "♮"
		}
		parts[i] = fmt.Sprintf("%s%d", sym, degreeNumbers[i])
	}

	return strings.Join(parts, " ")
}

// exampleString renders a scale's notes with Unicode accidentals, e.g. "C D♭ E♭ F G♭ A♭ B♭♭".
func exampleString(notes []scales.Note) string {
	parts := make([]string, len(notes))
	for i, n := range notes {
		parts[i] = unicodeNoteName(n.Name)
	}

	return strings.Join(parts, " ")
}
