package scales

import (
	"fmt"
	"slices"
)

// scaleRegistry maps every scale in docs/scales.md to its constructor, under a
// lowerCamelCase name derived from the scale's name in that doc (accidentals spelled out,
// e.g. "Dorian ♭2" -> "dorianFlat2", "Locrian ♭♭7" -> "locrianDoubleFlat7").
var scaleRegistry = map[string]func(string) (*Scale, error){
	// Modes of the major scale
	"ionian":     NewMajorScale,
	"dorian":     NewDorianScale,
	"phrygian":   NewPhrygianScale,
	"lydian":     NewLydianScale,
	"mixolydian": NewMixolydianScale,
	"aeolian":    NewNaturalMinorScale,
	"locrian":    NewLocrianScale,

	// Modes of the melodic minor scale
	"melodicMinor":    NewMelodicMinorScale,
	"dorianFlat2":     NewDorianFlat2Scale,
	"lydianAugmented": NewLydianAugmentedScale,
	"lydianDominant":  NewLydianDominantScale,
	"mixolydianFlat6": NewMixolydianFlat6Scale,
	"locrianNatural2": NewLocrianNatural2Scale,
	"altered":         NewAlteredScale,

	// Modes of the harmonic minor scale
	"harmonicMinor":    NewHarmonicMinorScale,
	"locrianNatural6":  NewLocrianNatural6Scale,
	"ionianSharp5":     NewIonianSharp5Scale,
	"dorianSharp4":     NewDorianSharp4Scale,
	"phrygianDominant": NewPhrygianDominantScale,
	"lydianSharp2":     NewLydianSharp2Scale,
	"ultralocrian":     NewUltralocrianScale,

	// Modes of the harmonic major scale
	"harmonicMajor":         NewHarmonicMajorScale,
	"dorianFlat5":           NewDorianFlat5Scale,
	"phrygianFlat4":         NewPhrygianFlat4Scale,
	"lydianFlat3":           NewLydianFlat3Scale,
	"mixolydianFlat2":       NewMixolydianFlat2Scale,
	"lydianAugmentedSharp2": NewLydianAugmentedSharp2Scale,
	"locrianDoubleFlat7":    NewLocrianDoubleFlat7Scale,

	// Modes of the double harmonic major scale
	"doubleHarmonicMajor":           NewDoubleHarmonicMajorScale,
	"lydianSharp2Sharp6":            NewLydianSharp2Sharp6Scale,
	"ultraphrygian":                 NewUltraphrygianScale,
	"hungarianMinor":                NewHungarianMinorScale,
	"oriental":                      NewOrientalScale,
	"ionianSharp2Sharp5":            NewIonianSharp2Sharp5Scale,
	"locrianDoubleFlat3DoubleFlat7": NewLocrianDoubleFlat3DoubleFlat7Scale,

	// Other heptatonic scales
	"neapolitanMinor":  NewNeapolitanMinorScale,
	"neapolitanMajor":  NewNeapolitanMajorScale,
	"leadingWholeTone": NewLeadingWholeToneScale,
	"majorLocrian":     NewMajorLocrianScale,
	"gypsy":            NewGypsyScale,
	"hungarianMajor":   NewHungarianMajorScale,
	"persian":          NewPersianScale,
	"enigmatic":        NewEnigmaticScale,

	// Pentatonic scales
	"majorPentatonic":      func(key string) (*Scale, error) { return pentatonicAsScale(NewMajorPentatonicScale, key) },
	"suspendedPentatonic":  NewSuspendedPentatonicScale,
	"bluesMinorPentatonic": NewBluesMinorPentatonicScale,
	"bluesMajorPentatonic": NewBluesMajorPentatonicScale,
	"minorPentatonic":      func(key string) (*Scale, error) { return pentatonicAsScale(NewMinorPentatonicScale, key) },
	"iwato":                NewIwatoScale,
	"miyakoBushi":          NewMiyakoBushiScale,
	"kumoi":                NewKumoiScale,
	"insen":                NewInsenScale,
	"chinesePentatonic":    NewChinesePentatonicScale,
	"balinesePelog":        NewBalinesePelogScale,

	// Hexatonic scales
	"wholeTone":            NewWholeToneScale,
	"augmented":            NewAugmentedScale,
	"blues":                NewBluesScale,
	"majorBlues":           NewMajorBluesScale,
	"majorHexatonic":       NewMajorHexatonicScale,
	"minorHexatonic":       NewMinorHexatonicScale,
	"prometheus":           NewPrometheusScale,
	"prometheusNeapolitan": NewPrometheusNeapolitanScale,
	"tritone":              NewTritoneScale,
	"istrian":              NewIstrianScale,
	"scaleOfHarmonics":     NewScaleOfHarmonicsScale,
	"twoSemitoneTritone":   NewTwoSemitoneTritoneScale,
	"messiaenMode5":        NewMessiaenMode5Scale,

	// Octatonic scales
	"halfWholeDiminished": NewHalfWholeDiminishedScale,
	"wholeHalfDiminished": NewWholeHalfDiminishedScale,
	"messiaenMode4":       NewMessiaenMode4Scale,
	"messiaenMode6":       NewMessiaenMode6Scale,
	"spanishEightTone":    NewSpanishEightToneScale,

	// Bebop scales
	"bebopDominant":      NewBebopDominantScale,
	"bebopMajor":         NewBebopMajorScale,
	"bebopDorian":        NewBebopDorianScale,
	"bebopMelodicMinor":  NewBebopMelodicMinorScale,
	"bebopHarmonicMinor": NewBebopHarmonicMinorScale,

	// Nine or more notes
	"messiaenMode3": NewMessiaenMode3Scale,
	"messiaenMode7": NewMessiaenMode7Scale,
	"chromatic":     NewChromaticScale,

	// Alternative names (see docs/scales.md "Other names")
	"acoustic":                   NewAcousticScale,
	"adonaiMalakh":               NewAdonaiMalakhScale,
	"aeolianDominant":            NewAeolianDominantScale,
	"aeolianFlat5":               NewAeolianFlat5Scale,
	"ahavaRabba":                 NewAhavaRabbaScale,
	"alteredDoubleFlat7":         NewAlteredDoubleFlat7Scale,
	"alteredNatural5":            NewAlteredNatural5Scale,
	"andalusian":                 NewAndalusianScale,
	"arabic":                     NewArabicScale,
	"augmentedMajor":             NewAugmentedMajorScale,
	"byzantine":                  NewByzantineScale,
	"chinese":                    NewChineseScale,
	"dominantDiminished":         NewDominantDiminishedScale,
	"dorianAeolianWithoutThe6th": NewDorianAeolianWithoutThe6thScale,
	"doubleHarmonicMinor":        NewDoubleHarmonicMinorScale,
	"freygish":                   NewFreygishScale,
	"gypsyMajor":                 NewGypsyMajorScale,
	"gypsyMinor":                 NewGypsyMinorScale,
	"halfDiminished":             NewHalfDiminishedScale,
	"hawaiian":                   NewHawaiianScale,
	"hijaz":                      NewHijazScale,
	"hindu":                      NewHinduScale,
	"ionianFlat6":                NewIonianFlat6Scale,
	"jazzMinor":                  NewJazzMinorScale,
	"korean":                     NewKoreanScale,
	"locrianDiminished":          NewLocrianDiminishedScale,
	"locrianNatural2Natural6":    NewLocrianNatural2Natural6Scale,
	"lydianDiminished":           NewLydianDiminishedScale,
	"lydianFlat7":                NewLydianFlat7Scale,
	"major":                      NewMajorScale,
	"majorScaleWithoutThe7th":    NewMajorScaleWithoutThe7thScale,
	"manGong":                    NewManGongScale,
	"mayamalavagowla":            NewMayamalavagowlaScale,
	"melodicMinorFlat2":          NewMelodicMinorFlat2Scale,
	"melodicMinorSharp4":         NewMelodicMinorSharp4Scale,
	"misheberak":                 NewMisheberakScale,
	"mixolydianSharp4":           NewMixolydianSharp4Scale,
	"mongolian":                  NewMongolianScale,
	"mysticChord":                NewMysticChordScale,
	"naturalMinor":               NewNaturalMinorScale,
	"overtone":                   NewOvertoneScale,
	"phrygianNatural6":           NewPhrygianNatural6Scale,
	"prometheusLiszt":            NewPrometheusLisztScale,
	"ragaBhairav":                NewRagaBhairavScale,
	"ragaBhupali":                NewRagaBhupaliScale,
	"ragaDurga":                  NewRagaDurgaScale,
	"ragaMalkauns":               NewRagaMalkaunsScale,
	"ritsu":                      NewRitsuScale,
	"ritsusen":                   NewRitsusenScale,
	"romanianMinor":              NewRomanianMinorScale,
	"sakura":                     NewSakuraScale,
	"slendro":                    NewSlendroScale,
	"spanishGypsy":               NewSpanishGypsyScale,
	"superLocrian":               NewSuperLocrianScale,
	"superlocrianDoubleFlat7":    NewSuperlocrianDoubleFlat7Scale,
	"thai":                       NewThaiScale,
	"ukrainianDorian":            NewUkrainianDorianScale,
	"verdisAscendingForm":        NewVerdisAscendingFormScale,
	"yo":                         NewYoScale,
}

func pentatonicAsScale(fn func(string) (*Pentatonic, error), key string) (*Scale, error) {
	p, err := fn(key)
	if err != nil {
		return nil, err
	}

	return &Scale{notes: p.notes}, nil
}

// ScaleNames returns the sorted list of scale names accepted by NewScaleByName.
func ScaleNames() []string {
	names := make([]string, 0, len(scaleRegistry))
	for name := range scaleRegistry {
		names = append(names, name)
	}

	slices.Sort(names)

	return names
}

// NewScaleByName builds any scale documented in docs/scales.md by its registry name (see
// ScaleNames) in the given key.
func NewScaleByName(name string, key string) (*Scale, error) {
	fn, ok := scaleRegistry[name]
	if !ok {
		return nil, fmt.Errorf("unknown scale name %q", name)
	}

	return fn(key)
}
