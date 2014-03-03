package security

import (
	"fmt"

	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"
)

func travelCostType(security universe.TrueSecurity) string {
	return fmt.Sprintf("security%02d", int(security.Rounded()*10))
}

func travelCost(security universe.TrueSecurity, value float64) universe.TravelCost {
	return travel.AddingTravelCost(travelCostType(security), value)
}

func sumSecurityCosts(costSum *travel.TravelCostSum, from int, to int) float64 {
	var result = 0.0

	for i := from; i <= to; i++ {
		var security = universe.TrueSecurity(float64(i) / 10.0)
		var nullCost = travelCost(security, 0)

		result += costSum.Cost(nullCost).Value()
	}

	return result
}
