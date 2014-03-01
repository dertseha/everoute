package travel

type Path interface {
	DestinationKey() string
	CostSum() *TravelCostSum
	IsStart() bool
	Previous() Path
	Extend(step *Step) Path
	Step() *Step
	Steps() []*Step
}

type chainedPath struct {
	step     *Step
	costSum  *TravelCostSum
	previous Path
}

type startPath struct {
	step    *Step
	costSum *TravelCostSum
}

func extendPath(path Path, step *Step) Path {
	var costs = append(path.Step().ContinueCosts(), step.EnterCosts()...)
	var result = &chainedPath{
		step:     step,
		previous: path,
		costSum:  NewTravelCostSum(costs)}

	return result
}

func (path *chainedPath) DestinationKey() string {
	return path.step.Key()
}

func (path *chainedPath) CostSum() *TravelCostSum {
	return path.costSum
}

func (path *chainedPath) IsStart() bool {
	return false
}

func (path *chainedPath) Previous() Path {
	return path.previous
}

func (path *chainedPath) Extend(step *Step) Path {
	return extendPath(path, step)
}

func (path *chainedPath) Step() *Step {
	return path.step
}

func (path *chainedPath) Steps() []*Step {
	var paths = make([]Path, 0)
	var temp Path = path

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

func (path *startPath) DestinationKey() string {
	return path.step.Key()
}

func (path *startPath) CostSum() *TravelCostSum {
	return path.costSum
}

func (path *startPath) IsStart() bool {
	return true
}

func (path *startPath) Previous() Path {
	panic("Start of Path has no predecessor")
}

func (path *startPath) Extend(step *Step) Path {
	return extendPath(path, step)
}

func (path *startPath) Step() *Step {
	return path.step
}

func (path *startPath) Steps() []*Step {
	var result = make([]*Step, 1)

	result[0] = path.step
	return result
}

func NewPath(step *Step) Path {
	var path = &startPath{
		step:    step,
		costSum: NewTravelCostSum(step.EnterCosts())}

	return path
}
