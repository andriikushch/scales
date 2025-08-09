package scales

import (
	"github.com/andriikushch/scales/pkg/internal"
)

var GuitarChordShapes map[string][]ChordShape

func init() {
	GuitarChordShapes = initShapes(allGuitarChordShapes)
}

// Open position chords
var (
	guitar__0_Schema     = []int{0, 1, 0, 2, 3, MutedNote}
	guitar__0_ChordShape = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, guitar__0_Schema, 4)
)

var (
	guitar__1_Schema     = []int{0, 2, 2, 2, 0, MutedNote}
	guitar__1_ChordShape = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, guitar__1_Schema, 4)
)

var (
	guitar_m_0_Schema     = []int{0, 1, 2, 2, 0, MutedNote}
	guitar_m_0_ChordShape = newChordShape("guitar", []int{internal.IUnison, internal.Im3, internal.IP5}, guitar_m_0_Schema, 4)
)

var (
	guitar__2_Schema     = []int{2, 3, 2, 0, MutedNote, MutedNote}
	guitar__2_ChordShape = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, guitar__2_Schema, 3)
)

var (
	guitar_m_1_Schema     = []int{1, 3, 2, 0, MutedNote, MutedNote}
	guitar_m_1_ChordShape = newChordShape("guitar", []int{internal.IUnison, internal.Im3, internal.IP5}, guitar_m_1_Schema, 3)
)

var (
	guitar__3_Schema     = []int{0, 0, 1, 2, 2, 0}
	guitar__3_ChordShape = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, guitar__3_Schema, 5)
)

var (
	guitar_m_2_Schema     = []int{0, 0, 0, 2, 2, 0}
	guitar_m_2_ChordShape = newChordShape("guitar", []int{internal.IUnison, internal.Im3, internal.IP5}, guitar_m_2_Schema, 5)
)

var (
	guitar__4_Schema     = []int{3, 0, 0, 0, 2, 3}
	guitar__4_ChordShape = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, guitar__4_Schema, 5)
)

var (
	guitar_m_maj7_0_Schema     = []int{MutedNote, 0, 0, 1, MutedNote, 0}
	guitar_m_maj7_0_ChordShape = newChordShape("guitar", []int{0, 3, 7, 11}, guitar_m_maj7_0_Schema, 5)
)

var (
	guitar_m_maj7_1_Schema     = []int{0, 1, 1, MutedNote, 0, MutedNote}
	guitar_m_maj7_1_ChordShape = newChordShape("guitar", []int{0, 3, 7, 11}, guitar_m_maj7_1_Schema, 4)
)

var (
	guitar_m_maj7_2_Schema     = []int{MutedNote, 1, 1, 2, 0, MutedNote}
	guitar_m_maj7_2_ChordShape = newChordShape("guitar", []int{0, 3, 7, 11}, guitar_m_maj7_2_Schema, 4)
)

var (
	guitar_aug6_0_Schema     = []int{} // TODO
	guitar_aug6_0_ChordShape = newChordShape("guitar", []int{0, 4, 8, 9}, guitar_aug6_0_Schema, -1)
)

var (
	guitar_sus_0_Schema     = []int{} // TODO
	guitar_sus_0_ChordShape = newChordShape("guitar", []int{0, 5, 7}, guitar_sus_0_Schema, -1)
)

var (
	guitar_dim6_0_Schema     = []int{} // TODO
	guitar_dim6_0_ChordShape = newChordShape("guitar", []int{0, 3, 6, 8}, guitar_dim6_0_Schema, -1)
)

var (
	guitar_dim7_0_Schema     = []int{MutedNote, 0, 1, 0, MutedNote, 1, MutedNote}
	guitar_dim7_0_ChordShape = newChordShape("guitar", []int{0, 3, 6, 9}, guitar_dim7_0_Schema, 1)
)

var (
	guitar_dim7_1_Schema     = []int{MutedNote, 0, 1, 0, MutedNote, 1, MutedNote}
	guitar_dim7_1_ChordShape = newChordShape("guitar", []int{0, 3, 6, 9}, guitar_dim7_1_Schema, 2)
)

var (
	guitar_dim7_2_Schema     = []int{MutedNote, 0, 1, 0, MutedNote, 1, MutedNote}
	guitar_dim7_2_ChordShape = newChordShape("guitar", []int{0, 3, 6, 9}, guitar_dim7_2_Schema, 3)
)

var (
	guitar_dim7_3_Schema     = []int{MutedNote, 0, 1, 0, MutedNote, 1, MutedNote}
	guitar_dim7_3_ChordShape = newChordShape("guitar", []int{0, 3, 6, 9}, guitar_dim7_3_Schema, 5)
)

var (
	guitar_dim7_4_Schema     = []int{0, 2, 0, MutedNote, 1, MutedNote}
	guitar_dim7_4_ChordShape = newChordShape("guitar", []int{0, 3, 6, 9}, guitar_dim7_4_Schema, 0)
)

var (
	guitar_dim7_5_Schema     = []int{0, 2, 0, MutedNote, 1, MutedNote}
	guitar_dim7_5_ChordShape = newChordShape("guitar", []int{0, 3, 6, 9}, guitar_dim7_5_Schema, 1)
)

var (
	guitar_dim7_6_Schema     = []int{0, 2, 0, MutedNote, 1, MutedNote}
	guitar_dim7_6_ChordShape = newChordShape("guitar", []int{0, 3, 6, 9}, guitar_dim7_6_Schema, 2)
)

var (
	guitar_dim7_7_Schema     = []int{0, 2, 0, MutedNote, 1, MutedNote}
	guitar_dim7_7_ChordShape = newChordShape("guitar", []int{0, 3, 6, 9}, guitar_dim7_7_Schema, 4)
)

var (
	guitar_dim7_8_Schema     = []int{MutedNote, 2, 0, 2, 1, MutedNote}
	guitar_dim7_8_ChordShape = newChordShape("guitar", []int{0, 3, 6, 9}, guitar_dim7_8_Schema, 1)
)

var (
	guitar_dim7_9_Schema     = []int{MutedNote, 2, 0, 2, 1, MutedNote}
	guitar_dim7_9_ChordShape = newChordShape("guitar", []int{0, 3, 6, 9}, guitar_dim7_9_Schema, 2)
)

var (
	guitar_dim7_10_Schema     = []int{MutedNote, 2, 0, 2, 1, MutedNote}
	guitar_dim7_10_ChordShape = newChordShape("guitar", []int{0, 3, 6, 9}, guitar_dim7_10_Schema, 3)
)

var (
	guitar_dim7_11_Schema     = []int{MutedNote, 2, 0, 2, 1, MutedNote}
	guitar_dim7_11_ChordShape = newChordShape("guitar", []int{0, 3, 6, 9}, guitar_dim7_11_Schema, 4)
)

var (
	guitar_dim_0_Schema     = []int{} // TODO
	guitar_dim_0_ChordShape = newChordShape("guitar", []int{0, 3, 6}, guitar_dim_0_Schema, -1)
)

var (
	guitar_7_0_Schema     = []int{MutedNote, 0, 1, 0, MutedNote, 0}
	guitar_7_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10}, guitar_7_0_Schema, 5)
)

var (
	guitar_7_1_Schema     = []int{MutedNote, 0, 2, 1, MutedNote, 2}
	guitar_7_1_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10}, guitar_7_1_Schema, 1)
)

var (
	guitar_7_2_Schema     = []int{MutedNote, 2, 0, 2, 0, MutedNote}
	guitar_7_2_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10}, guitar_7_2_Schema, 4)
)

var (
	guitar_sus2_0_Schema     = []int{} // TODO
	guitar_sus2_0_ChordShape = newChordShape("guitar", []int{0, 2, 7}, guitar_sus2_0_Schema, -1)
)

var (
	guitar_dim11_0_Schema     = []int{} // TODO
	guitar_dim11_0_ChordShape = newChordShape("guitar", []int{0, 3, 6, 9, 14, 17}, guitar_dim11_0_Schema, -1)
)

var (
	guitar_dim9_0_Schema     = []int{} // TODO
	guitar_dim9_0_ChordShape = newChordShape("guitar", []int{0, 3, 6, 9, 14}, guitar_dim9_0_Schema, -1)
)

var (
	guitar_m11_0_Schema     = []int{MutedNote, 0, 2, 2, MutedNote, 2}
	guitar_m11_0_ChordShape = newChordShape("guitar", []int{0, 3, 7, 10, 14, 17}, guitar_m11_0_Schema, 5)
)

var (
	guitar_m11_1_Schema     = []int{0, 1, 0, 0, 0, MutedNote}
	guitar_m11_1_ChordShape = newChordShape("guitar", []int{0, 3, 7, 10, 14, 17}, guitar_m11_1_Schema, 4)
)

var (
	guitar_9_0_Schema     = []int{1, 1, 0, 1, 0, MutedNote}
	guitar_9_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10, 14}, guitar_9_0_Schema, 0)
)

var (
	guitar_9_1_Schema     = []int{1, 1, 1, 0, 1, MutedNote}
	guitar_9_1_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10, 14}, guitar_9_1_Schema, 4)
)

var (
	guitar_7b5_b9_0_Schema     = []int{} // TODO
	guitar_7b5_b9_0_ChordShape = newChordShape("guitar", []int{0, 4, 6, 10, 13}, guitar_7b5_b9_0_Schema, -1)
)

var (
	guitar_9b5_0_Schema     = []int{} // TODO
	guitar_9b5_0_ChordShape = newChordShape("guitar", []int{0, 4, 6, 10, 14}, guitar_9b5_0_Schema, -1)
)

var (
	guitar_7b9_0_Schema     = []int{0, 0, 1, 0, MutedNote, 1}
	guitar_7b9_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10, 13}, guitar_7b9_0_Schema, 0)
)

var (
	guitar_7b9_1_Schema     = []int{MutedNote, 0, 1, 0, 1, MutedNote}
	guitar_7b9_1_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10, 13}, guitar_7b9_1_Schema, 4)
)

var (
	guitar_6_0_Schema     = []int{} // TODO
	guitar_6_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 9}, guitar_6_0_Schema, -1)
)

var (
	guitar_7sus_0_Schema     = []int{} // TODO
	guitar_7sus_0_ChordShape = newChordShape("guitar", []int{0, 5, 7, 10}, guitar_7sus_0_Schema, -1)
)

var (
	guitar_m_add9_0_Schema     = []int{} // TODO
	guitar_m_add9_0_ChordShape = newChordShape("guitar", []int{0, 3, 7, 14}, guitar_m_add9_0_Schema, -1)
)

var (
	guitar_aug_0_Schema     = []int{} // TODO
	guitar_aug_0_ChordShape = newChordShape("guitar", []int{0, 4, 8}, guitar_aug_0_Schema, -1)
)

var (
	guitar_m7_0_Schema     = []int{MutedNote, 0, 0, 0, MutedNote, 0}
	guitar_m7_0_ChordShape = newChordShape("guitar", []int{0, 3, 7, 10}, guitar_m7_0_Schema, 5)
)

var (
	guitar_m7_1_Schema     = []int{MutedNote, 0, 2, 0, MutedNote, 2}
	guitar_m7_1_ChordShape = newChordShape("guitar", []int{0, 3, 7, 10}, guitar_m7_1_Schema, 1)
)

var (
	guitar_m7_2_Schema     = []int{MutedNote, 1, 0, 2, 0, MutedNote}
	guitar_m7_2_ChordShape = newChordShape("guitar", []int{0, 3, 7, 10}, guitar_m7_2_Schema, 4)
)

var (
	guitar_7b5_0_Schema     = []int{MutedNote, 0, 2, 1, MutedNote, 1}
	guitar_7b5_0_ChordShape = newChordShape("guitar", []int{0, 4, 6, 10}, guitar_7b5_0_Schema, 5)
)

var (
	guitar_7b5_1_Schema     = []int{MutedNote, 0, 2, 1, MutedNote, 1}
	guitar_7b5_1_ChordShape = newChordShape("guitar", []int{0, 4, 6, 10}, guitar_7b5_1_Schema, 1)
)

var (
	guitar_7b5_2_Schema     = []int{MutedNote, 2, 0, 1, 0, MutedNote}
	guitar_7b5_2_ChordShape = newChordShape("guitar", []int{0, 4, 6, 10}, guitar_7b5_2_Schema, 4)
)

var (
	guitar_maj7b5_0_Schema     = []int{} // TODO
	guitar_maj7b5_0_ChordShape = newChordShape("guitar", []int{0, 4, 6, 11}, guitar_maj7b5_0_Schema, -1)
)

var (
	guitar_add9_0_Schema     = []int{} // TODO
	guitar_add9_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 14}, guitar_add9_0_Schema, -1)
)

var (
	guitar_6_add9_0_Schema     = []int{} // TODO
	guitar_6_add9_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 9, 14}, guitar_6_add9_0_Schema, -1)
)

var (
	guitar_m7b5_0_Schema     = []int{MutedNote, 0, 1, 1, MutedNote, 1}
	guitar_m7b5_0_ChordShape = newChordShape("guitar", []int{0, 3, 6, 10}, guitar_m7b5_0_Schema, 5)
)

var (
	guitar_m7b5_1_Schema     = []int{MutedNote, 0, 2, 0, MutedNote, 1}
	guitar_m7b5_1_ChordShape = newChordShape("guitar", []int{0, 3, 6, 10}, guitar_m7b5_1_Schema, 1)
)

var (
	guitar_m7b5_2_Schema     = []int{MutedNote, 1, 0, 1, 0, MutedNote}
	guitar_m7b5_2_ChordShape = newChordShape("guitar", []int{0, 3, 6, 10}, guitar_m7b5_2_Schema, 4)
)

var (
	guitar_11_0_Schema     = []int{MutedNote, 0, 3, 2, MutedNote, 2}
	guitar_11_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10, 14, 17}, guitar_11_0_Schema, 5)
)

var (
	guitar_11_1_Schema     = []int{0, 2, 0, 0, 0, MutedNote}
	guitar_11_1_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10, 14, 17}, guitar_11_1_Schema, 4)
)

var (
	guitar_maj9_0_Schema     = []int{MutedNote, 0, 2, 4, MutedNote, 3}
	guitar_maj9_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 11, 14}, guitar_maj9_0_Schema, 5)
)

var (
	guitar_maj9_1_Schema     = []int{MutedNote, 1, 2, 0, 1, MutedNote}
	guitar_maj9_1_ChordShape = newChordShape("guitar", []int{0, 4, 7, 11, 14}, guitar_maj9_1_Schema, 4)
)

var (
	guitar_maj7_sharp_5_0_Schema     = []int{MutedNote, 0, 1, 1, MutedNote, 0}
	guitar_maj7_sharp_5_0_ChordShape = newChordShape("guitar", []int{0, 4, 8, 11}, guitar_maj7_sharp_5_0_Schema, 5)
)

var (
	guitar_maj7_sharp_5_1_Schema     = []int{MutedNote, 2, 1, 2, 0, MutedNote}
	guitar_maj7_sharp_5_1_ChordShape = newChordShape("guitar", []int{0, 4, 8, 11}, guitar_maj7_sharp_5_1_Schema, 4)
)

var (
	guitar_maj7_sharp_5_2_Schema     = []int{MutedNote, 0, 3, 1, MutedNote, 2}
	guitar_maj7_sharp_5_2_ChordShape = newChordShape("guitar", []int{0, 4, 8, 11}, guitar_maj7_sharp_5_2_Schema, 1)
)

var (
	guitar_aug7_0_Schema     = []int{MutedNote, 1, 1, 0, MutedNote, 0}
	guitar_aug7_0_ChordShape = newChordShape("guitar", []int{0, 4, 8, 10}, guitar_aug7_0_Schema, 5)
)

var (
	guitar_aug7_1_Schema     = []int{1, 2, 0, MutedNote, 0, MutedNote}
	guitar_aug7_1_ChordShape = newChordShape("guitar", []int{0, 4, 8, 10}, guitar_aug7_1_Schema, 4)
)

var (
	guitar_aug7_2_Schema     = []int{MutedNote, 2, 0, 3, 0, MutedNote}
	guitar_aug7_2_ChordShape = newChordShape("guitar", []int{0, 4, 8, 10}, guitar_aug7_2_Schema, 4)
)

var (
	guitar_m9_0_Schema     = []int{0, 0, 0, 4, 2, 0}
	guitar_m9_0_ChordShape = newChordShape("guitar", []int{0, 3, 7, 10, 14}, guitar_m9_0_Schema, 5)
)

var (
	guitar_m9_1_Schema     = []int{MutedNote, 2, 2, 0, 2, MutedNote}
	guitar_m9_1_ChordShape = newChordShape("guitar", []int{0, 3, 7, 10, 14}, guitar_m9_1_Schema, 4)
)

var (
	guitar_aug11_0_Schema     = []int{} // TODO
	guitar_aug11_0_ChordShape = newChordShape("guitar", []int{0, 4, 8, 10, 14, 17}, guitar_aug11_0_Schema, -1)
)

var (
	guitar_m7_sharp_5_0_Schema     = []int{}
	guitar_m7_sharp_5_0_ChordShape = newChordShape("guitar", []int{0, 3, 8, 10}, guitar_m7_sharp_5_0_Schema, -1)
)

var (
	guitar_7b5_sharp_9_0_Schema     = []int{}
	guitar_7b5_sharp_9_0_ChordShape = newChordShape("guitar", []int{0, 4, 6, 10, 15}, guitar_7b5_sharp_9_0_Schema, -1)
)

var (
	guitar_7_sharp_11_0_Schema     = []int{}
	guitar_7_sharp_11_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10, 18}, guitar_7_sharp_11_0_Schema, -1)
)

var (
	guitar_9_sharp_11_0_Schema     = []int{}
	guitar_9_sharp_11_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10, 14, 18}, guitar_9_sharp_11_0_Schema, -1)
)

var (
	guitar_maj9_sharp_11_0_Schema     = []int{}
	guitar_maj9_sharp_11_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 11, 14, 18}, guitar_maj9_sharp_11_0_Schema, -1)
)

var (
	guitar_7_sharp_5_b9_0_Schema     = []int{}
	guitar_7_sharp_5_b9_0_ChordShape = newChordShape("guitar", []int{0, 4, 8, 10, 13}, guitar_7_sharp_5_b9_0_Schema, -1)
)

var (
	guitar_7_sharp_5_sharp_9_0_Schema     = []int{}
	guitar_7_sharp_5_sharp_9_0_ChordShape = newChordShape("guitar", []int{0, 4, 8, 10, 15}, guitar_7_sharp_5_sharp_9_0_Schema, -1)
)

var (
	guitar_7_sharp_9_0_Schema     = []int{MutedNote, MutedNote, 1, 1, 0, 1}
	guitar_7_sharp_9_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10, 15}, guitar_7_sharp_9_0_Schema, 5)
)

var (
	guitar_7_sharp_9_1_Schema     = []int{MutedNote, 2, 1, 0, 1, MutedNote}
	guitar_7_sharp_9_1_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10, 15}, guitar_7_sharp_9_1_Schema, 4)
)

var (
	guitar_maj7_sharp_11_0_Schema     = []int{}
	guitar_maj7_sharp_11_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 11, 18}, guitar_maj7_sharp_11_0_Schema, -1)
)

var (
	guitar_9_sharp_5_0_Schema     = []int{}
	guitar_9_sharp_5_0_ChordShape = newChordShape("guitar", []int{0, 4, 8, 10, 14}, guitar_9_sharp_5_0_Schema, -1)
)

var (
	guitar_maj7_0_Schema     = []int{}
	guitar_maj7_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 11}, guitar_maj7_0_Schema, -1)
)
var allGuitarChordShapes []ChordShape = []ChordShape{guitar_11_0_ChordShape, guitar_11_1_ChordShape, guitar_6_0_ChordShape, guitar_6_add9_0_ChordShape, guitar_7_0_ChordShape, guitar_7_1_ChordShape, guitar_7_2_ChordShape, guitar_7_sharp_11_0_ChordShape, guitar_7_sharp_5_b9_0_ChordShape, guitar_7_sharp_5_sharp_9_0_ChordShape, guitar_7_sharp_9_0_ChordShape, guitar_7_sharp_9_1_ChordShape, guitar_7b5_0_ChordShape, guitar_7b5_1_ChordShape, guitar_7b5_2_ChordShape, guitar_7b5_b9_0_ChordShape, guitar_7b5_sharp_9_0_ChordShape, guitar_7b9_0_ChordShape, guitar_7b9_1_ChordShape, guitar_7sus_0_ChordShape, guitar_9_0_ChordShape, guitar_9_1_ChordShape, guitar_9_sharp_11_0_ChordShape, guitar_9_sharp_5_0_ChordShape, guitar_9b5_0_ChordShape, guitar__0_ChordShape, guitar__1_ChordShape, guitar__2_ChordShape, guitar__3_ChordShape, guitar__4_ChordShape, guitar_add9_0_ChordShape, guitar_aug11_0_ChordShape, guitar_aug6_0_ChordShape, guitar_aug7_0_ChordShape, guitar_aug7_1_ChordShape, guitar_aug7_2_ChordShape, guitar_aug_0_ChordShape, guitar_dim11_0_ChordShape, guitar_dim6_0_ChordShape, guitar_dim7_0_ChordShape, guitar_dim7_10_ChordShape, guitar_dim7_11_ChordShape, guitar_dim7_1_ChordShape, guitar_dim7_2_ChordShape, guitar_dim7_3_ChordShape, guitar_dim7_4_ChordShape, guitar_dim7_5_ChordShape, guitar_dim7_6_ChordShape, guitar_dim7_7_ChordShape, guitar_dim7_8_ChordShape, guitar_dim7_9_ChordShape, guitar_dim9_0_ChordShape, guitar_dim_0_ChordShape, guitar_m11_0_ChordShape, guitar_m11_1_ChordShape, guitar_m7_0_ChordShape, guitar_m7_1_ChordShape, guitar_m7_2_ChordShape, guitar_m7_sharp_5_0_ChordShape, guitar_m7b5_0_ChordShape, guitar_m7b5_1_ChordShape, guitar_m7b5_2_ChordShape, guitar_m9_0_ChordShape, guitar_m9_1_ChordShape, guitar_m_0_ChordShape, guitar_m_1_ChordShape, guitar_m_2_ChordShape, guitar_m_add9_0_ChordShape, guitar_m_maj7_0_ChordShape, guitar_m_maj7_1_ChordShape, guitar_m_maj7_2_ChordShape, guitar_maj7_0_ChordShape, guitar_maj7_sharp_11_0_ChordShape, guitar_maj7_sharp_5_0_ChordShape, guitar_maj7_sharp_5_1_ChordShape, guitar_maj7_sharp_5_2_ChordShape, guitar_maj7b5_0_ChordShape, guitar_maj9_0_ChordShape, guitar_maj9_1_ChordShape, guitar_maj9_sharp_11_0_ChordShape, guitar_sus2_0_ChordShape, guitar_sus_0_ChordShape}
