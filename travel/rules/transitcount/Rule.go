package transitcount

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/travel/rules"
	"github.com/dertseha/everoute/universe"
)

// Rule returns a travel rule that orders costs by their transit count.
func Rule() travel.TravelRule {
	var nullCost = universe.AddingTravelCost(CostType, 0)

	return rules.NaturalOrderTravelRule(nullCost)
}
