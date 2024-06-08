package game

type Apple struct {
	X int
	Y int
}

func NewApple() *Apple {
	return &Apple{
		X: defaultAppleX,
		Y: defaultAppleY,
	}
}
