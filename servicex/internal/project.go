package internal

import (
	"strings"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"reflect"
	"github.com/johnnyeven/libtools/executil"
)

var ProjectFile = "profzone.yml"

var (
	DOCKER_REGISTRY_KEY = "PROFZONE_DOCKER_REGISTRY"
	DOCKER_REGISTRY     = "registry.profzone.net:5000"
)

type Project struct {
	Name            string            `env:"name" yaml:"name"`
	Group           string            `env:"group" yaml:"group,omitempty"`
	Version         Version           `env:"version" yaml:"version"`
	Desc            string            `env:"description" yaml:"description"`
	ProgramLanguage string            `env:"program_language" yaml:"program_language"`
	Scripts         map[string]Script `yaml:"scripts,omitempty"`
	Feature         string            `yaml:"feature,omitempty"`
}

func (p *Project) UnmarshalFromFile() error {
	bytes, err := ioutil.ReadFile(ProjectFile)
	if err != nil {
		return err
	}
	errForUnmarshal := yaml.Unmarshal(bytes, p)
	if errForUnmarshal != nil {
		return errForUnmarshal
	}
	return nil
}

func (p Project) WithVersion(s string) Project {
	v, err := FromVersionString(s)
	if err != nil {
		panic(err)
	}
	p.Version = *v
	return p
}

func (p Project) WithGroup(group string) Project {
	p.Group = group
	return p
}

func (p Project) WithDesc(desc string) Project {
	p.Desc = desc
	return p
}

func (p Project) WithName(name string) Project {
	p.Name = name
	return p
}

func (p Project) WithLanguage(pl string) Project {
	p.ProgramLanguage = pl
	return p
}

func (p Project) WithFeature(f string) Project {
	p.Feature = f
	return p
}

func (p Project) WithScripts(key string, scripts ...string) Project {
	if p.Scripts == nil {
		p.Scripts = map[string]Script{}
	}
	p.Scripts[key] = append(Script{}, scripts...)
	return p
}

func WrapEnv(s string) string {
	return strings.ToUpper("PROJECT_" + s)
}

func SetEnv(k string, v string) {
	os.Setenv(k, v)
	fmt.Printf("export %s=%s\n", k, v)
}

func (p *Project) SetEnviron() {
	if os.Getenv(DOCKER_REGISTRY_KEY) == "" {
		SetEnv(DOCKER_REGISTRY_KEY, DOCKER_REGISTRY)
	}

	tpe := reflect.TypeOf(p).Elem()
	rv := reflect.Indirect(reflect.ValueOf(p))

	for i := 0; i < tpe.NumField(); i++ {
		field := tpe.Field(i)
		env := field.Tag.Get("env")

		if len(env) > 0 {
			value := rv.FieldByName(field.Name)

			if stringer, ok := value.Interface().(fmt.Stringer); ok {
				v := stringer.String()
				if len(v) > 0 {
					SetEnv(WrapEnv(env), v)
				}
			} else {
				SetEnv(WrapEnv(env), value.String())
			}
		}
	}
}

func (p *Project) Command(args ...string) *exec.Cmd {
	p.SetEnviron()

	sh := "sh"
	if runtime.GOOS == "windows" {
		sh = "bash"
	}

	envVars := executil.EnvVars{}
	envVars.LoadFromEnviron()

	return exec.Command(sh, "-c", envVars.Parse(strings.Join(args, " ")))
}

func (p *Project) Run(commands ...*exec.Cmd) {
	for _, cmd := range commands {
		if cmd != nil {
			executil.StdRun(cmd)
		}
	}
}

func (p *Project) Execute(args ...string) {
	p.Run(p.Command(args...))
}

func (p *Project) WriteToFile() {
	bytes, err := yaml.Marshal(p)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(ProjectFile, bytes, os.ModePerm)
}

type Script []string

func (s Script) IsZero() bool {
	return len(s) == 0
}

func (s Script) String() string {
	return strings.Join(s, " && ")
}

func (s Script) MarshalYAML() (interface{}, error) {
	if len(s) > 1 {
		return s, nil
	}
	return s[0], nil
}

func (s *Script) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	err := unmarshal(&str)
	if err == nil {
		*s = []string{str}
	} else {
		var values []string
		err := unmarshal(&values)
		if err != nil {
			return err
		}
		*s = values
	}
	return nil
}
