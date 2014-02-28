package travel

type PathContest interface {
	Enter(path *Path) bool
}
