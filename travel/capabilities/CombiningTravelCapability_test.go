package capabilities

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"

	check "gopkg.in/check.v1"
)

type CombiningTravelCapabilityTestSuite struct{}

type testingCapability struct {
	path travel.Path
}

func (cap *testingCapability) NextPaths(origin travel.Path) []travel.Path {
	return []travel.Path{cap.path}
}

var _ = check.Suite(&CombiningTravelCapabilityTestSuite{})

func (suite *CombiningTravelCapabilityTestSuite) TestNextPathsReturnsEmptySliceWhenNoCapabilitiesProvided(c *check.C) {
	step := travel.NewStepBuilder(universe.Id(1)).Build()
	origin := travel.NewPath(step)
	capability := CombiningTravelCapability()
	result := capability.NextPaths(origin)

	c.Assert(result, check.HasLen, 0)
}

func (suite *CombiningTravelCapabilityTestSuite) TestNextPathsReturnsCombinationOfAllPaths(c *check.C) {
	step := travel.NewStepBuilder(universe.Id(1)).Build()
	origin := travel.NewPath(step)
	path1 := origin.Extend(travel.NewStepBuilder(universe.Id(2)).Build())
	path2 := origin.Extend(travel.NewStepBuilder(universe.Id(3)).Build())
	cap1 := &testingCapability{path1}
	cap2 := &testingCapability{path2}
	capability := CombiningTravelCapability(cap1, cap2)
	result := capability.NextPaths(origin)

	c.Assert(result, check.DeepEquals, []travel.Path{path1, path2})
}
