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
		extension1 := &SolarSystemExtension{newSolarSystemExtensionData(baseSystem)}
		extension1.AddCost(AddingTravelCost("test", 10.0))
		extension1.AddJump(newJumpBuilder("type1", Id(1001)).Build())
		extension1.BuildJump("type1", Id(1002))
		extension1.BuildJump("type2", Id(2001))

		extendedSystem1 := extension1.data.solarSystem()

		// build second extension
		extension2 := &SolarSystemExtension{newSolarSystemExtensionData(extendedSystem1)}
		extension2.AddCost(AddingTravelCost("test", 5.0))
		extension2.BuildJump("type2", Id(2002))

		return extension2.data.solarSystem()
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
