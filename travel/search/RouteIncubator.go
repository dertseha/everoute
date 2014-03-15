package search

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/util"
)

type routeIncubatorListener interface {
	RouteFound(route *Route)
	IncubatorEmpty()
}

type agentsByDestinationKeyMap map[string]*pathSearchAgent

type routeIncubator struct {
	capability  travel.TravelCapability
	rule        travel.TravelRule
	waypoints   []SearchCriterion
	destination SearchCriterion
	rand        util.Randomizer

	executor util.Executor

	cultures int

	waypointsAgentsBySource   map[int]agentsByDestinationKeyMap
	destinationAgentsBySource agentsByDestinationKeyMap

	listener routeIncubatorListener
}

type routeIncubatorAgentListener struct {
	executor util.Executor

	onSearchFailed    func()
	onSearchCompleted func(result pathSearchAgentResult)
}

func (listener *routeIncubatorAgentListener) searchFailed() {
	listener.executor.Execute(func() { listener.onSearchFailed() })
}

func (listener *routeIncubatorAgentListener) searchCompleted(result pathSearchAgentResult) {
	listener.executor.Execute(func() { listener.onSearchCompleted(result) })
}

func newRouteIncubator(capability travel.TravelCapability, rule travel.TravelRule,
	waypoints []SearchCriterion, destination SearchCriterion, rand util.Randomizer, listener routeIncubatorListener) *routeIncubator {
	incubator := &routeIncubator{
		capability:  capability,
		rule:        rule,
		waypoints:   append(make([]SearchCriterion, 0, len(waypoints)), waypoints...),
		destination: destination,
		rand:        rand,

		executor: util.SingleThreadExecutor(100),

		cultures: 0,

		waypointsAgentsBySource:   make(map[int]agentsByDestinationKeyMap),
		destinationAgentsBySource: make(agentsByDestinationKeyMap),

		listener: listener}

	for i := 0; i < len(waypoints); i++ {
		incubator.waypointsAgentsBySource[i] = make(agentsByDestinationKeyMap)
	}

	return incubator
}

func (incubator *routeIncubator) Request(chromosomes []*routeChromosome) {
	go incubator.executor.Execute(func() {
		incubator.cultures += len(chromosomes)
		for _, chromosome := range chromosomes {
			culture := newRouteIncubatorCulture(chromosome)
			incubator.continueCulture(culture, 0, culture.chromosome.startPath)
		}
	})
}

func (incubator *routeIncubator) continueCulture(culture *routeIncubatorCulture, finishedWaypoints int, lastPath travel.Path) {
	startPath := travel.NewPath(lastPath.Step().AsFirstStep())
	chromosome := culture.chromosome

	if finishedWaypoints < len(chromosome.waypoints) {
		waypointIndex := chromosome.waypoints[finishedWaypoints].index
		destinationKey := chromosome.waypoints[waypointIndex].destinationKey
		agent := incubator.pathSearchAgent(incubator.waypointsAgentsBySource[waypointIndex], startPath, incubator.waypoints[waypointIndex])
		onSearchCompleted := func(result pathSearchAgentResult) {
			path := result.Path(destinationKey, incubator.rand)

			culture.addWaypointPath(waypointIndex, path)
			incubator.continueCulture(culture, finishedWaypoints+1, path)
		}

		incubator.searchForCulture(agent, culture, onSearchCompleted)
	} else if incubator.destination != nil {
		agent := incubator.pathSearchAgent(incubator.destinationAgentsBySource, startPath, incubator.destination)
		onSearchCompleted := func(result pathSearchAgentResult) {
			path := result.BestPath()

			culture.setDestinationPath(path)
			incubator.onCultureCompleted(culture)
		}

		incubator.searchForCulture(agent, culture, onSearchCompleted)
	} else {
		incubator.onCultureCompleted(culture)
	}
}

func (incubator *routeIncubator) searchForCulture(agent *pathSearchAgent, culture *routeIncubatorCulture,
	onSearchCompleted func(result pathSearchAgentResult)) {
	listener := &routeIncubatorAgentListener{
		executor:          incubator.executor,
		onSearchFailed:    func() { incubator.onCultureFailed(culture) },
		onSearchCompleted: onSearchCompleted}

	agent.Request(listener)
}

func (incubator *routeIncubator) pathSearchAgent(agentsByKey agentsByDestinationKeyMap,
	startPath travel.Path, criterion SearchCriterion) *pathSearchAgent {
	startKey := startPath.DestinationKey()
	agent, existing := agentsByKey[startKey]

	if !existing {
		agent = newPathSearchAgent(startPath, incubator.capability, incubator.rule, criterion)
		agentsByKey[startKey] = agent
	}

	return agent
}

func (incubator *routeIncubator) onCultureFailed(culture *routeIncubatorCulture) {
	incubator.onCultureDone()
}

func (incubator *routeIncubator) onCultureCompleted(culture *routeIncubatorCulture) {
	incubator.listener.RouteFound(culture.toRoute())
	incubator.onCultureDone()
}

func (incubator *routeIncubator) onCultureDone() {
	incubator.cultures--
	if incubator.cultures == 0 {
		incubator.listener.IncubatorEmpty()
	}
}
