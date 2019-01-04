package main

import (
	"github.com/johnnyeven/libtools/task"
	"github.com/johnnyeven/libtools/task/constants"
)

func main() {
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

	agent.Start("service-test.dev", 0)
	select {}
}
