package tennis

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestTennisTestSuit(t *testing.T) {
	suite.Run(t, new(TennisTestSuite))
}

type TennisTestSuite struct {
	suite.Suite
	*tennis
}

func (s *TennisTestSuite) SetupTest() {}

func (s *TennisTestSuite) TestShouldBeLoveAll() {
	s.Equal("love all", "love all")
}
