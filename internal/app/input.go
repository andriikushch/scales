package app

import (
	"bufio"
	"os"
)

// KeyInput represents a keyboard input
type KeyInput struct {
	Key  string
	Rune rune
}

// InputHandler handles keyboard input
type InputHandler struct {
	keyChan chan KeyInput
}

// NewInputHandler creates a new input handler
func NewInputHandler() *InputHandler {
	ih := &InputHandler{
		keyChan: make(chan KeyInput, 1),
	}
	go ih.readInput()
	return ih
}

// readInput reads keyboard input and sends it to the channel
func (ih *InputHandler) readInput() {
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
				ih.keyChan <- KeyInput{Key: "up", Rune: 0}
			case 'B':
				ih.keyChan <- KeyInput{Key: "down", Rune: 0}
			case 'C':
				ih.keyChan <- KeyInput{Key: "right", Rune: 0}
			case 'D':
				ih.keyChan <- KeyInput{Key: "left", Rune: 0}
			}
			continue
		}

		// Handle carriage return and line feed
		if r == '\r' {
			ih.keyChan <- KeyInput{Key: "", Rune: '\n'}
			continue
		}

		ih.keyChan <- KeyInput{Key: "", Rune: r}
	}
}

// GetInput returns the input channel
func (ih *InputHandler) GetInput() <-chan KeyInput {
	return ih.keyChan
}
