package search

import "github.com/dertseha/everoute/travel"

type routeList struct {
	rule   travel.TravelRule
	routes []*Route
}

func emptyRouteList(rule travel.TravelRule) *routeList {
	return newRouteList(rule, 0)
}

func newRouteList(rule travel.TravelRule, size int) *routeList {
	result := &routeList{
		rule:   rule,
		routes: make([]*Route, size)}

	return result
}

func (list *routeList) Size() int {
	return len(list.routes)
}

func (list *routeList) Route(index int) *Route {
	return list.routes[index]
}

func (list *routeList) Add(route *Route) *routeList {
	result := newRouteList(list.rule, len(list.routes)+1)
	costSum := route.CostSum()
	position := len(list.routes) - 1
	isBetter := true

	for isBetter && position >= 0 {
		other := list.routes[position]

		isBetter = list.rule.Compare(costSum, other.CostSum()) < 0
		if isBetter {
			result.routes[position+1] = other
			position--
		}
	}
	result.routes[position+1] = route
	copy(result.routes[0:position+1], list.routes[0:position+1])

	return result
}

func (list *routeList) Limit(size int) *routeList {
	result := list

	if size < list.Size() {
		result = &routeList{
			rule:   list.rule,
			routes: list.routes[0:size]}
	}

	return result
}
