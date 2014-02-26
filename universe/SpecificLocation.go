package universe

import "github.com/dertseha/everoute/util"

type SpecificLocation struct {
	pos util.Vector3d
}

func NewSpecificLocation(x float64, y float64, z float64) *SpecificLocation {
	return &SpecificLocation{util.Vector3d{x, y, z}}
}

func (this *SpecificLocation) String() string {
	return this.pos.String()
}

func (this *SpecificLocation) PositionRelativeTo(origin util.Vector3d) *util.Vector3d {
	return &util.Vector3d{this.pos[0] - origin[0], this.pos[1] - origin[1], this.pos[2] - origin[2]}
}

func (this *SpecificLocation) DistanceTo(other Location) float64 {
	var pos = other.PositionRelativeTo(this.pos)

	return pos.Length()
}
