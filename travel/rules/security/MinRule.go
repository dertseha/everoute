package security

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"
)

type minRule struct {
	limit int
}

func (rule *minRule) Compare(sumA *universe.TravelCostSum, sumB *universe.TravelCostSum) float64 {
	var valueA = sumSecurityCosts(sumA, 0, rule.limit)
	var valueB = sumSecurityCosts(sumB, 0, rule.limit)

	return valueA - valueB
}

// MinRule returns a travel rule instance that compares the amount of systems that have a security value below
// the given limit. limit is expressed in a floating point value.
func MinRule(limit float64) travel.TravelRule {
	var rule = &minRule{limit: int((limit * 10) - 1)}

	return rule
}
