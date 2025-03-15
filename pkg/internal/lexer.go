package internal

import (
	"errors"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

const (
	m         = "m"
	maj       = "maj"
	aug       = "aug"
	dim       = "dim"
	b         = "b"
	sharpSign = "#"
	add       = "add"
	slash     = "/"
	sus       = "sus"
)

var roots = []string{
	C,
	D,
	E,
	F,
	G,
	A,
	B,
}

var re = regexp.MustCompile(`[^0-9]`)

type Lexer struct {
	input      string
	currentPos int
	lookupPos  int
}

func (l *Lexer) Tokenize(input string) ([]Token, error) {
	if len(input) == 0 {
		return nil, nil
	}

	l.input = input

	var result []Token

	for {
		token, err := l.NextToken()
		if err != nil {
			return nil, err
		}

		if token.Type == END {
			break
		}

		if token.Type == IGNORED {
			continue
		}

		result = append(result, token)
	}

	return result, nil
}

func (l *Lexer) NextToken() (Token, error) {
	tokenType, rawValue, err := l.readChars()
	if err != nil {
		return Token{}, err
	}

	if rawValue == "" {
		return Token{Type: END}, nil
	}

	var t Token

	// first Token should be root
	switch tokenType {
	case ROOT, NUMBER:
		t = Token{
			Type:  tokenType,
			Value: rawValue,
		}
	case SHARP, FLAT, MAJ, DIM, AUG, IGNORED, MINOR, ADD, SUS:
		t = Token{
			Type:  tokenType,
			Value: l.removeNonNumbers(rawValue),
		}
	case BASS:
		t = Token{
			Type:  BASS,
			Value: strings.ReplaceAll(rawValue, slash, ""),
		}
	}

	l.currentPos += len(rawValue)
	l.lookupPos = l.currentPos

	return t, nil
}

var ErrNoRootToken = errors.New("no ROOT Token found")

func (l *Lexer) readChars() (tokenType, string, error) {
	// try to read root
	if l.currentPos == 0 {
		index := slices.Index(roots, string(l.input[0]))
		if index < 0 {
			return "", "", ErrNoRootToken
		}

		return ROOT, roots[index], nil

	}

	for l.lookupPos < len(l.input) {
		l.lookupPos++
		ch := l.input[l.currentPos:l.lookupPos]

		return l.getToken(ch)
	}

	return END, "", nil
}

func (l *Lexer) getToken(ch string) (tokenType, string, error) {
	switch ch {
	case m:
		// check if it is "m" or "maj"
		if strings.HasPrefix(l.input[l.currentPos:], maj) {
			l.lookupPos += len(maj)
			// check if maj is followed by numbers
			l.takeFollowingNumbers()

			return MAJ, l.input[l.currentPos:l.lookupPos], nil
		}

		l.takeFollowingNumbers()

		return MINOR, l.input[l.currentPos:l.lookupPos], nil
	case sharpSign:
		// check if sharpSign is followed by numbers
		l.takeFollowingNumbers()

		return SHARP, l.input[l.currentPos:l.lookupPos], nil
	case b:
		// check if flat is followed by numbers
		l.takeFollowingNumbers()

		return FLAT, l.input[l.currentPos:l.lookupPos], nil
	case slash:
		// check if flat is followed by numbers
		l.takeFollowingNumbers()

		_, err := strconv.Atoi(strings.ReplaceAll(l.input[l.currentPos:l.lookupPos], slash, ""))
		if err == nil {
			// followed by number
			return ADD, l.input[l.currentPos:l.lookupPos], nil
		}

		// it should appear only at the end of the chord description
		return BASS, l.input[l.currentPos:], nil
	default:
		if unicode.IsDigit(rune(ch[0])) {
			l.takeFollowingNumbers()

			return NUMBER, l.input[l.currentPos:l.lookupPos], nil
		}

		if strings.HasPrefix(l.input[l.currentPos:], dim) {
			l.lookupPos += len(dim)
			l.takeFollowingNumbers()

			return DIM, l.input[l.currentPos:l.lookupPos], nil
		}

		if strings.HasPrefix(l.input[l.currentPos:], aug) {
			l.lookupPos += len(aug)
			l.takeFollowingNumbers()

			return AUG, l.input[l.currentPos:l.lookupPos], nil
		}

		if strings.HasPrefix(l.input[l.currentPos:], add) {
			l.lookupPos += len(add)
			// check if maj is followed by numbers
			l.takeFollowingNumbers()

			return ADD, l.input[l.currentPos:l.lookupPos], nil
		}

		if strings.HasPrefix(l.input[l.currentPos:], sus) {
			l.lookupPos += len(sus)
			// check if maj is followed by numbers
			l.takeFollowingNumbers()

			return SUS, l.input[l.currentPos:l.lookupPos], nil
		}

		return IGNORED, l.input[l.currentPos:l.lookupPos], nil
	}
}

func (l *Lexer) takeFollowingNumbers() {
	if l.lookupPos > len(l.input) {
		l.lookupPos--

		return
	}

	for ; l.lookupPos < len(l.input); l.lookupPos++ {
		if unicode.IsDigit(rune(l.input[l.lookupPos])) {
			continue
		}

		break
	}
}

func (l *Lexer) removeNonNumbers(s string) string {
	// Matches everything except digits
	return re.ReplaceAllString(s, "")
}
