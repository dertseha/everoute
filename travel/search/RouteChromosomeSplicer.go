package search

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/util"
)

type routeChromosomeSplicer struct {
	rand util.Randomizer
}

func newChromosomeSplicer(rand util.Randomizer) *routeChromosomeSplicer {
	return &routeChromosomeSplicer{rand: rand}
}

func (splicer *routeChromosomeSplicer) random(startPaths []travel.Path, waypointCount int) *routeChromosome {
	chromosome := newRouteChromosome(startPaths[splicer.rand.Index(len(startPaths))], waypointCount)

	for i := 0; i < waypointCount; i++ {
		chromosome.addWaypoint(splicer.findUnusedIndex(chromosome.waypoints, waypointCount), "")
	}

	return chromosome
}

func (splicer *routeChromosomeSplicer) createOffspring(parent1 *routeChromosome,
	parent2 *routeChromosome, crossoverIndex int) *routeChromosome {
	waypointCount := len(parent1.waypoints)
	child := newRouteChromosome(parent1.startPath, waypointCount)
	index := 0

	for index < crossoverIndex {
		child.waypoints = append(child.waypoints, parent1.waypoints[index])
		index++
	}
	for index < waypointCount {
		temp := parent2.waypoints[index]
		if isWaypointIndexUsed(child.waypoints, temp.index) {
			child.addWaypoint(splicer.findUnusedIndex(child.waypoints, waypointCount), "")
		} else {
			child.waypoints = append(child.waypoints, temp)
		}
		index++
	}

	return child
}

func isWaypointIndexUsed(waypoints []*waypointChromosome, index int) bool {
	result := false

	for _, info := range waypoints {
		if info.index == index {
			result = true
		}
	}

	return result
}

func (splicer *routeChromosomeSplicer) findUnusedIndex(waypoints []*waypointChromosome, limit int) int {
	result := splicer.rand.Index(limit)

	for isWaypointIndexUsed(waypoints, result) {
		result = splicer.rand.Index(limit)
	}

	return result
}
