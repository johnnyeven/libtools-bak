package task

import "github.com/johnnyeven/libtools/task/constants"

type Broker interface {
	RegisterChannel(channel string, processor constants.TaskProcessor) error
	Work()
	Stop()
}
