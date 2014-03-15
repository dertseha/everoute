package search

import (
	"sync"

	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/util"
)

type pathSearchAgentResult interface {
	BestPath() travel.Path
	Path(destinationKey string, rand util.Randomizer) travel.Path
}

type pathSearchAgentListener interface {
	searchFailed()
	searchCompleted(result pathSearchAgentResult)
}

type pathSearchAgentState interface {
	request(listener pathSearchAgentListener)
}

type basicPathSearchAgentResult struct {
	bestPath    travel.Path
	paths       map[string]travel.Path
	pathsAsList []travel.Path
}

func (result *basicPathSearchAgentResult) BestPath() travel.Path {
	return result.bestPath
}

func (result *basicPathSearchAgentResult) Path(destinationKey string, rand util.Randomizer) travel.Path {
	path, existing := result.paths[destinationKey]

	if !existing {
		path = result.pathsAsList[rand.Index(len(result.pathsAsList))]
	}

	return path
}

type pathSearchAgent struct {
	state     pathSearchAgentState
	stateLock sync.Locker
}

type pathSearchAgentActiveState struct {
	rule travel.TravelRule

	requests []pathSearchAgentListener

	results       map[string]travel.Path
	resultsAsList []travel.Path
}

func (state *pathSearchAgentActiveState) request(listener pathSearchAgentListener) {
	state.requests = append(state.requests, listener)
}

func (state *pathSearchAgentActiveState) Collect(result travel.Path) {
	state.results[result.DestinationKey()] = result
	state.resultsAsList = make([]travel.Path, 0, len(state.results))
	for _, path := range state.results {
		state.resultsAsList = append(state.resultsAsList, path)
	}
}

func (state *pathSearchAgentActiveState) Results() []travel.Path {
	return state.resultsAsList
}

type pathSearchAgentFailedState struct{}

func (state *pathSearchAgentFailedState) request(listener pathSearchAgentListener) {
	go listener.searchFailed()
}

type pathSearchAgentCompletedState struct {
	result pathSearchAgentResult
}

func (state *pathSearchAgentCompletedState) request(listener pathSearchAgentListener) {
	go listener.searchCompleted(state.result)
}

func newPathSearchAgent(startPath travel.Path, capability travel.TravelCapability,
	rule travel.TravelRule, criterion SearchCriterion) *pathSearchAgent {
	state := &pathSearchAgentActiveState{
		rule: rule,

		requests: make([]pathSearchAgentListener, 0),

		results:       make(map[string]travel.Path),
		resultsAsList: make([]travel.Path, 0)}
	agent := &pathSearchAgent{state: state, stateLock: &sync.Mutex{}}

	FindPath(startPath, capability, rule, criterion, state, func() { agent.searchDone(state) })

	return agent
}

func (agent *pathSearchAgent) Request(listener pathSearchAgentListener) {
	defer agent.stateLock.Unlock()
	agent.stateLock.Lock()

	agent.state.request(listener)
}

func (agent *pathSearchAgent) searchDone(state *pathSearchAgentActiveState) {
	defer agent.stateLock.Unlock()
	agent.stateLock.Lock()

	if len(state.resultsAsList) > 0 {
		result := &basicPathSearchAgentResult{
			bestPath:    nil,
			paths:       state.results,
			pathsAsList: state.resultsAsList}

		for _, path := range state.resultsAsList {
			if (result.bestPath == nil) || (state.rule.Compare(path.CostSum(), result.bestPath.CostSum()) < 0) {
				result.bestPath = path
			}
		}

		agent.state = &pathSearchAgentCompletedState{result: result}
	} else {
		agent.state = &pathSearchAgentFailedState{}
	}
	for _, listener := range state.requests {
		agent.state.request(listener)
	}
}
