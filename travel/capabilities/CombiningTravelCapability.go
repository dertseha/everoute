package capabilities

import "github.com/dertseha/everoute/travel"

type combiningTravelCapability struct {
	capabilities []TravelCapability
}

func CombiningTravelCapability(capabilities []TravelCapability) TravelCapability {
	return &combiningTravelCapability{append(make([]TravelCapability, 0, len(capabilities)), capabilities...)}
}

func (combine *combiningTravelCapability) NextPaths(origin *travel.Path) []*travel.Path {
	var result = make([]*travel.Path, 0)

	for _, capability := range combine.capabilities {
		result = append(result, capability.NextPaths(origin)...)
	}

	return result
}
