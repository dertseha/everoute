package universe

type SolarSystem interface {
	Id() Id
	GalaxyId() Id
	RegionId() Id
	ConstellationId() Id
	Location() Location
}
