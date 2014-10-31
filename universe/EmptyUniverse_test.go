package universe

import (
	check "gopkg.in/check.v1"
)

type EmptyUniverseTestSuite struct {
	universe Universe
}

var _ = check.Suite(&EmptyUniverseTestSuite{})

func (suite *EmptyUniverseTestSuite) SetUpTest(c *check.C) {
	suite.universe = New()
}

func (suite *EmptyUniverseTestSuite) TestExtendReturnsNewBuilder(c *check.C) {
	builder := suite.universe.Extend()

	c.Assert(builder, check.NotNil)
}

func (suite *EmptyUniverseTestSuite) TestHasSolarSystemReturnsFalse(c *check.C) {
	c.Assert(suite.universe.HasSolarSystem(Id(1234)), check.Equals, false)
}

func (suite *EmptyUniverseTestSuite) TestSolarSystemPanics(c *check.C) {
	c.Assert(func() { suite.universe.SolarSystem(Id(1234)) }, check.Panics, "SolarSystem with ID <1234> not found")
}

func (suite *EmptyUniverseTestSuite) TestSolarSystemIdsReturnsEmptySlice(c *check.C) {
	c.Assert(suite.universe.SolarSystemIds(), check.HasLen, 0)
}
