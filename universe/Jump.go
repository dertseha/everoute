package universe

// Jump describes a connection between two solar systems. A jump originates from a specific solar system.
type Jump interface {
	// Type returns the classification of the jump. Jump types are tightly bound to travel capabilities.
	Type() string
	// SourceLocation specifies the location from where within the originating solar system this jump must be performed.
	SourceLocation() Location
	// DestinationId identifies the destination solar system where this jump ends.
	DestinationId() Id
	// DestinationLocation specifies the location in the destination solar system where the jump lands.
	DestinationLocation() Location
	// Costs returns the sum of costs this jump requires for performing.
	Costs() *TravelCostSum
}

// dataJump is a basic implementation of the Jump interface.
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
