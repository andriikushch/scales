// Command mcpserver exposes the scales library as an MCP server over stdio.
package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "scales", Version: "0.1.0"}, nil)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_scale",
		Description: "Get the notes (and, for major/minor scales, the scale-degree chords) of a musical scale in a given key.",
	}, getScale)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "parse_chord",
		Description: "Parse a chord notation (e.g. Cmaj7, G7, Dm7b5) into its root, description, and notes.",
	}, parseChord)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "render_fretboard",
		Description: "Render a set of notes (e.g. from get_scale or parse_chord) as a PNG fretboard diagram for guitar, bass guitar, ukulele, or mandolin.",
	}, renderFretboard)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
