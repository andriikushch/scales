package main

// scaleEntry pairs a pkg/registry.go registry key with the display name docs/scales.md shows for
// it (accidentals spelled out with Unicode glyphs, matching the registry key comment convention
// at pkg/registry.go:8-10).
type scaleEntry struct {
	Key         string
	DisplayName string
}

// scaleCategory is one section of docs/scales.md's per-family tables, mirroring a comment section
// in pkg/registry.go's scaleRegistry map, in the same scale order.
type scaleCategory struct {
	Title  string
	Scales []scaleEntry
}

// scaleCategories mirrors pkg/registry.go's canonical (non-alias) section comments and key order
// exactly, with one deliberate change: chinesePentatonic's display name is "Chinese Pentatonic"
// rather than the previously ambiguous bare "Chinese", which collided with the unrelated
// "Chinese" alias for Major Pentatonic (see pkg/pentatonic.go:153-157 and the "chinese" entry in
// scaleAliases below).
var scaleCategories = []scaleCategory{
	{
		Title: "Modes of the major scale",
		Scales: []scaleEntry{
			{"ionian", "Ionian"},
			{"dorian", "Dorian"},
			{"phrygian", "Phrygian"},
			{"lydian", "Lydian"},
			{"mixolydian", "Mixolydian"},
			{"aeolian", "Aeolian"},
			{"locrian", "Locrian"},
		},
	},
	{
		Title: "Modes of the melodic minor scale",
		Scales: []scaleEntry{
			{"melodicMinor", "Melodic Minor"},
			{"dorianFlat2", "Dorian ♭2"},
			{"lydianAugmented", "Lydian Augmented"},
			{"lydianDominant", "Lydian Dominant"},
			{"mixolydianFlat6", "Mixolydian ♭6"},
			{"locrianNatural2", "Locrian ♮2"},
			{"altered", "Altered"},
		},
	},
	{
		Title: "Modes of the harmonic minor scale",
		Scales: []scaleEntry{
			{"harmonicMinor", "Harmonic Minor"},
			{"locrianNatural6", "Locrian ♮6"},
			{"ionianSharp5", "Ionian ♯5"},
			{"dorianSharp4", "Dorian ♯4"},
			{"phrygianDominant", "Phrygian Dominant"},
			{"lydianSharp2", "Lydian ♯2"},
			{"ultralocrian", "Ultralocrian"},
		},
	},
	{
		Title: "Modes of the harmonic major scale",
		Scales: []scaleEntry{
			{"harmonicMajor", "Harmonic Major"},
			{"dorianFlat5", "Dorian ♭5"},
			{"phrygianFlat4", "Phrygian ♭4"},
			{"lydianFlat3", "Lydian ♭3"},
			{"mixolydianFlat2", "Mixolydian ♭2"},
			{"lydianAugmentedSharp2", "Lydian Augmented ♯2"},
			{"locrianDoubleFlat7", "Locrian ♭♭7"},
		},
	},
	{
		Title: "Modes of the double harmonic major scale",
		Scales: []scaleEntry{
			{"doubleHarmonicMajor", "Double Harmonic Major"},
			{"lydianSharp2Sharp6", "Lydian ♯2 ♯6"},
			{"ultraphrygian", "Ultraphrygian"},
			{"hungarianMinor", "Hungarian Minor"},
			{"oriental", "Oriental"},
			{"ionianSharp2Sharp5", "Ionian ♯2 ♯5"},
			{"locrianDoubleFlat3DoubleFlat7", "Locrian ♭♭3 ♭♭7"},
		},
	},
	{
		Title: "Other heptatonic scales",
		Scales: []scaleEntry{
			{"neapolitanMinor", "Neapolitan Minor"},
			{"neapolitanMajor", "Neapolitan Major"},
			{"leadingWholeTone", "Leading Whole Tone"},
			{"majorLocrian", "Major Locrian"},
			{"gypsy", "Gypsy"},
			{"hungarianMajor", "Hungarian Major"},
			{"persian", "Persian"},
			{"enigmatic", "Enigmatic"},
		},
	},
	{
		Title: "Pentatonic scales",
		Scales: []scaleEntry{
			{"majorPentatonic", "Major Pentatonic"},
			{"suspendedPentatonic", "Suspended Pentatonic"},
			{"bluesMinorPentatonic", "Blues Minor Pentatonic"},
			{"bluesMajorPentatonic", "Blues Major Pentatonic"},
			{"minorPentatonic", "Minor Pentatonic"},
			{"iwato", "Iwato"},
			{"miyakoBushi", "Miyako-bushi"},
			{"kumoi", "Kumoi"},
			{"insen", "Insen"},
			{"chinesePentatonic", "Chinese Pentatonic"},
			{"balinesePelog", "Balinese Pelog"},
		},
	},
	{
		Title: "Hexatonic scales",
		Scales: []scaleEntry{
			{"wholeTone", "Whole Tone"},
			{"augmented", "Augmented"},
			{"blues", "Blues"},
			{"majorBlues", "Major Blues"},
			{"majorHexatonic", "Major Hexatonic"},
			{"minorHexatonic", "Minor Hexatonic"},
			{"prometheus", "Prometheus"},
			{"prometheusNeapolitan", "Prometheus Neapolitan"},
			{"tritone", "Tritone"},
			{"istrian", "Istrian"},
			{"scaleOfHarmonics", "Scale of Harmonics"},
			{"twoSemitoneTritone", "Two-Semitone Tritone"},
			{"messiaenMode5", "Messiaen Mode 5"},
		},
	},
	{
		Title: "Octatonic scales",
		Scales: []scaleEntry{
			{"halfWholeDiminished", "Half-Whole Diminished"},
			{"wholeHalfDiminished", "Whole-Half Diminished"},
			{"messiaenMode4", "Messiaen Mode 4"},
			{"messiaenMode6", "Messiaen Mode 6"},
			{"spanishEightTone", "Spanish Eight-Tone"},
		},
	},
	{
		Title: "Bebop scales (8 notes: a 7-note scale plus one chromatic passing tone)",
		Scales: []scaleEntry{
			{"bebopDominant", "Bebop Dominant"},
			{"bebopMajor", "Bebop Major"},
			{"bebopDorian", "Bebop Dorian"},
			{"bebopMelodicMinor", "Bebop Melodic Minor"},
			{"bebopHarmonicMinor", "Bebop Harmonic Minor"},
		},
	},
	{
		Title: "Nine or more notes",
		Scales: []scaleEntry{
			{"messiaenMode3", "Messiaen Mode 3"},
			{"messiaenMode7", "Messiaen Mode 7"},
			{"chromatic", "Chromatic"},
		},
	},
}

// scaleAliasEntry pairs a pkg/registry.go alias-section key with its display name and the
// display name of the canonical scale it aliases (see pkg/scale_aliases.go).
type scaleAliasEntry struct {
	Key                  string
	DisplayName          string
	CanonicalDisplayName string
}

// scaleAliases mirrors pkg/registry.go's "Alternative names" section exactly (59 keys).
var scaleAliases = []scaleAliasEntry{
	{"acoustic", "Acoustic", "Lydian Dominant"},
	{"adonaiMalakh", "Adonai malakh", "Mixolydian"},
	{"aeolianDominant", "Aeolian dominant", "Mixolydian ♭6"},
	{"aeolianFlat5", "Aeolian ♭5", "Locrian ♮2"},
	{"ahavaRabba", "Ahava Rabba", "Phrygian Dominant"},
	{"alteredDoubleFlat7", "Altered ♭♭7", "Ultralocrian"},
	{"alteredNatural5", "Altered ♮5", "Phrygian ♭4"},
	{"andalusian", "Andalusian", "Phrygian"},
	{"arabic", "Arabic (one usage)", "Double Harmonic Major"},
	{"augmentedMajor", "Augmented major", "Ionian ♯5"},
	{"byzantine", "Byzantine", "Double Harmonic Major"},
	{"chinese", "Chinese", "Major Pentatonic"},
	{"dominantDiminished", "Dominant diminished", "Half-Whole Diminished"},
	{"dorianAeolianWithoutThe6th", "Dorian/Aeolian without the 6th", "Minor Hexatonic"},
	{"doubleHarmonicMinor", "Double harmonic minor", "Hungarian Minor"},
	{"freygish", "Freygish", "Phrygian Dominant"},
	{"gypsyMajor", "Gypsy major", "Double Harmonic Major"},
	{"gypsyMinor", "Gypsy minor", "Hungarian Minor"},
	{"halfDiminished", "Half-diminished", "Locrian ♮2"},
	{"hawaiian", "Hawaiian", "Melodic Minor"},
	{"hijaz", "Hijaz", "Phrygian Dominant"},
	{"hindu", "Hindu", "Mixolydian ♭6"},
	{"ionianFlat6", "Ionian ♭6", "Harmonic Major"},
	{"jazzMinor", "Jazz minor", "Melodic Minor"},
	{"korean", "Korean", "Blues Major Pentatonic"},
	{"locrianDiminished", "Locrian diminished", "Locrian ♭♭7"},
	{"locrianNatural2Natural6", "Locrian ♮2 ♮6", "Dorian ♭5"},
	{"lydianDiminished", "Lydian diminished", "Lydian ♭3"},
	{"lydianFlat7", "Lydian ♭7", "Lydian Dominant"},
	{"major", "Major", "Ionian"},
	{"majorScaleWithoutThe7th", "Major scale without the 7th", "Major Hexatonic"},
	{"manGong", "Man Gong", "Blues Minor Pentatonic"},
	{"mayamalavagowla", "Mayamalavagowla", "Double Harmonic Major"},
	{"melodicMinorFlat2", "Melodic minor ♭2", "Neapolitan Major"},
	{"melodicMinorSharp4", "Melodic minor ♯4", "Lydian ♭3"},
	{"misheberak", "Misheberak", "Dorian ♯4"},
	{"mixolydianSharp4", "Mixolydian ♯4", "Lydian Dominant"},
	{"mongolian", "Mongolian", "Major Pentatonic"},
	{"mysticChord", "Mystic chord", "Prometheus"},
	{"naturalMinor", "Natural minor", "Aeolian"},
	{"overtone", "Overtone", "Lydian Dominant"},
	{"phrygianNatural6", "Phrygian ♮6", "Dorian ♭2"},
	{"prometheusLiszt", "Prometheus Liszt", "Prometheus Neapolitan"},
	{"ragaBhairav", "Raga Bhairav", "Double Harmonic Major"},
	{"ragaBhupali", "Raga Bhupali", "Major Pentatonic"},
	{"ragaDurga", "Raga Durga; 4th mode", "Blues Major Pentatonic"},
	{"ragaMalkauns", "Raga Malkauns; 3rd mode", "Blues Minor Pentatonic"},
	{"ritsu", "Ritsu", "Blues Major Pentatonic"},
	{"ritsusen", "Ritsusen", "Blues Major Pentatonic"},
	{"romanianMinor", "Romanian minor", "Dorian ♯4"},
	{"sakura", "Sakura", "Miyako-bushi"},
	{"slendro", "Slendro (approx.)", "Major Pentatonic"},
	{"spanishGypsy", "Spanish Gypsy", "Phrygian Dominant"},
	{"superLocrian", "Super-Locrian", "Altered"},
	{"superlocrianDoubleFlat7", "Superlocrian ♭♭7", "Ultralocrian"},
	{"thai", "Thai", "Blues Major Pentatonic"},
	{"ukrainianDorian", "Ukrainian Dorian", "Dorian ♯4"},
	{"verdisAscendingForm", "Verdi's ascending form", "Enigmatic"},
	{"yo", "Yo", "Blues Major Pentatonic"},
}
