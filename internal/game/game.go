package game

import (
	"log"
	"math/rand/v2"
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
		Apple:    &Coord{X: DefaultAppleX, Y: DefaultAppleY},
		Canvas:   NewCanvas(width, height),
		NextMove: Up,
		Snake:    NewSnake(),
		State:    Running,
		Stats:    NewStats(),
	}
}

func (game *Game) Tick(direction Direction) {
	if !game.nextMoveIsValid(direction) {
		game.State = GameOver
		return
	}

	game.NextMove = direction
	err := game.Snake.Move(game.NextMove)
	if err != nil {
		log.Fatalf("Invalid m.NextMove: %v", err)
	}

	if game.Snake.Body[0] == *game.Apple {
		game.eatApple()
	}
}

func (game Game) nextMoveIsValid(dir Direction) bool {
	nextHead := game.Snake.NextHead(dir)
	return game.Snake.IsValidMove(dir) && game.Canvas.InBounds(nextHead)
}

func (game Game) getRandApple() *Coord {
	width := game.Canvas.Width
	height := game.Canvas.Height

	X := rand.IntN(width)
	Y := rand.IntN(height)

	return &Coord{X: X, Y: Y}
}

func (game *Game) eatApple() {
	game.Snake.Add()
	game.Stats.EatApple()
	game.Stats.UpdateScore(game.Snake.Speed)

	game.Apple = game.getRandApple()
	for game.Snake.Contains(*game.Apple) {
		game.Apple = game.getRandApple()
	}
}
