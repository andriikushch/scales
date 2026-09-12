package scales

// This file provides one constructor per alias name in docs/scales.md's "Other names" table.
// Each is a thin passthrough to the constructor of the canonical scale it names. Two aliases in
// that table ("Major" -> Ionian, "Natural minor" -> Aeolian) already have an identically-named
// canonical constructor (NewMajorScale, NewNaturalMinorScale) and so are registered directly in
// scaleRegistry without a wrapper here.

// NewAcousticScale builds the Acoustic scale, an alias for the Lydian Dominant scale.
func NewAcousticScale(key string) (*Scale, error) {
	return NewLydianDominantScale(key)
}

// NewAdonaiMalakhScale builds the Adonai malakh scale, an alias for the Mixolydian scale.
func NewAdonaiMalakhScale(key string) (*Scale, error) {
	return NewMixolydianScale(key)
}

// NewAeolianDominantScale builds the Aeolian dominant scale, an alias for the Mixolydian ♭6 scale.
func NewAeolianDominantScale(key string) (*Scale, error) {
	return NewMixolydianFlat6Scale(key)
}

// NewAeolianFlat5Scale builds the Aeolian ♭5 scale, an alias for the Locrian ♮2 scale.
func NewAeolianFlat5Scale(key string) (*Scale, error) {
	return NewLocrianNatural2Scale(key)
}

// NewAhavaRabbaScale builds the Ahava Rabba scale, an alias for the Phrygian Dominant scale.
func NewAhavaRabbaScale(key string) (*Scale, error) {
	return NewPhrygianDominantScale(key)
}

// NewAlteredDoubleFlat7Scale builds the Altered ♭♭7 scale, an alias for the Ultralocrian scale.
func NewAlteredDoubleFlat7Scale(key string) (*Scale, error) {
	return NewUltralocrianScale(key)
}

// NewAlteredNatural5Scale builds the Altered ♮5 scale, an alias for the Phrygian ♭4 scale.
func NewAlteredNatural5Scale(key string) (*Scale, error) {
	return NewPhrygianFlat4Scale(key)
}

// NewAndalusianScale builds the Andalusian scale, an alias for the Phrygian scale.
func NewAndalusianScale(key string) (*Scale, error) {
	return NewPhrygianScale(key)
}

// NewArabicScale builds the Arabic scale (one usage of the name), an alias for the Double
// Harmonic Major scale.
func NewArabicScale(key string) (*Scale, error) {
	return NewDoubleHarmonicMajorScale(key)
}

// NewAugmentedMajorScale builds the Augmented major scale, an alias for the Ionian ♯5 scale.
func NewAugmentedMajorScale(key string) (*Scale, error) {
	return NewIonianSharp5Scale(key)
}

// NewByzantineScale builds the Byzantine scale, an alias for the Double Harmonic Major scale.
func NewByzantineScale(key string) (*Scale, error) {
	return NewDoubleHarmonicMajorScale(key)
}

// NewChineseScale builds the Chinese scale, an alias for the Major Pentatonic scale.
//
// This is distinct from NewChinesePentatonicScale, which is a different, unrelated note
// collection that some sources also call "Chinese" — see the comment there.
func NewChineseScale(key string) (*Scale, error) {
	return pentatonicAsScale(NewMajorPentatonicScale, key)
}

// NewDominantDiminishedScale builds the Dominant diminished scale, an alias for the Half-Whole
// Diminished scale.
func NewDominantDiminishedScale(key string) (*Scale, error) {
	return NewHalfWholeDiminishedScale(key)
}

// NewDorianAeolianWithoutThe6thScale builds the "Dorian/Aeolian without the 6th" scale, an alias
// for the Minor Hexatonic scale.
func NewDorianAeolianWithoutThe6thScale(key string) (*Scale, error) {
	return NewMinorHexatonicScale(key)
}

// NewDoubleHarmonicMinorScale builds the Double harmonic minor scale, an alias for the Hungarian
// Minor scale.
func NewDoubleHarmonicMinorScale(key string) (*Scale, error) {
	return NewHungarianMinorScale(key)
}

// NewFreygishScale builds the Freygish scale, an alias for the Phrygian Dominant scale.
func NewFreygishScale(key string) (*Scale, error) {
	return NewPhrygianDominantScale(key)
}

// NewGypsyMajorScale builds the Gypsy major scale, an alias for the Double Harmonic Major scale.
func NewGypsyMajorScale(key string) (*Scale, error) {
	return NewDoubleHarmonicMajorScale(key)
}

// NewGypsyMinorScale builds the Gypsy minor scale, an alias for the Hungarian Minor scale.
func NewGypsyMinorScale(key string) (*Scale, error) {
	return NewHungarianMinorScale(key)
}

// NewHalfDiminishedScale builds the Half-diminished scale, an alias for the Locrian ♮2 scale.
func NewHalfDiminishedScale(key string) (*Scale, error) {
	return NewLocrianNatural2Scale(key)
}

// NewHawaiianScale builds the Hawaiian scale, an alias for the Melodic Minor scale.
func NewHawaiianScale(key string) (*Scale, error) {
	return NewMelodicMinorScale(key)
}

// NewHijazScale builds the Hijaz scale, an alias for the Phrygian Dominant scale.
func NewHijazScale(key string) (*Scale, error) {
	return NewPhrygianDominantScale(key)
}

// NewHinduScale builds the Hindu scale, an alias for the Mixolydian ♭6 scale.
func NewHinduScale(key string) (*Scale, error) {
	return NewMixolydianFlat6Scale(key)
}

// NewIonianFlat6Scale builds the Ionian ♭6 scale, an alias for the Harmonic Major scale.
func NewIonianFlat6Scale(key string) (*Scale, error) {
	return NewHarmonicMajorScale(key)
}

// NewJazzMinorScale builds the Jazz minor scale, an alias for the Melodic Minor scale.
func NewJazzMinorScale(key string) (*Scale, error) {
	return NewMelodicMinorScale(key)
}

// NewKoreanScale builds the Korean scale, an alias for the Blues Major Pentatonic scale.
func NewKoreanScale(key string) (*Scale, error) {
	return NewBluesMajorPentatonicScale(key)
}

// NewLocrianDiminishedScale builds the Locrian diminished scale, an alias for the Locrian ♭♭7
// scale.
func NewLocrianDiminishedScale(key string) (*Scale, error) {
	return NewLocrianDoubleFlat7Scale(key)
}

// NewLocrianNatural2Natural6Scale builds the Locrian ♮2 ♮6 scale, an alias for the Dorian ♭5
// scale.
func NewLocrianNatural2Natural6Scale(key string) (*Scale, error) {
	return NewDorianFlat5Scale(key)
}

// NewLydianDiminishedScale builds the Lydian diminished scale, an alias for the Lydian ♭3 scale.
func NewLydianDiminishedScale(key string) (*Scale, error) {
	return NewLydianFlat3Scale(key)
}

// NewLydianFlat7Scale builds the Lydian ♭7 scale, an alias for the Lydian Dominant scale.
func NewLydianFlat7Scale(key string) (*Scale, error) {
	return NewLydianDominantScale(key)
}

// NewMajorScaleWithoutThe7thScale builds the "Major scale without the 7th" scale, an alias for
// the Major Hexatonic scale.
func NewMajorScaleWithoutThe7thScale(key string) (*Scale, error) {
	return NewMajorHexatonicScale(key)
}

// NewManGongScale builds the Man Gong scale, an alias for the Blues Minor Pentatonic scale.
func NewManGongScale(key string) (*Scale, error) {
	return NewBluesMinorPentatonicScale(key)
}

// NewMayamalavagowlaScale builds the Mayamalavagowla scale, an alias for the Double Harmonic
// Major scale.
func NewMayamalavagowlaScale(key string) (*Scale, error) {
	return NewDoubleHarmonicMajorScale(key)
}

// NewMelodicMinorFlat2Scale builds the Melodic minor ♭2 scale, an alias for the Neapolitan Major
// scale.
func NewMelodicMinorFlat2Scale(key string) (*Scale, error) {
	return NewNeapolitanMajorScale(key)
}

// NewMelodicMinorSharp4Scale builds the Melodic minor ♯4 scale, an alias for the Lydian ♭3 scale.
func NewMelodicMinorSharp4Scale(key string) (*Scale, error) {
	return NewLydianFlat3Scale(key)
}

// NewMisheberakScale builds the Misheberak scale, an alias for the Dorian ♯4 scale.
func NewMisheberakScale(key string) (*Scale, error) {
	return NewDorianSharp4Scale(key)
}

// NewMixolydianSharp4Scale builds the Mixolydian ♯4 scale, an alias for the Lydian Dominant
// scale.
func NewMixolydianSharp4Scale(key string) (*Scale, error) {
	return NewLydianDominantScale(key)
}

// NewMongolianScale builds the Mongolian scale, an alias for the Major Pentatonic scale.
func NewMongolianScale(key string) (*Scale, error) {
	return pentatonicAsScale(NewMajorPentatonicScale, key)
}

// NewMysticChordScale builds the Mystic chord scale, an alias for the Prometheus scale.
func NewMysticChordScale(key string) (*Scale, error) {
	return NewPrometheusScale(key)
}

// NewOvertoneScale builds the Overtone scale, an alias for the Lydian Dominant scale.
func NewOvertoneScale(key string) (*Scale, error) {
	return NewLydianDominantScale(key)
}

// NewPhrygianNatural6Scale builds the Phrygian ♮6 scale, an alias for the Dorian ♭2 scale.
func NewPhrygianNatural6Scale(key string) (*Scale, error) {
	return NewDorianFlat2Scale(key)
}

// NewPrometheusLisztScale builds the Prometheus Liszt scale, an alias for the Prometheus
// Neapolitan scale.
func NewPrometheusLisztScale(key string) (*Scale, error) {
	return NewPrometheusNeapolitanScale(key)
}

// NewRagaBhairavScale builds the Raga Bhairav scale, an alias for the Double Harmonic Major
// scale.
func NewRagaBhairavScale(key string) (*Scale, error) {
	return NewDoubleHarmonicMajorScale(key)
}

// NewRagaBhupaliScale builds the Raga Bhupali scale, an alias for the Major Pentatonic scale.
func NewRagaBhupaliScale(key string) (*Scale, error) {
	return pentatonicAsScale(NewMajorPentatonicScale, key)
}

// NewRagaDurgaScale builds the Raga Durga scale (its 4th mode), an alias for the Blues Major
// Pentatonic scale.
func NewRagaDurgaScale(key string) (*Scale, error) {
	return NewBluesMajorPentatonicScale(key)
}

// NewRagaMalkaunsScale builds the Raga Malkauns scale (its 3rd mode), an alias for the Blues
// Minor Pentatonic scale.
func NewRagaMalkaunsScale(key string) (*Scale, error) {
	return NewBluesMinorPentatonicScale(key)
}

// NewRitsuScale builds the Ritsu scale, an alias for the Blues Major Pentatonic scale.
func NewRitsuScale(key string) (*Scale, error) {
	return NewBluesMajorPentatonicScale(key)
}

// NewRitsusenScale builds the Ritsusen scale, an alias for the Blues Major Pentatonic scale.
func NewRitsusenScale(key string) (*Scale, error) {
	return NewBluesMajorPentatonicScale(key)
}

// NewRomanianMinorScale builds the Romanian minor scale, an alias for the Dorian ♯4 scale.
func NewRomanianMinorScale(key string) (*Scale, error) {
	return NewDorianSharp4Scale(key)
}

// NewSakuraScale builds the Sakura scale, an alias for the Miyako-bushi scale.
func NewSakuraScale(key string) (*Scale, error) {
	return NewMiyakoBushiScale(key)
}

// NewSlendroScale builds the Slendro scale (a 12-tone-equal-temperament approximation), an alias
// for the Major Pentatonic scale.
func NewSlendroScale(key string) (*Scale, error) {
	return pentatonicAsScale(NewMajorPentatonicScale, key)
}

// NewSpanishGypsyScale builds the Spanish Gypsy scale, an alias for the Phrygian Dominant scale.
func NewSpanishGypsyScale(key string) (*Scale, error) {
	return NewPhrygianDominantScale(key)
}

// NewSuperLocrianScale builds the Super-Locrian scale, an alias for the Altered scale.
func NewSuperLocrianScale(key string) (*Scale, error) {
	return NewAlteredScale(key)
}

// NewSuperlocrianDoubleFlat7Scale builds the Superlocrian ♭♭7 scale, an alias for the
// Ultralocrian scale.
func NewSuperlocrianDoubleFlat7Scale(key string) (*Scale, error) {
	return NewUltralocrianScale(key)
}

// NewThaiScale builds the Thai scale, an alias for the Blues Major Pentatonic scale.
func NewThaiScale(key string) (*Scale, error) {
	return NewBluesMajorPentatonicScale(key)
}

// NewUkrainianDorianScale builds the Ukrainian Dorian scale, an alias for the Dorian ♯4 scale.
func NewUkrainianDorianScale(key string) (*Scale, error) {
	return NewDorianSharp4Scale(key)
}

// NewVerdisAscendingFormScale builds Verdi's ascending form scale, an alias for the Enigmatic
// scale.
func NewVerdisAscendingFormScale(key string) (*Scale, error) {
	return NewEnigmaticScale(key)
}

// NewYoScale builds the Yo scale, an alias for the Blues Major Pentatonic scale.
func NewYoScale(key string) (*Scale, error) {
	return NewBluesMajorPentatonicScale(key)
}
