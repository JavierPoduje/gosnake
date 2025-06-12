package game

import (
	"slices"
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
	Body                []Coord // head is at index 0
	Dir                 Direction
	Speed               float64
	increaseAfterMoving bool
}

func NewSnake() *Snake {
	return &Snake{
		Body:  []Coord{{DefaultSnakeX, DefaultSnakeY}},
		Dir:   DefaultSnakeDir,
		Speed: DefaultSnakeSpeed,
	}
}

func (snake *Snake) Add() {
	snake.increaseAfterMoving = true
}

func (snake Snake) Head() Coord {
	return snake.Body[0]
}

func (snake *Snake) NextHead(direction Direction) Coord {
	offset := directions()[direction]
	return Coord{
		X: snake.Head().X + offset[0],
		Y: snake.Head().Y + offset[1],
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

	var previousCoordinate Coord
	offset := directions()[direction]

	newBody := make([]Coord, len(snake.Body))
	for i, snakeCoordinate := range snake.Body {
		if i == 0 {
			previousCoordinate = snakeCoordinate
			newBody[i] = Coord{
				snakeCoordinate.X + offset[0],
				snakeCoordinate.Y + offset[1],
			}
		} else {
			tempCoord := snakeCoordinate
			newBody[i] = previousCoordinate
			previousCoordinate = tempCoord
		}
	}

	if snake.increaseAfterMoving {
		newBody = append(newBody, previousCoordinate)
		snake.Speed *= SpeedIncreateRate
		snake.increaseAfterMoving = false
	}

	snake.Body = newBody
	snake.Dir = direction

	return nil
}

func (snake *Snake) Contains(coordinate Coord) bool {
	return slices.Contains(snake.Body, coordinate)
}

func (snake *Snake) IsHead(coordinate Coord) bool {
	return snake.Head().Equals(coordinate)
}
