package universe

import (
	check "gopkg.in/check.v1"
)

type AddingTravelCostTestSuite struct{}

var _ = check.Suite(&AddingTravelCostTestSuite{})

func (suite *AddingTravelCostTestSuite) TestTypeReturnsType(c *check.C) {
	costType := "type1"
	cost := AddingTravelCost(costType, 0.0)
	result := cost.Type()

	c.Assert(result, check.Equals, costType)
}

func (suite *AddingTravelCostTestSuite) TestValueReturnsValue(c *check.C) {
	value := 1234.0
	cost := AddingTravelCost("test2", value)
	result := cost.Value()

	c.Assert(result, check.Equals, value)
}

func (suite *AddingTravelCostTestSuite) TestJoinPanicsForMismatchedType(c *check.C) {
	cost1 := AddingTravelCost("test3", 0.0)
	cost2 := AddingTravelCost("test1", 0.0)

	c.Assert(func() { cost1.Join(cost2) }, check.Panics, "Cost type mismatch")
}

func (suite *AddingTravelCostTestSuite) TestJoinReturnsCombinedCost(c *check.C) {
	cost1 := AddingTravelCost("test4", 10.0)
	cost2 := AddingTravelCost("test4", 5.0)
	result := cost1.Join(cost2)

	c.Assert(result.Value(), check.Equals, 15.0)
}

func (suite *AddingTravelCostTestSuite) TestJoinReturnsCombinedCostWithSameType(c *check.C) {
	cost1 := AddingTravelCost("test4", 10.0)
	cost2 := AddingTravelCost("test4", 5.0)
	result := cost1.Join(cost2)

	c.Assert(result.Type(), check.Equals, "test4")
}
