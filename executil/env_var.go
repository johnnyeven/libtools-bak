package executil

import (
	"os"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

var reEnvVar = regexp.MustCompile("(\\$?\\$)\\{?([A-Z0-9_]+)\\}?")

type EnvVars map[string]string

func (envVars EnvVars) LoadFromEnv() {
	for _, keyPair := range os.Environ() {
		keyValue := strings.Split(keyPair, "=")
		envVars[keyValue[0]] = keyValue[1]
	}
}

func ParseByEnv(s string) string {
	envVars := EnvVars{}

	envVars.LoadFromEnv()

	result := reEnvVar.ReplaceAllStringFunc(s, func(str string) string {
		matched := reEnvVar.FindAllStringSubmatch(str, -1)[0]

		// skip $${ }
		if matched[1] == "$$" {
			return "${" + matched[2] + "}"
		}

		if value, ok := envVars[matched[2]]; ok {
			return value
		} else {
			logrus.Errorf("Missing environment variable ${%s}", matched[2])
		}

		return ""
	})

	return result
}
