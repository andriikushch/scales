package scales

import (
	"slices"

	"github.com/andriikushch/scales/pkg/internal/chromatic"
)

type chromaticScale [][]Note

func (s chromaticScale) next(previousNote Note, steps int) []Note {
	previousNoteIndex := -1

	for i, notes := range s {
		if slices.Index(notes, previousNote) >= 0 {
			previousNoteIndex = i

			break
		}
	}

	return s[(previousNoteIndex+steps)%len(s)]
}

var defaultChromaticScale = chromaticScale{
	{NewNote(chromatic.BSharp), NewNote(chromatic.C)},
	{NewNote(chromatic.CSharp), NewNote(chromatic.DFlat)},
	{NewNote(chromatic.CDoubleSharp), NewNote(chromatic.D)},
	{NewNote(chromatic.DSharp), NewNote(chromatic.EFlat)},
	{NewNote(chromatic.E)},
	{NewNote(chromatic.ESharp), NewNote(chromatic.F)},
	{NewNote(chromatic.FSharp), NewNote(chromatic.GFlat)},
	{NewNote(chromatic.FDoubleSharp), NewNote(chromatic.G)},
	{NewNote(chromatic.GSharp), NewNote(chromatic.AFlat)},
	{NewNote(chromatic.A)},
	{NewNote(chromatic.ASharp), NewNote(chromatic.BFlat)},
	{NewNote(chromatic.B)},
}
