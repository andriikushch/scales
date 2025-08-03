package scales

import (
	"fmt"

	"github.com/andriikushch/scales/pkg/internal"
)

var UkuleleChordShapes = map[string][]ChordShape{
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.IM3, internal.IP5): {
		ukulele__0_ChordShape,
		ukulele__1_ChordShape,
		ukulele__2_ChordShape,
		ukulele__3_ChordShape,
		ukulele__4_ChordShape,
	},
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.Im3, internal.IP5): {
		ukulele_m_0_ChordShape,
		ukulele_m_1_ChordShape,
		ukulele_m_2_ChordShape,
		ukulele_m_3_ChordShape,
	},
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
