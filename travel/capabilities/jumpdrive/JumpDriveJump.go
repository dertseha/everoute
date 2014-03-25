package jumpdrive

import (
	"github.com/dertseha/everoute/travel/rules/jumpdistance"
	"github.com/dertseha/everoute/universe"
)

type jumpDriveJump struct {
	toSystemId universe.Id
	distance   float64
}

func Jump(toSytemId universe.Id, distance float64) universe.Jump {
	return &jumpDriveJump{toSystemId: toSytemId, distance: distance}
}

func (jump *jumpDriveJump) Type() string {
	return JumpType
}

func (jump *jumpDriveJump) SourceLocation() universe.Location {
	return universe.AnyLocation()
}

func (jump *jumpDriveJump) DestinationId() universe.Id {
	return jump.toSystemId
}

func (jump *jumpDriveJump) DestinationLocation() universe.Location {
	return universe.AnyLocation()
}

func (jump *jumpDriveJump) Costs() *universe.TravelCostSum {
	return universe.SingleTravelCostSum(jumpdistance.Cost(jump.distance))
}
