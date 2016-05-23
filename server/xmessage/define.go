package xmessage

import (
	"encoding/json"
	"errors"
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
	res        *Res
	req        *Req
	reqParams  map[string]interface{}
	echoParams []string
	Method     string
	Middleware map[string]interface{}
	Params     map[string]string
	Data       string
	Error      CtxError
}
type ParamConfig struct {
	Key      string
	Required bool
	Default  string
	Echo     bool
	// Type     string
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
	init
*/

func (c *Ctx) Init() {
	c.Middleware = msg.Middleware.Map
	c.Params = make(map[string]string)
	c.reqParams = make(map[string]interface{})
	c.parseParams()
}
func (c *Ctx) parseParams() {
	s := c.req.Params
	err := json.Unmarshal([]byte(s), &c.reqParams)
	if err != nil {
		c.Error.NewFatal("Params parse error.")
	}
	// fmt.Println(c.reqParams)
}

/*
	context methods used in
*/

func (c *Ctx) getReqParamString(key string) string {
	switch c.reqParams[key].(type) {
	case string:
		return c.reqParams[key].(string)
	case float64:
		return fmt.Sprint(c.reqParams[key].(float64))
	default:
		c.Error.NewWarn(fmt.Sprint("Param type error, not a known type."))
		return fmt.Sprint(c.reqParams[key])
	}
}

func (c *Ctx) Set(data interface{}) {
	var setWitchConfig = func(p *ParamConfig) {
		switch {
		case p.Echo:
			c.echoParams = append(c.echoParams, p.Key)
			fallthrough
		case p.Default != "":
			c.Params[p.Key] = p.Default
		case p.Required:
			c.Error.NewFatal("Lack required param.")
		default:
			c.Params[p.Key] = c.getReqParamString(p.Key)
		}
	}
	switch d := data.(type) {
	case *ParamConfig:
		setWitchConfig(d)
	case []*ParamConfig:
		for _, c := range d {
			setWitchConfig(c)
		}
	default:
		c.Error.NewWarn("Error params to *Ctx.Set().")
	}
}

func (c *Ctx) Get(key string) string {
	return c.Params[key]
}

/*
	set res
*/
func (c *Ctx) setResParams() {
	var toEscaped = func(s string) string {
		return strings.Replace(s, `"`, `\"`, -1)
	}
	// var params []string
	// var stringMapMarshal = func(m map[string]string) string {
	// 	var kvs []string
	// 	for k, v := range m {
	// 		kvs = append(kvs, `"`+toEscaped(k)+`":"`+toEscaped(v)+`"`)
	// 	}
	// 	return "{" + strings.Join(kvs, ",") + "}"
	// }
	var stringMapPartlyMarshal = func(m map[string]string, keys []string) (string, error) {
		var kvs []string
		var err error
		for _, k := range keys {
			if v, ok := m[k]; ok {
				kvs = append(kvs, `"`+toEscaped(k)+`":"`+toEscaped(v)+`"`)
			} else {
				err = errors.New("Cannot find one given key in the map.")
			}
		}
		return "{" + strings.Join(kvs, ",") + "}", err
	}
	json, _ := stringMapPartlyMarshal(c.Params, c.echoParams)
	c.res.Params = json
}
