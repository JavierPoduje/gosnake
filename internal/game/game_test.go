package game

import (
	"testing"
)

func TestGame_CrushSnakeAgainstWall(t *testing.T) {
	g := NewGame(10, 10)
	g.Snake.Body = []Coord{
		{0, 0},
		{0, 1},
		{0, 2},
	}
	g.NextMove = Up

	g.Tick(Up)

	if g.State != GameOver {
		t.Errorf("Expected %v but got %v", GameOver, g.State)
	}
}

func TestGame_SnakeHitsItSelf(t *testing.T) {
	game := NewGame(10, 10)
	game.Snake.Body = []Coord{
		{0, 1},
		{1, 1},
		{1, 0},
		{0, 0},
	}
	game.NextMove = Up

	game.Tick(Up)

	if game.State != GameOver {
		t.Errorf("Expected %v but got %v", GameOver, game.State)
	}
}
