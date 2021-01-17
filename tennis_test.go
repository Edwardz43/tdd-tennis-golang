package tennis_test

import (
	"Edwardz43/tennis"
	"testing"

	"github.com/bmizerany/assert"
)

var tenn *tennis.Tennis

func init() {
	tenn = new(tennis.Tennis)
}

func addFirstPlayerScore() {
	for i := 0; i < 2; i++ {
		tenn.FirstPlayerScore()
	}
}

func TestShouldBeLoveAll(t *testing.T) {
	assert.Equal(t, "love all", tenn.Score())
}

func TestShouldBeFifteenLove(t *testing.T) {
	tenn.FirstPlayerScore()
	assert.Equal(t, "fifteen love", tenn.Score())
}

func TestShouldBeThirtyLove(t *testing.T) {
	addFirstPlayerScore()
	assert.Equal(t, "thirty love", tenn.Score())
}
