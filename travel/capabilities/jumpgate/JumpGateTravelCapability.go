package jumpgate

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/travel/rules/warpdistance"
	"github.com/dertseha/everoute/universe"
)

type jumpGateTravelCapability struct {
	universe universe.Universe

	avoidHighSec bool
}

// JumpGateTravelCapability returns a travel capability that tries to extend paths by the use of jump gates.
func JumpGateTravelCapability(universe universe.Universe, avoidHighSec bool) travel.TravelCapability {
	return &jumpGateTravelCapability{
		universe:     universe,
		avoidHighSec: avoidHighSec}
}

func (capability *jumpGateTravelCapability) NextPaths(origin travel.Path) []travel.Path {
	solarSystem := capability.universe.SolarSystem(origin.Step().SolarSystemId())
	jumps := solarSystem.Jumps(JumpType)
	result := make([]travel.Path, 0, len(jumps))

	for _, jump := range jumps {
		destination := capability.universe.SolarSystem(jump.DestinationId())

		if !capability.avoidHighSec || !destination.Security().IsHighSec() {
			warpDistance := origin.Step().Location().DistanceTo(jump.SourceLocation())
			warpCosts := universe.SingleTravelCostSum(warpdistance.Cost(warpDistance))
			builder := travel.NewStepBuilder(destination.Id())

			builder.WithEnterCosts(warpCosts.Add(jump.Costs()))
			builder.WithContinueCosts(destination.Costs())
			builder.To(jump.DestinationLocation())

			result = append(result, origin.Extend(builder.Build()))
		}
	}

	return result
}
