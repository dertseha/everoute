package transitcount

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/travel/rules"
)

func Rule() travel.TravelRule {
	var nullCost = travel.AddingTravelCost(CostType, 0)

	return rules.NaturalOrderTravelRule(nullCost)
}
