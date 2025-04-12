package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	scales "github.com/andriikushch/scales/pkg"
	"golang.org/x/term"
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

// View types
const (
	viewScale  = "scale"
	viewChords = "chords"
)

// Supported scales and instruments
var (
	validScales      = []string{scaleMajor, scaleMinor, scalePentatonic}
	validInstruments = []string{instrumentBassGuitar, instrumentGuitar, instrumentUkulele}
	minorScales      = []string{minorTypeNatural, minorTypeHarmonic, minorTypeMelodic}
	pentatonicScales = []string{pentatonicTypeMajor, pentatonicTypeMinor}
)

// Available keys for the scale
var keys = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

type scale interface {
	GetNotes() []scales.Note
	String() string
}

type keyInput struct {
	key  string
	rune rune
}

// Terminal state
type terminalState struct {
	width    int
	height   int
	oldState *term.State
}

func newTerminalState() (*terminalState, error) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return nil, err
	}

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}

	return &terminalState{
		width:    width,
		height:   height,
		oldState: oldState,
	}, nil
}

func (ts *terminalState) cleanup() {
	term.Restore(int(os.Stdin.Fd()), ts.oldState)
	fmt.Fprint(os.Stdout, "\r\n👋 Goodbye!\r\n")
	os.Stdout.Sync()
}

func (ts *terminalState) clearScreen() {
	fmt.Fprint(os.Stdout, "\033[H\033[2J")
	os.Stdout.Sync()
}

func (ts *terminalState) displayHelp(currentView string, scaleType string) {
	fmt.Fprint(os.Stdout, "Controls:\r\n")
	fmt.Fprint(os.Stdout, "← → : Change key\r\n")
	fmt.Fprint(os.Stdout, "↑ ↓ : Switch between scale and chords view\r\n")
	if currentView == viewChords && scaleType != scalePentatonic {
		fmt.Fprint(os.Stdout, "A D : Navigate between chords\r\n")
	}
	fmt.Fprint(os.Stdout, "s   : Change scale type\r\n")
	fmt.Fprint(os.Stdout, "i   : Change instrument\r\n")
	if scaleType == scalePentatonic {
		fmt.Fprint(os.Stdout, "t   : Toggle pentatonic type (major/minor)\r\n")
	} else if scaleType == scaleMinor {
		fmt.Fprint(os.Stdout, "t   : Toggle minor type (natural/harmonic/melodic)\r\n")
	}
	fmt.Fprint(os.Stdout, "q   : Quit\r\n")
	fmt.Fprint(os.Stdout, "\r\n")
	os.Stdout.Sync()
}

func (ts *terminalState) displayScaleTypes(current string) {
	fmt.Fprint(os.Stdout, "Available Scale Types:\r\n")
	fmt.Fprint(os.Stdout, "Use ↑↓ to navigate, Enter to select, or 's' to cancel\r\n")
	fmt.Fprint(os.Stdout, "┌─────────────────────┐\r\n")
	for _, s := range validScales {
		prefix := "│   "
		if s == current {
			prefix = "│ ▶ "
		}
		fmt.Fprintf(os.Stdout, "%s%-16s│\r\n", prefix, s)
	}
	fmt.Fprint(os.Stdout, "└─────────────────────┘\r\n")
	fmt.Fprint(os.Stdout, "\r\n")
	os.Stdout.Sync()
}

func (ts *terminalState) displayInstruments(current string) {
	fmt.Fprint(os.Stdout, "Available Instruments:\r\n")
	fmt.Fprint(os.Stdout, "Use ↑↓ to navigate, Enter to select, or 'i' to cancel\r\n")
	fmt.Fprint(os.Stdout, "┌─────────────────────┐\r\n")
	for _, s := range validInstruments {
		prefix := "│   "
		if s == current {
			prefix = "│ ▶ "
		}
		fmt.Fprintf(os.Stdout, "%s%-16s│\r\n", prefix, s)
	}
	fmt.Fprint(os.Stdout, "└─────────────────────┘\r\n")
	fmt.Fprint(os.Stdout, "\r\n")
	os.Stdout.Sync()
}

func (ts *terminalState) displayChords(chords []scales.Chord, currentIndex int) {
	fmt.Fprint(os.Stdout, "Available Chords:\r\n")
	for i, ch := range chords {
		prefix := "  "
		if i == currentIndex {
			prefix = "▶"
		}
		fmt.Fprintf(os.Stdout, "%s %s %s\r\n", prefix, ch.Description(), ch.Notes())
	}
	fmt.Fprint(os.Stdout, "\r\n")
	os.Stdout.Sync()
}

func main() {
	// Set up signal handling for Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Initialize terminal state
	ts, err := newTerminalState()
	if err != nil {
		log.Fatal(err)
	}
	defer ts.cleanup()

	// Define command-line flags
	scaleFlag := flag.String("scale", scaleMajor, "Specify the scale (major, minor, pentatonic)")
	keyFlag := flag.String("key", "C", "Specify the key (e.g., C, D#, F#)")
	instrumentFlag := flag.String("instrument", instrumentGuitar, "Select the instrument for visualization (guitar)")

	// Additional flags for minor and pentatonic types
	minorTypeFlag := flag.String("minorType", minorTypeNatural, "Specify the minor scale type (natural, harmonic, melodic)")
	pentatonicTypeFlag := flag.String("pentatonicType", pentatonicTypeMinor, "Specify the pentatonic scale type (major, minor)")

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
	}

	// Additional validation for pentatonic scale types
	if *scaleFlag == scalePentatonic {
		if !contains(pentatonicScales, *pentatonicTypeFlag) {
			fmt.Printf("❌ Invalid pentatonic scale type: %s. Supported: %s\n", *pentatonicTypeFlag, strings.Join(pentatonicScales, ", "))
			return
		}
	}

	// Find initial key index
	currentKeyIndex := 0
	for i, k := range keys {
		if k == *keyFlag {
			currentKeyIndex = i
			break
		}
	}

	currentView := viewScale
	var currentChordIndex int
	var showScaleTypes bool
	var showInstruments bool

	// Create a channel for keyboard input
	keyChan := make(chan keyInput, 1)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			r := rune(scanner.Bytes()[0])

			// Handle special keys
			if r == '\x1b' {
				// Read the next rune
				if !scanner.Scan() {
					continue
				}
				r2 := rune(scanner.Bytes()[0])
				if r2 != '[' {
					continue
				}

				// Read the final rune
				if !scanner.Scan() {
					continue
				}
				r3 := rune(scanner.Bytes()[0])

				switch r3 {
				case 'A':
					keyChan <- keyInput{key: "up", rune: 0}
				case 'B':
					keyChan <- keyInput{key: "down", rune: 0}
				case 'C':
					keyChan <- keyInput{key: "right", rune: 0}
				case 'D':
					keyChan <- keyInput{key: "left", rune: 0}
				}
				continue
			}

			// Handle carriage return and line feed
			if r == '\r' {
				keyChan <- keyInput{key: "", rune: '\n'}
				continue
			}

			keyChan <- keyInput{key: "", rune: r}
		}
	}()

	// Initial screen setup
	ts.clearScreen()
	ts.displayHelp(currentView, *scaleFlag)

	for {
		ts.clearScreen()
		ts.displayHelp(currentView, *scaleFlag)

		if showScaleTypes {
			ts.displayScaleTypes(*scaleFlag)
		} else if showInstruments {
			ts.displayInstruments(*instrumentFlag)
		} else {
			// Create scale based on current settings
			var s scale
			var err error

			switch *scaleFlag {
			case scaleMajor:
				s, err = scales.NewMajorScale(keys[currentKeyIndex])
			case scaleMinor:
				switch *minorTypeFlag {
				case minorTypeNatural:
					s, err = scales.NewNaturalMinorScale(keys[currentKeyIndex])
				case minorTypeHarmonic:
					s, err = scales.NewHarmonicMinorScale(keys[currentKeyIndex])
				case minorTypeMelodic:
					s, err = scales.NewMelodicMinorScale(keys[currentKeyIndex])
				}
			case scalePentatonic:
				switch *pentatonicTypeFlag {
				case pentatonicTypeMajor:
					s, err = scales.NewMajorPentatonicScale(keys[currentKeyIndex])
				case pentatonicTypeMinor:
					s, err = scales.NewMinorPentatonicScale(keys[currentKeyIndex])
				}
			}

			if err != nil {
				log.Fatal(err)
			}

			if s == nil {
				log.Fatal("something went wrong")
			}

			// Display current settings
			fmt.Fprintf(os.Stdout, "✔️ Key: %s\r\n", keys[currentKeyIndex])
			if *scaleFlag == scalePentatonic {
				fmt.Fprintf(os.Stdout, "✔️ Scale: %s (%s) : %s\r\n", *scaleFlag, *pentatonicTypeFlag, s.String())
			} else if *scaleFlag == scaleMinor {
				fmt.Fprintf(os.Stdout, "✔️ Scale: %s (%s) : %s\r\n", *scaleFlag, *minorTypeFlag, s.String())
			} else {
				fmt.Fprintf(os.Stdout, "✔️ Scale: %s : %s\r\n", *scaleFlag, s.String())
			}
			fmt.Fprintf(os.Stdout, "✔️ Instrument: %s\r\n", *instrumentFlag)
			fmt.Fprint(os.Stdout, "\r\n")
			os.Stdout.Sync()

			// Draw the current view
			switch currentView {
			case viewScale:
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
				if sc, ok := s.(*scales.Scale); ok && *scaleFlag != scalePentatonic {
					ts.displayChords(sc.GetChords(), -1)
				}
			case viewChords:
				if *scaleFlag != scalePentatonic {
					if sc, ok := s.(*scales.Scale); ok {
						chords := sc.GetChords()
						if currentChordIndex >= len(chords) {
							currentChordIndex = 0
						}
						chord := chords[currentChordIndex]
						ts.displayChords(chords, currentChordIndex)
						fmt.Fprintf(os.Stdout, "Current Chord: %s %s\r\n", chord.Description(), chord.Notes())
						fmt.Fprint(os.Stdout, "\r\n")
						switch *instrumentFlag {
						case instrumentGuitar:
							guitar := scales.NewGuitarWithStandardTuning()
							err = guitar.Draw(chord.Notes(), os.Stdout)
						case instrumentBassGuitar:
							guitar := scales.NewBassGuitarWithStandardTuning()
							err = guitar.Draw(chord.Notes(), os.Stdout)
						case instrumentUkulele:
							ukulele := scales.NewUkuleleWithStandardTuning()
							err = ukulele.Draw(chord.Notes(), os.Stdout)
						}
					}
				} else {
					fmt.Fprint(os.Stdout, "⚠️  Pentatonic scales do not have associated chords\r\n")
				}
			}

			if err != nil {
				log.Fatal(err)
			}
		}

		// Handle keyboard input and signals
		select {
		case <-sigChan:
			return
		case input := <-keyChan:
			if showScaleTypes {
				switch input.key {
				case "up":
					currentIndex := 0
					for i, s := range validScales {
						if s == *scaleFlag {
							currentIndex = i
							break
						}
					}
					*scaleFlag = validScales[(currentIndex-1+len(validScales))%len(validScales)]
				case "down":
					currentIndex := 0
					for i, s := range validScales {
						if s == *scaleFlag {
							currentIndex = i
							break
						}
					}
					*scaleFlag = validScales[(currentIndex+1)%len(validScales)]
				}
				switch input.rune {
				case '\r', '\n': // Enter key
					showScaleTypes = false
				case 's':
					showScaleTypes = false
				case 'q':
					return
				}
			} else if showInstruments {
				switch input.key {
				case "up":
					currentIndex := 0
					for i, s := range validInstruments {
						if s == *instrumentFlag {
							currentIndex = i
							break
						}
					}
					*instrumentFlag = validInstruments[(currentIndex-1+len(validInstruments))%len(validInstruments)]
				case "down":
					currentIndex := 0
					for i, s := range validInstruments {
						if s == *instrumentFlag {
							currentIndex = i
							break
						}
					}
					*instrumentFlag = validInstruments[(currentIndex+1)%len(validInstruments)]
				}
				switch input.rune {
				case '\r', '\n': // Enter key
					showInstruments = false
				case 'i':
					showInstruments = false
				case 'q':
					return
				}
			} else {
				switch input.key {
				case "left":
					currentKeyIndex = (currentKeyIndex - 1 + len(keys)) % len(keys)
				case "right":
					currentKeyIndex = (currentKeyIndex + 1) % len(keys)
				case "up":
					currentView = viewScale
				case "down":
					if *scaleFlag != scalePentatonic {
						currentView = viewChords
					}
				}

				switch input.rune {
				case 'q':
					return
				case 's':
					showScaleTypes = true
				case 'i':
					showInstruments = true
				case 't':
					if *scaleFlag == scalePentatonic {
						if *pentatonicTypeFlag == pentatonicTypeMajor {
							*pentatonicTypeFlag = pentatonicTypeMinor
						} else {
							*pentatonicTypeFlag = pentatonicTypeMajor
						}
					} else if *scaleFlag == scaleMinor {
						// Cycle through minor scale types
						currentIndex := 0
						for i, s := range minorScales {
							if s == *minorTypeFlag {
								currentIndex = i
								break
							}
						}
						*minorTypeFlag = minorScales[(currentIndex+1)%len(minorScales)]
					}
				}

				if currentView == viewChords && *scaleFlag != scalePentatonic {
					switch input.rune {
					case 'a', 'A':
						currentChordIndex = (currentChordIndex - 1 + 7) % 7
					case 'd', 'D':
						currentChordIndex = (currentChordIndex + 1) % 7
					}
				}
			}
		}
	}
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
