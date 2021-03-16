package tpl

import (
	"fmt"
	"reflect"
)

//TplTypeAssignment
func TplTypeAssignment(tplinter interface{}, params ...interface{}) interface{} {
	in := make([]reflect.Value, len(params))
	for k, v := range params {
		in[k] = reflect.ValueOf(v)
	}
	switch tplinter.(type) {
	case GoModTemplate:
		a := tplinter.(GoModTemplate)
		reflect.ValueOf(&a).MethodByName("ReturnInfoStruct").Call(in)
		return a
	case GoSumTemplate:
		a := tplinter.(GoSumTemplate)
		reflect.ValueOf(&a).MethodByName("ReturnInfoStruct").Call(in)
		return a
	default:
		fmt.Println("mobanbucunzai ")
	}
	return tplinter
}
