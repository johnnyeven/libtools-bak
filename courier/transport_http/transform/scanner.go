package transform

import (
	"fmt"
	"go/ast"
	"reflect"

	"golib/json"

	"golib/tools/reflectx"
	"golib/tools/strutil"
	"golib/tools/validate"
)

type Validator func() (isValid bool, msg string)

func NewScanner() *Scanner {
	return &Scanner{}
}

type Scanner struct {
	walker    PathWalker
	errMsgMap ErrMsgMap
}

func (vs *Scanner) Validate(rv reflect.Value, tpe reflect.Type) (bool, ErrMsgMap) {
	vs.marshalAndValidate(rv, tpe, "", false, "", "")
	return len(vs.errMsgMap) == 0, vs.errMsgMap
}

func (vs *Scanner) setErrMsg(path string, msg string) {
	if vs.errMsgMap == nil {
		vs.errMsgMap = ErrMsgMap{}
	}
	vs.errMsgMap[path] = msg
}

func (vs *Scanner) marshalAndValidate(rv reflect.Value, tpe reflect.Type, defaultValue string, required bool, tagValidate string, tagErrMsg string) {
	if _, ok := rv.Addr().Interface().(json.Unmarshaler); ok {
		errMsg := MarshalAndValidate(rv, tpe, defaultValue, required, tagValidate, tagErrMsg)
		if errMsg != "" {
			vs.setErrMsg(vs.walker.String(), errMsg)
		}
		return
	}

	tpe = reflectx.IndirectType(tpe)

	switch tpe.Kind() {
	case reflect.Struct:
		rv = reflectx.Indirect(rv)

		for i := 0; i < tpe.NumField(); i++ {
			field := tpe.Field(i)
			if !ast.IsExported(field.Name) {
				continue
			}

			jsonTag, exists := field.Tag.Lookup("json")
			if (exists && jsonTag != "-") || field.Anonymous {
				if !field.Anonymous {
					vs.walker.Enter(GetStructFieldDisplayName(&field))
				}

				tagValidate, _ := GetTagValidate(&field)
				defaultValue, notRequired := GetTagDefault(&field)
				tagErrMsg, _ := GetTagErrMsg(&field)

				vs.marshalAndValidate(rv.Field(i), field.Type, defaultValue, !notRequired, tagValidate, tagErrMsg)

				if rv.NumMethod() > 0 {
					validateHook := rv.MethodByName(fmt.Sprintf("Validate%s", field.Name))
					if validateHook.IsValid() {
						if validateFn, ok := validateHook.Interface().(func() string); ok {
							msg := validateFn()
							if msg != "" {
								vs.setErrMsg(vs.walker.String(), msg)
							}
						}
					}
				}

				if !field.Anonymous {
					vs.walker.Exit()
				}
			}
		}
	case reflect.Slice, reflect.Array:
		if tagValidate != "" {
			isValid, msg := validate.ValidateItem(tagValidate, rv.Interface(), tagErrMsg)
			if !isValid {
				vs.setErrMsg(vs.walker.String(), msg)
			}
		}
		for i := 0; i < rv.Len(); i++ {
			vs.walker.Enter(i)
			vs.marshalAndValidate(rv.Index(i), tpe.Elem(), "", false, "", tagErrMsg)
			vs.walker.Exit()
		}
	default:
		errMsg := MarshalAndValidate(rv, tpe, defaultValue, required, tagValidate, tagErrMsg)
		if errMsg != "" {
			vs.setErrMsg(vs.walker.String(), errMsg)
		}
	}
}

var ErrMsgForRequired = "缺失必填字段"

func MarshalAndValidate(rv reflect.Value, tpe reflect.Type, defaultValue string, isRequired bool, tagValidate string, tagErrMsg string) string {
	isPtr := rv.Kind() == reflect.Ptr
	if isPtr && rv.IsNil() {
		if isRequired {
			return ErrMsgForRequired
		}
		if defaultValue == "" {
			// key nil for patch check
			return ""
		}
		// when not required，should initial value
		if rv.CanSet() {
			rv.Set(reflect.New(reflectx.IndirectType(tpe)))
		}
	}

	rv = reflect.Indirect(rv)

	isEmptyValue := reflectx.IsEmptyValue(rv)

	if !isPtr && isEmptyValue && isRequired {
		return ErrMsgForRequired
	}

	// only empty value can set default
	if isEmptyValue && defaultValue != "" && rv.CanSet() {
		err := strutil.ConvertFromStr(defaultValue, rv)
		if err != nil {
			return fmt.Sprintf("%s can't set wrong default value %s", rv.Type().Name(), defaultValue)
		}
	}

	if tagValidate != "" {
		isValid, msg := validate.ValidateItem(tagValidate, rv.Interface(), tagErrMsg)
		if !isValid {
			return msg
		}
	}

	return ""
}

type PathWalker struct {
	path []interface{}
}

func (pw *PathWalker) Enter(i interface{}) {
	pw.path = append(pw.path, i)
}

func (pw *PathWalker) Exit() {
	pw.path = pw.path[:len(pw.path)-1]
}

func (pw *PathWalker) String() string {
	pathString := ""
	for i := 0; i < len(pw.path); i++ {
		switch pw.path[i].(type) {
		case string:
			if pathString != "" {
				pathString += "."
			}
			pathString += pw.path[i].(string)
		case int:
			pathString += fmt.Sprintf("[%d]", pw.path[i].(int))
		}
	}
	return pathString
}
