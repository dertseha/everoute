package security

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"
)

type maxRule struct {
	limit int
}

func (rule *maxRule) Compare(sumA *universe.TravelCostSum, sumB *universe.TravelCostSum) float64 {
	var valueA = sumSecurityCosts(sumA, rule.limit, 10)
	var valueB = sumSecurityCosts(sumB, rule.limit, 10)

	return valueA - valueB
}

// MaxRule returns a travel rule instance that compares the amount of systems that have a security value equal
// or above the given limit. limit is expressed in a floating point value.
func MaxRule(limit float64) travel.TravelRule {
	var rule = &maxRule{limit: int(limit * 10)}

	return rule
}
