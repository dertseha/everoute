package universe

import (
	check "gopkg.in/check.v1"
)

type TravelCostSumTestSuite struct{}

var _ = check.Suite(&TravelCostSumTestSuite{})

func (suite *TravelCostSumTestSuite) TestAddTakesCostsNotKnownInSecondDirectly(c *check.C) {
	sum1 := SingleTravelCostSum(AddingTravelCost("test1", 10.0))
	sum2 := SingleTravelCostSum(AddingTravelCost("test2", 20.0))
	result := sum1.Add(sum2)

	c.Assert(result.Cost(AddingTravelCost("test1", 0.0)).Value(), check.Equals, 10.0)
}
