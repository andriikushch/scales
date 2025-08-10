package scales

import (
	"github.com/andriikushch/scales/pkg/internal"
)

var BassGuitarChordShapes = map[string][]ChordShape{}

func init() {
	BassGuitarChordShapes = initShapes(allBassChordShapes)
}

var (
	bassGuitar__0_Schema     = []int{0, 2, 3, MutedNote}
	bassGuitar__0_ChordShape = newChordShape("bassGuitar", []int{internal.IUnison, internal.IM3, internal.IP5}, bassGuitar__0_Schema, 2)
)

var (
	bassGuitar__1_Schema     = []int{1, 2, 2, 0}
	bassGuitar__1_ChordShape = newChordShape("bassGuitar", []int{internal.IUnison, internal.IM3, internal.IP5}, bassGuitar__1_Schema, 3)
)

var (
	bassGuitar__2_Schema     = []int{0, 0, 2, 3}
	bassGuitar__2_ChordShape = newChordShape("bassGuitar", []int{internal.IUnison, internal.IM3, internal.IP5}, bassGuitar__2_Schema, 3)
)

var (
	bassGuitar_m_0_Schema     = []int{0, 2, 2, 0}
	bassGuitar_m_0_ChordShape = newChordShape("bassGuitar", []int{internal.IUnison, internal.Im3, internal.IP5}, bassGuitar_m_0_Schema, 3)
)

var (
	bassGuitar_m_1_Schema     = []int{0, 0, 1, 3}
	bassGuitar_m_1_ChordShape = newChordShape("bassGuitar", []int{internal.IUnison, internal.Im3, internal.IP5}, bassGuitar_m_1_Schema, 3)
)

var (
	bassGuitar_m_2_Schema     = []int{0, 1, 3, MutedNote}
	bassGuitar_m_2_ChordShape = newChordShape("bassGuitar", []int{internal.IUnison, internal.Im3, internal.IP5}, bassGuitar_m_2_Schema, 2)
)

var (
	bassGuitar_maj7_sharp_5_0_Schema     = []int{1, 1, 3, 0}
	bassGuitar_maj7_sharp_5_0_ChordShape = newChordShape("bassGuitar", []int{0, 4, 8, 11}, bassGuitar_maj7_sharp_5_0_Schema, 3)
)

var (
	// todo: is there is a good shape?
	bassGuitar_aug6_0_Schema     = []int{}
	bassGuitar_aug6_0_ChordShape = newChordShape("bassGuitar", []int{0, 4, 8, 9}, bassGuitar_aug6_0_Schema, -1)
)

var (
	bassGuitar_m_add9_0_Schema     = []int{}
	bassGuitar_m_add9_0_ChordShape = newChordShape("bassGuitar", []int{0, 3, 7, 14}, bassGuitar_m_add9_0_Schema, -1)
)

var (
	bassGuitar_m7_sharp_5_0_Schema     = []int{}
	bassGuitar_m7_sharp_5_0_ChordShape = newChordShape("bassGuitar", []int{0, 3, 8, 10}, bassGuitar_m7_sharp_5_0_Schema, -1)
)

var (
	bassGuitar_7b5_0_Schema     = []int{} // TODO
	bassGuitar_7b5_0_ChordShape = newChordShape("bassGuitar", []int{0, 4, 6, 10}, bassGuitar_7b5_0_Schema, -1)
)

var (
	bassGuitar_maj7b5_0_Schema     = []int{} // TODO
	bassGuitar_maj7b5_0_ChordShape = newChordShape("bassGuitar", []int{0, 4, 6, 11}, bassGuitar_maj7b5_0_Schema, -1)
)

var (
	bassGuitar_m_maj7_0_Schema     = []int{} // TODO
	bassGuitar_m_maj7_0_ChordShape = newChordShape("bassGuitar", []int{0, 3, 7, 11}, bassGuitar_m_maj7_0_Schema, -1)
)

var (
	bassGuitar_sus_0_Schema     = []int{} // TODO
	bassGuitar_sus_0_ChordShape = newChordShape("bassGuitar", []int{0, 5, 7}, bassGuitar_sus_0_Schema, -1)
)

var (
	bassGuitar_dim7_0_Schema     = []int{} // TODO
	bassGuitar_dim7_0_ChordShape = newChordShape("bassGuitar", []int{0, 3, 6, 9}, bassGuitar_dim7_0_Schema, -1)
)

var (
	bassGuitar_maj7_0_Schema     = []int{} // TODO
	bassGuitar_maj7_0_ChordShape = newChordShape("bassGuitar", []int{0, 4, 7, 11}, bassGuitar_maj7_0_Schema, -1)
)

var (
	bassGuitar_aug7_0_Schema     = []int{} // TODO
	bassGuitar_aug7_0_ChordShape = newChordShape("bassGuitar", []int{0, 4, 8, 10}, bassGuitar_aug7_0_Schema, -1)
)

var (
	bassGuitar_aug_0_Schema     = []int{0, 1, 2, MutedNote}
	bassGuitar_aug_0_ChordShape = newChordShape("bassGuitar", []int{0, 4, 8}, bassGuitar_aug_0_Schema, 2)
)

var (
	bassGuitar_m7_0_Schema     = []int{} // TODO
	bassGuitar_m7_0_ChordShape = newChordShape("bassGuitar", []int{0, 3, 7, 10}, bassGuitar_m7_0_Schema, -1)
)

var (
	bassGuitar_7_0_Schema     = []int{} // TODO
	bassGuitar_7_0_ChordShape = newChordShape("bassGuitar", []int{0, 4, 7, 10}, bassGuitar_7_0_Schema, -1)
)

var (
	bassGuitar_7sus_0_Schema     = []int{} // TODO
	bassGuitar_7sus_0_ChordShape = newChordShape("bassGuitar", []int{0, 5, 7, 10}, bassGuitar_7sus_0_Schema, -1)
)

var (
	bassGuitar_sus2_0_Schema     = []int{} // TODO
	bassGuitar_sus2_0_ChordShape = newChordShape("bassGuitar", []int{0, 2, 7}, bassGuitar_sus2_0_Schema, -1)
)

var (
	bassGuitar_m7b5_0_Schema     = []int{} // TODO
	bassGuitar_m7b5_0_ChordShape = newChordShape("bassGuitar", []int{0, 3, 6, 10}, bassGuitar_m7b5_0_Schema, -1)
)

var (
	bassGuitar_dim_0_Schema     = []int{} // TODO
	bassGuitar_dim_0_ChordShape = newChordShape("bassGuitar", []int{0, 3, 6}, bassGuitar_dim_0_Schema, -1)
)

var (
	bassGuitar_6_0_Schema     = []int{} // TODO
	bassGuitar_6_0_ChordShape = newChordShape("bassGuitar", []int{0, 4, 7, 9}, bassGuitar_6_0_Schema, -1)
)

var (
	bassGuitar_dim6_0_Schema     = []int{} // TODO
	bassGuitar_dim6_0_ChordShape = newChordShape("bassGuitar", []int{0, 3, 6, 8}, bassGuitar_dim6_0_Schema, -1)
)

var (
	bassGuitar_add9_0_Schema     = []int{} // TODO
	bassGuitar_add9_0_ChordShape = newChordShape("bassGuitar", []int{0, 4, 7, 14}, bassGuitar_add9_0_Schema, -1)
)

var (
	bassGuitar_m6_0_Schema     = []int{} // TODO
	bassGuitar_m6_0_ChordShape = newChordShape("bassGuitar", []int{0, 3, 7, 9}, bassGuitar_m6_0_Schema, -1)
)
var allBassChordShapes []ChordShape = []ChordShape{bassGuitar_6_0_ChordShape, bassGuitar_7_0_ChordShape, bassGuitar_7b5_0_ChordShape, bassGuitar_7sus_0_ChordShape, bassGuitar__0_ChordShape, bassGuitar__1_ChordShape, bassGuitar__2_ChordShape, bassGuitar_add9_0_ChordShape, bassGuitar_aug6_0_ChordShape, bassGuitar_aug7_0_ChordShape, bassGuitar_aug_0_ChordShape, bassGuitar_dim6_0_ChordShape, bassGuitar_dim7_0_ChordShape, bassGuitar_dim_0_ChordShape, bassGuitar_m6_0_ChordShape, bassGuitar_m7_0_ChordShape, bassGuitar_m7_sharp_5_0_ChordShape, bassGuitar_m7b5_0_ChordShape, bassGuitar_m_0_ChordShape, bassGuitar_m_1_ChordShape, bassGuitar_m_2_ChordShape, bassGuitar_m_add9_0_ChordShape, bassGuitar_m_maj7_0_ChordShape, bassGuitar_maj7_0_ChordShape, bassGuitar_maj7_sharp_5_0_ChordShape, bassGuitar_maj7b5_0_ChordShape, bassGuitar_sus2_0_ChordShape, bassGuitar_sus_0_ChordShape}
