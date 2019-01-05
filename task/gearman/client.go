package gearman

import (
	"github.com/mikespook/gearman-go/client"
	"github.com/johnnyeven/libtools/task/constants"
	"fmt"
	"github.com/sirupsen/logrus"
)

type GearmanClient struct {
	client *client.Client
}

func NewGearmanClient(info constants.ConnectionInfo) *GearmanClient {
	c, err := client.New(info.Protocol, fmt.Sprintf("%s:%d", info.Host, info.Port))
	if err != nil {
		logrus.Panicf("NewGearmanClient err: %v", err)
	}

	return &GearmanClient{
		client: c,
	}
}

func (c *GearmanClient) SendTask(task *constants.Task) error {
	data, err := constants.MarshalData(task)
	if err != nil {
		logrus.Errorf("GearmanClient.SendTask err: %v", err)
		return err
	}
	_, err = c.client.Do(task.Channel, data, client.JobNormal, func(response *client.Response) {

	})
	if err != nil {
		logrus.Errorf("GearmanClient.Do err: %v", err)
		return err
	}

	return nil
}

func (c *GearmanClient) Stop() {
	c.client.Close()
}
