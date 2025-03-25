package scales

type contextNotes interface {
	GetNotes() []Note
}

func parseChord(description string, context contextNotes) (Chord, error) {
	p := parser{}

	c, err := p.parse(description)
	if err != nil {
		return Chord{}, err
	}

	err = c.finish(context)
	if err != nil {
		return Chord{}, err
	}

	return *c, nil
}
