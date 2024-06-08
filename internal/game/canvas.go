package game

type Canvas struct {
	Width  int
	Height int
}

func NewCanvas(width, height int) *Canvas {
	return &Canvas{
		Width:  width,
		Height: height,
	}
}

func (c *Canvas) InBounds(coord Coord) bool {
	return coord.X >= 0 && coord.X < c.Width &&
		coord.Y >= 0 && coord.Y < c.Height
}
