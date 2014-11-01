package search

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"
)

type systemAvoidingSearchCriterion struct {
	ignoredIds map[universe.Id]bool
}

// SystemAvoidingSearchCriterion returns a SearchCriterion instance that aborts searches if a path ends up in one
// of the avoided solar systems. Note that this criterion does not limit desired paths - should the destination be
// one of the avoided systems, this is not prohibited.
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
