package universe

// SolarSystemExtension is a mutable extension to extend a solar system.
type SolarSystemExtension struct {
	data *solarSystemExtensionData
}

// Id returns the unique identifier of the solar system.
func (extension *SolarSystemExtension) Id() Id {
	return extension.data.base.Id()
}

// GalaxyId returns the unique identifier of the galaxy the solar system is in.
func (extension *SolarSystemExtension) GalaxyId() GalaxyId {
	return extension.data.base.GalaxyId()
}

// Location returns the position of the star of the solar system.
func (extension *SolarSystemExtension) Location() Location {
	return extension.data.base.Location()
}

// Security returns the security value of the solar system.
func (extension *SolarSystemExtension) Security() TrueSecurity {
	return extension.data.base.Security()
}

// AddJump adds the provided jump to the solar system.
func (extension *SolarSystemExtension) AddJump(jump Jump) {
	extension.data.jumps = append(extension.data.jumps, jump)
}

// BuildJump returns a JumpBuilder instance that can be used to define a new Jump instance.
// The returned builder will be already registered with the extension and it is not necessary to additionally
// build a Jump instance from the builder and register it in this extension.
func (extension *SolarSystemExtension) BuildJump(jumpType string, destinationId Id) *JumpBuilder {
	result := newJumpBuilder(jumpType, destinationId)

	extension.data.jumpBuilder = append(extension.data.jumpBuilder, result)

	return result
}

// AddCost adds the provided travel cost to the costs of the solar system.
func (extension *SolarSystemExtension) AddCost(cost TravelCost) {
	extension.data.costs = extension.data.costs.Add(SingleTravelCostSum(cost))
}
