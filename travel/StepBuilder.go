package travel

import "github.com/dertseha/everoute/universe"

// StepBuilder is a mutable entity used to create new step instances.
type StepBuilder struct {
	solarSystemId universe.Id
	to            universe.Location
	enterCosts    *universe.TravelCostSum
	continueCosts *universe.TravelCostSum
}

// NewStepBuilder returns a new StepBuilder instance with defaults for the parameters.
func NewStepBuilder(solarSystemId universe.Id) *StepBuilder {
	var builder = &StepBuilder{
		solarSystemId: solarSystemId,
		to:            universe.AnyLocation(),
		enterCosts:    universe.EmptyTravelCostSum(),
		continueCosts: universe.EmptyTravelCostSum()}

	return builder
}

// Build returns a new Step instance based on the current parameters.
func (builder *StepBuilder) Build() *Step {
	return NewStep(builder.solarSystemId, builder.to, builder.enterCosts, builder.continueCosts)
}

// To sets the destination location. Default is AnyLocation. The returned instance can be used as a fluent interface.
func (builder *StepBuilder) To(location universe.Location) *StepBuilder {
	builder.to = location

	return builder
}

// WithEnterCosts sets the costs necessary to perform the step. Default is an empty cost. The returned instance
// can be used as a fluent interface.
func (builder *StepBuilder) WithEnterCosts(costs *universe.TravelCostSum) *StepBuilder {
	builder.enterCosts = costs

	return builder
}

// WithContinueCosts sets the costs necessary to continue a path. Default is an empty cost. The returned instance
// can be used as a fluent interface.
func (builder *StepBuilder) WithContinueCosts(costs *universe.TravelCostSum) *StepBuilder {
	builder.continueCosts = costs

	return builder
}
