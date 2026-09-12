package terminal

import "strings"

// maxVisibleListItems caps how many rows a selection box renders at once; longer lists scroll,
// windowed around the current selection.
const maxVisibleListItems = 15

// DisplayHelp shows the help menu
func (ts *State) DisplayHelp() {
	ts.Print("Controls:\r\n")
	ts.Print("← → : Change key\r\n")
	ts.Print("s   : Change scale\r\n")
	ts.Print("i   : Change instrument\r\n")
	ts.Print("q   : Quit\r\n")
	ts.Print("\r\n")
}

// displayList renders a bordered selection box for validOptions, highlighting current and
// scrolling to keep it visible when the list is longer than maxVisibleListItems.
func (ts *State) displayList(title, cancelKey, current string, validOptions []string) {
	currentIndex := 0
	maxLen := 0
	for i, o := range validOptions {
		if o == current {
			currentIndex = i
		}
		if len(o) > maxLen {
			maxLen = len(o)
		}
	}

	start, end := 0, len(validOptions)
	if end > maxVisibleListItems {
		start = currentIndex - maxVisibleListItems/2
		if start < 0 {
			start = 0
		}
		end = start + maxVisibleListItems
		if end > len(validOptions) {
			end = len(validOptions)
			start = end - maxVisibleListItems
		}
	}

	border := strings.Repeat("─", maxLen+4)

	ts.Print("Available %s:\r\n", title)
	ts.Print("Use ↑↓ to navigate, Enter to select, or '%s' to cancel\r\n", cancelKey)
	if start > 0 {
		ts.Print("┌%s┐ (%d more above)\r\n", border, start)
	} else {
		ts.Print("┌%s┐\r\n", border)
	}
	for i := start; i < end; i++ {
		prefix := "│   "
		if i == currentIndex {
			prefix = "│ ▶ "
		}
		ts.Print("%s%-*s│\r\n", prefix, maxLen+1, validOptions[i])
	}
	if remaining := len(validOptions) - end; remaining > 0 {
		ts.Print("└%s┘ (%d more below)\r\n", border, remaining)
	} else {
		ts.Print("└%s┘\r\n", border)
	}
	ts.Print("\r\n")
}

// DisplayScaleTypes shows the scale selection menu
func (ts *State) DisplayScaleTypes(current string, validScales []string) {
	ts.displayList("Scales", "s", current, validScales)
}

// DisplayInstruments shows the instrument selection menu
func (ts *State) DisplayInstruments(current string, validInstruments []string) {
	ts.displayList("Instruments", "i", current, validInstruments)
}

// DisplaySettings shows the current settings
func (ts *State) DisplaySettings(key, scaleName, scale, instrument string) {
	ts.Print("✔️ Key: %s\r\n", key)
	ts.Print("✔️ Scale: %s : %s\r\n", scaleName, scale)
	ts.Print("✔️ Instrument: %s\r\n", instrument)
	ts.Print("\r\n")
}
