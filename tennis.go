package tennis

type Tennis struct {
	firstPlayerScoreTimes int
}

func (t *Tennis) Score() string {
	if t.firstPlayerScoreTimes == 1 {
		return "fifteen love"
	}
	return "love all"
}

func (t *Tennis) FirstPlayerScore() {
	t.firstPlayerScoreTimes++
}
