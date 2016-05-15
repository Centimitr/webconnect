package xmessage

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
)

func do(ws *websocket.Conn, req *Req) {
	// format req
	// match
	// exec
	// return
	res := Res{Id: req.Id}
	resBytes, _ := json.Marshal(res)
	// res := `
	//        {
	//    		"id":"` + req.Id + `",
	//    		"params":{},
	//    		"data":{
	//    			"articles":[]
	//    		},
	//    		"error":{}
	//    	}`
	if err := websocket.Message.Send(ws, string(resBytes)); err != nil {
		fmt.Println("SEND ERROR.")
		return
	}
}
