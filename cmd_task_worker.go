package main

import (
	"github.com/johnnyeven/libtools/task"
	"github.com/johnnyeven/libtools/task/constants"
	"github.com/sirupsen/logrus"
	"github.com/johnnyeven/libtools/task/test_routers"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	agent := task.Agent{
		ConnectionInfo: constants.ConnectionInfo{
			Protocol: "tcp",
			Host:     "www.profzone.net",
			Port:     4730,
			UserName: "",
			Password: "",
		},
		Channel:    "service-test.dev",
		BrokerType: constants.BROKER_TYPE__GEARMAN,
	}
	agent.MarshalDefaults(&agent)
	agent.Init()
	agent.RegisterRoutes(test_routers.Router.Routes()...)
	agent.Start(5)
	select {}
}
