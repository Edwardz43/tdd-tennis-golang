package tennis_test

import (
	"testing"

	"github.com/bmizerany/assert"
)

type Tennis struct {
}

func (t *Tennis) score() string {
	return "love all"
}

func TestShouldBeLoveAll(t *testing.T) {
	tenn := new(Tennis)
	assert.Equal(t, "love all", tenn.score())
}
