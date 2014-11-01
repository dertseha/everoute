package universe

import (
	check "gopkg.in/check.v1"
)

type UniverseTestSuite struct {
	createUniverse func() Universe

	universe Universe
}

func (suite *UniverseTestSuite) Universe() Universe {
	return suite.universe
}

func (suite *UniverseTestSuite) SetUpTest(c *check.C) {
	suite.universe = suite.createUniverse()
}

func (suite *UniverseTestSuite) TestExtendReturnsNewBuilder(c *check.C) {
	builder := suite.universe.Extend()

	c.Assert(builder, check.NotNil)
}

func (suite *UniverseTestSuite) TestHasSolarSystemReturnsFalseForUnknownSolarSystem(c *check.C) {
	c.Assert(suite.universe.HasSolarSystem(Id(10203040)), check.Equals, false)
}

func (suite *UniverseTestSuite) TestSolarSystemPanicsForUnknownSolarSystem(c *check.C) {
	c.Assert(func() { suite.universe.SolarSystem(Id(10203040)) }, check.Panics, "SolarSystem with ID <10203040> not found")
}
