package search

import "github.com/dertseha/everoute/travel"

type simplePathSearchResultCollector struct {
	results map[string]*travel.Path
}

func SimplePathSearchResultCollector() PathSearchResultCollector {
	return &simplePathSearchResultCollector{results: make(map[string]*travel.Path)}
}

func (collector *simplePathSearchResultCollector) Collect(path *travel.Path) {
	collector.results[path.DestinationKey()] = path
}

func (collector *simplePathSearchResultCollector) Results() []*travel.Path {
	var result = make([]*travel.Path, 0, len(collector.results))

	for _, path := range collector.results {
		result = append(result, path)
	}

	return result
}
