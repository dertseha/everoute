package jumpdistance

import (
	"github.com/dertseha/everoute/universe"
)

// Cost returns a travel cost instance that describes jump distances in light years.
func Cost(lightYears float64) universe.TravelCost {
	return universe.AddingTravelCost(CostType, lightYears)
}

var nullCost = Cost(0)

// NullCost returns a travel cost instance describing no cost.
func NullCost() universe.TravelCost {
	return nullCost
}
