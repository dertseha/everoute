package search

import "github.com/dertseha/everoute/travel"

// PathSearchResultCollector describes an entity that keeps results from a search.
type PathSearchResultCollector interface {
	// Collect receives a found path
	Collect(result travel.Path)
	// Results() returns a list of paths that have been kept so far.
	Results() []travel.Path
}
