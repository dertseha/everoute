package universe

type Jump interface {
	Type() string
	SourceLocation() Location
	DestinationId() Id
	DestinationLocation() Location
	Costs() *TravelCostSum
}

type dataJump struct {
	jumpType     string
	fromLocation Location
	toSystemId   Id
	toLocation   Location
	costs        *TravelCostSum
}

func (jump *dataJump) Type() string {
	return jump.jumpType
}

func (jump *dataJump) SourceLocation() Location {
	return jump.fromLocation
}

func (jump *dataJump) DestinationId() Id {
	return jump.toSystemId
}

func (jump *dataJump) DestinationLocation() Location {
	return jump.toLocation
}

func (jump *dataJump) Costs() *TravelCostSum {
	return jump.costs
}
