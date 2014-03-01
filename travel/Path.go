package travel

type Path struct {
	step     *Step
	isStart  bool
	previous func() *Path
	costSum  *TravelCostSum
}

func NewPath(step *Step) *Path {
	var path = &Path{
		step:     step,
		isStart:  true,
		previous: func() *Path { panic("Start of Path has no predecessor") },
		costSum:  NewTravelCostSum(step.EnterCosts())}

	return path
}

func (path *Path) DestinationKey() string {
	return path.step.Key()
}

func (path *Path) IsStart() bool {
	return path.isStart
}

func (path *Path) Previous() *Path {
	return path.previous()
	/*
		if path.IsStart() {
			panic("Start of Path has no predecessor")
		}

		return path.previous
	*/
}

func (path *Path) Extend(step *Step) *Path {
	var costs = append(path.step.ContinueCosts(), step.EnterCosts()...)
	var result = &Path{
		step:     step,
		isStart:  false,
		previous: func() *Path { return path },
		costSum:  NewTravelCostSum(costs)}

	return result
}

func (path *Path) Step() *Step {
	return path.step
}

func (path *Path) Steps() []*Step {
	var paths = make([]*Path, 0)
	var temp = path

	for !temp.IsStart() {
		paths = append(paths, temp)
		temp = temp.Previous()
	}
	paths = append(paths, temp)

	pathCount := len(paths)
	result := make([]*Step, pathCount)
	for index, temp := range paths {
		result[pathCount-1-index] = temp.Step()
	}

	return result
}

func (path *Path) CostSum() *TravelCostSum {
	return path.costSum
}
