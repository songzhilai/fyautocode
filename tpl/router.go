package tpl

import (
	"reflect"
)

type TemplateBase interface {
	ReturnInfoStruct(string) map[string]interface{}
}

//TplTypeAssignment
func TplTypeAssignment(tplinter interface{}, params ...interface{}) map[string]interface{} {
	in := make([]reflect.Value, len(params))
	for k, v := range params {
		in[k] = reflect.ValueOf(v)
	}
	base, ok := tplinter.(TemplateBase)
	if !ok {
		return make(map[string]interface{})
	}
	return base.ReturnInfoStruct("123")
}
