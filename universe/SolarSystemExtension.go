package universe

type SolarSystemExtension struct {
	data *solarSystemExtensionData
}

func (extension *SolarSystemExtension) Id() Id {
	return extension.data.base.Id()
}

func (extension *SolarSystemExtension) GalaxyId() Id {
	return extension.data.base.GalaxyId()
}

func (extension *SolarSystemExtension) Location() Location {
	return extension.data.base.Location()
}

func (extension *SolarSystemExtension) AddJump(jumpType string, destinationId Id) *JumpBuilder {
	result := newJumpBuilder(jumpType, destinationId)

	extension.data.jumpBuilder = append(extension.data.jumpBuilder, result)

	return result
}
