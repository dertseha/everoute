package jumpdistance

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/travel/rules"
)

// Rule returns a travel rule that orders jump distances by their value.
func Rule() travel.TravelRule {
	return rules.NaturalOrderTravelRule(NullCost())
}
