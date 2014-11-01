package universe

import (
	check "gopkg.in/check.v1"
)

type ExtendedUniverseTestSuite struct {
	UniverseTestSuite
}

var _ = check.Suite(&ExtendedUniverseTestSuite{UniverseTestSuite{
	createUniverse: func() Universe {
		baseUniverse := New()

		builder1 := baseUniverse.Extend()

		builder1.AddSolarSystem(Id(2), Id(10), Id(100), NewEdenId, AnyLocation(), TrueSecurity(0.0))

		verse2 := builder1.Build()
		builder2 := verse2.Extend()

		builder2.AddSolarSystem(Id(1), Id(10), Id(100), NewEdenId, AnyLocation(), TrueSecurity(0.0))

		return builder2.Build()
	}}})

func (suite *ExtendedUniverseTestSuite) TestSolarSystemIdsReturnsCombinedResults(c *check.C) {
	c.Assert(suite.Universe().SolarSystemIds(), check.HasLen, 2)
}

func (suite *ExtendedUniverseTestSuite) TestSolarSystemIdsReturnsIdsSorted(c *check.C) {
	c.Assert(suite.Universe().SolarSystemIds(), check.DeepEquals, []Id{1, 2})
}

func (suite *ExtendedUniverseTestSuite) TestSolarSystemReturnsSystemFromBase(c *check.C) {
	result := suite.Universe().SolarSystem(2)

	c.Assert(result, check.NotNil)
}

func (suite *ExtendedUniverseTestSuite) TestSolarSystemReturnsSystemFromExtension(c *check.C) {
	result := suite.Universe().SolarSystem(1)

	c.Assert(result, check.NotNil)
}
