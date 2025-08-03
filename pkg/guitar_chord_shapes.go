package scales

import (
	"fmt"

	"github.com/andriikushch/scales/pkg/internal"
)

// Open position chords
var (
	guitarOpenCMaj       = []int{0, 1, 0, 2, 3, MutedNote}
	guitarChordShapeCMaj = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, guitarOpenCMaj, 4)
)

var (
	guitarOpenAMaj       = []int{0, 2, 2, 2, 0, MutedNote}
	guitarChordShapeAMaj = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, guitarOpenAMaj, 4)
)

var (
	guitarOpenAMin       = []int{0, 1, 2, 2, 0, MutedNote}
	guitarChordShapeAMin = newChordShape("guitar", []int{internal.IUnison, internal.Im3, internal.IP5}, guitarOpenAMin, 4)
)

var (
	guitarOpenDMaj       = []int{2, 3, 2, 0, MutedNote, MutedNote}
	guitarChordShapeDMaj = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, guitarOpenDMaj, 3)
)

var (
	guitarOpenDMin       = []int{1, 3, 2, 0, MutedNote, MutedNote}
	guitarChordShapeDMin = newChordShape("guitar", []int{internal.IUnison, internal.Im3, internal.IP5}, guitarOpenDMin, 3)
)

var (
	guitarOpenEMaj       = []int{0, 0, 1, 2, 2, 0}
	guitarChordShapeEMaj = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, guitarOpenEMaj, 5)
)

var (
	guitarOpenEMin       = []int{0, 0, 0, 2, 2, 0}
	guitarChordShapeEMin = newChordShape("guitar", []int{internal.IUnison, internal.Im3, internal.IP5}, guitarOpenEMin, 5)
)

var (
	guitarOpenGMaj       = []int{3, 0, 0, 0, 2, 3}
	guitarChordShapeGMaj = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, guitarOpenGMaj, 5)
)

var GuitarChordShapes = map[string][]ChordShape{
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.Im3, internal.IP5): {
		guitarChordShapeAMin,
		guitarChordShapeDMin,
		guitarChordShapeEMin,
	},
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.IM3, internal.IP5): {
		guitarChordShapeCMaj,
		guitarChordShapeAMaj,
		guitarChordShapeDMaj,
		guitarChordShapeEMaj,
		guitarChordShapeGMaj,
	},
}
