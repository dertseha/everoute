package universe

type TravelCost interface {
	Type() string
	Value() float64
	Join(other TravelCost) TravelCost
}
