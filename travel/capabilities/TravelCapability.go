package capabilities

import "github.com/dertseha/everoute/travel"

type TravelCapability interface {
	NextPaths(origin *travel.Path) []*travel.Path
}
