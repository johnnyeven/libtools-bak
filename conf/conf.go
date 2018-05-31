package conf

import (
	"go/ast"
	"os"
	"reflect"
	"strings"

	"golib/tools/strutil"
)

func UnmarshalConf(c interface{}, prefix string) EnvVars {
	rv := reflect.Indirect(reflect.ValueOf(c))
	if !rv.CanSet() || rv.Type().Kind() != reflect.Struct {
		panic("UnmarshalConf need an variable which can set")
	}
	err := Unmarshal(rv, prefix)
	if err != nil {
		panic(err)
	}
	envVars, err := CollectEnvVars(rv, prefix)
	if err != nil {
		panic(err)
	}
	InitialRoot(rv)
	return envVars
}

type ICanInit interface {
	Init()
}

func InitialRoot(rv reflect.Value) {
	tpe := rv.Type()
	for i := 0; i < tpe.NumField(); i++ {
		value := rv.Field(i)
		if conf, ok := value.Interface().(ICanInit); ok {
			conf.Init()
		}
	}
}

// check and modify value
type IDefaultsMarshaller interface {
	MarshalDefaults(v interface{})
}

func Unmarshal(rv reflect.Value, envKey string) (err error) {
	rv = reflect.Indirect(rv)

	if rv.CanAddr() {
		v := rv.Addr().Interface()
		if defaultsMarshaller, ok := v.(IDefaultsMarshaller); ok {
			defaultsMarshaller.MarshalDefaults(v)
		}
	}

	switch rv.Kind() {
	case reflect.Func:
		// skip func
	case reflect.Struct:
		walkStructField(rv, true, func(fieldValue reflect.Value, field reflect.StructField) bool {
			nextEnvKey := resolveEnvVarKeyByField(envKey, field)
			err = Unmarshal(fieldValue, nextEnvKey)
			return err == nil
		})
	default:
		envValue, exist := os.LookupEnv(envKey)
		if exist {
			err = strutil.ConvertFromStr(envValue, rv)
			if err != nil {
				return
			}
		}
	}
	return nil
}

func walkStructField(rv reflect.Value, forSet bool, fn func(fieldValue reflect.Value, field reflect.StructField) bool) {
	reflectType := rv.Type()
	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		fieldValue := rv.Field(i)
		if !ast.IsExported(field.Name) {
			continue
		}

		if forSet && (!fieldValue.IsValid() || !fieldValue.CanSet()) {
			continue
		}

		next := fn(fieldValue, field)
		if !next {
			break
		}
	}
}

func resolveEnvVarKeyByField(pre string, field reflect.StructField) string {
	if field.Anonymous {
		return pre
	}
	if pre == "" {
		return strings.ToUpper(field.Name)
	}
	return strings.ToUpper(pre + "_" + field.Name)
}
