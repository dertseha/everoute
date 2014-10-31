package universe

// TravelCostSum is a collection of travel costs. One instance is immutable and generates new sums when
// combined with further costs.
type TravelCostSum struct {
	costs map[string]TravelCost
}

func newTravelCostSum() *TravelCostSum {
	return &TravelCostSum{costs: make(map[string]TravelCost)}
}

var emptyTravelCostSum = newTravelCostSum()

// EmptyTravelCostSum returns a sum that is without any costs.
func EmptyTravelCostSum() *TravelCostSum {
	return emptyTravelCostSum
}

// SingleTravelCostSum returns a new cost sum instance that is initialized with the provided single TravelCost.
func SingleTravelCostSum(cost TravelCost) *TravelCostSum {
	sum := newTravelCostSum()

	sum.costs[cost.Type()] = cost

	return sum
}

// Add combines this travel cost sum with another. The resulting sum contains all costs joined together.
func (sum *TravelCostSum) Add(other *TravelCostSum) *TravelCostSum {
	var result = newTravelCostSum()

	for costType, cost := range sum.costs {
		if _, existing := other.costs[costType]; !existing {
			result.costs[costType] = cost
		}
	}
	for costType, cost := range other.costs {
		if thisCost, existing := sum.costs[costType]; existing {
			result.costs[costType] = thisCost.Join(cost)
		} else {
			result.costs[costType] = cost
		}
	}

	return result
}

// Cost returns a single TravelCost instance identified by the provided nullCost. If the sum does not contain the
// specified cost, the nullCost will be returned.
func (sum *TravelCostSum) Cost(nullCost TravelCost) TravelCost {
	var result = nullCost
	var cost, existing = sum.costs[nullCost.Type()]

	if existing {
		result = cost
	}

	return result
}
