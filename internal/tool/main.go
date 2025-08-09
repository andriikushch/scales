package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"regexp"
	"sort"
	"strings"

	scales "github.com/andriikushch/scales/pkg"
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
				for j := range shapes {
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

		// After processing all chords, append the AllChordShapes slice
		sliceName := fmt.Sprintf("all%sChordShapes", strings.Title(instr.Name))
		err := updateChordShapeSliceAST(instr.ChordShapesFilePath, sliceName)
		if err != nil {
			fmt.Printf("Failed to update %s: %v\n", instr.ChordShapesFilePath, err)
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
	input = strings.ReplaceAll(input, "#", "_sharp_")
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
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString("\n" + content + "\n")
	return err
}

// updateChordShapeSliceAST finds all *_ChordShape variable declarations and creates/updates a single AllXxxChordShapes slice
func updateChordShapeSliceAST(filePath, sliceName string) error {
	fset := token.NewFileSet()
	src, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Parse source file
	node, err := parser.ParseFile(fset, filePath, src, parser.ParseComments)
	if err != nil {
		return err
	}

	var shapeVars []string
	sliceDeclIdx := -1

	// Inspect top-level declarations
	for i, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.VAR {
			continue
		}
		for _, spec := range genDecl.Specs {
			valueSpec := spec.(*ast.ValueSpec)
			for _, name := range valueSpec.Names {
				if name.Name == sliceName {
					sliceDeclIdx = i // Found existing slice definition
				}
				if strings.HasSuffix(name.Name, "ChordShape") {
					shapeVars = append(shapeVars, name.Name)
				}
			}
		}
	}

	sort.Strings(shapeVars) // Optional: keep the list sorted

	// Build the new slice declaration
	sliceExpr := &ast.GenDecl{
		Tok: token.VAR,
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Names: []*ast.Ident{ast.NewIdent(sliceName)},
				Type: &ast.ArrayType{
					Elt: ast.NewIdent("ChordShape"),
				},
				Values: []ast.Expr{
					&ast.CompositeLit{
						Type: &ast.ArrayType{Elt: ast.NewIdent("ChordShape")},
						Elts: buildIdentList(shapeVars),
					},
				},
			},
		},
	}

	// Replace or append the slice declaration
	if sliceDeclIdx != -1 {
		node.Decls[sliceDeclIdx] = sliceExpr
	} else {
		node.Decls = append(node.Decls, sliceExpr)
	}

	// Write the modified AST back to the file
	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	err = printer.Fprint(outFile, fset, node)
	if err != nil {
		return err
	}

	return nil
}

// buildIdentList builds a comma-separated list of identifiers (e.g. guitar_CMaj_0ChordShape)
func buildIdentList(names []string) []ast.Expr {
	var list []ast.Expr
	for _, name := range names {
		list = append(list, ast.NewIdent(name))
	}
	return list
}
