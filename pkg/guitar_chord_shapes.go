package scales

import (
	"fmt"

	"github.com/andriikushch/scales/pkg/internal"
)

var GuitarChordShapes = map[string][]ChordShape{
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.Im3, internal.IP5): {
		guitar_m_0_ChordShape,
		guitar_m_1_ChordShape,
		guitar_m_2_ChordShape,
	},
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.IM3, internal.IP5): {
		guitar__0_ChordShape,
		guitar__1_ChordShape,
		guitar__2_ChordShape,
		guitar__3_ChordShape,
		guitar__4_ChordShape,
	},
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
	guitar_m_maj7_0_Schema     = []int{} // TODO
	guitar_m_maj7_0_ChordShape = newChordShape("guitar", []int{0, 3, 7, 11}, guitar_m_maj7_0_Schema, -1)
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
	guitar_dim_0_Schema     = []int{} // TODO
	guitar_dim_0_ChordShape = newChordShape("guitar", []int{0, 3, 6}, guitar_dim_0_Schema, -1)
)

var (
	guitar_7_0_Schema     = []int{} // TODO
	guitar_7_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10}, guitar_7_0_Schema, -1)
)

var (
	guitar_7_5_9_0_Schema     = []int{} // TODO
	guitar_7_5_9_0_ChordShape = newChordShape("guitar", []int{0, 4, 8, 10, 15}, guitar_7_5_9_0_Schema, -1)
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
	guitar_m11_0_Schema     = []int{} // TODO
	guitar_m11_0_ChordShape = newChordShape("guitar", []int{0, 3, 7, 10, 14, 17}, guitar_m11_0_Schema, -1)
)

var (
	guitar_9_11_0_Schema     = []int{} // TODO
	guitar_9_11_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10, 14, 18}, guitar_9_11_0_Schema, -1)
)

var (
	guitar_9_0_Schema     = []int{} // TODO
	guitar_9_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10, 14}, guitar_9_0_Schema, -1)
)

var (
	guitar_dim6_0_Schema     = []int{} // TODO
	guitar_dim6_0_ChordShape = newChordShape("guitar", []int{0, 3, 6, 9}, guitar_dim6_0_Schema, -1)
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
	guitar_7b9_0_Schema     = []int{} // TODO
	guitar_7b9_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10, 13}, guitar_7b9_0_Schema, -1)
)

var (
	guitar_maj9_11_0_Schema     = []int{} // TODO
	guitar_maj9_11_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 11, 14, 18}, guitar_maj9_11_0_Schema, -1)
)

var (
	guitar_6_0_Schema     = []int{} // TODO
	guitar_6_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 9}, guitar_6_0_Schema, -1)
)

var (
	guitar_maj7_5_0_Schema     = []int{} // TODO
	guitar_maj7_5_0_ChordShape = newChordShape("guitar", []int{0, 4, 8, 11}, guitar_maj7_5_0_Schema, -1)
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
	guitar_9_5_0_Schema     = []int{} // TODO
	guitar_9_5_0_ChordShape = newChordShape("guitar", []int{0, 4, 8, 10, 14}, guitar_9_5_0_Schema, -1)
)

var (
	guitar_aug_0_Schema     = []int{} // TODO
	guitar_aug_0_ChordShape = newChordShape("guitar", []int{0, 4, 8}, guitar_aug_0_Schema, -1)
)

var (
	guitar_m7_0_Schema     = []int{} // TODO
	guitar_m7_0_ChordShape = newChordShape("guitar", []int{0, 3, 7, 10}, guitar_m7_0_Schema, -1)
)

var (
	guitar_m7_5_0_Schema     = []int{} // TODO
	guitar_m7_5_0_ChordShape = newChordShape("guitar", []int{0, 3, 8, 10}, guitar_m7_5_0_Schema, -1)
)

var (
	guitar_7b5_0_Schema     = []int{} // TODO
	guitar_7b5_0_ChordShape = newChordShape("guitar", []int{0, 4, 6, 10}, guitar_7b5_0_Schema, -1)
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
	guitar_m7b5_0_Schema     = []int{} // TODO
	guitar_m7b5_0_ChordShape = newChordShape("guitar", []int{0, 3, 6, 10}, guitar_m7b5_0_Schema, -1)
)

var (
	guitar_11_0_Schema     = []int{} // TODO
	guitar_11_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10, 14, 17}, guitar_11_0_Schema, -1)
)

var (
	guitar_maj9_0_Schema     = []int{} // TODO
	guitar_maj9_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 11, 14}, guitar_maj9_0_Schema, -1)
)

var (
	guitar_maj7_11_0_Schema     = []int{} // TODO
	guitar_maj7_11_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 11, 18}, guitar_maj7_11_0_Schema, -1)
)

var (
	guitar_maj7_0_Schema     = []int{} // TODO
	guitar_maj7_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 11}, guitar_maj7_0_Schema, -1)
)

var (
	guitar_7_5_b9_0_Schema     = []int{} // TODO
	guitar_7_5_b9_0_ChordShape = newChordShape("guitar", []int{0, 4, 8, 10, 13}, guitar_7_5_b9_0_Schema, -1)
)

var (
	guitar_aug7_0_Schema     = []int{} // TODO
	guitar_aug7_0_ChordShape = newChordShape("guitar", []int{0, 4, 8, 10}, guitar_aug7_0_Schema, -1)
)

var (
	guitar_m9_0_Schema     = []int{} // TODO
	guitar_m9_0_ChordShape = newChordShape("guitar", []int{0, 3, 7, 10, 14}, guitar_m9_0_Schema, -1)
)

var (
	guitar_7b5_9_0_Schema     = []int{} // TODO
	guitar_7b5_9_0_ChordShape = newChordShape("guitar", []int{0, 4, 6, 10, 15}, guitar_7b5_9_0_Schema, -1)
)

var (
	guitar_7_9_0_Schema     = []int{} // TODO
	guitar_7_9_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10, 15}, guitar_7_9_0_Schema, -1)
)

var (
	guitar_7_11_0_Schema     = []int{} // TODO
	guitar_7_11_0_ChordShape = newChordShape("guitar", []int{0, 4, 7, 10, 18}, guitar_7_11_0_Schema, -1)
)

var (
	guitar_aug11_0_Schema     = []int{} // TODO
	guitar_aug11_0_ChordShape = newChordShape("guitar", []int{0, 4, 8, 10, 14, 17}, guitar_aug11_0_Schema, -1)
)
