package game

func (direction Direction) String() string {
	return [...]string{"Up", "Right", "Down", "Left"}[direction]
}
