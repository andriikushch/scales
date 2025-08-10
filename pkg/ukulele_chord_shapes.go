package scales

import (
	"github.com/andriikushch/scales/pkg/internal"
)

var UkuleleChordShapes = map[string][]ChordShape{}

func init() {
	UkuleleChordShapes = initShapes(allUkuleleChordShapes)
}

// Open position ukulele chords (G–C–E–A tuning)
var (
	ukulele__0_Schema     = []int{3, 0, 0, 0}
	ukulele__0_ChordShape = newChordShape("ukulele", []int{internal.IUnison, internal.IM3, internal.IP5}, ukulele__0_Schema, 0)
)

var (
	ukulele__1_Schema     = []int{0, 0, 1, 2}
	ukulele__1_ChordShape = newChordShape("ukulele", []int{internal.IUnison, internal.IM3, internal.IP5}, ukulele__1_Schema, 0)
)

var (
	ukulele_m_0_Schema     = []int{0, 0, 0, 2}
	ukulele_m_0_ChordShape = newChordShape("ukulele", []int{internal.IUnison, internal.Im3, internal.IP5}, ukulele_m_0_Schema, 3)
)

var (
	ukulele__2_Schema     = []int{2, 3, 2, 0}
	ukulele__2_ChordShape = newChordShape("ukulele", []int{internal.IUnison, internal.IM3, internal.IP5}, ukulele__2_Schema, 3)
)

var (
	ukulele_m_1_Schema     = []int{1, 3, 2, 0}
	ukulele_m_1_ChordShape = newChordShape("ukulele", []int{internal.IUnison, internal.Im3, internal.IP5}, ukulele_m_1_Schema, 3)
)

var (
	ukulele__3_Schema     = []int{0, 2, 2, 2}
	ukulele__3_ChordShape = newChordShape("ukulele", []int{internal.IUnison, internal.IM3, internal.IP5}, ukulele__3_Schema, 2)
)

var (
	ukulele_m_2_Schema     = []int{0, 1, 2, 2}
	ukulele_m_2_ChordShape = newChordShape("ukulele", []int{internal.IUnison, internal.Im3, internal.IP5}, ukulele_m_2_Schema, 2)
)

var (
	ukulele_m_3_Schema     = []int{2, 3, 4, 0}
	ukulele_m_3_ChordShape = newChordShape("ukulele", []int{internal.IUnison, internal.Im3, internal.IP5}, ukulele_m_3_Schema, 2)
)

var (
	ukulele__4_Schema     = []int{0, 1, 0, 2}
	ukulele__4_ChordShape = newChordShape("ukulele", []int{internal.IUnison, internal.IM3, internal.IP5}, ukulele__4_Schema, 1)
)

var allUkuleleChordShapes []ChordShape = []ChordShape{
	ukulele_6_0_ChordShape, ukulele_7_0_ChordShape, ukulele_7b5_0_ChordShape, ukulele_7sus_0_ChordShape, ukulele__0_ChordShape, ukulele__1_ChordShape, ukulele__2_ChordShape, ukulele__3_ChordShape, ukulele__4_ChordShape, ukulele_add9_0_ChordShape, ukulele_aug6_0_ChordShape, ukulele_aug7_0_ChordShape, ukulele_aug_0_ChordShape, ukulele_dim6_0_ChordShape, ukulele_dim7_0_ChordShape, ukulele_dim_0_ChordShape, ukulele_m6_0_ChordShape, ukulele_m7_0_ChordShape, ukulele_m7_sharp_5_0_ChordShape, ukulele_m7b5_0_ChordShape, ukulele_m_0_ChordShape, ukulele_m_1_ChordShape, ukulele_m_2_ChordShape, ukulele_m_3_ChordShape, ukulele_m_add9_0_ChordShape, ukulele_m_maj7_0_ChordShape, ukulele_maj7_0_ChordShape, ukulele_maj7_sharp_5_0_ChordShape, ukulele_maj7b5_0_ChordShape,

	// TODO
	ukulele_sus2_0_ChordShape, ukulele_sus_0_ChordShape,
}

// TODO

// TODO

// TODO

// TODO

var (
	ukulele_maj7_sharp_5_0_Schema     = []int{}
	ukulele_maj7_sharp_5_0_ChordShape = newChordShape("ukulele", []int{0, 4, 8, 11}, ukulele_maj7_sharp_5_0_Schema, -1)
)

var (
	ukulele_m7_sharp_5_0_Schema     = []int{}
	ukulele_m7_sharp_5_0_ChordShape = newChordShape("ukulele", []int{0, 3, 8, 10}, ukulele_m7_sharp_5_0_Schema, -1)
)

var (
	ukulele_aug6_0_Schema     = []int{}
	ukulele_aug6_0_ChordShape = newChordShape("ukulele", []int{0, 4, 8, 9}, ukulele_aug6_0_Schema, -1)
)

var (
	ukulele_m_add9_0_Schema     = []int{}
	ukulele_m_add9_0_ChordShape = newChordShape("ukulele", []int{0, 3, 7, 14}, ukulele_m_add9_0_Schema, -1)
)

var (
	ukulele_7b5_0_Schema     = []int{}
	ukulele_7b5_0_ChordShape = newChordShape("ukulele", []int{0, 4, 6, 10}, ukulele_7b5_0_Schema, -1)
)

var (
	ukulele_maj7b5_0_Schema     = []int{} // TODO
	ukulele_maj7b5_0_ChordShape = newChordShape("ukulele", []int{0, 4, 6, 11}, ukulele_maj7b5_0_Schema, -1)
)

var (
	ukulele_m_maj7_0_Schema     = []int{} // TODO
	ukulele_m_maj7_0_ChordShape = newChordShape("ukulele", []int{0, 3, 7, 11}, ukulele_m_maj7_0_Schema, -1)
)

var (
	ukulele_sus_0_Schema     = []int{} // TODO
	ukulele_sus_0_ChordShape = newChordShape("ukulele", []int{0, 5, 7}, ukulele_sus_0_Schema, -1)
)

var (
	ukulele_dim7_0_Schema     = []int{} // TODO
	ukulele_dim7_0_ChordShape = newChordShape("ukulele", []int{0, 3, 6, 9}, ukulele_dim7_0_Schema, -1)
)

var (
	ukulele_maj7_0_Schema     = []int{} // TODO
	ukulele_maj7_0_ChordShape = newChordShape("ukulele", []int{0, 4, 7, 11}, ukulele_maj7_0_Schema, -1)
)

var (
	ukulele_aug7_0_Schema     = []int{} // TODO
	ukulele_aug7_0_ChordShape = newChordShape("ukulele", []int{0, 4, 8, 10}, ukulele_aug7_0_Schema, -1)
)

var (
	ukulele_aug_0_Schema     = []int{} // TODO
	ukulele_aug_0_ChordShape = newChordShape("ukulele", []int{0, 4, 8}, ukulele_aug_0_Schema, -1)
)

var (
	ukulele_m7_0_Schema     = []int{} // TODO
	ukulele_m7_0_ChordShape = newChordShape("ukulele", []int{0, 3, 7, 10}, ukulele_m7_0_Schema, -1)
)

var (
	ukulele_7_0_Schema     = []int{} // TODO
	ukulele_7_0_ChordShape = newChordShape("ukulele", []int{0, 4, 7, 10}, ukulele_7_0_Schema, -1)
)

var (
	ukulele_7sus_0_Schema     = []int{} // TODO
	ukulele_7sus_0_ChordShape = newChordShape("ukulele", []int{0, 5, 7, 10}, ukulele_7sus_0_Schema, -1)
)

var (
	ukulele_sus2_0_Schema     = []int{} // TODO
	ukulele_sus2_0_ChordShape = newChordShape("ukulele", []int{0, 2, 7}, ukulele_sus2_0_Schema, -1)
)

var (
	ukulele_m7b5_0_Schema     = []int{} // TODO
	ukulele_m7b5_0_ChordShape = newChordShape("ukulele", []int{0, 3, 6, 10}, ukulele_m7b5_0_Schema, -1)
)

var (
	ukulele_dim_0_Schema     = []int{} // TODO
	ukulele_dim_0_ChordShape = newChordShape("ukulele", []int{0, 3, 6}, ukulele_dim_0_Schema, -1)
)

var (
	ukulele_6_0_Schema     = []int{} // TODO
	ukulele_6_0_ChordShape = newChordShape("ukulele", []int{0, 4, 7, 9}, ukulele_6_0_Schema, -1)
)

var (
	ukulele_dim6_0_Schema     = []int{} // TODO
	ukulele_dim6_0_ChordShape = newChordShape("ukulele", []int{0, 3, 6, 8}, ukulele_dim6_0_Schema, -1)
)

var (
	ukulele_add9_0_Schema     = []int{} // TODO
	ukulele_add9_0_ChordShape = newChordShape("ukulele", []int{0, 4, 7, 14}, ukulele_add9_0_Schema, -1)
)

var (
	ukulele_m6_0_Schema     = []int{} // TODO
	ukulele_m6_0_ChordShape = newChordShape("ukulele", []int{0, 3, 7, 9}, ukulele_m6_0_Schema, -1)
)
