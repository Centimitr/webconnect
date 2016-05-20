package main

import (
	_ "github.com/Centimitr/xapp"
	msg "github.com/Centimitr/xmessage"

	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

func main() {
	// fmt.Printf("%-40s %-10s %-30s\n", "AP", "Module", "Package Path")
	// for _, proc := range msg.Table {
	// 	fmt.Printf("%-40s %-10s %-30s\n", proc.Module+"."+proc.Name, proc.Module, proc.PkgPath)
	// }

	// SERVER
	m := msg.Ins()
	fmt.Printf("%-40s %-10s %-30s\n", "API", "Module", "Package Path")
	for _, proc := range m.ProcessorMap {
		fmt.Printf("%-40s %-10s %-30s\n", proc.Module+"."+proc.Name, proc.Module, proc.PkgPath)
	}

	http.Handle("/echo", websocket.Handler(m.Server))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
