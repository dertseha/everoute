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

func reduceCosts(costs []universe.TravelCost, f func(bool, universe.TravelCost) bool, init bool) bool {
	var result = init

	for _, cost := range costs {
		result = f(result, cost)
	}

	return result
}

func (capability *jumpDriveTravelCapability) jumpsWithinLimit(jumps []*universe.Jump) []*universe.Jump {
	var jumpsWithinLimit = make([]*universe.Jump, 0, len(jumps))

	for _, jump := range jumps {
		if reduceCosts(jump.Costs(), func(accumulator bool, cost universe.TravelCost) bool {
			return accumulator && ((cost.Type() != jumpdistance.CostType) || (cost.Value() <= capability.distanceLimitLy))
		}, true) {
			jumpsWithinLimit = append(jumpsWithinLimit, jump)
		}
	}

	return jumpsWithinLimit
}

func (capability *jumpDriveTravelCapability) NextPaths(origin travel.Path) []travel.Path {
	var solarSystem = capability.universe.SolarSystem(origin.Step().SolarSystemId())
	var jumps = capability.jumpsWithinLimit(solarSystem.Jumps(JumpType))
	var result = make([]travel.Path, len(jumps))

	for i, jump := range jumps {
		var destination = capability.universe.SolarSystem(jump.DestinationId())
		var builder = travel.NewStepBuilder(destination.Id()).WithEnterCosts(jump.Costs()).WithContinueCosts(destination.Costs())

		result[i] = origin.Extend(builder.Build())
	}

	return result
}
