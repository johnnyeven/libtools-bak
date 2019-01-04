package main

import (
	"github.com/mikespook/gearman-go/client"
	"github.com/sirupsen/logrus"
	"github.com/johnnyeven/libtools/task/constants"
	"time"
	"github.com/google/uuid"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	c, err := client.New("tcp", "www.profzone.net:4730")
	if err != nil {
		logrus.Panic(err)
	}
	defer c.Close()

	t := constants.Task{
		ID:         uuid.New().String(),
		Channel:    "service-test.dev",
		Subject:    "FindUser",
		Data:       nil,
		CreateTime: time.Now(),
	}
	args, _ := constants.MarshalData(t)
	handle, err := c.Do("service-test.dev", args, client.JobNormal, func(response *client.Response) {
		logrus.Info(string(response.Data))
	})
	if err != nil {
		logrus.Panic(err)
	}

	status, err := c.Status(handle)
	if err != nil {
		logrus.Panic(err)
	}

	logrus.Infof("%+v", status)
	select{}
}
