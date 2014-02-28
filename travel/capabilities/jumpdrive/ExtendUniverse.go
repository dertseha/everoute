package jumpdrive

import "github.com/dertseha/everoute/universe"

const MetersPerAu float64 = 149597870700
const MetersPerLy float64 = MetersPerAu * 63241

func ExtendUniverse(builder *universe.UniverseBuilder, limit float64) {
	highSecSystems := make([]universe.SolarSystemExtension, 0)
	nonHighSecSystems := make([]universe.SolarSystemExtension, 0)
	ids := builder.SolarSystemIds()

	for _, id := range ids {
		extension := builder.ExtendSolarSystem(id)

		if extension.GalaxyId() == universe.NewEdenId {
			if extension.Security().IsHighSec() {
				highSecSystems = append(highSecSystems, extension)
			} else {
				nonHighSecSystems = append(nonHighSecSystems, extension)
			}
		}
	}

	createJumpsFromHighSec := func(source *universe.SolarSystemExtension) {
		for _, other := range nonHighSecSystems {
			distance := source.Location().DistanceTo(other.Location()) / MetersPerLy
			if distance <= limit {
				source.AddJump(JumpType, other.Id())
			}
		}
	}

	createJumpsBetween := func(source *universe.SolarSystemExtension, startIndex int) {
		indexLimit := len(nonHighSecSystems)

		for i := startIndex; i < indexLimit; i++ {
			other := nonHighSecSystems[i]
			distance := source.Location().DistanceTo(other.Location()) / MetersPerLy

			if distance <= limit {
				source.AddJump(JumpType, other.Id())
				other.AddJump(JumpType, source.Id())
			}
		}
	}

	for _, source := range highSecSystems {
		createJumpsFromHighSec(&source)
	}
	for i, source := range nonHighSecSystems {
		createJumpsBetween(&source, i+1)
	}
}
