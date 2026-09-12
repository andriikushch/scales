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
	ScaleType string `json:"scaleType" jsonschema:"scale type: major, minor, pentatonic, or any name from scales.ScaleNames() (e.g. dorian, harmonicMajor, bebopDominant, or an alias like freygish, jazzMinor, manGong) — see docs/scales.md for the full list, including its 'Other names' alias table"`
	Detail    string `json:"detail,omitempty" jsonschema:"for minor: natural, harmonic, or melodic (default natural); for pentatonic: major or minor (default minor); ignored otherwise"`
}

// GetScaleOutput is the output of the get_scale tool.
type GetScaleOutput struct {
	Notes []string `json:"notes"`
}

type notesProvider interface {
	GetNotes() []scales.Note
}

func getScale(_ context.Context, _ *mcp.CallToolRequest, in GetScaleInput) (*mcp.CallToolResult, GetScaleOutput, error) {
	var (
		notes notesProvider
		err   error
	)

	switch in.ScaleType {
	case "major":
		var sc *scales.Scale
		sc, err = scales.NewMajorScale(in.Key)
		notes = sc
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
		notes = sc
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
		var sc *scales.Scale
		sc, err = scales.NewScaleByName(in.ScaleType, in.Key)
		notes = sc
	}

	if err != nil {
		return nil, GetScaleOutput{}, err
	}

	out := GetScaleOutput{}
	for _, n := range notes.GetNotes() {
		out.Notes = append(out.Notes, n.Name)
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
