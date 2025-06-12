package game

import (
	"errors"
	"strings"
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
	Body  []Coord // head is at index 0
	Dir   Direction
	Speed float64
	add   bool
}

func NewSnake() *Snake {
	return &Snake{
		Body:  []Coord{{DefaultSnakeX, DefaultSnakeY}},
		Dir:   DefaultSnakeDir,
		Speed: DefaultSnakeSpeed,
	}
}

func (s *Snake) Add() {
	s.add = true
}

func (s Snake) Head() Coord {
	return s.Body[0]
}

func (s *Snake) NextHead(dir Direction) Coord {
	dirCoord := dirs()[dir]
	head := s.Body[0]
	return Coord{
		X: head.X + dirCoord[0],
		Y: head.Y + dirCoord[1],
	}
}

func (s *Snake) IsValidMove(dir Direction) bool {
	nextHead := s.NextHead(dir)
	return !s.Contains(nextHead)
}

func (s *Snake) Move(dir Direction) error {
	if dir < 0 && dir >= 4 {
		return errors.New("Invalid direction")
	}

	if s.Dir.IsOpposite(dir) {
		return nil
	}

	var prev Coord
	dirCoord := dirs()[dir]

	newBody := make([]Coord, len(s.Body))
	for i, coord := range s.Body {
		if i == 0 {
			prev = coord
			newBody[i] = Coord{
				coord.X + dirCoord[0],
				coord.Y + dirCoord[1],
			}
		} else {
			tempCoord := coord
			newBody[i] = prev
			prev = tempCoord
		}
	}

	if s.add {
		newBody = append(newBody, prev)
		s.Speed *= SpeedIncreateRate
		s.add = false
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

func (s *Snake) BodyAsString() string {
	body := strings.Builder{}
	for _, coord := range s.Body {
		body.WriteString(coord.String())
	}
	return body.String()
}
