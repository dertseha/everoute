package util

// Executor represents any strategy for running a task (a function).
// A typical executor would defer the task execution into another thread.
type Executor interface {
	// Execute registers the provided task at the executor. Depending on the Executor's implementation, this task will
	// either be executed immediately (blocking) or concurrently.
	Execute(task func())
}
