package task

import (
	"github.com/johnnyeven/libtools/conf"
	"os"
	"strings"
	"fmt"
	"github.com/johnnyeven/libtools/env"
	"github.com/johnnyeven/libtools/task/constants"
	"sync"
	"github.com/johnnyeven/libtools/courier"
	"reflect"
	"context"
)

type Agent struct {
	ConnectionInfo constants.ConnectionInfo
	Channel        string
	BrokerType     constants.BrokerType `conf:"env"`
	worker         *Worker
	backend        *Backend
	jobs           sync.Map
}

func (Agent) MarshalDefaults(v interface{}) {
	if agent, ok := v.(*Agent); ok {
		if agent.BrokerType == constants.BROKER_TYPE_UNKNOWN {
			agent.BrokerType = constants.BROKER_TYPE__GEARMAN
		}
		if agent.Channel == "" {
			agent.Channel = getDefaultChannel()
		}
		if agent.ConnectionInfo.Port == 0 {
			agent.ConnectionInfo.Port = 4730
		}
		if agent.ConnectionInfo.Protocol == "" {
			agent.ConnectionInfo.Protocol = "tcp"
		}
	}
}

func (a *Agent) DockerDefaults() conf.DockerDefaults {
	return conf.DockerDefaults{
		"Protocol":   "tcp",
		"Host":       conf.RancherInternal("dep-tools", "gearmand"),
		"Port":       4730,
		"BrokerType": 1,
	}
}

func getDefaultChannel() string {
	serviceName := "anonymous"
	if projectName, exists := os.LookupEnv("PROJECT_NAME"); exists {
		serviceName = projectName
	}
	return strings.ToLower(fmt.Sprintf("%s.%s", serviceName, env.GetRuntimeEnv()))
}

func (a *Agent) Register(subject string, processor constants.TaskProcessor) {
	a.jobs.Store(subject, processor)
}

func (a *Agent) RegisterRoutes(routes ...*courier.Route) {
	for _, route := range routes {
		if len(route.Operators) == 0 {
			continue
		}

		operatorMetas := courier.ToOperatorMetaList(route.Operators...)

		lastOpIndex := len(operatorMetas) - 1
		lastOpMeta := operatorMetas[lastOpIndex]
		subject := lastOpMeta.Type.Name()

		a.Register(subject, func(task *constants.Task) (interface{}, error) {
			ctx := context.Background()

			for i, opMeta := range operatorMetas {
				op := reflect.New(opMeta.Type).Interface().(courier.IOperator)

				if err := constants.UnmarshalData(task.Data, op); err != nil {
					return nil, err
				}

				ret, err := op.Output(ctx)
				if err != nil {
					return nil, err
				}

				if i != lastOpIndex {
					if ctxProvider, ok := op.(courier.IContextProvider); ok {
						ctx = context.WithValue(ctx, ctxProvider.ContextKey(), ret)
					}
					continue
				}

				return ret, nil
			}

			return nil, nil
		})
	}
}

func (a *Agent) Start(channel string, numWorker int) {
	a.Channel = channel

	if a.backend == nil {
		a.backend = NewBackend()
	}
	if a.worker == nil {
		a.worker = NewWorker(a.BrokerType, a.ConnectionInfo)
	}
	a.worker.Start(a.Channel, a.workerProcessor)
}

func (a *Agent) Stop() {
	if a.worker != nil {
		a.worker.Stop()
	}
}

func (a *Agent) workerProcessor(task *constants.Task) (interface{}, error) {
	a.backend.Feedback(task.Processing())
	subject := task.Subject
	p, ok := a.jobs.Load(subject)
	if !ok {
		err := fmt.Errorf("subject %s not registered", subject)
		a.backend.Feedback(task.Fail(err))
		return nil, err
	}

	ret, err := p.(constants.TaskProcessor)(task)
	if err != nil {
		a.backend.Feedback(task.Fail(err))
		return nil, err
	}
	a.backend.Feedback(task.Success(ret))
	return nil, nil
}
