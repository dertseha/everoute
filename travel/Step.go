package travel

import (
	"fmt"
	"github.com/dertseha/everoute/universe"
)

// Step represents one entry in a path.
type Step struct {
	solarSystemId universe.Id
	location      universe.Location
	enterCosts    *universe.TravelCostSum
	continueCosts *universe.TravelCostSum
	key           string
}

// NewStep creates a new Step instance with the provided information
func NewStep(solarSystemId universe.Id, location universe.Location, enterCosts *universe.TravelCostSum, continueCosts *universe.TravelCostSum) *Step {
	step := &Step{
		solarSystemId: solarSystemId,
		location:      location,
		enterCosts:    enterCosts,
		continueCosts: continueCosts,
		key:           fmt.Sprintf("%d@%s", solarSystemId, location)}

	return step
}

// AsFirstStep creates a clone of this step that has no costs.
func (step *Step) AsFirstStep() *Step {
	return NewStep(step.solarSystemId, step.location, universe.EmptyTravelCostSum(), universe.EmptyTravelCostSum())
}

// Key is a value uniquely identifying the step. It is based on the solar system ID and the location.
func (step *Step) Key() string {
	return step.key
}

// SolarSystemId identifies the solar system to which this step was made.
func (step *Step) SolarSystemId() universe.Id {
	return step.solarSystemId
}

// Location describes the location within the solar system where this step ends.
func (step *Step) Location() universe.Location {
	return step.location
}

// EnterCosts describes the costs that were necessary to perform this step.
func (step *Step) EnterCosts() *universe.TravelCostSum {
	return step.enterCosts
}

// ContinueCosts describes the costs that will be necessary when a path is continued with this step.
func (step *Step) ContinueCosts() *universe.TravelCostSum {
	return step.continueCosts
}
