package xmessage

import (
	"fmt"
	"strings"
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
	reqParams map[string]string
	Params    map[string]string
	Data      string
	Error     CtxError
}
type ParamConfig struct {
	Key      string
	Required bool
	Default  string
	// Echo     bool
}

/*
	init
*/

func (c *Ctx) Init() {
	c.Params = make(map[string]string)
	c.reqParams = make(map[string]string)
	c.parseParams()
}
func (c *Ctx) parseParams() {
	fmt.Println(c.req.Params)
	s := c.req.Params
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "{")
	s = strings.TrimSuffix(s, "}")
	// kvs := strings.Split(s, ",")
	// fmt.Println(kvs)
	// fmt.Println("ReqParamsString:", c.req.Params)
	// res.Params
	// fmt.Println("ReqParams", c.reqParams)
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

/*
	context methods used in
*/

func (c *Ctx) Set(p *ParamConfig) *Ctx {
	switch {
	// case c.Params == nil:
	// 	c.parseParams()
	// 	fallthrough
	case p.Default != "":
		c.Params[p.Key] = p.Default
	case p.Required:
		c.Error.NewFatal("Lack required param.")
	default:
		c.Params[p.Key] = c.reqParams[p.Key]
	}
	return c
}
