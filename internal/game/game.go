package game

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
