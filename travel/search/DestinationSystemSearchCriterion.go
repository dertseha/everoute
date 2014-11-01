package search

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"
)

type destinationSystemSearchCriterion struct {
	solarSystemId universe.Id
}

// DestinationSystemSearchCriterion returns a SearchCriterion instance that desires all paths that end in the specified
// solar system and aborts searches if a path ends in this system.
func DestinationSystemSearchCriterion(solarSystemId universe.Id) SearchCriterion {
	return &destinationSystemSearchCriterion{solarSystemId: solarSystemId}
}

func (criterion *destinationSystemSearchCriterion) IsDesired(path travel.Path) bool {
	return path.Step().SolarSystemId() == criterion.solarSystemId
}

func (criterion *destinationSystemSearchCriterion) ShouldSearchContinueWith(path travel.Path, results []travel.Path) bool {
	return path.Step().SolarSystemId() != criterion.solarSystemId
}
