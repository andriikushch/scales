package scales

import (
	"fmt"

	"github.com/andriikushch/scales/pkg/internal"
)

// Open position ukulele chords (G–C–E–A tuning)
var (
	openUkuleleC_Maj_0_Schema           = []int{3, 0, 0, 0}
	ukuleleChordShapeC_Maj_0_ChordShape = newChordShape("ukulele", []int{internal.IUnison, internal.IM3, internal.IP5}, openUkuleleC_Maj_0_Schema, 0)
)

var (
	openUkuleleA_Maj_1_Schema           = []int{0, 0, 1, 2}
	ukuleleChordShapeA_Maj_1_ChordShape = newChordShape("ukulele", []int{internal.IUnison, internal.IM3, internal.IP5}, openUkuleleA_Maj_1_Schema, 0)
)

var (
	openUkuleleA_Min_0_Schema           = []int{0, 0, 0, 2}
	ukuleleChordShapeA_Min_0_ChordShape = newChordShape("ukulele", []int{internal.IUnison, internal.Im3, internal.IP5}, openUkuleleA_Min_0_Schema, 3)
)

var (
	openUkuleleG_Maj_2_Schema           = []int{2, 3, 2, 0}
	ukuleleChordShapeG_Maj_2_ChordShape = newChordShape("ukulele", []int{internal.IUnison, internal.IM3, internal.IP5}, openUkuleleG_Maj_2_Schema, 3)
)

var (
	openUkuleleG_Min_1_Schema           = []int{1, 3, 2, 0}
	ukuleleChordShapeG_Min_1_ChordShape = newChordShape("ukulele", []int{internal.IUnison, internal.Im3, internal.IP5}, openUkuleleG_Min_1_Schema, 3)
)

var (
	openUkuleleD_Maj_3_Schema           = []int{0, 2, 2, 2}
	ukuleleChordShapeD_Maj_3_ChordShape = newChordShape("ukulele", []int{internal.IUnison, internal.IM3, internal.IP5}, openUkuleleD_Maj_3_Schema, 2)
)

var (
	openUkuleleD_Min_2_Schema           = []int{0, 1, 2, 2}
	ukuleleChordShapeD_Min_2_ChordShape = newChordShape("ukulele", []int{internal.IUnison, internal.Im3, internal.IP5}, openUkuleleD_Min_2_Schema, 2)
)

var (
	openUkuleleE_Min_3_Schema           = []int{2, 3, 4, 0}
	ukuleleChordShapeE_Min_3_ChordShape = newChordShape("ukulele", []int{internal.IUnison, internal.Im3, internal.IP5}, openUkuleleE_Min_3_Schema, 2)
)

var (
	openUkuleleF_Maj_4_Schema           = []int{0, 1, 0, 2}
	ukuleleChordShapeF_Maj_4_ChordShape = newChordShape("ukulele", []int{internal.IUnison, internal.IM3, internal.IP5}, openUkuleleF_Maj_4_Schema, 1)
)

var UkuleleChordShapes = map[string][]ChordShape{
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.IM3, internal.IP5): {
		ukuleleChordShapeC_Maj_0_ChordShape,
		ukuleleChordShapeA_Maj_1_ChordShape,
		ukuleleChordShapeG_Maj_2_ChordShape,
		ukuleleChordShapeD_Maj_3_ChordShape,
		ukuleleChordShapeF_Maj_4_ChordShape,
	},
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.Im3, internal.IP5): {
		ukuleleChordShapeA_Min_0_ChordShape,
		ukuleleChordShapeG_Min_1_ChordShape,
		ukuleleChordShapeD_Min_2_ChordShape,
		ukuleleChordShapeE_Min_3_ChordShape,
	},
}
