package scales

import (
	"fmt"

	"github.com/andriikushch/scales/pkg/internal"
)

// Open position chords
var (
	guitarOpenC_Maj_0_Schema           = []int{0, 1, 0, 2, 3, MutedNote}
	guitarChordShapeC_Maj_0_ChordShape = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, guitarOpenC_Maj_0_Schema, 4)
)

var (
	guitarOpenA_Maj_1_Schema           = []int{0, 2, 2, 2, 0, MutedNote}
	guitarChordShapeA_Maj_1_ChordShape = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, guitarOpenA_Maj_1_Schema, 4)
)

var (
	guitarOpenA_Min_0_Schema           = []int{0, 1, 2, 2, 0, MutedNote}
	guitarChordShapeA_Min_0_ChordShape = newChordShape("guitar", []int{internal.IUnison, internal.Im3, internal.IP5}, guitarOpenA_Min_0_Schema, 4)
)

var (
	guitarOpenD_Maj_2_Schema           = []int{2, 3, 2, 0, MutedNote, MutedNote}
	guitarChordShapeD_Maj_2_ChordShape = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, guitarOpenD_Maj_2_Schema, 3)
)

var (
	guitarOpenD_Min_1_Schema           = []int{1, 3, 2, 0, MutedNote, MutedNote}
	guitarChordShapeD_Min_1_ChordShape = newChordShape("guitar", []int{internal.IUnison, internal.Im3, internal.IP5}, guitarOpenD_Min_1_Schema, 3)
)

var (
	guitarOpenE_Maj_3_Schema           = []int{0, 0, 1, 2, 2, 0}
	guitarChordShapeE_Maj_3_ChordShape = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, guitarOpenE_Maj_3_Schema, 5)
)

var (
	guitarOpenE_Min_2_Schema           = []int{0, 0, 0, 2, 2, 0}
	guitarChordShapeE_Min_2_ChordShape = newChordShape("guitar", []int{internal.IUnison, internal.Im3, internal.IP5}, guitarOpenE_Min_2_Schema, 5)
)

var (
	guitarOpenG_Maj_4_Schema           = []int{3, 0, 0, 0, 2, 3}
	guitarChordShapeG_Maj_4_ChordShape = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, guitarOpenG_Maj_4_Schema, 5)
)

var GuitarChordShapes = map[string][]ChordShape{
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.Im3, internal.IP5): {
		guitarChordShapeA_Min_0_ChordShape,
		guitarChordShapeD_Min_1_ChordShape,
		guitarChordShapeE_Min_2_ChordShape,
	},
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.IM3, internal.IP5): {
		guitarChordShapeC_Maj_0_ChordShape,
		guitarChordShapeA_Maj_1_ChordShape,
		guitarChordShapeD_Maj_2_ChordShape,
		guitarChordShapeE_Maj_3_ChordShape,
		guitarChordShapeG_Maj_4_ChordShape,
	},
}
