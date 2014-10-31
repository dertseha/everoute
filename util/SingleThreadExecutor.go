package util

// singleThreadExecutor runs provided tasks within one dedicated go-routine.
type singleThreadExecutor struct {
	tasks chan func()
}

// SingleThreadExecutor creates an executor with the provided queueSize.
// The executor performs the requested tasks sequentially through a dedicated go-routine.
// The Execute() implementation will block if queueSize tasks are already pending.
func SingleThreadExecutor(queueSize int) Executor {
	executor := &singleThreadExecutor{
		tasks: make(chan func(), queueSize)}

	run := func() {
		for {
			select {
			case task := <-executor.tasks:
				task()
			}
		}
	}

	go run()

	return executor
}

// Execute is the Executor interface implementation.
func (executor *singleThreadExecutor) Execute(task func()) {
	executor.tasks <- task
}
