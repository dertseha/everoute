package universe

import "math"

type TrueSecurity float64

func (security TrueSecurity) Rounded() float64 {
	value := 0.0

	if security > 0.0 {
		value = math.Floor((float64(security)+0.05)*10.0) / 10.0
	}

	return value
}

func (security TrueSecurity) IsHighSec() bool {
	return security.Rounded() >= 0.5
}
