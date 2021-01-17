package tennis

import "fmt"

var lookUp = make(map[int]string, 4)

func init() {
	lookUp[1] = "fifteen"
	lookUp[2] = "thirty"
	lookUp[3] = "forty"
}

type Tennis struct {
	firstPlayerScoreTimes int
}

func (t *Tennis) Score() string {
	if t.firstPlayerScoreTimes > 0 {
		return fmt.Sprintf("%s love", lookUp[t.firstPlayerScoreTimes])
	}
	return "love all"
}

func (t *Tennis) FirstPlayerScore() {
	t.firstPlayerScoreTimes++
}
