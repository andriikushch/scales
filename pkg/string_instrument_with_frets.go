package scales

import (
	"fmt"
	"github.com/andriikushch/scales/pkg/internal"
	"github.com/andriikushch/scales/pkg/internal/colors"
	"io"
	"slices"
)

func (g *stringInstrumentWithFrets) drawOrNot(n Note, scale []Note) (bool, Note, int) {
	index := slices.IndexFunc(scale, n.Equal)

	if index == -1 {
		return false, Note{}, -1
	}

	return true, scale[index], index
}

func (g *stringInstrumentWithFrets) Draw(notesToDraw []Note, w io.Writer) error {
	var structure []int
	for range 25 {
		structure = append(structure, internal.HalfStep)
	}

	for i, note := range g.tuning {
		scale, err := newScale(note.Name, structure, []string{})
		if err != nil {
			return err
		}

		if i == 0 {
			g.printFretMarkers(0, -1, scale, w)
		}

		for i, note := range scale.GetNotes() {
			draw, noteFromTheScale, colorIndex := g.drawOrNot(note, notesToDraw)

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

		if i == len(g.tuning)-1 {
			g.printFretMarkers(0, -1, scale, w)
		}
	}

	return nil
}

func (g *stringInstrumentWithFrets) DrawChord(cs internal.ChordShape, c Chord, w io.Writer) error {
	rootString := cs.RootNotePosition

	var structure []int
	for range 24 {
		structure = append(structure, internal.HalfStep)
	}

	allNotesOnTheString, err := newScale(g.tuning[rootString].Name, structure, []string{})
	if err != nil {
		return err
	}

	rootNotePositionOnTheString := slices.Index(allNotesOnTheString.GetNotes(), c.root)
	if rootNotePositionOnTheString == -1 {
		return fmt.Errorf("note was not found on the string")
	}

	leftFret := rootNotePositionOnTheString - cs.Schema[cs.RootNotePosition]
	rightFret := rootNotePositionOnTheString + slices.Max(cs.Schema)

	g.printFretMarkers(leftFret, rightFret, allNotesOnTheString, w)

	for str, note := range g.tuning {
		stringAlreadyUsed := false
		scale, err := newScale(note.Name, structure, []string{})
		if err != nil {
			return err
		}

		for fret, note := range scale.GetNotes() {
			if fret < leftFret || fret > rightFret {
				continue
			}

			draw, noteFromTheScale, colorIndex := g.drawOrNot(note, c.Notes())

			nName := ""
			color := colors.GetColor(colorIndex)
			if cs.Schema[str] == internal.MutedNote {
				nName = "-"
				color = colors.GetColor(-1)
			}

			if stringAlreadyUsed {
				color = colors.GetColor(-1)
			}

			draw = draw && (cs.Schema[str]+leftFret == fret) && !stringAlreadyUsed
			if draw {
				stringAlreadyUsed = true
			}
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

	g.printFretMarkers(leftFret, rightFret, allNotesOnTheString, w)

	return nil
}

func (g *stringInstrumentWithFrets) printFretMarkers(startingPosition int, endPosition int, scale *Scale, w io.Writer) {
	for i := startingPosition; i < len(scale.GetNotes()) && i <= endPosition; i++ {
		if i == 5 || i == 7 || i == 12 || i == 17 || i == 19 || i == 24 {
			_, _ = fmt.Fprintf(w, "%s%s_%02d__", colors.BGBlack, colors.White, i)
		} else {
			_, _ = fmt.Fprintf(w, "%s%s_%02d__", colors.BGYellow, colors.Black, i)
		}
	}
	_, _ = fmt.Fprint(w, colors.End+"\r\n")
}
