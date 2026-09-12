// Command gendocs regenerates the scale-name list in README.md and the full scale reference in
// docs/scales.md from scales.ScaleNames()/NewScaleByName, so neither doc drifts from the
// registry.
package main

import (
	"fmt"
	"os"
	"strings"

	scales "github.com/andriikushch/scales/pkg"
)

const (
	readmePath  = "README.md"
	beginMarker = "<!-- BEGIN GENERATED: scale-names (run `make generate-docs` to update) -->"
	endMarker   = "<!-- END GENERATED: scale-names -->"
)

// GenerateScaleListBlock returns the markdown that belongs between beginMarker and endMarker.
func GenerateScaleListBlock(names []string) string {
	quoted := make([]string, len(names))
	for i, n := range names {
		quoted[i] = "`" + n + "`"
	}

	return fmt.Sprintf(
		"**%d scale names** are available via `scales.ScaleNames()` / `scales.NewScaleByName`, "+
			"the `-scale` flag, and the MCP `get_scale` tool — canonical scales and alternate/folk "+
			"names alike (see `docs/scales.md` for structures and aliases):\n\n%s",
		len(names), strings.Join(quoted, ", "),
	)
}

// replaceBlock rewrites the markdown between beginMarker and endMarker in content with block.
func replaceBlock(content, block string) (string, error) {
	start := strings.Index(content, beginMarker)
	end := strings.Index(content, endMarker)
	if start == -1 || end == -1 || end < start {
		return "", fmt.Errorf("could not find generated-block markers in %s", readmePath)
	}

	before := content[:start+len(beginMarker)]
	after := content[end:]

	return before + "\n\n" + block + "\n\n" + after, nil
}

func runReadme(path string, names []string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	updated, err := replaceBlock(string(content), GenerateScaleListBlock(names))
	if err != nil {
		return err
	}

	return os.WriteFile(path, []byte(updated), 0o644)
}

func main() {
	if err := runReadme(readmePath, scales.ScaleNames()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := writeScalesMarkdown(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
