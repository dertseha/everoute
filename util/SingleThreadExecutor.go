package util

type singleThreadExecutor struct {
	tasks chan func()
}

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

func (executor *singleThreadExecutor) Execute(task func()) {
	executor.tasks <- task
}
