package warpdistance

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"
)

// Cost returns a travel cost instance that describes warp distances in meters.
func Cost(meters float64) universe.TravelCost {
	return travel.AddingTravelCost(CostType, meters)
}

var nullCost = Cost(0)

// NullCost returns a travel cost instance describing no cost.
func NullCost() universe.TravelCost {
	return nullCost
}
