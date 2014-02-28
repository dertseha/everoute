package transitcount

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"
)

func ExtendUniverse(builder *universe.UniverseBuilder) {
	var solarSystemIds = builder.SolarSystemIds()
	var cost = travel.AddingTravelCost(CostType, 1)

	for _, id := range solarSystemIds {
		extension := builder.ExtendSolarSystem(id)
		extension.AddCost(cost)
	}
}
