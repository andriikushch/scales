package main

import (
	"bufio"
	"fmt"
	scales "github.com/andriikushch/scales/pkg"
	"os"
	"regexp"
	"strings"
)

type StringInstrument interface {
	GetStringsAmount() int
}

type Instrument struct {
	Name                string
	Instance            StringInstrument
	ChordShapes         map[string][]scales.ChordShape
	ChordShapesFilePath string
}

func main() {

	if len(os.Args) < 6 {
		panic("missing some parameters")
	}

	fileWithUniqQualities := os.Args[1]

	fileWithGuitarChordShapes := os.Args[2]
	fileWithUkuleleChordShapes := os.Args[3]
	fileWithBassGuitarChordShapes := os.Args[4]
	fileWithMandolinChordShapes := os.Args[5]

	f, err := os.Open(fileWithUniqQualities)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	parsableChords := map[string][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		sl := strings.Split(line, ";")

		_, ok := parsableChords[sl[0]]
		if ok {
			parsableChords[sl[0]] = append(parsableChords[sl[0]], sl[1])
		} else {
			parsableChords[sl[0]] = []string{sl[1]}
		}
	}

	instruments := []Instrument{
		{
			Name:                "guitar",
			Instance:            scales.NewGuitarWithStandardTuning(),
			ChordShapes:         scales.GuitarChordShapes,
			ChordShapesFilePath: fileWithGuitarChordShapes,
		},
		{
			Name:                "ukulele",
			Instance:            scales.NewUkuleleWithStandardTuning(),
			ChordShapes:         scales.UkuleleChordShapes,
			ChordShapesFilePath: fileWithUkuleleChordShapes,
		},
		{
			Name:                "bass guitar",
			Instance:            scales.NewBassGuitarWithStandardTuning(),
			ChordShapes:         scales.BassGuitarChordShapes,
			ChordShapesFilePath: fileWithBassGuitarChordShapes,
		},
		{
			Name:                "mandolin",
			Instance:            scales.NewMandolinWithStandardTuning(),
			ChordShapes:         scales.MandolinChordShapes,
			ChordShapesFilePath: fileWithMandolinChordShapes,
		},
	}

	for _, instr := range instruments {
		fmt.Println(strings.ToLower(instr.Name))
		for structure, v := range parsableChords {
			if instr.Instance.GetStringsAmount() < strings.Count(structure, "-")+1 {
				fmt.Printf("Not enough strings on the %s for %v %v\n", instr.Name, structure, v)
				continue
			}
			if shapes, ok := instr.ChordShapes[structure]; ok {
				for j := range v {
					shapes[j].PossibleNotations = v
				}
			} else {
				fmt.Printf("Add %s chord shape for '%v' '%v' \n", instr.Name, structure, v)
				fmt.Println(createChordShapeCode(instr.Name, ReplaceSpecialCharacters(v[0]), 0, structure))
			}
		}
	}
}

func createChordShapeCode(instrumentName string, chordName string, position int, structure string) (string, string) {
	// Sanitize names to be identifier-safe
	capitalizedChord := strings.Title(strings.ReplaceAll(chordName, "-", ""))
	shapeName := fmt.Sprintf("%s_%s_%d_", instrumentName, capitalizedChord, position)

	// Create the input slice as string

	code := fmt.Sprintf(`
var (
	%[1]s_schema = []int{} // TODO
	%[1]sChordShape = newChordShape("%[2]s", []int{%[3]s}, %[1]s_schema, %[4]d)
)
`, shapeName, instrumentName, strings.ReplaceAll(structure, "-", ","), -1)

	return code, shapeName
}

// ReplaceSpecialCharacters replaces all non-alphanumeric characters with underscores
func ReplaceSpecialCharacters(input string) string {
	// Replace all non-alphanumeric characters with "_"
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	result := re.ReplaceAllString(input, "_")

	// Optional: remove leading/trailing underscores and lowercase the result
	result = strings.Trim(result, "_")
	return result
}
