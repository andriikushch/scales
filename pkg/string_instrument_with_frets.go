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

func (inst *stringInstrumentWithFrets) Draw(notesToDraw []Note, w io.Writer) error {
	var structure []int
	for range 25 {
		structure = append(structure, internal.HalfStep)
	}

	for i, note := range inst.tuning {
		scale, err := newScale(note.Name, structure, []string{})
		if err != nil {
			return err
		}

		if i == 0 {
			inst.printFretMarkers(0, 24, scale, w)
		}

		for i, note := range scale.GetNotes() {
			draw, noteFromTheScale, colorIndex := inst.drawOrNot(note, notesToDraw)

			if i == 0 {
				if draw {
					_, _ = fmt.Fprintf(w, "%s%-3s%s||", colors.GetColor(colorIndex), noteFromTheScale.Name, colors.End)
				} else {
					_, _ = fmt.Fprintf(w, "%s%-3s%s||", colors.GetColor(colorIndex), "", colors.End)
				}
				continue
			}
			if draw {
				_, _ = fmt.Fprintf(w, "%s %-3s%s|", colors.GetColor(colorIndex), noteFromTheScale.Name, colors.End)
			} else {
				_, _ = fmt.Fprintf(w, "%s %-3s%s|", colors.GetColor(colorIndex), "", colors.End)
			}

		}
		_, _ = fmt.Fprint(w, "\r\n")

		if i == len(inst.tuning)-1 {
			inst.printFretMarkers(0, 24, scale, w)
		}
	}

	return nil
}

func (inst *stringInstrumentWithFrets) drawChord(cs chordShape, c Chord, w io.Writer) error {
	rootString := cs.RootNotePosition

	var structure []int
	for range 24 {
		structure = append(structure, internal.HalfStep)
	}

	allNotesOnTheString, err := newScale(inst.tuning[rootString].Name, structure, []string{})
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
		scale, err := newScale(note.Name, structure, []string{})
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
