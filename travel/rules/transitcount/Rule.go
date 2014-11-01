package transitcount

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/travel/rules"
)

// Rule returns a travel rule that orders costs by their transit count.
func Rule() travel.TravelRule {
	var nullCost = travel.AddingTravelCost(CostType, 0)

	return rules.NaturalOrderTravelRule(nullCost)
}
