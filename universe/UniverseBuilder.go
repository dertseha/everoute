package universe

import (
	"fmt"
	"sort"
)

// UniverseBuilder is a helper to extend an existing Universe instance with further content.
type UniverseBuilder struct {
	base                  Universe
	solarSystemExtensions map[Id]*solarSystemExtensionData
}

func newUniverseBuilder(universe Universe) *UniverseBuilder {
	var builder = &UniverseBuilder{
		base: universe,
		solarSystemExtensions: make(map[Id]*solarSystemExtensionData)}

	return builder
}

// Build returns an immutable Universe instance that is based on the current values of the builder.
func (builder *UniverseBuilder) Build() Universe {
	var universe = &extendedUniverse{
		base:         builder.base,
		solarSystems: make(map[Id]SolarSystem)}

	for id, data := range builder.solarSystemExtensions {
		universe.solarSystems[id] = data.solarSystem()
	}

	return universe
}

// AddSolarSystem adds a basic solar system to the future universe.
// This method panics if the universe already contains a solar system with this ID.
func (builder *UniverseBuilder) AddSolarSystem(id Id, constellationId Id, regionId Id, galaxyId GalaxyId, location Location, security TrueSecurity) {
	var data, exists = builder.solarSystemExtensions[id]

	if !exists && !builder.base.HasSolarSystem(id) {
		data = newSolarSystemExtensionData(newSolarSystem(id, constellationId, regionId, galaxyId, location, security))
		builder.solarSystemExtensions[id] = data
	} else {
		panic(fmt.Sprintf("Solar System %d already exists.", id))
	}
}

// ExtendSolarSystem returns a SolarSystemExtension instance meant to extend a solar system already existing in the
// universe.
// This method panics if the solar sytem can not be found.
func (builder *UniverseBuilder) ExtendSolarSystem(id Id) (extension SolarSystemExtension) {
	var data, exists = builder.solarSystemExtensions[id]

	if !exists && builder.base.HasSolarSystem(id) {
		data = newSolarSystemExtensionData(builder.base.SolarSystem(id))
		builder.solarSystemExtensions[id] = data
	}
	_, exists = builder.solarSystemExtensions[id]
	if exists {
		extension = SolarSystemExtension{}
		extension.data = data
	} else {
		panic(fmt.Sprintf("Solar System %d doesn't exist.", id))
	}

	return
}

// SolarSystemIds returns a sorted slice of all solar system ID values currently known to the builder. Both the
// underlying universe and the already added solar systems are taken into account.
func (builder *UniverseBuilder) SolarSystemIds() []Id {
	var baseIds = builder.base.SolarSystemIds()
	var ids = make([]Id, 0, len(builder.solarSystemExtensions))

	for id := range builder.solarSystemExtensions {
		ids = append(ids, id)
	}
	ids = append(ids, baseIds...)
	sort.Sort(IdOrder(ids))

	return ids
}
