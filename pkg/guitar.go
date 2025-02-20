package scales

import (
	"fmt"
	"github.com/andriikushch/scales/pkg/internal"
	"github.com/andriikushch/scales/pkg/internal/colors"
	"slices"
)

type Guitar struct {
	tuning []Note
}

func (g *Guitar) drawOrNot(n Note, scale []Note) (bool, Note, int) {
	index := slices.IndexFunc(scale, func(note Note) bool {
		return note.Equal(n)
	})

	if index == -1 {
		return false, Note{}, -1
	}

	return true, scale[index], index
}

func (g *Guitar) Draw(notesToDraw []Note) error {
	var structure []int
	for _ = range 25 {
		structure = append(structure, internal.HalfStep)
	}

	for i, note := range g.tuning {
		scale, err := newScale(note.Name, structure)
		if err != nil {
			return err
		}

		if i == 0 {
			for i := range scale.GetNotes() {
				if i == 0 {
					fmt.Printf("%s%s_____", colors.Black, colors.BGYellow)
					continue
				}
				fmt.Printf("_____")
			}
			fmt.Println(colors.End)
		}

		for i, note := range scale.GetNotes() {
			if i == 0 {
				if draw, noteFromTheScale, colorIndex := g.drawOrNot(note, notesToDraw); draw {
					fmt.Printf("%s%-3s%s||", colors.GetColor(colorIndex), noteFromTheScale.Name, colors.End)
				} else {
					fmt.Printf("%s%-3s%s||", colors.GetColor(colorIndex), "", colors.End)
				}

				continue
			}
			if draw, noteFromTheScale, colorIndex := g.drawOrNot(note, notesToDraw); draw {
				fmt.Printf("%s %-3s%s|", colors.GetColor(colorIndex), noteFromTheScale.Name, colors.End)
			} else {
				fmt.Printf("%s %-3s%s|", colors.GetColor(colorIndex), "", colors.End)
			}

		}
		fmt.Println()

		if i == len(g.tuning)-1 {
			for i := range scale.GetNotes() {

				fmt.Printf("%s%s_%02d__", colors.BGYellow, colors.Black, i)
			}
			fmt.Println(colors.End)
		}
	}

	return nil
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
