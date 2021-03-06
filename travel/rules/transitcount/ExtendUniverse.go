package transitcount

import (
	"github.com/dertseha/everoute/universe"
)

// ExtendUniverse adds to all solar systems the cost of one transit value.
func ExtendUniverse(builder *universe.UniverseBuilder) {
	var solarSystemIds = builder.SolarSystemIds()
	var cost = universe.AddingTravelCost(CostType, 1)

	for _, id := range solarSystemIds {
		extension := builder.ExtendSolarSystem(id)
		extension.AddCost(cost)
	}
}
