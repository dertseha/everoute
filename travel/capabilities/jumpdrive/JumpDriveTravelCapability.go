package jumpdrive

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/travel/rules/jumpdistance"
	"github.com/dertseha/everoute/universe"
)

type jumpDriveTravelCapability struct {
	universe        universe.Universe
	distanceLimitLy float64
}

func JumpDriveTravelCapability(universe universe.Universe, distanceLimitLy float64) travel.TravelCapability {
	return &jumpDriveTravelCapability{universe: universe, distanceLimitLy: distanceLimitLy}
}

func (capability *jumpDriveTravelCapability) NextPaths(origin travel.Path) []travel.Path {
	var solarSystem = capability.universe.SolarSystem(origin.Step().SolarSystemId())
	var jumps = solarSystem.Jumps(JumpType)
	var result = make([]travel.Path, 0, len(jumps))
	var nullCost = jumpdistance.NullCost()

	for _, jump := range jumps {
		var cost = jump.Costs().Cost(nullCost)

		if cost.Value() <= capability.distanceLimitLy {
			var destination = capability.universe.SolarSystem(jump.DestinationId())
			var builder = travel.NewStepBuilder(destination.Id()).WithEnterCosts(jump.Costs()).WithContinueCosts(destination.Costs())

			result = append(result, origin.Extend(builder.Build()))
		}
	}

	return result
}
