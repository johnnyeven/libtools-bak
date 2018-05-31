package transport_http

import (
	"net/http"
	"reflect"

	"github.com/julienschmidt/httprouter"

	"golib/tools/courier"
	"golib/tools/courier/transport_http/transform"
	"golib/tools/reflectx"
)

func createHttpRequestDecoder(r *http.Request, params *httprouter.Params) courier.OperatorDecoder {
	return func(op courier.IOperator, rv reflect.Value) (err error) {
		if httpRequestTransformer, ok := op.(IHttpRequestTransformer); ok {
			httpRequestTransformer.TransformHttpRequest(r)
		}

		return transform.MarshalParameters(transform.ParameterGroupFromReflectValue(rv), &ParameterValuesGetter{
			ParameterValuesGetter: transform.NewParameterValuesGetter(r),
			Params:                params,
		})
	}
}

type ParameterValuesGetter struct {
	*transform.ParameterValuesGetter
	Params *httprouter.Params
}

func (getter *ParameterValuesGetter) Param(name string) string {
	return getter.Params.ByName(name)
}

func MarshalOperator(r *http.Request, operator courier.IOperator) (err error) {
	params := httprouter.Params{}

	if canPath, ok := (operator).(IPath); ok {
		params, err = GetParams(canPath.Path(), r.URL.Path)
	}

	opDecode := createHttpRequestDecoder(r, &params)
	op, err := courier.NewOperatorBy(reflectx.IndirectType(reflect.TypeOf(operator)), operator, opDecode)
	if err != nil {
		return err
	}

	reflect.Indirect(reflect.ValueOf(operator)).Set(reflect.ValueOf(op).Elem())
	return nil
}
