package app

import (
	scales "github.com/andriikushch/scales/pkg"
)

// Scale types
const (
	ScaleMajor      = "major"
	ScaleMinor      = "minor"
	ScalePentatonic = "pentatonic"
)

// Instrument types
const (
	InstrumentBassGuitar = "bassGuitar"
	InstrumentGuitar     = "guitar"
	InstrumentUkulele    = "ukulele"
)

// Minor scale types
const (
	MinorTypeNatural  = "natural"
	MinorTypeHarmonic = "harmonic"
	MinorTypeMelodic  = "melodic"
)

// Pentatonic scale types
const (
	PentatonicTypeMajor = "major"
	PentatonicTypeMinor = "minor"
)

// View types
const (
	ViewScale  = "scale"
	ViewChords = "chords"
)

// scale interface represents a musical scale
type scale interface {
	GetNotes() []scales.Note
	String() string
}

// State represents the application state
type State struct {
	CurrentKeyIndex   int
	CurrentView       string
	CurrentChordIndex int
	ShowScaleTypes    bool
	ShowInstruments   bool
	Scale             scale
	ScaleType         string
	ScaleTypeDetail   string
	Instrument        string
	ValidScales       []string
	ValidInstruments  []string
	MinorScales       []string
	PentatonicScales  []string
	Keys              []string
}

// New creates a new application state
func New() *State {
	return &State{
		CurrentView:      ViewScale,
		ValidScales:      []string{ScaleMajor, ScaleMinor, ScalePentatonic},
		ValidInstruments: []string{InstrumentBassGuitar, InstrumentGuitar, InstrumentUkulele},
		MinorScales:      []string{MinorTypeNatural, MinorTypeHarmonic, MinorTypeMelodic},
		PentatonicScales: []string{PentatonicTypeMajor, PentatonicTypeMinor},
		Keys:             []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"},
	}
}

// CreateScale creates a new scale based on current settings
func (s *State) CreateScale() error {
	var err error
	switch s.ScaleType {
	case ScaleMajor:
		s.Scale, err = scales.NewMajorScale(s.Keys[s.CurrentKeyIndex])
	case ScaleMinor:
		switch s.ScaleTypeDetail {
		case MinorTypeNatural:
			s.Scale, err = scales.NewNaturalMinorScale(s.Keys[s.CurrentKeyIndex])
		case MinorTypeHarmonic:
			s.Scale, err = scales.NewHarmonicMinorScale(s.Keys[s.CurrentKeyIndex])
		case MinorTypeMelodic:
			s.Scale, err = scales.NewMelodicMinorScale(s.Keys[s.CurrentKeyIndex])
		}
	case ScalePentatonic:
		switch s.ScaleTypeDetail {
		case PentatonicTypeMajor:
			s.Scale, err = scales.NewMajorPentatonicScale(s.Keys[s.CurrentKeyIndex])
		case PentatonicTypeMinor:
			s.Scale, err = scales.NewMinorPentatonicScale(s.Keys[s.CurrentKeyIndex])
		}
	}
	return err
}

// GetChords returns the chords for the current scale
func (s *State) GetChords() []scales.Chord {
	if sc, ok := s.Scale.(*scales.Scale); ok && s.ScaleType != ScalePentatonic {
		return sc.GetChords()
	}
	return nil
}

// GetNotes returns the notes for the current scale
func (s *State) GetNotes() []scales.Note {
	return s.Scale.GetNotes()
}

// GetCurrentChord returns the current chord
func (s *State) GetCurrentChord() scales.Chord {
	chords := s.GetChords()
	if chords != nil && s.CurrentChordIndex < len(chords) {
		return chords[s.CurrentChordIndex]
	}
	return scales.Chord{}
}
