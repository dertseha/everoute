package universe

import (
	check "gopkg.in/check.v1"
)

type EmptySolarSystemTestSuite struct {
	SolarSystemTestSuite
}

var _ = check.Suite(&EmptySolarSystemTestSuite{SolarSystemTestSuite{createSystem: newSolarSystem}})

func (suite *EmptySolarSystemTestSuite) TestCostsReturnsAnEmptyCostInstance(c *check.C) {
	result := suite.SolarSystem().Costs()

	c.Assert(result, check.DeepEquals, EmptyTravelCostSum())
}
