package search

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"
)

type Route struct {
	costSum         *universe.TravelCostSum
	steps           []*travel.Step
	routeChromosome *routeChromosome
}

func (route *Route) CostSum() *universe.TravelCostSum {
	return route.costSum
}

func (route *Route) Steps() []*travel.Step {
	return route.steps
}

func (route *Route) chromosome() *routeChromosome {
	return route.routeChromosome
}
