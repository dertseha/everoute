package universe

import (
	"math/rand"
	"time"

	check "gopkg.in/check.v1"
)

type SolarSystemExtensionTestSuite struct {
	id              Id
	constellationId Id
	regionId        Id
	galaxyId        GalaxyId
	location        Location
	security        TrueSecurity

	system    SolarSystem
	extension *SolarSystemExtension
}

var _ = check.Suite(&SolarSystemExtensionTestSuite{})

func (suite *SolarSystemExtensionTestSuite) SetUpTest(c *check.C) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	suite.id = Id(random.Int63())
	suite.constellationId = Id(random.Int63())
	suite.regionId = Id(random.Int63())
	suite.galaxyId = GalaxyId(random.Int63())
	suite.location = NewSpecificLocation(random.Float64(), random.Float64(), random.Float64())
	suite.security = TrueSecurity((random.Intn(201) - 100) / 100.0)

	suite.system = newSolarSystem(suite.id, suite.constellationId, suite.regionId, suite.galaxyId, suite.location, suite.security)
	suite.extension = &SolarSystemExtension{newSolarSystemExtensionData(suite.system)}
}
func (suite *SolarSystemExtensionTestSuite) TestIdReturnsSolarSystemId(c *check.C) {
	c.Assert(suite.extension.Id(), check.Equals, suite.id)
}

func (suite *SolarSystemExtensionTestSuite) TestGalaxyIdReturnsGalaxyId(c *check.C) {
	c.Assert(suite.extension.GalaxyId(), check.Equals, suite.galaxyId)
}

func (suite *SolarSystemExtensionTestSuite) TestLocationReturnsLocation(c *check.C) {
	c.Assert(suite.extension.Location(), check.Equals, suite.location)
}

func (suite *SolarSystemExtensionTestSuite) TestSecurityReturnsSecurity(c *check.C) {
	c.Assert(suite.extension.Security(), check.Equals, suite.security)
}
