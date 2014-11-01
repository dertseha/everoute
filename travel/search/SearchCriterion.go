package search

import "github.com/dertseha/everoute/travel"

// SearchCriterion describes a criterion for finding a path to a destination.
type SearchCriterion interface {
	// IsDesired should return true if the provided path is one that fulfills the criterion.
	IsDesired(path travel.Path) bool
	// ShouldSearchContinueWith returns true if the path (compared to the list of already found paths) is suitable for
	// continuing the search.
	ShouldSearchContinueWith(path travel.Path, results []travel.Path) bool
}
