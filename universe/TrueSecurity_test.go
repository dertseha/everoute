package universe

import (
	check "gopkg.in/check.v1"
)

type TrueSecurityTestSuite struct{}

var _ = check.Suite(&TrueSecurityTestSuite{})

func (suite *TrueSecurityTestSuite) TestIsHighSecReturnsFalseForRoundedBelowPointFive(c *check.C) {
	sec := TrueSecurity(0.4)

	c.Assert(sec.IsHighSec(), check.Equals, false)
}

func (suite *TrueSecurityTestSuite) TestIsHighSecReturnsTrueForRoundedPointFive(c *check.C) {
	sec := TrueSecurity(0.5)

	c.Assert(sec.IsHighSec(), check.Equals, true)
}

func (suite *TrueSecurityTestSuite) TestIsHighSecReturnsTrueForRoundedOne(c *check.C) {
	sec := TrueSecurity(1.0)

	c.Assert(sec.IsHighSec(), check.Equals, true)
}

func (suite *TrueSecurityTestSuite) TestRoundedReturnsZeroForValueBelowZero(c *check.C) {
	sec := TrueSecurity(-0.4)

	c.Assert(sec.Rounded(), check.Equals, 0.0)
}

func (suite *TrueSecurityTestSuite) TestRoundedReturnsPoint5ForPoint45(c *check.C) {
	sec := TrueSecurity(0.45)

	c.Assert(sec.Rounded(), check.Equals, 0.5)
}

func (suite *TrueSecurityTestSuite) TestRoundedReturnsPoint4ForPoint449(c *check.C) {
	sec := TrueSecurity(0.449)

	c.Assert(sec.Rounded(), check.Equals, 0.4)
}
