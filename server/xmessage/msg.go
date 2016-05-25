package xmessage

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"io/ioutil"
)

type middlewareConfig struct {
	Name  string `json:"name"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type config struct {
	Middleware []middlewareConfig `json:"middleware"`
}

type middlewareSupport struct {
	Name          string
	AfterReceive  bool
	BeforeProcess bool
	AfterProcess  bool
	BeforeSend    bool
	AfterSend     bool
}

type middleware struct {
	Map               map[string]interface{}
	Support           []middlewareSupport
	AfterReceiveFunc  []func(*Req)
	BeforeProcessFunc []func(*Ctx)
	AfterProcessFunc  []func(*Ctx)
	BeforeSendFunc    []func(*Res)
	AfterSendFunc     []func(*Res)
}

type Msg struct {
	Config       config
	Middleware   middleware
	ProcessorMap map[string]*Processor
}

// return a http.Handler
func (m *Msg) Server(ws *websocket.Conn) {
	var err error
	for {
		var req Req
		req.Init()
		if err = websocket.JSON.Receive(ws, &req); err != nil {
			break
		}
		go m.do(ws, &req)
	}
}

func (m *Msg) loadConfig() {
	var err error
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Read config file error:", err)
	}
	err = json.Unmarshal(data, &m.Config)
	if err != nil {
		fmt.Println("Unmarshal config file content error:", err)
	}
}

/*
	init
*/

var msg *Msg

func init() {
	msg = &Msg{
		Middleware: middleware{
			Map: make(map[string]interface{}),
		},
		ProcessorMap: make(map[string]*Processor),
	}
	msg.loadConfig()
}

/*
	global
*/

// return msg package's global object: msg
func Ins() *Msg {
	return msg
}

// used in Processor Module package to load itself
func LoadModule(x interface{}) {
	msg.loadModule(x)
}

// used in Middleware package to load itself
func LoadMiddleware(x interface{}) {
	msg.loadMiddleware(x)
}
