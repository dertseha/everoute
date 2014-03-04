package universe

type SolarSystem interface {
	Id() Id
	GalaxyId() GalaxyId
	RegionId() Id
	ConstellationId() Id
	Location() Location
	Security() TrueSecurity
	Jumps(jumpType string) []*Jump
	Costs() *TravelCostSum
}
