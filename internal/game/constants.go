package game

type Direction int

const (
	Up = iota
	Right
	Down
	Left
)

const (
	Running = iota
	GameOver
	Paused
)

const (
	defaultSnakeX   = 4
	defaultSnakeY   = 4
	defaultSnakeDir = Right
)

const (
	defaultAppleX = 2
	defaultAppleY = 2
)
