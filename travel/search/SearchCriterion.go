package search

import "github.com/dertseha/everoute/travel"

type SearchCriterion interface {
	IsDesired(path *travel.Path) bool
	ShouldSearchContinueWith(path *travel.Path, results []*travel.Path) bool
}
