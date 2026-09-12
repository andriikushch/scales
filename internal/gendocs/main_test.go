package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	scales "github.com/andriikushch/scales/pkg"
)

// Test_ReadmeScaleListUpToDate fails if README.md's generated scale-name block doesn't match
// what `make generate-docs` would currently produce - i.e. someone added/removed a scale or
// alias without re-running the generator.
func Test_ReadmeScaleListUpToDate(t *testing.T) {
	t.Parallel()

	committed, err := os.ReadFile("../../README.md")
	require.NoError(t, err)

	regenerated, err := replaceBlock(string(committed), GenerateScaleListBlock(scales.ScaleNames()))
	require.NoError(t, err)

	require.Equal(t, string(committed), regenerated, "README.md's scale-name list is stale — run `make generate-docs`")
}
