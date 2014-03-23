package search

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"
)

type systemAvoidingSearchCriterion struct {
	ignoredIds map[universe.Id]bool
}

func SystemAvoidingSearchCriterion(solarSystemIds ...universe.Id) SearchCriterion {
	ignoredIds := make(map[universe.Id]bool)

	for _, id := range solarSystemIds {
		ignoredIds[id] = true
	}

	return &systemAvoidingSearchCriterion{ignoredIds: ignoredIds}
}

func (criterion *systemAvoidingSearchCriterion) IsDesired(path travel.Path) bool {
	return true
}

func (criterion *systemAvoidingSearchCriterion) ShouldSearchContinueWith(path travel.Path, results []travel.Path) bool {
	return path.IsStart() || !criterion.ignoredIds[path.Step().SolarSystemId()]
}
