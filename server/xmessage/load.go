package xmessage

import (
	"fmt"
	"reflect"
	// "strings"
)

var ProcessorTable map[string]*Processor

func init() {
	ProcessorTable = make(map[string]*Processor)
}

type Processor struct {
	Module string
	Name   string
	Func   func([]reflect.Value) []reflect.Value
}

func registerProcessor(p *Processor) {
	ProcessorTable[p.Module+"."+p.Name] = p
}

func Load(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Println("ADD:", t, t.Name())
	for i := 0; i < v.NumMethod(); i++ {
		fmt.Println("----", t.Method(i).Name)
		registerProcessor(&Processor{Module: t.PkgPath(), Name: t.Method(i).Name, Func: t.Method(i).Func.Call})
	}
}
