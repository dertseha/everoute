package travel

// PathContest represents a competition between path instances according to their costs.
type PathContest interface {
	// Enter returns true if the provided path is accepted based on its costs.
	// When it returns false, the provided path is considered to be too expensive.
	Enter(path Path) bool
}
