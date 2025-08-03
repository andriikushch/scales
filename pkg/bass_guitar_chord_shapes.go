package scales

import (
	"fmt"

	"github.com/andriikushch/scales/pkg/internal"
)

var (
	bassGuitarOpenC_Maj_0_Schema           = []int{0, 2, 3, MutedNote}
	bassGuitarChordShapeC_Maj_0_ChordShape = newChordShape("bassGuitar", []int{internal.IUnison, internal.IM3, internal.IP5}, bassGuitarOpenC_Maj_0_Schema, 2)
)

var (
	bassGuitarOpenE_Maj_1_Schema           = []int{1, 2, 2, 0}
	bassGuitarChordShapeE_Maj_1_ChordShape = newChordShape("bassGuitar", []int{internal.IUnison, internal.IM3, internal.IP5}, bassGuitarOpenE_Maj_1_Schema, 3)
)

var (
	bassGuitarOpenG_Maj_2_Schema           = []int{0, 0, 2, 3}
	bassGuitarChordShapeG_Maj_2_ChordShape = newChordShape("bassGuitar", []int{internal.IUnison, internal.IM3, internal.IP5}, bassGuitarOpenG_Maj_2_Schema, 3)
)

var (
	bassGuitarOpenE_Min_0_Schema           = []int{0, 2, 2, 0}
	bassGuitarChordShapeE_Min_0_ChordShape = newChordShape("bassGuitar", []int{internal.IUnison, internal.Im3, internal.IP5}, bassGuitarOpenE_Min_0_Schema, 3)
)

var (
	bassGuitarOpenG_Min_0_Schema           = []int{0, 0, 1, 3}
	bassGuitarChordShapeG_Min_1_ChordShape = newChordShape("bassGuitar", []int{internal.IUnison, internal.Im3, internal.IP5}, bassGuitarOpenG_Min_0_Schema, 3)
)

var (
	bassGuitarOpenC_Min_0_Schema           = []int{0, 1, 3, MutedNote}
	bassGuitarChordShapeC_Min_2_ChordShape = newChordShape("bassGuitar", []int{internal.IUnison, internal.Im3, internal.IP5}, bassGuitarOpenC_Min_0_Schema, 2)
)

var BassGuitarChordShapes = map[string][]ChordShape{
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.Im3, internal.IP5): {
		bassGuitarChordShapeE_Min_0_ChordShape,
		bassGuitarChordShapeG_Min_1_ChordShape,
		bassGuitarChordShapeC_Min_2_ChordShape,
	},
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.IM3, internal.IP5): {
		bassGuitarChordShapeC_Maj_0_ChordShape,
		bassGuitarChordShapeE_Maj_1_ChordShape,
		bassGuitarChordShapeG_Maj_2_ChordShape,
	},
}
