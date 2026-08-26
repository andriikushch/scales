package main

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	scales "github.com/andriikushch/scales/pkg"
)

// RenderFretboardInput is the input for the render_fretboard tool.
type RenderFretboardInput struct {
	Notes      []string `json:"notes" jsonschema:"note names to highlight on the fretboard, e.g. the notes returned by get_scale or parse_chord"`
	Instrument string   `json:"instrument,omitempty" jsonschema:"instrument: guitar, bassGuitar, ukulele, or mandolin (default guitar)"`
}

// RenderFretboardOutput is the (empty) structured output of the
// render_fretboard tool; the diagram itself is returned as image content.
type RenderFretboardOutput struct{}

type fretboard interface {
	Positions(notes []scales.Note) ([][]scales.FretCell, error)
}

const (
	fretCount    = 24
	cellWidth    = 40
	cellHeight   = 40
	leftMargin   = 50
	topMargin    = 20
	bottomMargin = 30
	noteRadius   = 14
)

var markerFrets = map[int]bool{5: true, 7: true, 12: true, 17: true, 19: true, 24: true}

func renderFretboard(_ context.Context, _ *mcp.CallToolRequest, in RenderFretboardInput) (*mcp.CallToolResult, RenderFretboardOutput, error) {
	instrumentName := in.Instrument
	if instrumentName == "" {
		instrumentName = "guitar"
	}

	var inst fretboard
	switch instrumentName {
	case "guitar":
		inst = scales.NewGuitarWithStandardTuning()
	case "bassGuitar":
		inst = scales.NewBassGuitarWithStandardTuning()
	case "ukulele":
		inst = scales.NewUkuleleWithStandardTuning()
	case "mandolin":
		inst = scales.NewMandolinWithStandardTuning()
	default:
		return nil, RenderFretboardOutput{}, fmt.Errorf("invalid instrument %q: must be guitar, bassGuitar, ukulele, or mandolin", in.Instrument)
	}

	notes := make([]scales.Note, len(in.Notes))
	for i, n := range in.Notes {
		notes[i] = scales.NewNote(n)
	}

	grid, err := inst.Positions(notes)
	if err != nil {
		return nil, RenderFretboardOutput{}, err
	}

	pngBytes, err := drawFretboardPNG(grid)
	if err != nil {
		return nil, RenderFretboardOutput{}, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.ImageContent{MIMEType: "image/png", Data: pngBytes},
		},
	}, RenderFretboardOutput{}, nil
}

// drawFretboardPNG renders a fretboard grid (as produced by
// (*stringInstrumentWithFrets).Positions, truncated to fretCount frets) to a
// PNG image, highlighting matched notes.
func drawFretboardPNG(grid [][]scales.FretCell) ([]byte, error) {
	numStrings := len(grid)
	width := leftMargin + (fretCount+1)*cellWidth
	height := topMargin + numStrings*cellHeight + bottomMargin

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), image.NewUniform(color.White), image.Point{}, draw.Src)

	// String lines (horizontal).
	for s := 0; s < numStrings; s++ {
		y := topMargin + s*cellHeight + cellHeight/2
		drawHLine(img, leftMargin, width-cellWidth/2, y, color.Black)
	}

	// Fret lines (vertical) and fret markers below the last string.
	markerY := topMargin + numStrings*cellHeight + bottomMargin/2
	for f := 0; f <= fretCount; f++ {
		x := leftMargin + f*cellWidth
		drawVLine(img, x, topMargin, topMargin+(numStrings-1)*cellHeight, color.Black)

		if markerFrets[f] {
			drawLabel(img, x-4, markerY, fmt.Sprintf("%d", f), color.Gray{Y: 96})
		}
	}

	// Notes.
	for s, row := range grid {
		y := topMargin + s*cellHeight + cellHeight/2
		for f, cell := range row {
			if f > fretCount || !cell.Match {
				continue
			}
			x := leftMargin + f*cellWidth

			fill := noteColor(cell.MatchIndex)
			drawFilledCircle(img, x, y, noteRadius, fill)
			drawLabel(img, x-len(cell.Note.Name)*3, y+4, cell.Note.Name, color.White)
		}
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// noteColor picks a fill color for a matched note: the root (index 0) is
// always red, other scale/chord degrees cycle through a small palette.
func noteColor(matchIndex int) color.Color {
	if matchIndex == 0 {
		return color.RGBA{R: 200, G: 30, B: 30, A: 255}
	}

	palette := []color.Color{
		color.RGBA{R: 30, G: 100, B: 200, A: 255},
		color.RGBA{R: 30, G: 150, B: 60, A: 255},
		color.RGBA{R: 180, G: 130, B: 20, A: 255},
		color.RGBA{R: 130, G: 60, B: 180, A: 255},
		color.RGBA{R: 20, G: 150, B: 150, A: 255},
	}

	return palette[(matchIndex-1)%len(palette)]
}

func drawHLine(img *image.RGBA, x0, x1, y int, c color.Color) {
	for x := x0; x <= x1; x++ {
		img.Set(x, y, c)
	}
}

func drawVLine(img *image.RGBA, x, y0, y1 int, c color.Color) {
	for y := y0; y <= y1; y++ {
		img.Set(x, y, c)
	}
}

func drawFilledCircle(img *image.RGBA, cx, cy, r int, c color.Color) {
	for y := -r; y <= r; y++ {
		for x := -r; x <= r; x++ {
			if x*x+y*y <= r*r {
				img.Set(cx+x, cy+y, c)
			}
		}
	}
}

func drawLabel(img *image.RGBA, x, y int, text string, c color.Color) {
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(c),
		Face: basicfont.Face7x13,
		Dot:  fixed.P(x, y),
	}
	d.DrawString(text)
}
