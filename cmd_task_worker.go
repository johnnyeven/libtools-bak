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
		BrokerType: constants.BROKER_TYPE__GEARMAN,
	}
	agent.RegisterRoutes(test_routers.Router.Routes()...)
	agent.Start("service-test.dev", 0)
	select {}
}
