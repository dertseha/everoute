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

func FindPath(start travel.Path, capability travel.TravelCapability, rule travel.TravelRule,
	criterion SearchCriterion, collector PathSearchResultCollector, searchDone chan int) {
	var ruleContest = travel.RuleBasedPathContest(rule)
	var contestCallback = make(chan DeferredPathContestRequest)
	var contest = DeferredPathContest(ruleContest, contestCallback)
	var optimizingCapability = newOptimizingTravelCapability(capability, contest)

	var search = func() {
		var leafRequests = make(chan leafRequest)
		var leavesActive = 0

		var leaf func(path travel.Path)
		var startLeaf = func(path travel.Path) {
			leavesActive++
			go leaf(path)
		}

		leaf = func(path travel.Path) {
			var nextPaths = optimizingCapability.NextPaths(path)

			for _, nextPath := range nextPaths {
				var request = &leafFoundRequest{
					criterion:   criterion,
					collector:   collector,
					path:        nextPath,
					leafStarter: startLeaf}

				leafRequests <- request
			}
			leafRequests <- leafFinishedRequest{leavesActive: &leavesActive}
		}

		startLeaf(start)
		for leavesActive > 0 {
			select {
			case request := <-contestCallback:
				request.Process()
			case request := <-leafRequests:
				request.Process()
			}
		}

		searchDone <- 0
		close(contestCallback)
	}

	go search()
}
