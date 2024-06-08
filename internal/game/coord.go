package game

import "fmt"

type Coord struct {
	X int
	Y int
}

func (p Coord) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}
