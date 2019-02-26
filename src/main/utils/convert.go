package utils

import (
	"fmt"
	"reflect"
)

func InterfaceToStruct(itf interface{}) interface{} {
	getType := reflect.TypeOf(itf)
	fmt.Println("get Type is :", getType.Name())
	getValue := reflect.ValueOf(itf)
	fmt.Println("get all Fields is:", getValue)
	fmt.Println(reflect.TypeOf(itf))
	return itf
}
