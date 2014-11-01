package universe

import (
	"math/rand"
	"time"

	check "gopkg.in/check.v1"
)

type SolarSystemTestSuite struct {
	id              Id
	constellationId Id
	regionId        Id
	galaxyId        GalaxyId
	location        Location
	security        TrueSecurity

	createSystem func(id Id, constellationId Id, regionId Id, galaxyId GalaxyId, location Location, security TrueSecurity) SolarSystem

	system SolarSystem
}

func (suite *SolarSystemTestSuite) SetUpTest(c *check.C) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	suite.id = Id(random.Int63())
	suite.constellationId = Id(random.Int63())
	suite.regionId = Id(random.Int63())
	suite.galaxyId = GalaxyId(random.Int63())
	suite.location = NewSpecificLocation(random.Float64(), random.Float64(), random.Float64())
	suite.security = TrueSecurity((random.Intn(201) - 100) / 100.0)

	suite.system = suite.createSystem(suite.id, suite.constellationId, suite.regionId, suite.galaxyId, suite.location, suite.security)
}

func (suite *SolarSystemTestSuite) SolarSystem() SolarSystem {
	return suite.system
}

func (suite *SolarSystemTestSuite) TestIdReturnsSolarSystemId(c *check.C) {
	c.Assert(suite.system.Id(), check.Equals, suite.id)
}

func (suite *SolarSystemTestSuite) TestConstellationIdReturnsConstellationId(c *check.C) {
	c.Assert(suite.system.ConstellationId(), check.Equals, suite.constellationId)
}

func (suite *SolarSystemTestSuite) TestRegionIdReturnsRegionId(c *check.C) {
	c.Assert(suite.system.RegionId(), check.Equals, suite.regionId)
}

func (suite *SolarSystemTestSuite) TestGalaxyIdReturnsGalaxyId(c *check.C) {
	c.Assert(suite.system.GalaxyId(), check.Equals, suite.galaxyId)
}

func (suite *SolarSystemTestSuite) TestLocationReturnsLocation(c *check.C) {
	c.Assert(suite.system.Location(), check.Equals, suite.location)
}

func (suite *SolarSystemTestSuite) TestSecurityReturnsSecurity(c *check.C) {
	c.Assert(suite.system.Security(), check.Equals, suite.security)
}

func (suite *SolarSystemTestSuite) TestCostsReturnsAnInstance(c *check.C) {
	c.Assert(suite.system.Costs(), check.NotNil)
}

func (suite *SolarSystemTestSuite) TestJumpsReturnsEmptySliceForUnknownJumpType(c *check.C) {
	c.Assert(suite.system.Jumps("test_unkonwn"), check.HasLen, 0)
}
