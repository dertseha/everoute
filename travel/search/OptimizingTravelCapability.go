package search

import (
	"github.com/dertseha/everoute/travel"
)

type optimizingTravelCapability struct {
	capability travel.TravelCapability
	contest    travel.PathContest
}

func newOptimizingTravelCapability(capability travel.TravelCapability, contest travel.PathContest) travel.TravelCapability {
	var result = &optimizingTravelCapability{
		capability: capability,
		contest:    contest}

	return result
}

func (optimizing *optimizingTravelCapability) NextPaths(origin *travel.Path) []*travel.Path {
	var best = make(map[string]*travel.Path)

	if optimizing.contest.Enter(origin) {
		var temp = optimizing.capability.NextPaths(origin)

		for _, foundPath := range temp {
			var destinationKey = foundPath.DestinationKey()

			if (origin.IsStart() || (origin.Previous().DestinationKey() != destinationKey)) && optimizing.contest.Enter(foundPath) {
				best[destinationKey] = foundPath
			}
		}
	}

	result := make([]*travel.Path, 0, len(best))
	for _, foundPath := range best {
		result = append(result, foundPath)
	}

	return result
}
