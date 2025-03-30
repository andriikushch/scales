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
	scaleMajor      = "major"
	scaleMinor      = "minor"
	scalePentatonic = "pentatonic"
)

// Instrument types
const (
	instrumentBassGuitar = "bassGuitar"
	instrumentGuitar     = "guitar"
	instrumentUkulele    = "ukulele"
)

// Minor scale types
const (
	minorTypeNatural  = "natural"
	minorTypeHarmonic = "harmonic"
	minorTypeMelodic  = "melodic"
)

// Pentatonic scale types
const (
	pentatonicTypeMajor = "major"
	pentatonicTypeMinor = "minor"
)

// Supported scales and instruments
var (
	validScales      = []string{scaleMajor, scaleMinor, scalePentatonic}
	validInstruments = []string{instrumentBassGuitar, instrumentGuitar, instrumentUkulele}
	minorScales      = []string{minorTypeNatural, minorTypeHarmonic, minorTypeMelodic}
	pentatonicScales = []string{pentatonicTypeMajor, pentatonicTypeMinor}
)

type scale interface {
	GetNotes() []scales.Note
	String() string
}

func main() {
	// Define command-line flags
	scaleFlag := flag.String("scale", scaleMajor, "Specify the scale (major, minor, pentatonic)")
	keyFlag := flag.String("key", "C", "Specify the key (e.g., C, D#, F#)")
	instrumentFlag := flag.String("instrument", instrumentGuitar, "Select the instrument for visualization (guitar)")

	// Additional flags for minor and pentatonic types
	minorTypeFlag := flag.String("minorType", minorTypeNatural, "Specify the minor scale type (natural, harmonic, melodic)")
	pentatonicTypeFlag := flag.String("pentatonicType", pentatonicTypeMinor, "Specify the pentatonic scale type (major, minor, blues)")

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
	if *scaleFlag == scaleMinor {
		if !contains(minorScales, *minorTypeFlag) {
			fmt.Printf("❌ Invalid minor scale type: %s. Supported: %s\n", *minorTypeFlag, strings.Join(minorScales, ", "))
			return
		}
		fmt.Printf("✔️ Minor Scale Type: %s\n", *minorTypeFlag)
	}

	// Additional validation for pentatonic scale types
	if *scaleFlag == scalePentatonic {
		if !contains(pentatonicScales, *pentatonicTypeFlag) {
			fmt.Printf("❌ Invalid pentatonic scale type: %s. Supported: %s\n", *pentatonicTypeFlag, strings.Join(pentatonicScales, ", "))
			return
		}
		fmt.Printf("✔️ Pentatonic Scale Type: %s\n", *pentatonicTypeFlag)
	}

	var s scale
	var err error

	switch *scaleFlag {
	case scaleMajor:
		s, err = scales.NewMajorScale(*keyFlag)
	case scaleMinor:
		switch *minorTypeFlag {
		case minorTypeNatural:
			s, err = scales.NewNaturalMinorScale(*keyFlag)
		case minorTypeHarmonic:
			s, err = scales.NewHarmonicMinorScale(*keyFlag)
		case minorTypeMelodic:
			s, err = scales.NewMelodicMinorScale(*keyFlag)
		}
	case scalePentatonic:
		switch *pentatonicTypeFlag {
		case pentatonicTypeMajor:
			s, err = scales.NewMajorPentatonicScale(*keyFlag)
		case pentatonicTypeMinor:
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
	case instrumentGuitar:
		guitar := scales.NewGuitarWithStandardTuning()
		err = guitar.Draw(s.GetNotes(), os.Stdout)
	case instrumentBassGuitar:
		guitar := scales.NewBassGuitarWithStandardTuning()
		err = guitar.Draw(s.GetNotes(), os.Stdout)
	case instrumentUkulele:
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
