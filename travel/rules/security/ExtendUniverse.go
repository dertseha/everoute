package security

import "github.com/dertseha/everoute/universe"

func ExtendUniverse(builder *universe.UniverseBuilder) {
	var solarSystemIds []universe.Id = builder.SolarSystemIds()

	for _, id := range solarSystemIds {
		var extension = builder.ExtendSolarSystem(id)
		var security = extension.Security()

		extension.AddCost(travelCost(security, 1))
	}
}
