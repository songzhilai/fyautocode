package main

import (
	"fmt"
	"reflect"
)

type animal struct{}

func (a animal) PrintName(name string) {
	fmt.Println(name)
}

func main() {
	//使用反射的方式调用方法
	reflect.ValueOf(animal{}).MethodByName("PrintName").Call([]reflect.Value{reflect.ValueOf("123")})
}
