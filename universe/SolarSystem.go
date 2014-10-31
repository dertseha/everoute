package universe

// SolarSystem represents a system within a Universe. SolarSystem instances are immutable and meant to be extended
// only through extending the complete universe.
type SolarSystem interface {
	// Id returns the unique identifier of this solar system.
	Id() Id
	// GalaxyId returns the unique identifier of the galaxy this system is in.
	GalaxyId() GalaxyId
	// RegionId returns the unique identifier of the region this system is in.
	RegionId() Id
	// ConstellationId returns the unique identifier of the constellation this system is in.
	ConstellationId() Id
	// Location specifies the position of the star of this system within the universe.
	Location() Location
	// Security returns the security value of this solar system.
	Security() TrueSecurity
	// Jumps returns all known jumps of given type from this solar system. Returns an empty slice if none of this type
	// are known.
	Jumps(jumpType string) []Jump
	// Costs returns the costs of this solar system for travelling. These costs are typically used for continuing a
	// journey from this system.
	Costs() *TravelCostSum
}
