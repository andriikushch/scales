package scales

import (
	"errors"
	"fmt"
	"slices"
	"strconv"

	"github.com/andriikushch/scales/pkg/internal"
)

type parser struct {
	l internal.Lexer
}

var (
	errIsNotInt             = errors.New("is not integer")
	errMissingDegreeMapping = errors.New("missing degree mapping")
	errUnexpectedInterval   = errors.New("unexpected interval")
	errUnknownNote          = errors.New("unknown note")
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
			err = p.parseNumber(token, chord, tokenIndex, tokens)
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
			err = p.initDimChord(chord, token, isLastToken)
		case internal.AUG:
			err = p.initAugChord(chord, token, isLastToken)
		case internal.MINOR:
			err = p.initMinorChord(chord, token, isLastToken)
		case internal.BASS:
			err = p.parseBass(token, chord)
		case internal.ADD:
			err = p.parseAdd(token, chord, tokenIndex, tokens, internal.Degrees)
		default:
			return nil, fmt.Errorf("unsupported token type %q", token.Type)
		}

		if err != nil {
			return nil, err
		}
	}

	return chord, nil
}

func (p parser) pickQualityString(tokenValue int, useShortForm bool) (string, error) {
	m := internal.NthMapping
	if useShortForm {
		m = internal.Numbers
	}

	val, ok := m[tokenValue]
	if !ok {
		return "", errMissingDegreeMapping
	}

	return val, nil
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

	if len(chord.quality) == 1 && chord.quality[0] == internal.Major {
		chord.quality[0] = internal.Added
	} else {
		chord.addType(internal.Added)
	}

	useShortForm := len(chord.quality) == 1 && tokenIndex+1 < len(tokens)

	typeString, err := p.pickQualityString(tokenValue, useShortForm)
	if err != nil {
		return err
	}

	chord.addType(typeString)

	return nil
}

func (p parser) parseBass(token internal.Token, chord *Chord) error {
	bassNote := NewNote(token.Value)
	if defaultChromaticScale.position(bassNote) == -1 {
		return fmt.Errorf("%w: %q is not a known note", errUnknownNote, token.Value)
	}

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
			return errors.Join(errIsNotInt, err)
		}
	} else {
		tokenValue = 4 // default values
	}

	chord.setChordBasicType(internal.Suspended)

	if len(chord.quality) == 1 {
		chord.quality[0] = internal.Suspended
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
		return errors.Join(errIsNotInt, err)
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
	if chord.quality[len(chord.quality)-1] != internal.Major {
		chord.addType(internal.Major)
	}

	if token.Value == "" {
		return nil
	}

	tokenValue, err := strconv.Atoi(token.Value)
	if err != nil {
		return err
	}

	useShortForm := len(chord.quality) == 1 && i+1 < len(tokens)

	typeString, err := p.pickQualityString(tokenValue, useShortForm)
	if err != nil {
		return err
	}

	chord.addType(typeString)

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

func (p parser) parseNumber(token internal.Token, chord *Chord, i int, tokens []internal.Token) error {
	tokenValue, err := strconv.Atoi(token.Value)
	if err != nil {
		return errIsNotInt
	}

	if len(chord.quality) == 1 && chord.quality[0] == internal.Major && tokenValue >= 7 {
		chord.quality[0] = internal.Dominant
	}

	if len(chord.quality) == 1 && chord.quality[0] == internal.Major && tokenValue <= 7 {
		chord.quality = []string{}
	}

	useShortForm := len(chord.quality) == 1 && i+1 < len(tokens) && tokens[i+1].Type != internal.SUS

	chordType, err := p.pickQualityString(tokenValue, useShortForm)
	if err != nil {
		return err
	}

	chord.addType(chordType)

	if tokenValue >= 7 {
		for degree := 7; degree <= tokenValue; degree += 2 {
			switch degree {
			case 7:
				if chord.quality[0] == internal.Dominant {
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

var minorExtensions = map[int][]int{
	6:  {internal.IM6},
	7:  {internal.Im7},
	9:  {internal.Im7, internal.IM9},
	11: {internal.Im7, internal.IM9, internal.IP11},
	13: {internal.Im7, internal.IM9, internal.IP11, internal.IM13},
}

var augExtensions = map[int][]int{
	6:  {internal.IM6},
	7:  {internal.Im7},
	9:  {internal.Im7, internal.IM9},
	11: {internal.Im7, internal.IM9, internal.IP11},
	13: {internal.Im7, internal.IM9, internal.IP11, internal.IM13},
}

var dimExtensions = map[int][]int{
	6:  {internal.Im6},
	7:  {internal.ID7},
	9:  {internal.ID7, internal.IM9},
	11: {internal.ID7, internal.IM9, internal.IP11},
	13: {internal.ID7, internal.IM9, internal.IP11, internal.IM13},
}

func (p parser) initChordWithExtension(chord *Chord, token internal.Token, isLastToken bool, basicType string, structuralAdjust func(*Chord), extensions map[int][]int) error {
	chord.setChordBasicType(basicType)

	if len(chord.quality) > 0 {
		chord.quality[0] = basicType
	}

	structuralAdjust(chord)

	if token.Value == "" {
		return nil
	}

	tokenValue, err := strconv.Atoi(token.Value)
	if err != nil {
		return errors.Join(errIsNotInt, err)
	}

	intervals, ok := extensions[tokenValue]
	if !ok {
		return errUnexpectedInterval
	}

	if err := chord.addIntervals(intervals...); err != nil {
		return err
	}

	typeString, err := p.pickQualityString(tokenValue, !isLastToken)
	if err != nil {
		return err
	}

	chord.addType(typeString)

	return nil
}

func (p parser) initMinorChord(chord *Chord, token internal.Token, isLastToken bool) error {
	return p.initChordWithExtension(chord, token, isLastToken, internal.Minor,
		func(c *Chord) { c.flatFirst(internal.IM3) }, minorExtensions)
}

func (p parser) initMajorChord(chord *Chord, token internal.Token) error {
	chord.root = NewNote(token.Value)

	// assume that this is major chord
	chord.setChordBasicType(internal.Major)
	chord.addType(internal.Major)

	return chord.addIntervals(internal.IUnison, internal.IM3, internal.IP5)
}

func (p parser) initAugChord(chord *Chord, token internal.Token, isLastToken bool) error {
	return p.initChordWithExtension(chord, token, isLastToken, internal.Augmented,
		func(c *Chord) { c.sharpFirst(internal.IP5) }, augExtensions)
}

func (p parser) initDimChord(chord *Chord, token internal.Token, isLastToken bool) error {
	return p.initChordWithExtension(chord, token, isLastToken, internal.Diminished,
		func(c *Chord) {
			c.flatFirst(internal.IM3)
			c.flatFirst(internal.IP5)
		}, dimExtensions)
}
