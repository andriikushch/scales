package terminal

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

// State represents the terminal state and operations
type State struct {
	width    int
	height   int
	oldState *term.State
}

// New creates a new terminal state
func New() (*State, error) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return nil, err
	}

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}

	return &State{
		width:    width,
		height:   height,
		oldState: oldState,
	}, nil
}

// Cleanup restores the terminal state
func (ts *State) Cleanup() {
	term.Restore(int(os.Stdin.Fd()), ts.oldState)
	fmt.Fprint(os.Stdout, "\r\n👋 Goodbye!\r\n")
	os.Stdout.Sync()
}

// ClearScreen clears the terminal screen
func (ts *State) ClearScreen() {
	fmt.Fprint(os.Stdout, "\033[H\033[2J")
	os.Stdout.Sync()
}

// Print prints a string to stdout and flushes
func (ts *State) Print(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, format, args...)
	os.Stdout.Sync()
}
