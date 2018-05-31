package transform

import (
	"encoding"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"

	"golib/tools/env"
	"golib/tools/reflectx"
	"golib/tools/strutil"
)

type TagFlags map[string]bool

func GetTagValueAndFlags(tagString string) (v string, tagFlags TagFlags) {
	valueAndFlags := strings.Split(tagString, ",")
	v = valueAndFlags[0]
	tagFlags = TagFlags{}
	if len(valueAndFlags) > 1 {
		for _, flag := range valueAndFlags[1:] {
			tagFlags[flag] = true
		}
	}
	return
}

func GetStructFieldDisplayName(field *reflect.StructField) string {
	if jsonName, exists, _ := GetTagJSON(field); exists && jsonName != "" {
		return jsonName
	}
	return field.Name
}

func GetParameterDisplayName(field *reflect.StructField) string {
	if fieldName, exists, _ := GetTagName(field); exists && fieldName != "" {
		return fieldName
	}
	if jsonName, exists, _ := GetTagJSON(field); exists && jsonName != "" {
		if !env.IsOnline() {
			logrus.Warnf("%s `%s`, deprecated `json` tag for naming parameter, please use `name` tag instead", field.Name, field.Tag)
		}
		return jsonName
	}
	return field.Name
}

func GetTagJSON(field *reflect.StructField) (name string, exists bool, tagFlags TagFlags) {
	name, exists = field.Tag.Lookup("json")
	if exists {
		name, tagFlags = GetTagValueAndFlags(name)
	}
	return
}

func GetTagName(field *reflect.StructField) (name string, exists bool, tagFlags TagFlags) {
	name, exists = field.Tag.Lookup("name")
	if exists {
		name, tagFlags = GetTagValueAndFlags(name)
	}
	return
}

func GetTagStyle(field *reflect.StructField) (name string, exists bool, tagFlags TagFlags) {
	name, exists = field.Tag.Lookup("style")
	if exists {
		name, tagFlags = GetTagValueAndFlags(name)
	}
	return
}

func GetTagFmt(field *reflect.StructField) (name string, exists bool, tagFlags TagFlags) {
	name, exists = field.Tag.Lookup("fmt")
	if exists {
		name, tagFlags = GetTagValueAndFlags(name)
	}
	return
}

func GetTagDefault(field *reflect.StructField) (string, bool) {
	return field.Tag.Lookup("default")
}

func GetTagValidate(field *reflect.StructField) (string, bool) {
	return field.Tag.Lookup("validate")
}

func GetTagErrMsg(field *reflect.StructField) (string, bool) {
	return field.Tag.Lookup("errMsg")
}

func GetTagIn(field *reflect.StructField) (tagIn string, hasIn bool, tagInFlags map[string]bool) {
	tagIn, hasIn = field.Tag.Lookup("in")
	if hasIn {
		tagIn, tagInFlags = GetTagValueAndFlags(tagIn)
	}
	return
}

func MustStructValue(functionName string, rv reflect.Value) {
	if rv.Kind() != reflect.Struct {
		logrus.Panicf("%s args must a struct value, but got %+v %+v", functionName, rv.Type(), rv.Interface())
	}
}

func Stringify(v interface{}) ([]byte, error) {
	if marshaler, ok := v.(json.Marshaler); ok {
		return Unquote(marshaler.MarshalJSON())
	}
	if textMarshaler, ok := v.(encoding.TextMarshaler); ok {
		return textMarshaler.MarshalText()
	}
	s, err := strutil.ConvertToStr(v)
	if err != nil {
		return nil, err
	}
	return []byte(s), nil
}

func Unquote(data []byte, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	if len(data) > 2 {
		if string(data[0]) == `"` && string(data[len(data)-1]) == `"` {
			return data[1 : len(data)-1], nil
		}
	}
	return data, err
}

func ResolveCommaSplitValues(tpe reflect.Type, ss ...string) (values []string) {
	tpe = reflectx.IndirectType(tpe)
	if tpe.Kind() == reflect.Slice || tpe.Kind() == reflect.Array {
		itemTpe := reflectx.IndirectType(tpe.Elem())
		if !(itemTpe.Kind() == reflect.Struct || itemTpe.Kind() == reflect.Map) {
			for _, s := range ss {
				values = append(values, strings.Split(s, ",")...)
			}
			return
		}
	}
	return ss
}
