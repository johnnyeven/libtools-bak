package gearman

import (
	"github.com/mikespook/gearman-go/worker"
	"github.com/mikespook/gearman-go/client"
	"github.com/johnnyeven/libtools/task/constants"
	"fmt"
	"github.com/sirupsen/logrus"
)

type GearmanBroker struct {
	worker          *worker.Worker
	workerProcessor constants.TaskProcessor
	client          *client.Client
}

func NewGearmanBroker(info constants.ConnectionInfo) *GearmanBroker {
	w := worker.New(worker.Unlimited)
	w.AddServer(info.Protocol, fmt.Sprintf("%s:%d", info.Host, info.Port))
	return &GearmanBroker{
		worker: w,
	}
}

func (b *GearmanBroker) RegisterChannel(channel string, processor constants.TaskProcessor) error {
	b.workerProcessor = processor
	return b.worker.AddFunc(channel, b.processorJob, worker.Unlimited)
}

func (b *GearmanBroker) processorJob(job worker.Job) ([]byte, error) {
	logrus.Infof("receive job: %+v", job)
	t := &constants.Task{}
	err := constants.UnmarshalData(job.Data(), t)
	if err != nil {
		return nil, err
	}
	ret, err := b.workerProcessor(t)
	if err != nil {
		return nil, err
	}
	return constants.MarshalData(ret)
}

func (b *GearmanBroker) Work() {
	logrus.Info("GearmanBroker.Working...")
	b.worker.Work()
}

func (b *GearmanBroker) Stop() {
	b.worker.Close()
	logrus.Info("GearmanBroker.Stopped")
}
