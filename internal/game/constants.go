package game

type Direction int

const (
	Up = iota
	Right
	Down
	Left
)

func (d Direction) IsOpposite(other Direction) bool {
	return (d == Up && other == Down) ||
		(d == Down && other == Up) ||
		(d == Left && other == Right) ||
		(d == Right && other == Left)
}

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
