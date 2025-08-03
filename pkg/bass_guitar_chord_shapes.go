package scales

import (
	"fmt"

	"github.com/andriikushch/scales/pkg/internal"
)

var (
	bassGuitarOpenCMaj       = []int{0, 2, 3, MutedNote}
	bassGuitarChordShapeCMaj = newChordShape("bassGuitar", []int{internal.IUnison, internal.IM3, internal.IP5}, bassGuitarOpenCMaj, 2)
)

var (
	bassGuitarOpenEMaj       = []int{1, 2, 2, 0}
	bassGuitarChordShapeEMaj = newChordShape("bassGuitar", []int{internal.IUnison, internal.IM3, internal.IP5}, bassGuitarOpenEMaj, 3)
)

var (
	bassGuitarOpenGMaj       = []int{0, 0, 2, 3}
	bassGuitarChordShapeGMaj = newChordShape("bassGuitar", []int{internal.IUnison, internal.IM3, internal.IP5}, bassGuitarOpenGMaj, 3)
)

var (
	bassGuitarOpenEMin       = []int{0, 2, 2, 0}
	bassGuitarChordShapeEMin = newChordShape("bassGuitar", []int{internal.IUnison, internal.Im3, internal.IP5}, bassGuitarOpenEMin, 3)
)

var (
	bassGuitarOpenGMin       = []int{0, 0, 1, 3}
	bassGuitarChordShapeGMin = newChordShape("bassGuitar", []int{internal.IUnison, internal.Im3, internal.IP5}, bassGuitarOpenGMin, 3)
)

var (
	bassGuitarOpenCMin       = []int{0, 1, 3, MutedNote}
	bassGuitarChordShapeCMin = newChordShape("bassGuitar", []int{internal.IUnison, internal.Im3, internal.IP5}, bassGuitarOpenCMin, 2)
)

var BassGuitarChordShapes = map[string][]ChordShape{
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.Im3, internal.IP5): {
		bassGuitarChordShapeEMin,
		bassGuitarChordShapeGMin,
		bassGuitarChordShapeCMin,
	},
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.IM3, internal.IP5): {
		bassGuitarChordShapeCMaj,
		bassGuitarChordShapeEMaj,
		bassGuitarChordShapeGMaj,
	},
}
