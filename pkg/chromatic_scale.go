package scales

import (
	"slices"
	"strings"

	"github.com/andriikushch/scales/pkg/internal"
)

type chromaticScale [][]Note

func (s chromaticScale) next(previousNote Note, steps int) []Note {
	for steps < 0 {
		steps += len(s)
	}

	previousNoteIndex := -1

	for i, notes := range s {
		if slices.Index(notes, previousNote) >= 0 {
			previousNoteIndex = i

			break
		}
	}

	return s[(previousNoteIndex+steps)%len(s)]
}

func (s chromaticScale) position(n Note) int {
	pos := -1

	for i, notes := range s {
		for _, note := range notes {
			if note.Name == n.Name {
				pos = i
				break
			}
		}
	}

	return pos
}

func (s chromaticScale) normalize(noteToNormalize Note) Note {
	stepsForward := strings.Count(noteToNormalize.Name, "##") * internal.Step
	stepsBackward := -1 * strings.Count(noteToNormalize.Name, "bb") * internal.Step

	noteToNormalize.Name = strings.ReplaceAll(noteToNormalize.Name, "##", "")
	noteToNormalize.Name = strings.ReplaceAll(noteToNormalize.Name, "bb", "")
	return defaultChromaticScale.next(noteToNormalize, stepsForward+stepsBackward)[0]
}

var defaultChromaticScale = chromaticScale{
	{NewNote(internal.BSharp), NewNote(internal.C)},
	{NewNote(internal.CSharp), NewNote(internal.DFlat)},
	{NewNote(internal.CDoubleSharp), NewNote(internal.D)},
	{NewNote(internal.DSharp), NewNote(internal.EFlat)},
	{NewNote(internal.E)},
	{NewNote(internal.ESharp), NewNote(internal.F)},
	{NewNote(internal.FSharp), NewNote(internal.GFlat)},
	{NewNote(internal.FDoubleSharp), NewNote(internal.G)},
	{NewNote(internal.GSharp), NewNote(internal.AFlat)},
	{NewNote(internal.A)},
	{NewNote(internal.ASharp), NewNote(internal.BFlat)},
	{NewNote(internal.B)},
}
