package universe

// JumpBuilder is used to create a new Jump. A new jump builder has suitable defaults for a generated jump which
// can be overridden by called the corresponding setter.
type JumpBuilder struct {
	jumpType     string
	fromLocation Location
	toSystemId   Id
	toLocation   Location
	costs        *TravelCostSum
}

func newJumpBuilder(jumpType string, destinationId Id) *JumpBuilder {
	result := &JumpBuilder{
		jumpType:     jumpType,
		toSystemId:   destinationId,
		fromLocation: AnyLocation(),
		toLocation:   AnyLocation(),
		costs:        EmptyTravelCostSum()}

	return result
}

// Build creates a Jump instance based on the current values. This method can be called more than once, each time
// creating a new Jump instance representing the current builder values.
func (builder *JumpBuilder) Build() Jump {
	result := &dataJump{
		jumpType:     builder.jumpType,
		fromLocation: builder.fromLocation,
		toSystemId:   builder.toSystemId,
		toLocation:   builder.toLocation,
		costs:        builder.costs}

	return result
}

// From sets the originating location in the source system. Defaults to AnyLocation. The returned builder instance
// can be used as a fluent interface.
func (builder *JumpBuilder) From(location Location) *JumpBuilder {
	builder.fromLocation = location

	return builder
}

// To sets the destination location in the destination system. Defaults to AnyLocation. The returned builder instance
// can be used as a fluent interface.
func (builder *JumpBuilder) To(location Location) *JumpBuilder {
	builder.toLocation = location

	return builder
}

// AddCost adds the provided cost to the current sum of costs. The builder starts with empty costs. The returned
// builder instance can be used as a fluent interface.
func (builder *JumpBuilder) AddCost(cost TravelCost) *JumpBuilder {
	builder.costs = builder.costs.Add(SingleTravelCostSum(cost))

	return builder
}
