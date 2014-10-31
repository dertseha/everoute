package travel

import "github.com/dertseha/everoute/universe"

type addingTravelCost struct {
	costType string
	value    float64
}

// AddingTravelCost returns a TravelCost that can simply add its value to others of the same type.
func AddingTravelCost(costType string, value float64) universe.TravelCost {
	cost := &addingTravelCost{
		costType: costType,
		value:    value}

	return cost
}

// Type returns the type identification passed at construction.
func (cost *addingTravelCost) Type() string {
	return cost.costType
}

// Value returns the cost value passed at construction.
func (cost *addingTravelCost) Value() float64 {
	return cost.value
}

// Join returns a travel cost of the same type that is the sum of this and the other value.
// This method panics if the type of the given cost is not of the same as from this cost.
func (cost *addingTravelCost) Join(other universe.TravelCost) universe.TravelCost {
	if cost.costType != other.Type() {
		panic("Cost type mismatch")
	}

	return AddingTravelCost(cost.costType, cost.value+other.Value())
}
