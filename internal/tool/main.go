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
		key, val := sl[0], sl[1]
		parsableChords[key] = append(parsableChords[key], val)
	}

	instruments := []Instrument{
		{"guitar", scales.NewGuitarWithStandardTuning(), scales.GuitarChordShapes, fileWithGuitarChordShapes},
		{"ukulele", scales.NewUkuleleWithStandardTuning(), scales.UkuleleChordShapes, fileWithUkuleleChordShapes},
		{"bassGuitar", scales.NewBassGuitarWithStandardTuning(), scales.BassGuitarChordShapes, fileWithBassGuitarChordShapes},
		{"mandolin", scales.NewMandolinWithStandardTuning(), scales.MandolinChordShapes, fileWithMandolinChordShapes},
	}

	for _, instr := range instruments {
		fmt.Println(strings.ToLower(instr.Name))
		for structure, notations := range parsableChords {
			if instr.Instance.GetStringsAmount() < strings.Count(structure, "-")+1 {
				fmt.Printf("Not enough strings on the %s for %v %v\n", instr.Name, structure, notations)
				continue
			}
			if shapes, ok := instr.ChordShapes[structure]; ok {
				for j := range notations {
					shapes[j].PossibleNotations = notations
				}
			} else {
				fmt.Printf("Add %s chord shape for '%v' '%v' \n", instr.Name, structure, notations)

				code, variableName := createChordShapeCode(instr.Name, ReplaceSpecialCharacters(notations[0]), 0, structure)

				// Only check for variable name, not whole template
				exists, err := checkIfVariableExists(instr.ChordShapesFilePath, variableName+"ChordShape")
				if err != nil {
					fmt.Printf("Error reading %s: %v\n", instr.ChordShapesFilePath, err)
					continue
				}

				if !exists {
					err := appendToFile(instr.ChordShapesFilePath, code)
					if err != nil {
						fmt.Printf("Error writing to %s: %v\n", instr.ChordShapesFilePath, err)
					} else {
						fmt.Printf("Added template for %s to %s\n", structure, instr.ChordShapesFilePath)
					}
				}
			}
		}
	}
}

func createChordShapeCode(instrumentName, sanitizedChordName string, position int, structure string) (code string, varName string) {
	varName = fmt.Sprintf("%s_%s_%d_", instrumentName, sanitizedChordName, position)
	frets := strings.ReplaceAll(structure, "-", ",")
	code = fmt.Sprintf(`
var (
	%[1]sSchema = []int{} // TODO
	%[1]sChordShape = newChordShape("%[2]s", []int{%[3]s}, %[1]sSchema, %[4]d)
)
`, varName, instrumentName, frets, -1)
	return code, varName
}

func ReplaceSpecialCharacters(input string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	return strings.Trim(re.ReplaceAllString(input, "_"), "_")
}

func checkIfVariableExists(filePath, variableName string) (bool, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return false, err
	}
	return strings.Contains(string(data), variableName), nil
}

func appendToFile(filePath, content string) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString("\n" + content + "\n")
	return err
}
