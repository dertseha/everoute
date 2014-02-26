package universe

import "sort"

type extendedUniverse struct {
	base         Universe
	solarSystems map[Id]SolarSystem
}

func (this *extendedUniverse) Extend() *UniverseBuilder {
	return newUniverseBuilder(this)
}

func (this *extendedUniverse) HasSolarSystem(id Id) bool {
	_, ok := this.solarSystems[id]

	return ok || this.base.HasSolarSystem(id)
}

func (this *extendedUniverse) SolarSystem(id Id) SolarSystem {
	system, ok := this.solarSystems[id]

	if !ok {
		system = this.base.SolarSystem(id)
	}

	return system
}

func (this *extendedUniverse) SolarSystemIds() []Id {
	var baseIds = this.base.SolarSystemIds()
	var ids = make([]Id, 0, len(this.solarSystems))

	for id := range this.solarSystems {
		ids = append(ids, id)
	}
	ids = append(ids, baseIds...)
	sort.Sort(IdOrder(ids))

	return ids
}
