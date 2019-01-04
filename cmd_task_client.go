package main

import (
	"github.com/mikespook/gearman-go/client"
	"github.com/sirupsen/logrus"
)

func main() {
	c, err := client.New("tcp", "www.profzone.net:4730")
	if err != nil {
		logrus.Panic(err)
	}
	defer c.Close()

	c.Do("service-test.dev", []byte{}, client.JobNormal, func(response *client.Response) {

	})
}
