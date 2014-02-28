package travel

import "github.com/dertseha/everoute/universe"

type TravelCostSum struct {
	costs map[string]universe.TravelCost
}

func NewTravelCostSum(initCosts []universe.TravelCost) *TravelCostSum {
	var sum = &TravelCostSum{
		costs: make(map[string]universe.TravelCost)}

	for _, cost := range initCosts {
		costType := cost.Type()
		sumCost, existing := sum.costs[cost.Type()]

		if existing {
			sum.costs[costType] = sumCost.Join(cost)
		} else {
			sum.costs[costType] = cost
		}
	}

	return sum
}

func (sum *TravelCostSum) Add(cost universe.TravelCost) *TravelCostSum {
	return NewTravelCostSum(append(sum.Total(), cost))
}

func (sum *TravelCostSum) Total() []universe.TravelCost {
	var result = make([]universe.TravelCost, 0, len(sum.costs)+1)

	for _, cost := range sum.costs {
		result = append(result, cost)
	}

	return result
}

func (sum *TravelCostSum) Cost(nullCost universe.TravelCost) universe.TravelCost {
	var result = nullCost
	var cost, existing = sum.costs[nullCost.Type()]

	if existing {
		result = cost
	}

	return result
}
