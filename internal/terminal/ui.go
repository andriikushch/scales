package terminal

import (
	scales "github.com/andriikushch/scales/pkg"
)

// DisplayHelp shows the help menu
func (ts *State) DisplayHelp(currentView string, scaleType string) {
	ts.Print("Controls:\r\n")
	ts.Print("← → : Change key\r\n")
	ts.Print("↑ ↓ : Switch between scale and chords view\r\n")
	if currentView == "chords" && scaleType != "pentatonic" {
		ts.Print("A D : Navigate between chords\r\n")
	}
	ts.Print("s   : Change scale type\r\n")
	ts.Print("i   : Change instrument\r\n")
	if scaleType == "pentatonic" {
		ts.Print("t   : Toggle pentatonic type (major/minor)\r\n")
	} else if scaleType == "minor" {
		ts.Print("t   : Toggle minor type (natural/harmonic/melodic)\r\n")
	}
	ts.Print("q   : Quit\r\n")
	ts.Print("\r\n")
}

// DisplayScaleTypes shows the scale type selection menu
func (ts *State) DisplayScaleTypes(current string, validScales []string) {
	ts.Print("Available Scale Types:\r\n")
	ts.Print("Use ↑↓ to navigate, Enter to select, or 's' to cancel\r\n")
	ts.Print("┌─────────────────────┐\r\n")
	for _, s := range validScales {
		prefix := "│   "
		if s == current {
			prefix = "│ ▶ "
		}
		ts.Print("%s%-16s│\r\n", prefix, s)
	}
	ts.Print("└─────────────────────┘\r\n")
	ts.Print("\r\n")
}

// DisplayInstruments shows the instrument selection menu
func (ts *State) DisplayInstruments(current string, validInstruments []string) {
	ts.Print("Available Instruments:\r\n")
	ts.Print("Use ↑↓ to navigate, Enter to select, or 'i' to cancel\r\n")
	ts.Print("┌─────────────────────┐\r\n")
	for _, s := range validInstruments {
		prefix := "│   "
		if s == current {
			prefix = "│ ▶ "
		}
		ts.Print("%s%-16s│\r\n", prefix, s)
	}
	ts.Print("└─────────────────────┘\r\n")
	ts.Print("\r\n")
}

// DisplayChords shows the chord list
func (ts *State) DisplayChords(chords []scales.Chord, currentIndex int) {
	ts.Print("Available Chords:\r\n")
	for i, ch := range chords {
		prefix := "  "
		if i == currentIndex {
			prefix = "▶"
		}
		ts.Print("%s %s %s\r\n", prefix, ch.Description(), ch.Notes())
	}
	ts.Print("\r\n")
}

// DisplaySettings shows the current settings
func (ts *State) DisplaySettings(key, scaleType, scaleTypeDetail, scale, instrument string) {
	ts.Print("✔️ Key: %s\r\n", key)
	if scaleType == "pentatonic" {
		ts.Print("✔️ Scale: %s (%s) : %s\r\n", scaleType, scaleTypeDetail, scale)
	} else if scaleType == "minor" {
		ts.Print("✔️ Scale: %s (%s) : %s\r\n", scaleType, scaleTypeDetail, scale)
	} else {
		ts.Print("✔️ Scale: %s : %s\r\n", scaleType, scale)
	}
	ts.Print("✔️ Instrument: %s\r\n", instrument)
	ts.Print("\r\n")
}
