package universe

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

func (builder *JumpBuilder) Build() Jump {
	result := &dataJump{
		jumpType:     builder.jumpType,
		fromLocation: builder.fromLocation,
		toSystemId:   builder.toSystemId,
		toLocation:   builder.toLocation,
		costs:        builder.costs}

	return result
}

func (builder *JumpBuilder) From(location Location) *JumpBuilder {
	builder.fromLocation = location

	return builder
}

func (builder *JumpBuilder) To(location Location) *JumpBuilder {
	builder.toLocation = location

	return builder
}

func (builder *JumpBuilder) AddCost(cost TravelCost) *JumpBuilder {
	builder.costs = builder.costs.Add(NewTravelCostSum(cost))

	return builder
}
