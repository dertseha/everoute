package travel

import "github.com/dertseha/everoute/universe"

// Path represents a single-linked list of steps. A path instance is always the end of the path so far.
type Path interface {
	// DestinationKey returns a value that uniquely represents the end of this path.
	DestinationKey() string
	// CostSum returns the total cost of the path so far. It is the total of all contained steps.
	CostSum() *universe.TravelCostSum
	// IsStart returns true for the start of a path.
	IsStart() bool
	// Previous returns the path instance of the previous step. This method panics if the path instance is the start.
	Previous() Path
	// Extend returns a new path instance based on this, combined with a new step.
	Extend(step *Step) Path
	// Step returns the step instance representing this path extension.
	Step() *Step
	// Steps returns a slice of steps of this path from its start to this instance.
	Steps() []*Step
}

type chainedPath struct {
	step     *Step
	costSum  *universe.TravelCostSum
	previous Path
}

type startPath struct {
	step    *Step
	costSum *universe.TravelCostSum
}

func extendPath(path Path, step *Step) Path {
	var result = &chainedPath{
		step:     step,
		previous: path,
		costSum:  path.CostSum().Add(path.Step().ContinueCosts()).Add(step.EnterCosts())}

	return result
}

func (path *chainedPath) DestinationKey() string {
	return path.step.Key()
}

func (path *chainedPath) CostSum() *universe.TravelCostSum {
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

func (path *startPath) CostSum() *universe.TravelCostSum {
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

// NewPath returns a new path that starts with the provided step.
func NewPath(step *Step) Path {
	var path = &startPath{
		step:    step,
		costSum: step.EnterCosts()}

	return path
}
