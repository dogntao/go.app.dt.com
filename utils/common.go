package utils

import (
	"fmt"
	"reflect"
)

type StructMapInterface interface {
	Test()
}

func StructChangeToMap(s StructMapInterface) {
	fmt.Println(reflect.TypeOf(s).Elem())
	// switch reflect.TypeOf(s).Kind() {
	// case reflect.Struct:
	// 	fmt.Println("struct")
	rt := reflect.TypeOf(s).Elem()
	for i := 0; i < rt.NumField(); i++ {
		fmt.Println(rt.Field(i).Name)
	}
	// }
}
