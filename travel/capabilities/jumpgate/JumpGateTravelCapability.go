package jumpgate

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"
)

type jumpGateTravelCapability struct {
	universe universe.Universe
}

func JumpGateTravelCapability(universe universe.Universe) travel.TravelCapability {
	return &jumpGateTravelCapability{universe: universe}
}

func (capability *jumpGateTravelCapability) NextPaths(origin *travel.Path) []*travel.Path {
	var solarSystem = capability.universe.SolarSystem(origin.Step().SolarSystemId())
	var jumps = solarSystem.Jumps(JumpType)
	var result = make([]*travel.Path, len(jumps))

	for i, jump := range jumps {
		var destination = capability.universe.SolarSystem(jump.DestinationId())
		var builder = travel.NewStepBuilder(destination.Id()).WithEnterCosts(jump.Costs()).WithEnterCosts(destination.Costs())

		result[i] = origin.Extend(builder.Build())
	}

	return result
}
