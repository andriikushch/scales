package scales

import (
	"fmt"
	"github.com/andriikushch/scales/pkg/internal"
)

// Open position chords

// C major (C E G) — root on string 5 (index 4)
var openCMaj = []int{0, 1, 0, 2, 3, MutedNote}
var guitarChordShapeCMaj = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, openCMaj, 4)

// A major (A C# E) — root on string 5 (index 4)
var openAMaj = []int{0, 2, 2, 2, 0, MutedNote}
var guitarChordShapeAMaj = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, openAMaj, 4)

// A minor (A C E) — root on string 5 (index 4)
var openAMin = []int{0, 1, 2, 2, 0, MutedNote}
var guitarChordShapeAMin = newChordShape("guitar", []int{internal.IUnison, internal.Im3, internal.IP5}, openAMin, 4)

// D major (D F# A) — root on string 4 (index 3)
var openDMaj = []int{2, 3, 2, 0, MutedNote, MutedNote}
var guitarChordShapeDMaj = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, openDMaj, 3)

// D minor (D F A) — root on string 4 (index 3)
var openDMin = []int{1, 3, 2, 0, MutedNote, MutedNote}
var guitarChordShapeDMin = newChordShape("guitar", []int{internal.IUnison, internal.Im3, internal.IP5}, openDMin, 3)

// E major (E G# B) — root on string 6 (index 5)
var openEMaj = []int{0, 0, 1, 2, 2, 0}
var guitarChordShapeEMaj = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, openEMaj, 5)

// E minor (E G B) — root on string 6 (index 5)
var openEMin = []int{0, 0, 0, 2, 2, 0}
var guitarChordShapeEMin = newChordShape("guitar", []int{internal.IUnison, internal.Im3, internal.IP5}, openEMin, 5)

// G major (G B D) — root on string 6 (index 5)
var openGMaj = []int{3, 0, 0, 0, 2, 3}
var guitarChordShapeGMaj = newChordShape("guitar", []int{internal.IUnison, internal.IM3, internal.IP5}, openGMaj, 5)

var guitarChordShapes = map[string][]chordShape{
	// []int{IUnison, Im3, IP5} - minor chord
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.Im3, internal.IP5): []chordShape{
		guitarChordShapeAMin,
		guitarChordShapeDMin,
		guitarChordShapeEMin,
	},
	fmt.Sprintf("%d-%d-%d", internal.IUnison, internal.IM3, internal.IP5): []chordShape{
		guitarChordShapeCMaj,
		guitarChordShapeAMaj,
		guitarChordShapeDMaj,
		guitarChordShapeEMaj,
		guitarChordShapeGMaj,
	},
}
