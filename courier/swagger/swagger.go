package swagger

import (
	"bytes"
	"context"
	"io/ioutil"

	"github.com/profzone/libtools/courier"
	"github.com/profzone/libtools/courier/httpx"
	"github.com/profzone/libtools/env"
)

func getSwaggerJSON() []byte {
	data, err := ioutil.ReadFile("./swagger.json")
	if err != nil {
		return data
	}
	return data
}

var SwaggerRouter = courier.NewRouter(Swagger{})

type Swagger struct {
	httpx.MethodGet
}

func (s Swagger) Output(c context.Context) (interface{}, error) {
	if !env.IsOnline() {
		json := &JSONBytes{}
		json.Write(getSwaggerJSON())
		return json, nil
	}
	return &JSONBytes{}, nil
}

// swagger:strfmt json
type JSONBytes struct {
	bytes.Buffer
}

func (JSONBytes) ContentType() string {
	return httpx.MIMEJSON
}
