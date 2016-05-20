package xmessage

import (
// "fmt"
// "reflect"
// "strings"
)

// var Table map[string]*Processor

// func init() {
// 	Table = make(map[string]*Processor)
// }

// // module packages use LoadModule to load itselft
// func LoadModule(x interface{}) {
// 	v := reflect.ValueOf(x)
// 	t := v.Type()
// 	for i := 0; i < v.NumMethod(); i++ {
// 		index := i
// 		// registerProcessor(&Processor{
// 		// 	PkgPath: t.PkgPath(),
// 		// 	Module:  t.Name(),
// 		// 	Name:    t.Method(i).Name,
// 		// 	Func: func(ctx *Ctx) []reflect.Value {
// 		// 		return t.Method(index).Func.Call([]reflect.Value{v, reflect.ValueOf(ctx)})
// 		// 	},
// 		// })
// 		msg.LoadModule(&Processor{
// 			PkgPath: t.PkgPath(),
// 			Module:  t.Name(),
// 			Name:    t.Method(i).Name,
// 			Func: func(ctx *Ctx) []reflect.Value {
// 				return t.Method(index).Func.Call([]reflect.Value{v, reflect.ValueOf(ctx)})
// 			},
// 		})
// 	}
// }

// // middleware packages use LoadMiddleware to load itself
// func LoadMiddleware(x interface{}) {
// 	v := reflect.ValueOf(x)
// 	t := v.Type()
// 	// fmt.Println("\nLOAD:", t.Name())
// 	for i := 0; i < v.NumMethod(); i++ {
// 		// fmt.Println("-----", t.Method(i).Name)
// 		// fmt.Println(t.Method(i).Name[0])
// 		index := i
// 		registerMiddleware(&Processor{
// 			PkgPath: t.PkgPath(),
// 			Module:  t.Name(),
// 			Name:    t.Method(i).Name,
// 			Func: func(ctx *Ctx) []reflect.Value {
// 				return t.Method(index).Func.Call([]reflect.Value{v, reflect.ValueOf(ctx)})
// 			},
// 		})
// 	}
// }
