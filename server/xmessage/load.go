package xmessage

import (
	"fmt"
	"reflect"
	"strings"
)

func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)
	fmt.Println(v.NumMethod())

	for i := 0; i < v.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
		t.Method(i).Func.Call([]reflect.Value{})
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name, strings.TrimPrefix(methType.String(), "func"))
	}
}
