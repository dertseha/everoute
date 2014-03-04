package travel

import (
	"fmt"
	"github.com/dertseha/everoute/universe"
)

type Step struct {
	solarSystemId universe.Id
	location      universe.Location
	enterCosts    *universe.TravelCostSum
	continueCosts *universe.TravelCostSum
	key           string
}

func NewStep(solarSystemId universe.Id, location universe.Location, enterCosts *universe.TravelCostSum, continueCosts *universe.TravelCostSum) *Step {
	step := &Step{
		solarSystemId: solarSystemId,
		location:      location,
		enterCosts:    enterCosts,
		continueCosts: continueCosts,
		key:           fmt.Sprintf("%d@%s", solarSystemId, location)}

	return step
}

func (step *Step) AsFirstStep() *Step {
	return NewStep(step.solarSystemId, step.location, nil, nil)
}

func (step *Step) Key() string {
	return step.key
}

func (step *Step) SolarSystemId() universe.Id {
	return step.solarSystemId
}

func (step *Step) Location() universe.Location {
	return step.location
}

func (step *Step) EnterCosts() *universe.TravelCostSum {
	return step.enterCosts
}

func (step *Step) ContinueCosts() *universe.TravelCostSum {
	return step.continueCosts
}
