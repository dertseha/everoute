package universe

import (
	"github.com/dertseha/everoute/util"
	"github.com/stvp/assert"
	"math"
	"testing"
)

func TestPositionRelativeToShouldReturnItselfForZero(t *testing.T) {
	var location = SpecificLocation{util.Vector3d{10, 20, 30}}
	var result = location.PositionRelativeTo(util.Vector3d{0, 0, 0})
	var expected = &util.Vector3d{10, 20, 30}

	assert.Equal(t, expected, result)
}

func TestPositionRelativeToShouldReturnRelativeVectorToAnotherPosition(t *testing.T) {
	var location = SpecificLocation{util.Vector3d{10, -20, 30}}
	var result = location.PositionRelativeTo(util.Vector3d{-5, 5, 5})
	var expected = &util.Vector3d{15, -25, 25}

	assert.Equal(t, expected, result)
}

func TestDistanceToShouldReturnZeroForItself(t *testing.T) {
	var location = SpecificLocation{util.Vector3d{10, -20, 30}}
	var result = location.DistanceTo(&location)

	assert.Equal(t, 0.0, result)
}

func TestDistanceToShouldReturnZeroForAnyLocation(t *testing.T) {
	var location = SpecificLocation{util.Vector3d{10, -20, 30}}
	var result = location.DistanceTo(AnyLocation())

	assert.Equal(t, 0.0, result)
}

func TestDistanceToShouldCalculateDistance1(t *testing.T) {
	var location1 = SpecificLocation{util.Vector3d{10, 20, 30}}
	var location2 = SpecificLocation{util.Vector3d{40, 80, -30}}
	var result = location1.DistanceTo(&location2)

	assert.Equal(t, 90.0, result)
}

func TestDistanceToShouldCalculateDistance2(t *testing.T) {
	var location1 = SpecificLocation{util.Vector3d{15, 49, 140}}
	var location2 = SpecificLocation{util.Vector3d{120.5, -33, -9}}
	var result = location1.DistanceTo(&location2)
	var expected = 200.138

	assert.Equal(t, expected*1000, math.Trunc(result*1000))
}
