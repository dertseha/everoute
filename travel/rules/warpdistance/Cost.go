package warpdistance

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"
)

func Cost(meters float64) universe.TravelCost {
	return travel.AddingTravelCost(CostType, meters)
}

var nullCost = Cost(0)

func NullCost() universe.TravelCost {
	return nullCost
}
