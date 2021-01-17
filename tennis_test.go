package tennis_test

import (
	"Edwardz43/tennis"
	"testing"

	"github.com/bmizerany/assert"
)

func TestShouldBeLoveAll(t *testing.T) {
	tenn := new(tennis.Tennis)
	assert.Equal(t, "love all", tenn.Score())
}

func TestShouldBeFifteenLove(t *testing.T) {
	tenn := new(tennis.Tennis)
	tenn.FirstPlayerScore()
	assert.Equal(t, "fifteen love", tenn.Score())
}

func TestShouldBeThirtyLove(t *testing.T) {
	tenn := new(tennis.Tennis)
	addFirstPlayerScore(tenn)
	assert.Equal(t, "thirty love", tenn.Score())
}

func addFirstPlayerScore(tenn *tennis.Tennis) {
	for i := 0; i < 2; i++ {
		tenn.FirstPlayerScore()
	}
}
