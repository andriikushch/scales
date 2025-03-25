package scales

func parseChord(description string) (Chord, error) {
	p := parser{}

	c, err := p.parse(description)
	if err != nil {
		return Chord{}, err
	}

	return *c, c.finish()
}
