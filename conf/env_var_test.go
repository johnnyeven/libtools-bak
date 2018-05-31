package conf

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"golib/tools/conf/presets"
)

type SubConfig struct {
	Key  string `conf:"env"`
	Bool bool
	Func func() error
}

func (s SubConfig) DockerDefaults() DockerDefaults {
	return DockerDefaults{
		"Key":  "test",
		"Bool": false,
	}
}

func TestEnvVar(t *testing.T) {
	tt := assert.New(t)

	c := struct {
		Array    []string
		Password presets.Password `conf:"env"`
		SubConfig
	}{}

	c.Password = "123456"
	c.Key = "123456"
	c.Array = []string{"1", "2"}

	rv := reflect.Indirect(reflect.ValueOf(c))

	envVars, err := CollectEnvVars(rv, "T")
	tt.Nil(err)

	envVars.Print()
}
