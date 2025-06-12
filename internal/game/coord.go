package game

import "fmt"

type Coord struct {
	X int
	Y int
}

func (coord Coord) Equals(other Coord) bool {
	return coord.X == other.X && coord.Y == other.Y
}

func (p Coord) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}
