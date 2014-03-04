package rules

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"
)

type travelRuleset struct {
	rules []travel.TravelRule
}

func TravelRuleset(rules ...travel.TravelRule) travel.TravelRule {
	var result = &travelRuleset{
		rules: append(make([]travel.TravelRule, 0, len(rules)), rules...)}

	return result
}

func (ruleset *travelRuleset) Compare(sumA *universe.TravelCostSum, sumB *universe.TravelCostSum) float64 {
	var result = 0.0
	var ruleCount = len(ruleset.rules)

	for i := 0; (result == 0.0) && (i < ruleCount); i++ {
		var rule = ruleset.rules[i]

		result = rule.Compare(sumA, sumB)
	}

	return result
}
