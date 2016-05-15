package main

import (
	msg "github.com/Centimitr/xmessage"
	// "golang.org/x/net/websocket"
	// "net/http"
	"fmt"
)

type Module struct{}

func (m Module) GetIndexArticles() {
	fmt.Println("A")
}
func (m Module) GetMessages() {
	fmt.Println("M")
}
func main() {
	// EXPERIMENT
	m := Module{}
	msg.Print(m)
	// SERVER
	// http.Handle("/echo", websocket.Handler(msg.Server))
	// err := http.ListenAndServe(":12345", nil)
	// if err != nil {
	// 	panic("ListenAndServe: " + err.Error())
	// }
}
