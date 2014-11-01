package search

import "github.com/dertseha/everoute/travel"

type simplePathSearchResultCollector struct {
	results map[string]travel.Path
}

// SimplePathSearchResultCollector returns a collector instance that keeps all paths by their destination key.
// If a path with an already known destination key is collected, the new instance is kept and the old one discarded.
func SimplePathSearchResultCollector() PathSearchResultCollector {
	return &simplePathSearchResultCollector{results: make(map[string]travel.Path)}
}

func (collector *simplePathSearchResultCollector) Collect(path travel.Path) {
	collector.results[path.DestinationKey()] = path
}

func (collector *simplePathSearchResultCollector) Results() []travel.Path {
	var result = make([]travel.Path, 0, len(collector.results))

	for _, path := range collector.results {
		result = append(result, path)
	}

	return result
}
