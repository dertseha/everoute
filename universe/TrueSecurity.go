package universe

import "math"

// TrueSecurity represents the security level of a solar system. Its own value represents the actual security value.
type TrueSecurity float64

// Rounded returns the security value rounded (and clipped) to a value range that is known to the end-user.
// The possible returned values are {0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0}
func (security TrueSecurity) Rounded() float64 {
	value := 0.0

	if security > 0.0 {
		value = math.Floor((float64(security)+0.05)*10.0) / 10.0
	}

	return value
}

// IsHighSec returns true if the security value represents one that classifies the system as 'High Security'. This is
// for systems with a rounded security value of 0.5 or higher.
func (security TrueSecurity) IsHighSec() bool {
	return security.Rounded() >= 0.5
}
