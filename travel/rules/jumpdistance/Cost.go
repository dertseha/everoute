package jumpdistance

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"
)

func Cost(lightYears float64) universe.TravelCost {
	return travel.AddingTravelCost(CostType, lightYears)
}
