package universe

type solarSystemExtensionData struct {
	base        SolarSystem
	jumpBuilder []*JumpBuilder
	costs       *TravelCostSum
}

func newSolarSystemExtensionData(base SolarSystem) *solarSystemExtensionData {
	result := &solarSystemExtensionData{
		base:        base,
		jumpBuilder: make([]*JumpBuilder, 0),
		costs:       EmptyTravelCostSum()}

	return result
}

func (data *solarSystemExtensionData) solarSystem() SolarSystem {
	return newExtendedSolarSystem(data)
}
