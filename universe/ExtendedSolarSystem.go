package universe

type extendedSolarSystem struct {
	base  SolarSystem
	jumps []*Jump
}

func newExtendedSolarSystem(data *solarSystemExtensionData) SolarSystem {
	result := &extendedSolarSystem{
		base:  data.base,
		jumps: make([]*Jump, len(data.jumpBuilder))}

	for i, builder := range data.jumpBuilder {
		result.jumps[i] = builder.Build()
	}

	return result
}

func (this *extendedSolarSystem) Id() Id {
	return this.base.Id()
}

func (this *extendedSolarSystem) ConstellationId() Id {
	return this.base.ConstellationId()
}

func (this *extendedSolarSystem) RegionId() Id {
	return this.base.RegionId()
}

func (this *extendedSolarSystem) GalaxyId() GalaxyId {
	return this.base.GalaxyId()
}

func (this *extendedSolarSystem) Location() Location {
	return this.base.Location()
}

func (this *extendedSolarSystem) Security() TrueSecurity {
	return this.base.Security()
}
