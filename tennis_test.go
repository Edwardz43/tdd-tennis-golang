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
