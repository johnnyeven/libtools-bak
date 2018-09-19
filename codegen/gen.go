package codegen

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"golang.org/x/tools/imports"

	"github.com/profzone/libtools/duration"
	"github.com/profzone/libtools/format"
)

type Outputs map[string]string

func (outputs Outputs) Add(filename string, content string) {
	outputs[filename] = content
}

func (outputs Outputs) WriteFile(filename string, content string) {
	if IsGoFile(filename) {
		bytes, err := imports.Process(filename, []byte(content), nil)
		if err != nil {
			lines := strings.Split(content, "\n")
			lengthOfLines := len(lines)
			lengthOfLastLineString := len(fmt.Sprintf("%d", lengthOfLines+1))
			for i, line := range lines {
				lineString := fmt.Sprintf("%d", i+1)
				lineString = strings.Repeat(" ", lengthOfLastLineString-len(lineString)) + lineString

				fmt.Printf("%s: %s\n", lineString, line)
			}
			panic(err.Error())
		} else {
			bytes, err := format.Process(filename, bytes)
			if err != nil {
				panic(err.Error())
			}
			content = string(bytes)
		}
	}
	WriteFile(filename, content)
	delete(outputs, filename)
}

func (outputs Outputs) WriteFiles() {
	for filename, content := range outputs {
		outputs.WriteFile(filename, content)
	}
}

type Generator interface {
	Load(cwd string)
	Pick()
	Output(cwd string) Outputs
}

func Generate(generator Generator) {
	cost := duration.NewDuration()
	defer func() {
		cost.ToLogger().Infof("generate by %s done", reflect.TypeOf(generator).String())
	}()

	cwd, _ := os.Getwd()
	generator.Load(cwd)
	generator.Pick()
	outputs := generator.Output(cwd)
	outputs.WriteFiles()
}
