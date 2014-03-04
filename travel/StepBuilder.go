package travel

import "github.com/dertseha/everoute/universe"

type StepBuilder struct {
	solarSystemId universe.Id
	to            universe.Location
	enterCosts    *universe.TravelCostSum
	continueCosts *universe.TravelCostSum
}

func NewStepBuilder(solarSystemId universe.Id) *StepBuilder {
	var builder = &StepBuilder{
		solarSystemId: solarSystemId,
		to:            universe.AnyLocation(),
		enterCosts:    universe.EmptyTravelCostSum(),
		continueCosts: universe.EmptyTravelCostSum()}

	return builder
}

func (builder *StepBuilder) Build() *Step {
	return NewStep(builder.solarSystemId, builder.to, builder.enterCosts, builder.continueCosts)
}

func (builder *StepBuilder) To(location universe.Location) *StepBuilder {
	builder.to = location

	return builder
}

func (builder *StepBuilder) WithEnterCosts(costs *universe.TravelCostSum) *StepBuilder {
	builder.enterCosts = costs

	return builder
}

func (builder *StepBuilder) WithContinueCosts(costs *universe.TravelCostSum) *StepBuilder {
	builder.continueCosts = costs

	return builder
}
