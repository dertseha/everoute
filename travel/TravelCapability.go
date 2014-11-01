package travel

// TravelCapability describes a means to continue a Path.
type TravelCapability interface {
	// NextPaths receives an origin from which this capability shall calculate the next possible Paths. The paths in
	// the returned slice should be based on the passed origin path.
	NextPaths(origin Path) []Path
}
