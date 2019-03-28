package kafka

import (
	"context"
	"fmt"
	"github.com/johnnyeven/libtools/task/constants"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	r               *kafka.Reader
	workerProcessor constants.TaskProcessor
}

func NewKafkaConsumer(info constants.ConnectionInfo) *KafkaConsumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{fmt.Sprintf("%s:%d", info.Host, info.Port)},
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	return &KafkaConsumer{
		r: r,
	}
}

func (c *KafkaConsumer) RegisterChannel(channel string, processor constants.TaskProcessor) error {
	c.r.Config().Topic = channel
	c.workerProcessor = processor

	return nil
}

func (c *KafkaConsumer) Work() {
	logrus.Debug("KafkaConsumer.Working...")
	for {
		m, err := c.r.ReadMessage(context.Background())
		if err != nil {
			logrus.Errorf("kafka.Reader.ReadMessage err: %v", err)
		}

		t := &constants.Task{}
		err = constants.UnmarshalData(m.Value, t)
		if err != nil {
			logrus.Errorf("Work UnmarshalData err: %v", err)
			continue
		}
		ret, err := c.workerProcessor(t)
		if err != nil {
			continue
		}
	}
}

func (c *KafkaConsumer) Stop() {
	c.r.Close()
	logrus.Debug("KafkaConsumer.Stopped")
}
