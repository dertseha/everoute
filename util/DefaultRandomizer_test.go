package util

import (
	check "gopkg.in/check.v1"
)

type DefaultRandomizerTestSuite struct {
	rand Randomizer
}

var _ = check.Suite(&DefaultRandomizerTestSuite{})

func (suite *DefaultRandomizerTestSuite) SetUpTest(c *check.C) {
	suite.rand = DefaultRandomizer()
}

func (suite *DefaultRandomizerTestSuite) TestIndexReturnsMinusOneWhenLenIsZero(c *check.C) {
	result := suite.rand.Index(0)

	c.Assert(result, check.Equals, -1)
}

func (suite *DefaultRandomizerTestSuite) TestIndexReturnsNoNegativeNumber(c *check.C) {
	negative := false

	suite.sample(100, func(value int) {
		if value < 0 {
			negative = true
		}
	})

	c.Assert(negative, check.Equals, false)
}

func (suite *DefaultRandomizerTestSuite) TestIndexReturnsNoNumberAboveOrEqualLimit(c *check.C) {
	limit := 100
	overflow := false

	suite.sample(limit, func(value int) {
		if value >= limit {
			overflow = true
		}
	})

	c.Assert(overflow, check.Equals, false)
}

func (suite *DefaultRandomizerTestSuite) TestIndexReturnsDifferentResultsWhenCalledSeveralTimes(c *check.C) {
	limit := 200
	first := suite.rand.Index(limit)
	unique := true

	suite.sample(limit, func(value int) {
		if value != first {
			unique = false
		}
	})

	c.Assert(unique, check.Equals, false)
}

func (suite *DefaultRandomizerTestSuite) sample(limit int, test func(int)) {
	tries := 1000

	for i := 0; i < tries; i++ {
		value := suite.rand.Index(limit)
		test(value)
	}
}
