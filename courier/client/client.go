package client

import (
	"fmt"
	"os"
	"strings"
	"time"

	"golib/tools/conf"
	"golib/tools/courier"
	"golib/tools/courier/httpx"
	"golib/tools/courier/transport_grpc"
	"golib/tools/courier/transport_http"
	"golib/tools/env"
	"golib/tools/log/context"
)

type Client struct {
	Name string
	// used in service
	Service       string
	Version       string
	Host          string `conf:"env,upstream"`
	Mode          string
	Port          int16
	Timeout       time.Duration
	WrapTransport transport_http.TransportWrapper
}

func (c Client) DockerDefaults() conf.DockerDefaults {
	return conf.DockerDefaults{
		// todo make switch in docker or expose
		"Host": conf.RancherInternal("service-"+c.Name, "service-"+c.Name),
		"Port": 80,
	}
}

func (Client) MarshalDefaults(v interface{}) {
	if client, ok := v.(*Client); ok {
		if client.Service == "" {
			client.Service = os.Getenv("PROJECT_NAME")
		}
		if client.Version == "" {
			client.Version = os.Getenv("PROJECT_REF")
		}
		if client.Mode == "" {
			client.Mode = "http"
		}
		if client.Host == "" {
			client.Host = fmt.Sprintf("service-%s.staging.g7pay.net", client.Name)
		}
		if client.Port == 0 {
			client.Port = 80
		}
		if client.Timeout == 0 {
			client.Timeout = 5 * time.Second
		}
	}
}

func (c Client) GetBaseURL(protocol string) (url string) {
	url = c.Host
	if protocol != "" {
		url = fmt.Sprintf("%s://%s", protocol, c.Host)
	}
	if c.Port > 0 {
		url = fmt.Sprintf("%s:%d", url, c.Port)
	}
	return
}

func (c *Client) Request(id, httpMethod, uri string, req interface{}, metas ...courier.Metadata) IRequest {
	requestID := context.GetLogID()
	if !env.IsOnline() {
		if requestIDInMeta := courier.MetadataMerge(metas...).Get(httpx.HeaderRequestID); requestIDInMeta != "" {
			requestID = requestIDInMeta
		}
		mocker, err := ParseMockID(c.Service, requestID)
		if err == nil {
			if m, exists := mocker.Mocks[id]; exists {
				return &MockRequest{
					MockData: m,
				}
			}
		}
	}

	switch strings.ToLower(c.Mode) {
	case "grpc":
		serverName, method := parseID(id)
		return &transport_grpc.GRPCRequest{
			BaseURL:    c.GetBaseURL(""),
			ServerName: serverName,
			Method:     method,
			Timeout:    c.Timeout,
			RequestID:  requestID,
			Req:        req,
		}
	default:
		return &transport_http.HttpRequest{
			UserAgent:     c.Service + " " + c.Version,
			BaseURL:       c.GetBaseURL(c.Mode),
			Method:        httpMethod,
			URI:           uri,
			ID:            id,
			Timeout:       c.Timeout,
			RequestID:     requestID,
			WrapTransport: c.WrapTransport,
			Req:           req,
		}
	}
}

func parseID(id string) (serverName string, method string) {
	values := strings.Split(id, ".")
	if len(values) == 2 {
		serverName = strings.ToLower(strings.Replace(values[0], "Client", "", -1))
		method = values[1]
	}
	return
}
