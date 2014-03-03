package jumpdistance

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/travel/rules"
)

func Rule() travel.TravelRule {
	return rules.NaturalOrderTravelRule(Cost(0.0))
}
