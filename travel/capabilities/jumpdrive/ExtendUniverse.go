package jumpdrive

import (
	"github.com/dertseha/everoute/universe"
	"github.com/dertseha/everoute/util"
)

// ExtendUniverse uses the provided builder to add jump drive jumps between all applicable solar systems that are
// at most the given limit in light years apart.
// Only the NewEden galaxy is considered, jumps are created from high-sec systems as well as between any suitable
// non-high-sec system pair.
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
			distance := source.Location().DistanceTo(other.Location()) / util.MetersPerLy
			if distance <= limit {
				source.AddJump(Jump(other.Id(), distance))
			}
		}
	}

	createJumpsBetween := func(source *universe.SolarSystemExtension, startIndex int) {
		indexLimit := len(nonHighSecSystems)

		for i := startIndex; i < indexLimit; i++ {
			other := nonHighSecSystems[i]
			distance := source.Location().DistanceTo(other.Location()) / util.MetersPerLy

			if distance <= limit {
				source.AddJump(Jump(other.Id(), distance))
				other.AddJump(Jump(source.Id(), distance))
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
