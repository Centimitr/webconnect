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
	// Params     map[string]interface{}
	Params map[string]string
	Data   string
	Error  CtxError
}
type ParamConfig struct {
	Key      string
	Required bool
	// Default  interface{}
	Default string
	// Type     string
	Echo bool
}

/*
	set res
*/
func (c *Ctx) setResParams() {
	var toEscaped = func(s string) string {
		return strings.Replace(s, `"`, `\"`, -1)
	}
	// var params []string
	var stringMapMarshal = func(m map[string]string) string {
		var kvs []string
		for k, v := range m {
			kvs = append(kvs, `"`+toEscaped(k)+`":"`+toEscaped(v)+`"`)
		}
		return "{" + strings.Join(kvs, ",") + "}"
	}
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
	fmt.Println(stringMapMarshal(c.Params))
	json, _ := stringMapPartlyMarshal(c.Params, c.echoParams)
	fmt.Println(json)
	c.res.Params = json
}

/*
	init
*/

func (c *Ctx) Init() {
	c.Params = make(map[string]string)
	// c.Params = make(map[string]interface{})
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

// func (c *Ctx) Set(p *ParamConfig) *Ctx {
// 	switch {
// 	case p.Echo:
// 		c.echoParams = append(c.echoParams, p.Key)
// 		fallthrough
// 	case p.Default != "":
// 		c.Params[p.Key] = p.Default
// 	case p.Required:
// 		c.Error.NewFatal("Lack required param.")
// 	default:
// 		c.Params[p.Key] = c.get(p.Key)
// 		// c.Params[p.Key] = "value"
// 	}
// 	return c
// }
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
			// c.Params[p.Key] = "value"
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

// func (c *Ctx) getNumber(key string) float64 {
// 	v, ok := c.Params[key].(float64)
// 	if !ok {
// 		c.Error.NewFatal(fmt.Sprint(`Param type error, "`, key, `"`, "expected to be float64."))
// 	}
// 	return v
// }

// func (c *Ctx) getString(key string) string {
// 	v, ok := c.Params[key].(string)
// 	if !ok {
// 		c.Error.NewFatal(fmt.Sprint(`Param type error, "`, key, `"`, "expected to be string."))
// 	}
// 	return v
// }

// func (c *Ctx) get(key string) string {
// 	switch c.Params[key].(type) {
// 	case string:
// 		return c.Params[key].(string)
// 	case float64:
// 		return fmt.Sprint(c.Params[key].(float64))
// 	default:
// 		c.Error.NewWarn(fmt.Sprint("Param type error, not a known type."))
// 		return fmt.Sprint(c.Params[key])
// 	}
// }
