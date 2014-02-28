package rules

import "github.com/dertseha/everoute/travel"

type travelRuleset struct {
	rules []TravelRule
}

func TravelRuleset(rules []TravelRule) TravelRule {
	var result = &travelRuleset{
		rules: append(make([]TravelRule, 0, len(rules)), rules...)}

	return result
}

func (ruleset *travelRuleset) Compare(sumA travel.TravelCostSum, sumB travel.TravelCostSum) float64 {
	var result = 0.0

	for i := 0; (result == 0) && (i < len(ruleset.rules)); i++ {
		var rule = ruleset.rules[i]

		result = rule.Compare(sumA, sumB)
	}

	return result
}
