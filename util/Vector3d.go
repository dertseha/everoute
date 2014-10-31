package util

import (
	"fmt"
	"math"
)

// A Vector3d is a group of three floating point values specifying the offsets along the three axes X, Y and Z.
type Vector3d [3]float64

// String is the Stringer interface implementation.
func (vec *Vector3d) String() string {
	return fmt.Sprintf("[%f, %f, %f]", vec[0], vec[1], vec[2])
}

// Length returns the length of the vector.
func (vec *Vector3d) Length() float64 {
	return math.Sqrt(vec[0]*vec[0] + vec[1]*vec[1] + vec[2]*vec[2])
}
