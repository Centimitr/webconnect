package main

import (

	// framework
	msg "github.com/Centimitr/xmessage"

	// middlewares
	_ "github.com/Centimitr/xapp"

	// processor modules
	_ "github.com/Centimitr/xstatistics"

	// system library
	"fmt"
	// "golang.org/x/net/websocket"
	// "net/http"
)

func main() {

	m := msg.Ins()

	fmt.Printf("%-40s %-10s %-10s %-10s %-10s %-10s\n", "Middleware", "AR", "BP", "AP", "BS", "AS")
	for _, m := range m.Middleware.Support {
		fmt.Printf("%-40s %-10v %-10v %-10v %-10v %-10v\n", m.Name, m.AfterReceive, m.BeforeProcess, m.AfterProcess, m.BeforeSend, m.AfterSend)
	}

	fmt.Printf("\n%-40s %-10s %-40s\n", "API", "Module", "Package Path")
	for _, proc := range m.ProcessorMap {
		fmt.Printf("%-40s %-10s %-30s\n", proc.Module+"."+proc.Name, proc.Module, proc.PkgPath)
	}

	// // SERVER
	// http.Handle("/echo", websocket.Handler(m.Server))
	// err := http.ListenAndServe(":12345", nil)
	// if err != nil {
	// 	panic("ListenAndServe: " + err.Error())
	// }
}
