package scales

import (
	"fmt"
	"io"
	"slices"

	"github.com/andriikushch/scales/pkg/internal"
	"github.com/andriikushch/scales/pkg/internal/colors"
)

type stringInstrumentWithFrets struct {
	tuning []Note
}

func (inst *stringInstrumentWithFrets) drawOrNot(n Note, scale []Note) (bool, Note, int) {
	index := slices.IndexFunc(scale, n.Equal)

	if index == -1 {
		return false, Note{}, -1
	}

	return true, scale[index], index
}

// FretCell describes a single string/fret position: the note that sounds
// there, and whether it is one of the notes being highlighted.
type FretCell struct {
	Note Note
	// Match is true if this cell is one of the notes passed to Positions.
	Match bool
	// MatchIndex is the index of Note within the notesToDraw slice passed to
	// Positions (useful for consistent coloring/root detection), or -1 if
	// Match is false.
	MatchIndex int
}

// Positions returns, for each of the instrument's strings, the note and
// match state at frets 0-24 for notesToDraw.
func (inst *stringInstrumentWithFrets) Positions(notesToDraw []Note) ([][]FretCell, error) {
	var structure []int
	for range 25 {
		structure = append(structure, internal.HalfStep)
	}

	grid := make([][]FretCell, len(inst.tuning))

	for str, note := range inst.tuning {
		scale, err := newScale(note.Name, structure)
		if err != nil {
			return nil, err
		}

		notes := scale.GetNotes()
		row := make([]FretCell, len(notes))
		for fret, note := range notes {
			draw, noteFromTheScale, colorIndex := inst.drawOrNot(note, notesToDraw)
			if draw {
				row[fret] = FretCell{Note: noteFromTheScale, Match: true, MatchIndex: colorIndex}
			} else {
				row[fret] = FretCell{Note: note, Match: false, MatchIndex: -1}
			}
		}
		grid[str] = row
	}

	return grid, nil
}

func (inst *stringInstrumentWithFrets) Draw(notesToDraw []Note, w io.Writer) error {
	grid, err := inst.Positions(notesToDraw)
	if err != nil {
		return err
	}

	var structure []int
	for range 25 {
		structure = append(structure, internal.HalfStep)
	}

	// printFretMarkers only cares about how many frets there are (always 25
	// here), not which string's scale is passed in.
	markerScale, err := newScale(inst.tuning[0].Name, structure)
	if err != nil {
		return err
	}

	for str, row := range grid {
		if str == 0 {
			inst.printFretMarkers(0, 24, markerScale, w)
		}

		for fret, cell := range row {
			colorIndex := -1
			name := ""
			if cell.Match {
				colorIndex = cell.MatchIndex
				name = cell.Note.Name
			}

			if fret == 0 {
				_, _ = fmt.Fprintf(w, "%s%-3s%s||", colors.GetColor(colorIndex), name, colors.End)
				continue
			}
			_, _ = fmt.Fprintf(w, "%s %-3s%s|", colors.GetColor(colorIndex), name, colors.End)
		}
		_, _ = fmt.Fprint(w, "\r\n")

		if str == len(grid)-1 {
			inst.printFretMarkers(0, 24, markerScale, w)
		}
	}

	return nil
}

func (inst *stringInstrumentWithFrets) drawChord(cs ChordShape, c Chord, w io.Writer) error {
	rootString := cs.RootNotePosition

	var structure []int
	for range 24 {
		structure = append(structure, internal.HalfStep)
	}

	allNotesOnTheString, err := newScale(inst.tuning[rootString].Name, structure)
	if err != nil {
		return err
	}

	rootNotePositionOnTheString := slices.Index(allNotesOnTheString.GetNotes(), c.root)
	if rootNotePositionOnTheString == -1 {
		return fmt.Errorf("note was not found on the string")
	}

	leftFret := rootNotePositionOnTheString - cs.Schema[cs.RootNotePosition]
	rightFret := leftFret + slices.Max(cs.Schema)

	if leftFret < 0 {
		rootNotePositionOnTheString += 12
		leftFret += 12
		rightFret += 12
	}

	inst.printFretMarkers(leftFret, rightFret, allNotesOnTheString, w)

	for str, note := range inst.tuning {
		scale, err := newScale(note.Name, structure)
		if err != nil {
			return err
		}

		for fret, note := range scale.GetNotes() {
			if fret < leftFret || fret > rightFret {
				continue
			}

			draw, noteFromTheScale, colorIndex := inst.drawOrNot(note, c.Notes())

			nName := ""
			color := colors.GetColor(colorIndex)
			if cs.Schema[str] == MutedNote {
				nName = "-"
				color = colors.GetColor(-1)
			}

			if cs.Schema[str]+leftFret != fret {
				color = colors.GetColor(-1)
			}

			draw = draw && (cs.Schema[str]+leftFret == fret)
			if fret == 0 {
				if draw {
					_, _ = fmt.Fprintf(w, "%s%-3s%s||", color, noteFromTheScale.Name, colors.End)
				} else {
					_, _ = fmt.Fprintf(w, "%s%-3s%s||", color, nName, colors.End)
				}
				continue
			}
			if draw {
				_, _ = fmt.Fprintf(w, "%s %-3s%s|", color, noteFromTheScale.Name, colors.End)
			} else {
				_, _ = fmt.Fprintf(w, "%s %-3s%s|", color, nName, colors.End)
			}

		}

		_, _ = fmt.Fprint(w, "\r\n")
	}

	inst.printFretMarkers(leftFret, rightFret, allNotesOnTheString, w)

	return nil
}

func (inst *stringInstrumentWithFrets) printFretMarkers(startingPosition int, endPosition int, scale *Scale, w io.Writer) {
	for i := startingPosition; i < len(scale.GetNotes()) && i <= endPosition; i++ {
		if i == 5 || i == 7 || i == 12 || i == 17 || i == 19 || i == 24 {
			_, _ = fmt.Fprintf(w, "%s%s_%02d__", colors.BGBlack, colors.White, i)
		} else {
			_, _ = fmt.Fprintf(w, "%s%s_%02d__", colors.BGYellow, colors.Black, i)
		}
	}
	_, _ = fmt.Fprint(w, colors.End+"\r\n")
}
