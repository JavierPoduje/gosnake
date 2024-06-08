package game

import "log"

type Game struct {
	Snake  *Snake
	Canvas *Canvas
	Apple  *Apple

	NextMove Direction
}

func NewGame(width, height int) *Game {
	return &Game{
		Snake:  NewSnake(),
		Canvas: NewCanvas(width, height),
		Apple:  NewApple(),
	}
}

func (m Game) UpdateSnake(dir Direction) {
	if !m.snakeCanMove(dir) {
		return
	}

	m.NextMove = dir
	err := m.Snake.Move(m.NextMove)
	if err != nil {
		log.Fatalf("Invalid m.NextMove: %v", err)
	}
}

func (m Game) snakeCanMove(dir Direction) bool {
	nextHead := m.Snake.NextHead(dir)
	return m.Canvas.InBounds(nextHead)
}

