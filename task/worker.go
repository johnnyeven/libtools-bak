package task

import (
	"github.com/johnnyeven/libtools/task/constants"
	"github.com/johnnyeven/libtools/task/gearman"
)

type Worker struct {
	broker    Consumer
	processor constants.TaskProcessor
}

func NewWorker(brokerType constants.BrokerType, connInfo constants.ConnectionInfo) *Worker {
	var b Consumer
	if brokerType == constants.BROKER_TYPE__GEARMAN {
		b = gearman.NewGearmanBroker(connInfo)
	}
	return &Worker{
		broker: b,
	}
}

func (mq *Worker) Stop() {
	if mq.broker == nil {
		return
	}
	mq.broker.Stop()
}

func (mq *Worker) Start(channel string, processor constants.TaskProcessor) {
	if mq.broker == nil {
		return
	}
	mq.processor = processor
	mq.broker.RegisterChannel(channel, processor)
	go mq.broker.Work()
}
