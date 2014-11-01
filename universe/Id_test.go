package universe

import (
	check "gopkg.in/check.v1"
)

type IdOrderTestSuite struct{}

var _ = check.Suite(&IdOrderTestSuite{})

func (suite *IdOrderTestSuite) TestLenReturnsLength(c *check.C) {
	order := IdOrder{Id(1), Id(2), Id(3)}

	c.Assert(order.Len(), check.Equals, 3)
}

func (suite *IdOrderTestSuite) TestLessReturnsResultOfComparison(c *check.C) {
	order := IdOrder{Id(1), Id(2), Id(3)}

	c.Assert(order.Less(0, 2), check.Equals, true)
}

func (suite *IdOrderTestSuite) TestSwapMovesItems(c *check.C) {
	order := IdOrder{Id(1), Id(2), Id(3)}

	order.Swap(2, 0)

	c.Assert(order, check.DeepEquals, IdOrder{Id(3), Id(2), Id(1)})
}
