package universe

// Universe is an interface for accessing universe data.
type Universe interface {
	// Extend allows the user to extend this universe with further additions. Extending a universe does not modify the
	// originating universe, which can be used without interruption. The new universe will contain everything from the
	// original universe plus any addition.
	Extend() *UniverseBuilder
	// HasSolarSystem returns true if it contains a solar system with the provided ID.
	HasSolarSystem(id Id) bool
	// SolarSystem returns the solar system instance for given ID. This method will panic if the ID is unknown.
	SolarSystem(id Id) SolarSystem
	// SolarSystemIds returns a slice of all IDs this universe currently has. The returned slice will contain the IDs in
	// a sorted way.
	SolarSystemIds() []Id
}
