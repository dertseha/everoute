package util

type Executor interface {
	Execute(task func())
}
