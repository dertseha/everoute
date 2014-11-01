package universe

import (
	check "gopkg.in/check.v1"
)

type JumpBuilderTestSuite struct {
	builder *JumpBuilder
}

var _ = check.Suite(&JumpBuilderTestSuite{})

func (suite *JumpBuilderTestSuite) SetUpTest(c *check.C) {
	suite.builder = newJumpBuilder("testType", Id(1))
}

func (suite *JumpBuilderTestSuite) TestBuildSetsJumpType(c *check.C) {
	result := suite.builder.Build()

	c.Assert(result.Type(), check.Equals, "testType")
}

func (suite *JumpBuilderTestSuite) TestBuildSetsDestinationId(c *check.C) {
	result := suite.builder.Build()

	c.Assert(result.DestinationId(), check.Equals, Id(1))
}

func (suite *JumpBuilderTestSuite) TestDefaultBuildUsesAnyLocationForSourceLocation(c *check.C) {
	result := suite.builder.Build()

	c.Assert(result.SourceLocation(), check.DeepEquals, AnyLocation())
}

func (suite *JumpBuilderTestSuite) TestDefaultBuildUsesAnyLocationForDestinationLocation(c *check.C) {
	result := suite.builder.Build()

	c.Assert(result.DestinationLocation(), check.DeepEquals, AnyLocation())
}

func (suite *JumpBuilderTestSuite) TestDefaultBuildUsesEmptyTravelCostSumForCosts(c *check.C) {
	result := suite.builder.Build()

	c.Assert(result.Costs(), check.DeepEquals, EmptyTravelCostSum())
}

func (suite *JumpBuilderTestSuite) TestFromSetsSourceLocation(c *check.C) {
	location := NewSpecificLocation(10, 20, 30)
	result := suite.builder.From(location).Build()

	c.Assert(result.SourceLocation(), check.DeepEquals, location)
}

func (suite *JumpBuilderTestSuite) TestToSetsDestinationLocation(c *check.C) {
	location := NewSpecificLocation(10, 20, 30)
	result := suite.builder.To(location).Build()

	c.Assert(result.DestinationLocation(), check.DeepEquals, location)
}

func (suite *JumpBuilderTestSuite) TestAddCostCombinesToCosts(c *check.C) {
	result := suite.builder.AddCost(AddingTravelCost("type1", 10.0)).AddCost(AddingTravelCost("type1", 20.0)).Build()

	c.Assert(result.Costs().Cost(AddingTravelCost("type1", 0.0)).Value(), check.Equals, 30.0)
}
