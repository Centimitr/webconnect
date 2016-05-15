package main

import (
	_ "github.com/Centimitr/xapp"
	msg "github.com/Centimitr/xmessage"

	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

func main() {
	fmt.Println()
	fmt.Println("TABL:")
	for _, proc := range msg.Table {
		fmt.Printf("%20s %10s %30s\n", proc.Name, proc.Module, proc.PkgPath)
	}
	// SERVER
	http.Handle("/echo", websocket.Handler(msg.Server))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
