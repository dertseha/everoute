package universe

type solarSystemExtensionData struct {
	base        SolarSystem
	jumpBuilder []*JumpBuilder
}

func newSolarSystemExtensionData(base SolarSystem) *solarSystemExtensionData {
	result := &solarSystemExtensionData{
		base:        base,
		jumpBuilder: make([]*JumpBuilder, 0)}

	return result
}

func (data *solarSystemExtensionData) solarSystem() SolarSystem {
	return newExtendedSolarSystem(data)
}
