package tennis

type tennis struct {
	firstPlayerScoreTimes int
}

func (t *tennis) FirstPlayerScore() {
	t.firstPlayerScoreTimes++
}

func (t *tennis) Score() interface{} {
	if t.firstPlayerScoreTimes == 1 {
		return "fifteen love"
	}
	return "love all"
}

func New() *tennis {
	return &tennis{firstPlayerScoreTimes: 0}
}
