package jumpgate

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/travel/rules/warpdistance"
	"github.com/dertseha/everoute/universe"
)

type jumpGateTravelCapability struct {
	universe universe.Universe
}

// JumpGateTravelCapability returns a travel capability that tries to extend paths by the use of jump gates.
func JumpGateTravelCapability(universe universe.Universe) travel.TravelCapability {
	return &jumpGateTravelCapability{universe: universe}
}

func (capability *jumpGateTravelCapability) NextPaths(origin travel.Path) []travel.Path {
	var solarSystem = capability.universe.SolarSystem(origin.Step().SolarSystemId())
	var jumps = solarSystem.Jumps(JumpType)
	var result = make([]travel.Path, len(jumps))

	for i, jump := range jumps {
		destination := capability.universe.SolarSystem(jump.DestinationId())
		warpDistance := origin.Step().Location().DistanceTo(jump.SourceLocation())
		warpCosts := universe.SingleTravelCostSum(warpdistance.Cost(warpDistance))
		builder := travel.NewStepBuilder(destination.Id())

		builder.WithEnterCosts(warpCosts.Add(jump.Costs()))
		builder.WithContinueCosts(destination.Costs())
		builder.To(jump.DestinationLocation())

		result[i] = origin.Extend(builder.Build())
	}

	return result
}
