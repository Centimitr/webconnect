package main

import (
	_ "github.com/Centimitr/xapp1"
	msg "github.com/Centimitr/xmessage"
	// "golang.org/x/net/websocket"
	// "net/http"
	"fmt"
)

func main() {
	// EXPERIMENT
	// m := Module{}
	// msg.Print(m)
	// msg.Load(m)
	fmt.Println(msg.ProcessorTable)
	// SERVER
	// http.Handle("/echo", websocket.Handler(msg.Server))
	// err := http.ListenAndServe(":12345", nil)
	// if err != nil {
	// 	panic("ListenAndServe: " + err.Error())
	// }
}
