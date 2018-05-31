package httplib

import (
	"go/ast"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"golib/tools/courier/transport_http/transform"
	"golib/tools/reflectx"
)

type Request interface {
	Handle(c *gin.Context)
}

func TypeOfRequest(request Request) reflect.Type {
	requestInterface := reflect.Indirect(reflect.ValueOf(request)).Interface()
	return reflect.Indirect(reflect.ValueOf(requestInterface)).Type()
}

func TransToReq(c *gin.Context, reqModel interface{}) error {
	reflectValue := reflect.ValueOf(reqModel)
	return TransToReqByReflectValue(c, reflectValue)
}

func TransToReqByReflectValue(c *gin.Context, rv reflect.Value) error {
	group := transform.ParameterGroupFromReflectValue(rv)
	return transform.MarshalParameters(group, &ParameterValuesGetter{
		ParameterValuesGetter: transform.NewParameterValuesGetter(c.Request),
		Group: group,
		C:     c,
	})
}

type ParameterValuesGetter struct {
	*transform.ParameterValuesGetter
	Group *transform.ParameterGroup
	C     *gin.Context
}

func (getter *ParameterValuesGetter) Initial() error {
	if getter.Group.Context.Len() > 0 {
		for _, parameterMeta := range getter.Group.Context {
			contextValue, contextExists := getter.C.Get(parameterMeta.Name)
			if !contextExists {
				logrus.Panicf("read req %s context failed!", parameterMeta.Name)
			}
			reflectx.SetValue(parameterMeta.Value, parameterMeta.Field.Type, contextValue)
		}
	}
	return getter.ParameterValuesGetter.Initial()
}

func (getter *ParameterValuesGetter) Param(name string) string {
	return getter.C.Param(name)
}

func checkRequestTypeSettings(tpe reflect.Type) {
	for tpe.Kind() == reflect.Ptr {
		tpe = tpe.Elem()
	}

	for i := 0; i < tpe.NumField(); i++ {
		field := tpe.Field(i)
		if !ast.IsExported(field.Name) {
			continue
		}

		if field.Anonymous {
			checkRequestTypeSettings(field.Type)
			continue
		}

		_, exists := field.Tag.Lookup("in")
		if !exists {
			logrus.Panicf("missing tag `in` on struct field `%s.%s[%s]`", tpe.PkgPath(), tpe.Name(), field.Name)
		}
	}
}

func FromRequest(request Request) gin.HandlerFunc {
	requestType := TypeOfRequest(request)
	checkRequestTypeSettings(requestType)

	return func(c *gin.Context) {
		requestReflectValue := reflect.New(requestType)

		err := TransToReqByReflectValue(c, requestReflectValue)

		if err != nil {
			WriteError(c, err)
			c.Abort()
			return
		}

		requestReflectValue.Interface().(Request).Handle(c)
	}
}
