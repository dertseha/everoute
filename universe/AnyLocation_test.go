package universe

import (
	"fmt"

	"github.com/dertseha/everoute/util"

	check "gopkg.in/check.v1"
)

type AnyLocationTestSuite struct {
	location Location
}

var _ = check.Suite(&AnyLocationTestSuite{})

func (suite *AnyLocationTestSuite) SetUpTest(c *check.C) {
	suite.location = AnyLocation()
}

func (suite *AnyLocationTestSuite) TestStringReturnsASimpleString(c *check.C) {
	result := fmt.Sprintf("%v", suite.location)

	c.Assert(result, check.Equals, "*")
}

func (suite *AnyLocationTestSuite) TestDistanceToReturnsZero(c *check.C) {
	result := suite.location.DistanceTo(NewSpecificLocation(10, 20, 30))

	c.Assert(result, check.Equals, 0.0)
}

func (suite *AnyLocationTestSuite) TestPositionRelativeToReturnsZeroVector(c *check.C) {
	result := suite.location.PositionRelativeTo(util.Vector3d{10, 20, 30})

	c.Assert(result, check.Equals, util.Vector3d{0.0, 0.0, 0.0})
}
