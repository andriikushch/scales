package internal

type tokenType string

const (
	ROOT    tokenType = "ROOT"
	MAJ     tokenType = "MAJ"
	AUG     tokenType = "AUG"
	DIM     tokenType = "DIM"
	ADD     tokenType = "ADD"
	SUS     tokenType = "SUS"
	FLAT    tokenType = "FLAT"
	SHARP   tokenType = "SHARP"
	MINOR   tokenType = "MINOR"
	END     tokenType = "END"
	NUMBER  tokenType = "NUMBER"
	BASS    tokenType = "BASS"
	IGNORED tokenType = "IGNORED"
)

type Token struct {
	Type  tokenType
	Value string
}
