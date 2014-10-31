package universe

import (
	"math/rand"
	"time"

	check "gopkg.in/check.v1"
)

type EmptySolarSystemTestSuite struct {
	id              Id
	constellationId Id
	regionId        Id
	galaxyId        GalaxyId
	location        Location
	security        TrueSecurity

	system SolarSystem
}

var _ = check.Suite(&EmptySolarSystemTestSuite{})

func (suite *EmptySolarSystemTestSuite) SetUpTest(c *check.C) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	suite.id = Id(random.Int63())
	suite.constellationId = Id(random.Int63())
	suite.regionId = Id(random.Int63())
	suite.galaxyId = GalaxyId(random.Int63())
	suite.location = NewSpecificLocation(random.Float64(), random.Float64(), random.Float64())
	suite.security = TrueSecurity((random.Intn(201) - 100) / 100.0)

	suite.system = newSolarSystem(suite.id, suite.constellationId, suite.regionId, suite.galaxyId, suite.location, suite.security)
}

func (suite *EmptySolarSystemTestSuite) TestIdReturnsSolarSystemId(c *check.C) {
	c.Assert(suite.system.Id(), check.Equals, suite.id)
}

func (suite *EmptySolarSystemTestSuite) TestConstellationIdReturnsConstellationId(c *check.C) {
	c.Assert(suite.system.ConstellationId(), check.Equals, suite.constellationId)
}

func (suite *EmptySolarSystemTestSuite) TestRegionIdReturnsRegionId(c *check.C) {
	c.Assert(suite.system.RegionId(), check.Equals, suite.regionId)
}

func (suite *EmptySolarSystemTestSuite) TestGalaxyIdReturnsGalaxyId(c *check.C) {
	c.Assert(suite.system.GalaxyId(), check.Equals, suite.galaxyId)
}

func (suite *EmptySolarSystemTestSuite) TestLocationReturnsLocation(c *check.C) {
	c.Assert(suite.system.Location(), check.Equals, suite.location)
}

func (suite *EmptySolarSystemTestSuite) TestSecurityReturnsSecurity(c *check.C) {
	c.Assert(suite.system.Security(), check.Equals, suite.security)
}
