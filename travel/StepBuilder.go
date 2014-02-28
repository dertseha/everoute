package travel

import "github.com/dertseha/everoute/universe"

type StepBuilder struct {
	solarSystemId universe.Id
	to            universe.Location
	enterCosts    []universe.TravelCost
	continueCosts []universe.TravelCost
}

func NewStepBuilder(solarSystemId universe.Id) *StepBuilder {
	var builder = &StepBuilder{solarSystemId: solarSystemId,
		to:            universe.AnyLocation(),
		enterCosts:    nil,
		continueCosts: nil}

	return builder
}

func (builder *StepBuilder) Build() *Step {
	return NewStep(builder.solarSystemId, builder.to, builder.enterCosts, builder.continueCosts)
}

func (builder *StepBuilder) To(location universe.Location) *StepBuilder {
	builder.to = location

	return builder
}

func (builder *StepBuilder) WithEnterCosts(costs []universe.TravelCost) *StepBuilder {
	builder.enterCosts = append(make([]universe.TravelCost, 0), costs...)

	return builder
}

func (builder *StepBuilder) WithContinueCosts(costs []universe.TravelCost) *StepBuilder {
	builder.continueCosts = append(make([]universe.TravelCost, 0), costs...)

	return builder
}
