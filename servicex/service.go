package servicex

import (
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/cobra"

	"github.com/profzone/libtools/conf"
	"github.com/profzone/libtools/service/dockerizier"
)

var (
	EnvVarKeyProjectName    = "PROJECT_NAME"
	EnvVarKeyProjectFeature = "PROJECT_FEATURE"
)

func init() {
	if projectFeature, exists := os.LookupEnv(EnvVarKeyProjectFeature); exists && projectFeature != "" {
		SetServiceName(ServiceName + "--" + projectFeature)
	}

	command.AddCommand(run)

	command.PersistentFlags().
		StringVarP(&envConfigPrefix, "env-config-prefix", "e", "S", "prefix for env var config")
	command.PersistentFlags().
		BoolVarP(&outputDockerConfig, "output-docker-config", "c", true, "output configuration of docker")
	command.PersistentFlags().
		BoolVarP(&AutoMigrate, "auto-migrate", "m", os.Getenv("GOENV") == "DEV" || os.Getenv("GOENV") == "TEST", "auto migrate database if need")
}

var command = &cobra.Command{
	Use: ServiceName,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var run = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		Args = args
	},
}

var Args []string
var ServiceName string
var AutoMigrate bool

var envConfigPrefix string
var outputDockerConfig bool
var envVars conf.EnvVars

func SetServiceName(serviceName string) {
	command.Use = serviceName
	ServiceName = serviceName
}

func ConfP(c interface{}) {
	tpe := reflect.TypeOf(c)
	if tpe.Kind() != reflect.Ptr {
		panic(fmt.Errorf("ConfP pass ptr for setting value"))
	}

	os.Setenv(EnvVarKeyProjectName, ServiceName)

	envVars = conf.UnmarshalConf(c, envConfigPrefix)
	envVars.Print()
}

func Execute() {
	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if outputDockerConfig {
		dockerizier.Dockerize(envVars, ServiceName)
	}
}
