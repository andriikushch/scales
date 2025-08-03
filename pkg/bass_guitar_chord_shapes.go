package scales

import (
	"fmt"
	"github.com/andriikushch/scales/pkg/internal"
)

var BassGuitarChordShapes = map[string][]ChordShape{
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.Im3, internal.IP5): {
		bassGuitar_m_0_ChordShape,
		bassGuitar_m_1_ChordShape,
		bassGuitar_m_2_ChordShape,
	},
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.IM3, internal.IP5): {
		bassGuitar__0_ChordShape,
		bassGuitar__1_ChordShape,
		bassGuitar__2_ChordShape,
	},
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
