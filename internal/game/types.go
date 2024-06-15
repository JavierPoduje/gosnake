package game

func (d Direction) String() string {
	return [...]string{"Up", "Right", "Down", "Left"}[d]
}
