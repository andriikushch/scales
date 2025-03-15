package scales

func parseChord(description string) (chord, error) {
	p := parser{}

	c, err := p.parse(description)
	if err != nil {
		return chord{}, err
	}

	return *c, c.finish()
}
