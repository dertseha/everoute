package travel

type ruleBasedPathContest struct {
	rule                  TravelRule
	pathsByDestinationKey map[string]Path
}

// RuleBasedPathContest returns a PathContest that is based on travel rules.
// This contest will only accept paths that have not yet been superseded by cheaper paths.
func RuleBasedPathContest(rule TravelRule) PathContest {
	contest := &ruleBasedPathContest{
		rule: rule,
		pathsByDestinationKey: make(map[string]Path)}

	return contest
}

func (contest *ruleBasedPathContest) Enter(path Path) bool {
	var result = true
	var destinationKey = path.DestinationKey()
	var oldPath, existing = contest.pathsByDestinationKey[destinationKey]

	if !path.IsStart() && !contest.isPathStillCurrent(path.Previous()) {
		result = false
	} else if existing && (oldPath != path) && (contest.rule.Compare(path.CostSum(), oldPath.CostSum()) >= 0) {
		result = false
	} else {
		contest.pathsByDestinationKey[destinationKey] = path
	}

	return result
}

func (contest *ruleBasedPathContest) isPathStillCurrent(path Path) bool {
	var entry = path

	for (contest.pathsByDestinationKey[entry.DestinationKey()] == entry) && !entry.IsStart() {
		entry = entry.Previous()
	}

	return entry.IsStart()
}
