package app

import (
	scales "github.com/andriikushch/scales/pkg"
)

// Instrument types
const (
	InstrumentBassGuitar = "bassGuitar"
	InstrumentGuitar     = "guitar"
	InstrumentUkulele    = "ukulele"
	InstrumentMandolin   = "mandolin"
)

// State represents the application state
type State struct {
	CurrentKeyIndex  int
	ShowScaleTypes   bool
	ShowInstruments  bool
	Scale            *scales.Scale
	ScaleName        string
	Instrument       string
	ScaleNames       []string
	ValidInstruments []string
	Keys             []string
}

// New creates a new application state
func New() *State {
	return &State{
		ScaleNames:       scales.ScaleNames(),
		ValidInstruments: []string{InstrumentBassGuitar, InstrumentGuitar, InstrumentUkulele, InstrumentMandolin},
		Keys:             []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"},
	}
}

// CreateScale creates a new scale based on current settings
func (s *State) CreateScale() error {
	var err error
	s.Scale, err = scales.NewScaleByName(s.ScaleName, s.Keys[s.CurrentKeyIndex])
	return err
}

// GetNotes returns the notes for the current scale
func (s *State) GetNotes() []scales.Note {
	return s.Scale.GetNotes()
}
