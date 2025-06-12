package game

import (
	"slices"
	"strings"
)

func directions() [][]int {
	return [][]int{
		{0, -1}, // Up
		{1, 0},  // Right
		{0, 1},  // Down
		{-1, 0}, // Left
	}
}

type Snake struct {
	Body         []Coord // head is at index 0
	Dir          Direction
	Speed        float64
	addAfterMove bool
}

func NewSnake() *Snake {
	return &Snake{
		Body:  []Coord{{DefaultSnakeX, DefaultSnakeY}},
		Dir:   DefaultSnakeDir,
		Speed: DefaultSnakeSpeed,
	}
}

func (s *Snake) Add() {
	s.addAfterMove = true
}

func (s Snake) Head() Coord {
	return s.Body[0]
}

func (s *Snake) NextHead(direction Direction) Coord {
	coordinate := directions()[direction]
	head := s.Body[0]
	return Coord{
		X: head.X + coordinate[0],
		Y: head.Y + coordinate[1],
	}
}

func (snake *Snake) IsValidMove(direction Direction) bool {
	nextHead := snake.NextHead(direction)
	return !snake.Contains(nextHead)
}

func (snake *Snake) Move(direction Direction) error {
	if direction < 0 && direction >= 4 {
		panic("invalid direction")
	}

	if snake.Dir.IsOpposite(direction) {
		return nil
	}

	var prev Coord
	offset := directions()[direction]

	newBody := make([]Coord, len(snake.Body))
	for i, snakeCoordinate := range snake.Body {
		if i == 0 {
			prev = snakeCoordinate
			newBody[i] = Coord{
				snakeCoordinate.X + offset[0],
				snakeCoordinate.Y + offset[1],
			}
		} else {
			tempCoord := snakeCoordinate
			newBody[i] = prev
			prev = tempCoord
		}
	}

	if snake.addAfterMove {
		newBody = append(newBody, prev)
		snake.Speed *= SpeedIncreateRate
		snake.addAfterMove = false
	}

	snake.Body = newBody
	snake.Dir = direction

	return nil
}

func (snake *Snake) Contains(coordinate Coord) bool {
	return slices.Contains(snake.Body, coordinate)
}

func (snake *Snake) BodyAsString() string {
	body := strings.Builder{}
	for _, coordinate := range snake.Body {
		body.WriteString(coordinate.String())
	}
	return body.String()
}
