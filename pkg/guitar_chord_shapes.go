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
