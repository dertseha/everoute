package universe

import (
	check "gopkg.in/check.v1"
)

type UniverseBuilderTestSuite struct {
	builder *UniverseBuilder
}

var _ = check.Suite(&UniverseBuilderTestSuite{})

func (suite *UniverseBuilderTestSuite) SetUpTest(c *check.C) {
	base := New()
	extension1 := base.Extend()

	extension1.AddSolarSystem(Id(2), Id(20), Id(200), NewEdenId, AnyLocation(), TrueSecurity(0.0))

	suite.builder = extension1.Build().Extend()
}

func (suite *UniverseBuilderTestSuite) TestSolarSystemIdsReturnsCombinedListSorted(c *check.C) {
	suite.builder.AddSolarSystem(Id(1), Id(10), Id(100), NewEdenId, AnyLocation(), TrueSecurity(0.0))
	result := suite.builder.SolarSystemIds()

	c.Assert(result, check.DeepEquals, []Id{1, 2})
}

func (suite *UniverseBuilderTestSuite) TestAddSolarSystemPanicsIfAlreadyExistingInBase(c *check.C) {
	addFunc := func() {
		suite.builder.AddSolarSystem(Id(2), Id(0), Id(0), NewEdenId, AnyLocation(), TrueSecurity(0.0))
	}
	c.Assert(addFunc, check.Panics, "Solar System 2 already exists.")
}

func (suite *UniverseBuilderTestSuite) TestAddSolarSystemPanicsIfAlreadyAdded(c *check.C) {
	addFunc := func() {
		suite.builder.AddSolarSystem(Id(1), Id(0), Id(0), NewEdenId, AnyLocation(), TrueSecurity(0.0))
	}
	suite.builder.AddSolarSystem(Id(1), Id(10), Id(100), NewEdenId, AnyLocation(), TrueSecurity(0.0))

	c.Assert(addFunc, check.Panics, "Solar System 1 already exists.")
}

func (suite *UniverseBuilderTestSuite) TestExtendSolarSystemPanicsIfSystemNotKnown(c *check.C) {
	panicFunc := func() {
		suite.builder.ExtendSolarSystem(Id(1234))
	}

	c.Assert(panicFunc, check.Panics, "Solar System 1234 doesn't exist.")
}

func (suite *UniverseBuilderTestSuite) TestExtendSolarSystemCanExtendBaseSystem(c *check.C) {
	result := suite.builder.ExtendSolarSystem(Id(2))

	c.Assert(result, check.NotNil)
}

func (suite *UniverseBuilderTestSuite) TestExtendSolarSystemCanExtendAddedSystem(c *check.C) {
	suite.builder.AddSolarSystem(Id(1), Id(10), Id(100), NewEdenId, AnyLocation(), TrueSecurity(0.0))
	result := suite.builder.ExtendSolarSystem(Id(1))

	c.Assert(result, check.NotNil)
}

func (suite *UniverseBuilderTestSuite) TestBuildCreatesUniverseThatKnowsBase(c *check.C) {
	result := suite.builder.Build()

	c.Assert(result.HasSolarSystem(Id(2)), check.Equals, true)
}

func (suite *UniverseBuilderTestSuite) TestBuildCreatesUniverseThatKnowsNewSystems(c *check.C) {
	suite.builder.AddSolarSystem(Id(1), Id(10), Id(100), NewEdenId, AnyLocation(), TrueSecurity(0.0))
	result := suite.builder.Build()

	c.Assert(result.HasSolarSystem(Id(1)), check.Equals, true)
}
