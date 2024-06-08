package game

type Direction int

func (d Direction) String() string {
	return [...]string{"Up", "Right", "Down", "Left"}[d]
}
