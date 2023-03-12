package ikayaku

const initTime = 60

type Timer struct {
	time int
}

func NewTimer() *Timer {
	return &Timer{
		time: initTime,
	}
}

func (t *Timer) GetTime() int {
	return t.time
}

func (t *Timer) AddTime(time int) {
	t.time += time
}

func (t *Timer) SubTime(time int) {
	t.time -= time
}
