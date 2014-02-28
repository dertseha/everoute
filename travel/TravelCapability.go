package travel

type TravelCapability interface {
	NextPaths(origin *Path) []*Path
}
