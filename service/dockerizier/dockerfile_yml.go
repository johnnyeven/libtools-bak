package dockerizier

import (
	"gopkg.in/yaml.v2"

	"golib/tools/conf"
	"golib/tools/docker"
)

var (
	Image     = "${G7PAY_DOCKER_REGISTRY}/${PROJECT_GROUP}/${PROJECT_NAME}:${PROJECT_VERSION}"
	FromImage = "${G7PAY_DOCKER_REGISTRY}/g7/env-golang:latest"
)

func toDockerFileYML(envVars conf.EnvVars, serviceName string) string {
	d := &docker.Dockerfile{
		From:  FromImage,
		Image: Image,
	}

	d = d.AddEnv("GOENV", "DEV")

	for key, envVar := range envVars {
		strValue := envVar.GetValue(false)
		if strValue == "./swagger.json" {
			d = d.AddContent(strValue, "./")
		}
		if envVar.FallbackValue != nil {
			d = d.AddEnv(key, envVar.GetFallbackValue(false))
		}
	}

	d = d.WithWorkDir("/go/bin")
	d = d.WithCmd("./"+serviceName, "-c", "false")
	d = d.WithExpose("80")

	d = d.AddContent("./config", "./config")
	d = d.AddContent("./"+serviceName, "./")

	bytes, _ := yaml.Marshal(d)
	return string(bytes)
}
