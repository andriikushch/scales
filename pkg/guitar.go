package scales

import (
	"fmt"
	"io"
	"slices"

	"github.com/andriikushch/scales/pkg/internal"
	"github.com/andriikushch/scales/pkg/internal/colors"
)

type Guitar struct {
	tuning []Note
}

func (g *Guitar) drawOrNot(n Note, scale []Note) (bool, Note, int) {
	index := slices.IndexFunc(scale, n.Equal)

	if index == -1 {
		return false, Note{}, -1
	}

	return true, scale[index], index
}

func (g *Guitar) Draw(notesToDraw []Note, w io.Writer) error {
	var structure []int
	for range 25 {
		structure = append(structure, internal.HalfStep)
	}

	for i, note := range g.tuning {
		scale, err := newScale(note.Name, structure)
		if err != nil {
			return err
		}

		if i == 0 {
			g.printFretMarkers(scale, w)
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
		_, _ = fmt.Fprintln(w)

		if i == len(g.tuning)-1 {
			g.printFretMarkers(scale, w)
		}
	}

	return nil
}

func (g *Guitar) printFretMarkers(scale *Scale, w io.Writer) {
	for i := range scale.GetNotes() {
		_, _ = fmt.Fprintf(w, "%s%s_%02d__", colors.BGYellow, colors.Black, i)
	}
	_, _ = fmt.Fprintln(w, colors.End)
}

func NewGuitarWithStandardTuning() *Guitar {
	return &Guitar{
		tuning: []Note{
			NewNote(internal.E),
			NewNote(internal.B),
			NewNote(internal.G),
			NewNote(internal.D),
			NewNote(internal.A),
			NewNote(internal.E),
		},
	}
}

func NewBassGuitarWithStandardTuning() *Guitar {
	return &Guitar{
		tuning: []Note{
			NewNote(internal.G),
			NewNote(internal.D),
			NewNote(internal.A),
			NewNote(internal.E),
		},
	}
}
