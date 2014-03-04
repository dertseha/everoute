package universe

type Jump struct {
	jumpType     string
	fromLocation Location
	toSystemId   Id
	toLocation   Location
	costs        *TravelCostSum
}

func (jump *Jump) Type() string {
	return jump.jumpType
}

func (jump *Jump) SourceLocation() Location {
	return jump.fromLocation
}

func (jump *Jump) DestinationId() Id {
	return jump.toSystemId
}

func (jump *Jump) DestinationLocation() Location {
	return jump.toLocation
}

func (jump *Jump) Costs() *TravelCostSum {
	return jump.costs
}
