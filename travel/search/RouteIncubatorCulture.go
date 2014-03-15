package search

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"
)

type routeIncubatorCulture struct {
	chromosome      *routeChromosome
	destinationPath travel.Path

	steps           []*travel.Step
	routeChromosome *routeChromosome
	costSum         *universe.TravelCostSum
}

func newRouteIncubatorCulture(chromosome *routeChromosome) *routeIncubatorCulture {
	startSteps := chromosome.startPath.Steps()
	culture := &routeIncubatorCulture{
		chromosome:      chromosome,
		destinationPath: nil,

		steps:           make([]*travel.Step, 0, len(startSteps)),
		routeChromosome: newRouteChromosome(chromosome.startPath, len(chromosome.waypoints)),
		costSum:         chromosome.startPath.CostSum()}

	culture.addPath(chromosome.startPath)
	culture.steps = append(culture.steps, startSteps...)

	return culture
}

func (culture *routeIncubatorCulture) addPath(path travel.Path) {
	culture.costSum = culture.costSum.Add(path.CostSum())
	culture.steps = append(culture.steps, path.Steps()[1:]...)
}

func (culture *routeIncubatorCulture) addWaypointPath(index int, path travel.Path) {
	culture.addPath(path)
	culture.routeChromosome.addWaypoint(index, path.DestinationKey())
}

func (culture *routeIncubatorCulture) setDestinationPath(path travel.Path) {
	culture.addPath(path)
	culture.destinationPath = path
}

func (culture *routeIncubatorCulture) toRoute() *Route {
	route := &Route{
		costSum:         culture.costSum,
		steps:           culture.steps,
		routeChromosome: culture.routeChromosome}

	return route
}
