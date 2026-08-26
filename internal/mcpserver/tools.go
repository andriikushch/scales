package main

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	scales "github.com/andriikushch/scales/pkg"
)

// GetScaleInput is the input for the get_scale tool.
type GetScaleInput struct {
	Key       string `json:"key" jsonschema:"root note/key, e.g. C, D#, F#, Bb"`
	ScaleType string `json:"scaleType" jsonschema:"scale type: major, minor, or pentatonic"`
	Detail    string `json:"detail,omitempty" jsonschema:"for minor: natural, harmonic, or melodic (default natural); for pentatonic: major or minor (default minor); ignored for major"`
}

// GetScaleOutput is the output of the get_scale tool.
type GetScaleOutput struct {
	Notes []string `json:"notes"`
	// Chords are the scale-degree chords (e.g. "Cmaj", "Dm", ...). Omitted for
	// pentatonic scales, which have none.
	Chords []string `json:"chords,omitempty"`
}

type notesProvider interface {
	GetNotes() []scales.Note
}

func getScale(_ context.Context, _ *mcp.CallToolRequest, in GetScaleInput) (*mcp.CallToolResult, GetScaleOutput, error) {
	var (
		notes  notesProvider
		chords *scales.Scale
		err    error
	)

	switch in.ScaleType {
	case "major":
		var sc *scales.Scale
		sc, err = scales.NewMajorScale(in.Key)
		notes, chords = sc, sc
	case "minor":
		detail := in.Detail
		if detail == "" {
			detail = "natural"
		}

		var sc *scales.Scale
		switch detail {
		case "natural":
			sc, err = scales.NewNaturalMinorScale(in.Key)
		case "harmonic":
			sc, err = scales.NewHarmonicMinorScale(in.Key)
		case "melodic":
			sc, err = scales.NewMelodicMinorScale(in.Key)
		default:
			return nil, GetScaleOutput{}, fmt.Errorf("invalid detail %q for a minor scale: must be natural, harmonic, or melodic", in.Detail)
		}
		notes, chords = sc, sc
	case "pentatonic":
		detail := in.Detail
		if detail == "" {
			detail = "minor"
		}

		var pent *scales.Pentatonic
		switch detail {
		case "major":
			pent, err = scales.NewMajorPentatonicScale(in.Key)
		case "minor":
			pent, err = scales.NewMinorPentatonicScale(in.Key)
		default:
			return nil, GetScaleOutput{}, fmt.Errorf("invalid detail %q for a pentatonic scale: must be major or minor", in.Detail)
		}
		notes = pent
	default:
		return nil, GetScaleOutput{}, fmt.Errorf("invalid scaleType %q: must be major, minor, or pentatonic", in.ScaleType)
	}

	if err != nil {
		return nil, GetScaleOutput{}, err
	}

	out := GetScaleOutput{}
	for _, n := range notes.GetNotes() {
		out.Notes = append(out.Notes, n.Name)
	}
	if chords != nil {
		for _, c := range chords.GetChords() {
			out.Chords = append(out.Chords, c.Description())
		}
	}

	return nil, out, nil
}

// ParseChordInput is the input for the parse_chord tool.
type ParseChordInput struct {
	Chord string `json:"chord" jsonschema:"chord notation, e.g. Cmaj7, G7, Dm7b5, F#dim"`
}

// ParseChordOutput is the output of the parse_chord tool.
type ParseChordOutput struct {
	Description string   `json:"description"`
	Notes       []string `json:"notes"`
}

func parseChord(_ context.Context, _ *mcp.CallToolRequest, in ParseChordInput) (*mcp.CallToolResult, ParseChordOutput, error) {
	chord, err := scales.NewChord(in.Chord)
	if err != nil {
		return nil, ParseChordOutput{}, fmt.Errorf("parsing chord %q: %w", in.Chord, err)
	}

	out := ParseChordOutput{Description: chord.Description()}
	for _, n := range chord.Notes() {
		out.Notes = append(out.Notes, n.Name)
	}

	return nil, out, nil
}
