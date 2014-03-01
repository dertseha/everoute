package search

import "github.com/dertseha/everoute/travel"

type DeferredPathContestRequest interface {
	Process()
}

type deferredPathContestRequest struct {
	contest travel.PathContest
	path    travel.Path
	result  chan bool
}

func (request *deferredPathContestRequest) Process() {
	request.result <- request.contest.Enter(request.path)
	close(request.result)
}

type deferredPathContest struct {
	contest  travel.PathContest
	callback chan DeferredPathContestRequest
}

func DeferredPathContest(contest travel.PathContest, callback chan DeferredPathContestRequest) travel.PathContest {
	var deferredContest = &deferredPathContest{
		contest:  contest,
		callback: callback}

	return deferredContest
}

func (contest *deferredPathContest) Enter(path travel.Path) bool {
	request := &deferredPathContestRequest{
		contest: contest.contest,
		path:    path,
		result:  make(chan bool)}

	contest.callback <- request

	return <-request.result
}
