package scales

import (
	"fmt"

	"github.com/andriikushch/scales/pkg/internal"
)

// Open position ukulele chords (G–C–E–A tuning)
var (
	openUkuleleCMaj       = []int{3, 0, 0, 0}
	ukuleleChordShapeCMaj = newChordShape("ukulele", []int{internal.IUnison, internal.IM3, internal.IP5}, openUkuleleCMaj, 0)
)

var (
	openUkuleleAMaj       = []int{0, 0, 1, 2}
	ukuleleChordShapeAMaj = newChordShape("ukulele", []int{internal.IUnison, internal.IM3, internal.IP5}, openUkuleleAMaj, 0)
)

var (
	openUkuleleAMin       = []int{0, 0, 0, 2}
	ukuleleChordShapeAMin = newChordShape("ukulele", []int{internal.IUnison, internal.Im3, internal.IP5}, openUkuleleAMin, 3)
)

var (
	openUkuleleGMaj       = []int{2, 3, 2, 0}
	ukuleleChordShapeGMaj = newChordShape("ukulele", []int{internal.IUnison, internal.IM3, internal.IP5}, openUkuleleGMaj, 3)
)

var (
	openUkuleleGMin       = []int{1, 3, 2, 0}
	ukuleleChordShapeGMin = newChordShape("ukulele", []int{internal.IUnison, internal.Im3, internal.IP5}, openUkuleleGMin, 3)
)

var (
	openUkuleleDMaj       = []int{2, 2, 2, 0}
	ukuleleChordShapeDMaj = newChordShape("ukulele", []int{internal.IUnison, internal.IM3, internal.IP5}, openUkuleleDMaj, 2)
)

var (
	openUkuleleDMin       = []int{0, 1, 2, 2}
	ukuleleChordShapeDMin = newChordShape("ukulele", []int{internal.IUnison, internal.Im3, internal.IP5}, openUkuleleDMin, 2)
)

var (
	openUkuleleEMin       = []int{2, 3, 4, 0}
	ukuleleChordShapeEMin = newChordShape("ukulele", []int{internal.IUnison, internal.Im3, internal.IP5}, openUkuleleEMin, 2)
)

var (
	openUkuleleFMaj       = []int{0, 1, 0, 2}
	ukuleleChordShapeFMaj = newChordShape("ukulele", []int{internal.IUnison, internal.IM3, internal.IP5}, openUkuleleFMaj, 1)
)

var ukuleleChordShapes = map[string][]chordShape{
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.IM3, internal.IP5): {
		ukuleleChordShapeCMaj,
		ukuleleChordShapeAMaj,
		ukuleleChordShapeGMaj,
		ukuleleChordShapeDMaj,
		ukuleleChordShapeFMaj,
	},
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.Im3, internal.IP5): {
		ukuleleChordShapeAMin,
		ukuleleChordShapeGMin,
		ukuleleChordShapeDMin,
		ukuleleChordShapeEMin,
	},
}
