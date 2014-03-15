package search

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/util"
)

type RouteFinder interface {
	Stop()
}

type routeFinder struct {
	startPaths    []travel.Path
	waypointCount int
	collector     RouteSearchResultCollector

	searchDone func()

	executor util.Executor
	rand     util.Randomizer

	splicer    *routeChromosomeSplicer
	incubator  *routeIncubator
	population *routeList

	populationLimit    int
	generationLimit    int
	uncontestedLimit   int
	mutationPercentage int

	generationCount  int
	uncontestedCount int
}

func (finder *routeFinder) shallContinue() bool {
	return (finder.generationCount < finder.generationLimit) && (finder.uncontestedCount < finder.uncontestedLimit)
}

func (finder *routeFinder) notifyDone() {
	go finder.searchDone()
	finder.searchDone = func() {}
}

func (finder *routeFinder) Stop() {
	go finder.executor.Execute(func() {
		finder.generationCount = finder.generationLimit
		finder.notifyDone()
	})
}

func (finder *routeFinder) RouteFound(route *Route) {
	finder.executor.Execute(func() {
		finder.population = finder.population.Add(route)
		if finder.population.Route(0) == route {
			finder.uncontestedCount = 0
			finder.collector.Collect(route)
		}
	})
}

func (finder *routeFinder) IncubatorEmpty() {
	finder.executor.Execute(func() {
		finder.generationCount++
		finder.uncontestedCount++

		if finder.shallContinue() {
			finder.limitPopulation()

			newChromosomes := make([]*routeChromosome, 0, 5)
			newChromosomes = finder.ensurePopulationSize(newChromosomes)
			newChromosomes = finder.createOffsprings(newChromosomes)
			finder.incubator.Request(newChromosomes)
		} else {
			finder.notifyDone()
		}
	})
}

func (finder *routeFinder) limitPopulation() {
	finder.population = finder.population.Limit(finder.populationLimit)
}

func (finder *routeFinder) ensurePopulationSize(chromosomes []*routeChromosome) []*routeChromosome {
	result := chromosomes

	for missing := finder.populationLimit - finder.population.Size(); missing > 0; missing-- {
		chromosome := finder.splicer.random(finder.startPaths, finder.waypointCount)
		result = append(result, chromosome)
	}

	return result
}

func (finder *routeFinder) createOffsprings(chromosomes []*routeChromosome) []*routeChromosome {
	result := chromosomes
	size := finder.population.Size()

	if size >= 2 {
		crossoverIndex := finder.rand.Index(finder.waypointCount) + 1
		shouldMutate := finder.rand.Index(100) < finder.mutationPercentage

		parent1 := finder.population.Route(0).chromosome()
		parent2 := finder.population.Route(1 + finder.rand.Index(size-1)).chromosome()

		if shouldMutate {
			mutation := finder.splicer.random(finder.startPaths, finder.waypointCount)

			result = append(result, mutation)
			result = append(result, finder.splicer.createOffspring(parent1, mutation, crossoverIndex))
			result = append(result, finder.splicer.createOffspring(mutation, parent1, crossoverIndex))
			result = append(result, finder.splicer.createOffspring(parent2, mutation, crossoverIndex))
			result = append(result, finder.splicer.createOffspring(mutation, parent2, crossoverIndex))
		} else {
			result = append(result, finder.splicer.createOffspring(parent1, parent2, crossoverIndex))
			result = append(result, finder.splicer.createOffspring(parent2, parent1, crossoverIndex))
		}
	}

	return result
}
