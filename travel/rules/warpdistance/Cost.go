package warpdistance

import (
	"github.com/dertseha/everoute/universe"
)

// Cost returns a travel cost instance that describes warp distances in meters.
func Cost(meters float64) universe.TravelCost {
	return universe.AddingTravelCost(CostType, meters)
}

var nullCost = Cost(0)

// NullCost returns a travel cost instance describing no cost.
func NullCost() universe.TravelCost {
	return nullCost
}
