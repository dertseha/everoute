package universe

import (
	"fmt"
	"sort"
)

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

func (builder *UniverseBuilder) Build() Universe {
	var universe = &extendedUniverse{
		base:         builder.base,
		solarSystems: make(map[Id]SolarSystem)}

	for id, data := range builder.solarSystemExtensions {
		universe.solarSystems[id] = data.solarSystem()
	}

	return universe
}

func (builder *UniverseBuilder) AddSolarSystem(id Id, constellationId Id, regionId Id, galaxyId Id, location Location) {
	var data, ok = builder.solarSystemExtensions[id]

	if !ok {
		data = newSolarSystemExtensionData(newSolarSystem(id, constellationId, regionId, galaxyId, location))
		builder.solarSystemExtensions[id] = data
	} else {
		panic(fmt.Sprintf("Solar System %d already exists.", id))
	}
}

func (builder *UniverseBuilder) ExtendSolarSystem(id Id) (extension SolarSystemExtension) {
	var data, ok = builder.solarSystemExtensions[id]

	if !ok && builder.base.HasSolarSystem(id) {
		data = newSolarSystemExtensionData(builder.base.SolarSystem(id))
		builder.solarSystemExtensions[id] = data
	}
	_, ok = builder.solarSystemExtensions[id]
	if ok {
		extension = SolarSystemExtension{}
		extension.data = data
	} else {
		panic(fmt.Sprintf("Solar System %d doesn't exist.", id))
	}

	return
}

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
