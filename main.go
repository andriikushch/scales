package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	scales "github.com/andriikushch/scales/pkg"
)

// Scale types
const (
	ScaleMajor      = "major"
	ScaleMinor      = "minor"
	ScalePentatonic = "pentatonic"
)

// Instrument types
const (
	InstrumentBassGuitar = "bassGuitar"
	InstrumentGuitar     = "guitar"
	InstrumentUkulele    = "ukulele"
)

// Minor scale types
const (
	MinorTypeNatural  = "natural"
	MinorTypeHarmonic = "harmonic"
	MinorTypeMelodic  = "melodic"
)

// Pentatonic scale types
const (
	PentatonicTypeMajor = "major"
	PentatonicTypeMinor = "minor"
)

// Supported scales and instruments
var (
	validScales      = []string{ScaleMajor, ScaleMinor, ScalePentatonic}
	validInstruments = []string{InstrumentBassGuitar, InstrumentGuitar, InstrumentUkulele}
	minorScales      = []string{MinorTypeNatural, MinorTypeHarmonic, MinorTypeMelodic}
	pentatonicScales = []string{PentatonicTypeMajor, PentatonicTypeMinor}
)

type scale interface {
	GetNotes() []scales.Note
	String() string
}

func main() {
	// Define command-line flags
	scaleFlag := flag.String("scale", ScaleMajor, "Specify the scale (major, minor, pentatonic)")
	keyFlag := flag.String("key", "C", "Specify the key (e.g., C, D#, F#)")
	instrumentFlag := flag.String("instrument", InstrumentGuitar, "Select the instrument for visualization (guitar)")

	// Additional flags for minor and pentatonic types
	minorTypeFlag := flag.String("minorType", MinorTypeNatural, "Specify the minor scale type (natural, harmonic, melodic)")
	pentatonicTypeFlag := flag.String("pentatonicType", PentatonicTypeMinor, "Specify the pentatonic scale type (major, minor, blues)")

	flag.Parse() // Parse command-line flags

	// Validate scale selection
	if !contains(validScales, *scaleFlag) {
		fmt.Printf("❌ Invalid scale: %s. Supported: %s\n", *scaleFlag, strings.Join(validScales, ", "))
		return
	}

	// Validate instrument selection
	if !contains(validInstruments, *instrumentFlag) {
		fmt.Printf("❌ Invalid instrument: %s. Supported: %s\n", *instrumentFlag, strings.Join(validInstruments, ", "))
		return
	}

	// Additional validation for minor scale types
	if *scaleFlag == ScaleMinor {
		if !contains(minorScales, *minorTypeFlag) {
			fmt.Printf("❌ Invalid minor scale type: %s. Supported: %s\n", *minorTypeFlag, strings.Join(minorScales, ", "))
			return
		}
		fmt.Printf("✔️ Minor Scale Type: %s\n", *minorTypeFlag)
	}

	// Additional validation for pentatonic scale types
	if *scaleFlag == ScalePentatonic {
		if !contains(pentatonicScales, *pentatonicTypeFlag) {
			fmt.Printf("❌ Invalid pentatonic scale type: %s. Supported: %s\n", *pentatonicTypeFlag, strings.Join(pentatonicScales, ", "))
			return
		}
		fmt.Printf("✔️ Pentatonic Scale Type: %s\n", *pentatonicTypeFlag)
	}

	var s scale
	var err error

	switch *scaleFlag {
	case ScaleMajor:
		s, err = scales.NewMajorScale(*keyFlag)
	case ScaleMinor:
		switch *minorTypeFlag {
		case MinorTypeNatural:
			s, err = scales.NewNaturalMinorScale(*keyFlag)
		case MinorTypeHarmonic:
			s, err = scales.NewHarmonicMinorScale(*keyFlag)
		case MinorTypeMelodic:
			s, err = scales.NewMelodicMinorScale(*keyFlag)
		}
	case ScalePentatonic:
		switch *pentatonicTypeFlag {
		case PentatonicTypeMajor:
			s, err = scales.NewMajorPentatonicScale(*keyFlag)
		case PentatonicTypeMinor:
			s, err = scales.NewMinorPentatonicScale(*keyFlag)
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	if s == nil {
		log.Fatal("something went wrong")
	}

	// Display user selection
	fmt.Println("🎵 Scale Generator 🎵")
	fmt.Printf("✔️ Key: %s\n", *keyFlag)
	fmt.Printf("✔️ Scale: %s : %s\n", *scaleFlag, s.String())
	fmt.Printf("✔️ Instrument: %s\n", *instrumentFlag)

	switch *instrumentFlag {
	case InstrumentGuitar:
		guitar := scales.NewGuitarWithStandardTuning()
		err = guitar.Draw(s.GetNotes(), os.Stdout)
	case InstrumentBassGuitar:
		guitar := scales.NewBassGuitarWithStandardTuning()
		err = guitar.Draw(s.GetNotes(), os.Stdout)
	case InstrumentUkulele:
		ukulele := scales.NewUkuleleWithStandardTuning()
		err = ukulele.Draw(s.GetNotes(), os.Stdout)
	}

	if err != nil {
		log.Fatal(err)
	}

	if sc, ok := s.(*scales.Scale); ok {
		chords := ""
		for _, ch := range sc.GetChords() {
			chords += fmt.Sprintf("%s %s, ", ch.Description(), ch.Notes())
		}
		fmt.Printf("✔️ Chords: %s\n", strings.TrimSuffix(chords, ", "))
	}

	return
}

// Helper function to check if a value is in a slice
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
