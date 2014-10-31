package util

import (
	"time"

	check "gopkg.in/check.v1"
)

type SingleThreadExecutorTestSuite struct {
	executor Executor
}

var _ = check.Suite(&SingleThreadExecutorTestSuite{})

func (suite *SingleThreadExecutorTestSuite) SetUpTest(c *check.C) {
	suite.executor = SingleThreadExecutor(1)
}

func (suite *SingleThreadExecutorTestSuite) TestExecuteWillRunPassedTask(c *check.C) {
	taskDone := make(chan bool, 1)
	timeout := time.After(5 * time.Second)
	task := func() { taskDone <- true }
	executed := false

	suite.executor.Execute(task)

	select {
	case <-taskDone:
		executed = true
	case <-timeout:
	}

	c.Assert(executed, check.Equals, true)
}
