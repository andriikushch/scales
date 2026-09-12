package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	scales "github.com/andriikushch/scales/pkg"
)

const scalesMdPath = "docs/scales.md"

const scalesMdFooter = `## Two things worth knowing

1. Alternate/folk names are collected in the "Other names" table above rather than duplicated as their own rows in the family tables earlier in this document — e.g. the "Chinese" alias and Raga Bhupali are both Major Pentatonic; Man Gong is Blues Minor Pentatonic.

2. Spellings like F♭ and B♭♭ are deliberate. They sound like E and A but each scale degree needs its own letter.
`

// GenerateScalesMarkdown builds the full docs/scales.md content from scaleCategories and
// scaleAliases, computing each scale's Structure/Degrees/C-example columns from
// scales.NewScaleByName via the arithmetic in scale_arithmetic.go.
func GenerateScalesMarkdown() (string, error) {
	total := 0
	for _, cat := range scaleCategories {
		total += len(cat.Scales)
	}

	var b strings.Builder

	fmt.Fprintf(&b, "# Scale Reference\n\n")
	fmt.Fprintf(&b, "%d scales, each a different set of notes when rooted on C.\n\n", total)
	fmt.Fprintf(&b, "**Structure** is the gap between each note in semitones, including the gap from the last note back to the root — so the numbers always total 12. Every row is checked: totals, matching counts across columns, and the example notes actually following the structure.\n\n")

	for _, cat := range scaleCategories {
		fmt.Fprintf(&b, "## %s\n\n", cat.Title)
		fmt.Fprintf(&b, "| Scale | Structure | Degrees | C example |\n|---|---|---|---|\n")
		for _, entry := range cat.Scales {
			scale, err := scales.NewScaleByName(entry.Key, "C")
			if err != nil {
				return "", fmt.Errorf("building %s: %w", entry.Key, err)
			}

			notes := scale.GetNotes()
			fmt.Fprintf(&b, "| %s | %s | %s | %s |\n",
				entry.DisplayName, structureString(notes), degreesString(notes), exampleString(notes))
		}
		fmt.Fprintf(&b, "\n")
	}

	fmt.Fprintf(&b, "## Other names\n\n")
	fmt.Fprintf(&b, "| Name | Scale here |\n|---|---|\n")
	for _, alias := range scaleAliases {
		fmt.Fprintf(&b, "| %s | %s |\n", alias.DisplayName, alias.CanonicalDisplayName)
	}
	fmt.Fprintf(&b, "\n")

	b.WriteString(scalesMdFooter)

	return b.String(), nil
}

// writeScalesMarkdown regenerates docs/scales.md in place.
func writeScalesMarkdown() error {
	content, err := GenerateScalesMarkdown()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(scalesMdPath), 0o755); err != nil {
		return err
	}

	return os.WriteFile(scalesMdPath, []byte(content), 0o644)
}
