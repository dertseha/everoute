package search

import "github.com/dertseha/everoute/travel"

type costAwareSearchCriterion struct {
	rule travel.TravelRule
}

func CostAwareSearchCriterion(rule travel.TravelRule) SearchCriterion {
	return &costAwareSearchCriterion{rule: rule}
}

func (criterion *costAwareSearchCriterion) IsDesired(path travel.Path) bool {
	return true
}

func (criterion *costAwareSearchCriterion) ShouldSearchContinueWith(path travel.Path, results []travel.Path) bool {
	var resultCount = len(results)
	var isCheaper = true

	if resultCount > 0 {
		var costSum = path.CostSum()

		for i := 0; isCheaper && (i < resultCount); i++ {
			previousResult := results[i]
			isCheaper = criterion.rule.Compare(costSum, previousResult.CostSum()) < 0
		}
	}

	return isCheaper
}
