package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	scales "github.com/andriikushch/scales/pkg"
)

// Supported scales and instruments
var (
	validScales      = []string{"major", "minor", "pentatonic"}
	validInstruments = []string{"bass", "guitar"}
	minorScales      = []string{"natural", "harmonic", "melodic"}
	pentatonicScales = []string{"major", "minor"}
)

func main() {
	// Define command-line flags
	scaleFlag := flag.String("scale", "major", "Specify the scale (major, minor, pentatonic)")
	keyFlag := flag.String("key", "C", "Specify the key (e.g., C, D#, F#)")
	instrumentFlag := flag.String("instrument", "guitar", "Select the instrument for visualization (guitar)")

	// Additional flags for minor and pentatonic types
	minorTypeFlag := flag.String("minorType", "natural", "Specify the minor scale type (natural, harmonic, melodic)")
	pentatonicTypeFlag := flag.String("pentatonicType", "minor", "Specify the pentatonic scale type (major, minor, blues)")

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
	if *scaleFlag == "minor" {
		if !contains(minorScales, *minorTypeFlag) {
			fmt.Printf("❌ Invalid minor scale type: %s. Supported: %s\n", *minorTypeFlag, strings.Join(minorScales, ", "))
			return
		}
		fmt.Printf("✔️ Minor Scale Type: %s\n", *minorTypeFlag)
	}

	// Additional validation for pentatonic scale types
	if *scaleFlag == "pentatonic" {
		if !contains(pentatonicScales, *pentatonicTypeFlag) {
			fmt.Printf("❌ Invalid pentatonic scale type: %s. Supported: %s\n", *pentatonicTypeFlag, strings.Join(pentatonicScales, ", "))
			return
		}
		fmt.Printf("✔️ Pentatonic Scale Type: %s\n", *pentatonicTypeFlag)
	}

	var scale *scales.Scale
	var err error

	switch *scaleFlag {
	case "major":
		scale, err = scales.NewMajorScale(*keyFlag)
	case "minor":
		switch *minorTypeFlag {
		case "natural":
			scale, err = scales.NewNaturalMinorScale(*keyFlag)
		case "harmonic":
			scale, err = scales.NewHarmonicMinorScale(*keyFlag)
		case "melodic":
			scale, err = scales.NewMelodicMinorScale(*keyFlag)
		}
	case "pentatonic":
		switch *pentatonicTypeFlag {
		case "major":
			scale, err = scales.NewMajorPentatonicScale(*keyFlag)
		case "minor":
			scale, err = scales.NewMinorPentatonicScale(*keyFlag)
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	if scale == nil {
		log.Fatal("something went wrong")
	}

	// Display user selection
	fmt.Println("🎵 Scale Generator 🎵")
	fmt.Printf("✔️ Key: %s\n", *keyFlag)
	fmt.Printf("✔️ Scale: %s : %s\n", *scaleFlag, scale.String())
	fmt.Printf("✔️ Instrument: %s\n", *instrumentFlag)

	switch *instrumentFlag {
	case "guitar":
		guitar := scales.NewGuitarWithStandardTuning()
		err = guitar.Draw(scale.GetNotes(), os.Stdout)
	case "bass":
		guitar := scales.NewBassWithStandardTuning()
		err = guitar.Draw(scale.GetNotes(), os.Stdout)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
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
