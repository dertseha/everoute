package universe

type emptySolarSystem struct {
	id              Id
	constellationId Id
	regionId        Id
	galaxyId        Id
	location        Location
}

func newSolarSystem(id Id, constellationId Id, regionId Id, galaxyId Id, location Location) SolarSystem {
	var system = &emptySolarSystem{
		id:              id,
		constellationId: constellationId,
		regionId:        regionId,
		galaxyId:        galaxyId,
		location:        location}

	return system
}

func (this *emptySolarSystem) Id() Id {
	return this.id
}

func (this *emptySolarSystem) ConstellationId() Id {
	return this.constellationId
}

func (this *emptySolarSystem) RegionId() Id {
	return this.regionId
}

func (this *emptySolarSystem) GalaxyId() Id {
	return this.galaxyId
}

func (this *emptySolarSystem) Location() Location {
	return this.location
}
