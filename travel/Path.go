package travel

type Path struct {
	step     *Step
	previous *Path
	costSum  *TravelCostSum
}

func NewPath(step *Step) *Path {
	var path = &Path{
		step:     step,
		previous: nil,
		costSum:  NewTravelCostSum(step.EnterCosts())}

	return path
}

func (path *Path) DestinationKey() string {
	return path.step.Key()
}

func (path *Path) IsStart() bool {
	return path.previous == nil
}

func (path *Path) Previous() *Path {
	if path.IsStart() {
		panic("Start of Path has no predecessor")
	}

	return path.previous
}

func (path *Path) Extend(step *Step) *Path {
	var costs = append(path.step.ContinueCosts(), step.EnterCosts()...)
	var result = &Path{
		step:     step,
		previous: path,
		costSum:  NewTravelCostSum(costs)}

	return result
}

func (path *Path) Step() *Step {
	return path.step
}

func (path *Path) CostSum() *TravelCostSum {
	return path.costSum
}
