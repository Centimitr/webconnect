package xmessage

import (
	"fmt"
	// "strings"
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
}
type CtxError struct {
	Warn  []string
	Fatal []string
}

func (c *CtxError) NewFatal(info string) {
	c.Fatal = append(c.Fatal, info)
}
func (c *Ctx) parseParams() {
	fmt.Println("ReqParamsString:", c.req.Params)
	// res.Params
	fmt.Println("ReqParams", c.reqParams)
}

func (c *Ctx) Set(p *ParamConfig) *Ctx {
	// if c.Params == nil {
	// 	c.parseParams()
	// }
	// if p.Required {
	// 	c.Error.NewFatal("Lack required param.")
	// 	return c
	// } else if c.reqParams[p.Key] != nil {
	// 	c.Params[p.Key] = c.reqParams[p.Key]
	// } else if p.Default != nil {
	// 	c.Params[p.Key] = p.Default
	// }
	switch {
	case c.Params == nil:
		c.parseParams()
		fallthrough
	case p.Default != "":
		c.Params[p.Key] = p.Default
	case p.Required:
		c.Error.NewFatal("Lack required param.")
	default:
		c.Params[p.Key] = c.reqParams[p.Key]
	}
	return c
}

// type Controller func(Req, *Res)

// type Router map[string]Controller
