package game

import "strconv"

type Stats struct {
	EatenApples int
	Score       float64
}

func NewStats() *Stats {
	return &Stats{
		EatenApples: 0,
		Score:       0,
	}
}

func (s *Stats) EatApple() {
	s.EatenApples++
}

func (s *Stats) UpdateScore(speed float64) {
	s.Score = float64(s.EatenApples) * speed
}

func (s Stats) ScoreAsString() string {
	return strconv.FormatFloat(s.Score, 'f', 2, 64)
}

func (s Stats) RoundedScoreAsString() string {
	return strconv.FormatFloat(s.Score, 'f', 0, 64)
}
