package universe

import "github.com/dertseha/everoute/util"

type specificLocation struct {
	pos util.Vector3d
}

// NewSpecificLocation returns a location instance representing a specific location in three dimensional space.
// The three parameters x, y and z will be used to define the location.
func NewSpecificLocation(x float64, y float64, z float64) Location {
	return &specificLocation{util.Vector3d{x, y, z}}
}

func (this *specificLocation) String() string {
	return this.pos.String()
}

func (this *specificLocation) PositionRelativeTo(origin util.Vector3d) *util.Vector3d {
	return &util.Vector3d{this.pos[0] - origin[0], this.pos[1] - origin[1], this.pos[2] - origin[2]}
}

func (this *specificLocation) DistanceTo(other Location) float64 {
	var pos = other.PositionRelativeTo(this.pos)

	return pos.Length()
}
