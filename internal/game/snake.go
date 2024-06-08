package game

import (
	"errors"
)

func dirs() [][]int {
	return [][]int{
		{0, -1}, // Up
		{1, 0},  // Right
		{0, 1},  // Down
		{-1, 0}, // Left
	}
}

type Snake struct {
	Body []Coord // head is at index 0
	Dir  int
}

func NewSnake() *Snake {
	return &Snake{
		Body: []Coord{{defaultSnakeX, defaultSnakeY}},
		Dir:  defaultSnakeDir,
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
		return nil
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
		if coord == c {
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
