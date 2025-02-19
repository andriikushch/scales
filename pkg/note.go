package scales

type Note struct {
	Name string
}

func NewNote(name string) Note {
	return Note{
		Name: name,
	}
}
