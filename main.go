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
	scaleFlag := flag.String("scale", "major", "Specify the scale name (any name from scales.ScaleNames(), e.g. major, minorPentatonic, harmonicMinor, freygish)")
	keyFlag := flag.String("key", "C", "Specify the key (e.g., C, D#, F#)")
	instrumentFlag := flag.String("instrument", app.InstrumentGuitar, "Select the instrument for visualization (guitar)")

	flag.Parse()

	// Initialize application state
	state := app.New()
	state.ScaleName = *scaleFlag
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
	ts.DisplayHelp()

	for {
		ts.ClearScreen()
		ts.DisplayHelp()

		if state.ShowScaleTypes {
			ts.DisplayScaleTypes(state.ScaleName, state.ScaleNames)
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
				state.ScaleName,
				state.Scale.String(),
				state.Instrument,
			)

			// Draw the scale on the selected instrument
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
			case app.InstrumentMandolin:
				mandolin := scales.NewMandolinWithStandardTuning()
				err = mandolin.Draw(state.GetNotes(), os.Stdout)
			}
			ts.Print("\r\n")

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
					for i, s := range state.ScaleNames {
						if s == state.ScaleName {
							currentIndex = i
							break
						}
					}
					state.ScaleName = state.ScaleNames[(currentIndex-1+len(state.ScaleNames))%len(state.ScaleNames)]
				case "down":
					currentIndex := 0
					for i, s := range state.ScaleNames {
						if s == state.ScaleName {
							currentIndex = i
							break
						}
					}
					state.ScaleName = state.ScaleNames[(currentIndex+1)%len(state.ScaleNames)]
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
				}

				switch input.Rune {
				case 'q':
					return
				case 's':
					state.ShowScaleTypes = true
				case 'i':
					state.ShowInstruments = true
				}
			}
		}
	}
}
