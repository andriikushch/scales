package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/andriikushch/scales/internal/app"
	"github.com/andriikushch/scales/internal/terminal"
	scales "github.com/andriikushch/scales/pkg"
)

func main() {
	// Set up signal handling for Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Initialize terminal state
	ts, err := terminal.New()
	if err != nil {
		log.Fatal(err)
	}
	defer ts.Cleanup()

	// Define command-line flags
	scaleFlag := flag.String("scale", app.ScaleMajor, "Specify the scale (major, minor, pentatonic)")
	keyFlag := flag.String("key", "C", "Specify the key (e.g., C, D#, F#)")
	instrumentFlag := flag.String("instrument", app.InstrumentGuitar, "Select the instrument for visualization (guitar)")

	// Additional flags for minor and pentatonic types
	minorTypeFlag := flag.String("minorType", app.MinorTypeNatural, "Specify the minor scale type (natural, harmonic, melodic)")
	pentatonicTypeFlag := flag.String("pentatonicType", app.PentatonicTypeMinor, "Specify the pentatonic scale type (major, minor)")

	flag.Parse()

	// Initialize application state
	state := app.New()
	state.ScaleType = *scaleFlag
	state.ScaleTypeDetail = *minorTypeFlag
	if *scaleFlag == app.ScalePentatonic {
		state.ScaleTypeDetail = *pentatonicTypeFlag
	}
	state.Instrument = *instrumentFlag

	// Find initial key index
	for i, k := range state.Keys {
		if k == *keyFlag {
			state.CurrentKeyIndex = i
			break
		}
	}

	// Initialize input handler
	inputHandler := app.NewInputHandler()

	// Initial screen setup
	ts.ClearScreen()
	ts.DisplayHelp(state.CurrentView, state.ScaleType)

	for {
		ts.ClearScreen()
		ts.DisplayHelp(state.CurrentView, state.ScaleType)

		if state.ShowScaleTypes {
			ts.DisplayScaleTypes(state.ScaleType, state.ValidScales)
		} else if state.ShowInstruments {
			ts.DisplayInstruments(state.Instrument, state.ValidInstruments)
		} else {
			// Create scale based on current settings
			if err := state.CreateScale(); err != nil {
				log.Fatal(err)
			}

			// Display current settings
			ts.DisplaySettings(
				state.Keys[state.CurrentKeyIndex],
				state.ScaleType,
				state.ScaleTypeDetail,
				state.Scale.String(),
				state.Instrument,
			)

			// Always draw the scale first
			switch state.Instrument {
			case app.InstrumentGuitar:
				guitar := scales.NewGuitarWithStandardTuning()
				err = guitar.Draw(state.GetNotes(), os.Stdout)
			case app.InstrumentBassGuitar:
				guitar := scales.NewBassGuitarWithStandardTuning()
				err = guitar.Draw(state.GetNotes(), os.Stdout)
			case app.InstrumentUkulele:
				ukulele := scales.NewUkuleleWithStandardTuning()
				err = ukulele.Draw(state.GetNotes(), os.Stdout)
			}
			ts.Print("\r\n")

			// Draw the current view
			switch state.CurrentView {
			case app.ViewScale:
				if chords := state.GetChords(); chords != nil {
					ts.DisplayChords(chords, -1)
				}
			case app.ViewChords:
				if state.ScaleType != app.ScalePentatonic {
					if chords := state.GetChords(); chords != nil {
						ts.DisplayChords(chords, state.CurrentChordIndex)
						chord := state.GetCurrentChord()
						ts.Print("Current Chord: %s %s\r\n", chord.Description(), chord.Notes())
						ts.Print("\r\n")
						// Draw the current chord
						switch state.Instrument {
						case app.InstrumentGuitar:
							guitar := scales.NewGuitarWithStandardTuning()
							err = guitar.Draw(chord.Notes(), os.Stdout)
						case app.InstrumentBassGuitar:
							guitar := scales.NewBassGuitarWithStandardTuning()
							err = guitar.Draw(chord.Notes(), os.Stdout)
						case app.InstrumentUkulele:
							ukulele := scales.NewUkuleleWithStandardTuning()
							err = ukulele.Draw(chord.Notes(), os.Stdout)
						}
					}
				} else {
					ts.Print("⚠️  Pentatonic scales do not have associated chords\r\n")
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
		case input := <-inputHandler.GetInput():
			if state.ShowScaleTypes {
				switch input.Key {
				case "up":
					currentIndex := 0
					for i, s := range state.ValidScales {
						if s == state.ScaleType {
							currentIndex = i
							break
						}
					}
					state.ScaleType = state.ValidScales[(currentIndex-1+len(state.ValidScales))%len(state.ValidScales)]
				case "down":
					currentIndex := 0
					for i, s := range state.ValidScales {
						if s == state.ScaleType {
							currentIndex = i
							break
						}
					}
					state.ScaleType = state.ValidScales[(currentIndex+1)%len(state.ValidScales)]
				}
				switch input.Rune {
				case '\r', '\n': // Enter key
					state.ShowScaleTypes = false
				case 's':
					state.ShowScaleTypes = false
				case 'q':
					return
				}
			} else if state.ShowInstruments {
				switch input.Key {
				case "up":
					currentIndex := 0
					for i, s := range state.ValidInstruments {
						if s == state.Instrument {
							currentIndex = i
							break
						}
					}
					state.Instrument = state.ValidInstruments[(currentIndex-1+len(state.ValidInstruments))%len(state.ValidInstruments)]
				case "down":
					currentIndex := 0
					for i, s := range state.ValidInstruments {
						if s == state.Instrument {
							currentIndex = i
							break
						}
					}
					state.Instrument = state.ValidInstruments[(currentIndex+1)%len(state.ValidInstruments)]
				}
				switch input.Rune {
				case '\r', '\n': // Enter key
					state.ShowInstruments = false
				case 'i':
					state.ShowInstruments = false
				case 'q':
					return
				}
			} else {
				switch input.Key {
				case "left":
					state.CurrentKeyIndex = (state.CurrentKeyIndex - 1 + len(state.Keys)) % len(state.Keys)
				case "right":
					state.CurrentKeyIndex = (state.CurrentKeyIndex + 1) % len(state.Keys)
				case "up":
					state.CurrentView = app.ViewScale
				case "down":
					if state.ScaleType != app.ScalePentatonic {
						state.CurrentView = app.ViewChords
					}
				}

				switch input.Rune {
				case 'q':
					return
				case 's':
					state.ShowScaleTypes = true
				case 'i':
					state.ShowInstruments = true
				case 't':
					if state.ScaleType == app.ScalePentatonic {
						if state.ScaleTypeDetail == app.PentatonicTypeMajor {
							state.ScaleTypeDetail = app.PentatonicTypeMinor
						} else {
							state.ScaleTypeDetail = app.PentatonicTypeMajor
						}
					} else if state.ScaleType == app.ScaleMinor {
						// Cycle through minor scale types
						currentIndex := 0
						for i, s := range state.MinorScales {
							if s == state.ScaleTypeDetail {
								currentIndex = i
								break
							}
						}
						state.ScaleTypeDetail = state.MinorScales[(currentIndex+1)%len(state.MinorScales)]
					}
				}

				if state.CurrentView == app.ViewChords && state.ScaleType != app.ScalePentatonic {
					switch input.Rune {
					case 'a', 'A':
						state.CurrentChordIndex = (state.CurrentChordIndex - 1 + 7) % 7
					case 'd', 'D':
						state.CurrentChordIndex = (state.CurrentChordIndex + 1) % 7
					}
				}
			}
		}
	}
}
