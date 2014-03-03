package search

import (
	"github.com/dertseha/everoute/travel"
)

type combiningSearchCriterion struct {
	criteria []SearchCriterion
}

func CombiningSearchCriterion(criteria ...SearchCriterion) SearchCriterion {
	return &combiningSearchCriterion{criteria: append(make([]SearchCriterion, 0, len(criteria)), criteria...)}
}

func (criterion *combiningSearchCriterion) reduceCriteria(f func(bool, SearchCriterion) bool, init bool) bool {
	var result = init

	for _, nested := range criterion.criteria {
		result = f(result, nested)
	}

	return result
}

func (criterion *combiningSearchCriterion) IsDesired(path travel.Path) bool {
	var result = false

	if len(criterion.criteria) > 0 {
		result = criterion.reduceCriteria(func(desired bool, nestedCriterion SearchCriterion) bool {
			return desired && nestedCriterion.IsDesired(path)
		}, true)
	}

	return result
}

func (criterion *combiningSearchCriterion) ShouldSearchContinueWith(path travel.Path, results []travel.Path) bool {
	var result = false

	if len(criterion.criteria) > 0 {
		result = criterion.reduceCriteria(func(desired bool, nestedCriterion SearchCriterion) bool {
			return desired && nestedCriterion.ShouldSearchContinueWith(path, results)
		}, true)
	}

	return result
}
