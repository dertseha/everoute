package travel

import "github.com/dertseha/everoute/universe"

type AddingTravelCost struct {
	costType string
	value    float64
}

func (cost *AddingTravelCost) Type() string {
	return cost.costType
}

func (cost *AddingTravelCost) Value() float64 {
	return cost.value
}

func (cost *AddingTravelCost) Join(other universe.TravelCost) universe.TravelCost {
	if cost.costType != other.Type() {
		panic("Cost type mismatch")
	}

	result := AddingTravelCost{
		costType: cost.costType,
		value:    cost.value + other.Value()}

	return &result
}
