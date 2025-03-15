package internal

const (
	Over       = "over"
	Added      = "added"
	Major      = "major"
	Minor      = "minor"
	Suspended  = "suspended"
	Seventh    = "seventh"
	Diminished = "diminished"
	Augmented  = "augmented"
	Dominant   = "dominant"
	Flat       = "flat"
	Sharp      = "sharp"
	Five       = "five"
	Seven      = "seven"
	Nine       = "nine"
	Eleven     = "eleven"
	Thirteen   = "thirteen"
	Second     = "second"
	Fourth     = "fourth"
	Sixth      = "sixth"
	Ninth      = "ninth"
	Eleventh   = "eleventh"
	Thirteenth = "thirteenth"
)

var NthMapping = map[int]string{
	6:  Sixth,
	7:  Seventh,
	9:  Ninth,
	11: Eleventh,
	13: Thirteenth,
}

var Degrees = map[int]int{
	5:  IP5,
	6:  IM6,
	7:  IM7,
	9:  IM9,
	11: IP11,
	13: IM13,
}

var SharpDegrees = map[int]int{
	9:  IA9,
	11: IA11,
	13: IA13,
}

var FlatDegrees = map[int]int{
	7:  Im7,
	9:  Im9,
	11: ID11,
	13: ID13,
}

var Numbers = map[int]string{
	5:  Five,
	7:  Seven,
	9:  Nine,
	11: Eleven,
	13: Thirteen,
}
