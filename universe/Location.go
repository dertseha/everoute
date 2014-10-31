package universe

import "github.com/dertseha/everoute/util"

// Location represents a position in three dimensional space.
type Location interface {
	// PositionRelativeTo returns a vector that describes the difference between this location and the one provided.
	PositionRelativeTo(origin util.Vector3d) *util.Vector3d
	// DistanceTo returns the length between this location and the one provided.
	DistanceTo(other Location) float64
}
