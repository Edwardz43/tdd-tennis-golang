package tennis_test

import (
	"Edwardz43/tennis"
	"testing"

	"github.com/bmizerany/assert"
)

var tenn *tennis.Tennis

func beforeEach() {
	tenn = &tennis.Tennis{
		FirstPlayerName:  "Tom",
		SecondPlayerName: "Elsa",
	}
}

func addFirstPlayerScore(times int) {
	for i := 0; i < times; i++ {
		tenn.FirstPlayerScore()
	}
}

func addSecondPlayerScore(times int) {
	for i := 0; i < times; i++ {
		tenn.SecondPlayerScore()
	}
}

func deuce() {
	addFirstPlayerScore(3)
	addSecondPlayerScore(3)
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
	addFirstPlayerScore(2)
	assert.Equal(t, "thirty love", tenn.Score())
}

func TestShouldBeFortyLove(t *testing.T) {
	beforeEach()
	addFirstPlayerScore(3)
	assert.Equal(t, "forty love", tenn.Score())
}

func TestShouldBeLoveFifteen(t *testing.T) {
	beforeEach()
	tenn.SecondPlayerScore()
	assert.Equal(t, "love fifteen", tenn.Score())
}

func TestShouldBeLoveThirty(t *testing.T) {
	beforeEach()
	addSecondPlayerScore(2)
	assert.Equal(t, "love thirty", tenn.Score())
}

func TestShouldBeLoveForty(t *testing.T) {
	beforeEach()
	addSecondPlayerScore(3)
	assert.Equal(t, "love forty", tenn.Score())
}

func TestShouldBeFifteenAll(t *testing.T) {
	beforeEach()
	tenn.FirstPlayerScore()
	tenn.SecondPlayerScore()
	assert.Equal(t, "fifteen all", tenn.Score())
}

func TestShouldBeThirtyAll(t *testing.T) {
	beforeEach()
	addFirstPlayerScore(2)
	addSecondPlayerScore(2)
	assert.Equal(t, "thirty all", tenn.Score())
}

func TestShouldBeDeuce(t *testing.T) {
	beforeEach()
	deuce()
	assert.Equal(t, "deuce", tenn.Score())
}

func TestShouldBeTomAdv(t *testing.T) {
	beforeEach()
	deuce()
	tenn.FirstPlayerScore()
	assert.Equal(t, "Tom adv", tenn.Score())
}

func TestShouldBeElsaAdv(t *testing.T) {
	beforeEach()
	deuce()
	tenn.SecondPlayerScore()
	assert.Equal(t, "Elsa adv", tenn.Score())
}

func TestShouldBeElsaWin(t *testing.T) {
	beforeEach()
	deuce()
	addSecondPlayerScore(2)
	assert.Equal(t, "Elsa win", tenn.Score())
}

func TestShouldBeTomWin(t *testing.T) {
	beforeEach()
	deuce()
	addFirstPlayerScore(2)
	assert.Equal(t, "Tom win", tenn.Score())
}
