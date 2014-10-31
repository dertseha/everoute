package universe

// TravelCost represents some cost for travelling.
type TravelCost interface {
	// Type specifies the cost type. Cost types are bound to travel capabilities and rules.
	Type() string
	// Value returns the numerical representation of the cost. Unit and range is type specific.
	Value() float64
	// Join combines this cost with another into a resulting cost.
	Join(other TravelCost) TravelCost
}
