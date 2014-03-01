package search

import "github.com/dertseha/everoute/travel"

type PathSearchResultCollector interface {
	Collect(result travel.Path)
	Results() []travel.Path
}
