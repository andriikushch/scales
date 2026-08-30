package scales

func parseChord(description string) (Chord, error) {
	p := parser{}

	c, err := p.parse(description)
	if err != nil {
		return Chord{}, err
	}

	err = c.finish()
	if err != nil {
		return Chord{}, err
	}

	return *c, nil
}
