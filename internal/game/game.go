package game

import (
	"log"
	"math/rand/v2"
	"strings"
)

type Game struct {
	Snake  *Snake
	Canvas *Canvas
	Apple  *Coord
	Stats  *Stats

	NextMove Direction
	State    int
}

func NewGame(width, height int) *Game {
	return &Game{
		Snake:    NewSnake(),
		Canvas:   NewCanvas(width, height),
		Apple:    &Coord{X: DefaultAppleX, Y: DefaultAppleY},
		NextMove: Up,
		State:    Running,
		Stats:    NewStats(),
	}
}

func (g *Game) Tick(dir Direction) {
	if !g.canSnakeMove(dir) {
		g.State = GameOver
		return
	}

	g.NextMove = dir
	err := g.Snake.Move(g.NextMove)
	if err != nil {
		log.Fatalf("Invalid m.NextMove: %v", err)
	}

	if g.Snake.Body[0] == *g.Apple {
		g.eatApple()
	}
}

func (m Game) canSnakeMove(dir Direction) bool {
	nextHead := m.Snake.NextHead(dir)
	return m.Snake.IsValidMove(dir) && m.Canvas.InBounds(nextHead)
}

func (g Game) getRandApple() *Coord {
	width := g.Canvas.Width
	height := g.Canvas.Height

	X := rand.IntN(width)
	Y := rand.IntN(height)

	return &Coord{X: X, Y: Y}
}

func (g *Game) eatApple() {
	g.Snake.Add()
	g.Stats.EatApple()

	g.Apple = g.getRandApple()

	for g.Snake.Contains(*g.Apple) {
		g.Apple = g.getRandApple()
	}

	snakeBody := strings.Builder{}
	for _, coord := range g.Snake.Body {
		snakeBody.WriteString(coord.String())
	}
}
