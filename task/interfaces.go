package task

import "github.com/johnnyeven/libtools/task/constants"

type Broker interface {
	RegisterChannel(channel string, processor constants.TaskProcessor) error
	Work()
	Stop()
}

type Client interface {
	SendTask(task *constants.Task) error
	Stop()
}

type CronDescriber interface {
	CronSpec() string
}
