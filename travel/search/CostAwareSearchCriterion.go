package search

import "github.com/dertseha/everoute/travel"

type costAwareSearchCriterion struct {
	rule travel.TravelRule
}

// CostAwareSearchCriterion returns a SearchCriterion instance that desires all paths but aborts searches if a path
// is considered worse than any of the already found paths. It uses a travel rule for comparing the costs.
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
