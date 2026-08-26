package scales_test

import (
	"testing"

	scales "github.com/andriikushch/scales/pkg"
)

func Test_HexatonicScales(t *testing.T) {
	t.Parallel()

	runScaleCases(t, []scaleCase{
		{"WholeTone", scales.NewWholeToneScale, wantNotes("C", "D", "E", "F#", "G#", "A#")},
		{"Augmented", scales.NewAugmentedScale, wantNotes("C", "Eb", "E", "G", "Ab", "B")},
		{"Blues", scales.NewBluesScale, wantNotes("C", "Eb", "F", "Gb", "G", "Bb")},
		{"MajorBlues", scales.NewMajorBluesScale, wantNotes("C", "D", "Eb", "E", "G", "A")},
		{"MajorHexatonic", scales.NewMajorHexatonicScale, wantNotes("C", "D", "E", "F", "G", "A")},
		{"MinorHexatonic", scales.NewMinorHexatonicScale, wantNotes("C", "D", "Eb", "F", "G", "Bb")},
		{"Prometheus", scales.NewPrometheusScale, wantNotes("C", "D", "E", "F#", "A", "Bb")},
		{"PrometheusNeapolitan", scales.NewPrometheusNeapolitanScale, wantNotes("C", "Db", "E", "F#", "A", "Bb")},
		{"Tritone", scales.NewTritoneScale, wantNotes("C", "Db", "E", "Gb", "G", "Bb")},
		{"Istrian", scales.NewIstrianScale, wantNotes("C", "Db", "Eb", "Fb", "Gb", "G")},
		{"ScaleOfHarmonics", scales.NewScaleOfHarmonicsScale, wantNotes("C", "Eb", "E", "F", "G", "A")},
		{"TwoSemitoneTritone", scales.NewTwoSemitoneTritoneScale, wantNotes("C", "Db", "D", "F#", "G", "Ab")},
		{"MessiaenMode5", scales.NewMessiaenMode5Scale, wantNotes("C", "Db", "F", "Gb", "G", "B")},
	})
}
