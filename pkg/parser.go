package scales

import (
	"errors"
	"github.com/andriikushch/scales/pkg/internal"
	"slices"
	"strconv"
)

type parser struct {
	l internal.Lexer
}

var (
	errIsNotInt             = errors.New("is not integer")
	errMissingDegreeMapping = errors.New("missing degree mapping")
	errUnexpectedInterval   = errors.New("unexpected interval")
)

func (p parser) parse(input string) (*Chord, error) {
	tokens, err := p.l.Tokenize(input)
	if err != nil {
		return nil, err
	}

	chord := &Chord{}
	chord.description = input

	for tokenIndex, token := range tokens {

		isLastToken := tokenIndex == len(tokens)-1

		switch token.Type {
		case internal.ROOT:
			err = p.initMajorChord(chord, token)
		case internal.NUMBER:
			err = p.parseNumber(token, chord, tokenIndex, tokens, internal.Numbers, internal.NthMapping)
		case internal.MAJ:
			err = p.parseMaj(chord, token, tokenIndex, tokens)
		case internal.FLAT:
			if token.Value != "" {
				chord.addType(internal.Flat)
			}
			err = p.parseAlt(chord, token, internal.Degrees, internal.FlatDegrees, chord.flatFirst)
		case internal.SHARP:
			if token.Value != "" {
				chord.addType(internal.Sharp)
			}
			err = p.parseAlt(chord, token, internal.Degrees, internal.SharpDegrees, chord.sharpFirst)
		case internal.SUS:
			err = p.parseSus(token, chord)
		case internal.DIM:
			err = p.initDimChord(chord, token)
		case internal.AUG:
			err = p.initAugChord(chord, token)
		case internal.MINOR:
			err = p.initMinorChord(chord, token, isLastToken)
		case internal.BASS:
			err = p.parseBass(token, chord)
		case internal.ADD:
			err = p.parseAdd(token, chord, tokenIndex, tokens, internal.Degrees)
		}

		if err != nil {
			return nil, err
		}
	}

	return chord, nil
}

func (p parser) parseAdd(token internal.Token, chord *Chord, tokenIndex int, tokens []internal.Token, degrees map[int]int) error {
	tokenValue, err := strconv.Atoi(token.Value)
	if err != nil {
		return errors.Join(errIsNotInt, err)
	}

	// add degree
	if val, ok := degrees[tokenValue]; ok {
		err = chord.add(val)
		if err != nil {
			return err
		}
	} else {
		return errMissingDegreeMapping
	}

	if len(chord.cType) == 1 && chord.cType[0] == internal.Major {
		chord.cType[0] = internal.Added
	} else {
		chord.addType(internal.Added)
	}

	if len(chord.cType) == 1 && tokenIndex+1 < len(tokens) {
		if val, ok := internal.Numbers[tokenValue]; ok {
			chord.addType(val)
		} else {
			return errMissingDegreeMapping
		}
	} else {
		if val, ok := internal.NthMapping[tokenValue]; ok {
			chord.addType(val)
		} else {
			return errMissingDegreeMapping
		}
	}

	return nil
}

func (p parser) parseBass(token internal.Token, chord *Chord) error {
	bassNote := NewNote(token.Value)
	chord.setBase(bassNote)
	distance := defaultChromaticScale.findDistance(chord.root, bassNote)

	if slices.Index(chord.structure, distance) == -1 {
		err := chord.add(distance)
		if err != nil {
			return err
		}
	}

	chord.addType(internal.Over)
	chord.addType(bassNote.Name)

	return nil
}

func (p parser) parseSus(token internal.Token, chord *Chord) error {
	var tokenValue int

	var err error

	if token.Value != "" {
		tokenValue, err = strconv.Atoi(token.Value)
		if err != nil {
			return errIsNotInt
		}
	} else {
		tokenValue = 4 // default values
	}

	chord.setChordBasicType(internal.Suspended)

	if len(chord.cType) == 1 {
		chord.cType[0] = internal.Suspended
	} else {
		chord.addType(internal.Suspended)
	}

	sus2 := 2
	sus4 := 4

	switch tokenValue {
	case sus2:
		switch chord.structure[1] {
		case internal.IM3, internal.Im3:
			chord.structure[1] = internal.IM2
			chord.addType(internal.Second)
		default:
			return errUnexpectedInterval
		}

	case sus4:
		switch chord.structure[1] {
		case internal.IM3, internal.Im3:
			chord.structure[1] = internal.IP4
			chord.addType(internal.Fourth)
		default:
			return errUnexpectedInterval
		}
	default:
		return errUnexpectedInterval
	}

	return nil
}

func (p parser) parseAlt(chord *Chord, token internal.Token, degrees map[int]int, mDegrees map[int]int, modifyFunc func(i int)) error {
	if token.Value == "" {
		return nil
	}

	tokenValue, err := strconv.Atoi(token.Value)
	if err != nil {
		return errIsNotInt
	}

	toModify, ok := degrees[tokenValue]
	if !ok {
		return errMissingDegreeMapping
	}

	index := slices.Index(chord.structure, toModify)
	if index == -1 {
		if val, ok := mDegrees[tokenValue]; ok {
			err = chord.add(val)
			if err != nil {
				return err
			}
		} else {
			return errMissingDegreeMapping
		}
	} else {
		modifyFunc(toModify)
	}

	if val, ok := internal.Numbers[tokenValue]; ok {
		chord.addType(val)
	} else {
		return errMissingDegreeMapping
	}

	return nil
}

func (p parser) parseMaj(chord *Chord, token internal.Token, i int, tokens []internal.Token) error {
	if chord.cType[len(chord.cType)-1] != internal.Major {
		chord.addType(internal.Major)
	}

	if token.Value == "" {
		return nil
	}

	tokenValue, err := strconv.Atoi(token.Value)
	if err != nil {
		return err
	}

	if len(chord.cType) == 1 && i+1 < len(tokens) {
		if val, ok := internal.Numbers[tokenValue]; ok {
			chord.addType(val)
		} else {
			return errMissingDegreeMapping
		}
	} else {
		if val, ok := internal.NthMapping[tokenValue]; ok {
			chord.addType(val)
		} else {
			return errMissingDegreeMapping
		}
	}

	for i := 7; i <= tokenValue; i += 2 {
		if val, ok := internal.Degrees[i]; ok {
			err = chord.add(val)
			if err != nil {
				return err
			}
		} else {
			return errMissingDegreeMapping
		}
	}

	return nil
}

func (p parser) parseNumber(token internal.Token, chord *Chord, i int, tokens []internal.Token, numbers, nthMapping map[int]string) error {
	tokenValue, err := strconv.Atoi(token.Value)
	if err != nil {
		return errIsNotInt
	}

	if len(chord.cType) == 1 && chord.cType[0] == internal.Major && tokenValue >= 7 {
		chord.cType[0] = internal.Dominant
	}

	if len(chord.cType) == 1 && chord.cType[0] == internal.Major && tokenValue <= 7 {
		chord.cType = []string{}
	}

	var chordType string

	var mappingExists bool

	if len(chord.cType) == 1 && i+1 < len(tokens) && tokens[i+1].Type != internal.SUS {
		chordType, mappingExists = numbers[tokenValue]
	} else {
		chordType, mappingExists = nthMapping[tokenValue]
	}

	if !mappingExists {
		return errMissingDegreeMapping
	}

	chord.addType(chordType)

	if tokenValue >= 7 {
		for degree := 7; degree <= tokenValue; degree += 2 {
			switch degree {
			case 7:
				if chord.cType[0] == internal.Dominant {
					err = chord.add(internal.Im7)
				}
			default:
				if val, ok := internal.Degrees[degree]; ok {
					err = chord.add(val)
				} else {
					return errMissingDegreeMapping
				}
			}

			if err != nil {
				return err
			}
		}
	} else {
		// add degree
		if val, ok := internal.Degrees[tokenValue]; ok {
			err = chord.add(val)
			if err != nil {
				return err
			}
		} else {
			return errMissingDegreeMapping
		}
	}

	return nil
}

func (p parser) initMinorChord(chord *Chord, token internal.Token, isLastToken bool) error {
	chord.setChordBasicType(internal.Minor)

	if len(chord.cType) > 0 {
		chord.cType[0] = internal.Minor
	}

	// convert first M3 tp m3
	chord.flatFirst(internal.IM3)

	if token.Value == "" {
		return nil
	}

	tokenValue, err := strconv.Atoi(token.Value)
	if err != nil {
		return errIsNotInt
	}

	switch tokenValue {
	case 7:
		err = chord.add(internal.Im7)
	case 9:
		err = chord.addIntervals(internal.Im7, internal.IM9)
	case 11:
		err = chord.addIntervals(internal.Im7, internal.IM9, internal.IP11)
	case 13:
		err = chord.addIntervals(internal.Im7, internal.IM9, internal.IP11, internal.IM13)
	default:
		err = errUnexpectedInterval
	}

	if err != nil {
		return err
	}

	var typeString string
	var ok bool

	if isLastToken {
		typeString, ok = internal.NthMapping[tokenValue]
		if !ok {
			return errMissingDegreeMapping
		}
	} else {
		typeString, ok = internal.Numbers[tokenValue]
		if !ok {
			return errMissingDegreeMapping
		}
	}

	chord.addType(typeString)
	return nil

}

func (p parser) initMajorChord(chord *Chord, token internal.Token) error {
	chord.root = NewNote(token.Value)

	// assume that this is major chord
	chord.setChordBasicType(internal.Major)
	chord.addType(internal.Major)

	return chord.addIntervals(internal.IUnison, internal.IM3, internal.IP5)
}

func (p parser) initAugChord(chord *Chord, token internal.Token) error {
	chord.setChordBasicType(internal.Augmented)

	if len(chord.cType) > 0 {
		chord.cType[0] = internal.Augmented
	}

	// add half step to the 5th
	chord.sharpFirst(internal.IP5)

	tokenValue, err := strconv.Atoi(token.Value)
	if err != nil {
		return errIsNotInt
	}

	switch tokenValue {
	case 6:
		err = chord.add(internal.IM6)
	case 7:
		err = chord.add(internal.Im7)
	case 9:
		err = chord.addIntervals(internal.Im7, internal.IM9)
	case 11:
		err = chord.addIntervals(internal.Im7, internal.IM9, internal.IP11)
	default:
		err = errUnexpectedInterval
	}

	if err != nil {
		return err
	}
	typeString, ok := internal.NthMapping[tokenValue]
	if !ok {
		return errMissingDegreeMapping
	}

	chord.addType(typeString)
	return nil
}

func (p parser) initDimChord(chord *Chord, token internal.Token) error {
	chord.setChordBasicType(internal.Diminished)

	if len(chord.cType) > 0 {
		chord.cType[0] = internal.Diminished
	}

	// convert first M3 tp m3
	chord.flatFirst(internal.IM3)
	chord.flatFirst(internal.IP5)

	if token.Value == "" {
		return nil
	}

	tokenValue, err := strconv.Atoi(token.Value)
	if err != nil {
		return errIsNotInt
	}

	switch tokenValue {
	case 6:
		err = chord.add(internal.IM6)
	case 7:
		err = chord.add(internal.ID7)
	case 9:
		err = chord.addIntervals(internal.ID7, internal.IM9)
	case 11:
		err = chord.addIntervals(internal.ID7, internal.IM9, internal.IP11)
	default:
		err = errUnexpectedInterval
	}

	if err != nil {
		return err
	}
	typeString, ok := internal.NthMapping[tokenValue]
	if !ok {
		return errMissingDegreeMapping
	}

	chord.addType(typeString)
	return nil
}
