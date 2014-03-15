package search

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/util"
)

type RouteFinderBuilder interface {
	AddWaypoint(criterion SearchCriterion) RouteFinderBuilder
	ForDestination(criterion SearchCriterion) RouteFinderBuilder

	Build() RouteFinder
}

type routeFinderBuilder struct {
	capability  travel.TravelCapability
	rule        travel.TravelRule
	startPaths  []travel.Path
	waypoints   []SearchCriterion
	destination SearchCriterion
	collector   RouteSearchResultCollector

	searchDone func()

	populationLimit  int
	generationLimit  int
	uncontestedLimit int

	mutationPercentage int

	rand util.Randomizer
}

func NewRouteFinder(capability travel.TravelCapability, rule travel.TravelRule,
	startPaths []travel.Path, collector RouteSearchResultCollector, searchDone func()) RouteFinderBuilder {
	builder := &routeFinderBuilder{
		capability: capability,
		rule:       rule,
		startPaths: startPaths,
		waypoints:  make([]SearchCriterion, 0),
		collector:  collector,

		searchDone: searchDone,

		populationLimit:    50,
		generationLimit:    40000,
		mutationPercentage: 20,

		rand: util.DefaultRandomizer()}

	return builder
}

func (builder *routeFinderBuilder) AddWaypoint(criterion SearchCriterion) RouteFinderBuilder {
	builder.waypoints = append(builder.waypoints, criterion)

	return builder
}

func (builder *routeFinderBuilder) ForDestination(criterion SearchCriterion) RouteFinderBuilder {
	builder.destination = criterion

	return builder
}

func (builder *routeFinderBuilder) Build() RouteFinder {
	finder := &routeFinder{
		startPaths:    builder.startPaths,
		waypointCount: len(builder.waypoints),
		collector:     builder.collector,

		searchDone: builder.searchDone,

		populationLimit:    builder.populationLimit,
		generationLimit:    builder.generationLimit,
		uncontestedLimit:   builder.populationLimit * 20,
		mutationPercentage: builder.mutationPercentage,

		executor: util.SingleThreadExecutor(builder.populationLimit * 4),
		rand:     builder.rand,

		splicer:    newChromosomeSplicer(builder.rand),
		population: emptyRouteList(builder.rule)}

	finder.incubator = newRouteIncubator(builder.capability, builder.rule, builder.waypoints, builder.destination, builder.rand, finder)

	finder.IncubatorEmpty()

	return finder
}
