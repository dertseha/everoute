package universe

type emptySolarSystem struct {
	id              Id
	constellationId Id
	regionId        Id
	galaxyId        GalaxyId
	location        Location
	security        TrueSecurity
}

// newSolarSystem returns a SolarSystem instance with provided parameters that is empty.
func newSolarSystem(id Id, constellationId Id, regionId Id, galaxyId GalaxyId, location Location, security TrueSecurity) SolarSystem {
	var system = &emptySolarSystem{
		id:              id,
		constellationId: constellationId,
		regionId:        regionId,
		galaxyId:        galaxyId,
		location:        location,
		security:        security}

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

func (this *emptySolarSystem) GalaxyId() GalaxyId {
	return this.galaxyId
}

func (this *emptySolarSystem) Location() Location {
	return this.location
}

func (this *emptySolarSystem) Security() TrueSecurity {
	return this.security
}

func (this *emptySolarSystem) Jumps(jumpType string) []Jump {
	return nil
}

func (this *emptySolarSystem) Costs() *TravelCostSum {
	return EmptyTravelCostSum()
}
