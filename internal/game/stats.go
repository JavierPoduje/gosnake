package game

type Stats struct {
	EatenApples int
}

func NewStats() *Stats {
	return &Stats{
		EatenApples: 0,
	}
}

func (s *Stats) EatApple() {
	s.EatenApples++
}
