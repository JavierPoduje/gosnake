package game

import (
	"log"
	"math/rand/v2"
)

type Game struct {
	Snake  *Snake
	Canvas *Canvas
	Apple  *Coord

	NextMove Direction
}

func NewGame(width, height int) *Game {
	return &Game{
		Snake:  NewSnake(),
		Canvas: NewCanvas(width, height),
		Apple:  &Coord{X: defaultAppleX, Y: defaultAppleY},
	}
}

func (g *Game) Tick(dir Direction) {
	if !g.canSnakeMove(dir) {
		return
	}

	g.NextMove = dir
	err := g.Snake.Move(g.NextMove)
	if err != nil {
		log.Fatalf("Invalid m.NextMove: %v", err)
	}

	head := g.Snake.Body[0]

	if head == *g.Apple {
		g.Snake.Add()

		g.Apple = g.getRandApple()

		for g.Snake.Contains(*g.Apple) {
			g.Apple = g.getRandApple()
		}
	}
}

func (m Game) canSnakeMove(dir Direction) bool {
	nextHead := m.Snake.NextHead(dir)
	return m.Canvas.InBounds(nextHead)
}

func (g Game) getRandApple() *Coord {
	width := g.Canvas.Width
	height := g.Canvas.Height

	X := rand.IntN(width)
	Y := rand.IntN(height)

	return &Coord{X: X, Y: Y}
}
