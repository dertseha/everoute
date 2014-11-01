package search

import (
	"fmt"

	"github.com/dertseha/everoute/travel/capabilities/jumpdrive"
	"github.com/dertseha/everoute/travel/capabilities/jumpgate"
	"github.com/dertseha/everoute/travel/rules/security"
	"github.com/dertseha/everoute/travel/rules/transitcount"
	"github.com/dertseha/everoute/universe"
)

func buildSolarSystems(builder *universe.UniverseBuilder) {
	for _, system := range SolarSystems {
		trueSec := universe.TrueSecurity(system.Security)
		galaxyId := universe.NewEdenId

		builder.AddSolarSystem(system.SolarSystemId, system.ConstellationId, system.RegionId, galaxyId,
			universe.NewSpecificLocation(system.X, system.Y, system.Z), trueSec)
	}
}

func getSolarSystemIdsByName() map[string]universe.Id {
	result := make(map[string]universe.Id)

	for _, system := range SolarSystems {
		result[system.Name] = system.SolarSystemId
	}

	return result
}

/*
func getJumpGateDestinationName(gate JumpGateData) string {
  destNameStart := strings.Index(gate.Name, "(") + 1
  destNameEnd := strings.Index(gate.Name, ")")

  return gate.Name[destNameStart:destNameEnd]
}
*/
func getJumpGateKey(fromSolarSystemId, toSolarSystemId universe.Id) string {
	return fmt.Sprintf("%d->%d", fromSolarSystemId, toSolarSystemId)
}

func getJumpGateLocations() map[string]universe.Location {
	result := make(map[string]universe.Location)
	//solarSystemIdsByName := getSolarSystemIdsByName()

	/*
	   for _, gate := range JumpGates {
	     destName := getJumpGateDestinationName(gate)
	     key := getJumpGateKey(gate.SolarSystemId, solarSystemIdsByName[destName])
	     location := universe.NewSpecificLocation(gate.X, gate.Y, gate.Z)

	     result[key] = location
	   }
	*/

	return result
}

func buildJumpGates(builder *universe.UniverseBuilder) {
	//jumpGateLocations := getJumpGateLocations()

	for _, jumpData := range SolarSystemJumps {
		extension1 := builder.ExtendSolarSystem(jumpData.FromSolarSystemId)
		var _ = extension1.BuildJump(jumpgate.JumpType, jumpData.ToSolarSystemId)

		extension2 := builder.ExtendSolarSystem(jumpData.ToSolarSystemId)
		var _ = extension2.BuildJump(jumpgate.JumpType, jumpData.FromSolarSystemId)

		//jumpBuilder.From(jumpGateLocations[getJumpGateKey(jumpData.FromSolarSystemId, jumpData.ToSolarSystemId)])
		//jumpBuilder.To(jumpGateLocations[getJumpGateKey(jumpData.ToSolarSystemId, jumpData.FromSolarSystemId)])
	}
}

func prepareUniverse() *universe.UniverseBuilder {
	builder := universe.New().Extend()

	buildSolarSystems(builder)
	buildJumpGates(builder)
	transitcount.ExtendUniverse(builder)
	security.ExtendUniverse(builder)

	jumpdrive.ExtendUniverse(builder, 10.0)

	return builder
}

func BuildHeimatar() universe.Universe {
	universeBuilder := prepareUniverse()
	verse := universeBuilder.Build()

	return verse
}
