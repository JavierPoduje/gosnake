package game

// score = length (eaten apples) * speed
// score increases each time the snake eats an apple
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
