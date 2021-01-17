package tennis_test

import (
	"Edwardz43/tennis"
	"testing"

	"github.com/bmizerany/assert"
)

var tenn *tennis.Tennis

func beforeEach() {
	tenn = new(tennis.Tennis)
}

func addFirstPlayerScore() {
	for i := 0; i < 2; i++ {
		tenn.FirstPlayerScore()
	}
}

func TestShouldBeLoveAll(t *testing.T) {
	beforeEach()
	assert.Equal(t, "love all", tenn.Score())
}

func TestShouldBeFifteenLove(t *testing.T) {
	beforeEach()
	tenn.FirstPlayerScore()
	assert.Equal(t, "fifteen love", tenn.Score())
}

func TestShouldBeThirtyLove(t *testing.T) {
	beforeEach()
	addFirstPlayerScore()
	assert.Equal(t, "thirty love", tenn.Score())
}
