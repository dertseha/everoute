package travel

type ruleBasedPathContest struct {
	rule                  TravelRule
	pathsByDestinationKey map[string]*Path
}

func RuleBasedPathContest(rule TravelRule) PathContest {
	contest := &ruleBasedPathContest{
		rule: rule,
		pathsByDestinationKey: make(map[string]*Path)}

	return contest
}

func (contest *ruleBasedPathContest) Enter(path *Path) bool {
	var result = true
	var destinationKey = path.DestinationKey()
	var oldPath, existing = contest.pathsByDestinationKey[destinationKey]

	if existing && (contest.rule.Compare(path.CostSum(), oldPath.CostSum()) >= 0) {
		result = false
	} else if !path.IsStart() && !contest.isPathStillCurrent(path.Previous()) {
		result = false
	} else {
		contest.pathsByDestinationKey[destinationKey] = path
	}

	return result
}

func (contest *ruleBasedPathContest) isPathStillCurrent(path *Path) bool {
	var entry = path
	var result = false
	var isEntryCurrent = func() bool {
		var key = entry.DestinationKey()
		var path, existing = contest.pathsByDestinationKey[key]

		return existing && path == entry
	}

	result = isEntryCurrent()
	for result && !entry.IsStart() {
		entry = entry.Previous()
		result = isEntryCurrent()
	}

	return result
}
