package capabilities

import "github.com/dertseha/everoute/travel"

type combiningTravelCapability struct {
	capabilities []travel.TravelCapability
}

// CombiningTravelCapability returns a travel capability that is based on a list of other travel capabilities.
// The returned travel capability returns the sum of all contained travel capabilities for continuing a path.
func CombiningTravelCapability(capabilities ...travel.TravelCapability) travel.TravelCapability {
	return &combiningTravelCapability{append(make([]travel.TravelCapability, 0, len(capabilities)), capabilities...)}
}

func (combine *combiningTravelCapability) NextPaths(origin travel.Path) []travel.Path {
	var result = make([]travel.Path, 0)

	for _, capability := range combine.capabilities {
		result = append(result, capability.NextPaths(origin)...)
	}

	return result
}
