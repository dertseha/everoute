package universe

type extendedSolarSystem struct {
	base  SolarSystem
	jumps map[string][]Jump
	costs *TravelCostSum
}

func newExtendedSolarSystem(data *solarSystemExtensionData) SolarSystem {
	result := &extendedSolarSystem{
		base:  data.base,
		jumps: make(map[string][]Jump),
		costs: data.costs}
	addJump := func(jump Jump) {
		jumps, existing := result.jumps[jump.Type()]

		if !existing {
			jumps = make([]Jump, 0, 1)
		}
		result.jumps[jump.Type()] = append(jumps, jump)
	}

	for _, jump := range data.jumps {
		addJump(jump)
	}
	for _, builder := range data.jumpBuilder {
		jump := builder.Build()
		addJump(jump)
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

func (this *extendedSolarSystem) Jumps(jumpType string) []Jump {
	var jumps, existing = this.jumps[jumpType]
	var result = this.base.Jumps(jumpType)

	if existing {
		result = append(result, jumps...)
	}

	return result
}

func (this *extendedSolarSystem) Costs() *TravelCostSum {
	return this.base.Costs().Add(this.costs)
}
