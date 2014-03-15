package search

import "github.com/dertseha/everoute/travel"

type leafRequest interface {
	Process()
}

type leafFinishedRequest struct {
	leavesActive *int
}

func (request leafFinishedRequest) Process() {
	*request.leavesActive--
}

type leafFoundRequest struct {
	criterion   SearchCriterion
	collector   PathSearchResultCollector
	path        travel.Path
	leafStarter func(travel.Path)
}

func (request leafFoundRequest) Process() {
	if request.criterion.IsDesired(request.path) {
		request.collector.Collect(request.path)
	}
	if request.criterion.ShouldSearchContinueWith(request.path, request.collector.Results()) {
		request.leafStarter(request.path)
	}
}

func startOptimizingContest(rule travel.TravelRule, capability travel.TravelCapability, contestQuit chan int) travel.TravelCapability {
	var ruleContest = travel.RuleBasedPathContest(rule)
	var contestCallback = make(chan DeferredPathContestRequest, 50)
	var contest = DeferredPathContest(ruleContest, contestCallback)
	var optimizingCapability = newOptimizingTravelCapability(capability, contest)

	var runContest = func() {
		var active = true

		for active {
			select {
			case request := <-contestCallback:
				request.Process()
			case <-contestQuit:
				active = false
			}
		}
		close(contestCallback)
		close(contestQuit)
	}
	go runContest()

	return optimizingCapability
}

func FindPath(start travel.Path, capability travel.TravelCapability, rule travel.TravelRule,
	criterion SearchCriterion, collector PathSearchResultCollector, searchDone func()) {
	var contestQuit = make(chan int)
	var optimizingCapability = startOptimizingContest(rule, capability, contestQuit)

	var search = func() {
		var leafRequests = make(chan leafRequest, 50)
		var leavesActive = 0

		var leaf func(path travel.Path)
		var startLeaf = func(path travel.Path) {
			leavesActive++
			go leaf(path)
		}
		var foundLeaf = func(path travel.Path) {
			var request = &leafFoundRequest{
				criterion:   criterion,
				collector:   collector,
				path:        path,
				leafStarter: startLeaf}

			leafRequests <- request
		}

		leaf = func(path travel.Path) {
			var nextPaths = optimizingCapability.NextPaths(path)

			for _, nextPath := range nextPaths {
				foundLeaf(nextPath)
			}
			leafRequests <- &leafFinishedRequest{leavesActive: &leavesActive}
		}

		leavesActive++
		foundLeaf(start)
		leafRequests <- &leafFinishedRequest{leavesActive: &leavesActive}

		for leavesActive > 0 {
			select {
			case request := <-leafRequests:
				request.Process()
			}
		}

		go searchDone()
		contestQuit <- 0
		close(leafRequests)
	}

	go search()
}
