package rules

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"
)

type naturalOrderTravelRule struct {
	nullCost universe.TravelCost
}

func NaturalOrderTravelRule(nullCost universe.TravelCost) TravelRule {
	return &naturalOrderTravelRule{nullCost}
}

func (rule *naturalOrderTravelRule) Compare(sumA travel.TravelCostSum, sumB travel.TravelCostSum) float64 {
	return sumA.Cost(rule.nullCost).Value() - sumB.Cost(rule.nullCost).Value()
}
