package travel

import "github.com/dertseha/everoute/universe"

type addingTravelCost struct {
	costType string
	value    float64
}

func AddingTravelCost(costType string, value float64) universe.TravelCost {
	cost := &addingTravelCost{
		costType: costType,
		value:    value}

	return cost
}

func (cost *addingTravelCost) Type() string {
	return cost.costType
}

func (cost *addingTravelCost) Value() float64 {
	return cost.value
}

func (cost *addingTravelCost) Join(other universe.TravelCost) universe.TravelCost {
	if cost.costType != other.Type() {
		panic("Cost type mismatch")
	}

	return AddingTravelCost(cost.costType, cost.value+other.Value())
}
