package universe

type Universe interface {
	Extend() *UniverseBuilder
	HasSolarSystem(id Id) bool
	SolarSystem(id Id) SolarSystem
	SolarSystemIds() []Id
}
