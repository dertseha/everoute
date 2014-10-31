package universe

import (
	"math"

	"github.com/dertseha/everoute/util"

	check "gopkg.in/check.v1"
)

type SpecificLocationTestSuite struct{}

var _ = check.Suite(&SpecificLocationTestSuite{})

func (suite *SpecificLocationTestSuite) TestPositionRelativeToShouldReturnItselfForZero(c *check.C) {
	location := NewSpecificLocation(10, 20, 30)
	result := location.PositionRelativeTo(util.Vector3d{0, 0, 0})
	expected := &util.Vector3d{10, 20, 30}

	c.Assert(result, check.Equals, expected)
}

func (suite *SpecificLocationTestSuite) TestPositionRelativeToShouldReturnRelativeVectorToAnotherPosition(c *check.C) {
	location := SpecificLocation{util.Vector3d{10, -20, 30}}
	result := location.PositionRelativeTo(util.Vector3d{-5, 5, 5})
	expected := &util.Vector3d{15, -25, 25}

	assert.Equal(result, check.Equals, expected)
}

func (suite *SpecificLocationTestSuite) TestDistanceToShouldReturnZeroForItself(c *check.C) {
	location := SpecificLocation{util.Vector3d{10, -20, 30}}
	result := location.DistanceTo(&location)

	c.Assert(result, check.Equals, 0.0)
}

func (suite *SpecificLocationTestSuite) TestDistanceToShouldReturnZeroForAnyLocation(c *check.C) {
	location := SpecificLocation{util.Vector3d{10, -20, 30}}
	result := location.DistanceTo(AnyLocation())

	c.Assert(result, check.Equals, 0.0)
}

func (suite *SpecificLocationTestSuite) TestDistanceToShouldCalculateDistance1(c *check.C) {
	location1 := SpecificLocation{util.Vector3d{10, 20, 30}}
	location2 := SpecificLocation{util.Vector3d{40, 80, -30}}
	result := location1.DistanceTo(&location2)

	c.Assert(result, check.Equals, 90.0)
}

func (suite *SpecificLocationTestSuite) TestDistanceToShouldCalculateDistance2(c *check.C) {
	location1 := SpecificLocation{util.Vector3d{15, 49, 140}}
	location2 := SpecificLocation{util.Vector3d{120.5, -33, -9}}
	result := location1.DistanceTo(&location2)
	expected := 200.138

	c.Assert(math.Trunc(result*1000), check.Equals, expected*1000)
}
