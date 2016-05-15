package main

import (
	"fmt"
    "net/http"
    "golang.org/x/net/websocket"
    "encoding/json"
)
type Req struct{
	Id string `json:'id'`
	Method string `json:'method'`
	Params string `json:'params'`
	Data string `json:'data'`
}
type Res struct{
	Id string `json:'id'`
	Method string `json:'method'`
	Params string `json:'params'`
	Data string `json:'data'`
	Error string `json:'error'`
}
type Controller func(Req,*Res)
type Router map[string]Controller

func init(){
	r:=Router
	r.add('getIndexArticles',)
}




// Echo the data received on the WebSocket.
func Server(ws *websocket.Conn) {
    var err error
    for {
        var req string
        if err = websocket.Message.Receive(ws, &req); err != nil {
            fmt.Println("WAITING.")
            break
        }
        var r Req
        json.Unmarshal([]byte(req),&r)
        fmt.Println("RECEIVED.",r.Params)
        res := `
        {
    		"id":"`+r.Id+`",
    		"params":{},
    		"data":{
    			"articles":[]
    		},
    		"error":{}
    	}`
        fmt.Println("SEND.")

        if err = websocket.Message.Send(ws, res); err != nil {
            fmt.Println("SEND ERROR.")
            break
        }
    }
}

func main() {
    http.Handle("/echo", websocket.Handler(Server))
    err := http.ListenAndServe(":12345", nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
}