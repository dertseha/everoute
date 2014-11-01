package universe

import (
	check "gopkg.in/check.v1"
)

type EmptyUniverseTestSuite struct {
	UniverseTestSuite
}

var _ = check.Suite(&EmptyUniverseTestSuite{UniverseTestSuite{createUniverse: New}})

func (suite *EmptyUniverseTestSuite) TestSolarSystemIdsReturnsEmptySlice(c *check.C) {
	c.Assert(suite.Universe().SolarSystemIds(), check.HasLen, 0)
}
