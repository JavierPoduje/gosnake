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
	DefaultSnakeX     = 4
	DefaultSnakeY     = 4
	DefaultSnakeDir   = Right
	DefaultSnakeSpeed = float64(3)
	SpeedIncreateRate = 1.062
)

const (
	DefaultAppleX = 2
	DefaultAppleY = 2
)
