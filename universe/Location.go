package universe

import "github.com/dertseha/everoute/util"

type Location interface {
	PositionRelativeTo(origin util.Vector3d) *util.Vector3d
	DistanceTo(other Location) float64
}
