package rules

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"
)

type naturalOrderTravelRule struct {
	nullCost universe.TravelCost
}

// NaturalOrderTravelRule returns a travel rule instance that compares costs using the provided nullCost.
// This rule compares the values of these costs according to their natural order.
func NaturalOrderTravelRule(nullCost universe.TravelCost) travel.TravelRule {
	return &naturalOrderTravelRule{nullCost}
}

func (rule *naturalOrderTravelRule) Compare(sumA *universe.TravelCostSum, sumB *universe.TravelCostSum) float64 {
	return sumA.Cost(rule.nullCost).Value() - sumB.Cost(rule.nullCost).Value()
}
