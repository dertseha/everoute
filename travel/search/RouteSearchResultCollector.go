package search

type RouteSearchResultCollector interface {
	Collect(route *Route)
}
