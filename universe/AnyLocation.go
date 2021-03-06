package universe

import "github.com/dertseha/everoute/util"

type anyLocation struct{}

func (this *anyLocation) String() string {
	return "*"
}

func (this *anyLocation) PositionRelativeTo(origin util.Vector3d) *util.Vector3d {
	return &util.Vector3d{0.0, 0.0, 0.0}
}

func (this *anyLocation) DistanceTo(other Location) float64 {
	return 0.0
}

var instance Location = new(anyLocation)

// AnyLocation returns a location instance that represents a point that can be anywhere.
// When compared to other locations, this location will always appear as if it were right at the other location.
func AnyLocation() Location {
	return instance
}
