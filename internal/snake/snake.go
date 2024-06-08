package snake

import (
	"errors"
)

const (
	defaultX = 1
	defaultY = 1
)

const (
	Up = iota
	Right
	Down
	Left
)

func dirs() [][]int {
	return [][]int{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Down
	}
}

type Snake struct {
	Body []Coord // head is at index 0
	Dir  int
}

func NewSnake() *Snake {
	return &Snake{
		Body: []Coord{{defaultX, defaultY}},
		Dir:  Right,
	}
}

func (s *Snake) Add() {
	lastCoord := s.Body[len(s.Body)-1]
	s.Body = append(s.Body, Coord{
		X: lastCoord.X,
		Y: lastCoord.Y,
	})
}

func (s *Snake) Move(dir int) error {
	if dir < 0 && dir >= 4 {
		return errors.New("Invalid direction")
	}

	if s.isOppositeDir(dir) {
		return errors.New("Can't move in opposite direction")
	}

	var head Coord
	var prev Coord
	dirCoord := dirs()[dir]

	newBody := make([]Coord, len(s.Body))
	for i, coord := range s.Body {
		if i == 0 {
			head = Coord{
				coord.X + dirCoord[0],
				coord.Y + dirCoord[1],
			}
			newBody[i] = head
		} else {
			prev = newBody[i-1]
			newBody[i] = prev
		}
	}

	s.Body = newBody
	s.Dir = dir

	return nil
}

func (s *Snake) Contains(c Coord) bool {
	for _, coord := range s.Body {
		if coord.X == c.X && coord.Y == c.Y {
			return true
		}
	}
	return false
}

func (s *Snake) isOppositeDir(dir int) bool {
	return (s.Dir == Up && dir == Down) ||
		(s.Dir == Down && dir == Up) ||
		(s.Dir == Left && dir == Right) ||
		(s.Dir == Right && dir == Left)
}
