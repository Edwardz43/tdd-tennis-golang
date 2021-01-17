package tennis

import "fmt"

var lookUp = make(map[int]string, 4)

func init() {
	lookUp[0] = "love"
	lookUp[1] = "fifteen"
	lookUp[2] = "thirty"
	lookUp[3] = "forty"
}

type Tennis struct {
	firstPlayerScoreTimes  int
	secondPlayerScoreTimes int
	FirstPlayerName        string
	SecondPlayerName       string
}

func (t *Tennis) Score() string {
	if t.samScore() {
		if t.firstPlayerScoreTimes >= 3 {
			return "deuce"
		}
		return fmt.Sprintf("%s all", lookUp[t.firstPlayerScoreTimes])
	}
	if t.isReadyForGamePoint() {
		if t.scoreGapAbsEqualOne() {
			return fmt.Sprintf("%s adv", t.advPlayerName())
		}
		return fmt.Sprintf("%s win", t.advPlayerName())
	}
	return t.lookUpScore()
}

func (t *Tennis) lookUpScore() string {
	return fmt.Sprintf("%s %s", lookUp[t.firstPlayerScoreTimes], lookUp[t.secondPlayerScoreTimes])
}

func (t *Tennis) scoreGapAbsEqualOne() bool {
	return t.firstPlayerScoreTimes-t.secondPlayerScoreTimes == 1 ||
		t.firstPlayerScoreTimes-t.secondPlayerScoreTimes == -1
}

func (t *Tennis) advPlayerName() string {
	var advPlayer string
	if t.firstPlayerScoreTimes > t.secondPlayerScoreTimes {
		advPlayer = t.FirstPlayerName
	} else {
		advPlayer = t.SecondPlayerName
	}
	return advPlayer
}
func (t *Tennis) isReadyForGamePoint() bool {
	return t.firstPlayerScoreTimes >= 3 && t.secondPlayerScoreTimes >= 3
}

func (t *Tennis) samScore() bool {
	return t.secondPlayerScoreTimes == t.firstPlayerScoreTimes
}

func (t *Tennis) FirstPlayerScore() {
	t.firstPlayerScoreTimes++
}

func (t *Tennis) SecondPlayerScore() {
	t.secondPlayerScoreTimes++
}
