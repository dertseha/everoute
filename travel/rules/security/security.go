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

var sumSecurityCosts func(*travel.TravelCostSum, int, int) float64 = (func() func(*travel.TravelCostSum, int, int) float64 {
	var nullCosts = make(map[int]universe.TravelCost)

	for i := 0; i <= 10; i++ {
		nullCosts[i] = travelCost(universe.TrueSecurity(float64(i)/10.0), 0)
	}

	return func(costSum *travel.TravelCostSum, from int, to int) float64 {
		var result = 0.0

		for i := from; i <= to; i++ {
			result += costSum.Cost(nullCosts[i]).Value()
		}

		return result
	}
})()
