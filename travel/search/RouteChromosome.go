package search

import "github.com/dertseha/everoute/travel"

type waypointChromosome struct {
	index          int
	destinationKey string
}

type routeChromosome struct {
	startPath travel.Path
	waypoints []*waypointChromosome
}

func newRouteChromosome(startPath travel.Path, waypointCount int) *routeChromosome {
	chromosome := &routeChromosome{
		startPath: startPath,
		waypoints: make([]*waypointChromosome, 0, waypointCount)}

	return chromosome
}

func (chromosome *routeChromosome) addWaypoint(index int, destinationKey string) {
	entry := &waypointChromosome{index: index, destinationKey: destinationKey}

	chromosome.waypoints = append(chromosome.waypoints, entry)
}
