package universe

import (
	check "gopkg.in/check.v1"
)

type ExtendedSolarSystemTestSuite struct {
	SolarSystemTestSuite
}

var _ = check.Suite(&ExtendedSolarSystemTestSuite{SolarSystemTestSuite{
	createSystem: func(id Id, constellationId Id, regionId Id, galaxyId GalaxyId, location Location, security TrueSecurity) SolarSystem {
		baseSystem := newSolarSystem(id, constellationId, regionId, galaxyId, location, security)

		// build first extension
		extensionData1 := newSolarSystemExtensionData(baseSystem)

		extensionData1.costs = extensionData1.costs.Add(SingleTravelCostSum(AddingTravelCost("test", 10.0)))

		builder1a := newJumpBuilder("type1", Id(1001))
		extensionData1.jumps = append(extensionData1.jumps, builder1a.Build())
		builder1b := newJumpBuilder("type1", Id(1002))
		extensionData1.jumpBuilder = append(extensionData1.jumpBuilder, builder1b)
		builder1c := newJumpBuilder("type2", Id(2001))
		extensionData1.jumps = append(extensionData1.jumps, builder1c.Build())

		extendedSystem1 := extensionData1.solarSystem()

		// build second extension
		extensionData2 := newSolarSystemExtensionData(extendedSystem1)

		extensionData2.costs = extensionData2.costs.Add(SingleTravelCostSum(AddingTravelCost("test", 5.0)))

		builder2a := newJumpBuilder("type2", Id(2002))
		extensionData2.jumps = append(extensionData2.jumps, builder2a.Build())

		return extensionData2.solarSystem()
	}}})

func (suite *ExtendedSolarSystemTestSuite) TestCostsContainsAddedAndCombinedTravelCost(c *check.C) {
	result := suite.SolarSystem().Costs().Cost(AddingTravelCost("test", 0.0))

	c.Assert(result.Value(), check.Equals, 15.0)
}

func (suite *ExtendedSolarSystemTestSuite) TestJumpsReturnsResultOfBothBuilderAndLiteralJumps(c *check.C) {
	result := suite.SolarSystem().Jumps("type1")

	c.Assert(result, check.HasLen, 2)
}

func (suite *ExtendedSolarSystemTestSuite) TestJumpsReturnsCombinedJumps(c *check.C) {
	result := suite.SolarSystem().Jumps("type2")

	c.Assert(result, check.HasLen, 2)
}
