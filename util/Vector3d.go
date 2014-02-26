package util

import (
	"fmt"
	"math"
)

type Vector3d [3]float64

func (this *Vector3d) String() string {
	return fmt.Sprintf("[%f, %f, %f]", this[0], this[1], this[2])
}

func (this *Vector3d) Length() float64 {
	return math.Sqrt(this[0]*this[0] + this[1]*this[1] + this[2]*this[2])
}
