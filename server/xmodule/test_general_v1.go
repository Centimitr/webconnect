package xmodule

import (
	"github.com/Centimitr/xjsonbase"
	msg "github.com/Centimitr/xmessage"
	"github.com/Centimitr/xstatistics"
)

type Test struct {
}

func (t Test) PrintStat(c *msg.Ctx) {
	stat := c.Middleware["Statistics"].(xstatistics.Statistics)
	if c.Error.Fatal == nil {
		stat.Get()
	}
}

func (t Test) TestJSONBaseLoad(c *msg.Ctx) {
	j := c.Middleware["JSONBase"].(xjsonbase.JSONBase)
	if c.Error.Fatal == nil {
		j.Load(c)
	}
}

func (t Test) TestJSONBaseSave(c *msg.Ctx) {
	j := c.Middleware["JSONBase"].(xjsonbase.JSONBase)
	if c.Error.Fatal == nil {
		j.Save(c)
	}
}

func init() {
	var m Test
	msg.LoadModule(m)
}
