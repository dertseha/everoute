package security

import (
	"github.com/dertseha/everoute/universe"

	"github.com/stvp/assert"

	"testing"
)

func TestTravelCostTypeForSecurity00(t *testing.T) {
	var security = universe.TrueSecurity(0.0)
	var result = travelCostType(security)

	assert.Equal(t, "security00", result)
}

func TestTravelCostTypeForSecurity0346(t *testing.T) {
	var security = universe.TrueSecurity(0.346)
	var result = travelCostType(security)

	assert.Equal(t, "security03", result)
}

func TestTravelCostTypeForSecurity099(t *testing.T) {
	var security = universe.TrueSecurity(0.99)
	var result = travelCostType(security)

	assert.Equal(t, "security10", result)
}

func TestTravelCostTypeForSecurity10(t *testing.T) {
	var security = universe.TrueSecurity(1.0)
	var result = travelCostType(security)

	assert.Equal(t, "security10", result)
}
