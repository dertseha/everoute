package universe

type TravelCostSum struct {
	costs map[string]TravelCost
}

func newTravelCostSum() *TravelCostSum {
	return &TravelCostSum{costs: make(map[string]TravelCost)}
}

var emptyTravelCostSum = newTravelCostSum()

func EmptyTravelCostSum() *TravelCostSum {
	return emptyTravelCostSum
}

func NewTravelCostSum(initCosts ...TravelCost) *TravelCostSum {
	var sum = newTravelCostSum()

	for _, cost := range initCosts {
		costType := cost.Type()
		sumCost, existing := sum.costs[costType]

		if existing {
			sum.costs[costType] = sumCost.Join(cost)
		} else {
			sum.costs[costType] = cost
		}
	}

	return sum
}

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

func (sum *TravelCostSum) Cost(nullCost TravelCost) TravelCost {
	var result = nullCost
	var cost, existing = sum.costs[nullCost.Type()]

	if existing {
		result = cost
	}

	return result
}
