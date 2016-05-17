package xmessage

import (
	"fmt"
	// "strings"
	"encoding/json"
)

type Req struct {
	Id     string `json:"id"`
	Method string `json:"method"`
	Params string `json:"params"`
	Data   string `json:"data"`
}

type Res struct {
	Id     string `json:"id"`
	Method string `json:"method"`
	Params string `json:"params"`
	Data   string `json:"data"`
	Error  string `json:"error"`
}

type Ctx struct {
	res       *Res
	req       *Req
	reqParams map[string]interface{}
	Params    map[string]interface{}
	Data      string
	Error     CtxError
}
type ParamConfig struct {
	Key      string
	Required bool
	Default  interface{}
	// Type     string
	// Echo     bool
}

/*
	init
*/

func (c *Ctx) Init() {
	c.Params = make(map[string]interface{})
	c.reqParams = make(map[string]interface{})
	c.parseParams()
}
func (c *Ctx) parseParams() {
	// fmt.Println(c.req.Params)
	s := c.req.Params
	err := json.Unmarshal([]byte(s), &c.reqParams)
	if err != nil {
		c.Error.NewFatal("Params parse error.")
	}
	// fmt.Println(c.reqParams)
}

/*
	error
*/

type CtxError struct {
	Warn  []string
	Fatal []string
}

func (c *CtxError) NewFatal(info string) {
	c.Fatal = append(c.Fatal, info)
}
func (c *CtxError) NewWarn(info string) {
	c.Warn = append(c.Warn, info)
}

/*
	context methods used in
*/

func (c *Ctx) Set(p *ParamConfig) *Ctx {
	switch {
	case p.Default != nil && p.Default != "":
		c.Params[p.Key] = p.Default
	case p.Required:
		c.Error.NewFatal("Lack required param.")
	default:
		c.Params[p.Key] = c.reqParams[p.Key]
		// c.Params[p.Key] = "value"
	}
	return c
}

func (c *Ctx) GetNumber(key string) int {
	v, ok := c.Params[key].(float64)
	if !ok {
		c.Error.NewFatal(fmt.Sprint(`Param type error, "`, key, `"`, "expected to be int."))
	}
	return v
}

func (c *Ctx) GetString(key string) string {
	v, ok := c.Params[key].(string)
	if !ok {
		c.Error.NewFatal(fmt.Sprint(`Param type error, "`, key, `"`, "expected to be string."))
	}
	return v
}

func (c *Ctx) Get(key string) string {
	switch c.Params[key].(type) {
	case string:
		return c.Params[key].(string)
	case float64:
		return fmt.Sprint(c.Params[key].(float64))
	default:
		c.Error.NewWarn(fmt.Sprint("Param type error, not a known type."))
		return fmt.Sprint(c.Params[key])
	}
}
