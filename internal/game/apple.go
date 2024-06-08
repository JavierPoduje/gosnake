package game

type Apple struct {
	X int
	Y int
}

func NewApple() *Apple {
	return &Apple{
		X: 2,
		Y: 2,
	}
}
