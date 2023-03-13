package ikayaku

type Score struct {
	score int
}

func NewScore() *Score {
	return &Score{
		score: 0,
	}
}

func (s *Score) GetScore() int {
	return s.score
}

func (s *Score) AddScore(point int) {
	s.score += point
}

func (s *Score) SubScore(point int) {
	if s.score <= 0 {
		return
	}
	s.score -= point
}
